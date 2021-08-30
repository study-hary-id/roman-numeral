package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/study-hary-id/roman-numeral/romannumerals"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		if urlPathElements[1] == "roman-numbers" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))

			if number == 0 || number > 10 {
				// Response failure because it's not in the data.
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				// Response success with response html text.
				fmt.Fprintf(w, "%q", html.EscapeString(romannumerals.Numerals[number]))
			}

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
