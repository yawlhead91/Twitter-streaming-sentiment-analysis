package main

import (
	"fmt"
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

	session, err := r.CreateSession(datastoreAddr)
	if err != nil {
		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		panic(fmt.Errorf("Could not connect to datastore with host %s - %v", datastoreAddr, err))
	}

	err = dc.StreamTweets(session.Clone())
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error streaming: %s", err))
	}

	err = dc.StreamRss(session.Clone())
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error streaming: %s", err))
	}
}
