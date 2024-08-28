# \ShopHolidayPreferencesAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHolidayPreferences**](ShopHolidayPreferencesAPI.md#GetHolidayPreferences) | **Get** /v3/application/shops/{shop_id}/holiday-preferences | 
[**UpdateHolidayPreferences**](ShopHolidayPreferencesAPI.md#UpdateHolidayPreferences) | **Put** /v3/application/shops/{shop_id}/holiday-preferences/{holiday_id} | 



## GetHolidayPreferences

> []ShopHolidayPreference GetHolidayPreferences(ctx, shopId).Execute()





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
	resp, r, err := apiClient.ShopHolidayPreferencesAPI.GetHolidayPreferences(context.Background(), shopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopHolidayPreferencesAPI.GetHolidayPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHolidayPreferences`: []ShopHolidayPreference
	fmt.Fprintf(os.Stdout, "Response from `ShopHolidayPreferencesAPI.GetHolidayPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHolidayPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ShopHolidayPreference**](ShopHolidayPreference.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHolidayPreferences

> ShopHolidayPreference UpdateHolidayPreferences(ctx, shopId, holidayId).IsWorking(isWorking).Execute()





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
	holidayId := openapiclient.updateHolidayPreferences_holiday_id_parameter("1") // UpdateHolidayPreferencesHolidayIdParameter | The unique id that maps to the holiday a country observes. See the [Fulfillment Tutorial docs](https://developer.etsy.com/documentation/tutorials/fulfillment/#country-holidays) for more info
	isWorking := true // bool | A boolean value for whether the shop will process orders on a particular holiday.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopHolidayPreferencesAPI.UpdateHolidayPreferences(context.Background(), shopId, holidayId).IsWorking(isWorking).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopHolidayPreferencesAPI.UpdateHolidayPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHolidayPreferences`: ShopHolidayPreference
	fmt.Fprintf(os.Stdout, "Response from `ShopHolidayPreferencesAPI.UpdateHolidayPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**holidayId** | [**UpdateHolidayPreferencesHolidayIdParameter**](.md) | The unique id that maps to the holiday a country observes. See the [Fulfillment Tutorial docs](https://developer.etsy.com/documentation/tutorials/fulfillment/#country-holidays) for more info | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHolidayPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isWorking** | **bool** | A boolean value for whether the shop will process orders on a particular holiday. | 

### Return type

[**ShopHolidayPreference**](ShopHolidayPreference.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

