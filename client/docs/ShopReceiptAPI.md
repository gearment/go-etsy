# \ShopReceiptAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReceiptShipment**](ShopReceiptAPI.md#CreateReceiptShipment) | **Post** /v3/application/shops/{shop_id}/receipts/{receipt_id}/tracking | 
[**GetShopReceipt**](ShopReceiptAPI.md#GetShopReceipt) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id} | 
[**GetShopReceipts**](ShopReceiptAPI.md#GetShopReceipts) | **Get** /v3/application/shops/{shop_id}/receipts | 
[**UpdateShopReceipt**](ShopReceiptAPI.md#UpdateShopReceipt) | **Put** /v3/application/shops/{shop_id}/receipts/{receipt_id} | 



## CreateReceiptShipment

> ShopReceipt CreateReceiptShipment(ctx, shopId, receiptId).TrackingCode(trackingCode).CarrierName(carrierName).SendBcc(sendBcc).NoteToBuyer(noteToBuyer).Execute()





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
	receiptId := int64(56) // int64 | The receipt to submit tracking for.
	trackingCode := "trackingCode_example" // string | The tracking code for this receipt. (optional)
	carrierName := "carrierName_example" // string | The carrier name for this receipt. (optional)
	sendBcc := true // bool | If true, the shipping notification will be sent to the seller as well (optional)
	noteToBuyer := "noteToBuyer_example" // string | Message to include in notification to the buyer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopReceiptAPI.CreateReceiptShipment(context.Background(), shopId, receiptId).TrackingCode(trackingCode).CarrierName(carrierName).SendBcc(sendBcc).NoteToBuyer(noteToBuyer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReceiptAPI.CreateReceiptShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReceiptShipment`: ShopReceipt
	fmt.Fprintf(os.Stdout, "Response from `ShopReceiptAPI.CreateReceiptShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**receiptId** | **int64** | The receipt to submit tracking for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **trackingCode** | **string** | The tracking code for this receipt. | 
 **carrierName** | **string** | The carrier name for this receipt. | 
 **sendBcc** | **bool** | If true, the shipping notification will be sent to the seller as well | 
 **noteToBuyer** | **string** | Message to include in notification to the buyer. | 

### Return type

[**ShopReceipt**](ShopReceipt.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopReceipt

> ShopReceipt GetShopReceipt(ctx, shopId, receiptId).Execute()





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
	receiptId := int64(56) // int64 | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopReceiptAPI.GetShopReceipt(context.Background(), shopId, receiptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReceiptAPI.GetShopReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopReceipt`: ShopReceipt
	fmt.Fprintf(os.Stdout, "Response from `ShopReceiptAPI.GetShopReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**receiptId** | **int64** | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShopReceipt**](ShopReceipt.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopReceipts

> ShopReceipts GetShopReceipts(ctx, shopId).MinCreated(minCreated).MaxCreated(maxCreated).MinLastModified(minLastModified).MaxLastModified(maxLastModified).Limit(limit).Offset(offset).SortOn(sortOn).SortOrder(sortOrder).WasPaid(wasPaid).WasShipped(wasShipped).WasDelivered(wasDelivered).WasCanceled(wasCanceled).Execute()





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
	minCreated := int64(56) // int64 | The earliest unix timestamp for when a record was created. (optional)
	maxCreated := int64(56) // int64 | The latest unix timestamp for when a record was created. (optional)
	minLastModified := int64(56) // int64 | The earliest unix timestamp for when a record last changed. (optional)
	maxLastModified := int64(56) // int64 | The latest unix timestamp for when a record last changed. (optional)
	limit := int64(56) // int64 | The maximum number of results to return. (optional) (default to 25)
	offset := int64(56) // int64 | The number of records to skip before selecting the first result. (optional) (default to 0)
	sortOn := openapiclient.getShopReceipts_sort_on_parameter("created") // GetShopReceiptsSortOnParameter | The value to sort a search result of listings on. (optional) (default to "created")
	sortOrder := openapiclient.getShopReceipts_sort_order_parameter("asc") // GetShopReceiptsSortOrderParameter | The ascending(up) or descending(down) order to sort receipts by. (optional) (default to "desc")
	wasPaid := true // bool | When `true`, returns receipts where the seller has recieved payment for the receipt. When `false`, returns receipts where payment has not been received. (optional)
	wasShipped := true // bool | When `true`, returns receipts where the seller shipped the product(s) in this receipt. When `false`, returns receipts where shipment has not been set. (optional)
	wasDelivered := true // bool | When `true`, returns receipts that have been marked as delivered. When `false`, returns receipts where shipment has not been marked as delivered. (optional)
	wasCanceled := true // bool | When `true`, the endpoint will only return the canceled receipts. When `false`, the endpoint will only return non-canceled receipts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopReceiptAPI.GetShopReceipts(context.Background(), shopId).MinCreated(minCreated).MaxCreated(maxCreated).MinLastModified(minLastModified).MaxLastModified(maxLastModified).Limit(limit).Offset(offset).SortOn(sortOn).SortOrder(sortOrder).WasPaid(wasPaid).WasShipped(wasShipped).WasDelivered(wasDelivered).WasCanceled(wasCanceled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReceiptAPI.GetShopReceipts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopReceipts`: ShopReceipts
	fmt.Fprintf(os.Stdout, "Response from `ShopReceiptAPI.GetShopReceipts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minCreated** | **int64** | The earliest unix timestamp for when a record was created. | 
 **maxCreated** | **int64** | The latest unix timestamp for when a record was created. | 
 **minLastModified** | **int64** | The earliest unix timestamp for when a record last changed. | 
 **maxLastModified** | **int64** | The latest unix timestamp for when a record last changed. | 
 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]
 **sortOn** | [**GetShopReceiptsSortOnParameter**](GetShopReceiptsSortOnParameter.md) | The value to sort a search result of listings on. | [default to &quot;created&quot;]
 **sortOrder** | [**GetShopReceiptsSortOrderParameter**](GetShopReceiptsSortOrderParameter.md) | The ascending(up) or descending(down) order to sort receipts by. | [default to &quot;desc&quot;]
 **wasPaid** | **bool** | When &#x60;true&#x60;, returns receipts where the seller has recieved payment for the receipt. When &#x60;false&#x60;, returns receipts where payment has not been received. | 
 **wasShipped** | **bool** | When &#x60;true&#x60;, returns receipts where the seller shipped the product(s) in this receipt. When &#x60;false&#x60;, returns receipts where shipment has not been set. | 
 **wasDelivered** | **bool** | When &#x60;true&#x60;, returns receipts that have been marked as delivered. When &#x60;false&#x60;, returns receipts where shipment has not been marked as delivered. | 
 **wasCanceled** | **bool** | When &#x60;true&#x60;, the endpoint will only return the canceled receipts. When &#x60;false&#x60;, the endpoint will only return non-canceled receipts. | 

### Return type

[**ShopReceipts**](ShopReceipts.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShopReceipt

> ShopReceipt UpdateShopReceipt(ctx, shopId, receiptId).WasShipped(wasShipped).WasPaid(wasPaid).Execute()





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
	receiptId := int64(56) // int64 | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction.
	wasShipped := true // bool | When `true`, returns receipts where the seller shipped the product(s) in this receipt. When `false`, returns receipts where shipment has not been set. (optional)
	wasPaid := true // bool | When `true`, returns receipts where the seller has recieved payment for the receipt. When `false`, returns receipts where payment has not been received. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopReceiptAPI.UpdateShopReceipt(context.Background(), shopId, receiptId).WasShipped(wasShipped).WasPaid(wasPaid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopReceiptAPI.UpdateShopReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShopReceipt`: ShopReceipt
	fmt.Fprintf(os.Stdout, "Response from `ShopReceiptAPI.UpdateShopReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**receiptId** | **int64** | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShopReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wasShipped** | **bool** | When &#x60;true&#x60;, returns receipts where the seller shipped the product(s) in this receipt. When &#x60;false&#x60;, returns receipts where shipment has not been set. | 
 **wasPaid** | **bool** | When &#x60;true&#x60;, returns receipts where the seller has recieved payment for the receipt. When &#x60;false&#x60;, returns receipts where payment has not been received. | 

### Return type

[**ShopReceipt**](ShopReceipt.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

