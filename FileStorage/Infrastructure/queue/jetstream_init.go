package queue

import (
	"log"

	"github.com/nats-io/nats.go"
)

func InitializeJetStream(natsURL string, streamName string, streamSubjects []string) (nats.JetStreamContext, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to nats")
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to JetStream")

	if _, err := js.StreamInfo(streamName); err != nil {
		log.Printf("Creating stream %s with streamSubjects: %v", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: streamSubjects,
		})
		if err != nil {
			return nil, err
		}
		log.Printf("Stream %s created", streamName)
	} else {
		log.Printf("Stream %s already exists", streamName)
	}
	return js, nil
}
