package blog

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

type channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Items []item `xml:"item"`
}

type rss struct {
	Channel channel `xml:"channel"`
}

// It is often called as the magical reference date. Google it.
const (
	timeRSS       = "Mon, 2 Jan 2006 15:04:05 MST"
	timeFormatted = "January 02, 2006"
)

//Generate func queries my blog's rss feed, and generates an md file with the lates articles
func Generate(fileName, feedURL string) {

	// Fetch data from RSS feed
	resp, err := http.Get(feedURL)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	defer resp.Body.Close()

	// decode XML and assign to 'data' variable which is of 'rss' type. Refer strucks on top.
	data := rss{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
	}

	// Create an empty file, or clear the contents of an existing file.
	ioutil.WriteFile(fileName, []byte(""), 0644)

	// Heading of blog section
	blogHeading := fmt.Sprintf("# Latest Articles on my [blog](%v) :black_nib: :file_folder: :paperclip:\n", data.Channel.Link)
	writeToFile(fileName, blogHeading)

	for i, item := range data.Channel.Items {

		// Converting RFC822 format (typically used in RSS feeds) to human readable format
		// We parse the input time with timeRSS, and then format the result with timeFormatted.
		// Mon, 2 Jan 2006 15:04:05 MST => January 02, 2006
		format, _ := time.Parse(timeRSS, item.PubDate)
		formattedTimestamp := format.Format(timeFormatted)

		blogArticle := fmt.Sprintf("* [%v](%v) - %v\n", item.Title, item.Link, formattedTimestamp)
		writeToFile(fileName, blogArticle)

		if i == 4 {
			// only mentioning latest 5 articles
			break
		}
	}

}

func writeToFile(fileName, data string) {

	// Open blog md file
	openfile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
	}
	defer openfile.Close()

	// Write to blog md file, if err, log the error
	if _, err := openfile.WriteString(data); err != nil {
		fmt.Printf("Failed to write to file: %v", err)
	}

}
