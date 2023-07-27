package types

// SimpleAlbum contains basic data about an album.
type SimpleAlbum struct {
	// The name of the album.
	Name string `json:"name"`
	// A slice of SimpleArtists
	Artists []SimpleArtist `json:"artists"`
	// The field is present when getting an artist’s
	// albums. Possible values are “album”, “single”,
	// “compilation”, “appears_on”. Compare to album_type
	// this field represents relationship between the artist
	// and the album.
	AlbumGroup string `json:"album_group"`
	// The type of the album: one of "album",
	// "single", or "compilation".
	AlbumType string `json:"album_type"`
	// The SpotifyID for the album.
	ID ID `json:"id"`
	// The SpotifyURI for the album.
	URI URI `json:"uri"`
	// The markets in which the album is available,
	// identified using ISO 3166-1 alpha-2 country
	// codes.  Note that al album is considered
	// available in a market when at least 1 of its
	// tracks is available in that market.
	AvailableMarkets []string `json:"available_markets"`
	// A link to the Web API endpoint providing full
	// details of the album.
	Endpoint string `json:"href"`
	// The cover art for the album in various sizes,
	// widest first.
	Images []Image `json:"images"`
	// Known external URLs for this album.
	ExternalURLs map[string]string `json:"external_urls"`
	// The date the album was first released.  For example, "1981-12-15".
	// Depending on the ReleaseDatePrecision, it might be shown as
	// "1981" or "1981-12". You can use ReleaseDateTime to convert this
	// to a time.Time value.
	ReleaseDate string `json:"release_date"`
	// The precision with which ReleaseDate value is known: "year", "month", or "day"
	ReleaseDatePrecision string `json:"release_date_precision"`
}

// Copyright contains the copyright statement associated with an album.
type Copyright struct {
	// The copyright text for the album.
	Text string `json:"text"`
	// The type of copyright.
	Type string `json:"type"`
}

// FullAlbum provides extra album data in addition to the data provided by SimpleAlbum.
type FullAlbum struct {
	SimpleAlbum
	Copyrights []Copyright `json:"copyrights"`
	Genres     []string    `json:"genres"`
	// The popularity of the album, represented as an integer between 0 and 100,
	// with 100 being the most popular.  Popularity of an album is calculated
	// from the popularity of the album's individual tracks.
	Popularity  int               `json:"popularity"`
	Tracks      SimpleTrackPage   `json:"tracks"`
	ExternalIDs map[string]string `json:"external_ids"`
}

// SavedAlbum provides info about an album saved to an user's account.
type SavedAlbum struct {
	// The date and time the track was saved, represented as an ISO
	// 8601 UTC timestamp with a zero offset (YYYY-MM-DDTHH:MM:SSZ).
	// You can use the TimestampLayout constant to convert this to
	// a time.Time value.
	AddedAt   string `json:"added_at"`
	FullAlbum `json:"album"`
}

// AlbumType represents the type of an album. It can be used to filter
// results when searching for albums.
type AlbumType int

// AlbumType values that can be used to filter which types of albums are
// searched for.  These are flags that can be bitwise OR'd together
// to search for multiple types of albums simultaneously.
const (
	AlbumTypeAlbum AlbumType = 1 << iota
	AlbumTypeSingle
	AlbumTypeAppearsOn
	AlbumTypeCompilation
)
