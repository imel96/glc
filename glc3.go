package main

import (
	"fmt"
	dice "github.com/imel96/glc/dice"
	"net/http"
)

func main() {
	http.HandleFunc("/mf", mfHandler)
	http.ListenAndServe(":8080", nil)
}

func mfHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "Hello, %q", r.Form["r"])
	fmt.Fprintf(w, "Bye, %v", dice.MinimumFaces([]string{"137", "364", "115", "724"}))
}
