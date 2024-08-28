# ListingInventoryProductOffering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferingId** | Pointer to **int64** | The ID for the ProductOffering | [optional] 
**Quantity** | Pointer to **int64** | The quantity the ProductOffering | [optional] 
**IsEnabled** | Pointer to **bool** | Whether or not the offering can be shown to buyers. | [optional] 
**IsDeleted** | Pointer to **bool** | Whether or not the offering has been deleted. | [optional] 
**Price** | Pointer to [**Money**](Money.md) | Price data for this ProductOffering | [optional] 

## Methods

### NewListingInventoryProductOffering

`func NewListingInventoryProductOffering() *ListingInventoryProductOffering`

NewListingInventoryProductOffering instantiates a new ListingInventoryProductOffering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingInventoryProductOfferingWithDefaults

`func NewListingInventoryProductOfferingWithDefaults() *ListingInventoryProductOffering`

NewListingInventoryProductOfferingWithDefaults instantiates a new ListingInventoryProductOffering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferingId

`func (o *ListingInventoryProductOffering) GetOfferingId() int64`

GetOfferingId returns the OfferingId field if non-nil, zero value otherwise.

### GetOfferingIdOk

`func (o *ListingInventoryProductOffering) GetOfferingIdOk() (*int64, bool)`

GetOfferingIdOk returns a tuple with the OfferingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingId

`func (o *ListingInventoryProductOffering) SetOfferingId(v int64)`

SetOfferingId sets OfferingId field to given value.

### HasOfferingId

`func (o *ListingInventoryProductOffering) HasOfferingId() bool`

HasOfferingId returns a boolean if a field has been set.

### GetQuantity

`func (o *ListingInventoryProductOffering) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ListingInventoryProductOffering) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ListingInventoryProductOffering) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ListingInventoryProductOffering) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ListingInventoryProductOffering) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ListingInventoryProductOffering) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ListingInventoryProductOffering) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ListingInventoryProductOffering) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ListingInventoryProductOffering) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ListingInventoryProductOffering) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ListingInventoryProductOffering) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ListingInventoryProductOffering) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetPrice

`func (o *ListingInventoryProductOffering) GetPrice() Money`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListingInventoryProductOffering) GetPriceOk() (*Money, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListingInventoryProductOffering) SetPrice(v Money)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListingInventoryProductOffering) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


