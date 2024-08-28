# TaxonomyNodeProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of results. | [optional] 
**Results** | Pointer to [**[]TaxonomyNodeProperty**](TaxonomyNodeProperty.md) | The list of requested resources. | [optional] 

## Methods

### NewTaxonomyNodeProperties

`func NewTaxonomyNodeProperties() *TaxonomyNodeProperties`

NewTaxonomyNodeProperties instantiates a new TaxonomyNodeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxonomyNodePropertiesWithDefaults

`func NewTaxonomyNodePropertiesWithDefaults() *TaxonomyNodeProperties`

NewTaxonomyNodePropertiesWithDefaults instantiates a new TaxonomyNodeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TaxonomyNodeProperties) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TaxonomyNodeProperties) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TaxonomyNodeProperties) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TaxonomyNodeProperties) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *TaxonomyNodeProperties) GetResults() []TaxonomyNodeProperty`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TaxonomyNodeProperties) GetResultsOk() (*[]TaxonomyNodeProperty, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TaxonomyNodeProperties) SetResults(v []TaxonomyNodeProperty)`

SetResults sets Results field to given value.

### HasResults

`func (o *TaxonomyNodeProperties) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


