package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
)

func main() {
	// Load the config file using viper the
	// config file holds you access keys
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// The twitter package provides a Client for accessing the Twitter API.
	// Here we create a new twitter client using the credentials in the config file
	config := oauth1.NewConfig(viper.GetString("credentials.consumerKey"),
		viper.GetString("credentials.consumerSecret"))
	token := oauth1.NewToken(viper.GetString("credentials.accessToken"),
		viper.GetString("credentials.accessSecret"))

	// Authentication is handled by the http.Client which is
	// passed to twitter NewClient as Oauth1
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	params := &twitter.StreamFilterParams{
		Track:         []string{"bitcoin"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(params)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}

	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	stream.Stop()

}
