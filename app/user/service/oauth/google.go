package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func OauthGoogleCallback(ctx context.Context, oauthConfig *oauth2.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read oauthState from Cookie
		oauthState, _ := r.Cookie("oauthstate")

		if r.FormValue("state") != oauthState.Value {
			fmt.Println("invalid oauth google state")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		data, err := getUserDataFromGoogle(r.FormValue("code"), oauthConfig)
		if err != nil {
			fmt.Println(err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		// GetOrCreate User in your db.
		// Redirect or response with a token.
		// More code .....
		fmt.Fprintf(w, "UserInfo: %s\n", data)
	}
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
