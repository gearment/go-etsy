# CreateShopReadinessStateDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadinessState** | [**CreateShopReadinessStateDefinitionRequestReadinessState**](CreateShopReadinessStateDefinitionRequestReadinessState.md) |  | 
**MinProcessingTime** | **int64** | The minimum number of days or weeks for processing a specific product. | 
**MaxProcessingTime** | **int64** | The maximum number of days or weeks for processing a specific product. | 
**ProcessingTimeUnit** | Pointer to [**CreateShopReadinessStateDefinitionRequestProcessingTimeUnit**](CreateShopReadinessStateDefinitionRequestProcessingTimeUnit.md) |  | [optional] [default to CREATESHOPREADINESSSTATEDEFINITIONREQUESTPROCESSINGTIMEUNIT_DAYS]

## Methods

### NewCreateShopReadinessStateDefinitionRequest

`func NewCreateShopReadinessStateDefinitionRequest(readinessState CreateShopReadinessStateDefinitionRequestReadinessState, minProcessingTime int64, maxProcessingTime int64, ) *CreateShopReadinessStateDefinitionRequest`

NewCreateShopReadinessStateDefinitionRequest instantiates a new CreateShopReadinessStateDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShopReadinessStateDefinitionRequestWithDefaults

`func NewCreateShopReadinessStateDefinitionRequestWithDefaults() *CreateShopReadinessStateDefinitionRequest`

NewCreateShopReadinessStateDefinitionRequestWithDefaults instantiates a new CreateShopReadinessStateDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadinessState

`func (o *CreateShopReadinessStateDefinitionRequest) GetReadinessState() CreateShopReadinessStateDefinitionRequestReadinessState`

GetReadinessState returns the ReadinessState field if non-nil, zero value otherwise.

### GetReadinessStateOk

`func (o *CreateShopReadinessStateDefinitionRequest) GetReadinessStateOk() (*CreateShopReadinessStateDefinitionRequestReadinessState, bool)`

GetReadinessStateOk returns a tuple with the ReadinessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessState

`func (o *CreateShopReadinessStateDefinitionRequest) SetReadinessState(v CreateShopReadinessStateDefinitionRequestReadinessState)`

SetReadinessState sets ReadinessState field to given value.


### GetMinProcessingTime

`func (o *CreateShopReadinessStateDefinitionRequest) GetMinProcessingTime() int64`

GetMinProcessingTime returns the MinProcessingTime field if non-nil, zero value otherwise.

### GetMinProcessingTimeOk

`func (o *CreateShopReadinessStateDefinitionRequest) GetMinProcessingTimeOk() (*int64, bool)`

GetMinProcessingTimeOk returns a tuple with the MinProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProcessingTime

`func (o *CreateShopReadinessStateDefinitionRequest) SetMinProcessingTime(v int64)`

SetMinProcessingTime sets MinProcessingTime field to given value.


### GetMaxProcessingTime

`func (o *CreateShopReadinessStateDefinitionRequest) GetMaxProcessingTime() int64`

GetMaxProcessingTime returns the MaxProcessingTime field if non-nil, zero value otherwise.

### GetMaxProcessingTimeOk

`func (o *CreateShopReadinessStateDefinitionRequest) GetMaxProcessingTimeOk() (*int64, bool)`

GetMaxProcessingTimeOk returns a tuple with the MaxProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProcessingTime

`func (o *CreateShopReadinessStateDefinitionRequest) SetMaxProcessingTime(v int64)`

SetMaxProcessingTime sets MaxProcessingTime field to given value.


### GetProcessingTimeUnit

`func (o *CreateShopReadinessStateDefinitionRequest) GetProcessingTimeUnit() CreateShopReadinessStateDefinitionRequestProcessingTimeUnit`

GetProcessingTimeUnit returns the ProcessingTimeUnit field if non-nil, zero value otherwise.

### GetProcessingTimeUnitOk

`func (o *CreateShopReadinessStateDefinitionRequest) GetProcessingTimeUnitOk() (*CreateShopReadinessStateDefinitionRequestProcessingTimeUnit, bool)`

GetProcessingTimeUnitOk returns a tuple with the ProcessingTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeUnit

`func (o *CreateShopReadinessStateDefinitionRequest) SetProcessingTimeUnit(v CreateShopReadinessStateDefinitionRequestProcessingTimeUnit)`

SetProcessingTimeUnit sets ProcessingTimeUnit field to given value.

### HasProcessingTimeUnit

`func (o *CreateShopReadinessStateDefinitionRequest) HasProcessingTimeUnit() bool`

HasProcessingTimeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


