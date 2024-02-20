package main

import (
	"demo/strings"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	PORT = ""
)

func main() {

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "58080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			bytes, err := io.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}

			rstr := strings.Reverse(string(bytes))
			fmt.Fprintln(w, rstr)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
	})

	log.Println("The server is up and running on the port:", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}
