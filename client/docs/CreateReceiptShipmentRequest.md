# CreateReceiptShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingCode** | Pointer to **string** | The tracking code for this receipt. | [optional] 
**CarrierName** | Pointer to **string** | The carrier name for this receipt. | [optional] 
**SendBcc** | Pointer to **bool** | If true, the shipping notification will be sent to the seller as well | [optional] 
**NoteToBuyer** | Pointer to **string** | Message to include in notification to the buyer. | [optional] 

## Methods

### NewCreateReceiptShipmentRequest

`func NewCreateReceiptShipmentRequest() *CreateReceiptShipmentRequest`

NewCreateReceiptShipmentRequest instantiates a new CreateReceiptShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReceiptShipmentRequestWithDefaults

`func NewCreateReceiptShipmentRequestWithDefaults() *CreateReceiptShipmentRequest`

NewCreateReceiptShipmentRequestWithDefaults instantiates a new CreateReceiptShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingCode

`func (o *CreateReceiptShipmentRequest) GetTrackingCode() string`

GetTrackingCode returns the TrackingCode field if non-nil, zero value otherwise.

### GetTrackingCodeOk

`func (o *CreateReceiptShipmentRequest) GetTrackingCodeOk() (*string, bool)`

GetTrackingCodeOk returns a tuple with the TrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCode

`func (o *CreateReceiptShipmentRequest) SetTrackingCode(v string)`

SetTrackingCode sets TrackingCode field to given value.

### HasTrackingCode

`func (o *CreateReceiptShipmentRequest) HasTrackingCode() bool`

HasTrackingCode returns a boolean if a field has been set.

### GetCarrierName

`func (o *CreateReceiptShipmentRequest) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *CreateReceiptShipmentRequest) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *CreateReceiptShipmentRequest) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *CreateReceiptShipmentRequest) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetSendBcc

`func (o *CreateReceiptShipmentRequest) GetSendBcc() bool`

GetSendBcc returns the SendBcc field if non-nil, zero value otherwise.

### GetSendBccOk

`func (o *CreateReceiptShipmentRequest) GetSendBccOk() (*bool, bool)`

GetSendBccOk returns a tuple with the SendBcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBcc

`func (o *CreateReceiptShipmentRequest) SetSendBcc(v bool)`

SetSendBcc sets SendBcc field to given value.

### HasSendBcc

`func (o *CreateReceiptShipmentRequest) HasSendBcc() bool`

HasSendBcc returns a boolean if a field has been set.

### GetNoteToBuyer

`func (o *CreateReceiptShipmentRequest) GetNoteToBuyer() string`

GetNoteToBuyer returns the NoteToBuyer field if non-nil, zero value otherwise.

### GetNoteToBuyerOk

`func (o *CreateReceiptShipmentRequest) GetNoteToBuyerOk() (*string, bool)`

GetNoteToBuyerOk returns a tuple with the NoteToBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToBuyer

`func (o *CreateReceiptShipmentRequest) SetNoteToBuyer(v string)`

SetNoteToBuyer sets NoteToBuyer field to given value.

### HasNoteToBuyer

`func (o *CreateReceiptShipmentRequest) HasNoteToBuyer() bool`

HasNoteToBuyer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


