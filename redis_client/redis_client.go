package redis_client

import (
	"github.com/go-redis/redis"
	"os"
	"encoding/json"
)

func New() *redis.Client {
	options := redis.Options{}

	redisUrl := os.Getenv("REDIS_URL")
	if len(redisUrl) != 0 {
		options.Addr = redisUrl
	}

	return redis.NewClient(&options)
}

func ShortenUrl(url string, ownerId uint, options string) (string, error) {
	shortenedUrl := ShortenedUrl{Url: url, Clicks: 0, OwnerId: ownerId, Options: []byte(options)}
	marshaledShortenedUrl, err := json.Marshal(shortenedUrl)
	if err != nil {
		return "", err
	}

	var shUrl string

	redisClient := New()
	for {
		shUrl = Generate()
		_, err := redisClient.Get(shUrl).Result()
		if err == redis.Nil {
			_, err = redisClient.Set(shUrl, string(marshaledShortenedUrl), 0).Result()
			if err != nil {
				return "", err
			}
			break
		}
	}

	return shUrl, nil
}

func FindShortenedUrl(shUrl string) (ShortenedUrl, error) {
	redisClient := New()

	val, err := redisClient.Get(shUrl).Result()
	if err != nil {
		return ShortenedUrl{}, err
	}

	var shortenedUrl ShortenedUrl
	err2 := json.Unmarshal([]byte(val), &shortenedUrl)
	if err2 != nil {
		return ShortenedUrl{}, err2
	}

	return shortenedUrl, nil
}

func RemoveShortenedUrl(shUrl string) (bool, error) {
	redisClient := New()

	_, err := redisClient.Del(shUrl).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}
