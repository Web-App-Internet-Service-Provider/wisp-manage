package main

import (
	"fmt"
	"log"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/handlers"
	"github.com/joho/godotenv"
)

func loadenv() {
	if err := godotenv.Load("../build/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Welcome To WISP API!!!!")

	//Loading Environmental Variable
	loadenv()

	log.Fatal(handlers.RunAPI(":8090"))

}
