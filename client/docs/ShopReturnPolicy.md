# ShopReturnPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnPolicyId** | Pointer to **int64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). | [optional] 
**ShopId** | Pointer to **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | [optional] 
**AcceptsReturns** | Pointer to **bool** | return_policy_accepts_returns | [optional] 
**AcceptsExchanges** | Pointer to **bool** | return_policy_accepts_exchanges | [optional] 
**ReturnDeadline** | Pointer to **NullableInt64** | The deadline for the Return Policy, measured in days. The value must be one of the following: [7, 14, 21, 30, 45, 60, 90]. | [optional] 

## Methods

### NewShopReturnPolicy

`func NewShopReturnPolicy() *ShopReturnPolicy`

NewShopReturnPolicy instantiates a new ShopReturnPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopReturnPolicyWithDefaults

`func NewShopReturnPolicyWithDefaults() *ShopReturnPolicy`

NewShopReturnPolicyWithDefaults instantiates a new ShopReturnPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnPolicyId

`func (o *ShopReturnPolicy) GetReturnPolicyId() int64`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *ShopReturnPolicy) GetReturnPolicyIdOk() (*int64, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *ShopReturnPolicy) SetReturnPolicyId(v int64)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *ShopReturnPolicy) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### GetShopId

`func (o *ShopReturnPolicy) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ShopReturnPolicy) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ShopReturnPolicy) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ShopReturnPolicy) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetAcceptsReturns

`func (o *ShopReturnPolicy) GetAcceptsReturns() bool`

GetAcceptsReturns returns the AcceptsReturns field if non-nil, zero value otherwise.

### GetAcceptsReturnsOk

`func (o *ShopReturnPolicy) GetAcceptsReturnsOk() (*bool, bool)`

GetAcceptsReturnsOk returns a tuple with the AcceptsReturns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsReturns

`func (o *ShopReturnPolicy) SetAcceptsReturns(v bool)`

SetAcceptsReturns sets AcceptsReturns field to given value.

### HasAcceptsReturns

`func (o *ShopReturnPolicy) HasAcceptsReturns() bool`

HasAcceptsReturns returns a boolean if a field has been set.

### GetAcceptsExchanges

`func (o *ShopReturnPolicy) GetAcceptsExchanges() bool`

GetAcceptsExchanges returns the AcceptsExchanges field if non-nil, zero value otherwise.

### GetAcceptsExchangesOk

`func (o *ShopReturnPolicy) GetAcceptsExchangesOk() (*bool, bool)`

GetAcceptsExchangesOk returns a tuple with the AcceptsExchanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsExchanges

`func (o *ShopReturnPolicy) SetAcceptsExchanges(v bool)`

SetAcceptsExchanges sets AcceptsExchanges field to given value.

### HasAcceptsExchanges

`func (o *ShopReturnPolicy) HasAcceptsExchanges() bool`

HasAcceptsExchanges returns a boolean if a field has been set.

### GetReturnDeadline

`func (o *ShopReturnPolicy) GetReturnDeadline() int64`

GetReturnDeadline returns the ReturnDeadline field if non-nil, zero value otherwise.

### GetReturnDeadlineOk

`func (o *ShopReturnPolicy) GetReturnDeadlineOk() (*int64, bool)`

GetReturnDeadlineOk returns a tuple with the ReturnDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDeadline

`func (o *ShopReturnPolicy) SetReturnDeadline(v int64)`

SetReturnDeadline sets ReturnDeadline field to given value.

### HasReturnDeadline

`func (o *ShopReturnPolicy) HasReturnDeadline() bool`

HasReturnDeadline returns a boolean if a field has been set.

### SetReturnDeadlineNil

`func (o *ShopReturnPolicy) SetReturnDeadlineNil(b bool)`

 SetReturnDeadlineNil sets the value for ReturnDeadline to be an explicit nil

### UnsetReturnDeadline
`func (o *ShopReturnPolicy) UnsetReturnDeadline()`

UnsetReturnDeadline ensures that no value is present for ReturnDeadline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


