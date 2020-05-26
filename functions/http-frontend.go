package functions

import (
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/translate"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/text/language"
	"net/http"
	"os"
)

// TranslationHTTP is an entry point for the smbe
func TranslationHTTP(w http.ResponseWriter, r *http.Request) {

	gcpProject := os.Getenv("GCP_PROJECT")
	if gcpProject == "" {
		http.Error(w, fmt.Sprintf("GCP_PROJECT not set\n"), http.StatusInternalServerError)
		return

	}

	pubsubTopic := os.Getenv("PUBSUB_TOPIC_TRUNC")
	if pubsubTopic == "" {
		http.Error(w, fmt.Sprintf("PUBSUB_TOPIC_TRUNC not set\n"), http.StatusInternalServerError)
		return
	}

	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to decode request: %v\n", err), http.StatusInternalServerError)
			return
		}
	}

	pubsubTopicVersion := fmt.Sprintf("%s_%s", pubsubTopic, request.ClientVersion)

	taskId, err := uuid.NewRandom()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create new random UUID: %v\n", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("TaskId: %s\n", taskId)

	//traceInfoMap := make(map[string]string)
	//traceInfoMap["source"] = "http-frontend.go"
	//traceInfoMap["target"] = pubsubTopicVersion
	//
	//traceInfoSlice := make([]map[string]string, 0)
	//traceInfoSlice = append(traceInfoSlice, traceInfoMap)

	translationTask := TranslationTask{
		ClientVersion:  request.ClientVersion,
		ClientId:       request.ClientId,
		TaskId:         taskId.String(),
		Text:           request.Text,
		SourceLanguage: request.SourceLanguage,
		TargetLanguage: request.TargetLanguage,
		TranslatedText: "none",
		//TraceInfo:      traceInfoSlice,
	}

	ctx := context.Background()
	translateClient, err := translate.NewClient(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create new translate client: %v\n", err), http.StatusInternalServerError)
		return
	}

	// Use the client.
	ds, err := translateClient.DetectLanguage(ctx, []string{translationTask.Text})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to detect language: %v\n", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(ds)

	if ds[0][0].Language.String() != translationTask.SourceLanguage {
		http.Error(w, fmt.Sprintf("source language is %q, but expected is %q\n", ds[0][0].Language.String(), translationTask.SourceLanguage),
			http.StatusInternalServerError)
		return
	}

	if ds[0][0].Confidence < 0.9 {
		http.Error(w, fmt.Sprintf("source language detection's confidence for %q is below 90%s\n", ds[0][0].Language.String(), "%"),
			http.StatusInternalServerError)
		return
	}

	langs, err := translateClient.SupportedLanguages(ctx, language.English)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to retrieve supported languages: %v\n", err), http.StatusInternalServerError)
		return
	}
	//fmt.Println(langs)

	var targetTag language.Tag
	for _, lang := range langs {
		if lang.Tag.String() == translationTask.TargetLanguage {
			targetTag = lang.Tag
		}
	}

	translations, err := translateClient.Translate(ctx,
		[]string{translationTask.Text}, targetTag,
		&translate.Options{
			Source: ds[0][0].Language,
			Format: translate.Text,
		})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to translate text: %v\n", err), http.StatusInternalServerError)
		return
	}
	//fmt.Println(translations[0].Text)

	loadCommands := make([]string, 0)
	loadCommands = append(loadCommands, fmt.Sprintf("gsutil cat gs://hybrid-cloud-22365.appspot.com/%s/%s/%s | jq",
		translationTask.ClientVersion, translationTask.ClientId, taskId),
		fmt.Sprintf("bq query '%s'",
			fmt.Sprintf("SELECT * FROM migros_showcase.translations_v0_0_1 WHERE taskId = %q"), taskId),
		fmt.Sprintf("firebase database:get --pretty --instance %s --project %s /translations_v0_0_1/%s/%s",
			"migros-showcase", "hybrid-cloud-22365", translationTask.ClientId, taskId))

	// gcloud logging read 'resource.type="cloud_function" resource.labels.function_name="Translation" resource.labels.region="europe-west1"
	//textPayload="taskId: %q"    TaskId: '

	response := Response{
		TaskId:         taskId.String(),
		TranslatedText: translations[0].Text,
		LoadCommands:   loadCommands,
	}

	translationTask.TranslatedText = translations[0].Text

	pubsubClient, err := pubsub.NewClient(ctx, gcpProject)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create new pubsub client: %v\n", err), http.StatusInternalServerError)
		return
	}

	var translationJson []byte
	translationJson, err = json.Marshal(translationTask)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal translationTask: %v\n", err), http.StatusInternalServerError)
		return
	}

	topic := pubsubClient.Topic(pubsubTopicVersion)
	defer topic.Stop()
	var results []*pubsub.PublishResult
	res := topic.Publish(ctx, &pubsub.Message{
		Data: translationJson,
	})
	results = append(results, res)
	// Do other work ...
	for _, r := range results {
		messageId, err := r.Get(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get pubsub result: %v\n", err), http.StatusInternalServerError)
			return
		}
		_ = messageId // future use?

		responseJson, err := json.MarshalIndent(response, "", " ")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal response: %v\n", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s\n", responseJson)
	}
}
