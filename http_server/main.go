package main

import (
	"net/http"
	"shortener/shortener"
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
		// something went wrong
	}

	http.Redirect(w, r, shortenedUrl.Url, http.StatusMovedPermanently)
}

func main() {
	h := handler{shortener: shortener.Shortener{}}

	http.HandleFunc("/", h.resolveAndRedirect)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
