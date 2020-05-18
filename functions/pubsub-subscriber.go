package functions

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
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
