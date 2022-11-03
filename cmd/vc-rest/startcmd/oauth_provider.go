/*
Copyright Avast Software. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package startcmd

import (
	"context"
	"errors"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	fositeoauth2 "github.com/ory/fosite/handler/oauth2"
	"github.com/ory/fosite/token/hmac"

	"github.com/trustbloc/vcs/component/oidc/fositemongo"
	"github.com/trustbloc/vcs/pkg/storage/mongodb"
)

func bootstrapOAuthProvider(
	ctx context.Context,
	secret string,
	mongoClient *mongodb.Client,
) (fosite.OAuth2Provider, error) {
	if len(secret) == 0 {
		return nil, errors.New("invalid secret")
	}

	config := new(fosite.Config)
	config.GlobalSecret = []byte(secret)
	config.AuthorizeCodeLifespan = 30 * time.Minute
	config.AccessTokenLifespan = 30 * time.Minute

	var hmacStrategy = &fositeoauth2.HMACSHAStrategy{
		Enigma: &hmac.HMACStrategy{
			Config: config,
		},
		Config: config,
	}

	store, err := fositemongo.NewStore(ctx, mongoClient)
	if err != nil {
		return nil, err
	}

	return compose.Compose(config, store, hmacStrategy,
		compose.OAuth2AuthorizeExplicitFactory,
		compose.OAuth2PKCEFactory,
		compose.PushedAuthorizeHandlerFactory,
	), nil
}