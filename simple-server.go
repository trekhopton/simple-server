package main

import (
	"fmt"
	"net/http"
)

type page struct {
}

func (p page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World")
}

func main() {
	var p page
	err := http.ListenAndServe(":8080", p)
	if err != nil {
		panic(err)
	}
}
