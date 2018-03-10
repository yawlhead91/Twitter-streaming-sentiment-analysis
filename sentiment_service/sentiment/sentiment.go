package sentiment

import (
	"fmt"
	"log"

	"github.com/cdipaolo/sentiment"
	pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"
)

var model sentiment.Models

func init() {
	model, err := sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}
	log.Printf("Model init: %v", model)
}

// Run sentiment summary  on the  given tweet, the
// package is trained on IMDB movies text data so
// the accuracy is null but gives a rough estimate
func Run(tweet *pb.Tweet) error {
	analysis := model.SentimentAnalysis(tweet.Text, sentiment.English)
	log.Printf("Analysis: %d", analysis.Score)
	return nil
}
