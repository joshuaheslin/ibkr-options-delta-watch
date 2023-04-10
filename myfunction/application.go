package myfunction

import (
	"fmt"
	"os"
)

func MakePost() {
	fmt.Printf(os.Getenv("FOO"))
	fmt.Printf(os.Getenv("OPENAI_API_KEY"))
	fmt.Printf(os.Getenv("TWITTER_API_KEY"))
	fmt.Println("Making post")
}
