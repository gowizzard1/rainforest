package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const CHALLENGE_URL = "https://letsrevolutionizetesting.com/challenge.json"

type ChallengeResponse struct {
	Follow  string `json:"follow"`
	Message string `json:"message"`
}

func getQueryParam(u, param string) string {
	parsedUrl, _ := url.Parse(u)
	queryParams := parsedUrl.Query()
	return queryParams.Get(param)
}

func makeRequest(challengeID string) string {
	client := &http.Client{}

	reqUrl := CHALLENGE_URL
	if challengeID != "" {
		reqUrl += "?id=" + challengeID
	}

	resp, err := client.Get(reqUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var jsonResponse ChallengeResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		panic(err)
	}

	if jsonResponse.Follow != "" {
		id := getQueryParam(jsonResponse.Follow, "id")
		return makeRequest(id)
	}

	return jsonResponse.Message
}

func main() {
	message := makeRequest("")
	fmt.Println(message)
}
