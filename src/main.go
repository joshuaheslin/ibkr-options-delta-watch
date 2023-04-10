package myfunction

import (
	"fmt"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("22")
	fmt.Fprintf(w, "Hello, World!")
	fmt.Printf(os.Getenv("FOO"))
}
