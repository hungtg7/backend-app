package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"github.com/hungtg7/backend-app/lib/logging"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type GoogleAccount struct {
	Id             string
	Email          string
	Verified_email bool
	Picture        string
}

func OauthGoogleCallback(ctx context.Context, oauthConfig *oauth2.Config, w http.ResponseWriter, r *http.Request) (*GoogleAccount, error) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")
	acc := &GoogleAccount{}

	if r.FormValue("state") != oauthState.Value {
		logging.Log.Fatal("invalid oauth google state")
		return nil, errors.New("InternalError")
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"), oauthConfig)
	if err != nil {
		logging.Log.Fatal(err.Error())
		return nil, errors.New("InternalError")
	}

	err = json.Unmarshal(data, acc)
	if err != nil {
		logging.Log.Fatal(err.Error())
		http.Error(w, "InternalError", 500)
		return nil, errors.New("InternalError")
	}
	return acc, nil
}

func OauthGoogleLogin(ctx context.Context, oauthConfig *oauth2.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create oauthState cookie
		oauthState := generateStateOauthCookie(w)

		/*
			AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
			validate that it matches the the state query parameter on your redirect callback.
		*/
		u := oauthConfig.AuthCodeURL(oauthState)
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}

func getUserDataFromGoogle(code string, oauthConfig *oauth2.Config) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}
