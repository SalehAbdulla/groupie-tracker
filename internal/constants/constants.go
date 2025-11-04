package constants

const PORT = ":5171"
const TEMPLATES_PATH = "templates"

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

type RelationIndex struct {
	Index []RelationEntry `json:"index"`
}

type RelationEntry struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type HomePageData struct {
	Artists []ArtistView
}
