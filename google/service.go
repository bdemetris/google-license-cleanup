package google

import (
	"context"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	admin "google.golang.org/api/admin/directory/v1"
)

const (
	delegate = "bdemetris-admin@flexport.com"
)

type Service struct {
	googleClient *admin.Service
}

func NewService() (*Service, error) {
	var s *Service

	gc, err := createGoogleClient()
	if err != nil {
		return nil, err
	}
	s.googleClient = gc

	return s, nil
}

func createGoogleClient() (*admin.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile("credentials-dev.json")
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(b, admin.AdminDirectoryUserScope)
	if err != nil {
		return nil, err
	}
	config.Subject = delegate
	ts := config.TokenSource(ctx)

	srv, err := admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	return srv, nil
}
