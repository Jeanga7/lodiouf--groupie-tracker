package gt

import (
	"encoding/json"
	"io"
	"net/http"
)

func HandleError(err error) {
	panic(err)
}

type Artist struct {
	id           int      `json:"id"`
	image        string   `json:"image"`
	name         string   `json:"name"`
	members      []string `json:"members"`
	creationDate int      `json:"creationDate"`
	firstAlbum   string   `json:"firstAlbum"`
	locations    string   `json:"locations"`
	concertDates string   `json:"concertDates"`
	relations    string   `json:"relations"`
}

func ReadAPI() []Artist {
	r, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		HandleError(err)
	}

	data, err1 := io.ReadAll(r.Body)
	if err1 != nil {
		HandleError(err1)
	}

	allArtists := []Artist{}

	err = json.Unmarshal(data, &allArtists)
	if err != nil {
		HandleError(err)
	}
	return allArtists
}
