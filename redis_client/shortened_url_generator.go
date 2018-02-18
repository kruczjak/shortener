package redis_client

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"strconv"
)

func Generate() string {
	shortLength, _ := strconv.Atoi(os.Getenv("SHORTENED_URL_LENGTH"))
	shortenedUrl, _ := GenerateRandomString(shortLength)
	return shortenedUrl
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
