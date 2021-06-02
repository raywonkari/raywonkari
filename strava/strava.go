package strava

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type parameters struct {
	Distance int `json:"distance"`
}

type activity struct {
	ThisYearRuns  parameters `json:"ytd_run_totals"`
	ThisYearRides parameters `json:"ytd_ride_totals"`
	AllTimeRuns   parameters `json:"all_run_totals"`
	AllTimeRides  parameters `json:"all_ride_totals"`
}

type token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//Generate func writes details about my strava activities
func Generate(fileName string) {

	// Get statistics from strava API & Extract the data
	stravaData := getStravaData()
	var strava activity
	json.Unmarshal(stravaData, &strava)

	// fmt.Println(strava)
	yearride := fmt.Sprintf("%.2f", meterToKilometer(strava.ThisYearRides.Distance))
	yearrun := fmt.Sprintf("%.2f", meterToKilometer(strava.ThisYearRuns.Distance))
	alltimeride := fmt.Sprintf("%.2f", meterToKilometer(strava.AllTimeRides.Distance))
	alltimerun := fmt.Sprintf("%.2f", meterToKilometer(strava.AllTimeRuns.Distance))

	yearCycling := fmt.Sprintf("* Total cycling distance from this year: %v km\n", yearride)
	yearRunning := fmt.Sprintf("* Total running distance from this year: %v km\n", yearrun)
	allTimeCycling := fmt.Sprintf("* All time cycling distance: %v km\n", alltimeride)
	allTimeRunning := fmt.Sprintf("* All time running distance: %v km\n", alltimerun)

	// Create an empty file, or clear the contents of an existing file.
	ioutil.WriteFile(fileName, []byte(""), 0644)

	writeToFile(fileName, "# Check out my activities on strava ![logo](https://github.com/raywonkari/raywonkari/blob/master/logo/strava.png)\n")

	writeToFile(fileName, "* https://strava.com/athletes/raywonkari\n")
	writeToFile(fileName, yearCycling)
	writeToFile(fileName, yearRunning)
	writeToFile(fileName, allTimeCycling)
	writeToFile(fileName, allTimeRunning)
	writeToFile(fileName, "* My moto with this is to inspire at least a few people to start exercising.\n")
}

func writeToFile(fileName, data string) {

	// Open strava md file
	openfile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
	}
	defer openfile.Close()

	// Write to strava md file, if err, log the error
	if _, err := openfile.WriteString(data); err != nil {
		fmt.Printf("Failed to write to file: %v", err)
	}

}

func getStravaData() []byte {

	// Get Access Token
	accessToken := getAccessToken()

	// Endpoint to retreive my strava statistics
	url := "https://www.strava.com/api/v3/athletes/54372323/stats?access_token=" + accessToken
	method := "GET"

	// Init HTTP Client & Request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	// Extract data
	byteValue, _ := ioutil.ReadAll(res.Body)
	return byteValue

}

func getAccessToken() string {

	// Endpoint to retrieve access token from
	url := "https://www.strava.com/api/v3/oauth/token"
	method := "POST"

	// Strava credentials
	clientID := os.Getenv("STRAVA_CLIENT_ID")
	clientSecret := os.Getenv("STRAVA_CLIENT_SECRET")
	refreshToken := os.Getenv("STRAVA_REFRESH_TOKEN")

	// Payload to send to strava
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("client_id", clientID)
	_ = writer.WriteField("client_secret", clientSecret)
	_ = writer.WriteField("refresh_token", refreshToken)
	_ = writer.WriteField("grant_type", "refresh_token")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	// Init HTTP Client & Request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	// Extract access token
	var token token
	json.Unmarshal(body, &token)

	// write refresh_token value to a file, so we can update it as a github secret in the later GH actions step.
	// GitHub requires using LibSodium package to encrypt values, and using that lib within Golang seems to be a hassle, therefore I am skipping it but will use nodejs instead.
	// Create an empty file, or clear the contents of an existing file.
	refreshTokenFileName := "strava_refresh_token"
	ioutil.WriteFile(refreshTokenFileName, []byte(""), 0644)
	writeToFile(refreshTokenFileName, token.RefreshToken)

	return string(token.AccessToken)
}

func meterToKilometer(input int) float64 {
	return float64(input) / 1000
}
