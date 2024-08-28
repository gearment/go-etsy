# TaxonomyNodeProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyId** | Pointer to **int64** | The unique numeric ID of this product property. | [optional] 
**Name** | Pointer to **string** | The name string for this taxonomy node. | [optional] 
**DisplayName** | Pointer to **string** | The human-readable product property name string. | [optional] 
**Scales** | Pointer to [**[]TaxonomyPropertyScale**](TaxonomyPropertyScale.md) | A list of available scales. | [optional] 
**IsRequired** | Pointer to **bool** | When true, listings assigned eligible taxonomy IDs require this property. | [optional] 
**SupportsAttributes** | Pointer to **bool** | When true, you can use this property in listing attributes. | [optional] 
**SupportsVariations** | Pointer to **bool** | When true, you can use this property in listing variations. | [optional] 
**IsMultivalued** | Pointer to **bool** | When true, you can assign multiple property values to this property | [optional] 
**MaxValuesAllowed** | Pointer to **NullableInt64** | When true, you can assign multiple property values to this property | [optional] 
**PossibleValues** | Pointer to [**[]TaxonomyPropertyValue**](TaxonomyPropertyValue.md) | A list of supported property value strings for this property. | [optional] 
**SelectedValues** | Pointer to [**[]TaxonomyPropertyValue**](TaxonomyPropertyValue.md) | A list of property value strings automatically and always selected for the given property. | [optional] 

## Methods

### NewTaxonomyNodeProperty

`func NewTaxonomyNodeProperty() *TaxonomyNodeProperty`

NewTaxonomyNodeProperty instantiates a new TaxonomyNodeProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxonomyNodePropertyWithDefaults

`func NewTaxonomyNodePropertyWithDefaults() *TaxonomyNodeProperty`

NewTaxonomyNodePropertyWithDefaults instantiates a new TaxonomyNodeProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyId

`func (o *TaxonomyNodeProperty) GetPropertyId() int64`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *TaxonomyNodeProperty) GetPropertyIdOk() (*int64, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *TaxonomyNodeProperty) SetPropertyId(v int64)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *TaxonomyNodeProperty) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetName

`func (o *TaxonomyNodeProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxonomyNodeProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxonomyNodeProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxonomyNodeProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *TaxonomyNodeProperty) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TaxonomyNodeProperty) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TaxonomyNodeProperty) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TaxonomyNodeProperty) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetScales

`func (o *TaxonomyNodeProperty) GetScales() []TaxonomyPropertyScale`

GetScales returns the Scales field if non-nil, zero value otherwise.

### GetScalesOk

`func (o *TaxonomyNodeProperty) GetScalesOk() (*[]TaxonomyPropertyScale, bool)`

GetScalesOk returns a tuple with the Scales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScales

`func (o *TaxonomyNodeProperty) SetScales(v []TaxonomyPropertyScale)`

SetScales sets Scales field to given value.

### HasScales

`func (o *TaxonomyNodeProperty) HasScales() bool`

HasScales returns a boolean if a field has been set.

### GetIsRequired

`func (o *TaxonomyNodeProperty) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *TaxonomyNodeProperty) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *TaxonomyNodeProperty) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *TaxonomyNodeProperty) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetSupportsAttributes

`func (o *TaxonomyNodeProperty) GetSupportsAttributes() bool`

GetSupportsAttributes returns the SupportsAttributes field if non-nil, zero value otherwise.

### GetSupportsAttributesOk

`func (o *TaxonomyNodeProperty) GetSupportsAttributesOk() (*bool, bool)`

GetSupportsAttributesOk returns a tuple with the SupportsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAttributes

`func (o *TaxonomyNodeProperty) SetSupportsAttributes(v bool)`

SetSupportsAttributes sets SupportsAttributes field to given value.

### HasSupportsAttributes

`func (o *TaxonomyNodeProperty) HasSupportsAttributes() bool`

HasSupportsAttributes returns a boolean if a field has been set.

### GetSupportsVariations

`func (o *TaxonomyNodeProperty) GetSupportsVariations() bool`

GetSupportsVariations returns the SupportsVariations field if non-nil, zero value otherwise.

### GetSupportsVariationsOk

`func (o *TaxonomyNodeProperty) GetSupportsVariationsOk() (*bool, bool)`

GetSupportsVariationsOk returns a tuple with the SupportsVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsVariations

`func (o *TaxonomyNodeProperty) SetSupportsVariations(v bool)`

SetSupportsVariations sets SupportsVariations field to given value.

### HasSupportsVariations

`func (o *TaxonomyNodeProperty) HasSupportsVariations() bool`

HasSupportsVariations returns a boolean if a field has been set.

### GetIsMultivalued

`func (o *TaxonomyNodeProperty) GetIsMultivalued() bool`

GetIsMultivalued returns the IsMultivalued field if non-nil, zero value otherwise.

### GetIsMultivaluedOk

`func (o *TaxonomyNodeProperty) GetIsMultivaluedOk() (*bool, bool)`

GetIsMultivaluedOk returns a tuple with the IsMultivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultivalued

`func (o *TaxonomyNodeProperty) SetIsMultivalued(v bool)`

SetIsMultivalued sets IsMultivalued field to given value.

### HasIsMultivalued

`func (o *TaxonomyNodeProperty) HasIsMultivalued() bool`

HasIsMultivalued returns a boolean if a field has been set.

### GetMaxValuesAllowed

`func (o *TaxonomyNodeProperty) GetMaxValuesAllowed() int64`

GetMaxValuesAllowed returns the MaxValuesAllowed field if non-nil, zero value otherwise.

### GetMaxValuesAllowedOk

`func (o *TaxonomyNodeProperty) GetMaxValuesAllowedOk() (*int64, bool)`

GetMaxValuesAllowedOk returns a tuple with the MaxValuesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValuesAllowed

`func (o *TaxonomyNodeProperty) SetMaxValuesAllowed(v int64)`

SetMaxValuesAllowed sets MaxValuesAllowed field to given value.

### HasMaxValuesAllowed

`func (o *TaxonomyNodeProperty) HasMaxValuesAllowed() bool`

HasMaxValuesAllowed returns a boolean if a field has been set.

### SetMaxValuesAllowedNil

`func (o *TaxonomyNodeProperty) SetMaxValuesAllowedNil(b bool)`

 SetMaxValuesAllowedNil sets the value for MaxValuesAllowed to be an explicit nil

### UnsetMaxValuesAllowed
`func (o *TaxonomyNodeProperty) UnsetMaxValuesAllowed()`

UnsetMaxValuesAllowed ensures that no value is present for MaxValuesAllowed, not even an explicit nil
### GetPossibleValues

`func (o *TaxonomyNodeProperty) GetPossibleValues() []TaxonomyPropertyValue`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *TaxonomyNodeProperty) GetPossibleValuesOk() (*[]TaxonomyPropertyValue, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *TaxonomyNodeProperty) SetPossibleValues(v []TaxonomyPropertyValue)`

SetPossibleValues sets PossibleValues field to given value.

### HasPossibleValues

`func (o *TaxonomyNodeProperty) HasPossibleValues() bool`

HasPossibleValues returns a boolean if a field has been set.

### GetSelectedValues

`func (o *TaxonomyNodeProperty) GetSelectedValues() []TaxonomyPropertyValue`

GetSelectedValues returns the SelectedValues field if non-nil, zero value otherwise.

### GetSelectedValuesOk

`func (o *TaxonomyNodeProperty) GetSelectedValuesOk() (*[]TaxonomyPropertyValue, bool)`

GetSelectedValuesOk returns a tuple with the SelectedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedValues

`func (o *TaxonomyNodeProperty) SetSelectedValues(v []TaxonomyPropertyValue)`

SetSelectedValues sets SelectedValues field to given value.

### HasSelectedValues

`func (o *TaxonomyNodeProperty) HasSelectedValues() bool`

HasSelectedValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


