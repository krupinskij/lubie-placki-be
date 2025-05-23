package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvClientPath() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("CLIENT_PATH")
}

func EnvServerPath() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("SERVER_PATH")
}
func EnvMongoURI() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGODB_URI")
}

func EnvClientId() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("CLIENT_ID")
}

func EnvClientSecret() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("CLIENT_SECRET")
}

func EnvAuthState() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("AUTH_STATE")
}
