package main

import (
	"fmt"
	dice "github.com/imel96/glc/dice"
//	dice "glc/dice"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/mf", mfHandler)
	http.ListenAndServe(":8080", nil)
}

func mfHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "Rolls: %q", r.Form["r"])
	fmt.Fprintf(w, "Minimum faces:, %v", dice.MinimumFaces(strings.Split(r.Form["r"][0], ",")))
}
