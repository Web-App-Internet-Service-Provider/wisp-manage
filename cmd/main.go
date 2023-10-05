package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/handlers"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load("../build/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	return nil
}

func main() {
	fmt.Println("Web-App Internet Service Provider Management Service")

	// Loading Environmental Variable file
	if err := LoadEnv(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load .env file: %s\n", err)
		os.Exit(1)
	}

	if err := handlers.RunAPI(":8090"); err != nil {
		fmt.Fprintf(os.Stderr, "Server start failure: %s\n", err)
		os.Exit(1)
	}
}
