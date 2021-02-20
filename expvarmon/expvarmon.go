package main

import (
	"net/http"

	_ "expvar"
)

func main() {
	http.ListenAndServe("0.0.0.0:8080", nil)
}
