package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type personAgePayload struct {
	Birthday time.Time `json:"birthday"`
}

func handle(payload personAgePayload) (int, error) {
	return int(time.Since(payload.Birthday) / time.Hour / 24 / 365), nil
}

func main() {
	lambda.Start(handle)
}
