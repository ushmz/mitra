package store

import (
	"context"
	"fmt"
	"mitra/domain"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

// UserAuthentication : Interface for the user authentication.
type AuthenticationStore interface {
	// RegisterUser : RegisterUser user to authentication service.
	RegisterUser(ctx context.Context, externalID, secret string) (*domain.ImplicitUser, error)

	// Login(externalID, secret string) error

	// GenerateToken : Generate new token with externalID as an UID.
	GenerateToken(ctx context.Context, externalID string) (string, error)

	// GenerateSessionCookie : Generate new session cookie.
	GenerateSessionCookie(ctx context.Context, externalID string, expiresIn time.Duration) (string, error)

	// RevokeToken : Revoke generated token with externalID as an UID.
	RevokeToken(ctx context.Context, externalID string) error

	// DeleteUser : DeleteUser user from authentication service.
	DeleteUser(ctx context.Context, externalID string) error
}

// AuthenticationImpl : Implemention of user authentication.
type AuthenticationImpl struct {
	App *firebase.App
}

// NewUserAuthenticationImpl : Return new UserAuthentication implemention
func NewAuthenticationStore(app *firebase.App) AuthenticationStore {
	return &AuthenticationImpl{App: app}
}

// RegisterUser : Register user with externalID and password.
func (u *AuthenticationImpl) RegisterUser(ctx context.Context, externalID, secret string) (*domain.ImplicitUser, error) {
	if u == nil {
		return nil, ErrNilReceiver
	}

	client, err := u.App.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("Try to get auth client: %w", err)
	}

	params := (&auth.UserToCreate{}).
		UID(externalID).
		Email(externalID + "@mitra.ushmz.org").
		EmailVerified(true).
		Password(secret)

	user, err := client.CreateUser(ctx, params)
	if err != nil {
		if auth.IsUIDAlreadyExists(err) {
			return nil, ErrUIDAlreadyExists
		}
		return nil, fmt.Errorf("Try to create user: %w", err)
	}

	return &domain.ImplicitUser{ExternalID: externalID, FirebaseUID: user.UID}, nil
}

// DeleteUser : Delete user from application.
func (u *AuthenticationImpl) DeleteUser(ctx context.Context, externalID string) error {
	if u == nil {
		return ErrNilReceiver
	}

	client, err := u.App.Auth(ctx)
	if err != nil {
		return fmt.Errorf("Try to get auth client: %w", err)
	}

	if err := client.DeleteUser(context.Background(), externalID); err != nil {
		return fmt.Errorf("Try to delete user: %w", err)
	}
	return nil
}

// GenerateToken : Generate new token with externalID as an UID.
func (u *AuthenticationImpl) GenerateToken(ctx context.Context, externalID string) (string, error) {
	if u == nil {
		return "", ErrNilReceiver
	}

	client, err := u.App.Auth(ctx)
	if err != nil {
		return "", fmt.Errorf("Try to get auth client: %w", err)
	}

	token, err := client.CustomToken(ctx, externalID)
	if err != nil {
		return "", fmt.Errorf("Try to generate user token: %w", err)
	}
	return token, nil
}

// GenerateSessionCookie : Generate new session cookie with externalID as an UID.
func (u *AuthenticationImpl) GenerateSessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error) {
	if u == nil {
		return "", ErrNilReceiver
	}

	client, err := u.App.Auth(ctx)
	if err != nil {
		return "", fmt.Errorf("Try to get auth client: %w", err)
	}

	cookie, err := client.SessionCookie(ctx, idToken, expiresIn)
	if err != nil {
		return "", fmt.Errorf("Try to generate auth cookie: %w", err)
	}
	return cookie, nil
}

// RevokeToken : Revoke generated token with externalID as an UID.
func (u *AuthenticationImpl) RevokeToken(ctx context.Context, externalID string) error {
	if u == nil {
		return ErrNilReceiver
	}

	client, err := u.App.Auth(ctx)
	if err != nil {
		return fmt.Errorf("Try to get auth client: %w", err)
	}

	if err := client.RevokeRefreshTokens(ctx, externalID); err != nil {
		return fmt.Errorf("Try to revoke token: %w", err)
	}
	return nil
}
