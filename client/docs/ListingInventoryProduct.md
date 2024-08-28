# ListingInventoryProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **int64** | The numeric ID for a specific [product](/documentation/reference#tag/ShopListing-Product) purchased from a listing. | [optional] 
**Sku** | Pointer to **string** | The SKU string for the product | [optional] 
**IsDeleted** | Pointer to **bool** | When true, someone deleted this product. | [optional] 
**Offerings** | Pointer to [**[]ListingInventoryProductOffering**](ListingInventoryProductOffering.md) | A list of product offering entries for this product. | [optional] 
**PropertyValues** | Pointer to [**[]ListingPropertyValue**](ListingPropertyValue.md) | A list of property value entries for this product. Note: parenthesis characters (&#x60;(&#x60; and &#x60;)&#x60;) are not allowed. | [optional] 

## Methods

### NewListingInventoryProduct

`func NewListingInventoryProduct() *ListingInventoryProduct`

NewListingInventoryProduct instantiates a new ListingInventoryProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingInventoryProductWithDefaults

`func NewListingInventoryProductWithDefaults() *ListingInventoryProduct`

NewListingInventoryProductWithDefaults instantiates a new ListingInventoryProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ListingInventoryProduct) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ListingInventoryProduct) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ListingInventoryProduct) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ListingInventoryProduct) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSku

`func (o *ListingInventoryProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ListingInventoryProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ListingInventoryProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ListingInventoryProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ListingInventoryProduct) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ListingInventoryProduct) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ListingInventoryProduct) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ListingInventoryProduct) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetOfferings

`func (o *ListingInventoryProduct) GetOfferings() []ListingInventoryProductOffering`

GetOfferings returns the Offerings field if non-nil, zero value otherwise.

### GetOfferingsOk

`func (o *ListingInventoryProduct) GetOfferingsOk() (*[]ListingInventoryProductOffering, bool)`

GetOfferingsOk returns a tuple with the Offerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferings

`func (o *ListingInventoryProduct) SetOfferings(v []ListingInventoryProductOffering)`

SetOfferings sets Offerings field to given value.

### HasOfferings

`func (o *ListingInventoryProduct) HasOfferings() bool`

HasOfferings returns a boolean if a field has been set.

### GetPropertyValues

`func (o *ListingInventoryProduct) GetPropertyValues() []ListingPropertyValue`

GetPropertyValues returns the PropertyValues field if non-nil, zero value otherwise.

### GetPropertyValuesOk

`func (o *ListingInventoryProduct) GetPropertyValuesOk() (*[]ListingPropertyValue, bool)`

GetPropertyValuesOk returns a tuple with the PropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValues

`func (o *ListingInventoryProduct) SetPropertyValues(v []ListingPropertyValue)`

SetPropertyValues sets PropertyValues field to given value.

### HasPropertyValues

`func (o *ListingInventoryProduct) HasPropertyValues() bool`

HasPropertyValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


