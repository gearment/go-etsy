# \ShopAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindShops**](ShopAPI.md#FindShops) | **Get** /v3/application/shops | 
[**GetShop**](ShopAPI.md#GetShop) | **Get** /v3/application/shops/{shop_id} | 
[**GetShopByOwnerUserId**](ShopAPI.md#GetShopByOwnerUserId) | **Get** /v3/application/users/{user_id}/shops | 
[**UpdateShop**](ShopAPI.md#UpdateShop) | **Put** /v3/application/shops/{shop_id} | 



## FindShops

> Shops FindShops(ctx).ShopName(shopName).Limit(limit).Offset(offset).Execute()





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
	shopName := "shopName_example" // string | The shop's name string.
	limit := int64(56) // int64 | The maximum number of results to return. (optional) (default to 25)
	offset := int64(56) // int64 | The number of records to skip before selecting the first result. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopAPI.FindShops(context.Background()).ShopName(shopName).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.FindShops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindShops`: Shops
	fmt.Fprintf(os.Stdout, "Response from `ShopAPI.FindShops`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindShopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shopName** | **string** | The shop&#39;s name string. | 
 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]

### Return type

[**Shops**](Shops.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShop

> Shop GetShop(ctx, shopId).Execute()





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
	resp, r, err := apiClient.ShopAPI.GetShop(context.Background(), shopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.GetShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShop`: Shop
	fmt.Fprintf(os.Stdout, "Response from `ShopAPI.GetShop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Shop**](Shop.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopByOwnerUserId

> Shop GetShopByOwnerUserId(ctx, userId).Execute()





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
	userId := int64(56) // int64 | The numeric user ID of the [user](/documentation/reference#tag/User) who owns this shop.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopAPI.GetShopByOwnerUserId(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.GetShopByOwnerUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopByOwnerUserId`: Shop
	fmt.Fprintf(os.Stdout, "Response from `ShopAPI.GetShopByOwnerUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | The numeric user ID of the [user](/documentation/reference#tag/User) who owns this shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopByOwnerUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Shop**](Shop.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShop

> Shop UpdateShop(ctx, shopId).Title(title).Announcement(announcement).SaleMessage(saleMessage).DigitalSaleMessage(digitalSaleMessage).PolicyAdditional(policyAdditional).Execute()





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
	title := "title_example" // string | A brief heading string for the shop's main page. (optional)
	announcement := "announcement_example" // string | An announcement string to buyers that displays on the shop's homepage. (optional)
	saleMessage := "saleMessage_example" // string | A message string sent to users who complete a purchase from this shop. (optional)
	digitalSaleMessage := "digitalSaleMessage_example" // string | A message string sent to users who purchase a digital item from this shop. (optional)
	policyAdditional := "policyAdditional_example" // string | The shop's additional policies string (may be blank). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopAPI.UpdateShop(context.Background(), shopId).Title(title).Announcement(announcement).SaleMessage(saleMessage).DigitalSaleMessage(digitalSaleMessage).PolicyAdditional(policyAdditional).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopAPI.UpdateShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShop`: Shop
	fmt.Fprintf(os.Stdout, "Response from `ShopAPI.UpdateShop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** | A brief heading string for the shop&#39;s main page. | 
 **announcement** | **string** | An announcement string to buyers that displays on the shop&#39;s homepage. | 
 **saleMessage** | **string** | A message string sent to users who complete a purchase from this shop. | 
 **digitalSaleMessage** | **string** | A message string sent to users who purchase a digital item from this shop. | 
 **policyAdditional** | **string** | The shop&#39;s additional policies string (may be blank). | 

### Return type

[**Shop**](Shop.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

