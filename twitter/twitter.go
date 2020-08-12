package twitter

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Generate func writes details about my Twitter profile
func Generate(fileName, URL string) {

	// Create an empty file, or clear the contents of an existing file.
	ioutil.WriteFile(fileName, []byte(""), 0644)

	writeToFile(fileName, "# Check out my tweets ![logo](https://github.com/raywonkari/raywonkari/blob/master/logo/twitter.png)\n")

	writeToFile(fileName, "* https://twitter.com/raywonkari\n")
	writeToFile(fileName, "* TODO: My plan is to embed my latest tweets here.\n")

}

func writeToFile(fileName, data string) {

	// Open twitter md file
	openfile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
	}
	defer openfile.Close()

	// Write to twitter md file, if err, log the error
	if _, err := openfile.WriteString(data); err != nil {
		fmt.Printf("Failed to write to file: %v", err)
	}

}
