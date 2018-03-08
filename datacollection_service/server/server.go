package server

import (
	pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"
	twitter "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitterapi_client"
)

// TwitterRouteServer Implements the generated
// TwitterRouteServer interface made in the proto file
type TwitterRouteServer struct{}

// GetTweets creates a stream of tweets for the given params
// to be searched for from the twitter api
func (s *TwitterRouteServer) GetTweets(params *pb.Params, stream pb.TwitterRoute_GetTweetsServer) error {
	ts, err := twitter.GetTweets(params)
	if err != nil {
		return err
	}

	return nil
}
