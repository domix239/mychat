package structs

type ChannelHost struct {
	Id        int
	URI       string
	SessionID string
	Alive     bool
	Guest     any
}
