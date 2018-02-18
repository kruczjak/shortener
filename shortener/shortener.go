package shortener

import (
	"shortener/redis_client"
)

type Shortener struct {}

func (s Shortener) Shorten(url string) (string, error) {
	return redis_client.ShortenUrl(url, 0, "")
}

func (s Shortener) Resolve(shUrl string) (redis_client.ShortenedUrl, error) {
	return redis_client.FindShortenedUrl(shUrl)
}

func (s Shortener) Remove(shUrl string) (bool, error) {
	return redis_client.RemoveShortenedUrl(shUrl)
}
