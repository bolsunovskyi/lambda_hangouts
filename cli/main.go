package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"github.com/bolsunovskyi/lambda_hangouts_bot"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err)
	}

	lmb := hangouts.Make(sess)

	lambda.Start(lmb.Handler)
}