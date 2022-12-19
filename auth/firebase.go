package auth

import (
	"context"
	"errors"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// initFirebaseApp : Initialize project as firebase application
func InitFirebaseApp() (*firebase.App, error) {
	ctx := context.Background()
	credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_CREDENTIALS_JSON")))
	if err != nil {
		return nil, errors.New("Failed to load credential options from Env values")
	}
	opt := option.WithCredentials(credentials)
	return firebase.NewApp(ctx, nil, opt)
}

// getAuthClient : Return firebase authentication client to create new user and more
func GetAuthClient(app *firebase.App) (*auth.Client, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Failed to get auth client %v", err)
	}
	return client, nil
}
