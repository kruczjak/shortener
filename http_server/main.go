package main

import (
	"net/http"
	"shortener/shortener"
	"github.com/joho/godotenv"
	"os"
	"shortener/config"
)

type handler struct {
	shortener shortener.Shortener
}

func (h handler) resolveAndRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Path[1:]

	shortenedUrl, err := h.shortener.Resolve(url)
	if err != nil {
		http.Redirect(w, r, os.Getenv("FAILED_REDIRECTION_URL"), http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, shortenedUrl.Url, http.StatusMovedPermanently)
}

func main() {
	godotenv.Load()

	h := handler{shortener: shortener.Shortener{}}

	http.HandleFunc("/", h.resolveAndRedirect)
	if err := http.ListenAndServe(config.Getenv("HTTP_HOST", ":8080"), nil); err != nil {
		panic(err)
	}
}
