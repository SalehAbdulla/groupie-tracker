package constants

import (
	"errors"
	"net/http"
)

const PORT = ":5171"
const TEMPLATES_PATH = "templates"

var (
	InternalServerError = errors.New(http.StatusText(http.StatusInternalServerError))
	NotFoundPage = errors.New(http.StatusText(http.StatusNotFound))
)

type ArtistView struct {
	ArtistData
	Locs []string
	Rel  map[string][]string
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

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type HomePageData struct {
	Artists []ArtistView
}

type CardPageData struct {
	PageData []Relation `json:"pageData"`
}

type Error struct {
	Error string
}

type DateLocations map[string][]string

type IndexItem struct {
	ID             int             `json:"id"`
	DatesLocations DateLocations   `json:"datesLocations"`
}

type ArtistIndex struct {
	Index []IndexItem `json:"index"`
}
