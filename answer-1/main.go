package main

import (
	"log"
	"os"

	"m4j00/pkg/config"
	"m4j00/pkg/db"
	"m4j00/pkg/routes"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// set time zone
	os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	dbUrl := config.LoadDBConfig()
	h := db.Init(dbUrl)
	r := routes.SetupRouter(h)
	r.Run()
}
