package hangouts

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
)

type Lambda struct {
	sess *session.Session
}

func Make(sess *session.Session) Lambda {
	return Lambda{
		sess: sess,
	}
}

func (Lambda) Handler(update Update) (SimpleMessage, error) {
	log.Printf("%+v", update)

	return SimpleMessage{
		Text: "pong",
	}, nil
}
