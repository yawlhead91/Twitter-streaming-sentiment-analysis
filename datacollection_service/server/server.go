package server

import pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"

// TwitterRouteServer Implements the generated
// TwitterRouteServer interface made in the proto file
type TwitterRouteServer struct{}

// GetTweets creates a stream of tweets for the given params
// to be searched for from the twitter api
func (s *TwitterRouteServer) GetTweets(params *pb.Params, stream pb.TwitterRoute_GetTweetsServer) error {
	return nil
}
