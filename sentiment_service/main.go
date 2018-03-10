package main

import (
	"fmt"

	dc "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service/datacollection_client"
	_ "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service/sentiment"
)

func main() {

	err := dc.StreamTweets()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error streaming: %s", err))
	}
}
