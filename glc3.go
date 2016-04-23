package main

import (
	"fmt"
	dice "github.com/imel96/glc/dice"
	"net/http"
	"strings"
)

// http://quiksilver.local:8080/mf?r=24412,56316,66666,45625
func main() {
	http.HandleFunc("/mf", mfHandler)
	http.ListenAndServe(":8080", nil)
}

func mfHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "Rolls: %q", r.Form["r"])
	fmt.Fprintf(w, "Minimum faces:, %v", dice.MinimumFaces(strings.Split(r.Form["r"][0], ",")))
}
