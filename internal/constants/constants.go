package constants

const PORT = ":5171"
const TEMPLATES_PATH = "templates"

type ArtistData struct {
	Id  int
	Image   string
	Name string
	Members  []string
	CreationDate int
	FirstAlbum string
	Locations string
	ConcertDates string
	Relations string
}
