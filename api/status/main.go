package main

import (
    "github.com/aws/aws-lambda-go/lambda"
)

func status() (string, error) {
    return "status good", nil
}

func main() {
    lambda.Start(status)
}