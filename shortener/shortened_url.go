package shortener

type ShortenedUrl struct {
	Url string
	Clicks uint32
	OwnerId uint
	Options []byte // json
}
