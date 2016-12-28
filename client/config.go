package client

import (
	"fmt"
	"log"
	"net/http"
)

//BootClient ...
func BootClient() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "alive\n")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
