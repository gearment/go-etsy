# UpdateListingInventoryRequestProductsInnerOfferingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **float32** | The price of the product. | 
**Quantity** | **int64** | How many of this product are available? | 
**IsEnabled** | **bool** | True if the offering is shown to buyers | 
**ReadinessStateId** | **NullableInt64** | The numeric ID of the [processing profile](/documentation/reference#operation/getShopReadinessStateDefinition) associated with the listing. Returned only when the listing is &#x60;active&#x60; and of type &#x60;physical&#x60;, and the endpoint is either shop-scoped (path contains &#x60;shop_id&#x60;) or a single-listing request such as &#x60;getListing&#x60;. For every other case this field can be null. | 

## Methods

### NewUpdateListingInventoryRequestProductsInnerOfferingsInner

`func NewUpdateListingInventoryRequestProductsInnerOfferingsInner(price float32, quantity int64, isEnabled bool, readinessStateId NullableInt64, ) *UpdateListingInventoryRequestProductsInnerOfferingsInner`

NewUpdateListingInventoryRequestProductsInnerOfferingsInner instantiates a new UpdateListingInventoryRequestProductsInnerOfferingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListingInventoryRequestProductsInnerOfferingsInnerWithDefaults

`func NewUpdateListingInventoryRequestProductsInnerOfferingsInnerWithDefaults() *UpdateListingInventoryRequestProductsInnerOfferingsInner`

NewUpdateListingInventoryRequestProductsInnerOfferingsInnerWithDefaults instantiates a new UpdateListingInventoryRequestProductsInnerOfferingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetQuantity

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetIsEnabled

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetReadinessStateId

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetReadinessStateId() int64`

GetReadinessStateId returns the ReadinessStateId field if non-nil, zero value otherwise.

### GetReadinessStateIdOk

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetReadinessStateIdOk() (*int64, bool)`

GetReadinessStateIdOk returns a tuple with the ReadinessStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessStateId

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) SetReadinessStateId(v int64)`

SetReadinessStateId sets ReadinessStateId field to given value.


### SetReadinessStateIdNil

`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) SetReadinessStateIdNil(b bool)`

 SetReadinessStateIdNil sets the value for ReadinessStateId to be an explicit nil

### UnsetReadinessStateId
`func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) UnsetReadinessStateId()`

UnsetReadinessStateId ensures that no value is present for ReadinessStateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


