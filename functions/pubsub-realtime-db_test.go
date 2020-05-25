package functions

import (
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestE2eRealtimeDb(t *testing.T) {

	serviceUrl := "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net"

	//taskId := "17d55af7-ceb4-4f4a-bfa0-ddcffb46fcde"
	//clientId := "beab10c6-deee-4843-9757-719566214526"

	//projectID := "hybrid-cloud-22365"
	databaseURL := "https://migros-showcase.firebaseio.com"
	databaseTable := "translations_v0_0_1"

	//databaseTable := os.Getenv("RTDB_TABLE")
	//if databaseTable == "" {
	//	t.Errorf("RTDB_TABLE not set")
	//	return
	//}
	//databaseTable := "translations_v_0_0_1"

	//"_RTDB_URL": "https://migros-showcase.firebaseio.com",
	//	"_RTDB_TABLE": "translations_v0_0_1"

	ctx := context.Background()

	//rtdbCredentialFile := "hybrid-cloud-22365-firebase-adminsdk-ca37q-d1e808e47b.json"

	//opt := option.WithCredentialsFile(rtdbCredentialFile)

	conf := &firebase.Config{
		DatabaseURL: databaseURL,
	}

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		t.Errorf("failed to create new app: %v", err)
		return

	}

	var testCases map[string]struct {
		request Request
	}
	testCases = map[string]struct {
		request Request
	}{
		"RealtimeDb v0.0.1": {
			request: Request{
				ClientVersion:  "0.0.1",
				ClientId:       "beab10c6-deee-4843-9757-719566214526",
				Text:           "Today is ascension of Jesus",
				SourceLanguage: "en",
				TargetLanguage: "de",
			},
		},
	}
	for n, tc := range testCases {

		t.Run(n, func(t *testing.T) {

			requestJson, err := json.Marshal(tc.request)
			if err != nil {
				t.Errorf("failed to marshal response: %v\n", err)
				return
			}
			//fmt.Printf("%s\n", requestJson)

			// Send request to service
			res, err := http.Post(serviceUrl+"/Translation",
				"application/json",
				strings.NewReader(string(requestJson)))
			if err != nil {
				t.Errorf("failed to send POST request to %q: %v\n", serviceUrl, err)
				return
			}
			body, err := ioutil.ReadAll(res.Body)
			//fmt.Printf("body: %v\n", string(body))

			var response *Response
			err = json.Unmarshal(body, &response)
			if err != nil {
				t.Errorf("failed to unmarshal response: %v\n", err)
				return
			}
			//fmt.Printf("TaskId: %s\n", response.TaskId)

			err = res.Body.Close()
			if err != nil {
				t.Errorf("cannot close response body\n")
			}

			time.Sleep(time.Second * 25)

			ctx, cancel := context.WithTimeout(ctx, time.Second*50)
			defer cancel()

			client, err := app.Database(ctx)
			if err != nil {
				t.Errorf("failed to create database client: %v", err)
				return

			}

			var translationTask TranslationTask
			ref := client.NewRef("/" + databaseTable + "/" + tc.request.ClientId + "/" + response.TaskId)
			err = ref.Get(ctx, interface{}(&translationTask))
			if err != nil {
				t.Errorf("failed to get database result: %v", err)
				return
			}

			//fmt.Printf("translationTask: %v\n", translationTask)

			if translationTask.ClientVersion != tc.request.ClientVersion {
				t.Errorf("ClientVersion is %q and not as expected %q\n", translationTask.ClientVersion, tc.request.ClientVersion)
			}
			if translationTask.ClientId != tc.request.ClientId {
				t.Errorf("ClientId is %q and not as expected %q\n", translationTask.ClientId, tc.request.ClientId)
			}
			if translationTask.Text != tc.request.Text {
				t.Errorf("Text is %q and not as expected %q\n", translationTask.Text, tc.request.Text)
			}
			if translationTask.SourceLanguage != tc.request.SourceLanguage {
				t.Errorf("SourceLanguage is %q and not as expected %q\n", translationTask.SourceLanguage, tc.request.SourceLanguage)
			}
			if translationTask.TargetLanguage != tc.request.TargetLanguage {
				t.Errorf("TargetLanguage is %q and not as expected %q\n", translationTask.TargetLanguage, tc.request.TargetLanguage)
			}

			if translationTask.TaskId != response.TaskId {
				t.Errorf("TaskId is %q and not as expected %q\n", translationTask.TaskId, response.TaskId)
			}

			if translationTask.TranslatedText != response.TranslatedText {
				t.Errorf("TranslatedText is %q and not as expected %q\n", translationTask.TranslatedText, response.TranslatedText)
			}
		})
	}
}
