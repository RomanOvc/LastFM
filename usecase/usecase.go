package usecase

import (
	"LastFM/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetContent(nameArtist, albumName string) string { //err

	//req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	//if err != nil {
	//	return nil, errors.Wrap(err, "Error create new request")
	//}
	//
	//query := req.URL.Query()
	//query.Add("part", "id")
	//query.Add("part", "contentDetails")
	//query.Add("part", "snippet")
	//query.Add("mine", "true")
	//query.Add("maxResults", "100")
	//req.URL.RawQuery = query.Encode()
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	return nil, errors.Wrap(err, "Error request")
	//}
	nameArtistT := &url.URL{Path: nameArtist}
	nameArtistTEncoder := nameArtistT.String()
	albumNameT := &url.URL{Path: albumName}
	albumNameTEncoder := albumNameT.String()

	url := "http://ws.audioscrobbler.com/2.0/?method=album.getinfo&api_key=a8e35717f897047c80a1b447e9438370&artist=" +
		nameArtistTEncoder + "&album=" + albumNameTEncoder + "&format=json"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data models.AlbumU

	if err = json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	var parsedData []models.AlbumTrack
	for i := range data.Album.Tracks.Track {
		parsedData = append(parsedData, models.AlbumTrack{Id: i + 1, URL: data.Album.Tracks.Track[i].Url, TrackName: data.Album.Tracks.Track[i].Name})
	}
	result := models.Result{AlbumName: (data.Album.Name), AlbumArtist: data.Album.Artist, AlbumTrack: parsedData}
	u, err := json.Marshal(result) // handler

	if err != nil {
		panic(err)
	}

	return string(u)
}
