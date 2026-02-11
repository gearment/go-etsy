# ShopProcessingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | Pointer to **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | [optional] 
**ReadinessStateId** | Pointer to **int64** | The numeric ID of the [processing profile](/documentation/reference#operation/getShopReadinessStateDefinition) associated with the listing. Returned only when the listing is &#x60;active&#x60; and of type &#x60;physical&#x60;, and the endpoint is either shop-scoped (path contains &#x60;shop_id&#x60;) or a single-listing request such as &#x60;getListing&#x60;. For every other case this field can be null. | [optional] 
**ReadinessState** | Pointer to [**CreateShopReadinessStateDefinitionRequestReadinessState**](CreateShopReadinessStateDefinitionRequestReadinessState.md) |  | [optional] 
**MinProcessingDays** | Pointer to **int64** | The minimum number of days for processing a specific product. | [optional] 
**MaxProcessingDays** | Pointer to **int64** | The maximum number of days for processing a specific product. | [optional] 
**ProcessingDaysDisplayLabel** | Pointer to **string** | Translated display label string for processing days, for example \&quot;3 - 5 days\&quot;. | [optional] 

## Methods

### NewShopProcessingProfile

`func NewShopProcessingProfile() *ShopProcessingProfile`

NewShopProcessingProfile instantiates a new ShopProcessingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopProcessingProfileWithDefaults

`func NewShopProcessingProfileWithDefaults() *ShopProcessingProfile`

NewShopProcessingProfileWithDefaults instantiates a new ShopProcessingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *ShopProcessingProfile) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ShopProcessingProfile) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ShopProcessingProfile) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ShopProcessingProfile) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetReadinessStateId

`func (o *ShopProcessingProfile) GetReadinessStateId() int64`

GetReadinessStateId returns the ReadinessStateId field if non-nil, zero value otherwise.

### GetReadinessStateIdOk

`func (o *ShopProcessingProfile) GetReadinessStateIdOk() (*int64, bool)`

GetReadinessStateIdOk returns a tuple with the ReadinessStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessStateId

`func (o *ShopProcessingProfile) SetReadinessStateId(v int64)`

SetReadinessStateId sets ReadinessStateId field to given value.

### HasReadinessStateId

`func (o *ShopProcessingProfile) HasReadinessStateId() bool`

HasReadinessStateId returns a boolean if a field has been set.

### GetReadinessState

`func (o *ShopProcessingProfile) GetReadinessState() CreateShopReadinessStateDefinitionRequestReadinessState`

GetReadinessState returns the ReadinessState field if non-nil, zero value otherwise.

### GetReadinessStateOk

`func (o *ShopProcessingProfile) GetReadinessStateOk() (*CreateShopReadinessStateDefinitionRequestReadinessState, bool)`

GetReadinessStateOk returns a tuple with the ReadinessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessState

`func (o *ShopProcessingProfile) SetReadinessState(v CreateShopReadinessStateDefinitionRequestReadinessState)`

SetReadinessState sets ReadinessState field to given value.

### HasReadinessState

`func (o *ShopProcessingProfile) HasReadinessState() bool`

HasReadinessState returns a boolean if a field has been set.

### GetMinProcessingDays

`func (o *ShopProcessingProfile) GetMinProcessingDays() int64`

GetMinProcessingDays returns the MinProcessingDays field if non-nil, zero value otherwise.

### GetMinProcessingDaysOk

`func (o *ShopProcessingProfile) GetMinProcessingDaysOk() (*int64, bool)`

GetMinProcessingDaysOk returns a tuple with the MinProcessingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProcessingDays

`func (o *ShopProcessingProfile) SetMinProcessingDays(v int64)`

SetMinProcessingDays sets MinProcessingDays field to given value.

### HasMinProcessingDays

`func (o *ShopProcessingProfile) HasMinProcessingDays() bool`

HasMinProcessingDays returns a boolean if a field has been set.

### GetMaxProcessingDays

`func (o *ShopProcessingProfile) GetMaxProcessingDays() int64`

GetMaxProcessingDays returns the MaxProcessingDays field if non-nil, zero value otherwise.

### GetMaxProcessingDaysOk

`func (o *ShopProcessingProfile) GetMaxProcessingDaysOk() (*int64, bool)`

GetMaxProcessingDaysOk returns a tuple with the MaxProcessingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProcessingDays

`func (o *ShopProcessingProfile) SetMaxProcessingDays(v int64)`

SetMaxProcessingDays sets MaxProcessingDays field to given value.

### HasMaxProcessingDays

`func (o *ShopProcessingProfile) HasMaxProcessingDays() bool`

HasMaxProcessingDays returns a boolean if a field has been set.

### GetProcessingDaysDisplayLabel

`func (o *ShopProcessingProfile) GetProcessingDaysDisplayLabel() string`

GetProcessingDaysDisplayLabel returns the ProcessingDaysDisplayLabel field if non-nil, zero value otherwise.

### GetProcessingDaysDisplayLabelOk

`func (o *ShopProcessingProfile) GetProcessingDaysDisplayLabelOk() (*string, bool)`

GetProcessingDaysDisplayLabelOk returns a tuple with the ProcessingDaysDisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingDaysDisplayLabel

`func (o *ShopProcessingProfile) SetProcessingDaysDisplayLabel(v string)`

SetProcessingDaysDisplayLabel sets ProcessingDaysDisplayLabel field to given value.

### HasProcessingDaysDisplayLabel

`func (o *ShopProcessingProfile) HasProcessingDaysDisplayLabel() bool`

HasProcessingDaysDisplayLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


