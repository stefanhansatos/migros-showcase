package functions

import (
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"fmt"
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

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	var testCases map[string]struct {
		request              Request
		response             Response
		expectedError        bool
		expectedStatusCode   int
		expectedStatusPrefix string
	}
	testCases = map[string]struct {
		request              Request
		response             Response
		expectedError        bool
		expectedStatusCode   int
		expectedStatusPrefix string
	}{
		"Standard": {
			request: Request{
				ClientVersion:  "0.0.1",
				ClientId:       "beab10c6-deee-4843-9757-719566214526",
				Text:           "Today is ascension of Jesus",
				SourceLanguage: "en",
				TargetLanguage: "de",
			},
			response: Response{
				TranslatedText: "Heute ist die Himmelfahrt Jesu",
			},
			expectedError: false,
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
			fmt.Printf("body: %v\n", string(body))

			if tc.expectedError {

				fmt.Printf("Status: %v\n", res.Status)
				fmt.Printf("Body: %s\n", body)

				if res.StatusCode != tc.expectedStatusCode {
					t.Errorf("status code is %v and not as expected %v\n", res.StatusCode, tc.expectedStatusCode)
					return
				}
				if strings.HasPrefix(res.Status, tc.expectedStatusPrefix) {
					t.Errorf("status text is %v and has not as expected the prefix %v\n", res.StatusCode, tc.expectedStatusCode)
					return
				}
			} else {

				var response *Response
				err = json.Unmarshal(body, &response)
				if err != nil {
					t.Errorf("failed to unmarshal response: %v\n", err)
					return
				}

				fmt.Printf("TaskId: %s\n", response.TaskId)

				if response.TaskId == "" {
					t.Errorf("TaskId in response is empty\n")
				}

				if response.TranslatedText == "" {
					t.Errorf("TranslatedText in response is empty\n")
				}

				if response.TranslatedText != tc.response.TranslatedText {
					t.Errorf("TranslatedText is %q and not as expected %q\n", response.TranslatedText, tc.response.TranslatedText)
				}

				if response.LoadCommand == "" {
					t.Errorf("LoadCommand in response is empty\n")
				}

				//response.TaskId = "d09870fc-8e3e-4ec6-9808-501155e94915"
				time.Sleep(time.Second * 2)

				var translationTaskJson []byte
				rc, err := bucket.Object(tc.request.ClientVersion + "/" + response.TaskId).NewReader(ctx)
				if err != nil {
					t.Errorf("failed to create reader %q\n", err)
					return
				} else {
					_, err = rc.Read(translationTaskJson)
					if err != nil {
						t.Errorf("failed to read %q\n", err)
						return
					}
				}

				fmt.Printf("translationTaskJson: %s\n", translationTaskJson)
				fmt.Printf("translationTaskJson: %s\n", &translationTaskJson)

				var translationTask TranslationTask
				err = json.Unmarshal(translationTaskJson, &translationTask)
				if err != nil {
					t.Errorf("failed to unmarshal translationTask: %v\n", err)
					return
				}
				fmt.Printf("translationTask.TaskId: %v\n", translationTask.TaskId)

			}

			err = res.Body.Close()
			if err != nil {
				t.Errorf("cannot close response body\n")
			}
		})
	}
}
