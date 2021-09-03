package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/study-hary-id/roman-numeral-api/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		if urlPathElements[1] == "roman-numbers" {
			if len(urlPathElements) < 3 {
				handlers.IncompleteRouteHandler(w)

			} else {
				num, err := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
				if err != nil {
					handlers.WrongParamHandler(w)
				} else if num == 0 {
					handlers.ZeroParamHandler(w)
				} else {
					handlers.RomanNumberHandler(w, num)
				}
			}

		} else {
			// Response failure if using random endpoints.
			handlers.WrongRouteHandler(w, r)
		}
	})

	// Create a server and run it on 8000 port.
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server listening at http://localhost:8000")
	log.Fatal(s.ListenAndServe())
}
