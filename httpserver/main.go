package httpserver

import (
	"fmt"
	"net/http"
)

func StartServer() {
	fmt.Println("Starting server on port 10000")
	http.ListenAndServe(":10000", nil)
}
