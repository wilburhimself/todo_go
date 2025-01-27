package main

import (
	"net/http"
)

func main() {
	router := Router()
	http.ListenAndServe(":8080", router)
}
