package main

import (
	"fmt"
	"log"
	"os"

	dc "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service/datacollectionclient"
	r "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service/repository"
	_ "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service/sentiment"
)

var datastoreAddr = "127.0.0.1:27017"

func main() {

	host := os.Getenv("datastoreAddr")
	if host != "" {
		datastoreAddr = host
	}

	log.Print("Host:", host)

	session, err := r.CreateSession(datastoreAddr)
	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	err = dc.StreamTweets(session.Clone())
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error streaming: %s", err))
	}
}
