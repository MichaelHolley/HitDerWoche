package types

// SimpleArtist contains basic info about an artist.
type SimpleArtist struct {
	Name string `json:"name"`
	ID   ID     `json:"id"`
	// The Spotify URI for the artist.
	URI URI `json:"uri"`
	// A link to the Web API endpoint providing full details of the artist.
	Endpoint     string            `json:"href"`
	ExternalURLs map[string]string `json:"external_urls"`
}

// FullArtist provides extra artist data in addition to what is provided by SimpleArtist.
type FullArtist struct {
	SimpleArtist
	// The popularity of the artist, expressed as an integer between 0 and 100.
	// The artist's popularity is calculated from the popularity of the artist's tracks.
	Popularity int `json:"popularity"`
	// A list of genres the artist is associated with.  For example, "Prog Rock"
	// or "Post-Grunge".  If not yet classified, the slice is empty.
	Genres    []string  `json:"genres"`
	Followers Followers `json:"followers"`
	// Images of the artist in various sizes, widest first.
	Images []Image `json:"images"`
}
