# \ShopReturnPolicyAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsolidateShopReturnPolicies**](ShopReturnPolicyAPI.md#ConsolidateShopReturnPolicies) | **Post** /v3/application/shops/{shop_id}/policies/return/consolidate | 
[**CreateShopReturnPolicy**](ShopReturnPolicyAPI.md#CreateShopReturnPolicy) | **Post** /v3/application/shops/{shop_id}/policies/return | 
[**DeleteShopReturnPolicy**](ShopReturnPolicyAPI.md#DeleteShopReturnPolicy) | **Delete** /v3/application/shops/{shop_id}/policies/return/{return_policy_id} | 
[**GetShopReturnPolicies**](ShopReturnPolicyAPI.md#GetShopReturnPolicies) | **Get** /v3/application/shops/{shop_id}/policies/return | 
[**GetShopReturnPolicy**](ShopReturnPolicyAPI.md#GetShopReturnPolicy) | **Get** /v3/application/shops/{shop_id}/policies/return/{return_policy_id} | 
[**UpdateShopReturnPolicy**](ShopReturnPolicyAPI.md#UpdateShopReturnPolicy) | **Put** /v3/application/shops/{shop_id}/policies/return/{return_policy_id} | 



## ConsolidateShopReturnPolicies

> ShopReturnPolicy ConsolidateShopReturnPolicies(ctx, shopId).ConsolidateShopReturnPoliciesRequest(consolidateShopReturnPoliciesRequest).Execute()





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
	consolidateShopReturnPoliciesRequest := *openapiclient.NewConsolidateShopReturnPoliciesRequest(int64(123), int64(123)) // ConsolidateShopReturnPoliciesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopReturnPolicyAPI.ConsolidateShopReturnPolicies(context.Background(), shopId).ConsolidateShopReturnPoliciesRequest(consolidateShopReturnPoliciesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReturnPolicyAPI.ConsolidateShopReturnPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsolidateShopReturnPolicies`: ShopReturnPolicy
	fmt.Fprintf(os.Stdout, "Response from `ShopReturnPolicyAPI.ConsolidateShopReturnPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsolidateShopReturnPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consolidateShopReturnPoliciesRequest** | [**ConsolidateShopReturnPoliciesRequest**](ConsolidateShopReturnPoliciesRequest.md) |  | 

### Return type

[**ShopReturnPolicy**](ShopReturnPolicy.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShopReturnPolicy

> ShopReturnPolicy CreateShopReturnPolicy(ctx, shopId).CreateShopReturnPolicyRequest(createShopReturnPolicyRequest).Execute()





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
	createShopReturnPolicyRequest := *openapiclient.NewCreateShopReturnPolicyRequest(false, false) // CreateShopReturnPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopReturnPolicyAPI.CreateShopReturnPolicy(context.Background(), shopId).CreateShopReturnPolicyRequest(createShopReturnPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReturnPolicyAPI.CreateShopReturnPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShopReturnPolicy`: ShopReturnPolicy
	fmt.Fprintf(os.Stdout, "Response from `ShopReturnPolicyAPI.CreateShopReturnPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShopReturnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createShopReturnPolicyRequest** | [**CreateShopReturnPolicyRequest**](CreateShopReturnPolicyRequest.md) |  | 

### Return type

[**ShopReturnPolicy**](ShopReturnPolicy.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShopReturnPolicy

> DeleteShopReturnPolicy(ctx, shopId, returnPolicyId).Execute()





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
	returnPolicyId := int64(56) // int64 | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShopReturnPolicyAPI.DeleteShopReturnPolicy(context.Background(), shopId, returnPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReturnPolicyAPI.DeleteShopReturnPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**returnPolicyId** | **int64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShopReturnPolicyRequest struct via the builder pattern


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


## GetShopReturnPolicies

> ShopReturnPolicies GetShopReturnPolicies(ctx, shopId).Execute()





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
	resp, r, err := apiClient.ShopReturnPolicyAPI.GetShopReturnPolicies(context.Background(), shopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReturnPolicyAPI.GetShopReturnPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopReturnPolicies`: ShopReturnPolicies
	fmt.Fprintf(os.Stdout, "Response from `ShopReturnPolicyAPI.GetShopReturnPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopReturnPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShopReturnPolicies**](ShopReturnPolicies.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopReturnPolicy

> ShopReturnPolicy GetShopReturnPolicy(ctx, shopId, returnPolicyId).Execute()





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
	returnPolicyId := int64(56) // int64 | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopReturnPolicyAPI.GetShopReturnPolicy(context.Background(), shopId, returnPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReturnPolicyAPI.GetShopReturnPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopReturnPolicy`: ShopReturnPolicy
	fmt.Fprintf(os.Stdout, "Response from `ShopReturnPolicyAPI.GetShopReturnPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**returnPolicyId** | **int64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopReturnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShopReturnPolicy**](ShopReturnPolicy.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShopReturnPolicy

> ShopReturnPolicy UpdateShopReturnPolicy(ctx, shopId, returnPolicyId).CreateShopReturnPolicyRequest(createShopReturnPolicyRequest).Execute()





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
	returnPolicyId := int64(56) // int64 | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies).
	createShopReturnPolicyRequest := *openapiclient.NewCreateShopReturnPolicyRequest(false, false) // CreateShopReturnPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopReturnPolicyAPI.UpdateShopReturnPolicy(context.Background(), shopId, returnPolicyId).CreateShopReturnPolicyRequest(createShopReturnPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReturnPolicyAPI.UpdateShopReturnPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShopReturnPolicy`: ShopReturnPolicy
	fmt.Fprintf(os.Stdout, "Response from `ShopReturnPolicyAPI.UpdateShopReturnPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**returnPolicyId** | **int64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShopReturnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createShopReturnPolicyRequest** | [**CreateShopReturnPolicyRequest**](CreateShopReturnPolicyRequest.md) |  | 

### Return type

[**ShopReturnPolicy**](ShopReturnPolicy.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

