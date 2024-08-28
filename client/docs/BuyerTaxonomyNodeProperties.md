# BuyerTaxonomyNodeProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of results. | [optional] 
**Results** | Pointer to [**[]BuyerTaxonomyNodeProperty**](BuyerTaxonomyNodeProperty.md) | The list of requested resources. | [optional] 

## Methods

### NewBuyerTaxonomyNodeProperties

`func NewBuyerTaxonomyNodeProperties() *BuyerTaxonomyNodeProperties`

NewBuyerTaxonomyNodeProperties instantiates a new BuyerTaxonomyNodeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerTaxonomyNodePropertiesWithDefaults

`func NewBuyerTaxonomyNodePropertiesWithDefaults() *BuyerTaxonomyNodeProperties`

NewBuyerTaxonomyNodePropertiesWithDefaults instantiates a new BuyerTaxonomyNodeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BuyerTaxonomyNodeProperties) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BuyerTaxonomyNodeProperties) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BuyerTaxonomyNodeProperties) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *BuyerTaxonomyNodeProperties) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *BuyerTaxonomyNodeProperties) GetResults() []BuyerTaxonomyNodeProperty`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BuyerTaxonomyNodeProperties) GetResultsOk() (*[]BuyerTaxonomyNodeProperty, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BuyerTaxonomyNodeProperties) SetResults(v []BuyerTaxonomyNodeProperty)`

SetResults sets Results field to given value.

### HasResults

`func (o *BuyerTaxonomyNodeProperties) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


