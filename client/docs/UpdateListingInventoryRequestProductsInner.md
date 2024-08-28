# UpdateListingInventoryRequestProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **NullableString** | The SKU string for the product | [optional] 
**PropertyValues** | Pointer to [**[]UpdateListingInventoryRequestProductsInnerPropertyValuesInner**](UpdateListingInventoryRequestProductsInnerPropertyValuesInner.md) | A list of property value entries for this product. Note: parenthesis characters (&#x60;(&#x60; and &#x60;)&#x60;) are not allowed. | [optional] 
**Offerings** | [**[]UpdateListingInventoryRequestProductsInnerOfferingsInner**](UpdateListingInventoryRequestProductsInnerOfferingsInner.md) | A list of product offering entries for this product. | 

## Methods

### NewUpdateListingInventoryRequestProductsInner

`func NewUpdateListingInventoryRequestProductsInner(offerings []UpdateListingInventoryRequestProductsInnerOfferingsInner, ) *UpdateListingInventoryRequestProductsInner`

NewUpdateListingInventoryRequestProductsInner instantiates a new UpdateListingInventoryRequestProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListingInventoryRequestProductsInnerWithDefaults

`func NewUpdateListingInventoryRequestProductsInnerWithDefaults() *UpdateListingInventoryRequestProductsInner`

NewUpdateListingInventoryRequestProductsInnerWithDefaults instantiates a new UpdateListingInventoryRequestProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *UpdateListingInventoryRequestProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *UpdateListingInventoryRequestProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *UpdateListingInventoryRequestProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *UpdateListingInventoryRequestProductsInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *UpdateListingInventoryRequestProductsInner) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *UpdateListingInventoryRequestProductsInner) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetPropertyValues

`func (o *UpdateListingInventoryRequestProductsInner) GetPropertyValues() []UpdateListingInventoryRequestProductsInnerPropertyValuesInner`

GetPropertyValues returns the PropertyValues field if non-nil, zero value otherwise.

### GetPropertyValuesOk

`func (o *UpdateListingInventoryRequestProductsInner) GetPropertyValuesOk() (*[]UpdateListingInventoryRequestProductsInnerPropertyValuesInner, bool)`

GetPropertyValuesOk returns a tuple with the PropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValues

`func (o *UpdateListingInventoryRequestProductsInner) SetPropertyValues(v []UpdateListingInventoryRequestProductsInnerPropertyValuesInner)`

SetPropertyValues sets PropertyValues field to given value.

### HasPropertyValues

`func (o *UpdateListingInventoryRequestProductsInner) HasPropertyValues() bool`

HasPropertyValues returns a boolean if a field has been set.

### GetOfferings

`func (o *UpdateListingInventoryRequestProductsInner) GetOfferings() []UpdateListingInventoryRequestProductsInnerOfferingsInner`

GetOfferings returns the Offerings field if non-nil, zero value otherwise.

### GetOfferingsOk

`func (o *UpdateListingInventoryRequestProductsInner) GetOfferingsOk() (*[]UpdateListingInventoryRequestProductsInnerOfferingsInner, bool)`

GetOfferingsOk returns a tuple with the Offerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferings

`func (o *UpdateListingInventoryRequestProductsInner) SetOfferings(v []UpdateListingInventoryRequestProductsInnerOfferingsInner)`

SetOfferings sets Offerings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


