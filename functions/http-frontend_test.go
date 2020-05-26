package functions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

//serviceUrl := os.Getenv("BOOTSTRAP_DATA_SERVER")
//if serviceUrl != "https://europe-west1-bootstrap-data-cloudfunctions.cloudfunctions.net" {
//	t.Errorf("BOOTSTRAP_DATA_SERVER environment variable expected %q set %q",
//		"https://europe-west1-bootstrap-data-cloudfunctions.cloudfunctions.net", serviceUrl)
//}

func TestHttpFrontendAvailable(t *testing.T) {

	serviceUrl := "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net"

	// Send ping request to service
	res, err := http.Post(serviceUrl+"/Translation",
		"application/x-www-form-urlencoded",
		strings.NewReader(""))
	if err != nil {
		t.Errorf("failed to send POST request to %q: %v\n", serviceUrl, err)
		return
	}

	err = res.Body.Close()
	if err != nil {
		t.Errorf("cannot close response body\n")
	}
}

func TestHttpFrontendPostData(t *testing.T) {

	serviceUrl := "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net"

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
		"Unknown source language": {
			request: Request{
				ClientVersion:  "0.0.1",
				ClientId:       "beab10c6-deee-4843-9757-719566214526",
				Text:           "Today is ascension of Jesus",
				SourceLanguage: "xx",
				TargetLanguage: "de",
			},
			response: Response{
				TranslatedText: "Heute ist die Himmelfahrt Jesu",
			},
			expectedError:        true,
			expectedStatusCode:   500,
			expectedStatusPrefix: "failed to translate text",
		},
		"Unknown target language": {
			request: Request{
				ClientVersion:  "0.0.1",
				ClientId:       "beab10c6-deee-4843-9757-719566214526",
				Text:           "Today is ascension of Jesus",
				SourceLanguage: "en",
				TargetLanguage: "xx",
			},
			response: Response{
				TranslatedText: "Heute ist die Himmelfahrt Jesu",
			},
			expectedError:        true,
			expectedStatusCode:   500,
			expectedStatusPrefix: "failed to translate text",
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

			if tc.expectedError {

				//fmt.Printf("Status: %v\n", res.Status)
				//fmt.Printf("Body: %s\n", body)

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

				if response.TaskId == "" {
					t.Errorf("TaskId in response is empty\n")
				}

				if response.TranslatedText == "" {
					t.Errorf("TranslatedText in response is empty\n")
				}

				if response.TranslatedText != tc.response.TranslatedText {
					t.Errorf("TranslatedText is %q and not as expected %q\n", response.TranslatedText, tc.response.TranslatedText)
				}

				if response.LoadCommands == nil {
					t.Errorf("LoadCommands in response is nil\n")
				}
			}

			err = res.Body.Close()
			if err != nil {
				t.Errorf("cannot close response body\n")
			}
		})
	}
}
