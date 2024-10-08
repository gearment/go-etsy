/*
Etsy Open API v3

Testing OtherAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package goEtsy

import (
	"context"
	"testing"

	openapiclient "github.com/gearment/go-etsy/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_goEtsy_OtherAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OtherAPIService Ping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OtherAPI.Ping(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OtherAPIService TokenScopes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OtherAPI.TokenScopes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
