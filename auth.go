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
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/oauth2"
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

type cliTokenSource struct {
}

func (cts cliTokenSource) Token() (*oauth2.Token, error) {
	output, err := exec.Command("gcloud", "auth", "print-access-token").Output()
	if err != nil {
		return nil, fmt.Errorf("problem getting gcloud token: %v", err)
	}
	return &oauth2.Token{
		AccessToken: string(output),
	}, nil
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
		ClientID:     "32555940559.apps.googleusercontent.com",
		ClientSecret: "ZmssLNjJy2998hD4CTg2ejr2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
		Scopes: []string{"https://www.googleapis.com/auth/cloud-platform"},
	}

	var ts oauth2.TokenSource

	refreshToken := os.Getenv("GCLOUD_APIS_REFRESH_TOKEN")
	if refreshToken == "" {
		ts = cliTokenSource{}
	} else {
		token := &oauth2.Token{
			RefreshToken: refreshToken,
		}
		ts = conf.TokenSource(oauth2.NoContext, token)
	}

	oauth2Transport := &oauth2.Transport{
		Source: ts,
	}

	client := &http.Client{
		Transport: gcloudTransport{oauth2Transport},
	}

	return client, nil
}
