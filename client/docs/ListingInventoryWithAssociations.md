# ListingInventoryWithAssociations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]ListingInventoryProduct**](ListingInventoryProduct.md) | A JSON array of products available in a listing, even if only one product. All field names in the JSON blobs are lowercase. | [optional] 
**PriceOnProperty** | Pointer to **[]int64** | An array of unique [listing property](/documentation/reference#operation/getListingProperties) ID integers for the properties that change product prices, if any. For example, if you charge specific prices for different sized products in the same listing, then this array contains the property ID for size. | [optional] 
**QuantityOnProperty** | Pointer to **[]int64** | An array of unique [listing property](/documentation/reference#operation/getListingProperties) ID integers for the properties that change the quantity of the products, if any. For example, if you stock specific quantities of different colored products in the same listing, then this array contains the property ID for color. | [optional] 
**SkuOnProperty** | Pointer to **[]int64** | An array of unique [listing property](/documentation/reference#operation/getListingProperties) ID integers for the properties that change the product SKU, if any. For example, if you use specific skus for different colored products in the same listing, then this array contains the property ID for color. | [optional] 
**Listing** | Pointer to [**ShopListing**](ShopListing.md) | An enumerated string that attaches a valid association. Default value is null. | [optional] 

## Methods

### NewListingInventoryWithAssociations

`func NewListingInventoryWithAssociations() *ListingInventoryWithAssociations`

NewListingInventoryWithAssociations instantiates a new ListingInventoryWithAssociations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingInventoryWithAssociationsWithDefaults

`func NewListingInventoryWithAssociationsWithDefaults() *ListingInventoryWithAssociations`

NewListingInventoryWithAssociationsWithDefaults instantiates a new ListingInventoryWithAssociations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *ListingInventoryWithAssociations) GetProducts() []ListingInventoryProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ListingInventoryWithAssociations) GetProductsOk() (*[]ListingInventoryProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ListingInventoryWithAssociations) SetProducts(v []ListingInventoryProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ListingInventoryWithAssociations) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPriceOnProperty

`func (o *ListingInventoryWithAssociations) GetPriceOnProperty() []int64`

GetPriceOnProperty returns the PriceOnProperty field if non-nil, zero value otherwise.

### GetPriceOnPropertyOk

`func (o *ListingInventoryWithAssociations) GetPriceOnPropertyOk() (*[]int64, bool)`

GetPriceOnPropertyOk returns a tuple with the PriceOnProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceOnProperty

`func (o *ListingInventoryWithAssociations) SetPriceOnProperty(v []int64)`

SetPriceOnProperty sets PriceOnProperty field to given value.

### HasPriceOnProperty

`func (o *ListingInventoryWithAssociations) HasPriceOnProperty() bool`

HasPriceOnProperty returns a boolean if a field has been set.

### GetQuantityOnProperty

`func (o *ListingInventoryWithAssociations) GetQuantityOnProperty() []int64`

GetQuantityOnProperty returns the QuantityOnProperty field if non-nil, zero value otherwise.

### GetQuantityOnPropertyOk

`func (o *ListingInventoryWithAssociations) GetQuantityOnPropertyOk() (*[]int64, bool)`

GetQuantityOnPropertyOk returns a tuple with the QuantityOnProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnProperty

`func (o *ListingInventoryWithAssociations) SetQuantityOnProperty(v []int64)`

SetQuantityOnProperty sets QuantityOnProperty field to given value.

### HasQuantityOnProperty

`func (o *ListingInventoryWithAssociations) HasQuantityOnProperty() bool`

HasQuantityOnProperty returns a boolean if a field has been set.

### GetSkuOnProperty

`func (o *ListingInventoryWithAssociations) GetSkuOnProperty() []int64`

GetSkuOnProperty returns the SkuOnProperty field if non-nil, zero value otherwise.

### GetSkuOnPropertyOk

`func (o *ListingInventoryWithAssociations) GetSkuOnPropertyOk() (*[]int64, bool)`

GetSkuOnPropertyOk returns a tuple with the SkuOnProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuOnProperty

`func (o *ListingInventoryWithAssociations) SetSkuOnProperty(v []int64)`

SetSkuOnProperty sets SkuOnProperty field to given value.

### HasSkuOnProperty

`func (o *ListingInventoryWithAssociations) HasSkuOnProperty() bool`

HasSkuOnProperty returns a boolean if a field has been set.

### GetListing

`func (o *ListingInventoryWithAssociations) GetListing() ShopListing`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *ListingInventoryWithAssociations) GetListingOk() (*ShopListing, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *ListingInventoryWithAssociations) SetListing(v ShopListing)`

SetListing sets Listing field to given value.

### HasListing

`func (o *ListingInventoryWithAssociations) HasListing() bool`

HasListing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


