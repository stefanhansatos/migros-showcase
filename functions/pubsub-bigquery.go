package functions

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// PubsubBqPutTranslationTask stores the translation task in BigQuery
func PubsubBqPutTranslationTask(ctx context.Context, message pubsub.Message) error {

	var translationTask TranslationTask
	err := json.Unmarshal(message.Data, &translationTask)
	if err != nil {
		return fmt.Errorf("failed to unmarshal translationTask: %v", err)
	}

	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		return fmt.Errorf("GCP_PROJECT not set")
	}
	//projectID := "hybrid-cloud-22365"

	bqDataset := os.Getenv("BQ_DATASET")
	if projectID == "" {
		return fmt.Errorf("BQ_DATASET not set")
	}
	//bqDataset := "migros-showcase"

	bqTable := os.Getenv("BQ_TABLE")
	if projectID == "" {
		return fmt.Errorf("BQ_TABLE not set")
	}
	//bqTable := "translations_v0_0_1"

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("failed to create bigquery client: %v", err)
	}
	defer client.Close()

	translationTasks := []TranslationTask{
		translationTask,
	}

	smbeDataset := client.Dataset(bqDataset)
	translationsTable := smbeDataset.Table(bqTable)

	inserter := translationsTable.Inserter()
	err = inserter.Put(ctx, translationTasks)
	if err != nil {
		return fmt.Errorf("failed to insert data into bigquery: %v", err)
	}
	return nil
}
