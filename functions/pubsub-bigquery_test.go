package functions

import (
	"cloud.google.com/go/bigquery"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestE2eBigQuery(t *testing.T) {

	serviceUrl := "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net"

	//taskId := "17d55af7-ceb4-4f4a-bfa0-ddcffb46fcde"
	//clientId := "beab10c6-deee-4843-9757-719566214526"

	projectID := "hybrid-cloud-22365"
	bqDataset := "migros_showcase"
	bqTable := "translations_v0_0_1"

	//bqLocation := os.Getenv("BQ_LOCATION")
	//if bqLocation == "" {
	//	return fmt.Errorf("BQ_LOCATION not set")
	//}
	////bqLocation = "US"

	ctx := context.Background()

	bqCredentialFile := "hybrid-cloud-22365-firebase-bq-22365.json"

	opt := option.WithCredentialsFile(bqCredentialFile)

	client, err := bigquery.NewClient(ctx, projectID, opt)
	if err != nil {
		fmt.Printf("failed to create bigquery client: %v", err)
		return
	}
	defer func() {
		err := client.Close()
		if err != nil {
			fmt.Printf("failed to close bigquery client: %v", err)
			return
		}
	}()

	var testCases map[string]struct {
		request Request
	}
	testCases = map[string]struct {
		request Request
	}{
		"BigQuery v0.0.1": {
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

			queryStr := "SELECT ClientVersion,  " +
				"ClientId, " +
				"TaskId , " +
				"Text, " +
				"SourceLanguage, " +
				"TargetLanguage, " +
				"TranslatedText FROM `" +
				projectID + "." + bqDataset + "." + bqTable + "` " +
				"WHERE clientId = \"" + tc.request.ClientId + "\" " +
				"AND taskId = \"" + response.TaskId + "\" " +
				"LIMIT 1"

			fmt.Printf("queryStr: %s\n", queryStr)

			time.Sleep(time.Second * 25)

			ctx, cancel := context.WithTimeout(ctx, time.Second*50)
			defer cancel()

			q := client.Query(queryStr)

			// Location must match that of the dataset(s) referenced in the query.
			q.Location = "US"

			// Run the query and print results when the query job is completed.
			job, err := q.Run(ctx)
			if err != nil {
				fmt.Printf("failed q.Run(ctx): %v", err)
				return
			}
			status, err := job.Wait(ctx)
			if err != nil {
				fmt.Printf("failed job.Wait(ctx): %v", err)
				return
			}
			if err := status.Err(); err != nil {
				fmt.Printf("status.Err(): %v", err)
				return
			}
			it, err := job.Read(ctx)

			var queryResult struct {
				ClientVersion  string `json:"clientVersion"`
				ClientId       string `json:"clientId"`
				TaskId         string `json:"taskId"`
				Text           string `json:"text"`
				SourceLanguage string `json:"sourceLanguage"`
				TargetLanguage string `json:"targetLanguage"`
				TranslatedText string `json:"translatedText"`
			}

			for {
				err := it.Next(&queryResult)
				if err == iterator.Done {
					break
				}
				if err != nil {
					fmt.Printf("it.Next(&row): %v", err)
					return
				}
			}
			//fmt.Printf("queryResult: %v\n", queryResult)

			if queryResult.ClientVersion != tc.request.ClientVersion {
				t.Errorf("ClientVersion is %q and not as expected %q\n", queryResult.ClientVersion, tc.request.ClientVersion)
			}
			if queryResult.ClientId != tc.request.ClientId {
				t.Errorf("ClientId is %q and not as expected %q\n", queryResult.ClientId, tc.request.ClientId)
			}
			if queryResult.Text != tc.request.Text {
				t.Errorf("Text is %q and not as expected %q\n", queryResult.Text, tc.request.Text)
			}
			if queryResult.SourceLanguage != tc.request.SourceLanguage {
				t.Errorf("SourceLanguage is %q and not as expected %q\n", queryResult.SourceLanguage, tc.request.SourceLanguage)
			}
			if queryResult.TargetLanguage != tc.request.TargetLanguage {
				t.Errorf("TargetLanguage is %q and not as expected %q\n", queryResult.TargetLanguage, tc.request.TargetLanguage)
			}

			if queryResult.TaskId != response.TaskId {
				t.Errorf("TaskId is %q and not as expected %q\n", queryResult.TaskId, response.TaskId)
			}

			if queryResult.TranslatedText != response.TranslatedText {
				t.Errorf("TranslatedText is %q and not as expected %q\n", queryResult.TranslatedText, response.TranslatedText)
			}
		})
	}
}
