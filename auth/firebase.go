package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// initFirebaseApp : Initialize project as firebase application
func InitFirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile("mitra-firebase-adminsdk.json")
	return firebase.NewApp(context.Background(), nil, opt)
}

// getAuthClient : Return firebase authentication client to create new user and more
func GetAuthClient(app *firebase.App) (*auth.Client, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Failed to get auth client %v", err)
	}
	return client, nil
}
