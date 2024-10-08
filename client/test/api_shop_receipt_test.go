/*
Etsy Open API v3

Testing ShopReceiptAPIService

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

func Test_goEtsy_ShopReceiptAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShopReceiptAPIService CreateReceiptShipment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var receiptId int64

		resp, httpRes, err := apiClient.ShopReceiptAPI.CreateReceiptShipment(context.Background(), shopId, receiptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopReceiptAPIService GetShopReceipt", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var receiptId int64

		resp, httpRes, err := apiClient.ShopReceiptAPI.GetShopReceipt(context.Background(), shopId, receiptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopReceiptAPIService GetShopReceipts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64

		resp, httpRes, err := apiClient.ShopReceiptAPI.GetShopReceipts(context.Background(), shopId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopReceiptAPIService UpdateShopReceipt", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shopId int64
		var receiptId int64

		resp, httpRes, err := apiClient.ShopReceiptAPI.UpdateShopReceipt(context.Background(), shopId, receiptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
