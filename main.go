package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/study-hary-id/roman-numeral-api/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		// Handle only roman-numbers endpoint.
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
			handlers.NotFound(w, r)
		}
	})

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	// Create a server and run it on PORT environment variable.
	s := &http.Server{
		Addr:           PORT,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(s.ListenAndServe())
}
