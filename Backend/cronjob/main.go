package main

import (
	"encoding/json"
  "github.com/robfig/cron"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// URL of the authentication endpoint
const authURL = "https://accounts.spotify.com/api/token"

// Client credentials
const clientID = ""
const clientSecret = ""

// Playlist ID
const playlistID = ""

type PlaylistResponse struct {
	Items []struct {
		Track struct {
			Name string `json:"name"`
		} `json:"track"`
	} `json:"items"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func requestAccessToken() (token string, success bool) {

	// Create the form data with client credentials
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", clientID)
	formData.Set("client_secret", clientSecret)

	// Prepare the request
	req, err := http.NewRequest("POST", authURL, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return "", false
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return "", false
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Authentication failed. Status code:", resp.StatusCode)
		return "", false
	}

	var authResponse AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		fmt.Println("Failed to parse JSON response:", err)
		return
	}

	// Access the parsed response fields
	fmt.Println("Access Token:", authResponse.AccessToken)

	return authResponse.AccessToken, true
}

func updateData() {
	// Set your Spotify access token here
	accessToken, success := requestAccessToken()
	if success != true {
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request to retrieve the playlist
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", playlistID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Authorization header with the access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check if the request was successful (HTTP 200 status code)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	// Decode the response body into a PlaylistResponse struct
	var playlistResp PlaylistResponse
	err = json.NewDecoder(resp.Body).Decode(&playlistResp)
	if err != nil {
		log.Fatal(err)
	}

	// Display the song names
	fmt.Println("Playlist Songs:")
	for _, item := range playlistResp.Items {
		fmt.Println(item.Track.Name)
	}

	// TODO: Write to DB
}

func main() {
	// Create a new cron scheduler
	c := cron.New()

	// Schedule the cron job to run every 15 minutes
	c.AddFunc("*/15 * * * *", func() {
			fmt.Println("Running cronjob at:", time.Now())
			updateData()
	})

	// Start the cron scheduler
	c.Start()

	select {}
}
