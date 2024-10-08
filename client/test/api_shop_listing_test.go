/*
Etsy Open API v3

Testing ShopListingAPIService

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

func Test_goEtsy_ShopListingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShopListingAPIService CreateDraftListing", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64

		resp, httpRes, err := apiClient.ShopListingAPI.CreateDraftListing(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService DeleteListing", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listingId int64

		httpRes, err := apiClient.ShopListingAPI.DeleteListing(context.Background(), listingId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService DeleteListingProperty", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var listingId int64
		var propertyId int64

		httpRes, err := apiClient.ShopListingAPI.DeleteListingProperty(context.Background(), shopId, listingId, propertyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService FindAllActiveListingsByShop", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64

		resp, httpRes, err := apiClient.ShopListingAPI.FindAllActiveListingsByShop(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService FindAllListingsActive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ShopListingAPI.FindAllListingsActive(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService GetFeaturedListingsByShop", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64

		resp, httpRes, err := apiClient.ShopListingAPI.GetFeaturedListingsByShop(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService GetListing", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listingId int64

		resp, httpRes, err := apiClient.ShopListingAPI.GetListing(context.Background(), listingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService GetListingProperties", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var listingId int64

		resp, httpRes, err := apiClient.ShopListingAPI.GetListingProperties(context.Background(), shopId, listingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService GetListingProperty", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listingId int64
		var propertyId int64

		resp, httpRes, err := apiClient.ShopListingAPI.GetListingProperty(context.Background(), listingId, propertyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService GetListingsByListingIds", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ShopListingAPI.GetListingsByListingIds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService GetListingsByShop", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64

		resp, httpRes, err := apiClient.ShopListingAPI.GetListingsByShop(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService GetListingsByShopReceipt", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var receiptId int64
		var shopId int64

		resp, httpRes, err := apiClient.ShopListingAPI.GetListingsByShopReceipt(context.Background(), receiptId, shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService GetListingsByShopReturnPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var returnPolicyId int64
		var shopId int64

		resp, httpRes, err := apiClient.ShopListingAPI.GetListingsByShopReturnPolicy(context.Background(), returnPolicyId, shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService GetListingsByShopSectionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64

		resp, httpRes, err := apiClient.ShopListingAPI.GetListingsByShopSectionId(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService UpdateListing", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var listingId int64

		resp, httpRes, err := apiClient.ShopListingAPI.UpdateListing(context.Background(), shopId, listingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService UpdateListingDeprecated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var listingId int64

		resp, httpRes, err := apiClient.ShopListingAPI.UpdateListingDeprecated(context.Background(), shopId, listingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopListingAPIService UpdateListingProperty", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var listingId int64
		var propertyId int64

		resp, httpRes, err := apiClient.ShopListingAPI.UpdateListingProperty(context.Background(), shopId, listingId, propertyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
