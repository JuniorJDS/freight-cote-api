package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetSettings() map[string]string {
	ignoreEnviroment, err := strconv.ParseBool(os.Getenv("IGNORE_ENVIRONMENT"))
	if !ignoreEnviroment {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("Error loading .env file: ", err.Error())
	}

	settings := make(map[string]string)

	settings["API_V1"] = "/api/v1"
	settings["PORT"] = "5000"

	// Mongo Settings
	settings["MONGO_URI"] = os.Getenv("MONGO_URI")
	settings["MONGO_DATABASE"] = os.Getenv("MONGO_DATABASE")

	// External API Settings
	settings["FRETERAPIDO_API_URL"] = os.Getenv("FRETERAPIDO_API_URL")
	settings["TOKEN"] = os.Getenv("TOKEN")
	settings["PLATFORMCODE"] = os.Getenv("PLATFORMCODE")
	settings["DISPATCHERSZIPCODE"] = os.Getenv("DISPATCHERSZIPCODE")
	settings["REGISTEREDNUMBER"] = os.Getenv("REGISTEREDNUMBER")

	return settings
}
