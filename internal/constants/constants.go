package constants

const PORT = ":5171"
const TEMPLATES_PATH = "templates"

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type ArtistData struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type LocationsIndex struct {
	Index []LocationEntry `json:"index"`
}

type LocationEntry struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DatesIndex struct {
	Index []DateEntry `json:"index"`
}

type DateEntry struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationIndex struct {
	Index []RelationEntry `json:"index"`
}

type RelationEntry struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
