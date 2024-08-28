/*
Etsy Open API v3

Testing UserAddressAPIService

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

func Test_goEtsy_UserAddressAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserAddressAPIService DeleteUserAddress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userAddressId int64

		httpRes, err := apiClient.UserAddressAPI.DeleteUserAddress(context.Background(), userAddressId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAddressAPIService GetUserAddress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userAddressId int64

		resp, httpRes, err := apiClient.UserAddressAPI.GetUserAddress(context.Background(), userAddressId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAddressAPIService GetUserAddresses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserAddressAPI.GetUserAddresses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
