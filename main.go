package main

import (
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		if urlPathElements[1] == "roman-numbers" {
			//number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			//fmt.Fprintf(w, "%d", number)
			w.Write([]byte("Successfully access roman-numbers"))

		} else {
			// Response failure if using random endpoints.
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})

	// Create a server and run it on 8000 port.
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
