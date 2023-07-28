package services

import (
	"backend/types"
	"log"
)

func GetTracks() []Track {
	db, err := GetDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	var tracks = []Track{}
	db.Find(&tracks)

	return tracks
}

func SaveTracks(tracks []types.SimpleTrack) {
	db, err := GetDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// iterate tracks
	for i, t := range tracks {
		artistNames := ""

		// build artists-string
		for j, a := range t.Artists {
			if j != 0 {
				artistNames += ", "
			}

			artistNames += a.Name
		}

		var track = Track{
			ID:               t.ID,
			PlaylistPosition: i + 1,
			Name:             t.Name,
			Artists:          artistNames,
			DiscNumber:       t.DiscNumber,
			Duration:         t.Duration,
			Explicit:         t.Explicit,
			Endpoint:         t.Endpoint,
			PreviewURL:       t.PreviewURL,
			TrackNumber:      t.TrackNumber,
			URI:              t.URI,
			Type:             t.Type,
		}

		result := db.Save(&track)
		if result.Error != nil {
			log.Fatal("Failed to upsert user:", result.Error)
		}
	}
}
