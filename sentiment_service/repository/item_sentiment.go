package repository

import (
	mgo "gopkg.in/mgo.v2"
)

const (
	collection = "item_sentiment"
)

type Repository interface {
	Create(*SentimentScore) error
	GetAll() (*SentimentScore, error)
	Close()
}

// SentimentScore is a reflection of a sentiment score results type
type SentimentScore struct {
	Source    string
	CreatedAt string
	Text      string
	Score     int32
}

type ScoreRepository struct {
	Session *mgo.Session
}

// Create a new thing
func (repo *ScoreRepository) Create(item *SentimentScore) error {

	_, e := repo.collection().Upsert(item, item)
	return e
}

// GetAll consignments
func (repo *ScoreRepository) GetAll() ([]*SentimentScore, error) {
	var item []*SentimentScore
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	err := repo.collection().Find(nil).All(&item)
	return item, err
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
func (repo *ScoreRepository) Close() {
	repo.Session.Close()
}

func (repo *ScoreRepository) collection() *mgo.Collection {

	return repo.Session.DB(dbName).C(collection)
}
