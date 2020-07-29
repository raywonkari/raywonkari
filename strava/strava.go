package strava

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Generate func writes details about my Twitter profile
func Generate(fileName, URL string) {

	// Create an empty file, or clear the contents of an existing file.
	ioutil.WriteFile(fileName, []byte(""), 0644)

	writeToFile(fileName, "# Check out my activities on strava ![logo](./logo/strava.png)\n")

	writeToFile(fileName, "* https://strava.com/athletes/raywonkari\n")
	writeToFile(fileName, "* I recently started running and cycling.\n")
	writeToFile(fileName, "* At work, my colleagues used to form a team, and go out running every week, also some used to bike to work. I got inspired from them and took some time to prepare and get started.\n")
	writeToFile(fileName, "* My only moto with this is to inspire at least a few people to do so.\n")
	writeToFile(fileName, "* TODO: Embed latest activities here, instead of static data.\n")

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
