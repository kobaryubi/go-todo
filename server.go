package main

import (
	"net/http"
)

func main() {
	port := "8080"
	http.ListenAndServe(":"+port, nil)
}
