package types

// SimpleArtist contains basic info about an artist.
type SimpleArtist struct {
	Name string `json:"name"`
	ID   ID     `json:"id"`
	// The Spotify URI for the artist.
	URI URI `json:"uri"`
	// A link to the Web API endpoint providing full details of the artist.
	Endpoint string `json:"href"`
	// TODO ExternalURLs map[string]string `json:"external_urls"`
}
