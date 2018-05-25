package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	rs "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/rss_route"
	s "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/server"
	pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"
	t "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitterapi_client"
)

const (
	port = 5253
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

	t.Auth()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTwitterRouteServer(grpcServer, &s.TwitterRouteServer{})
	rs.RegisterRssRouteServer(grpcServer, &s.RssRouteServer{})
	// determine whether to use TLS
	log.Printf("Serving twitter route server on : %d", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}
