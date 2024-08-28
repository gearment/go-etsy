# \ShopProductionPartnerAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShopProductionPartners**](ShopProductionPartnerAPI.md#GetShopProductionPartners) | **Get** /v3/application/shops/{shop_id}/production-partners | 



## GetShopProductionPartners

> ShopProductionPartners GetShopProductionPartners(ctx, shopId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopProductionPartnerAPI.GetShopProductionPartners(context.Background(), shopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopProductionPartnerAPI.GetShopProductionPartners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopProductionPartners`: ShopProductionPartners
	fmt.Fprintf(os.Stdout, "Response from `ShopProductionPartnerAPI.GetShopProductionPartners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopProductionPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShopProductionPartners**](ShopProductionPartners.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

