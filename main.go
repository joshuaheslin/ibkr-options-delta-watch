package main

import (
	"fmt"
	"log"
	"os"

	"main/myfunction"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	fmt.Println(os.Getenv("FOO"))
	fmt.Println(os.Getenv("OPENAI_API_KEY"))
	fmt.Println(os.Getenv("TWITTER_API_KEY"))

	myfunction.MakePost()
}
