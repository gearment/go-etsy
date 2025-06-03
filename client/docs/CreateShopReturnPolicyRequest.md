# CreateShopReturnPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptsReturns** | **bool** |  | 
**AcceptsExchanges** | **bool** |  | 
**ReturnDeadline** | Pointer to **NullableInt64** | The deadline for the Return Policy, measured in days. The value must be one of the following: [7, 14, 21, 30, 45, 60, 90]. | [optional] 

## Methods

### NewCreateShopReturnPolicyRequest

`func NewCreateShopReturnPolicyRequest(acceptsReturns bool, acceptsExchanges bool, ) *CreateShopReturnPolicyRequest`

NewCreateShopReturnPolicyRequest instantiates a new CreateShopReturnPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShopReturnPolicyRequestWithDefaults

`func NewCreateShopReturnPolicyRequestWithDefaults() *CreateShopReturnPolicyRequest`

NewCreateShopReturnPolicyRequestWithDefaults instantiates a new CreateShopReturnPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptsReturns

`func (o *CreateShopReturnPolicyRequest) GetAcceptsReturns() bool`

GetAcceptsReturns returns the AcceptsReturns field if non-nil, zero value otherwise.

### GetAcceptsReturnsOk

`func (o *CreateShopReturnPolicyRequest) GetAcceptsReturnsOk() (*bool, bool)`

GetAcceptsReturnsOk returns a tuple with the AcceptsReturns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsReturns

`func (o *CreateShopReturnPolicyRequest) SetAcceptsReturns(v bool)`

SetAcceptsReturns sets AcceptsReturns field to given value.


### GetAcceptsExchanges

`func (o *CreateShopReturnPolicyRequest) GetAcceptsExchanges() bool`

GetAcceptsExchanges returns the AcceptsExchanges field if non-nil, zero value otherwise.

### GetAcceptsExchangesOk

`func (o *CreateShopReturnPolicyRequest) GetAcceptsExchangesOk() (*bool, bool)`

GetAcceptsExchangesOk returns a tuple with the AcceptsExchanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsExchanges

`func (o *CreateShopReturnPolicyRequest) SetAcceptsExchanges(v bool)`

SetAcceptsExchanges sets AcceptsExchanges field to given value.


### GetReturnDeadline

`func (o *CreateShopReturnPolicyRequest) GetReturnDeadline() int64`

GetReturnDeadline returns the ReturnDeadline field if non-nil, zero value otherwise.

### GetReturnDeadlineOk

`func (o *CreateShopReturnPolicyRequest) GetReturnDeadlineOk() (*int64, bool)`

GetReturnDeadlineOk returns a tuple with the ReturnDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDeadline

`func (o *CreateShopReturnPolicyRequest) SetReturnDeadline(v int64)`

SetReturnDeadline sets ReturnDeadline field to given value.

### HasReturnDeadline

`func (o *CreateShopReturnPolicyRequest) HasReturnDeadline() bool`

HasReturnDeadline returns a boolean if a field has been set.

### SetReturnDeadlineNil

`func (o *CreateShopReturnPolicyRequest) SetReturnDeadlineNil(b bool)`

 SetReturnDeadlineNil sets the value for ReturnDeadline to be an explicit nil

### UnsetReturnDeadline
`func (o *CreateShopReturnPolicyRequest) UnsetReturnDeadline()`

UnsetReturnDeadline ensures that no value is present for ReturnDeadline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


