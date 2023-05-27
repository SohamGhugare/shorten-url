package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

/*
	This function loads in all the environment variables stored in .env file
*/
func LoadEnvVars(){
	if err:=godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}
}