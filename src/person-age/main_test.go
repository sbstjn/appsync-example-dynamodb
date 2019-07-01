package main

import (
	"testing"
	"time"
)

func TestPersonAgeResolver(t *testing.T) {
	data, err := handle(personAgePayload{
		Birthday: time.Date(1987, time.August, 28, 0, 0, 0, 0, time.UTC),
	})

	if err != nil {
		t.Errorf("Error must be nil, got: %v", err)

		return
	}

	if data < 31 {
		t.Errorf("Must return age of at least 31, got: %d", data)
	}
}
