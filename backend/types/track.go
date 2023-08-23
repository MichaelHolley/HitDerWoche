package types

type TrackExternalIDs struct {
	ISRC string `json:"isrc"`
	EAN  string `json:"ean"`
	UPC  string `json:"upc"`
}

type SimpleTrack struct {
	Album            SimpleAlbum      `json:"album"`
	Artists          []SimpleArtist   `json:"artists"`
	AvailableMarkets []string         `json:"available_markets"`
	DiscNumber       int              `json:"disc_number"`
	Duration         int              `json:"duration_ms"`
	Explicit         bool             `json:"explicit"`
	ExternalIDs      TrackExternalIDs `json:"external_ids"`
	Endpoint         string           `json:"href"`
	ID               ID               `json:"id"`
	Name             string           `json:"name"`
	PreviewURL       string           `json:"preview_url"`
	TrackNumber      int              `json:"track_number"`
	URI              URI              `json:"uri"`
	Type             string           `json:"type"`
}
