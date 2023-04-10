package myfunction

import (
	"fmt"
	"os"
)

func MakePost() {
	fmt.Printf(os.Getenv("FOO"))
	fmt.Println("Making post")
}
