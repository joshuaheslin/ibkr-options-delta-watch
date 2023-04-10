package myfunction

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	MakePost()
	fmt.Fprintf(w, "Posting done!")
}
