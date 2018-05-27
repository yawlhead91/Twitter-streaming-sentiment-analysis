package server

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	rss "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/rss_client"
	rs "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/rss_route"
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

// RssRouteServer Implements the generated
// RssRouteServer interface made in the proto file
type RssRouteServer struct{}

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

	// Run go routine to close the
	// twitter channel
	go func() {
		err = <-errors
		if err != nil {
			log.Print(err)
		}
		log.Print("Closing twitters wew")
		ts.Stop()
	}()

	// Receives messages and type switches them to
	// call functions with typed messages. As we are
	// only interested in the tweets we only add the
	// 'Tweet' handler
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		log.Print("Tweets")
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

func (s *RssRouteServer) GetRss(params *rs.ParamsRss, stream rs.RssRoute_GetRssServer) error {
	feeds := rss.StreamRssFeeds()

	for feed := range feeds {
		for _, item := range feed.Items {

			i := &rs.FeedItem{}
			i.CreatedAt = item.Published
			i.Title = item.Title
			i.Text = item.Content

			if err := stream.Send(i); err != nil {
				return err
			}
		}
	}
	return nil
}
