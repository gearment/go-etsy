# \PaymentAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentAccountLedgerEntryPayments**](PaymentAPI.md#GetPaymentAccountLedgerEntryPayments) | **Get** /v3/application/shops/{shop_id}/payment-account/ledger-entries/payments | 
[**GetPayments**](PaymentAPI.md#GetPayments) | **Get** /v3/application/shops/{shop_id}/payments | 
[**GetShopPaymentByReceiptId**](PaymentAPI.md#GetShopPaymentByReceiptId) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id}/payments | 



## GetPaymentAccountLedgerEntryPayments

> Payments GetPaymentAccountLedgerEntryPayments(ctx, shopId).LedgerEntryIds(ledgerEntryIds).Execute()





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
	ledgerEntryIds := []int64{int64(123)} // []int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.GetPaymentAccountLedgerEntryPayments(context.Background(), shopId).LedgerEntryIds(ledgerEntryIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.GetPaymentAccountLedgerEntryPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentAccountLedgerEntryPayments`: Payments
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.GetPaymentAccountLedgerEntryPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentAccountLedgerEntryPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ledgerEntryIds** | **[]int64** |  | 

### Return type

[**Payments**](Payments.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayments

> Payments GetPayments(ctx, shopId).PaymentIds(paymentIds).Execute()





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
	paymentIds := []int64{int64(123)} // []int64 | A comma-separated array of Payment IDs numbers.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.GetPayments(context.Background(), shopId).PaymentIds(paymentIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.GetPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayments`: Payments
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.GetPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentIds** | **[]int64** | A comma-separated array of Payment IDs numbers. | 

### Return type

[**Payments**](Payments.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopPaymentByReceiptId

> Payments GetShopPaymentByReceiptId(ctx, shopId, receiptId).Execute()





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
	resp, r, err := apiClient.PaymentAPI.GetShopPaymentByReceiptId(context.Background(), shopId, receiptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.GetShopPaymentByReceiptId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopPaymentByReceiptId`: Payments
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.GetShopPaymentByReceiptId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**receiptId** | **int64** | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopPaymentByReceiptIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Payments**](Payments.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

