package service

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/hungtg7/backend-app/app/user/config"
	"github.com/hungtg7/backend-app/app/user/repo"
	"github.com/hungtg7/backend-app/app/user/service/oauth"
	"github.com/hungtg7/backend-app/lib/logging"
)

type Service struct {
	config            *config.Base
	googleOauthConfig *oauth2.Config
	repo              *repo.UserRepo
}

func NewService(cfg *config.Base, repo *repo.UserRepo) *Service {
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
		repo:              repo,
	}
}

func (s *Service) GoogleCallback(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		acc, err := oauth.OauthGoogleCallback(ctx, s.googleOauthConfig, w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !acc.Verified_email {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		eGoogleAccount, err := s.repo.GetGoogleAccountByID(ctx, acc.Id)
		if err != nil {
			logging.Log.Fatal("Can not get google account by id")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if eGoogleAccount == nil {
			// TODO: create account here
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

func (s *Service) GoogleLogin(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return oauth.OauthGoogleLogin(ctx, s.googleOauthConfig)
}
