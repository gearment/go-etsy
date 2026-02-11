# \ShopListingAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDraftListing**](ShopListingAPI.md#CreateDraftListing) | **Post** /v3/application/shops/{shop_id}/listings | 
[**DeleteListing**](ShopListingAPI.md#DeleteListing) | **Delete** /v3/application/listings/{listing_id} | 
[**DeleteListingProperty**](ShopListingAPI.md#DeleteListingProperty) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/properties/{property_id} | 
[**FindAllActiveListingsByShop**](ShopListingAPI.md#FindAllActiveListingsByShop) | **Get** /v3/application/shops/{shop_id}/listings/active | 
[**FindAllListingsActive**](ShopListingAPI.md#FindAllListingsActive) | **Get** /v3/application/listings/active | 
[**GetFeaturedListingsByShop**](ShopListingAPI.md#GetFeaturedListingsByShop) | **Get** /v3/application/shops/{shop_id}/listings/featured | 
[**GetListing**](ShopListingAPI.md#GetListing) | **Get** /v3/application/listings/{listing_id} | 
[**GetListingProperties**](ShopListingAPI.md#GetListingProperties) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/properties | 
[**GetListingProperty**](ShopListingAPI.md#GetListingProperty) | **Get** /v3/application/listings/{listing_id}/properties/{property_id} | 
[**GetListingsByListingIds**](ShopListingAPI.md#GetListingsByListingIds) | **Get** /v3/application/listings/batch | 
[**GetListingsByShop**](ShopListingAPI.md#GetListingsByShop) | **Get** /v3/application/shops/{shop_id}/listings | 
[**GetListingsByShopReceipt**](ShopListingAPI.md#GetListingsByShopReceipt) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id}/listings | 
[**GetListingsByShopReturnPolicy**](ShopListingAPI.md#GetListingsByShopReturnPolicy) | **Get** /v3/application/shops/{shop_id}/policies/return/{return_policy_id}/listings | 
[**GetListingsByShopSectionId**](ShopListingAPI.md#GetListingsByShopSectionId) | **Get** /v3/application/shops/{shop_id}/shop-sections/listings | 
[**UpdateListing**](ShopListingAPI.md#UpdateListing) | **Patch** /v3/application/shops/{shop_id}/listings/{listing_id} | 
[**UpdateListingProperty**](ShopListingAPI.md#UpdateListingProperty) | **Put** /v3/application/shops/{shop_id}/listings/{listing_id}/properties/{property_id} | 



## CreateDraftListing

> ShopListing CreateDraftListing(ctx, shopId).Legacy(legacy).CreateDraftListingRequest(createDraftListingRequest).Execute()





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
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)
	createDraftListingRequest := *openapiclient.NewCreateDraftListingRequest(int64(123), "Title_example", "Description_example", float32(123), openapiclient.createDraftListing_request_who_made("i_did"), openapiclient.createDraftListing_request_when_made("made_to_order"), int64(123)) // CreateDraftListingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.CreateDraftListing(context.Background(), shopId).Legacy(legacy).CreateDraftListingRequest(createDraftListingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.CreateDraftListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDraftListing`: ShopListing
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.CreateDraftListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDraftListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 
 **createDraftListingRequest** | [**CreateDraftListingRequest**](CreateDraftListingRequest.md) |  | 

### Return type

[**ShopListing**](ShopListing.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListing

> DeleteListing(ctx, listingId).Execute()





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
	r, err := apiClient.ShopListingAPI.DeleteListing(context.Background(), listingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.DeleteListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingRequest struct via the builder pattern


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


## DeleteListingProperty

> DeleteListingProperty(ctx, shopId, listingId, propertyId).Execute()





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
	propertyId := int64(789) // int64 | The unique ID of an Etsy [listing property](/documentation/reference#operation/getListingProperties).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShopListingAPI.DeleteListingProperty(context.Background(), shopId, listingId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.DeleteListingProperty``: %v\n", err)
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
**propertyId** | **int64** | The unique ID of an Etsy [listing property](/documentation/reference#operation/getListingProperties). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingPropertyRequest struct via the builder pattern


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


## FindAllActiveListingsByShop

> ShopListings FindAllActiveListingsByShop(ctx, shopId).Limit(limit).SortOn(sortOn).SortOrder(sortOrder).Offset(offset).Keywords(keywords).Legacy(legacy).Execute()





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
	sortOn := openapiclient.getListingsByShop_sort_on_parameter("created") // GetListingsByShopSortOnParameter | The value to sort a search result of listings on. NOTES: a) `sort_on` only works when combined with one of the search options (keywords, region, etc.). b) when using `score` the returned results will always be in _descending_ order, regardless of the `sort_order` parameter. (optional) (default to "created")
	sortOrder := openapiclient.getListingsByShop_sort_order_parameter("asc") // GetListingsByShopSortOrderParameter | The ascending(up) or descending(down) order to sort listings by. NOTE: sort_order only works when combined with one of the search options (keywords, region, etc.). (optional) (default to "desc")
	offset := int64(56) // int64 | The number of records to skip before selecting the first result. (optional) (default to 0)
	keywords := "keywords_example" // string | Search term or phrase that must appear in all results. (optional)
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.FindAllActiveListingsByShop(context.Background(), shopId).Limit(limit).SortOn(sortOn).SortOrder(sortOrder).Offset(offset).Keywords(keywords).Legacy(legacy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.FindAllActiveListingsByShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindAllActiveListingsByShop`: ShopListings
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.FindAllActiveListingsByShop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllActiveListingsByShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **sortOn** | [**GetListingsByShopSortOnParameter**](GetListingsByShopSortOnParameter.md) | The value to sort a search result of listings on. NOTES: a) &#x60;sort_on&#x60; only works when combined with one of the search options (keywords, region, etc.). b) when using &#x60;score&#x60; the returned results will always be in _descending_ order, regardless of the &#x60;sort_order&#x60; parameter. | [default to &quot;created&quot;]
 **sortOrder** | [**GetListingsByShopSortOrderParameter**](GetListingsByShopSortOrderParameter.md) | The ascending(up) or descending(down) order to sort listings by. NOTE: sort_order only works when combined with one of the search options (keywords, region, etc.). | [default to &quot;desc&quot;]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]
 **keywords** | **string** | Search term or phrase that must appear in all results. | 
 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 

### Return type

[**ShopListings**](ShopListings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllListingsActive

> ShopListings FindAllListingsActive(ctx).Limit(limit).Offset(offset).Keywords(keywords).SortOn(sortOn).SortOrder(sortOrder).MinPrice(minPrice).MaxPrice(maxPrice).TaxonomyId(taxonomyId).ShopLocation(shopLocation).Legacy(legacy).Execute()





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
	limit := int64(56) // int64 | The maximum number of results to return. (optional) (default to 25)
	offset := int64(56) // int64 | The number of records to skip before selecting the first result. (optional) (default to 0)
	keywords := "keywords_example" // string | Search term or phrase that must appear in all results. (optional)
	sortOn := openapiclient.getListingsByShop_sort_on_parameter("created") // GetListingsByShopSortOnParameter | The value to sort a search result of listings on. NOTES: a) `sort_on` only works when combined with one of the search options (keywords, region, etc.). b) when using `score` the returned results will always be in _descending_ order, regardless of the `sort_order` parameter. (optional) (default to "created")
	sortOrder := openapiclient.getListingsByShop_sort_order_parameter("asc") // GetListingsByShopSortOrderParameter | The ascending(up) or descending(down) order to sort listings by. NOTE: sort_order only works when combined with one of the search options (keywords, region, etc.). (optional) (default to "desc")
	minPrice := float32(3.4) // float32 | The minimum price of listings to be returned by a search result. (optional)
	maxPrice := float32(3.4) // float32 | The maximum price of listings to be returned by a search result. (optional)
	taxonomyId := int64(789) // int64 | The numerical taxonomy ID of the listing. See [SellerTaxonomy](/documentation/reference#tag/SellerTaxonomy) and [BuyerTaxonomy](/documentation/reference#tag/BuyerTaxonomy) for more information. (optional)
	shopLocation := "shopLocation_example" // string | Filters by shop location. If location cannot be parsed, Etsy responds with an error. (optional)
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.FindAllListingsActive(context.Background()).Limit(limit).Offset(offset).Keywords(keywords).SortOn(sortOn).SortOrder(sortOrder).MinPrice(minPrice).MaxPrice(maxPrice).TaxonomyId(taxonomyId).ShopLocation(shopLocation).Legacy(legacy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.FindAllListingsActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindAllListingsActive`: ShopListings
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.FindAllListingsActive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllListingsActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]
 **keywords** | **string** | Search term or phrase that must appear in all results. | 
 **sortOn** | [**GetListingsByShopSortOnParameter**](GetListingsByShopSortOnParameter.md) | The value to sort a search result of listings on. NOTES: a) &#x60;sort_on&#x60; only works when combined with one of the search options (keywords, region, etc.). b) when using &#x60;score&#x60; the returned results will always be in _descending_ order, regardless of the &#x60;sort_order&#x60; parameter. | [default to &quot;created&quot;]
 **sortOrder** | [**GetListingsByShopSortOrderParameter**](GetListingsByShopSortOrderParameter.md) | The ascending(up) or descending(down) order to sort listings by. NOTE: sort_order only works when combined with one of the search options (keywords, region, etc.). | [default to &quot;desc&quot;]
 **minPrice** | **float32** | The minimum price of listings to be returned by a search result. | 
 **maxPrice** | **float32** | The maximum price of listings to be returned by a search result. | 
 **taxonomyId** | **int64** | The numerical taxonomy ID of the listing. See [SellerTaxonomy](/documentation/reference#tag/SellerTaxonomy) and [BuyerTaxonomy](/documentation/reference#tag/BuyerTaxonomy) for more information. | 
 **shopLocation** | **string** | Filters by shop location. If location cannot be parsed, Etsy responds with an error. | 
 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 

### Return type

[**ShopListings**](ShopListings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeaturedListingsByShop

> ShopListings GetFeaturedListingsByShop(ctx, shopId).Limit(limit).Offset(offset).Legacy(legacy).Execute()





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
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.GetFeaturedListingsByShop(context.Background(), shopId).Limit(limit).Offset(offset).Legacy(legacy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.GetFeaturedListingsByShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeaturedListingsByShop`: ShopListings
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.GetFeaturedListingsByShop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturedListingsByShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]
 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 

### Return type

[**ShopListings**](ShopListings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListing

> ShopListingWithAssociations GetListing(ctx, listingId).Includes(includes).Language(language).Legacy(legacy).AllowSuggestedTitle(allowSuggestedTitle).Execute()





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
	includes := []openapiclient.GetListingsByShopIncludesParameterInner{openapiclient.getListingsByShop_includes_parameter_inner("Shipping")} // []GetListingsByShopIncludesParameterInner | An enumerated string that attaches a valid association. Acceptable inputs are 'Shipping', 'Shop', 'Images', 'User', 'Translations', 'Videos', 'Inventory' and 'Personalization'. (optional)
	language := "language_example" // string | The IETF language tag for the language of this translation. Ex: `de`, `en`, `es`, `fr`, `it`, `ja`, `nl`, `pl`, `pt`. (optional)
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)
	allowSuggestedTitle := true // bool | This parameter will include in the response a suggested title for the listing, if one is available. Since suggestions are only available to the listing's owner, client must submit an oauth_access_token scoped to the owner of the listing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.GetListing(context.Background(), listingId).Includes(includes).Language(language).Legacy(legacy).AllowSuggestedTitle(allowSuggestedTitle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.GetListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListing`: ShopListingWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.GetListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includes** | [**[]GetListingsByShopIncludesParameterInner**](GetListingsByShopIncludesParameterInner.md) | An enumerated string that attaches a valid association. Acceptable inputs are &#39;Shipping&#39;, &#39;Shop&#39;, &#39;Images&#39;, &#39;User&#39;, &#39;Translations&#39;, &#39;Videos&#39;, &#39;Inventory&#39; and &#39;Personalization&#39;. | 
 **language** | **string** | The IETF language tag for the language of this translation. Ex: &#x60;de&#x60;, &#x60;en&#x60;, &#x60;es&#x60;, &#x60;fr&#x60;, &#x60;it&#x60;, &#x60;ja&#x60;, &#x60;nl&#x60;, &#x60;pl&#x60;, &#x60;pt&#x60;. | 
 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 
 **allowSuggestedTitle** | **bool** | This parameter will include in the response a suggested title for the listing, if one is available. Since suggestions are only available to the listing&#39;s owner, client must submit an oauth_access_token scoped to the owner of the listing. | 

### Return type

[**ShopListingWithAssociations**](ShopListingWithAssociations.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingProperties

> ListingPropertyValues GetListingProperties(ctx, shopId, listingId).Execute()





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
	resp, r, err := apiClient.ShopListingAPI.GetListingProperties(context.Background(), shopId, listingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.GetListingProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingProperties`: ListingPropertyValues
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.GetListingProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListingPropertyValues**](ListingPropertyValues.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingProperty

> ListingPropertyValue GetListingProperty(ctx, listingId, propertyId).Execute()





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
	propertyId := int64(789) // int64 | The unique ID of an Etsy [listing property](/documentation/reference#operation/getListingProperties).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.GetListingProperty(context.Background(), listingId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.GetListingProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingProperty`: ListingPropertyValue
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.GetListingProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 
**propertyId** | **int64** | The unique ID of an Etsy [listing property](/documentation/reference#operation/getListingProperties). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListingPropertyValue**](ListingPropertyValue.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingsByListingIds

> ShopListingsWithAssociations GetListingsByListingIds(ctx).ListingIds(listingIds).Includes(includes).Legacy(legacy).Execute()





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
	listingIds := []int64{int64(123)} // []int64 | The list of numeric IDS for the listings in a specific Etsy shop.
	includes := []openapiclient.GetListingsByShopIncludesParameterInner{openapiclient.getListingsByShop_includes_parameter_inner("Shipping")} // []GetListingsByShopIncludesParameterInner | An enumerated string that attaches a valid association. Acceptable inputs are 'Shipping', 'Shop', 'Images', 'User', 'Translations', 'Videos', 'Inventory' and 'Personalization'. (optional)
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.GetListingsByListingIds(context.Background()).ListingIds(listingIds).Includes(includes).Legacy(legacy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.GetListingsByListingIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingsByListingIds`: ShopListingsWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.GetListingsByListingIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsByListingIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listingIds** | **[]int64** | The list of numeric IDS for the listings in a specific Etsy shop. | 
 **includes** | [**[]GetListingsByShopIncludesParameterInner**](GetListingsByShopIncludesParameterInner.md) | An enumerated string that attaches a valid association. Acceptable inputs are &#39;Shipping&#39;, &#39;Shop&#39;, &#39;Images&#39;, &#39;User&#39;, &#39;Translations&#39;, &#39;Videos&#39;, &#39;Inventory&#39; and &#39;Personalization&#39;. | 
 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 

### Return type

[**ShopListingsWithAssociations**](ShopListingsWithAssociations.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingsByShop

> ShopListingsWithAssociations GetListingsByShop(ctx, shopId).State(state).Limit(limit).Offset(offset).SortOn(sortOn).SortOrder(sortOrder).Includes(includes).Legacy(legacy).Execute()





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
	state := openapiclient.getListingsByShop_state_parameter("active") // GetListingsByShopStateParameter | When _updating_ a listing, this value can be either `active` or `inactive`. Note: Setting a `draft` listing to `active` will also publish the listing on etsy.com and requires that the listing have an image set. Setting a `sold_out` listing to active will update the quantity to 1 and renew the listing on etsy.com. (optional) (default to "active")
	limit := int64(56) // int64 | The maximum number of results to return. (optional) (default to 25)
	offset := int64(56) // int64 | The number of records to skip before selecting the first result. (optional) (default to 0)
	sortOn := openapiclient.getListingsByShop_sort_on_parameter("created") // GetListingsByShopSortOnParameter | The value to sort a search result of listings on. NOTES: a) `sort_on` only works when combined with one of the search options (keywords, region, etc.). b) when using `score` the returned results will always be in _descending_ order, regardless of the `sort_order` parameter. (optional) (default to "created")
	sortOrder := openapiclient.getListingsByShop_sort_order_parameter("asc") // GetListingsByShopSortOrderParameter | The ascending(up) or descending(down) order to sort listings by. NOTE: sort_order only works when combined with one of the search options (keywords, region, etc.). (optional) (default to "desc")
	includes := []openapiclient.GetListingsByShopIncludesParameterInner{openapiclient.getListingsByShop_includes_parameter_inner("Shipping")} // []GetListingsByShopIncludesParameterInner | An enumerated string that attaches a valid association. Acceptable inputs are 'Shipping', 'Shop', 'Images', 'User', 'Translations', 'Videos', 'Inventory' and 'Personalization'. (optional)
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.GetListingsByShop(context.Background(), shopId).State(state).Limit(limit).Offset(offset).SortOn(sortOn).SortOrder(sortOrder).Includes(includes).Legacy(legacy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.GetListingsByShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingsByShop`: ShopListingsWithAssociations
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.GetListingsByShop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsByShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | [**GetListingsByShopStateParameter**](GetListingsByShopStateParameter.md) | When _updating_ a listing, this value can be either &#x60;active&#x60; or &#x60;inactive&#x60;. Note: Setting a &#x60;draft&#x60; listing to &#x60;active&#x60; will also publish the listing on etsy.com and requires that the listing have an image set. Setting a &#x60;sold_out&#x60; listing to active will update the quantity to 1 and renew the listing on etsy.com. | [default to &quot;active&quot;]
 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]
 **sortOn** | [**GetListingsByShopSortOnParameter**](GetListingsByShopSortOnParameter.md) | The value to sort a search result of listings on. NOTES: a) &#x60;sort_on&#x60; only works when combined with one of the search options (keywords, region, etc.). b) when using &#x60;score&#x60; the returned results will always be in _descending_ order, regardless of the &#x60;sort_order&#x60; parameter. | [default to &quot;created&quot;]
 **sortOrder** | [**GetListingsByShopSortOrderParameter**](GetListingsByShopSortOrderParameter.md) | The ascending(up) or descending(down) order to sort listings by. NOTE: sort_order only works when combined with one of the search options (keywords, region, etc.). | [default to &quot;desc&quot;]
 **includes** | [**[]GetListingsByShopIncludesParameterInner**](GetListingsByShopIncludesParameterInner.md) | An enumerated string that attaches a valid association. Acceptable inputs are &#39;Shipping&#39;, &#39;Shop&#39;, &#39;Images&#39;, &#39;User&#39;, &#39;Translations&#39;, &#39;Videos&#39;, &#39;Inventory&#39; and &#39;Personalization&#39;. | 
 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 

### Return type

[**ShopListingsWithAssociations**](ShopListingsWithAssociations.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingsByShopReceipt

> ShopListings GetListingsByShopReceipt(ctx, receiptId, shopId).Limit(limit).Offset(offset).Legacy(legacy).Execute()





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
	receiptId := int64(789) // int64 | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction.
	shopId := int64(789) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	limit := int64(56) // int64 | The maximum number of results to return. (optional) (default to 25)
	offset := int64(56) // int64 | The number of records to skip before selecting the first result. (optional) (default to 0)
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.GetListingsByShopReceipt(context.Background(), receiptId, shopId).Limit(limit).Offset(offset).Legacy(legacy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.GetListingsByShopReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingsByShopReceipt`: ShopListings
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.GetListingsByShopReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **int64** | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction. | 
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsByShopReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]
 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 

### Return type

[**ShopListings**](ShopListings.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingsByShopReturnPolicy

> ShopListings GetListingsByShopReturnPolicy(ctx, returnPolicyId, shopId).Legacy(legacy).Execute()





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
	returnPolicyId := int64(789) // int64 | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies).
	shopId := int64(789) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.GetListingsByShopReturnPolicy(context.Background(), returnPolicyId, shopId).Legacy(legacy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.GetListingsByShopReturnPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingsByShopReturnPolicy`: ShopListings
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.GetListingsByShopReturnPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnPolicyId** | **int64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). | 
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsByShopReturnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 

### Return type

[**ShopListings**](ShopListings.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingsByShopSectionId

> ShopListings GetListingsByShopSectionId(ctx, shopId).ShopSectionIds(shopSectionIds).Limit(limit).Offset(offset).SortOn(sortOn).SortOrder(sortOrder).Legacy(legacy).Execute()





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
	shopSectionIds := []int64{int64(123)} // []int64 | A list of numeric IDS for all sections in a specific Etsy shop.
	limit := int64(56) // int64 | The maximum number of results to return. (optional) (default to 25)
	offset := int64(56) // int64 | The number of records to skip before selecting the first result. (optional) (default to 0)
	sortOn := openapiclient.getListingsByShop_sort_on_parameter("created") // GetListingsByShopSortOnParameter | The value to sort a search result of listings on. NOTES: a) `sort_on` only works when combined with one of the search options (keywords, region, etc.). b) when using `score` the returned results will always be in _descending_ order, regardless of the `sort_order` parameter. (optional) (default to "created")
	sortOrder := openapiclient.getListingsByShop_sort_order_parameter("asc") // GetListingsByShopSortOrderParameter | The ascending(up) or descending(down) order to sort listings by. NOTE: sort_order only works when combined with one of the search options (keywords, region, etc.). (optional) (default to "desc")
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.GetListingsByShopSectionId(context.Background(), shopId).ShopSectionIds(shopSectionIds).Limit(limit).Offset(offset).SortOn(sortOn).SortOrder(sortOrder).Legacy(legacy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.GetListingsByShopSectionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingsByShopSectionId`: ShopListings
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.GetListingsByShopSectionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsByShopSectionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopSectionIds** | **[]int64** | A list of numeric IDS for all sections in a specific Etsy shop. | 
 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]
 **sortOn** | [**GetListingsByShopSortOnParameter**](GetListingsByShopSortOnParameter.md) | The value to sort a search result of listings on. NOTES: a) &#x60;sort_on&#x60; only works when combined with one of the search options (keywords, region, etc.). b) when using &#x60;score&#x60; the returned results will always be in _descending_ order, regardless of the &#x60;sort_order&#x60; parameter. | [default to &quot;created&quot;]
 **sortOrder** | [**GetListingsByShopSortOrderParameter**](GetListingsByShopSortOrderParameter.md) | The ascending(up) or descending(down) order to sort listings by. NOTE: sort_order only works when combined with one of the search options (keywords, region, etc.). | [default to &quot;desc&quot;]
 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 

### Return type

[**ShopListings**](ShopListings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListing

> ShopListing UpdateListing(ctx, shopId, listingId).Legacy(legacy).UpdateListingRequest(updateListingRequest).Execute()





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
	legacy := true // bool | This parameter needed to enable new parameters and response values related to processing profiles. (optional)
	updateListingRequest := *openapiclient.NewUpdateListingRequest() // UpdateListingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.UpdateListing(context.Background(), shopId, listingId).Legacy(legacy).UpdateListingRequest(updateListingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.UpdateListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateListing`: ShopListing
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.UpdateListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **legacy** | **bool** | This parameter needed to enable new parameters and response values related to processing profiles. | 
 **updateListingRequest** | [**UpdateListingRequest**](UpdateListingRequest.md) |  | 

### Return type

[**ShopListing**](ShopListing.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListingProperty

> ListingPropertyValue UpdateListingProperty(ctx, shopId, listingId, propertyId).UpdateListingPropertyRequest(updateListingPropertyRequest).Execute()





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
	propertyId := int64(789) // int64 | The unique ID of an Etsy [listing property](/documentation/reference#operation/getListingProperties).
	updateListingPropertyRequest := *openapiclient.NewUpdateListingPropertyRequest([]int64{int64(123)}, []string{"Values_example"}) // UpdateListingPropertyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingAPI.UpdateListingProperty(context.Background(), shopId, listingId, propertyId).UpdateListingPropertyRequest(updateListingPropertyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingAPI.UpdateListingProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateListingProperty`: ListingPropertyValue
	fmt.Fprintf(os.Stdout, "Response from `ShopListingAPI.UpdateListingProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 
**propertyId** | **int64** | The unique ID of an Etsy [listing property](/documentation/reference#operation/getListingProperties). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListingPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateListingPropertyRequest** | [**UpdateListingPropertyRequest**](UpdateListingPropertyRequest.md) |  | 

### Return type

[**ListingPropertyValue**](ListingPropertyValue.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

