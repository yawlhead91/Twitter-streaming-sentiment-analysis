package server

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"
	t "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitterapi_client"
)

const (
	streamMax = 1000
	streamMin = 10
)

var limit int32

// TwitterRouteServer Implements the generated
// TwitterRouteServer interface made in the proto file
type TwitterRouteServer struct{}

// GetTweets creates a stream of tweets for the given params
// to be searched for from the twitter api
func (s *TwitterRouteServer) GetTweets(params *pb.Params, stream pb.TwitterRoute_GetTweetsServer) error {

	errors := make(chan error)

	var streamcount int32
	// Here add or own feature to mutate
	// the stream from the twitter client
	limit = streamMin
	if params.Maxcount != 0 {
		if params.Maxcount >= streamMin && params.Maxcount <= streamMax {
			limit = params.Maxcount
		}
	}

	ts, err := t.GetStream(params)
	if err != nil {
		return err
	}

	go func() {
		<-errors
		log.Print("Closing twitter stream")
		ts.Stop()
	}()

	// Receives messages and type switches them to
	// call functions with typed messages. As we are
	// only interested in the tweets we only add the
	// 'Tweet' handler
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {

		// Extract the desired data
		r := &pb.Tweet{
			CreatedAt:    tweet.CreatedAt,
			RetweetCount: int64(tweet.RetweetCount),
			Text:         tweet.Text,
		}

		// Send the data back on the stream
		if err = stream.Send(r); err != nil {
			errors <- err
		}
	}

	// Receive messages from twitter stream, ensure
	// stream does not exceed stream allowance
	for message := range ts.Messages {
		if streamcount >= limit {
			errors <- nil
			return nil
		}
		demux.Handle(message)
		streamcount++
	}

	return nil
}
