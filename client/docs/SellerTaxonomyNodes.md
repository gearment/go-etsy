# SellerTaxonomyNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of results. | [optional] 
**Results** | Pointer to [**[]SellerTaxonomyNode**](SellerTaxonomyNode.md) | The list of requested resources. | [optional] 

## Methods

### NewSellerTaxonomyNodes

`func NewSellerTaxonomyNodes() *SellerTaxonomyNodes`

NewSellerTaxonomyNodes instantiates a new SellerTaxonomyNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellerTaxonomyNodesWithDefaults

`func NewSellerTaxonomyNodesWithDefaults() *SellerTaxonomyNodes`

NewSellerTaxonomyNodesWithDefaults instantiates a new SellerTaxonomyNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SellerTaxonomyNodes) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SellerTaxonomyNodes) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SellerTaxonomyNodes) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SellerTaxonomyNodes) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *SellerTaxonomyNodes) GetResults() []SellerTaxonomyNode`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SellerTaxonomyNodes) GetResultsOk() (*[]SellerTaxonomyNode, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SellerTaxonomyNodes) SetResults(v []SellerTaxonomyNode)`

SetResults sets Results field to given value.

### HasResults

`func (o *SellerTaxonomyNodes) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


