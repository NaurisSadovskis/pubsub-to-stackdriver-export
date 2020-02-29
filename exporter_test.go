package exporter

import (
	"context"
	"testing"
)

func TestExporter(t *testing.T) {

	var testPayload PubSubMessage
	testPayload.Data = []byte(`"{"name":"John"}"`)

	err := Exporter(context.Background(), testPayload)
	if err != nil {
		t.Error(err)
	}

}
