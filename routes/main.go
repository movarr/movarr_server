package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home, Page!")
	})
	http.Handle("/", router)
}
