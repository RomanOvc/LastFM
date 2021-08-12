package models

// Unmarshal struct
type TrackInfoU struct {
	Url  string
	Name string
}
type TrackU struct {
	Track []TrackInfoU
}
type InfoU struct {
	Name   string
	Artist string
	Tracks TrackU
}
type AlbumU struct {
	Album InfoU
}

// Marshal struct
type AlbumTrack struct {
	Id        int    `json:"id"`
	TrackName string `json:"track_name"`
	URL       string `json:"url"`
}

type Result struct {
	AlbumName   string       `json:"album_mame"`
	AlbumArtist string       `json:"album_artist"`
	AlbumTrack  []AlbumTrack `json:"album_track"`
}

type Res struct {
	AlbumName   string `json:"album_mame"`
	AlbumArtist string `json:"album_artist"`
	AlbumTrack  []struct {
		Id        int    `json:"id"`
		TrackName string `json:"track_name"`
		URL       string `json:"url"`
	} `json:"album_track"`
}
