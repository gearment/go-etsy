# BuyerTaxonomyNodeProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyId** | Pointer to **int64** | The unique numeric ID of this product property. | [optional] 
**Name** | Pointer to **string** | The name string for this taxonomy node. | [optional] 
**DisplayName** | Pointer to **string** | The human-readable product property name string. | [optional] 
**Scales** | Pointer to [**[]BuyerTaxonomyPropertyScale**](BuyerTaxonomyPropertyScale.md) | A list of available scales. | [optional] 
**IsRequired** | Pointer to **bool** | When true, listings assigned eligible taxonomy IDs require this property. | [optional] 
**SupportsAttributes** | Pointer to **bool** | When true, you can use this property in listing attributes. | [optional] 
**SupportsVariations** | Pointer to **bool** | When true, you can use this property in listing variations. | [optional] 
**IsMultivalued** | Pointer to **bool** | When true, you can assign multiple property values to this property | [optional] 
**MaxValuesAllowed** | Pointer to **NullableInt64** | When true, you can assign multiple property values to this property | [optional] 
**PossibleValues** | Pointer to [**[]BuyerTaxonomyPropertyValue**](BuyerTaxonomyPropertyValue.md) | A list of supported property value strings for this property. | [optional] 
**SelectedValues** | Pointer to [**[]BuyerTaxonomyPropertyValue**](BuyerTaxonomyPropertyValue.md) | A list of property value strings automatically and always selected for the given property. | [optional] 

## Methods

### NewBuyerTaxonomyNodeProperty

`func NewBuyerTaxonomyNodeProperty() *BuyerTaxonomyNodeProperty`

NewBuyerTaxonomyNodeProperty instantiates a new BuyerTaxonomyNodeProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerTaxonomyNodePropertyWithDefaults

`func NewBuyerTaxonomyNodePropertyWithDefaults() *BuyerTaxonomyNodeProperty`

NewBuyerTaxonomyNodePropertyWithDefaults instantiates a new BuyerTaxonomyNodeProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyId

`func (o *BuyerTaxonomyNodeProperty) GetPropertyId() int64`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *BuyerTaxonomyNodeProperty) GetPropertyIdOk() (*int64, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *BuyerTaxonomyNodeProperty) SetPropertyId(v int64)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *BuyerTaxonomyNodeProperty) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetName

`func (o *BuyerTaxonomyNodeProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuyerTaxonomyNodeProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuyerTaxonomyNodeProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BuyerTaxonomyNodeProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *BuyerTaxonomyNodeProperty) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BuyerTaxonomyNodeProperty) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BuyerTaxonomyNodeProperty) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BuyerTaxonomyNodeProperty) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetScales

`func (o *BuyerTaxonomyNodeProperty) GetScales() []BuyerTaxonomyPropertyScale`

GetScales returns the Scales field if non-nil, zero value otherwise.

### GetScalesOk

`func (o *BuyerTaxonomyNodeProperty) GetScalesOk() (*[]BuyerTaxonomyPropertyScale, bool)`

GetScalesOk returns a tuple with the Scales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScales

`func (o *BuyerTaxonomyNodeProperty) SetScales(v []BuyerTaxonomyPropertyScale)`

SetScales sets Scales field to given value.

### HasScales

`func (o *BuyerTaxonomyNodeProperty) HasScales() bool`

HasScales returns a boolean if a field has been set.

### GetIsRequired

`func (o *BuyerTaxonomyNodeProperty) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *BuyerTaxonomyNodeProperty) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *BuyerTaxonomyNodeProperty) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *BuyerTaxonomyNodeProperty) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetSupportsAttributes

`func (o *BuyerTaxonomyNodeProperty) GetSupportsAttributes() bool`

GetSupportsAttributes returns the SupportsAttributes field if non-nil, zero value otherwise.

### GetSupportsAttributesOk

`func (o *BuyerTaxonomyNodeProperty) GetSupportsAttributesOk() (*bool, bool)`

GetSupportsAttributesOk returns a tuple with the SupportsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAttributes

`func (o *BuyerTaxonomyNodeProperty) SetSupportsAttributes(v bool)`

SetSupportsAttributes sets SupportsAttributes field to given value.

### HasSupportsAttributes

`func (o *BuyerTaxonomyNodeProperty) HasSupportsAttributes() bool`

HasSupportsAttributes returns a boolean if a field has been set.

### GetSupportsVariations

`func (o *BuyerTaxonomyNodeProperty) GetSupportsVariations() bool`

GetSupportsVariations returns the SupportsVariations field if non-nil, zero value otherwise.

### GetSupportsVariationsOk

`func (o *BuyerTaxonomyNodeProperty) GetSupportsVariationsOk() (*bool, bool)`

GetSupportsVariationsOk returns a tuple with the SupportsVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsVariations

`func (o *BuyerTaxonomyNodeProperty) SetSupportsVariations(v bool)`

SetSupportsVariations sets SupportsVariations field to given value.

### HasSupportsVariations

`func (o *BuyerTaxonomyNodeProperty) HasSupportsVariations() bool`

HasSupportsVariations returns a boolean if a field has been set.

### GetIsMultivalued

`func (o *BuyerTaxonomyNodeProperty) GetIsMultivalued() bool`

GetIsMultivalued returns the IsMultivalued field if non-nil, zero value otherwise.

### GetIsMultivaluedOk

`func (o *BuyerTaxonomyNodeProperty) GetIsMultivaluedOk() (*bool, bool)`

GetIsMultivaluedOk returns a tuple with the IsMultivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultivalued

`func (o *BuyerTaxonomyNodeProperty) SetIsMultivalued(v bool)`

SetIsMultivalued sets IsMultivalued field to given value.

### HasIsMultivalued

`func (o *BuyerTaxonomyNodeProperty) HasIsMultivalued() bool`

HasIsMultivalued returns a boolean if a field has been set.

### GetMaxValuesAllowed

`func (o *BuyerTaxonomyNodeProperty) GetMaxValuesAllowed() int64`

GetMaxValuesAllowed returns the MaxValuesAllowed field if non-nil, zero value otherwise.

### GetMaxValuesAllowedOk

`func (o *BuyerTaxonomyNodeProperty) GetMaxValuesAllowedOk() (*int64, bool)`

GetMaxValuesAllowedOk returns a tuple with the MaxValuesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValuesAllowed

`func (o *BuyerTaxonomyNodeProperty) SetMaxValuesAllowed(v int64)`

SetMaxValuesAllowed sets MaxValuesAllowed field to given value.

### HasMaxValuesAllowed

`func (o *BuyerTaxonomyNodeProperty) HasMaxValuesAllowed() bool`

HasMaxValuesAllowed returns a boolean if a field has been set.

### SetMaxValuesAllowedNil

`func (o *BuyerTaxonomyNodeProperty) SetMaxValuesAllowedNil(b bool)`

 SetMaxValuesAllowedNil sets the value for MaxValuesAllowed to be an explicit nil

### UnsetMaxValuesAllowed
`func (o *BuyerTaxonomyNodeProperty) UnsetMaxValuesAllowed()`

UnsetMaxValuesAllowed ensures that no value is present for MaxValuesAllowed, not even an explicit nil
### GetPossibleValues

`func (o *BuyerTaxonomyNodeProperty) GetPossibleValues() []BuyerTaxonomyPropertyValue`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *BuyerTaxonomyNodeProperty) GetPossibleValuesOk() (*[]BuyerTaxonomyPropertyValue, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *BuyerTaxonomyNodeProperty) SetPossibleValues(v []BuyerTaxonomyPropertyValue)`

SetPossibleValues sets PossibleValues field to given value.

### HasPossibleValues

`func (o *BuyerTaxonomyNodeProperty) HasPossibleValues() bool`

HasPossibleValues returns a boolean if a field has been set.

### GetSelectedValues

`func (o *BuyerTaxonomyNodeProperty) GetSelectedValues() []BuyerTaxonomyPropertyValue`

GetSelectedValues returns the SelectedValues field if non-nil, zero value otherwise.

### GetSelectedValuesOk

`func (o *BuyerTaxonomyNodeProperty) GetSelectedValuesOk() (*[]BuyerTaxonomyPropertyValue, bool)`

GetSelectedValuesOk returns a tuple with the SelectedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedValues

`func (o *BuyerTaxonomyNodeProperty) SetSelectedValues(v []BuyerTaxonomyPropertyValue)`

SetSelectedValues sets SelectedValues field to given value.

### HasSelectedValues

`func (o *BuyerTaxonomyNodeProperty) HasSelectedValues() bool`

HasSelectedValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


