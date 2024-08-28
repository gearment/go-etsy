# TransactionVariations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyId** | Pointer to **int64** | The variation property ID. | [optional] 
**ValueId** | Pointer to **NullableInt64** | The ID of the variation value selected. | [optional] 
**FormattedName** | Pointer to **string** | Formatted name of the variation. | [optional] 
**FormattedValue** | Pointer to **string** | Value of the variation entered by the buyer. | [optional] 

## Methods

### NewTransactionVariations

`func NewTransactionVariations() *TransactionVariations`

NewTransactionVariations instantiates a new TransactionVariations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionVariationsWithDefaults

`func NewTransactionVariationsWithDefaults() *TransactionVariations`

NewTransactionVariationsWithDefaults instantiates a new TransactionVariations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyId

`func (o *TransactionVariations) GetPropertyId() int64`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *TransactionVariations) GetPropertyIdOk() (*int64, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *TransactionVariations) SetPropertyId(v int64)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *TransactionVariations) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetValueId

`func (o *TransactionVariations) GetValueId() int64`

GetValueId returns the ValueId field if non-nil, zero value otherwise.

### GetValueIdOk

`func (o *TransactionVariations) GetValueIdOk() (*int64, bool)`

GetValueIdOk returns a tuple with the ValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueId

`func (o *TransactionVariations) SetValueId(v int64)`

SetValueId sets ValueId field to given value.

### HasValueId

`func (o *TransactionVariations) HasValueId() bool`

HasValueId returns a boolean if a field has been set.

### SetValueIdNil

`func (o *TransactionVariations) SetValueIdNil(b bool)`

 SetValueIdNil sets the value for ValueId to be an explicit nil

### UnsetValueId
`func (o *TransactionVariations) UnsetValueId()`

UnsetValueId ensures that no value is present for ValueId, not even an explicit nil
### GetFormattedName

`func (o *TransactionVariations) GetFormattedName() string`

GetFormattedName returns the FormattedName field if non-nil, zero value otherwise.

### GetFormattedNameOk

`func (o *TransactionVariations) GetFormattedNameOk() (*string, bool)`

GetFormattedNameOk returns a tuple with the FormattedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedName

`func (o *TransactionVariations) SetFormattedName(v string)`

SetFormattedName sets FormattedName field to given value.

### HasFormattedName

`func (o *TransactionVariations) HasFormattedName() bool`

HasFormattedName returns a boolean if a field has been set.

### GetFormattedValue

`func (o *TransactionVariations) GetFormattedValue() string`

GetFormattedValue returns the FormattedValue field if non-nil, zero value otherwise.

### GetFormattedValueOk

`func (o *TransactionVariations) GetFormattedValueOk() (*string, bool)`

GetFormattedValueOk returns a tuple with the FormattedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedValue

`func (o *TransactionVariations) SetFormattedValue(v string)`

SetFormattedValue sets FormattedValue field to given value.

### HasFormattedValue

`func (o *TransactionVariations) HasFormattedValue() bool`

HasFormattedValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


