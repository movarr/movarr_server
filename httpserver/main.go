package httpserver

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home, Page!")
	})
	fmt.Println("Starting server on port 10000")
	http.ListenAndServe(":10000", nil)
}
