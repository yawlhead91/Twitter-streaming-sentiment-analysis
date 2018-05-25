package rss

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mmcdole/gofeed"
)

type Website struct {
	Name string `json:"Name"`
	URL  string `json:"Url"`
}

var websites = []Website{}

// StreamRssFeeds :
func StreamRssFeeds() chan *gofeed.Feed {
	err := loadfeeds()
	if err != nil {
		panic(err)
	}

	stream := make(chan *gofeed.Feed, len(websites))
	defer close(stream)

	for _, website := range websites {
		feed, _ := gofeed.NewParser().ParseURL(website.URL)
		stream <- feed
	}

	return stream
}

//build a lists of website from websites.yaml, if no file exists, it creates a default file
func loadfeeds() error {
	if _, err := os.Stat("rss.json"); os.IsNotExist(err) {
		return err
	}
	//import file and read websites
	fileHandler, _ := ioutil.ReadFile("rss.json")
	json.Unmarshal([]byte(fileHandler), &websites)
	return nil
}
