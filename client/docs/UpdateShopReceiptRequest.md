# UpdateShopReceiptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WasShipped** | Pointer to **NullableBool** | When &#x60;true&#x60;, returns receipts where the seller shipped the product(s) in this receipt. When &#x60;false&#x60;, returns receipts where shipment has not been set. | [optional] 
**WasPaid** | Pointer to **NullableBool** | When &#x60;true&#x60;, returns receipts where the seller has recieved payment for the receipt. When &#x60;false&#x60;, returns receipts where payment has not been received. | [optional] 

## Methods

### NewUpdateShopReceiptRequest

`func NewUpdateShopReceiptRequest() *UpdateShopReceiptRequest`

NewUpdateShopReceiptRequest instantiates a new UpdateShopReceiptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShopReceiptRequestWithDefaults

`func NewUpdateShopReceiptRequestWithDefaults() *UpdateShopReceiptRequest`

NewUpdateShopReceiptRequestWithDefaults instantiates a new UpdateShopReceiptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWasShipped

`func (o *UpdateShopReceiptRequest) GetWasShipped() bool`

GetWasShipped returns the WasShipped field if non-nil, zero value otherwise.

### GetWasShippedOk

`func (o *UpdateShopReceiptRequest) GetWasShippedOk() (*bool, bool)`

GetWasShippedOk returns a tuple with the WasShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasShipped

`func (o *UpdateShopReceiptRequest) SetWasShipped(v bool)`

SetWasShipped sets WasShipped field to given value.

### HasWasShipped

`func (o *UpdateShopReceiptRequest) HasWasShipped() bool`

HasWasShipped returns a boolean if a field has been set.

### SetWasShippedNil

`func (o *UpdateShopReceiptRequest) SetWasShippedNil(b bool)`

 SetWasShippedNil sets the value for WasShipped to be an explicit nil

### UnsetWasShipped
`func (o *UpdateShopReceiptRequest) UnsetWasShipped()`

UnsetWasShipped ensures that no value is present for WasShipped, not even an explicit nil
### GetWasPaid

`func (o *UpdateShopReceiptRequest) GetWasPaid() bool`

GetWasPaid returns the WasPaid field if non-nil, zero value otherwise.

### GetWasPaidOk

`func (o *UpdateShopReceiptRequest) GetWasPaidOk() (*bool, bool)`

GetWasPaidOk returns a tuple with the WasPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasPaid

`func (o *UpdateShopReceiptRequest) SetWasPaid(v bool)`

SetWasPaid sets WasPaid field to given value.

### HasWasPaid

`func (o *UpdateShopReceiptRequest) HasWasPaid() bool`

HasWasPaid returns a boolean if a field has been set.

### SetWasPaidNil

`func (o *UpdateShopReceiptRequest) SetWasPaidNil(b bool)`

 SetWasPaidNil sets the value for WasPaid to be an explicit nil

### UnsetWasPaid
`func (o *UpdateShopReceiptRequest) UnsetWasPaid()`

UnsetWasPaid ensures that no value is present for WasPaid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


