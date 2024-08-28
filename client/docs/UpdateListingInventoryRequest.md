# UpdateListingInventoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]UpdateListingInventoryRequestProductsInner**](UpdateListingInventoryRequestProductsInner.md) | A JSON array of products available in a listing, even if only one product. All field names in the JSON blobs are lowercase. | 
**PriceOnProperty** | Pointer to **[]int64** | An array of unique [listing property](/documentation/reference#operation/getListingProperties) ID integers for the properties that change product prices, if any. For example, if you charge specific prices for different sized products in the same listing, then this array contains the property ID for size. | [optional] 
**QuantityOnProperty** | Pointer to **[]int64** | An array of unique [listing property](/documentation/reference#operation/getListingProperties) ID integers for the properties that change the quantity of the products, if any. For example, if you stock specific quantities of different colored products in the same listing, then this array contains the property ID for color. | [optional] 
**SkuOnProperty** | Pointer to **[]int64** | An array of unique [listing property](/documentation/reference#operation/getListingProperties) ID integers for the properties that change the product SKU, if any. For example, if you use specific skus for different colored products in the same listing, then this array contains the property ID for color. | [optional] 

## Methods

### NewUpdateListingInventoryRequest

`func NewUpdateListingInventoryRequest(products []UpdateListingInventoryRequestProductsInner, ) *UpdateListingInventoryRequest`

NewUpdateListingInventoryRequest instantiates a new UpdateListingInventoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListingInventoryRequestWithDefaults

`func NewUpdateListingInventoryRequestWithDefaults() *UpdateListingInventoryRequest`

NewUpdateListingInventoryRequestWithDefaults instantiates a new UpdateListingInventoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *UpdateListingInventoryRequest) GetProducts() []UpdateListingInventoryRequestProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *UpdateListingInventoryRequest) GetProductsOk() (*[]UpdateListingInventoryRequestProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *UpdateListingInventoryRequest) SetProducts(v []UpdateListingInventoryRequestProductsInner)`

SetProducts sets Products field to given value.


### GetPriceOnProperty

`func (o *UpdateListingInventoryRequest) GetPriceOnProperty() []int64`

GetPriceOnProperty returns the PriceOnProperty field if non-nil, zero value otherwise.

### GetPriceOnPropertyOk

`func (o *UpdateListingInventoryRequest) GetPriceOnPropertyOk() (*[]int64, bool)`

GetPriceOnPropertyOk returns a tuple with the PriceOnProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceOnProperty

`func (o *UpdateListingInventoryRequest) SetPriceOnProperty(v []int64)`

SetPriceOnProperty sets PriceOnProperty field to given value.

### HasPriceOnProperty

`func (o *UpdateListingInventoryRequest) HasPriceOnProperty() bool`

HasPriceOnProperty returns a boolean if a field has been set.

### GetQuantityOnProperty

`func (o *UpdateListingInventoryRequest) GetQuantityOnProperty() []int64`

GetQuantityOnProperty returns the QuantityOnProperty field if non-nil, zero value otherwise.

### GetQuantityOnPropertyOk

`func (o *UpdateListingInventoryRequest) GetQuantityOnPropertyOk() (*[]int64, bool)`

GetQuantityOnPropertyOk returns a tuple with the QuantityOnProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnProperty

`func (o *UpdateListingInventoryRequest) SetQuantityOnProperty(v []int64)`

SetQuantityOnProperty sets QuantityOnProperty field to given value.

### HasQuantityOnProperty

`func (o *UpdateListingInventoryRequest) HasQuantityOnProperty() bool`

HasQuantityOnProperty returns a boolean if a field has been set.

### GetSkuOnProperty

`func (o *UpdateListingInventoryRequest) GetSkuOnProperty() []int64`

GetSkuOnProperty returns the SkuOnProperty field if non-nil, zero value otherwise.

### GetSkuOnPropertyOk

`func (o *UpdateListingInventoryRequest) GetSkuOnPropertyOk() (*[]int64, bool)`

GetSkuOnPropertyOk returns a tuple with the SkuOnProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuOnProperty

`func (o *UpdateListingInventoryRequest) SetSkuOnProperty(v []int64)`

SetSkuOnProperty sets SkuOnProperty field to given value.

### HasSkuOnProperty

`func (o *UpdateListingInventoryRequest) HasSkuOnProperty() bool`

HasSkuOnProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


