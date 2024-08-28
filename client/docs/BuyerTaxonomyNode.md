# BuyerTaxonomyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique numeric ID of an Etsy taxonomy node, which is a metadata category for listings organized into the seller taxonomy hierarchy tree. For example, the \\\&quot;shoes\\\&quot; taxonomy node (ID: 1429, level: 1) is higher in the hierarchy than \\\&quot;girls&#39; shoes\\\&quot; (ID: 1440, level: 2). The taxonomy nodes assigned to a listing support access to specific standardized product scales and properties. For example, listings assigned the taxonomy nodes \\\&quot;shoes\\\&quot; or \\\&quot;girls&#39; shoes\\\&quot; support access to the \\\&quot;EU\\\&quot; shoe size scale with its associated property names and IDs for EU shoe sizes, such as property &#x60;value_id&#x60;:\\\&quot;1394\\\&quot;, and &#x60;name&#x60;:\\\&quot;38\\\&quot;. | [optional] 
**Level** | Pointer to **int64** | The integer depth of this taxonomy node in the seller taxonomy tree, with roots at level 0. | [optional] 
**Name** | Pointer to **string** | The name string for this taxonomy node. | [optional] 
**ParentId** | Pointer to **NullableInt64** | The numeric taxonomy ID of the parent of this node. | [optional] 
**Children** | Pointer to [**[]BuyerTaxonomyNode**](BuyerTaxonomyNode.md) | An array of taxonomy nodes for all the direct children of this taxonomy node in the seller taxonomy tree. | [optional] 
**FullPathTaxonomyIds** | Pointer to **[]int64** | An array of &#x60;taxonomy_id&#x60;s including this node and all of its direct parents in the seller taxonomy tree up to a root node. They are listed in order from root to leaf. | [optional] 

## Methods

### NewBuyerTaxonomyNode

`func NewBuyerTaxonomyNode() *BuyerTaxonomyNode`

NewBuyerTaxonomyNode instantiates a new BuyerTaxonomyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerTaxonomyNodeWithDefaults

`func NewBuyerTaxonomyNodeWithDefaults() *BuyerTaxonomyNode`

NewBuyerTaxonomyNodeWithDefaults instantiates a new BuyerTaxonomyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BuyerTaxonomyNode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuyerTaxonomyNode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuyerTaxonomyNode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BuyerTaxonomyNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *BuyerTaxonomyNode) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *BuyerTaxonomyNode) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *BuyerTaxonomyNode) SetLevel(v int64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *BuyerTaxonomyNode) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *BuyerTaxonomyNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuyerTaxonomyNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuyerTaxonomyNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BuyerTaxonomyNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *BuyerTaxonomyNode) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BuyerTaxonomyNode) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BuyerTaxonomyNode) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BuyerTaxonomyNode) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *BuyerTaxonomyNode) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *BuyerTaxonomyNode) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetChildren

`func (o *BuyerTaxonomyNode) GetChildren() []BuyerTaxonomyNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *BuyerTaxonomyNode) GetChildrenOk() (*[]BuyerTaxonomyNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *BuyerTaxonomyNode) SetChildren(v []BuyerTaxonomyNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *BuyerTaxonomyNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetFullPathTaxonomyIds

`func (o *BuyerTaxonomyNode) GetFullPathTaxonomyIds() []int64`

GetFullPathTaxonomyIds returns the FullPathTaxonomyIds field if non-nil, zero value otherwise.

### GetFullPathTaxonomyIdsOk

`func (o *BuyerTaxonomyNode) GetFullPathTaxonomyIdsOk() (*[]int64, bool)`

GetFullPathTaxonomyIdsOk returns a tuple with the FullPathTaxonomyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPathTaxonomyIds

`func (o *BuyerTaxonomyNode) SetFullPathTaxonomyIds(v []int64)`

SetFullPathTaxonomyIds sets FullPathTaxonomyIds field to given value.

### HasFullPathTaxonomyIds

`func (o *BuyerTaxonomyNode) HasFullPathTaxonomyIds() bool`

HasFullPathTaxonomyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


