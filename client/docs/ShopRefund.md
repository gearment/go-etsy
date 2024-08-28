# ShopRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Money**](Money.md) | A number equal to the refund total. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The date &amp; time of the refund, in epoch seconds. | [optional] 
**Reason** | Pointer to **NullableString** | The reason string given for the refund. | [optional] 
**NoteFromIssuer** | Pointer to **NullableString** | The note string created by the refund issuer. | [optional] 
**Status** | Pointer to **NullableString** | The status indication string for the refund. | [optional] 

## Methods

### NewShopRefund

`func NewShopRefund() *ShopRefund`

NewShopRefund instantiates a new ShopRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopRefundWithDefaults

`func NewShopRefundWithDefaults() *ShopRefund`

NewShopRefundWithDefaults instantiates a new ShopRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ShopRefund) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ShopRefund) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ShopRefund) SetAmount(v Money)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ShopRefund) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ShopRefund) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ShopRefund) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ShopRefund) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ShopRefund) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetReason

`func (o *ShopRefund) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ShopRefund) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ShopRefund) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ShopRefund) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ShopRefund) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ShopRefund) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetNoteFromIssuer

`func (o *ShopRefund) GetNoteFromIssuer() string`

GetNoteFromIssuer returns the NoteFromIssuer field if non-nil, zero value otherwise.

### GetNoteFromIssuerOk

`func (o *ShopRefund) GetNoteFromIssuerOk() (*string, bool)`

GetNoteFromIssuerOk returns a tuple with the NoteFromIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteFromIssuer

`func (o *ShopRefund) SetNoteFromIssuer(v string)`

SetNoteFromIssuer sets NoteFromIssuer field to given value.

### HasNoteFromIssuer

`func (o *ShopRefund) HasNoteFromIssuer() bool`

HasNoteFromIssuer returns a boolean if a field has been set.

### SetNoteFromIssuerNil

`func (o *ShopRefund) SetNoteFromIssuerNil(b bool)`

 SetNoteFromIssuerNil sets the value for NoteFromIssuer to be an explicit nil

### UnsetNoteFromIssuer
`func (o *ShopRefund) UnsetNoteFromIssuer()`

UnsetNoteFromIssuer ensures that no value is present for NoteFromIssuer, not even an explicit nil
### GetStatus

`func (o *ShopRefund) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShopRefund) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShopRefund) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShopRefund) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ShopRefund) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ShopRefund) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


