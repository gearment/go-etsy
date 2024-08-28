# PaymentAdjustment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentAdjustmentId** | Pointer to **int64** | The numeric ID for a payment adjustment. | [optional] 
**PaymentId** | Pointer to **int64** | A unique numeric ID for a payment to a specific Etsy [shop](/documentation/reference#tag/Shop). | [optional] 
**Status** | Pointer to **string** | The status string of the payment adjustment. | [optional] 
**IsSuccess** | Pointer to **bool** | When true, the payment adjustment was or is likely to complete successfully. | [optional] 
**UserId** | Pointer to **int64** | The numeric ID for the [user](/documentation/reference#tag/User) (seller) fulfilling the purchase. | [optional] 
**ReasonCode** | Pointer to **string** | A human-readable string describing the reason for the refund. | [optional] 
**TotalAdjustmentAmount** | Pointer to **NullableInt64** | The total numeric amount of the refund in the payment currency. | [optional] 
**ShopTotalAdjustmentAmount** | Pointer to **NullableInt64** | The numeric amount of the refund in the shop currency. | [optional] 
**BuyerTotalAdjustmentAmount** | Pointer to **NullableInt64** | The numeric amount of the refund in the buyer currency. | [optional] 
**TotalFeeAdjustmentAmount** | Pointer to **NullableInt64** | The numeric amount of card processing fees associated with a payment adjustment. | [optional] 
**CreateTimestamp** | Pointer to **int64** | The transaction\\&#39;s creation date and time, in epoch seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The transaction\\&#39;s creation date and time, in epoch seconds. | [optional] 
**UpdateTimestamp** | Pointer to **int64** | The date and time of the last change to the payment adjustment in epoch seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The date and time of the last change to the payment adjustment in epoch seconds. | [optional] 
**PaymentAdjustmentItems** | Pointer to [**[]PaymentAdjustmentItem**](PaymentAdjustmentItem.md) | List of payment adjustment line items. | [optional] 

## Methods

### NewPaymentAdjustment

`func NewPaymentAdjustment() *PaymentAdjustment`

NewPaymentAdjustment instantiates a new PaymentAdjustment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAdjustmentWithDefaults

`func NewPaymentAdjustmentWithDefaults() *PaymentAdjustment`

NewPaymentAdjustmentWithDefaults instantiates a new PaymentAdjustment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentAdjustmentId

`func (o *PaymentAdjustment) GetPaymentAdjustmentId() int64`

GetPaymentAdjustmentId returns the PaymentAdjustmentId field if non-nil, zero value otherwise.

### GetPaymentAdjustmentIdOk

`func (o *PaymentAdjustment) GetPaymentAdjustmentIdOk() (*int64, bool)`

GetPaymentAdjustmentIdOk returns a tuple with the PaymentAdjustmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAdjustmentId

`func (o *PaymentAdjustment) SetPaymentAdjustmentId(v int64)`

SetPaymentAdjustmentId sets PaymentAdjustmentId field to given value.

### HasPaymentAdjustmentId

`func (o *PaymentAdjustment) HasPaymentAdjustmentId() bool`

HasPaymentAdjustmentId returns a boolean if a field has been set.

### GetPaymentId

`func (o *PaymentAdjustment) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentAdjustment) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentAdjustment) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PaymentAdjustment) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentAdjustment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentAdjustment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentAdjustment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentAdjustment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsSuccess

`func (o *PaymentAdjustment) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *PaymentAdjustment) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *PaymentAdjustment) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *PaymentAdjustment) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetUserId

`func (o *PaymentAdjustment) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PaymentAdjustment) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PaymentAdjustment) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PaymentAdjustment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetReasonCode

`func (o *PaymentAdjustment) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *PaymentAdjustment) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *PaymentAdjustment) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *PaymentAdjustment) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetTotalAdjustmentAmount

`func (o *PaymentAdjustment) GetTotalAdjustmentAmount() int64`

GetTotalAdjustmentAmount returns the TotalAdjustmentAmount field if non-nil, zero value otherwise.

### GetTotalAdjustmentAmountOk

`func (o *PaymentAdjustment) GetTotalAdjustmentAmountOk() (*int64, bool)`

GetTotalAdjustmentAmountOk returns a tuple with the TotalAdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAdjustmentAmount

`func (o *PaymentAdjustment) SetTotalAdjustmentAmount(v int64)`

SetTotalAdjustmentAmount sets TotalAdjustmentAmount field to given value.

### HasTotalAdjustmentAmount

`func (o *PaymentAdjustment) HasTotalAdjustmentAmount() bool`

HasTotalAdjustmentAmount returns a boolean if a field has been set.

### SetTotalAdjustmentAmountNil

`func (o *PaymentAdjustment) SetTotalAdjustmentAmountNil(b bool)`

 SetTotalAdjustmentAmountNil sets the value for TotalAdjustmentAmount to be an explicit nil

### UnsetTotalAdjustmentAmount
`func (o *PaymentAdjustment) UnsetTotalAdjustmentAmount()`

UnsetTotalAdjustmentAmount ensures that no value is present for TotalAdjustmentAmount, not even an explicit nil
### GetShopTotalAdjustmentAmount

`func (o *PaymentAdjustment) GetShopTotalAdjustmentAmount() int64`

GetShopTotalAdjustmentAmount returns the ShopTotalAdjustmentAmount field if non-nil, zero value otherwise.

### GetShopTotalAdjustmentAmountOk

`func (o *PaymentAdjustment) GetShopTotalAdjustmentAmountOk() (*int64, bool)`

GetShopTotalAdjustmentAmountOk returns a tuple with the ShopTotalAdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopTotalAdjustmentAmount

`func (o *PaymentAdjustment) SetShopTotalAdjustmentAmount(v int64)`

SetShopTotalAdjustmentAmount sets ShopTotalAdjustmentAmount field to given value.

### HasShopTotalAdjustmentAmount

`func (o *PaymentAdjustment) HasShopTotalAdjustmentAmount() bool`

HasShopTotalAdjustmentAmount returns a boolean if a field has been set.

### SetShopTotalAdjustmentAmountNil

`func (o *PaymentAdjustment) SetShopTotalAdjustmentAmountNil(b bool)`

 SetShopTotalAdjustmentAmountNil sets the value for ShopTotalAdjustmentAmount to be an explicit nil

### UnsetShopTotalAdjustmentAmount
`func (o *PaymentAdjustment) UnsetShopTotalAdjustmentAmount()`

UnsetShopTotalAdjustmentAmount ensures that no value is present for ShopTotalAdjustmentAmount, not even an explicit nil
### GetBuyerTotalAdjustmentAmount

`func (o *PaymentAdjustment) GetBuyerTotalAdjustmentAmount() int64`

GetBuyerTotalAdjustmentAmount returns the BuyerTotalAdjustmentAmount field if non-nil, zero value otherwise.

### GetBuyerTotalAdjustmentAmountOk

`func (o *PaymentAdjustment) GetBuyerTotalAdjustmentAmountOk() (*int64, bool)`

GetBuyerTotalAdjustmentAmountOk returns a tuple with the BuyerTotalAdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerTotalAdjustmentAmount

`func (o *PaymentAdjustment) SetBuyerTotalAdjustmentAmount(v int64)`

SetBuyerTotalAdjustmentAmount sets BuyerTotalAdjustmentAmount field to given value.

### HasBuyerTotalAdjustmentAmount

`func (o *PaymentAdjustment) HasBuyerTotalAdjustmentAmount() bool`

HasBuyerTotalAdjustmentAmount returns a boolean if a field has been set.

### SetBuyerTotalAdjustmentAmountNil

`func (o *PaymentAdjustment) SetBuyerTotalAdjustmentAmountNil(b bool)`

 SetBuyerTotalAdjustmentAmountNil sets the value for BuyerTotalAdjustmentAmount to be an explicit nil

### UnsetBuyerTotalAdjustmentAmount
`func (o *PaymentAdjustment) UnsetBuyerTotalAdjustmentAmount()`

UnsetBuyerTotalAdjustmentAmount ensures that no value is present for BuyerTotalAdjustmentAmount, not even an explicit nil
### GetTotalFeeAdjustmentAmount

`func (o *PaymentAdjustment) GetTotalFeeAdjustmentAmount() int64`

GetTotalFeeAdjustmentAmount returns the TotalFeeAdjustmentAmount field if non-nil, zero value otherwise.

### GetTotalFeeAdjustmentAmountOk

`func (o *PaymentAdjustment) GetTotalFeeAdjustmentAmountOk() (*int64, bool)`

GetTotalFeeAdjustmentAmountOk returns a tuple with the TotalFeeAdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFeeAdjustmentAmount

`func (o *PaymentAdjustment) SetTotalFeeAdjustmentAmount(v int64)`

SetTotalFeeAdjustmentAmount sets TotalFeeAdjustmentAmount field to given value.

### HasTotalFeeAdjustmentAmount

`func (o *PaymentAdjustment) HasTotalFeeAdjustmentAmount() bool`

HasTotalFeeAdjustmentAmount returns a boolean if a field has been set.

### SetTotalFeeAdjustmentAmountNil

`func (o *PaymentAdjustment) SetTotalFeeAdjustmentAmountNil(b bool)`

 SetTotalFeeAdjustmentAmountNil sets the value for TotalFeeAdjustmentAmount to be an explicit nil

### UnsetTotalFeeAdjustmentAmount
`func (o *PaymentAdjustment) UnsetTotalFeeAdjustmentAmount()`

UnsetTotalFeeAdjustmentAmount ensures that no value is present for TotalFeeAdjustmentAmount, not even an explicit nil
### GetCreateTimestamp

`func (o *PaymentAdjustment) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *PaymentAdjustment) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *PaymentAdjustment) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *PaymentAdjustment) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *PaymentAdjustment) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentAdjustment) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentAdjustment) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *PaymentAdjustment) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *PaymentAdjustment) GetUpdateTimestamp() int64`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *PaymentAdjustment) GetUpdateTimestampOk() (*int64, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *PaymentAdjustment) SetUpdateTimestamp(v int64)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *PaymentAdjustment) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *PaymentAdjustment) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentAdjustment) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentAdjustment) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *PaymentAdjustment) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetPaymentAdjustmentItems

`func (o *PaymentAdjustment) GetPaymentAdjustmentItems() []PaymentAdjustmentItem`

GetPaymentAdjustmentItems returns the PaymentAdjustmentItems field if non-nil, zero value otherwise.

### GetPaymentAdjustmentItemsOk

`func (o *PaymentAdjustment) GetPaymentAdjustmentItemsOk() (*[]PaymentAdjustmentItem, bool)`

GetPaymentAdjustmentItemsOk returns a tuple with the PaymentAdjustmentItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAdjustmentItems

`func (o *PaymentAdjustment) SetPaymentAdjustmentItems(v []PaymentAdjustmentItem)`

SetPaymentAdjustmentItems sets PaymentAdjustmentItems field to given value.

### HasPaymentAdjustmentItems

`func (o *PaymentAdjustment) HasPaymentAdjustmentItems() bool`

HasPaymentAdjustmentItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


