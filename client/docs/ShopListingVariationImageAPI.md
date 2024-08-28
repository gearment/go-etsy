# \ShopListingVariationImageAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListingVariationImages**](ShopListingVariationImageAPI.md#GetListingVariationImages) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/variation-images | 
[**UpdateVariationImages**](ShopListingVariationImageAPI.md#UpdateVariationImages) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/variation-images | 



## GetListingVariationImages

> ListingVariationImages GetListingVariationImages(ctx, shopId, listingId).Execute()





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
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	listingId := int64(56) // int64 | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingVariationImageAPI.GetListingVariationImages(context.Background(), shopId, listingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingVariationImageAPI.GetListingVariationImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingVariationImages`: ListingVariationImages
	fmt.Fprintf(os.Stdout, "Response from `ShopListingVariationImageAPI.GetListingVariationImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingVariationImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListingVariationImages**](ListingVariationImages.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVariationImages

> ListingVariationImages UpdateVariationImages(ctx, shopId, listingId).UpdateVariationImagesRequest(updateVariationImagesRequest).Execute()





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
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	listingId := int64(56) // int64 | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.
	updateVariationImagesRequest := *openapiclient.NewUpdateVariationImagesRequest([]openapiclient.UpdateVariationImagesRequestVariationImagesInner{*openapiclient.NewUpdateVariationImagesRequestVariationImagesInner(int64(123), int64(123), int64(123))}) // UpdateVariationImagesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingVariationImageAPI.UpdateVariationImages(context.Background(), shopId, listingId).UpdateVariationImagesRequest(updateVariationImagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingVariationImageAPI.UpdateVariationImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVariationImages`: ListingVariationImages
	fmt.Fprintf(os.Stdout, "Response from `ShopListingVariationImageAPI.UpdateVariationImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVariationImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVariationImagesRequest** | [**UpdateVariationImagesRequest**](UpdateVariationImagesRequest.md) |  | 

### Return type

[**ListingVariationImages**](ListingVariationImages.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

