/*
Etsy Open API v3

Testing ShopShippingProfileAPIService

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

func Test_goEtsy_ShopShippingProfileAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShopShippingProfileAPIService CreateShopShippingProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.CreateShopShippingProfile(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService CreateShopShippingProfileDestination", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.CreateShopShippingProfileDestination(context.Background(), shopId, shippingProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService CreateShopShippingProfileUpgrade", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.CreateShopShippingProfileUpgrade(context.Background(), shopId, shippingProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService DeleteShopShippingProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64

		httpRes, err := apiClient.ShopShippingProfileAPI.DeleteShopShippingProfile(context.Background(), shopId, shippingProfileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService DeleteShopShippingProfileDestination", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64
		var shippingProfileDestinationId int64

		httpRes, err := apiClient.ShopShippingProfileAPI.DeleteShopShippingProfileDestination(context.Background(), shopId, shippingProfileId, shippingProfileDestinationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService DeleteShopShippingProfileUpgrade", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64
		var upgradeId int64

		httpRes, err := apiClient.ShopShippingProfileAPI.DeleteShopShippingProfileUpgrade(context.Background(), shopId, shippingProfileId, upgradeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService GetShippingCarriers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.GetShippingCarriers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService GetShopShippingProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.GetShopShippingProfile(context.Background(), shopId, shippingProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService GetShopShippingProfileDestinationsByShippingProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.GetShopShippingProfileDestinationsByShippingProfile(context.Background(), shopId, shippingProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService GetShopShippingProfileUpgrades", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.GetShopShippingProfileUpgrades(context.Background(), shopId, shippingProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService GetShopShippingProfiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.GetShopShippingProfiles(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService UpdateShopShippingProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.UpdateShopShippingProfile(context.Background(), shopId, shippingProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService UpdateShopShippingProfileDestination", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64
		var shippingProfileDestinationId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.UpdateShopShippingProfileDestination(context.Background(), shopId, shippingProfileId, shippingProfileDestinationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopShippingProfileAPIService UpdateShopShippingProfileUpgrade", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var shippingProfileId int64
		var upgradeId int64

		resp, httpRes, err := apiClient.ShopShippingProfileAPI.UpdateShopShippingProfileUpgrade(context.Background(), shopId, shippingProfileId, upgradeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
