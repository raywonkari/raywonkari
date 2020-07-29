package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/raywonkari/raywonkari/blog"
	"github.com/raywonkari/raywonkari/strava"
	"github.com/raywonkari/raywonkari/twitter"
)

const (
	blogMdFile  = "blog.md"
	blogFeedURL = "https://raywontalks.com/feed.xml"

	twitterMdFile = "twitter.md"
	twitterURL    = "https://twitter.com/raywonkari"

	stravaMdFile = "strava.md"

	readmeHeader     = "# Hi, I'm [Raywon](https://raywonkari.com) :wave:\n"
	readmeHowItWorks = "[How it works](https://github.com/raywonkari/raywonkari/blob/master/HOW_IT_WORKS.md)\n"
	readmeFooter     = "<a href='https://github.com/raywonkari/raywonkari/actions'><img src='https://github.com/raywonkari/raywonkari/workflows/Build%20README/badge.svg' alt='Build README'></a>\n"
	readmeFile       = "README.md"
)

func main() {
	// generate twitter.md
	fmt.Println("Generating Twitter MD File")
	twitter.Generate(twitterMdFile, twitterURL)

	// generate strava.md
	fmt.Println("Generating Strava MD File")
	strava.Generate(stravaMdFile)

	// generate blog.md
	fmt.Println("Generating Blog MD File")
	blog.Generate(blogMdFile, blogFeedURL)

	// generate README.md
	fmt.Println("Generating Main README MD File")
	generateReadme()
}

func generateReadme() {
	// Create an empty file, or clear the contents of an existing file.
	ioutil.WriteFile(readmeFile, []byte(""), 0644)

	// Write header
	writeToFile(readmeHeader)
	writeToFile("\n")

	//Append blog.md to README.md
	writeToFile(string(readFromFile(blogMdFile)))

	//Append twitter.md to README.md
	writeToFile(string(readFromFile(twitterMdFile)))

	//Append strava.md to README.md
	writeToFile(string(readFromFile(stravaMdFile)))

	// Write footer
	writeToFile("\n")
	writeToFile(readmeHowItWorks)
	writeToFile("\n")
	writeToFile(readmeFooter)
}

func writeToFile(data string) {

	// Open readme md file
	openfile, err := os.OpenFile(readmeFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
	}
	defer openfile.Close()

	// Write to blog md file, if err, log the error
	if _, err := openfile.WriteString(data); err != nil {
		fmt.Printf("Failed to write to file: %v", err)
	}

}

func readFromFile(file string) []byte {
	openfile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error Reading %v: %v", file, err)
	}

	return openfile
}
