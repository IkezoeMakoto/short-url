package main

import (
	"github.com/joho/godotenv"
	"os"
	"fmt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env")
	}

	fmt.Println("hello")
}