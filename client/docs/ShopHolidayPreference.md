# ShopHolidayPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | Pointer to **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | [optional] 
**HolidayId** | Pointer to [**UpdateHolidayPreferencesHolidayIdParameter**](UpdateHolidayPreferencesHolidayIdParameter.md) |  | [optional] 
**CountryIso** | Pointer to **string** | The country iso where the shop is located. | [optional] 
**IsWorking** | Pointer to **bool** | A boolean value for whether the shop will process orders on a particular holiday. | [optional] 
**HolidayName** | Pointer to **string** | The name of the holiday that a country observes. | [optional] 

## Methods

### NewShopHolidayPreference

`func NewShopHolidayPreference() *ShopHolidayPreference`

NewShopHolidayPreference instantiates a new ShopHolidayPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopHolidayPreferenceWithDefaults

`func NewShopHolidayPreferenceWithDefaults() *ShopHolidayPreference`

NewShopHolidayPreferenceWithDefaults instantiates a new ShopHolidayPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *ShopHolidayPreference) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ShopHolidayPreference) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ShopHolidayPreference) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ShopHolidayPreference) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetHolidayId

`func (o *ShopHolidayPreference) GetHolidayId() UpdateHolidayPreferencesHolidayIdParameter`

GetHolidayId returns the HolidayId field if non-nil, zero value otherwise.

### GetHolidayIdOk

`func (o *ShopHolidayPreference) GetHolidayIdOk() (*UpdateHolidayPreferencesHolidayIdParameter, bool)`

GetHolidayIdOk returns a tuple with the HolidayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayId

`func (o *ShopHolidayPreference) SetHolidayId(v UpdateHolidayPreferencesHolidayIdParameter)`

SetHolidayId sets HolidayId field to given value.

### HasHolidayId

`func (o *ShopHolidayPreference) HasHolidayId() bool`

HasHolidayId returns a boolean if a field has been set.

### GetCountryIso

`func (o *ShopHolidayPreference) GetCountryIso() string`

GetCountryIso returns the CountryIso field if non-nil, zero value otherwise.

### GetCountryIsoOk

`func (o *ShopHolidayPreference) GetCountryIsoOk() (*string, bool)`

GetCountryIsoOk returns a tuple with the CountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso

`func (o *ShopHolidayPreference) SetCountryIso(v string)`

SetCountryIso sets CountryIso field to given value.

### HasCountryIso

`func (o *ShopHolidayPreference) HasCountryIso() bool`

HasCountryIso returns a boolean if a field has been set.

### GetIsWorking

`func (o *ShopHolidayPreference) GetIsWorking() bool`

GetIsWorking returns the IsWorking field if non-nil, zero value otherwise.

### GetIsWorkingOk

`func (o *ShopHolidayPreference) GetIsWorkingOk() (*bool, bool)`

GetIsWorkingOk returns a tuple with the IsWorking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorking

`func (o *ShopHolidayPreference) SetIsWorking(v bool)`

SetIsWorking sets IsWorking field to given value.

### HasIsWorking

`func (o *ShopHolidayPreference) HasIsWorking() bool`

HasIsWorking returns a boolean if a field has been set.

### GetHolidayName

`func (o *ShopHolidayPreference) GetHolidayName() string`

GetHolidayName returns the HolidayName field if non-nil, zero value otherwise.

### GetHolidayNameOk

`func (o *ShopHolidayPreference) GetHolidayNameOk() (*string, bool)`

GetHolidayNameOk returns a tuple with the HolidayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayName

`func (o *ShopHolidayPreference) SetHolidayName(v string)`

SetHolidayName sets HolidayName field to given value.

### HasHolidayName

`func (o *ShopHolidayPreference) HasHolidayName() bool`

HasHolidayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


