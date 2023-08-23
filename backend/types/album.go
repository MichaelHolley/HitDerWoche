package types

type SimpleAlbum struct {
	Name                 string         `json:"name"`
	Artists              []SimpleArtist `json:"artists"`
	AlbumGroup           string         `json:"album_group"`
	AlbumType            string         `json:"album_type"`
	ID                   ID             `json:"id"`
	URI                  URI            `json:"uri"`
	AvailableMarkets     []string       `json:"available_markets"`
	Endpoint             string         `json:"href"`
	Images               []Image        `json:"images"`
	ReleaseDate          string         `json:"release_date"`
	ReleaseDatePrecision string         `json:"release_date_precision"`
}
