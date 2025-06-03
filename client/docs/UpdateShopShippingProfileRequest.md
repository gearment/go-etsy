# UpdateShopShippingProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The name string of this shipping profile. | [optional] 
**OriginCountryIso** | Pointer to **string** | The ISO code of the country from which the listing ships. | [optional] 
**MinProcessingTime** | Pointer to **int64** | The minimum time required to process to ship listings with this shipping profile. | [optional] 
**MaxProcessingTime** | Pointer to **int64** | The maximum processing time the listing needs to ship. | [optional] 
**ProcessingTimeUnit** | Pointer to [**CreateShopShippingProfileRequestProcessingTimeUnit**](CreateShopShippingProfileRequestProcessingTimeUnit.md) |  | [optional] [default to CREATESHOPSHIPPINGPROFILEREQUESTPROCESSINGTIMEUNIT_BUSINESS_DAYS]
**OriginPostalCode** | Pointer to **string** | The postal code string (not necessarily a number) for the location from which the listing ships. Required if the &#x60;origin_country_iso&#x60; supports postal codes. See the [Fulfillment Tutorial docs](https://developer.etsy.com/documentation/tutorials/fulfillment/#countries-requiring-postal-codes) for more info | [optional] 

## Methods

### NewUpdateShopShippingProfileRequest

`func NewUpdateShopShippingProfileRequest() *UpdateShopShippingProfileRequest`

NewUpdateShopShippingProfileRequest instantiates a new UpdateShopShippingProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShopShippingProfileRequestWithDefaults

`func NewUpdateShopShippingProfileRequestWithDefaults() *UpdateShopShippingProfileRequest`

NewUpdateShopShippingProfileRequestWithDefaults instantiates a new UpdateShopShippingProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateShopShippingProfileRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateShopShippingProfileRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateShopShippingProfileRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateShopShippingProfileRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOriginCountryIso

`func (o *UpdateShopShippingProfileRequest) GetOriginCountryIso() string`

GetOriginCountryIso returns the OriginCountryIso field if non-nil, zero value otherwise.

### GetOriginCountryIsoOk

`func (o *UpdateShopShippingProfileRequest) GetOriginCountryIsoOk() (*string, bool)`

GetOriginCountryIsoOk returns a tuple with the OriginCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountryIso

`func (o *UpdateShopShippingProfileRequest) SetOriginCountryIso(v string)`

SetOriginCountryIso sets OriginCountryIso field to given value.

### HasOriginCountryIso

`func (o *UpdateShopShippingProfileRequest) HasOriginCountryIso() bool`

HasOriginCountryIso returns a boolean if a field has been set.

### GetMinProcessingTime

`func (o *UpdateShopShippingProfileRequest) GetMinProcessingTime() int64`

GetMinProcessingTime returns the MinProcessingTime field if non-nil, zero value otherwise.

### GetMinProcessingTimeOk

`func (o *UpdateShopShippingProfileRequest) GetMinProcessingTimeOk() (*int64, bool)`

GetMinProcessingTimeOk returns a tuple with the MinProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProcessingTime

`func (o *UpdateShopShippingProfileRequest) SetMinProcessingTime(v int64)`

SetMinProcessingTime sets MinProcessingTime field to given value.

### HasMinProcessingTime

`func (o *UpdateShopShippingProfileRequest) HasMinProcessingTime() bool`

HasMinProcessingTime returns a boolean if a field has been set.

### GetMaxProcessingTime

`func (o *UpdateShopShippingProfileRequest) GetMaxProcessingTime() int64`

GetMaxProcessingTime returns the MaxProcessingTime field if non-nil, zero value otherwise.

### GetMaxProcessingTimeOk

`func (o *UpdateShopShippingProfileRequest) GetMaxProcessingTimeOk() (*int64, bool)`

GetMaxProcessingTimeOk returns a tuple with the MaxProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProcessingTime

`func (o *UpdateShopShippingProfileRequest) SetMaxProcessingTime(v int64)`

SetMaxProcessingTime sets MaxProcessingTime field to given value.

### HasMaxProcessingTime

`func (o *UpdateShopShippingProfileRequest) HasMaxProcessingTime() bool`

HasMaxProcessingTime returns a boolean if a field has been set.

### GetProcessingTimeUnit

`func (o *UpdateShopShippingProfileRequest) GetProcessingTimeUnit() CreateShopShippingProfileRequestProcessingTimeUnit`

GetProcessingTimeUnit returns the ProcessingTimeUnit field if non-nil, zero value otherwise.

### GetProcessingTimeUnitOk

`func (o *UpdateShopShippingProfileRequest) GetProcessingTimeUnitOk() (*CreateShopShippingProfileRequestProcessingTimeUnit, bool)`

GetProcessingTimeUnitOk returns a tuple with the ProcessingTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeUnit

`func (o *UpdateShopShippingProfileRequest) SetProcessingTimeUnit(v CreateShopShippingProfileRequestProcessingTimeUnit)`

SetProcessingTimeUnit sets ProcessingTimeUnit field to given value.

### HasProcessingTimeUnit

`func (o *UpdateShopShippingProfileRequest) HasProcessingTimeUnit() bool`

HasProcessingTimeUnit returns a boolean if a field has been set.

### GetOriginPostalCode

`func (o *UpdateShopShippingProfileRequest) GetOriginPostalCode() string`

GetOriginPostalCode returns the OriginPostalCode field if non-nil, zero value otherwise.

### GetOriginPostalCodeOk

`func (o *UpdateShopShippingProfileRequest) GetOriginPostalCodeOk() (*string, bool)`

GetOriginPostalCodeOk returns a tuple with the OriginPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPostalCode

`func (o *UpdateShopShippingProfileRequest) SetOriginPostalCode(v string)`

SetOriginPostalCode sets OriginPostalCode field to given value.

### HasOriginPostalCode

`func (o *UpdateShopShippingProfileRequest) HasOriginPostalCode() bool`

HasOriginPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


