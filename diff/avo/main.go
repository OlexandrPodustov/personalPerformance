package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start")

	if err := http.ListenAndServe(":8080", &mh{}); err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("end")
}

type mh struct{}

func (s *mh) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter echo handler server")
	fmt.Fprintf(w, "echo server %q", r.URL)
}
