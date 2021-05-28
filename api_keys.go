package rockset

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// CreateAPIKey creates a new API key.
// An admin can create an api key for another user with option.ForUser().
//
// REST API documentation https://docs.rockset.com/rest-api/#createapikey
func (rc *RockClient) CreateAPIKey(ctx context.Context, options ...option.APIKeyOption) (openapi.ApiKey, error) {
	var err error
	var resp openapi.CreateApiKeyResponse

	opts := option.APIKeyOptions{}
	for _, o := range options {
		o(&opts)
	}

	// create for current user
	if opts.User == nil {
		getReq := rc.APIKeysApi.CreateApiKey(ctx)

		err = rc.Retry(ctx, func() error {
			resp, _, err = getReq.Execute()
			return err
		})

		if err != nil {
			return openapi.ApiKey{}, err
		}

		return *resp.Data, nil
	}

	// an admin can create for another user
	getReq := rc.APIKeysApi.CreateApiKeyAdmin(ctx, *opts.User)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return openapi.ApiKey{}, err
	}

	return *resp.Data, nil
}

// DeleteAPIKey deletes an API key.
// An admin can delete an api key for another user with option.ForUser().
//
// REST API documentation https://docs.rockset.com/rest-api/#deleteapikey
func (rc *RockClient) DeleteAPIKey(ctx context.Context, keyName string, options ...option.APIKeyOption) error {
	var err error

	opts := option.APIKeyOptions{}
	for _, o := range options {
		o(&opts)
	}

	// delete for current user
	if opts.User == nil {
		getReq := rc.APIKeysApi.DeleteApiKey(ctx, keyName)

		err = rc.Retry(ctx, func() error {
			_, _, err = getReq.Execute()
			return err
		})

		if err != nil {
			return err
		}

		return nil
	}

	// an admin can delete for another user
	getReq := rc.APIKeysApi.DeleteApiKeyAdmin(ctx, keyName, *opts.User)

	err = rc.Retry(ctx, func() error {
		_, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

// ListAPIKeys list API keys.
// An admin can list api keys for another user with option.ForUser().
//
// REST API documentation https://docs.rockset.com/rest-api/#listapikey
func (rc *RockClient) ListAPIKeys(ctx context.Context, options ...option.APIKeyOption) ([]openapi.ApiKey, error) {
	var err error
	var resp openapi.ListApiKeysResponse

	opts := option.APIKeyOptions{}
	for _, o := range options {
		o(&opts)
	}

	// list for current user
	if opts.User == nil {
		getReq := rc.APIKeysApi.ListApiKeys(ctx)

		err = rc.Retry(ctx, func() error {
			resp, _, err = getReq.Execute()
			return err
		})

		if err != nil {
			return nil, err
		}

		return *resp.Data, nil
	}

	// an admin can list for another user
	getReq := rc.APIKeysApi.ListApiKeysAdmin(ctx, *opts.User)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return nil, err
	}

	return *resp.Data, nil
}
