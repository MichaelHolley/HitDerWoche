package types

type SimpleArtist struct {
	Name     string `json:"name"`
	ID       ID     `json:"id"`
	URI      URI    `json:"uri"`
	Endpoint string `json:"href"`
}
