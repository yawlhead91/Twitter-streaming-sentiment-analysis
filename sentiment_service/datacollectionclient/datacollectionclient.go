package datacollectionclient

import (
	"context"
	"flag"
	"io"
	"log"

	pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"
	sentiment "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service/sentiment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:5253", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.twitter.com", "The server name use to verify the hostname returned by TLS handshake")
)

var client pb.TwitterRouteClient

// StreamTweets : returns a stream of tweets from
// the data collection service
func StreamTweets() error {
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = "../authentication/CA.pem"
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewTwitterRouteClient(conn)

	params := &pb.Params{
		Track:         []string{"Bitcoin"},
		Language:      []string{"en"},
		StallWarnings: false,
		Maxcount:      100,
	}

	stream, err := client.GetTweets(context.Background(), params)
	if err != nil {
		return err
	}

	for {
		tweet, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		err = sentiment.Run(tweet)
		if err != nil {
			return err
		}
	}

	return nil
}
