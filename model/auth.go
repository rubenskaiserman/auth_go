package model

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var config map[string]interface{}

type AuthTokens struct {
	IdToken      string
	RefreshToken string
}

func init() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		panic(err)
	}
}

// TODO: Add GCP credentials

// A problem is that the JWT will need to be stored in a database. This is because the JWT will need to be validated and refreshed. The JWT will also need to be deleted when the user logs out and it must be deleted either way when a certain amount of time has passed or when a new JWT is issued.

func IdentityProviderLogin(email string, password string) (AuthTokens, error) {
	baseUrl := config["authUrl"].(string)
	apiKey := config["authToken"].(string)
	url := fmt.Sprintf(baseUrl+"/accounts:signInWithPassword?key=%s", apiKey)

	payload := map[string]interface{}{
		"email":             email,
		"password":          password,
		"returnSecureToken": true,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return AuthTokens{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making the POST request:", err)
		return AuthTokens{}, err
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	authTokens := AuthTokens{
		IdToken:      result["idToken"].(string),
		RefreshToken: result["refreshToken"].(string),
	}

	return authTokens, nil
}

func Logout(authTokens AuthTokens) {
	// TODO: Implement Logouts

	// It should:
	// 1. validate the JWT
	// 		remove it from existance if it is valid
	// 		return an error if it is not valid
	// 2. return a success message
}

func ValidateJWT(authToken AuthTokens) {
	// TODO: Implement JWT validation

	// It should:
	// 1. validate the JWT
	// 		return an error if it is not valid
	// 2. return a success message
}

func RefreshJWT(authToken AuthTokens) {
	// TODO: Implement JWT refresh

	// It should:
	// 1. validate the JWT
	// 		return an error if it is not valid
	// 2. return a new JWT
}

func RetrieveJWT(authCode string) (AuthTokens, error) {
	// TODO: Implement JWT retrieval
	ctx := context.Background()
	projectID := config["gcpProjectId"].(string)
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile("./credentials/gcp_firestore.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer client.Close()

	queryResponse := *client.Collection("auth").Where("authCode", "==", authCode).Documents(ctx)
	tokens, err := queryResponse.GetAll()
	if err != nil {
		log.Fatalf("Failed to get Auth info: %v", err)
		return AuthTokens{}, err
	}

	if len(tokens) != 1 {
		return AuthTokens{}, fmt.Errorf("Invalid auth code")
	}

	timestamp := tokens[0].Data()["timestamp"].(time.Time)
	if time.Now().Sub(timestamp).Minutes() > 3 {
		return AuthTokens{}, fmt.Errorf("Auth code expired")
	}

	authTokens := AuthTokens{
		IdToken:      tokens[0].Data()["idToken"].(string),
		RefreshToken: tokens[0].Data()["refreshToken"].(string),
	}

	// It should:

	// 3. Check db writing timestamp
	// 4. Delete the JWT from the database
	// 5. if everything is successful and timestamp < 3 minutes, return the JWT
	return authTokens, nil
}

func GenAuthCode() (string, error) {
	length := 32

	byteLength := length / 2

	randomBytes := make([]byte, byteLength)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	authCode := hex.EncodeToString(randomBytes)

	return authCode, nil
}

func SaveJWT(authTokens AuthTokens, authCode string) error {
	ctx := context.Background()
	projectID := config["gcpProjectId"].(string)

	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile("./credentials/gcp_firestore.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer client.Close()

	_, _, err = client.Collection("auth").Add(ctx, map[string]interface{}{
		"authCode":     authCode,
		"idToken":      authTokens.IdToken,
		"refreshToken": authTokens.RefreshToken,
		"timestamp":    firestore.ServerTimestamp,
	})
	if err != nil {
		log.Fatalf("Failed to add Auth info: %v", err)
		return err
	}

	return nil
}
