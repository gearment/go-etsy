# \ShopListingVideoAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteListingVideo**](ShopListingVideoAPI.md#DeleteListingVideo) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/videos/{video_id} | 
[**GetListingVideo**](ShopListingVideoAPI.md#GetListingVideo) | **Get** /v3/application/listings/{listing_id}/videos/{video_id} | 
[**GetListingVideos**](ShopListingVideoAPI.md#GetListingVideos) | **Get** /v3/application/listings/{listing_id}/videos | 
[**UploadListingVideo**](ShopListingVideoAPI.md#UploadListingVideo) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/videos | 



## DeleteListingVideo

> DeleteListingVideo(ctx, shopId, listingId, videoId).Execute()





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
	videoId := int64(56) // int64 | The unique ID of a video associated with a listing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShopListingVideoAPI.DeleteListingVideo(context.Background(), shopId, listingId, videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingVideoAPI.DeleteListingVideo``: %v\n", err)
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
**videoId** | **int64** | The unique ID of a video associated with a listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingVideoRequest struct via the builder pattern


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


## GetListingVideo

> ListingVideo GetListingVideo(ctx, videoId, listingId).Execute()





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
	videoId := int64(56) // int64 | The unique ID of a video associated with a listing.
	listingId := int64(56) // int64 | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingVideoAPI.GetListingVideo(context.Background(), videoId, listingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingVideoAPI.GetListingVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingVideo`: ListingVideo
	fmt.Fprintf(os.Stdout, "Response from `ShopListingVideoAPI.GetListingVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **int64** | The unique ID of a video associated with a listing. | 
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListingVideo**](ListingVideo.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingVideos

> ListingVideos GetListingVideos(ctx, listingId).Execute()





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
	listingId := int64(56) // int64 | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingVideoAPI.GetListingVideos(context.Background(), listingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingVideoAPI.GetListingVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListingVideos`: ListingVideos
	fmt.Fprintf(os.Stdout, "Response from `ShopListingVideoAPI.GetListingVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListingVideos**](ListingVideos.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadListingVideo

> ListingVideo UploadListingVideo(ctx, shopId, listingId).VideoId(videoId).Video(video).Name(name).Execute()





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
	videoId := int64(56) // int64 | The unique ID of a video associated with a listing. (optional)
	video := os.NewFile(1234, "some_file") // *os.File | A video file to upload. (optional)
	name := "name_example" // string | The file name string for the video to upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopListingVideoAPI.UploadListingVideo(context.Background(), shopId, listingId).VideoId(videoId).Video(video).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopListingVideoAPI.UploadListingVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadListingVideo`: ListingVideo
	fmt.Fprintf(os.Stdout, "Response from `ShopListingVideoAPI.UploadListingVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**listingId** | **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadListingVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **videoId** | **int64** | The unique ID of a video associated with a listing. | 
 **video** | ***os.File** | A video file to upload. | 
 **name** | **string** | The file name string for the video to upload. | 

### Return type

[**ListingVideo**](ListingVideo.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

