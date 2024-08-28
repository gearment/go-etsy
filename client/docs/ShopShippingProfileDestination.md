# ShopShippingProfileDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingProfileDestinationId** | Pointer to **int64** | The numeric ID of the shipping profile destination in the [shipping profile](/documentation/reference#tag/Shop-ShippingProfile) associated with the listing. | [optional] 
**ShippingProfileId** | Pointer to **int64** | The numeric ID of the shipping profile. | [optional] 
**OriginCountryIso** | Pointer to **string** | The ISO code of the country from which the listing ships. | [optional] 
**DestinationCountryIso** | Pointer to **string** | The ISO code of the country to which the listing ships. If null, request sets destination to destination_region. Required if destination_region is null or not provided. | [optional] 
**DestinationRegion** | Pointer to [**ShopShippingProfileDestinationDestinationRegion**](ShopShippingProfileDestinationDestinationRegion.md) |  | [optional] 
**PrimaryCost** | Pointer to [**Money**](Money.md) | The cost of shipping to this country/region alone, measured in the store&#39;s default currency. | [optional] 
**SecondaryCost** | Pointer to [**Money**](Money.md) | The cost of shipping to this country/region with another item, measured in the store&#39;s default currency. | [optional] 
**ShippingCarrierId** | Pointer to **NullableInt64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] 
**MailClass** | Pointer to **NullableString** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] 
**MinDeliveryDays** | Pointer to **NullableInt64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 
**MaxDeliveryDays** | Pointer to **NullableInt64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 

## Methods

### NewShopShippingProfileDestination

`func NewShopShippingProfileDestination() *ShopShippingProfileDestination`

NewShopShippingProfileDestination instantiates a new ShopShippingProfileDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopShippingProfileDestinationWithDefaults

`func NewShopShippingProfileDestinationWithDefaults() *ShopShippingProfileDestination`

NewShopShippingProfileDestinationWithDefaults instantiates a new ShopShippingProfileDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingProfileDestinationId

`func (o *ShopShippingProfileDestination) GetShippingProfileDestinationId() int64`

GetShippingProfileDestinationId returns the ShippingProfileDestinationId field if non-nil, zero value otherwise.

### GetShippingProfileDestinationIdOk

`func (o *ShopShippingProfileDestination) GetShippingProfileDestinationIdOk() (*int64, bool)`

GetShippingProfileDestinationIdOk returns a tuple with the ShippingProfileDestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileDestinationId

`func (o *ShopShippingProfileDestination) SetShippingProfileDestinationId(v int64)`

SetShippingProfileDestinationId sets ShippingProfileDestinationId field to given value.

### HasShippingProfileDestinationId

`func (o *ShopShippingProfileDestination) HasShippingProfileDestinationId() bool`

HasShippingProfileDestinationId returns a boolean if a field has been set.

### GetShippingProfileId

`func (o *ShopShippingProfileDestination) GetShippingProfileId() int64`

GetShippingProfileId returns the ShippingProfileId field if non-nil, zero value otherwise.

### GetShippingProfileIdOk

`func (o *ShopShippingProfileDestination) GetShippingProfileIdOk() (*int64, bool)`

GetShippingProfileIdOk returns a tuple with the ShippingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileId

`func (o *ShopShippingProfileDestination) SetShippingProfileId(v int64)`

SetShippingProfileId sets ShippingProfileId field to given value.

### HasShippingProfileId

`func (o *ShopShippingProfileDestination) HasShippingProfileId() bool`

HasShippingProfileId returns a boolean if a field has been set.

### GetOriginCountryIso

`func (o *ShopShippingProfileDestination) GetOriginCountryIso() string`

GetOriginCountryIso returns the OriginCountryIso field if non-nil, zero value otherwise.

### GetOriginCountryIsoOk

`func (o *ShopShippingProfileDestination) GetOriginCountryIsoOk() (*string, bool)`

GetOriginCountryIsoOk returns a tuple with the OriginCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountryIso

`func (o *ShopShippingProfileDestination) SetOriginCountryIso(v string)`

SetOriginCountryIso sets OriginCountryIso field to given value.

### HasOriginCountryIso

`func (o *ShopShippingProfileDestination) HasOriginCountryIso() bool`

HasOriginCountryIso returns a boolean if a field has been set.

### GetDestinationCountryIso

`func (o *ShopShippingProfileDestination) GetDestinationCountryIso() string`

GetDestinationCountryIso returns the DestinationCountryIso field if non-nil, zero value otherwise.

### GetDestinationCountryIsoOk

`func (o *ShopShippingProfileDestination) GetDestinationCountryIsoOk() (*string, bool)`

GetDestinationCountryIsoOk returns a tuple with the DestinationCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCountryIso

`func (o *ShopShippingProfileDestination) SetDestinationCountryIso(v string)`

SetDestinationCountryIso sets DestinationCountryIso field to given value.

### HasDestinationCountryIso

`func (o *ShopShippingProfileDestination) HasDestinationCountryIso() bool`

HasDestinationCountryIso returns a boolean if a field has been set.

### GetDestinationRegion

`func (o *ShopShippingProfileDestination) GetDestinationRegion() ShopShippingProfileDestinationDestinationRegion`

GetDestinationRegion returns the DestinationRegion field if non-nil, zero value otherwise.

### GetDestinationRegionOk

`func (o *ShopShippingProfileDestination) GetDestinationRegionOk() (*ShopShippingProfileDestinationDestinationRegion, bool)`

GetDestinationRegionOk returns a tuple with the DestinationRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRegion

`func (o *ShopShippingProfileDestination) SetDestinationRegion(v ShopShippingProfileDestinationDestinationRegion)`

SetDestinationRegion sets DestinationRegion field to given value.

### HasDestinationRegion

`func (o *ShopShippingProfileDestination) HasDestinationRegion() bool`

HasDestinationRegion returns a boolean if a field has been set.

### GetPrimaryCost

`func (o *ShopShippingProfileDestination) GetPrimaryCost() Money`

GetPrimaryCost returns the PrimaryCost field if non-nil, zero value otherwise.

### GetPrimaryCostOk

`func (o *ShopShippingProfileDestination) GetPrimaryCostOk() (*Money, bool)`

GetPrimaryCostOk returns a tuple with the PrimaryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCost

`func (o *ShopShippingProfileDestination) SetPrimaryCost(v Money)`

SetPrimaryCost sets PrimaryCost field to given value.

### HasPrimaryCost

`func (o *ShopShippingProfileDestination) HasPrimaryCost() bool`

HasPrimaryCost returns a boolean if a field has been set.

### GetSecondaryCost

`func (o *ShopShippingProfileDestination) GetSecondaryCost() Money`

GetSecondaryCost returns the SecondaryCost field if non-nil, zero value otherwise.

### GetSecondaryCostOk

`func (o *ShopShippingProfileDestination) GetSecondaryCostOk() (*Money, bool)`

GetSecondaryCostOk returns a tuple with the SecondaryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCost

`func (o *ShopShippingProfileDestination) SetSecondaryCost(v Money)`

SetSecondaryCost sets SecondaryCost field to given value.

### HasSecondaryCost

`func (o *ShopShippingProfileDestination) HasSecondaryCost() bool`

HasSecondaryCost returns a boolean if a field has been set.

### GetShippingCarrierId

`func (o *ShopShippingProfileDestination) GetShippingCarrierId() int64`

GetShippingCarrierId returns the ShippingCarrierId field if non-nil, zero value otherwise.

### GetShippingCarrierIdOk

`func (o *ShopShippingProfileDestination) GetShippingCarrierIdOk() (*int64, bool)`

GetShippingCarrierIdOk returns a tuple with the ShippingCarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCarrierId

`func (o *ShopShippingProfileDestination) SetShippingCarrierId(v int64)`

SetShippingCarrierId sets ShippingCarrierId field to given value.

### HasShippingCarrierId

`func (o *ShopShippingProfileDestination) HasShippingCarrierId() bool`

HasShippingCarrierId returns a boolean if a field has been set.

### SetShippingCarrierIdNil

`func (o *ShopShippingProfileDestination) SetShippingCarrierIdNil(b bool)`

 SetShippingCarrierIdNil sets the value for ShippingCarrierId to be an explicit nil

### UnsetShippingCarrierId
`func (o *ShopShippingProfileDestination) UnsetShippingCarrierId()`

UnsetShippingCarrierId ensures that no value is present for ShippingCarrierId, not even an explicit nil
### GetMailClass

`func (o *ShopShippingProfileDestination) GetMailClass() string`

GetMailClass returns the MailClass field if non-nil, zero value otherwise.

### GetMailClassOk

`func (o *ShopShippingProfileDestination) GetMailClassOk() (*string, bool)`

GetMailClassOk returns a tuple with the MailClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClass

`func (o *ShopShippingProfileDestination) SetMailClass(v string)`

SetMailClass sets MailClass field to given value.

### HasMailClass

`func (o *ShopShippingProfileDestination) HasMailClass() bool`

HasMailClass returns a boolean if a field has been set.

### SetMailClassNil

`func (o *ShopShippingProfileDestination) SetMailClassNil(b bool)`

 SetMailClassNil sets the value for MailClass to be an explicit nil

### UnsetMailClass
`func (o *ShopShippingProfileDestination) UnsetMailClass()`

UnsetMailClass ensures that no value is present for MailClass, not even an explicit nil
### GetMinDeliveryDays

`func (o *ShopShippingProfileDestination) GetMinDeliveryDays() int64`

GetMinDeliveryDays returns the MinDeliveryDays field if non-nil, zero value otherwise.

### GetMinDeliveryDaysOk

`func (o *ShopShippingProfileDestination) GetMinDeliveryDaysOk() (*int64, bool)`

GetMinDeliveryDaysOk returns a tuple with the MinDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDeliveryDays

`func (o *ShopShippingProfileDestination) SetMinDeliveryDays(v int64)`

SetMinDeliveryDays sets MinDeliveryDays field to given value.

### HasMinDeliveryDays

`func (o *ShopShippingProfileDestination) HasMinDeliveryDays() bool`

HasMinDeliveryDays returns a boolean if a field has been set.

### SetMinDeliveryDaysNil

`func (o *ShopShippingProfileDestination) SetMinDeliveryDaysNil(b bool)`

 SetMinDeliveryDaysNil sets the value for MinDeliveryDays to be an explicit nil

### UnsetMinDeliveryDays
`func (o *ShopShippingProfileDestination) UnsetMinDeliveryDays()`

UnsetMinDeliveryDays ensures that no value is present for MinDeliveryDays, not even an explicit nil
### GetMaxDeliveryDays

`func (o *ShopShippingProfileDestination) GetMaxDeliveryDays() int64`

GetMaxDeliveryDays returns the MaxDeliveryDays field if non-nil, zero value otherwise.

### GetMaxDeliveryDaysOk

`func (o *ShopShippingProfileDestination) GetMaxDeliveryDaysOk() (*int64, bool)`

GetMaxDeliveryDaysOk returns a tuple with the MaxDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveryDays

`func (o *ShopShippingProfileDestination) SetMaxDeliveryDays(v int64)`

SetMaxDeliveryDays sets MaxDeliveryDays field to given value.

### HasMaxDeliveryDays

`func (o *ShopShippingProfileDestination) HasMaxDeliveryDays() bool`

HasMaxDeliveryDays returns a boolean if a field has been set.

### SetMaxDeliveryDaysNil

`func (o *ShopShippingProfileDestination) SetMaxDeliveryDaysNil(b bool)`

 SetMaxDeliveryDaysNil sets the value for MaxDeliveryDays to be an explicit nil

### UnsetMaxDeliveryDays
`func (o *ShopShippingProfileDestination) UnsetMaxDeliveryDays()`

UnsetMaxDeliveryDays ensures that no value is present for MaxDeliveryDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


