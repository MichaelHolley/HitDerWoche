package services

import (
	"backend/types"
)

type Track struct {
	ID               types.ID  `json:"id" gorm:"primarykey"`
	PlaylistPosition int       `json:"playlistPosition"`
	Artists          string    `json:"artists"`
	DiscNumber       int       `json:"disc_number"`
	Duration         int       `json:"duration_ms"`
	Explicit         bool      `json:"explicit"`
	Endpoint         string    `json:"href"`
	Name             string    `json:"name"`
	PreviewURL       string    `json:"preview_url"`
	TrackNumber      int       `json:"track_number"`
	URI              types.URI `json:"uri"`
	Type             string    `json:"type"`
}
