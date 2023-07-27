package cronjob

import "backend/types"

type PlaylistResponse struct {
	Items []struct {
		Track types.SimpleTrack `json:"track"`
	} `json:"items"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
