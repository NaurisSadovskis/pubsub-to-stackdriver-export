package exporter

import (
	"cloud.google.com/go/logging"
	"context"
	"encoding/json"
	"log"
	"os"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// Exporter consumes a Pub/Sub message.
func Exporter(ctx context.Context, m PubSubMessage) error {
	client, err := logging.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	logName := os.Getenv("LOG_NAME")

	logger := client.Logger(logName)

	// Log textPayload
	if !isJSON(string(m.Data)) {
		logger.Log(logging.Entry{Payload: string(m.Data)})
		return nil
	}

	// Log jsonPayload
	logger.Log(logging.Entry{Payload: json.RawMessage([]byte(m.Data))})
	return nil
}
