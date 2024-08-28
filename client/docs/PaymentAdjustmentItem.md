# PaymentAdjustmentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentAdjustmentId** | Pointer to **int64** | The numeric ID for a payment adjustment. | [optional] 
**PaymentAdjustmentItemId** | Pointer to **int64** | Unique ID for the adjustment line item. | [optional] 
**AdjustmentType** | Pointer to **NullableString** | String indicating the type of adjustment for this line item. | [optional] 
**Amount** | Pointer to **int64** | Integer value for the amount of the adjustment in original currency. | [optional] [default to 0]
**ShopAmount** | Pointer to **int64** | Integer value for the amount of the adjustment in currency for the shop. | [optional] [default to 0]
**TransactionId** | Pointer to **NullableInt64** | The unique numeric ID for a transaction. | [optional] 
**BillPaymentId** | Pointer to **NullableInt64** | Unique ID for the bill payment adjustment. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The transaction\\&#39;s creation date and time, in epoch seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The update date and time the payment adjustment in epoch seconds. | [optional] 

## Methods

### NewPaymentAdjustmentItem

`func NewPaymentAdjustmentItem() *PaymentAdjustmentItem`

NewPaymentAdjustmentItem instantiates a new PaymentAdjustmentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAdjustmentItemWithDefaults

`func NewPaymentAdjustmentItemWithDefaults() *PaymentAdjustmentItem`

NewPaymentAdjustmentItemWithDefaults instantiates a new PaymentAdjustmentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentAdjustmentId

`func (o *PaymentAdjustmentItem) GetPaymentAdjustmentId() int64`

GetPaymentAdjustmentId returns the PaymentAdjustmentId field if non-nil, zero value otherwise.

### GetPaymentAdjustmentIdOk

`func (o *PaymentAdjustmentItem) GetPaymentAdjustmentIdOk() (*int64, bool)`

GetPaymentAdjustmentIdOk returns a tuple with the PaymentAdjustmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAdjustmentId

`func (o *PaymentAdjustmentItem) SetPaymentAdjustmentId(v int64)`

SetPaymentAdjustmentId sets PaymentAdjustmentId field to given value.

### HasPaymentAdjustmentId

`func (o *PaymentAdjustmentItem) HasPaymentAdjustmentId() bool`

HasPaymentAdjustmentId returns a boolean if a field has been set.

### GetPaymentAdjustmentItemId

`func (o *PaymentAdjustmentItem) GetPaymentAdjustmentItemId() int64`

GetPaymentAdjustmentItemId returns the PaymentAdjustmentItemId field if non-nil, zero value otherwise.

### GetPaymentAdjustmentItemIdOk

`func (o *PaymentAdjustmentItem) GetPaymentAdjustmentItemIdOk() (*int64, bool)`

GetPaymentAdjustmentItemIdOk returns a tuple with the PaymentAdjustmentItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAdjustmentItemId

`func (o *PaymentAdjustmentItem) SetPaymentAdjustmentItemId(v int64)`

SetPaymentAdjustmentItemId sets PaymentAdjustmentItemId field to given value.

### HasPaymentAdjustmentItemId

`func (o *PaymentAdjustmentItem) HasPaymentAdjustmentItemId() bool`

HasPaymentAdjustmentItemId returns a boolean if a field has been set.

### GetAdjustmentType

`func (o *PaymentAdjustmentItem) GetAdjustmentType() string`

GetAdjustmentType returns the AdjustmentType field if non-nil, zero value otherwise.

### GetAdjustmentTypeOk

`func (o *PaymentAdjustmentItem) GetAdjustmentTypeOk() (*string, bool)`

GetAdjustmentTypeOk returns a tuple with the AdjustmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentType

`func (o *PaymentAdjustmentItem) SetAdjustmentType(v string)`

SetAdjustmentType sets AdjustmentType field to given value.

### HasAdjustmentType

`func (o *PaymentAdjustmentItem) HasAdjustmentType() bool`

HasAdjustmentType returns a boolean if a field has been set.

### SetAdjustmentTypeNil

`func (o *PaymentAdjustmentItem) SetAdjustmentTypeNil(b bool)`

 SetAdjustmentTypeNil sets the value for AdjustmentType to be an explicit nil

### UnsetAdjustmentType
`func (o *PaymentAdjustmentItem) UnsetAdjustmentType()`

UnsetAdjustmentType ensures that no value is present for AdjustmentType, not even an explicit nil
### GetAmount

`func (o *PaymentAdjustmentItem) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentAdjustmentItem) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentAdjustmentItem) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentAdjustmentItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetShopAmount

`func (o *PaymentAdjustmentItem) GetShopAmount() int64`

GetShopAmount returns the ShopAmount field if non-nil, zero value otherwise.

### GetShopAmountOk

`func (o *PaymentAdjustmentItem) GetShopAmountOk() (*int64, bool)`

GetShopAmountOk returns a tuple with the ShopAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopAmount

`func (o *PaymentAdjustmentItem) SetShopAmount(v int64)`

SetShopAmount sets ShopAmount field to given value.

### HasShopAmount

`func (o *PaymentAdjustmentItem) HasShopAmount() bool`

HasShopAmount returns a boolean if a field has been set.

### GetTransactionId

`func (o *PaymentAdjustmentItem) GetTransactionId() int64`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PaymentAdjustmentItem) GetTransactionIdOk() (*int64, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PaymentAdjustmentItem) SetTransactionId(v int64)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PaymentAdjustmentItem) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *PaymentAdjustmentItem) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *PaymentAdjustmentItem) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetBillPaymentId

`func (o *PaymentAdjustmentItem) GetBillPaymentId() int64`

GetBillPaymentId returns the BillPaymentId field if non-nil, zero value otherwise.

### GetBillPaymentIdOk

`func (o *PaymentAdjustmentItem) GetBillPaymentIdOk() (*int64, bool)`

GetBillPaymentIdOk returns a tuple with the BillPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPaymentId

`func (o *PaymentAdjustmentItem) SetBillPaymentId(v int64)`

SetBillPaymentId sets BillPaymentId field to given value.

### HasBillPaymentId

`func (o *PaymentAdjustmentItem) HasBillPaymentId() bool`

HasBillPaymentId returns a boolean if a field has been set.

### SetBillPaymentIdNil

`func (o *PaymentAdjustmentItem) SetBillPaymentIdNil(b bool)`

 SetBillPaymentIdNil sets the value for BillPaymentId to be an explicit nil

### UnsetBillPaymentId
`func (o *PaymentAdjustmentItem) UnsetBillPaymentId()`

UnsetBillPaymentId ensures that no value is present for BillPaymentId, not even an explicit nil
### GetCreatedTimestamp

`func (o *PaymentAdjustmentItem) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentAdjustmentItem) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentAdjustmentItem) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *PaymentAdjustmentItem) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *PaymentAdjustmentItem) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentAdjustmentItem) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentAdjustmentItem) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *PaymentAdjustmentItem) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


