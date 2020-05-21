package functions

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"fmt"
	"time"
)

func PubsubStoreTranslationTask(ctx context.Context, message pubsub.Message) error {

	var translationTask *TranslationTask
	err := json.Unmarshal(message.Data, &translationTask)
	if err != nil {
		return fmt.Errorf("failed to unmarshal translationTask: %v", err)
	}

	bucketUrl := "hybrid-cloud-22365.appspot.com"

	//bucketUrl := os.Getenv("FIREBASE_BUCKET_URL")
	//if bucketUrl == "" {
	//	return fmt.Errorf("FIREBASE_BUCKET_URL not set\n")
	//}
	//fmt.Printf("FIREBASE_BUCKET_URL: %q\n", bucketUrl)

	config := &firebase.Config{
		StorageBucket: bucketUrl,
	}

	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		return fmt.Errorf("failed to create new firebase app: %v\n", err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return fmt.Errorf("failed to return storage instance: %v\n", err)
	}
	bucket, err := client.DefaultBucket()
	if err != nil {
		return fmt.Errorf("failed to return default bucket handle: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := bucket.Object(translationTask.TaskId).NewWriter(ctx)
	_, err = wc.Write(message.Data)
	if err != nil {
		return fmt.Errorf("failed to write to bucket: %v\n", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("failed to close bucket handle: %v\n", err)
	}

	return nil
}