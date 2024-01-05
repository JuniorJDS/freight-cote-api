package configs

import "os"

func GetSettings() map[string]string {
	// Global Settings

	// err := godotenv.Load()
	// if err != nil {
	//	log.Fatal("Error loading .env file: ", err.Error())
	// }

	settings := make(map[string]string)

	settings["API_V1"] = "/api/v1"
	settings["PORT"] = "5000"

	// Mongo Settings
	settings["MONGO_URI"] = os.Getenv("MONGO_URI")
	settings["MONGO_DATABASE"] = os.Getenv("MONGO_DATABASE")

	// settings["HOST"] = os.Getenv("HOST")

	return settings
}
