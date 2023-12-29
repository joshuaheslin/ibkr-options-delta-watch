package myfunction

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	Run()
	fmt.Fprintf(w, "Posting done!!!")
}
