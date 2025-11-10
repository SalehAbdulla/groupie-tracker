package constants

import (
	"errors"
	"net/http"
)

const PORT = ":5171"

var (
	ErrInternalServer = errors.New(http.StatusText(http.StatusInternalServerError))
	ErrNotFound       = errors.New(http.StatusText(http.StatusNotFound))
)

type ArtistView struct {
	ArtistData
	Rel map[string][]string
}

type ArtistData struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
}

type IndexItem struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type CardPageData struct {
	Index []IndexItem `json:"index"`
}

type Error struct {
	Error string
}

type HomePageData struct {
	Artists []ArtistView
}
