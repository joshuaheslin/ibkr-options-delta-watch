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

	fmt.Println(os.Getenv("MARKET_DATA_API_KEY"))

	myfunction.Run()
}
