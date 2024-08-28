# ShopShippingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingProfileId** | Pointer to **int64** | The numeric ID of the shipping profile. | [optional] 
**Title** | Pointer to **NullableString** | The name string of this shipping profile. | [optional] 
**UserId** | Pointer to **int64** | The numeric ID for the [user](/documentation/reference#tag/User) who owns the shipping profile. | [optional] 
**MinProcessingDays** | Pointer to **NullableInt64** | The minimum number of days for processing the listing. | [optional] 
**MaxProcessingDays** | Pointer to **NullableInt64** | The maximum number of days for processing the listing. | [optional] 
**ProcessingDaysDisplayLabel** | Pointer to **string** | Translated display label string for processing days. | [optional] 
**OriginCountryIso** | Pointer to **string** | The ISO code of the country from which the listing ships. | [optional] 
**IsDeleted** | Pointer to **bool** | When true, someone deleted this shipping profile. | [optional] 
**ShippingProfileDestinations** | Pointer to [**[]ShopShippingProfileDestination**](ShopShippingProfileDestination.md) | A list of [shipping profile destinations](/documentation/reference/#operation/createShopShippingProfileDestination) available for this shipping profile. | [optional] 
**ShippingProfileUpgrades** | Pointer to [**[]ShopShippingProfileUpgrade**](ShopShippingProfileUpgrade.md) | A list of [shipping profile upgrades](/documentation/reference/#operation/createShopShippingProfileUpgrade) available for this shipping profile. | [optional] 
**OriginPostalCode** | Pointer to **NullableString** | The postal code string (not necessarily a number) for the location from which the listing ships. Required if the &#x60;origin_country_iso&#x60; supports postal codes. See the [Fulfillment Tutorial docs](https://developer.etsy.com/documentation/tutorials/fulfillment/#countries-requiring-postal-codes) for more info | [optional] 
**ProfileType** | Pointer to [**ShopShippingProfileProfileType**](ShopShippingProfileProfileType.md) |  | [optional] [default to SHOPSHIPPINGPROFILEPROFILETYPE_MANUAL]
**DomesticHandlingFee** | Pointer to **float32** | The domestic handling fee added to buyer&#39;s shipping total - only available for calculated shipping profiles. | [optional] [default to 0]
**InternationalHandlingFee** | Pointer to **float32** | The international handling fee added to buyer&#39;s shipping total - only available for calculated shipping profiles. | [optional] [default to 0]

## Methods

### NewShopShippingProfile

`func NewShopShippingProfile() *ShopShippingProfile`

NewShopShippingProfile instantiates a new ShopShippingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopShippingProfileWithDefaults

`func NewShopShippingProfileWithDefaults() *ShopShippingProfile`

NewShopShippingProfileWithDefaults instantiates a new ShopShippingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingProfileId

`func (o *ShopShippingProfile) GetShippingProfileId() int64`

GetShippingProfileId returns the ShippingProfileId field if non-nil, zero value otherwise.

### GetShippingProfileIdOk

`func (o *ShopShippingProfile) GetShippingProfileIdOk() (*int64, bool)`

GetShippingProfileIdOk returns a tuple with the ShippingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileId

`func (o *ShopShippingProfile) SetShippingProfileId(v int64)`

SetShippingProfileId sets ShippingProfileId field to given value.

### HasShippingProfileId

`func (o *ShopShippingProfile) HasShippingProfileId() bool`

HasShippingProfileId returns a boolean if a field has been set.

### GetTitle

`func (o *ShopShippingProfile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShopShippingProfile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShopShippingProfile) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ShopShippingProfile) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ShopShippingProfile) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ShopShippingProfile) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserId

`func (o *ShopShippingProfile) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ShopShippingProfile) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ShopShippingProfile) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ShopShippingProfile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetMinProcessingDays

`func (o *ShopShippingProfile) GetMinProcessingDays() int64`

GetMinProcessingDays returns the MinProcessingDays field if non-nil, zero value otherwise.

### GetMinProcessingDaysOk

`func (o *ShopShippingProfile) GetMinProcessingDaysOk() (*int64, bool)`

GetMinProcessingDaysOk returns a tuple with the MinProcessingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProcessingDays

`func (o *ShopShippingProfile) SetMinProcessingDays(v int64)`

SetMinProcessingDays sets MinProcessingDays field to given value.

### HasMinProcessingDays

`func (o *ShopShippingProfile) HasMinProcessingDays() bool`

HasMinProcessingDays returns a boolean if a field has been set.

### SetMinProcessingDaysNil

`func (o *ShopShippingProfile) SetMinProcessingDaysNil(b bool)`

 SetMinProcessingDaysNil sets the value for MinProcessingDays to be an explicit nil

### UnsetMinProcessingDays
`func (o *ShopShippingProfile) UnsetMinProcessingDays()`

UnsetMinProcessingDays ensures that no value is present for MinProcessingDays, not even an explicit nil
### GetMaxProcessingDays

`func (o *ShopShippingProfile) GetMaxProcessingDays() int64`

GetMaxProcessingDays returns the MaxProcessingDays field if non-nil, zero value otherwise.

### GetMaxProcessingDaysOk

`func (o *ShopShippingProfile) GetMaxProcessingDaysOk() (*int64, bool)`

GetMaxProcessingDaysOk returns a tuple with the MaxProcessingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProcessingDays

`func (o *ShopShippingProfile) SetMaxProcessingDays(v int64)`

SetMaxProcessingDays sets MaxProcessingDays field to given value.

### HasMaxProcessingDays

`func (o *ShopShippingProfile) HasMaxProcessingDays() bool`

HasMaxProcessingDays returns a boolean if a field has been set.

### SetMaxProcessingDaysNil

`func (o *ShopShippingProfile) SetMaxProcessingDaysNil(b bool)`

 SetMaxProcessingDaysNil sets the value for MaxProcessingDays to be an explicit nil

### UnsetMaxProcessingDays
`func (o *ShopShippingProfile) UnsetMaxProcessingDays()`

UnsetMaxProcessingDays ensures that no value is present for MaxProcessingDays, not even an explicit nil
### GetProcessingDaysDisplayLabel

`func (o *ShopShippingProfile) GetProcessingDaysDisplayLabel() string`

GetProcessingDaysDisplayLabel returns the ProcessingDaysDisplayLabel field if non-nil, zero value otherwise.

### GetProcessingDaysDisplayLabelOk

`func (o *ShopShippingProfile) GetProcessingDaysDisplayLabelOk() (*string, bool)`

GetProcessingDaysDisplayLabelOk returns a tuple with the ProcessingDaysDisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingDaysDisplayLabel

`func (o *ShopShippingProfile) SetProcessingDaysDisplayLabel(v string)`

SetProcessingDaysDisplayLabel sets ProcessingDaysDisplayLabel field to given value.

### HasProcessingDaysDisplayLabel

`func (o *ShopShippingProfile) HasProcessingDaysDisplayLabel() bool`

HasProcessingDaysDisplayLabel returns a boolean if a field has been set.

### GetOriginCountryIso

`func (o *ShopShippingProfile) GetOriginCountryIso() string`

GetOriginCountryIso returns the OriginCountryIso field if non-nil, zero value otherwise.

### GetOriginCountryIsoOk

`func (o *ShopShippingProfile) GetOriginCountryIsoOk() (*string, bool)`

GetOriginCountryIsoOk returns a tuple with the OriginCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountryIso

`func (o *ShopShippingProfile) SetOriginCountryIso(v string)`

SetOriginCountryIso sets OriginCountryIso field to given value.

### HasOriginCountryIso

`func (o *ShopShippingProfile) HasOriginCountryIso() bool`

HasOriginCountryIso returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ShopShippingProfile) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ShopShippingProfile) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ShopShippingProfile) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ShopShippingProfile) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetShippingProfileDestinations

`func (o *ShopShippingProfile) GetShippingProfileDestinations() []ShopShippingProfileDestination`

GetShippingProfileDestinations returns the ShippingProfileDestinations field if non-nil, zero value otherwise.

### GetShippingProfileDestinationsOk

`func (o *ShopShippingProfile) GetShippingProfileDestinationsOk() (*[]ShopShippingProfileDestination, bool)`

GetShippingProfileDestinationsOk returns a tuple with the ShippingProfileDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileDestinations

`func (o *ShopShippingProfile) SetShippingProfileDestinations(v []ShopShippingProfileDestination)`

SetShippingProfileDestinations sets ShippingProfileDestinations field to given value.

### HasShippingProfileDestinations

`func (o *ShopShippingProfile) HasShippingProfileDestinations() bool`

HasShippingProfileDestinations returns a boolean if a field has been set.

### GetShippingProfileUpgrades

`func (o *ShopShippingProfile) GetShippingProfileUpgrades() []ShopShippingProfileUpgrade`

GetShippingProfileUpgrades returns the ShippingProfileUpgrades field if non-nil, zero value otherwise.

### GetShippingProfileUpgradesOk

`func (o *ShopShippingProfile) GetShippingProfileUpgradesOk() (*[]ShopShippingProfileUpgrade, bool)`

GetShippingProfileUpgradesOk returns a tuple with the ShippingProfileUpgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileUpgrades

`func (o *ShopShippingProfile) SetShippingProfileUpgrades(v []ShopShippingProfileUpgrade)`

SetShippingProfileUpgrades sets ShippingProfileUpgrades field to given value.

### HasShippingProfileUpgrades

`func (o *ShopShippingProfile) HasShippingProfileUpgrades() bool`

HasShippingProfileUpgrades returns a boolean if a field has been set.

### GetOriginPostalCode

`func (o *ShopShippingProfile) GetOriginPostalCode() string`

GetOriginPostalCode returns the OriginPostalCode field if non-nil, zero value otherwise.

### GetOriginPostalCodeOk

`func (o *ShopShippingProfile) GetOriginPostalCodeOk() (*string, bool)`

GetOriginPostalCodeOk returns a tuple with the OriginPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPostalCode

`func (o *ShopShippingProfile) SetOriginPostalCode(v string)`

SetOriginPostalCode sets OriginPostalCode field to given value.

### HasOriginPostalCode

`func (o *ShopShippingProfile) HasOriginPostalCode() bool`

HasOriginPostalCode returns a boolean if a field has been set.

### SetOriginPostalCodeNil

`func (o *ShopShippingProfile) SetOriginPostalCodeNil(b bool)`

 SetOriginPostalCodeNil sets the value for OriginPostalCode to be an explicit nil

### UnsetOriginPostalCode
`func (o *ShopShippingProfile) UnsetOriginPostalCode()`

UnsetOriginPostalCode ensures that no value is present for OriginPostalCode, not even an explicit nil
### GetProfileType

`func (o *ShopShippingProfile) GetProfileType() ShopShippingProfileProfileType`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *ShopShippingProfile) GetProfileTypeOk() (*ShopShippingProfileProfileType, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *ShopShippingProfile) SetProfileType(v ShopShippingProfileProfileType)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *ShopShippingProfile) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.

### GetDomesticHandlingFee

`func (o *ShopShippingProfile) GetDomesticHandlingFee() float32`

GetDomesticHandlingFee returns the DomesticHandlingFee field if non-nil, zero value otherwise.

### GetDomesticHandlingFeeOk

`func (o *ShopShippingProfile) GetDomesticHandlingFeeOk() (*float32, bool)`

GetDomesticHandlingFeeOk returns a tuple with the DomesticHandlingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticHandlingFee

`func (o *ShopShippingProfile) SetDomesticHandlingFee(v float32)`

SetDomesticHandlingFee sets DomesticHandlingFee field to given value.

### HasDomesticHandlingFee

`func (o *ShopShippingProfile) HasDomesticHandlingFee() bool`

HasDomesticHandlingFee returns a boolean if a field has been set.

### GetInternationalHandlingFee

`func (o *ShopShippingProfile) GetInternationalHandlingFee() float32`

GetInternationalHandlingFee returns the InternationalHandlingFee field if non-nil, zero value otherwise.

### GetInternationalHandlingFeeOk

`func (o *ShopShippingProfile) GetInternationalHandlingFeeOk() (*float32, bool)`

GetInternationalHandlingFeeOk returns a tuple with the InternationalHandlingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalHandlingFee

`func (o *ShopShippingProfile) SetInternationalHandlingFee(v float32)`

SetInternationalHandlingFee sets InternationalHandlingFee field to given value.

### HasInternationalHandlingFee

`func (o *ShopShippingProfile) HasInternationalHandlingFee() bool`

HasInternationalHandlingFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


