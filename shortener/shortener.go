package shortener

import "github.com/go-redis/redis"

type Config struct {
	redisOptions redis.Options
	uniqueKeyLength uint
}

type Shortener struct {}

func (s Shortener) Shorten(url string) (string, error) {
	return "", nil
}

func (s Shortener) Resolve(shortenedUrl string) (ShortenedUrl, error) {
	return ShortenedUrl{Url: "http://wykop.pl"}, nil
}

func (s Shortener) Remove(shortenedUrl string) (bool, error) {
	return true, nil
}
