# \ShopListingPersonalizationAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteListingPersonalization**](ShopListingPersonalizationAPI.md#DeleteListingPersonalization) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/personalization | 
[**GetListingPersonalization**](ShopListingPersonalizationAPI.md#GetListingPersonalization) | **Get** /v3/application/listings/{listing_id}/personalization | 
[**UpdateListingPersonalization**](ShopListingPersonalizationAPI.md#UpdateListingPersonalization) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/personalization | 



## DeleteListingPersonalization

> DeleteListingPersonalization(ctx, shopId, listingId).Execute()





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
	listingId := int64(789) // int64 | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShopListingPersonalizationAPI.DeleteListingPersonalization(context.Background(), shopId, listingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingPersonalizationAPI.DeleteListingPersonalization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingPersonalizationRequest struct via the builder pattern


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


## GetListingPersonalization

> EtsyModulesListingPersonalizationApiResourcesOpenApiListingPersonalization GetListingPersonalization(ctx, listingId).Execute()





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
	listingId := int64(789) // int64 | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingPersonalizationAPI.GetListingPersonalization(context.Background(), listingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingPersonalizationAPI.GetListingPersonalization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingPersonalization`: EtsyModulesListingPersonalizationApiResourcesOpenApiListingPersonalization
	fmt.Fprintf(os.Stdout, "Response from `ShopListingPersonalizationAPI.GetListingPersonalization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingPersonalizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EtsyModulesListingPersonalizationApiResourcesOpenApiListingPersonalization**](EtsyModulesListingPersonalizationApiResourcesOpenApiListingPersonalization.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListingPersonalization

> EtsyModulesListingPersonalizationApiResourcesOpenApiListingPersonalization UpdateListingPersonalization(ctx, shopId, listingId).SupportsMultiplePersonalizationQuestions(supportsMultiplePersonalizationQuestions).UpdateListingPersonalizationRequest(updateListingPersonalizationRequest).Execute()





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
	listingId := int64(789) // int64 | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.
	supportsMultiplePersonalizationQuestions := true // bool | This query parameter indicates that the caller supports up to 5 personalization questions and the following question types: 'text_input', 'dropdown', 'unlabeled_upload', 'labeled_upload'. Sending this param without updating your application can lead to inadvertantly deleting seller-entered data. (optional)
	updateListingPersonalizationRequest := *openapiclient.NewUpdateListingPersonalizationRequest([]openapiclient.UpdateListingPersonalizationRequestPersonalizationQuestionsInner{*openapiclient.NewUpdateListingPersonalizationRequestPersonalizationQuestionsInner(NullableInt64(123), "QuestionText_example", "Instructions_example", openapiclient.updateListingPersonalization_request_personalization_questions_inner_question_type("text_input"), false, NullableInt64(123), NullableInt64(123), []openapiclient.UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner{*openapiclient.NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner(NullableInt64(123), "Label_example")})}) // UpdateListingPersonalizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingPersonalizationAPI.UpdateListingPersonalization(context.Background(), shopId, listingId).SupportsMultiplePersonalizationQuestions(supportsMultiplePersonalizationQuestions).UpdateListingPersonalizationRequest(updateListingPersonalizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingPersonalizationAPI.UpdateListingPersonalization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateListingPersonalization`: EtsyModulesListingPersonalizationApiResourcesOpenApiListingPersonalization
	fmt.Fprintf(os.Stdout, "Response from `ShopListingPersonalizationAPI.UpdateListingPersonalization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListingPersonalizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportsMultiplePersonalizationQuestions** | **bool** | This query parameter indicates that the caller supports up to 5 personalization questions and the following question types: &#39;text_input&#39;, &#39;dropdown&#39;, &#39;unlabeled_upload&#39;, &#39;labeled_upload&#39;. Sending this param without updating your application can lead to inadvertantly deleting seller-entered data. | 
 **updateListingPersonalizationRequest** | [**UpdateListingPersonalizationRequest**](UpdateListingPersonalizationRequest.md) |  | 

### Return type

[**EtsyModulesListingPersonalizationApiResourcesOpenApiListingPersonalization**](EtsyModulesListingPersonalizationApiResourcesOpenApiListingPersonalization.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

