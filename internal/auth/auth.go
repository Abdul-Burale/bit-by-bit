package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// global
var firebaseAuth *auth.Client

type FirebaseConfig struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

func createFirebaseConfig(filename string) error {
	config := FirebaseConfig{
		Type:                    os.Getenv("FIREBASE_TYPE"),
		ProjectID:               os.Getenv("FIREBASE_PROJECT_ID"),
		PrivateKeyID:            os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("FIREBASE_PRIVATE_KEY"),
		ClientEmail:             os.Getenv("FIREBASE_CLIENT_EMAIL"),
		ClientID:                os.Getenv("FIREBASE_CLIENT_ID"),
		AuthURI:                 os.Getenv("FIREBASE_AUTH_URI"),
		TokenURI:                os.Getenv("FIREBASE_TOKEN_URI"),
		AuthProviderX509CertURL: os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertURL:       os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
		UniverseDomain:          os.Getenv("FIREBASE_UNIVERSE_DOMAIN"),
	}

	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling Firebase config to JSON: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing Fiebase config to file %v", err)
	}

	log.Printf("Firebase configation written to %s", filename)
	return nil
}

func InitFireBase() error {
	credentials := "firebase_keys.json"

	if _, err := os.Stat(credentials); os.IsNotExist(err) {
		log.Println("Firebase credentials file not found. Creating file")

		err := createFirebaseConfig(credentials)
		if err != nil {
			return fmt.Errorf("error creating Firebase config: %v", err)
		}
	} else {
		log.Println("Firebase credentials file found.")
	}

	opt := option.WithCredentialsFile(credentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initalizing app: %v\n", err)
	}

	firebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	//	testVerifyIDToken()
	if err != nil {
		log.Fatalf("error generating ID Token (jwt)")
	}
	return nil
}

// TODO: VerifyID Token
func VerifyIDToken(authHeader string) (*auth.Token, error) {

	if authHeader == "" {
		return nil, errors.New("missing Authorization header")
	}

	partitions := strings.Split(authHeader, " ")
	if len(partitions) != 2 || partitions[0] != "Bearer" {
		return nil, errors.New("invalid Authorization header format")
	}

	idToken := partitions[1]
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		fmt.Println("Error verifying ID Token:", err)
		log.Fatalf("error verifying ID Token: %v\n", err)
	}
	log.Printf("Verified ID Token: %v\n", token)

	return token, nil
}
