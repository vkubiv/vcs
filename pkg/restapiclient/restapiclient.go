/*
Copyright Avast Software. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package restapiclient

//go:generate mockgen -destination restapiclient_mocks_test.go -package restapiclient_test -source=restapiclient.go -mock_names httpClient=MockHttpClient

import (
	"context"
	"fmt"
	"net/http"
)

const (
	prepareClaimDataAuthEndpoint = "/issuer/interactions/prepare-claim-data-authz-request"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	hostURI string
	client  httpClient
}

func NewClient(
	hostURI string,
	client httpClient,
) *Client {
	return &Client{
		hostURI: hostURI,
		client:  client,
	}
}

func (c *Client) PrepareClaimDataAuthZ(
	ctx context.Context,
	req *PrepareClaimDataAuthorizationRequest,
) (*PrepareClaimDataAuthorizationResponse, error) {
	return sendInternal[PrepareClaimDataAuthorizationRequest, PrepareClaimDataAuthorizationResponse](
		ctx,
		c.client,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.hostURI, prepareClaimDataAuthEndpoint),
		req,
	)
}