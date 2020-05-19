package functions

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func PubsubTranslationTaskReceiver(ctx context.Context, message pubsub.Message) error {

	var translationTask *TranslationTask
	err := json.Unmarshal(message.Data, &translationTask)
	if err != nil {
		return fmt.Errorf("failed to unmarshal translationTask: %v", err)
	}
	translationTaskJson, err := json.Marshal(translationTask)
	if err != nil {
		return fmt.Errorf("failed to marshal translationTaskJson: %v", err)
	}
	fmt.Printf("%s\n", translationTaskJson)

	return nil
}

func PubsubTranslationTaskConverter(ctx context.Context, message pubsub.Message) error {

	var translationTask *TranslationTask
	err := json.Unmarshal(message.Data, &translationTask)
	if err != nil {
		return fmt.Errorf("failed to unmarshal translationTask: %v", err)
	}

	//translationTaskJson, err := json.Marshal(translationTask)
	//if err != nil {
	//	return fmt.Errorf("failed to marshal translationTaskJson: %v", err)
	//}
	//fmt.Printf("%s\n", translationTaskJson)

	translationQuery := TranslationQuery{
		Uuid:           translationTask.TaskId,
		Text:           translationTask.Text,
		SourceLanguage: translationTask.SourceLanguage,
		TargetLanguage: translationTask.TargetLanguage,
	}
	fmt.Printf("translationQuery: %v\n", translationQuery)

	wrappedData := WrappedData{
		Source: "projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslate-europe-west1-smbe_input",
		Translation: &Translation{
			TranslationQuery:  &translationQuery,
			TranslatedText:    translationTask.TranslatedText,
			TranslationErrors: []string{""},
		},
		LogFilter: "xxx",
		Timestamp: time.Now().String(),
		Unix:      time.Now().Unix(),
	}
	fmt.Printf("wrappedData: %v\n", wrappedData)

	var translationJson []byte
	translationJson, err = json.Marshal(wrappedData)
	if err != nil {
		return fmt.Errorf("failed to marshal wrappedData: %v\n", err)
	}
	fmt.Printf("translationJson: %s\n", translationJson)

	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		return fmt.Errorf("GCP_PROJECT not set.\n")
	}

	// Fetch the service account key JSON file contents
	//opt := option.WithCredentialsFile("./hybrid-cloud-22365-smbe-22365.json")

	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("failed to create new pubsub client: %v\n", err)
	}

	attributes := make(map[string]string)
	attributes["uuid"] = translationQuery.Uuid

	topic := pubsubClient.Topic("smbe_output")
	defer topic.Stop()
	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data:       translationJson,
		Attributes: attributes,
	})
	results = append(results, r)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			return fmt.Errorf("failed to get pubsub result: %v\n", err)
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}

	return nil
}

//
//// SmbeTranslationQueryLoad stores the translation query in smbe:translation-queries to store the pubsub message
//func SmbeTranslate(ctx context.Context, message pubsub.Message) error {
//
//	// resource.type="cloud_function" resource.labels.function_name="SmbeTranslate" resource.labels.region="europe-west1" severity=DEFAULT
//
//	var translationQuery *TranslationQuery
//	err := json.Unmarshal(message.Data, &translationQuery)
//
//	wrappedData := WrappedData{
//		Source: "projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslate-europe-west1-smbe_input",
//		Translation: &Translation{
//			TranslationQuery:  translationQuery,
//			TranslatedText:    "",
//			TranslationErrors: []string{""},
//		},
//		LogFilter: "xxx",
//		Timestamp: time.Now().String(),
//		Unix:      time.Now().Unix(),
//	}
//	//
//	//databaseURL := os.Getenv("RTDB_URL")
//	//if databaseURL == "" {
//	//	return fmt.Errorf("RTDB_URL not set")
//	//}
//
//	translateClient, err := translate.NewClient(ctx)
//	if err != nil {
//		return fmt.Errorf("failed to create new client: %v\n", err)
//	}
//
//	// Use the client.
//	ds, err := translateClient.DetectLanguage(ctx, []string{translationQuery.Text})
//	if err != nil {
//		return fmt.Errorf("failed to detect language: %v\n", err)
//	}
//	fmt.Println(ds)
//
//	if ds[0][0].Language.String() != translationQuery.SourceLanguage {
//		return fmt.Errorf("source language is %q, but expected is %q\n", ds[0][0].Language.String(), translationQuery.SourceLanguage)
//	}
//
//	if ds[0][0].Confidence < 0.9 {
//		return fmt.Errorf("source language detection's confidence for %q is below 90%\n", ds[0][0].Language.String())
//	}
//
//	langs, err := translateClient.SupportedLanguages(ctx, language.English)
//	if err != nil {
//		return fmt.Errorf("failed to retrieve supported languages: %v\n", err)
//	}
//	//fmt.Println(langs)
//
//	var targetTag language.Tag
//	for _, lang := range langs {
//		if lang.Tag.String() == translationQuery.TargetLanguage {
//			targetTag = lang.Tag
//		}
//	}
//
//	translations, err := translateClient.Translate(ctx,
//		[]string{translationQuery.Text}, targetTag,
//		&translate.Options{
//			Source: ds[0][0].Language,
//			Format: translate.Text,
//		})
//	if err != nil {
//		return fmt.Errorf("failed to translate text: %v\n", err)
//	}
//	fmt.Println(translations[0].Text)
//
//	wrappedData.Translation.TranslatedText = translations[0].Text
//
//	var translationJson []byte
//	translationJson, err = json.Marshal(wrappedData)
//	if err != nil {
//		return fmt.Errorf("failed to marshal wrappedData: %v\n", err)
//	}
//
//	fmt.Printf("translationJson: %s\n", translationJson)
//
//	// Close the client when finished.
//	if err := translateClient.Close(); err != nil {
//		return fmt.Errorf("failed to close client: %v\n", err)
//	}
//
//	projectID := os.Getenv("GCP_PROJECT")
//	if projectID == "" {
//		return fmt.Errorf("GCP_PROJECT not set.\n")
//	}
//
//	//pubsubCredentialFile := os.Getenv("SMBE_APPLICATION_CREDENTIALS")
//	//if pubsubCredentialFile == "" {
//	//	return fmt.Errorf("SMBE_APPLICATION_CREDENTIALS not set\n")
//	//}
//
//	// Fetch the service account key JSON file contents
//	opt := option.WithCredentialsFile("./hybrid-cloud-22365-smbe-22365.json")
//
//	pubsubClient, err := pubsub.NewClient(ctx, projectID, opt)
//	if err != nil {
//		return fmt.Errorf("failed to create new pubsub client: %v\n", err)
//	}
//
//	attributes := make(map[string]string)
//	attributes["uuid"] = translationQuery.Uuid
//
//	topic := pubsubClient.Topic("smbe_output")
//	defer topic.Stop()
//	var results []*pubsub.PublishResult
//	r := topic.Publish(ctx, &pubsub.Message{
//		Data:       translationJson,
//		Attributes: attributes,
//	})
//	results = append(results, r)
//	// Do other work ...
//	for _, r := range results {
//		id, err := r.Get(ctx)
//		if err != nil {
//			return fmt.Errorf("failed to get pubsub result: %v\n", err)
//		}
//		fmt.Printf("Published a message with a message ID: %s\n", id)
//	}
//
//	return nil
//}
