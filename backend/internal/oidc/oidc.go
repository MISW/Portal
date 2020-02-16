package oidc

import (
	"context"
	"net/http"

	"github.com/MISW/Portal/backend/internal/tokenutil"
	"github.com/coreos/go-oidc"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/xerrors"
)

// AuthResult - OpenID ConnectによるID Token取得の結果
type AuthResult struct {
	Token   *oauth2.Token
	IDToken *oidc.IDToken
	Profile map[string]interface{}
}

// Authenticator - OpenID Connectによる認証を行う
type Authenticator interface {
	// Login - OpenID Connect Providerへのredirect URLとstateを生成する
	Login(ctx context.Context) (redirectURL, state string, err error)

	// Callback - 外部からのCallbackを検証する
	Callback(ctx context.Context, expectedState, actualState, code string) (*AuthResult, error)
}

// NewAuthenticator initializes an Authenticator
func NewAuthenticator(ctx context.Context, clientID, clientSecret, redirectURL, providerURL string, scopes []string) (Authenticator, error) {
	provider, err := oidc.NewProvider(ctx, providerURL)
	if err != nil {
		return nil, xerrors.Errorf("failed to initialize oidc provider: %w", err)
	}

	conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       scopes,
	}

	return &authenticator{
		provider: provider,
		config:   conf,
	}, nil
}

type authenticator struct {
	provider *oidc.Provider
	config   oauth2.Config
}

// Callback - 外部からのCallbackを検証する
func (author *authenticator) Callback(ctx context.Context, expectedState, actualState, code string) (*AuthResult, error) {
	token, err := author.config.Exchange(ctx, code)
	if err != nil {
		return nil, xerrors.Errorf("failed to exchange code: %w", err)
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, xerrors.Errorf("No id_token field in oauth2 token: %w", http.StatusInternalServerError)
	}

	oidcConfig := &oidc.Config{
		ClientID: author.config.ClientID,
	}

	idToken, err := author.provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)

	if err != nil {
		return nil, xerrors.Errorf("Failed to verify ID Token: %w", err)
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return nil, xerrors.Errorf("failed to get userinfo: %w", err)
	}

	return &AuthResult{
		Token:   token,
		IDToken: idToken,
		Profile: profile,
	}, nil
}

// Login - OpenID Connect Providerへのredirect URLとstateを生成する
func (author *authenticator) Login(ctx context.Context) (redirectURL, state string, err error) {
	state, err = tokenutil.GenerateRandomToken()

	if err != nil {
		return "", "", xerrors.Errorf("failed to generate a random token: %w", err)
	}

	hashedState, err := bcrypt.GenerateFromPassword([]byte(state), bcrypt.DefaultCost)

	if err != nil {
		return "", "", xerrors.Errorf("failed to hash the state string: %w", err)
	}

	redirectURL = author.config.AuthCodeURL(string(hashedState))

	return
}
