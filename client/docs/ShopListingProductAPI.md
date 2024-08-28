# \ShopListingProductAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListingProduct**](ShopListingProductAPI.md#GetListingProduct) | **Get** /v3/application/listings/{listing_id}/inventory/products/{product_id} | 



## GetListingProduct

> ListingInventoryProduct GetListingProduct(ctx, listingId, productId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	listingId := int64(56) // int64 | The listing to return a ListingProduct for.
	productId := int64(56) // int64 | The numeric ID for a specific [product](/documentation/reference#tag/ShopListing-Product) purchased from a listing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingProductAPI.GetListingProduct(context.Background(), listingId, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingProductAPI.GetListingProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingProduct`: ListingInventoryProduct
	fmt.Fprintf(os.Stdout, "Response from `ShopListingProductAPI.GetListingProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | **int64** | The listing to return a ListingProduct for. | 
**productId** | **int64** | The numeric ID for a specific [product](/documentation/reference#tag/ShopListing-Product) purchased from a listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListingInventoryProduct**](ListingInventoryProduct.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

