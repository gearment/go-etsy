# CreateShopShippingProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The name string of this shipping profile. | 
**OriginCountryIso** | **string** | The ISO code of the country from which the listing ships. | 
**PrimaryCost** | **float32** | The cost of shipping to this country/region alone, measured in the store&#39;s default currency. | 
**SecondaryCost** | **float32** | The cost of shipping to this country/region with another item, measured in the store&#39;s default currency. | 
**MinProcessingTime** | **int64** | The minimum time required to process to ship listings with this shipping profile. | 
**MaxProcessingTime** | **int64** | The maximum processing time the listing needs to ship. | 
**ProcessingTimeUnit** | Pointer to [**CreateShopShippingProfileRequestProcessingTimeUnit**](CreateShopShippingProfileRequestProcessingTimeUnit.md) |  | [optional] [default to CREATESHOPSHIPPINGPROFILEREQUESTPROCESSINGTIMEUNIT_BUSINESS_DAYS]
**DestinationCountryIso** | Pointer to **string** | The ISO code of the country to which the listing ships. If null, request sets destination to destination_region. Required if destination_region is null or not provided. | [optional] 
**DestinationRegion** | Pointer to [**CreateShopShippingProfileRequestDestinationRegion**](CreateShopShippingProfileRequestDestinationRegion.md) |  | [optional] [default to CREATESHOPSHIPPINGPROFILEREQUESTDESTINATIONREGION_NONE]
**OriginPostalCode** | Pointer to **string** | The postal code string (not necessarily a number) for the location from which the listing ships. Required if the &#x60;origin_country_iso&#x60; supports postal codes. See the [Fulfillment Tutorial docs](https://developer.etsy.com/documentation/tutorials/fulfillment/#countries-requiring-postal-codes) for more info | [optional] [default to ""]
**ShippingCarrierId** | Pointer to **int64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] [default to 0]
**MailClass** | Pointer to **string** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] 
**MinDeliveryDays** | Pointer to **int64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 
**MaxDeliveryDays** | Pointer to **int64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 

## Methods

### NewCreateShopShippingProfileRequest

`func NewCreateShopShippingProfileRequest(title string, originCountryIso string, primaryCost float32, secondaryCost float32, minProcessingTime int64, maxProcessingTime int64, ) *CreateShopShippingProfileRequest`

NewCreateShopShippingProfileRequest instantiates a new CreateShopShippingProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShopShippingProfileRequestWithDefaults

`func NewCreateShopShippingProfileRequestWithDefaults() *CreateShopShippingProfileRequest`

NewCreateShopShippingProfileRequestWithDefaults instantiates a new CreateShopShippingProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateShopShippingProfileRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateShopShippingProfileRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateShopShippingProfileRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetOriginCountryIso

`func (o *CreateShopShippingProfileRequest) GetOriginCountryIso() string`

GetOriginCountryIso returns the OriginCountryIso field if non-nil, zero value otherwise.

### GetOriginCountryIsoOk

`func (o *CreateShopShippingProfileRequest) GetOriginCountryIsoOk() (*string, bool)`

GetOriginCountryIsoOk returns a tuple with the OriginCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountryIso

`func (o *CreateShopShippingProfileRequest) SetOriginCountryIso(v string)`

SetOriginCountryIso sets OriginCountryIso field to given value.


### GetPrimaryCost

`func (o *CreateShopShippingProfileRequest) GetPrimaryCost() float32`

GetPrimaryCost returns the PrimaryCost field if non-nil, zero value otherwise.

### GetPrimaryCostOk

`func (o *CreateShopShippingProfileRequest) GetPrimaryCostOk() (*float32, bool)`

GetPrimaryCostOk returns a tuple with the PrimaryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCost

`func (o *CreateShopShippingProfileRequest) SetPrimaryCost(v float32)`

SetPrimaryCost sets PrimaryCost field to given value.


### GetSecondaryCost

`func (o *CreateShopShippingProfileRequest) GetSecondaryCost() float32`

GetSecondaryCost returns the SecondaryCost field if non-nil, zero value otherwise.

### GetSecondaryCostOk

`func (o *CreateShopShippingProfileRequest) GetSecondaryCostOk() (*float32, bool)`

GetSecondaryCostOk returns a tuple with the SecondaryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCost

`func (o *CreateShopShippingProfileRequest) SetSecondaryCost(v float32)`

SetSecondaryCost sets SecondaryCost field to given value.


### GetMinProcessingTime

`func (o *CreateShopShippingProfileRequest) GetMinProcessingTime() int64`

GetMinProcessingTime returns the MinProcessingTime field if non-nil, zero value otherwise.

### GetMinProcessingTimeOk

`func (o *CreateShopShippingProfileRequest) GetMinProcessingTimeOk() (*int64, bool)`

GetMinProcessingTimeOk returns a tuple with the MinProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProcessingTime

`func (o *CreateShopShippingProfileRequest) SetMinProcessingTime(v int64)`

SetMinProcessingTime sets MinProcessingTime field to given value.


### GetMaxProcessingTime

`func (o *CreateShopShippingProfileRequest) GetMaxProcessingTime() int64`

GetMaxProcessingTime returns the MaxProcessingTime field if non-nil, zero value otherwise.

### GetMaxProcessingTimeOk

`func (o *CreateShopShippingProfileRequest) GetMaxProcessingTimeOk() (*int64, bool)`

GetMaxProcessingTimeOk returns a tuple with the MaxProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProcessingTime

`func (o *CreateShopShippingProfileRequest) SetMaxProcessingTime(v int64)`

SetMaxProcessingTime sets MaxProcessingTime field to given value.


### GetProcessingTimeUnit

`func (o *CreateShopShippingProfileRequest) GetProcessingTimeUnit() CreateShopShippingProfileRequestProcessingTimeUnit`

GetProcessingTimeUnit returns the ProcessingTimeUnit field if non-nil, zero value otherwise.

### GetProcessingTimeUnitOk

`func (o *CreateShopShippingProfileRequest) GetProcessingTimeUnitOk() (*CreateShopShippingProfileRequestProcessingTimeUnit, bool)`

GetProcessingTimeUnitOk returns a tuple with the ProcessingTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeUnit

`func (o *CreateShopShippingProfileRequest) SetProcessingTimeUnit(v CreateShopShippingProfileRequestProcessingTimeUnit)`

SetProcessingTimeUnit sets ProcessingTimeUnit field to given value.

### HasProcessingTimeUnit

`func (o *CreateShopShippingProfileRequest) HasProcessingTimeUnit() bool`

HasProcessingTimeUnit returns a boolean if a field has been set.

### GetDestinationCountryIso

`func (o *CreateShopShippingProfileRequest) GetDestinationCountryIso() string`

GetDestinationCountryIso returns the DestinationCountryIso field if non-nil, zero value otherwise.

### GetDestinationCountryIsoOk

`func (o *CreateShopShippingProfileRequest) GetDestinationCountryIsoOk() (*string, bool)`

GetDestinationCountryIsoOk returns a tuple with the DestinationCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCountryIso

`func (o *CreateShopShippingProfileRequest) SetDestinationCountryIso(v string)`

SetDestinationCountryIso sets DestinationCountryIso field to given value.

### HasDestinationCountryIso

`func (o *CreateShopShippingProfileRequest) HasDestinationCountryIso() bool`

HasDestinationCountryIso returns a boolean if a field has been set.

### GetDestinationRegion

`func (o *CreateShopShippingProfileRequest) GetDestinationRegion() CreateShopShippingProfileRequestDestinationRegion`

GetDestinationRegion returns the DestinationRegion field if non-nil, zero value otherwise.

### GetDestinationRegionOk

`func (o *CreateShopShippingProfileRequest) GetDestinationRegionOk() (*CreateShopShippingProfileRequestDestinationRegion, bool)`

GetDestinationRegionOk returns a tuple with the DestinationRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRegion

`func (o *CreateShopShippingProfileRequest) SetDestinationRegion(v CreateShopShippingProfileRequestDestinationRegion)`

SetDestinationRegion sets DestinationRegion field to given value.

### HasDestinationRegion

`func (o *CreateShopShippingProfileRequest) HasDestinationRegion() bool`

HasDestinationRegion returns a boolean if a field has been set.

### GetOriginPostalCode

`func (o *CreateShopShippingProfileRequest) GetOriginPostalCode() string`

GetOriginPostalCode returns the OriginPostalCode field if non-nil, zero value otherwise.

### GetOriginPostalCodeOk

`func (o *CreateShopShippingProfileRequest) GetOriginPostalCodeOk() (*string, bool)`

GetOriginPostalCodeOk returns a tuple with the OriginPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPostalCode

`func (o *CreateShopShippingProfileRequest) SetOriginPostalCode(v string)`

SetOriginPostalCode sets OriginPostalCode field to given value.

### HasOriginPostalCode

`func (o *CreateShopShippingProfileRequest) HasOriginPostalCode() bool`

HasOriginPostalCode returns a boolean if a field has been set.

### GetShippingCarrierId

`func (o *CreateShopShippingProfileRequest) GetShippingCarrierId() int64`

GetShippingCarrierId returns the ShippingCarrierId field if non-nil, zero value otherwise.

### GetShippingCarrierIdOk

`func (o *CreateShopShippingProfileRequest) GetShippingCarrierIdOk() (*int64, bool)`

GetShippingCarrierIdOk returns a tuple with the ShippingCarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCarrierId

`func (o *CreateShopShippingProfileRequest) SetShippingCarrierId(v int64)`

SetShippingCarrierId sets ShippingCarrierId field to given value.

### HasShippingCarrierId

`func (o *CreateShopShippingProfileRequest) HasShippingCarrierId() bool`

HasShippingCarrierId returns a boolean if a field has been set.

### GetMailClass

`func (o *CreateShopShippingProfileRequest) GetMailClass() string`

GetMailClass returns the MailClass field if non-nil, zero value otherwise.

### GetMailClassOk

`func (o *CreateShopShippingProfileRequest) GetMailClassOk() (*string, bool)`

GetMailClassOk returns a tuple with the MailClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClass

`func (o *CreateShopShippingProfileRequest) SetMailClass(v string)`

SetMailClass sets MailClass field to given value.

### HasMailClass

`func (o *CreateShopShippingProfileRequest) HasMailClass() bool`

HasMailClass returns a boolean if a field has been set.

### GetMinDeliveryDays

`func (o *CreateShopShippingProfileRequest) GetMinDeliveryDays() int64`

GetMinDeliveryDays returns the MinDeliveryDays field if non-nil, zero value otherwise.

### GetMinDeliveryDaysOk

`func (o *CreateShopShippingProfileRequest) GetMinDeliveryDaysOk() (*int64, bool)`

GetMinDeliveryDaysOk returns a tuple with the MinDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDeliveryDays

`func (o *CreateShopShippingProfileRequest) SetMinDeliveryDays(v int64)`

SetMinDeliveryDays sets MinDeliveryDays field to given value.

### HasMinDeliveryDays

`func (o *CreateShopShippingProfileRequest) HasMinDeliveryDays() bool`

HasMinDeliveryDays returns a boolean if a field has been set.

### GetMaxDeliveryDays

`func (o *CreateShopShippingProfileRequest) GetMaxDeliveryDays() int64`

GetMaxDeliveryDays returns the MaxDeliveryDays field if non-nil, zero value otherwise.

### GetMaxDeliveryDaysOk

`func (o *CreateShopShippingProfileRequest) GetMaxDeliveryDaysOk() (*int64, bool)`

GetMaxDeliveryDaysOk returns a tuple with the MaxDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveryDays

`func (o *CreateShopShippingProfileRequest) SetMaxDeliveryDays(v int64)`

SetMaxDeliveryDays sets MaxDeliveryDays field to given value.

### HasMaxDeliveryDays

`func (o *CreateShopShippingProfileRequest) HasMaxDeliveryDays() bool`

HasMaxDeliveryDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


