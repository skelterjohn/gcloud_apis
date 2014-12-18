/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

type gcloudTransport struct {
	transport http.RoundTripper
}

func (gt gcloudTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "Cloud SDK Command Line Tool")
	req.Header.Add("User-Agent", "gcloud")
	req.Header.Add("User-Agent", "gcloud_apis")
	return gt.transport.RoundTrip(req)
}

/*
This auth function is a bootstrap until better integration with
gcloud auth can be provided. To use, set (and export) the environment
variable GCLOUD_APIS_REFRESH_TOKEN to a valid refresh token, which can
usually be acquired via $(gcloud auth print-refresh-token). In the
future, gcloud_apis will be able to fetch credential information from
the same place as gcloud, and will be usable without setting this
environment variable.
*/
func getAuthenticatedClient() (*http.Client, error) {

	conf := oauth2.Config{
		ClientID: "32555940559.apps.googleusercontent.com",
		ClientSecret: "ZmssLNjJy2998hD4CTg2ejr2",
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
		Scopes: []string{"https://www.googleapis.com/auth/cloud-platform"},
	}

	refreshToken := os.Getenv("GCLOUD_APIS_REFRESH_TOKEN")
	if refreshToken == "" {
		err := fmt.Errorf("no refresh token found in GCLOUD_APIS_REFRESH_TOKEN")
		return nil, err
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	tokenSource := conf.TokenSource(oauth2.NoContext, token)

	oauth2Transport := &oauth2.Transport{
		Source: tokenSource,
	}

	client := &http.Client{
		Transport: gcloudTransport{oauth2Transport},
	}

	return client, nil
}
