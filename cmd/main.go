package main

import (
	"log"
	"net/http"

	"github.com/AbdulrahmanDaud10/wisp-manage/pkg/api"
	"github.com/AbdulrahmanDaud10/wisp-manage/pkg/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := api.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := app.SettingUpRoutes(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
