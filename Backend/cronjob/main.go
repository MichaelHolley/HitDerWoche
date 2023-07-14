package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PlaylistResponse struct {
	Items []struct {
		Track struct {
			Name string `json:"name"`
		} `json:"track"`
	} `json:"items"`
}

func main() {
	// Set your Spotify access token here
	accessToken := "<your-access-token>"

	// Set the Spotify playlist ID here
	playlistID := "4RnajpiFvsWCxU3xqwXXrc"

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request to retrieve the playlist
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/{playlist_id}", playlistID)
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
