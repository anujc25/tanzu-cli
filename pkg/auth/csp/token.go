// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package csp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"

	configapi "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"

	"github.com/vmware-tanzu/tanzu-cli/pkg/constants"
	"github.com/vmware-tanzu/tanzu-cli/pkg/interfaces"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
)

const (
	// AuthTokenDir is a directory where cluster access token and refresh tokens are stored.
	AuthTokenDir = "tokens"

	// ExtraIDToken is the key in the Extra fields map that contains id_token.
	ExtraIDToken = "id_token"

	// StgIssuer is the VMware CSP(VCSP) staging issuer.
	StgIssuer = "https://console-stg.cloud.vmware.com/csp/gateway/am/api"

	// ProdIssuer is the VMware CSP(VCSP) issuer.
	ProdIssuer = "https://console.cloud.vmware.com/csp/gateway/am/api"

	// StgIssuerTCSP is the Tanzu CSP (TCSP) staging issuer.
	StgIssuerTCSP = "https://console-stg.tanzu.broadcom.com/csp/gateway/am/api"

	// ProdIssuerTCSP is the Tanzu CSP (TCSP) issuer
	ProdIssuerTCSP = "https://console.tanzu.broadcom.com/csp/gateway/am/api"

	//nolint:gosec // Avoid "hardcoded credentials" false positive.
	// APITokenKey is the env var for an API token override.
	APITokenKey = "CSP_API_TOKEN"
)

// Token types
const (
	// APITokenType Token type to denote the token obtained using API token
	APITokenType = "api-token"
	// IDTokenType Token type to denote the token obtained using interactive login flow
	IDTokenType = "id-token"
)

var (
	// DefaultKnownIssuers are known OAuth2 endpoints in each CSP environment.
	DefaultKnownIssuers = map[string]oauth2.Endpoint{
		StgIssuer: {
			AuthURL:   "https://console-stg.cloud.vmware.com/csp/gateway/discovery",
			TokenURL:  "https://console-stg.cloud.vmware.com/csp/gateway/am/api/auth/authorize",
			AuthStyle: oauth2.AuthStyleInHeader,
		},
		ProdIssuer: {
			AuthURL:   "https://console.cloud.vmware.com/csp/gateway/discovery",
			TokenURL:  "https://console.cloud.vmware.com/csp/gateway/am/api/auth/authorize",
			AuthStyle: oauth2.AuthStyleInHeader,
		},
		StgIssuerTCSP: {
			AuthURL:   "https://console-stg.tanzu.broadcom.com/csp/gateway/discovery",
			TokenURL:  "https://console-stg.tanzu.broadcom.com/csp/gateway/am/api/auth/authorize",
			AuthStyle: oauth2.AuthStyleInHeader,
		},
		ProdIssuerTCSP: {
			AuthURL:   "https://console.tanzu.broadcom.com/csp/gateway/discovery",
			TokenURL:  "https://console.tanzu.broadcom.com/csp/gateway/am/api/auth/authorize",
			AuthStyle: oauth2.AuthStyleInHeader,
		},
	}
	httpRestClient interfaces.HTTPClient
)

func init() {
	httpRestClient = http.DefaultClient
}

// IDTokenFromTokenSource parses out the id token from extra info in tokensource if available, or returns empty string.
func IDTokenFromTokenSource(token *oauth2.Token) (idTok string) {
	extraTok := token.Extra("id_token")
	if extraTok != nil {
		idTok = extraTok.(string)
	}
	return
}

// Token is a CSP token.
type Token struct {
	// IDToken for OIDC.
	IDToken string `json:"id_token"`

	// TokenType is the type of token.
	TokenType string `json:"token_type"`

	// ExpiresIn is expiration in seconds.
	ExpiresIn int64 `json:"expires_in"`

	// Scope of the token.
	Scope string `json:"scope"`

	// AccessToken from CSP.
	AccessToken string `json:"access_token"`

	// RefreshToken for use with Refresh Token grant.
	RefreshToken string `json:"refresh_token"`
}

// GetAccessTokenFromAPIToken fetches CSP access token using the API-token.
func GetAccessTokenFromAPIToken(apiToken, issuer string) (*Token, error) {
	api := fmt.Sprintf("%s/auth/api-tokens/authorize", issuer)
	data := url.Values{}
	data.Set("refresh_token", apiToken)
	req, _ := http.NewRequestWithContext(context.Background(), "POST", api, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpRestClient.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to obtain access token. Please provide valid VMware Cloud Services API-token")
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, errors.Errorf("Failed to obtain access token. Please provide valid VMware Cloud Services API-token -- %s", string(body))
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	token := Token{}

	if err = json.Unmarshal(body, &token); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal auth token")
	}

	return &token, nil
}

// GetIssuer returns the appropriate CSP issuer based on the environment.
func GetIssuer(staging bool) string {
	cspMetadata := GetCSPMetadata()
	if staging {
		return cspMetadata.IssuerStaging
	}
	return cspMetadata.IssuerProduction
}

// IsExpired checks for the token expiry and returns true if the token has expired else will return false
func IsExpired(tokenExpiry time.Time) bool {
	// refresh at half token life
	two := 2
	now := time.Now().Unix()
	halfDur := -time.Duration((tokenExpiry.Unix()-now)/int64(two)) * time.Second
	return tokenExpiry.Add(halfDur).Unix() < now
}

// ParseToken parses the token.
func ParseToken(tkn *oauth2.Token) (*Claims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tkn.AccessToken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	c, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not parse claims")
	}
	perm := []string{}
	p, ok := c["perms"].([]interface{})
	if !ok {
		log.Warning("could not parse perms from token")
	}
	for _, i := range p {
		perm = append(perm, i.(string))
	}
	uname, ok := c["username"].(string)
	if !ok {
		return nil, fmt.Errorf("could not parse username from token")
	}
	orgID, ok := c["context_name"].(string)
	if !ok {
		return nil, fmt.Errorf("could not parse orgID from token")
	}
	claims := &Claims{
		Username:    uname,
		Permissions: perm,
		OrgID:       orgID,
		Raw:         c,
	}
	return claims, nil
}

// Claims are the jwt claims.
type Claims struct {
	Username    string
	Permissions []string
	OrgID       string
	Raw         map[string]interface{}
}

// GetToken fetches a token for the current auth context.
func GetToken(g *configapi.GlobalServerAuth) (*oauth2.Token, error) {
	var token *Token
	var err error
	var orgID string
	if !IsExpired(g.Expiration) {
		tok := &oauth2.Token{
			AccessToken: g.AccessToken,
			Expiry:      g.Expiration,
		}
		return tok.WithExtra(map[string]interface{}{
			"id_token": g.IDToken,
		}), nil
	}
	if g.Type == APITokenType {
		token, err = GetAccessTokenFromAPIToken(g.RefreshToken, g.Issuer)
		if err != nil {
			return nil, err
		}
	} else if g.Type == IDTokenType {
		orgID, err = getOrgIDFromAccessToken(g)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get the CSP OrgID from the existing access token")
		}
		loginOptions := []LoginOption{WithRefreshToken(g.RefreshToken), WithOrgID(orgID), WithListenerPortFromEnv(constants.TanzuCLIOAuthLocalListenerPort)}
		token, err = TanzuLogin(g.Issuer, loginOptions...)
		if err != nil {
			return nil, err
		}
	}
	claims, err := ParseToken(&oauth2.Token{AccessToken: token.AccessToken})
	if err != nil {
		return nil, err
	}
	expiration := time.Now().Local().Add(time.Second * time.Duration(token.ExpiresIn))
	g.Expiration = expiration
	g.RefreshToken = token.RefreshToken
	g.AccessToken = token.AccessToken
	g.IDToken = token.IDToken
	g.Permissions = claims.Permissions

	tok := &oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       expiration,
	}
	return tok.WithExtra(map[string]interface{}{
		"id_token": token.IDToken,
	}), nil
}

// getOrgIDFromAccessToken fetches the OrgID from the access token which is available in context's auth information
func getOrgIDFromAccessToken(g *configapi.GlobalServerAuth) (string, error) {
	token, err := ParseToken(&oauth2.Token{AccessToken: g.AccessToken})
	if err != nil {
		return "", err
	}
	return token.OrgID, nil
}
