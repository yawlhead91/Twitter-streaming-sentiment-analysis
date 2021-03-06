package datacollectionclient

import (
	"context"
	"flag"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	rss "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/rss_route"
	pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"
	r "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service/repository"
	sentiment "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/sentiment_service/sentiment"
	mgo "gopkg.in/mgo.v2"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverHostOverride = flag.String("server_host_override", "x.test.twitter.com", "The server name use to verify the hostname returned by TLS handshake")
	serverAddr         = "127.0.0.1:5253"
)

var client pb.TwitterRouteClient

// StreamTweets : returns a stream of tweets from
// the data collection service
func StreamTweets(session *mgo.Session) error {

	// Database host from the environment variables
	host := os.Getenv("serverAddr")
	if host != "" {
		serverAddr = host
	}

	repo := &r.ScoreRepository{session}
	defer repo.Close()

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

	conn, err := grpc.Dial(serverAddr, opts...)
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

		score, err := sentiment.Run(tweet.Text)
		if err != nil {
			return err
		}

		log.Print("Receving tweet")

		entry := &r.SentimentScore{}
		entry.Source = "Twitter"
		entry.CreatedAt = tweet.CreatedAt
		entry.Score = int32(score)
		entry.Text = tweet.Text

		err = repo.Create(entry)
		if err != nil {
			return err
		}
	}

	return nil
}

func StreamRss(session *mgo.Session) error {

	// Database host from the environment variablesx
	host := os.Getenv("serverAddr")
	if host != "" {
		serverAddr = host
	}

	repo := &r.ScoreRepository{session}
	defer repo.Close()

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

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := rss.NewRssRouteClient(conn)

	params := &rss.ParamsRss{
		Maxcount: 100,
	}

	stream, err := client.GetRss(context.Background(), params)
	if err != nil {
		return err
	}

	for {
		item, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		score, err := sentiment.Run(item.Text)
		if err != nil {
			return err
		}

		entry := &r.SentimentScore{}
		entry.Source = "Rss"
		entry.CreatedAt = item.CreatedAt
		entry.Score = int32(score)
		entry.Text = item.Title

		err = repo.Create(entry)
		if err != nil {
			return err
		}
	}

	return nil
}
