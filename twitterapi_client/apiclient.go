package twitterapiclient

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
	pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"
)

var client *twitter.Client

func init() {
	// The twitter package provides a Client for accessing the Twitter API.
	// Here we create a new twitter client using the credentials in the config file
	config := oauth1.NewConfig(viper.GetString("credentials.consumerKey"),
		viper.GetString("credentials.consumerSecret"))
	token := oauth1.NewToken(viper.GetString("credentials.accessToken"),
		viper.GetString("credentials.accessSecret"))

	// Authentication is handled by the http.Client which is
	// passed to twitter NewClient as Oauth1
	httpClient := config.Client(oauth1.NoContext, token)
	client = twitter.NewClient(httpClient)
}

// GetStream : returns new stream with the given params
func GetStream(params *pb.Params) (*twitter.Stream, error) {

	tp := &twitter.StreamFilterParams{}

	stream, err := client.Streams.Filter(tp)
	if err != nil { // Handle errors reading the config file
		return nil, err
	}

	return stream, nil

}
