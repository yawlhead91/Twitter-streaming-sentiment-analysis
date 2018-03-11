package repository

import (
	"log"

	pb "github.com/yawlhead91/Twitter-streaming-sentiment-analysis/datacollection_service/twitter_route"
	mgo "gopkg.in/mgo.v2"
)

const (
	collection = "tweet_sentiment"
)

type Repository interface {
	Create(*pb.Tweet) error
	GetAll() (*pb.Tweet, error)
	Close()
}

type TweetRepository struct {
	Session *mgo.Session
}

// Create a new thing
func (repo *TweetRepository) Create(tweet *pb.Tweet) error {

	log.Print("Collection: ", collection)
	_, e := repo.collection().Upsert(tweet, tweet)
	return e
}

// GetAll consignments
func (repo *TweetRepository) GetAll() ([]*pb.Tweet, error) {
	var tweet []*pb.Tweet
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	err := repo.collection().Find(nil).All(&tweet)
	return tweet, err
}

// Close closes the database session after each query has ran.
// Mgo creates a 'master' session on start-up, it's then good practice
// to copy a new session for each request that's made. This means that
// each request has its own database session. This is safer and more efficient,
// as under the hood each session has its own database socket and error handling.
// Using one main database socket means requests having to wait for that session.
// I.e this approach avoids locking and allows for requests to be processed concurrently. Nice!
// But... it does mean we need to ensure each session is closed on completion. Otherwise
// you'll likely build up loads of dud connections and hit a connection limit. Not nice!
func (repo *TweetRepository) Close() {
	repo.Session.Close()
}

func (repo *TweetRepository) collection() *mgo.Collection {

	return repo.Session.DB(dbName).C(collection)
}
