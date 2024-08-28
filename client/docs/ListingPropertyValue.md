# ListingPropertyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyId** | Pointer to **int64** | The numeric ID of the Property. | [optional] 
**PropertyName** | Pointer to **NullableString** | The name of the Property. | [optional] 
**ScaleId** | Pointer to **NullableInt64** | The numeric ID of the scale (if any). | [optional] 
**ScaleName** | Pointer to **NullableString** | The label used to describe the chosen scale (if any). | [optional] 
**ValueIds** | Pointer to **[]int64** | The numeric IDs of the Property values | [optional] 
**Values** | Pointer to **[]string** | The Property values | [optional] 

## Methods

### NewListingPropertyValue

`func NewListingPropertyValue() *ListingPropertyValue`

NewListingPropertyValue instantiates a new ListingPropertyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingPropertyValueWithDefaults

`func NewListingPropertyValueWithDefaults() *ListingPropertyValue`

NewListingPropertyValueWithDefaults instantiates a new ListingPropertyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyId

`func (o *ListingPropertyValue) GetPropertyId() int64`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *ListingPropertyValue) GetPropertyIdOk() (*int64, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *ListingPropertyValue) SetPropertyId(v int64)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *ListingPropertyValue) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetPropertyName

`func (o *ListingPropertyValue) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *ListingPropertyValue) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *ListingPropertyValue) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *ListingPropertyValue) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### SetPropertyNameNil

`func (o *ListingPropertyValue) SetPropertyNameNil(b bool)`

 SetPropertyNameNil sets the value for PropertyName to be an explicit nil

### UnsetPropertyName
`func (o *ListingPropertyValue) UnsetPropertyName()`

UnsetPropertyName ensures that no value is present for PropertyName, not even an explicit nil
### GetScaleId

`func (o *ListingPropertyValue) GetScaleId() int64`

GetScaleId returns the ScaleId field if non-nil, zero value otherwise.

### GetScaleIdOk

`func (o *ListingPropertyValue) GetScaleIdOk() (*int64, bool)`

GetScaleIdOk returns a tuple with the ScaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleId

`func (o *ListingPropertyValue) SetScaleId(v int64)`

SetScaleId sets ScaleId field to given value.

### HasScaleId

`func (o *ListingPropertyValue) HasScaleId() bool`

HasScaleId returns a boolean if a field has been set.

### SetScaleIdNil

`func (o *ListingPropertyValue) SetScaleIdNil(b bool)`

 SetScaleIdNil sets the value for ScaleId to be an explicit nil

### UnsetScaleId
`func (o *ListingPropertyValue) UnsetScaleId()`

UnsetScaleId ensures that no value is present for ScaleId, not even an explicit nil
### GetScaleName

`func (o *ListingPropertyValue) GetScaleName() string`

GetScaleName returns the ScaleName field if non-nil, zero value otherwise.

### GetScaleNameOk

`func (o *ListingPropertyValue) GetScaleNameOk() (*string, bool)`

GetScaleNameOk returns a tuple with the ScaleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleName

`func (o *ListingPropertyValue) SetScaleName(v string)`

SetScaleName sets ScaleName field to given value.

### HasScaleName

`func (o *ListingPropertyValue) HasScaleName() bool`

HasScaleName returns a boolean if a field has been set.

### SetScaleNameNil

`func (o *ListingPropertyValue) SetScaleNameNil(b bool)`

 SetScaleNameNil sets the value for ScaleName to be an explicit nil

### UnsetScaleName
`func (o *ListingPropertyValue) UnsetScaleName()`

UnsetScaleName ensures that no value is present for ScaleName, not even an explicit nil
### GetValueIds

`func (o *ListingPropertyValue) GetValueIds() []int64`

GetValueIds returns the ValueIds field if non-nil, zero value otherwise.

### GetValueIdsOk

`func (o *ListingPropertyValue) GetValueIdsOk() (*[]int64, bool)`

GetValueIdsOk returns a tuple with the ValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueIds

`func (o *ListingPropertyValue) SetValueIds(v []int64)`

SetValueIds sets ValueIds field to given value.

### HasValueIds

`func (o *ListingPropertyValue) HasValueIds() bool`

HasValueIds returns a boolean if a field has been set.

### GetValues

`func (o *ListingPropertyValue) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ListingPropertyValue) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ListingPropertyValue) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ListingPropertyValue) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


