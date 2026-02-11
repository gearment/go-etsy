# UpdateShopReadinessStateDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadinessState** | Pointer to [**CreateShopReadinessStateDefinitionRequestReadinessState**](CreateShopReadinessStateDefinitionRequestReadinessState.md) |  | [optional] 
**MinProcessingTime** | Pointer to **int64** | The minimum number of days or weeks for processing a specific product. | [optional] 
**MaxProcessingTime** | Pointer to **int64** | The maximum number of days or weeks for processing a specific product. | [optional] 
**ProcessingTimeUnit** | Pointer to [**CreateShopReadinessStateDefinitionRequestProcessingTimeUnit**](CreateShopReadinessStateDefinitionRequestProcessingTimeUnit.md) |  | [optional] [default to CREATESHOPREADINESSSTATEDEFINITIONREQUESTPROCESSINGTIMEUNIT_DAYS]

## Methods

### NewUpdateShopReadinessStateDefinitionRequest

`func NewUpdateShopReadinessStateDefinitionRequest() *UpdateShopReadinessStateDefinitionRequest`

NewUpdateShopReadinessStateDefinitionRequest instantiates a new UpdateShopReadinessStateDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShopReadinessStateDefinitionRequestWithDefaults

`func NewUpdateShopReadinessStateDefinitionRequestWithDefaults() *UpdateShopReadinessStateDefinitionRequest`

NewUpdateShopReadinessStateDefinitionRequestWithDefaults instantiates a new UpdateShopReadinessStateDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadinessState

`func (o *UpdateShopReadinessStateDefinitionRequest) GetReadinessState() CreateShopReadinessStateDefinitionRequestReadinessState`

GetReadinessState returns the ReadinessState field if non-nil, zero value otherwise.

### GetReadinessStateOk

`func (o *UpdateShopReadinessStateDefinitionRequest) GetReadinessStateOk() (*CreateShopReadinessStateDefinitionRequestReadinessState, bool)`

GetReadinessStateOk returns a tuple with the ReadinessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessState

`func (o *UpdateShopReadinessStateDefinitionRequest) SetReadinessState(v CreateShopReadinessStateDefinitionRequestReadinessState)`

SetReadinessState sets ReadinessState field to given value.

### HasReadinessState

`func (o *UpdateShopReadinessStateDefinitionRequest) HasReadinessState() bool`

HasReadinessState returns a boolean if a field has been set.

### GetMinProcessingTime

`func (o *UpdateShopReadinessStateDefinitionRequest) GetMinProcessingTime() int64`

GetMinProcessingTime returns the MinProcessingTime field if non-nil, zero value otherwise.

### GetMinProcessingTimeOk

`func (o *UpdateShopReadinessStateDefinitionRequest) GetMinProcessingTimeOk() (*int64, bool)`

GetMinProcessingTimeOk returns a tuple with the MinProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProcessingTime

`func (o *UpdateShopReadinessStateDefinitionRequest) SetMinProcessingTime(v int64)`

SetMinProcessingTime sets MinProcessingTime field to given value.

### HasMinProcessingTime

`func (o *UpdateShopReadinessStateDefinitionRequest) HasMinProcessingTime() bool`

HasMinProcessingTime returns a boolean if a field has been set.

### GetMaxProcessingTime

`func (o *UpdateShopReadinessStateDefinitionRequest) GetMaxProcessingTime() int64`

GetMaxProcessingTime returns the MaxProcessingTime field if non-nil, zero value otherwise.

### GetMaxProcessingTimeOk

`func (o *UpdateShopReadinessStateDefinitionRequest) GetMaxProcessingTimeOk() (*int64, bool)`

GetMaxProcessingTimeOk returns a tuple with the MaxProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProcessingTime

`func (o *UpdateShopReadinessStateDefinitionRequest) SetMaxProcessingTime(v int64)`

SetMaxProcessingTime sets MaxProcessingTime field to given value.

### HasMaxProcessingTime

`func (o *UpdateShopReadinessStateDefinitionRequest) HasMaxProcessingTime() bool`

HasMaxProcessingTime returns a boolean if a field has been set.

### GetProcessingTimeUnit

`func (o *UpdateShopReadinessStateDefinitionRequest) GetProcessingTimeUnit() CreateShopReadinessStateDefinitionRequestProcessingTimeUnit`

GetProcessingTimeUnit returns the ProcessingTimeUnit field if non-nil, zero value otherwise.

### GetProcessingTimeUnitOk

`func (o *UpdateShopReadinessStateDefinitionRequest) GetProcessingTimeUnitOk() (*CreateShopReadinessStateDefinitionRequestProcessingTimeUnit, bool)`

GetProcessingTimeUnitOk returns a tuple with the ProcessingTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeUnit

`func (o *UpdateShopReadinessStateDefinitionRequest) SetProcessingTimeUnit(v CreateShopReadinessStateDefinitionRequestProcessingTimeUnit)`

SetProcessingTimeUnit sets ProcessingTimeUnit field to given value.

### HasProcessingTimeUnit

`func (o *UpdateShopReadinessStateDefinitionRequest) HasProcessingTimeUnit() bool`

HasProcessingTimeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


