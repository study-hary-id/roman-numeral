package main

import (
	"fmt"
	"github.com/study-hary-id/roman-numeral-api/handlers"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		size := len(urlPathElements)

		if size >= 3 {
			num, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if num > 0 {
				handlers.RomanNumber(w, r, num)
				return
			}
		}
		handlers.NotFound(w, r)
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
