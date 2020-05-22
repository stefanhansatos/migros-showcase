package functions

import (
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestE2eStorage(t *testing.T) {

	serviceUrl := "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net"

	//bucketUrl := os.Getenv("FIREBASE_BUCKET_URL")
	//if bucketUrl == "" {
	//	fmt.Errorf("FIREBASE_BUCKET_URL not set\n")
	//}
	//fmt.Printf("FIREBASE_BUCKET_URL: %q\n", bucketUrl)

	bucketUrl := "hybrid-cloud-22365.appspot.com"

	config := &firebase.Config{
		StorageBucket: bucketUrl,
	}

	storageCredentialFile := "hybrid-cloud-22365-firebase-storage-22365.json"

	opt := option.WithCredentialsFile(storageCredentialFile)
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		t.Fatalf("failed to create new firebase app: %v\n", err)
	}

	ctx := context.Background()
	client, err := app.Storage(ctx)
	if err != nil {
		t.Fatalf("failed to return storage instance: %v\n", err)
	}
	bucket, err := client.DefaultBucket()
	if err != nil {
		t.Fatalf("failed to return default bucket handle: %v\n", err)
	}

	var testCases map[string]struct {
		request Request
	}
	testCases = map[string]struct {
		request Request
	}{
		"Storage v0.0.1": {
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
			//fmt.Printf("gsutil cat gs://%s/%s/%s\n", bucketUrl, tc.request.ClientVersion, response.TaskId)

			err = res.Body.Close()
			if err != nil {
				t.Errorf("cannot close response body\n")
			}

			time.Sleep(time.Second * 2)

			ctx, cancel := context.WithTimeout(ctx, time.Second*50)
			defer cancel()

			var translationTaskJson []byte
			rc, err := bucket.Object(tc.request.ClientVersion + "/" + response.TaskId).NewReader(ctx)
			if err != nil {
				t.Errorf("failed to create reader %q\n", err)
				return
			} else {
				translationTaskJson, err = ioutil.ReadAll(rc)
				if err != nil {
					t.Errorf("failed to read %q\n", err)
					return
				}
			}
			//fmt.Printf("translationTaskJson: %s\n", translationTaskJson)

			err = rc.Close()
			if err != nil {
				t.Errorf("cannot close reader\n")
			}

			var translationTask TranslationTask
			err = json.Unmarshal(translationTaskJson, &translationTask)
			if err != nil {
				t.Errorf("failed to unmarshal translationTask: %v\n", err)
				return
			}

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
