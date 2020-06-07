package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	_, ok := os.LookupEnv("LAMBDA_TASK_ROOT")
	if !ok {
		log.Fatal("missing required env 'LAMBDA_TASK_ROOT'")
	}

	lambda.Start(func(in json.RawMessage) (json.RawMessage, error) {
		return nil, nil
	})
}
