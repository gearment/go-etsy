/*
Etsy Open API v3

Testing ShopListingOfferingAPIService

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

func Test_goEtsy_ShopListingOfferingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShopListingOfferingAPIService GetListingOffering", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listingId int64
		var productId int64
		var productOfferingId int64

		resp, httpRes, err := apiClient.ShopListingOfferingAPI.GetListingOffering(context.Background(), listingId, productId, productOfferingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
