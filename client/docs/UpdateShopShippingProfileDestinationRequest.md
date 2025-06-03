# UpdateShopShippingProfileDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryCost** | Pointer to **float32** | The cost of shipping to this country/region alone, measured in the store&#39;s default currency. | [optional] 
**SecondaryCost** | Pointer to **float32** | The cost of shipping to this country/region with another item, measured in the store&#39;s default currency. | [optional] 
**DestinationCountryIso** | Pointer to **string** | The ISO code of the country to which the listing ships. If null, request sets destination to destination_region. Required if destination_region is null or not provided. | [optional] 
**DestinationRegion** | Pointer to [**CreateShopShippingProfileRequestDestinationRegion**](CreateShopShippingProfileRequestDestinationRegion.md) |  | [optional] [default to CREATESHOPSHIPPINGPROFILEREQUESTDESTINATIONREGION_NONE]
**ShippingCarrierId** | Pointer to **int64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] 
**MailClass** | Pointer to **string** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] 
**MinDeliveryDays** | Pointer to **int64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 
**MaxDeliveryDays** | Pointer to **int64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 

## Methods

### NewUpdateShopShippingProfileDestinationRequest

`func NewUpdateShopShippingProfileDestinationRequest() *UpdateShopShippingProfileDestinationRequest`

NewUpdateShopShippingProfileDestinationRequest instantiates a new UpdateShopShippingProfileDestinationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShopShippingProfileDestinationRequestWithDefaults

`func NewUpdateShopShippingProfileDestinationRequestWithDefaults() *UpdateShopShippingProfileDestinationRequest`

NewUpdateShopShippingProfileDestinationRequestWithDefaults instantiates a new UpdateShopShippingProfileDestinationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryCost

`func (o *UpdateShopShippingProfileDestinationRequest) GetPrimaryCost() float32`

GetPrimaryCost returns the PrimaryCost field if non-nil, zero value otherwise.

### GetPrimaryCostOk

`func (o *UpdateShopShippingProfileDestinationRequest) GetPrimaryCostOk() (*float32, bool)`

GetPrimaryCostOk returns a tuple with the PrimaryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCost

`func (o *UpdateShopShippingProfileDestinationRequest) SetPrimaryCost(v float32)`

SetPrimaryCost sets PrimaryCost field to given value.

### HasPrimaryCost

`func (o *UpdateShopShippingProfileDestinationRequest) HasPrimaryCost() bool`

HasPrimaryCost returns a boolean if a field has been set.

### GetSecondaryCost

`func (o *UpdateShopShippingProfileDestinationRequest) GetSecondaryCost() float32`

GetSecondaryCost returns the SecondaryCost field if non-nil, zero value otherwise.

### GetSecondaryCostOk

`func (o *UpdateShopShippingProfileDestinationRequest) GetSecondaryCostOk() (*float32, bool)`

GetSecondaryCostOk returns a tuple with the SecondaryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCost

`func (o *UpdateShopShippingProfileDestinationRequest) SetSecondaryCost(v float32)`

SetSecondaryCost sets SecondaryCost field to given value.

### HasSecondaryCost

`func (o *UpdateShopShippingProfileDestinationRequest) HasSecondaryCost() bool`

HasSecondaryCost returns a boolean if a field has been set.

### GetDestinationCountryIso

`func (o *UpdateShopShippingProfileDestinationRequest) GetDestinationCountryIso() string`

GetDestinationCountryIso returns the DestinationCountryIso field if non-nil, zero value otherwise.

### GetDestinationCountryIsoOk

`func (o *UpdateShopShippingProfileDestinationRequest) GetDestinationCountryIsoOk() (*string, bool)`

GetDestinationCountryIsoOk returns a tuple with the DestinationCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCountryIso

`func (o *UpdateShopShippingProfileDestinationRequest) SetDestinationCountryIso(v string)`

SetDestinationCountryIso sets DestinationCountryIso field to given value.

### HasDestinationCountryIso

`func (o *UpdateShopShippingProfileDestinationRequest) HasDestinationCountryIso() bool`

HasDestinationCountryIso returns a boolean if a field has been set.

### GetDestinationRegion

`func (o *UpdateShopShippingProfileDestinationRequest) GetDestinationRegion() CreateShopShippingProfileRequestDestinationRegion`

GetDestinationRegion returns the DestinationRegion field if non-nil, zero value otherwise.

### GetDestinationRegionOk

`func (o *UpdateShopShippingProfileDestinationRequest) GetDestinationRegionOk() (*CreateShopShippingProfileRequestDestinationRegion, bool)`

GetDestinationRegionOk returns a tuple with the DestinationRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRegion

`func (o *UpdateShopShippingProfileDestinationRequest) SetDestinationRegion(v CreateShopShippingProfileRequestDestinationRegion)`

SetDestinationRegion sets DestinationRegion field to given value.

### HasDestinationRegion

`func (o *UpdateShopShippingProfileDestinationRequest) HasDestinationRegion() bool`

HasDestinationRegion returns a boolean if a field has been set.

### GetShippingCarrierId

`func (o *UpdateShopShippingProfileDestinationRequest) GetShippingCarrierId() int64`

GetShippingCarrierId returns the ShippingCarrierId field if non-nil, zero value otherwise.

### GetShippingCarrierIdOk

`func (o *UpdateShopShippingProfileDestinationRequest) GetShippingCarrierIdOk() (*int64, bool)`

GetShippingCarrierIdOk returns a tuple with the ShippingCarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCarrierId

`func (o *UpdateShopShippingProfileDestinationRequest) SetShippingCarrierId(v int64)`

SetShippingCarrierId sets ShippingCarrierId field to given value.

### HasShippingCarrierId

`func (o *UpdateShopShippingProfileDestinationRequest) HasShippingCarrierId() bool`

HasShippingCarrierId returns a boolean if a field has been set.

### GetMailClass

`func (o *UpdateShopShippingProfileDestinationRequest) GetMailClass() string`

GetMailClass returns the MailClass field if non-nil, zero value otherwise.

### GetMailClassOk

`func (o *UpdateShopShippingProfileDestinationRequest) GetMailClassOk() (*string, bool)`

GetMailClassOk returns a tuple with the MailClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClass

`func (o *UpdateShopShippingProfileDestinationRequest) SetMailClass(v string)`

SetMailClass sets MailClass field to given value.

### HasMailClass

`func (o *UpdateShopShippingProfileDestinationRequest) HasMailClass() bool`

HasMailClass returns a boolean if a field has been set.

### GetMinDeliveryDays

`func (o *UpdateShopShippingProfileDestinationRequest) GetMinDeliveryDays() int64`

GetMinDeliveryDays returns the MinDeliveryDays field if non-nil, zero value otherwise.

### GetMinDeliveryDaysOk

`func (o *UpdateShopShippingProfileDestinationRequest) GetMinDeliveryDaysOk() (*int64, bool)`

GetMinDeliveryDaysOk returns a tuple with the MinDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDeliveryDays

`func (o *UpdateShopShippingProfileDestinationRequest) SetMinDeliveryDays(v int64)`

SetMinDeliveryDays sets MinDeliveryDays field to given value.

### HasMinDeliveryDays

`func (o *UpdateShopShippingProfileDestinationRequest) HasMinDeliveryDays() bool`

HasMinDeliveryDays returns a boolean if a field has been set.

### GetMaxDeliveryDays

`func (o *UpdateShopShippingProfileDestinationRequest) GetMaxDeliveryDays() int64`

GetMaxDeliveryDays returns the MaxDeliveryDays field if non-nil, zero value otherwise.

### GetMaxDeliveryDaysOk

`func (o *UpdateShopShippingProfileDestinationRequest) GetMaxDeliveryDaysOk() (*int64, bool)`

GetMaxDeliveryDaysOk returns a tuple with the MaxDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveryDays

`func (o *UpdateShopShippingProfileDestinationRequest) SetMaxDeliveryDays(v int64)`

SetMaxDeliveryDays sets MaxDeliveryDays field to given value.

### HasMaxDeliveryDays

`func (o *UpdateShopShippingProfileDestinationRequest) HasMaxDeliveryDays() bool`

HasMaxDeliveryDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


