package server

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"
	t "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitterapi_client"
)

// TwitterRouteServer Implements the generated
// TwitterRouteServer interface made in the proto file
type TwitterRouteServer struct{}

// GetTweets creates a stream of tweets for the given params
// to be searched for from the twitter api
func (s *TwitterRouteServer) GetTweets(params *pb.Params, stream pb.TwitterRoute_GetTweetsServer) error {
	ts, err := t.GetStream(params)
	if err != nil {
		return err
	}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		r := &pb.Tweet{
			CreatedAt:    tweet.CreatedAt,
			RetweetCount: int64(tweet.RetweetCount),
			Text:         tweet.Text,
		}
		if err = stream.Send(r); err != nil {

			ts.Stop()
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}

	demux.HandleChan(ts.Messages)

	return nil
}
