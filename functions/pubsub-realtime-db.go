package functions

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"fmt"
	"os"
)

// PubsubRealtimeDbInsertTranslationTask stores the translation task in Realtime Database
func PubsubRealtimeDbInsertTranslationTask(ctx context.Context, message pubsub.Message) error {

	var translationTask TranslationTask
	err := json.Unmarshal(message.Data, &translationTask)
	if err != nil {
		return fmt.Errorf("failed to unmarshal translationTask: %v", err)
	}

	for i, val := range os.Environ() {
		fmt.Printf("%v: %q\n", i, val)
	}

	databaseURL := os.Getenv("RTDB_URL")
	if databaseURL == "" {
		return fmt.Errorf("RTDB_URL not set")
	}
	//databaseURL := "https://migros-showcase.firebaseio.com"

	conf := &firebase.Config{
		DatabaseURL: databaseURL,
	}

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return fmt.Errorf("failed to create new app: %v", err)

	}

	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("failed to create database client: %v", err)

	}

	databaseTable := os.Getenv("RTDB_TABLE")
	if databaseTable == "" {
		return fmt.Errorf("RTDB_TABLE not set")
	}
	//databaseTable := "translations_v_0_0_1"

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("/" + databaseTable + "/" + translationTask.ClientId + "/" + translationTask.TaskId)
	err = ref.Set(ctx, interface{}(&translationTask))
	if err != nil {
		return fmt.Errorf("failed to push list node: %v", err)

	}
	return nil
}
