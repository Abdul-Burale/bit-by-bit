package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// MongoDB Methods
func GetMongoURI() string {
	// Return the value of the environment variable MONGO_URI
	return os.Getenv("MONGO_URI")
}

func GetMongoDatabaseName() string {
	return os.Getenv("DATABASE_NAME")
}

// Has to be in Block Capitals with underscores representing spaces. Example -> ORDER_COLLECTION
func GetMongoDatabaseCollection(collectionName string) string {
	return os.Getenv(collectionName)
}

// FireBase Methods

func GetFirebaseType() string {
	return os.Getenv("FIREBASE_TYPE")
}
func GetFirebaseProjectID() string {
	return os.Getenv("FIREBASE_PROJECT_ID")
}
func GetFirebasePrivateKeyID() string {
	return os.Getenv("FIREBASE_PRIVATE_KEY_ID")
}
func GetFirebasePrivateKey() string {
	return os.Getenv("FIREBASE_PRIVATE_KEY")
}
func GetFirebaseClientEmail() string {
	return os.Getenv("FIREBASE_CLIENT_EMAIL")
}
func GetFirebaseClientID() string {
	return os.Getenv("FIREBASE_CLIENT_ID")
}
func GetFirebaseAuth_URI() string {
	return os.Getenv("FIREBASE_AUTH_URI")
}
func GetFirebaseToken_URI() string {
	return os.Getenv("FIREBASE_TOKEN_URI")
}
func GetFirebaseAuthProviderX509Cert() string {
	return os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERTL_URL")
}
func GetFirebaseClientX509Cert() string {
	return os.Getenv("FIREBASE_CLIENT_X509_CERT_URL")
}
func GetFirebaseUniverseDomain() string {
	return os.Getenv("FIREBASE_UNIVERSE_DOMAIN")
}
