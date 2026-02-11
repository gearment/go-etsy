# \ShopProcessingProfilesAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShopReadinessStateDefinition**](ShopProcessingProfilesAPI.md#CreateShopReadinessStateDefinition) | **Post** /v3/application/shops/{shop_id}/readiness-state-definitions | 
[**DeleteShopReadinessStateDefinition**](ShopProcessingProfilesAPI.md#DeleteShopReadinessStateDefinition) | **Delete** /v3/application/shops/{shop_id}/readiness-state-definitions/{readiness_state_definition_id} | 
[**GetShopReadinessStateDefinition**](ShopProcessingProfilesAPI.md#GetShopReadinessStateDefinition) | **Get** /v3/application/shops/{shop_id}/readiness-state-definitions/{readiness_state_definition_id} | 
[**GetShopReadinessStateDefinitions**](ShopProcessingProfilesAPI.md#GetShopReadinessStateDefinitions) | **Get** /v3/application/shops/{shop_id}/readiness-state-definitions | 
[**UpdateShopReadinessStateDefinition**](ShopProcessingProfilesAPI.md#UpdateShopReadinessStateDefinition) | **Put** /v3/application/shops/{shop_id}/readiness-state-definitions/{readiness_state_definition_id} | 



## CreateShopReadinessStateDefinition

> ShopProcessingProfile CreateShopReadinessStateDefinition(ctx, shopId).CreateShopReadinessStateDefinitionRequest(createShopReadinessStateDefinitionRequest).Execute()





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
	shopId := int64(789) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	createShopReadinessStateDefinitionRequest := *openapiclient.NewCreateShopReadinessStateDefinitionRequest(openapiclient.createShopReadinessStateDefinition_request_readiness_state("ready_to_ship"), int64(123), int64(123)) // CreateShopReadinessStateDefinitionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopProcessingProfilesAPI.CreateShopReadinessStateDefinition(context.Background(), shopId).CreateShopReadinessStateDefinitionRequest(createShopReadinessStateDefinitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopProcessingProfilesAPI.CreateShopReadinessStateDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShopReadinessStateDefinition`: ShopProcessingProfile
	fmt.Fprintf(os.Stdout, "Response from `ShopProcessingProfilesAPI.CreateShopReadinessStateDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShopReadinessStateDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createShopReadinessStateDefinitionRequest** | [**CreateShopReadinessStateDefinitionRequest**](CreateShopReadinessStateDefinitionRequest.md) |  | 

### Return type

[**ShopProcessingProfile**](ShopProcessingProfile.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShopReadinessStateDefinition

> DeleteShopReadinessStateDefinition(ctx, shopId, readinessStateDefinitionId).Execute()





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
	shopId := int64(789) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	readinessStateDefinitionId := int64(789) // int64 | The numeric ID of the [processing profile](/documentation/reference#operation/getShopReadinessStateDefinition) associated with the listing. Returned only when the listing is `active` and of type `physical`, and the endpoint is either shop-scoped (path contains `shop_id`) or a single-listing request such as `getListing`. For every other case this field can be null.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShopProcessingProfilesAPI.DeleteShopReadinessStateDefinition(context.Background(), shopId, readinessStateDefinitionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopProcessingProfilesAPI.DeleteShopReadinessStateDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**readinessStateDefinitionId** | **int64** | The numeric ID of the [processing profile](/documentation/reference#operation/getShopReadinessStateDefinition) associated with the listing. Returned only when the listing is &#x60;active&#x60; and of type &#x60;physical&#x60;, and the endpoint is either shop-scoped (path contains &#x60;shop_id&#x60;) or a single-listing request such as &#x60;getListing&#x60;. For every other case this field can be null. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShopReadinessStateDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopReadinessStateDefinition

> ShopProcessingProfile GetShopReadinessStateDefinition(ctx, shopId, readinessStateDefinitionId).Execute()





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
	shopId := int64(789) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	readinessStateDefinitionId := int64(789) // int64 | The numeric ID of the [processing profile](/documentation/reference#operation/getShopReadinessStateDefinition) associated with the listing. Returned only when the listing is `active` and of type `physical`, and the endpoint is either shop-scoped (path contains `shop_id`) or a single-listing request such as `getListing`. For every other case this field can be null.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopProcessingProfilesAPI.GetShopReadinessStateDefinition(context.Background(), shopId, readinessStateDefinitionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopProcessingProfilesAPI.GetShopReadinessStateDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopReadinessStateDefinition`: ShopProcessingProfile
	fmt.Fprintf(os.Stdout, "Response from `ShopProcessingProfilesAPI.GetShopReadinessStateDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**readinessStateDefinitionId** | **int64** | The numeric ID of the [processing profile](/documentation/reference#operation/getShopReadinessStateDefinition) associated with the listing. Returned only when the listing is &#x60;active&#x60; and of type &#x60;physical&#x60;, and the endpoint is either shop-scoped (path contains &#x60;shop_id&#x60;) or a single-listing request such as &#x60;getListing&#x60;. For every other case this field can be null. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopReadinessStateDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShopProcessingProfile**](ShopProcessingProfile.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopReadinessStateDefinitions

> ShopProcessingProfiles GetShopReadinessStateDefinitions(ctx, shopId).Limit(limit).Offset(offset).Execute()





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
	shopId := int64(789) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	limit := int64(56) // int64 | The maximum number of results to return. (optional) (default to 25)
	offset := int64(56) // int64 | The number of records to skip before selecting the first result. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopProcessingProfilesAPI.GetShopReadinessStateDefinitions(context.Background(), shopId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopProcessingProfilesAPI.GetShopReadinessStateDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopReadinessStateDefinitions`: ShopProcessingProfiles
	fmt.Fprintf(os.Stdout, "Response from `ShopProcessingProfilesAPI.GetShopReadinessStateDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopReadinessStateDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]

### Return type

[**ShopProcessingProfiles**](ShopProcessingProfiles.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShopReadinessStateDefinition

> ShopProcessingProfile UpdateShopReadinessStateDefinition(ctx, shopId, readinessStateDefinitionId).UpdateShopReadinessStateDefinitionRequest(updateShopReadinessStateDefinitionRequest).Execute()





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
	shopId := int64(789) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	readinessStateDefinitionId := int64(789) // int64 | The numeric ID of the [processing profile](/documentation/reference#operation/getShopReadinessStateDefinition) associated with the listing. Returned only when the listing is `active` and of type `physical`, and the endpoint is either shop-scoped (path contains `shop_id`) or a single-listing request such as `getListing`. For every other case this field can be null.
	updateShopReadinessStateDefinitionRequest := *openapiclient.NewUpdateShopReadinessStateDefinitionRequest() // UpdateShopReadinessStateDefinitionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopProcessingProfilesAPI.UpdateShopReadinessStateDefinition(context.Background(), shopId, readinessStateDefinitionId).UpdateShopReadinessStateDefinitionRequest(updateShopReadinessStateDefinitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopProcessingProfilesAPI.UpdateShopReadinessStateDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShopReadinessStateDefinition`: ShopProcessingProfile
	fmt.Fprintf(os.Stdout, "Response from `ShopProcessingProfilesAPI.UpdateShopReadinessStateDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**readinessStateDefinitionId** | **int64** | The numeric ID of the [processing profile](/documentation/reference#operation/getShopReadinessStateDefinition) associated with the listing. Returned only when the listing is &#x60;active&#x60; and of type &#x60;physical&#x60;, and the endpoint is either shop-scoped (path contains &#x60;shop_id&#x60;) or a single-listing request such as &#x60;getListing&#x60;. For every other case this field can be null. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShopReadinessStateDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateShopReadinessStateDefinitionRequest** | [**UpdateShopReadinessStateDefinitionRequest**](UpdateShopReadinessStateDefinitionRequest.md) |  | 

### Return type

[**ShopProcessingProfile**](ShopProcessingProfile.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

