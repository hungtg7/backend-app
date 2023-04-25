package service

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/hungtg7/backend-app/app/user/config"
	"github.com/hungtg7/backend-app/app/user/service/oauth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Service struct {
	config            *config.Base
	googleOauthConfig *oauth2.Config
}

func NewService(cfg *config.Base) *Service {
	// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  fmt.Sprintf("http://localhost:%d/auth/google/callback", cfg.Server.Port),
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	return &Service{
		config:            cfg,
		googleOauthConfig: googleOauthConfig,
	}
}

func (s *Service) GoogleCallback(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return oauth.OauthGoogleCallback(ctx, s.googleOauthConfig)
}

func (s *Service) GoogleLogin(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return oauth.OauthGoogleLogin(ctx, s.googleOauthConfig)
}
