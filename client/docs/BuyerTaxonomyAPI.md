# \BuyerTaxonomyAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBuyerTaxonomyNodes**](BuyerTaxonomyAPI.md#GetBuyerTaxonomyNodes) | **Get** /v3/application/buyer-taxonomy/nodes | 
[**GetPropertiesByBuyerTaxonomyId**](BuyerTaxonomyAPI.md#GetPropertiesByBuyerTaxonomyId) | **Get** /v3/application/buyer-taxonomy/nodes/{taxonomy_id}/properties | 



## GetBuyerTaxonomyNodes

> BuyerTaxonomyNodes GetBuyerTaxonomyNodes(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerTaxonomyAPI.GetBuyerTaxonomyNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerTaxonomyAPI.GetBuyerTaxonomyNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuyerTaxonomyNodes`: BuyerTaxonomyNodes
	fmt.Fprintf(os.Stdout, "Response from `BuyerTaxonomyAPI.GetBuyerTaxonomyNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerTaxonomyNodesRequest struct via the builder pattern


### Return type

[**BuyerTaxonomyNodes**](BuyerTaxonomyNodes.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPropertiesByBuyerTaxonomyId

> BuyerTaxonomyNodeProperties GetPropertiesByBuyerTaxonomyId(ctx, taxonomyId).Execute()





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
	taxonomyId := int64(56) // int64 | The unique numeric ID of an Etsy taxonomy node, which is a metadata category for listings organized into the seller taxonomy hierarchy tree. For example, the \"shoes\" taxonomy node (ID: 1429, level: 1) is higher in the hierarchy than \"girls' shoes\" (ID: 1440, level: 2). The taxonomy nodes assigned to a listing support access to specific standardized product scales and properties. For example, listings assigned the taxonomy nodes \"shoes\" or \"girls' shoes\" support access to the \"EU\" shoe size scale with its associated property names and IDs for EU shoe sizes, such as property `value_id`:\"1394\", and `name`:\"38\".

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerTaxonomyAPI.GetPropertiesByBuyerTaxonomyId(context.Background(), taxonomyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerTaxonomyAPI.GetPropertiesByBuyerTaxonomyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPropertiesByBuyerTaxonomyId`: BuyerTaxonomyNodeProperties
	fmt.Fprintf(os.Stdout, "Response from `BuyerTaxonomyAPI.GetPropertiesByBuyerTaxonomyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonomyId** | **int64** | The unique numeric ID of an Etsy taxonomy node, which is a metadata category for listings organized into the seller taxonomy hierarchy tree. For example, the \&quot;shoes\&quot; taxonomy node (ID: 1429, level: 1) is higher in the hierarchy than \&quot;girls&#39; shoes\&quot; (ID: 1440, level: 2). The taxonomy nodes assigned to a listing support access to specific standardized product scales and properties. For example, listings assigned the taxonomy nodes \&quot;shoes\&quot; or \&quot;girls&#39; shoes\&quot; support access to the \&quot;EU\&quot; shoe size scale with its associated property names and IDs for EU shoe sizes, such as property &#x60;value_id&#x60;:\&quot;1394\&quot;, and &#x60;name&#x60;:\&quot;38\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesByBuyerTaxonomyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuyerTaxonomyNodeProperties**](BuyerTaxonomyNodeProperties.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

