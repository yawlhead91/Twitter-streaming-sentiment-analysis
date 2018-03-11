package sentiment

import (
	"github.com/cdipaolo/sentiment"
)

// Run sentiment summary  on the  given tweet, the
// package is trained on IMDB movies text data so
// the accuracy is null but gives a rough estimate
func Run(tweet string) (uint8, error) {

	model, err := sentiment.Restore()
	if err != nil {
		return 0, err
	}

	analysis := model.SentimentAnalysis(tweet, sentiment.English)

	return analysis.Score, nil
}
