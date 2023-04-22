package music

type DownloadLink string
type TrackLink string
type FileName string

type Track struct {
	Title     string
	TrackLink TrackLink
}

type Music interface {
	Auth(apiToken string) error
	FindTrack(title string) []Track
	DownloadTrack(link DownloadLink) FileName
}
