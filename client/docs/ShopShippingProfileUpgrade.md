# ShopShippingProfileUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingProfileId** | Pointer to **int64** | The numeric ID of the base shipping profile. | [optional] 
**UpgradeId** | Pointer to **int64** | The numeric ID that is associated with a shipping upgrade | [optional] 
**UpgradeName** | Pointer to **string** | Name for the shipping upgrade shown to shoppers at checkout, e.g. USPS Priority. | [optional] 
**Type** | Pointer to [**CreateShopShippingProfileUpgradeRequestType**](CreateShopShippingProfileUpgradeRequestType.md) |  | [optional] 
**Rank** | Pointer to **int64** | The positive non-zero numeric position in the images displayed in a listing, with rank 1 images appearing in the left-most position in a listing. | [optional] 
**Language** | Pointer to **string** | The IETF language tag for the language of the shipping profile. Ex: &#x60;de&#x60;, &#x60;en&#x60;, &#x60;es&#x60;, &#x60;fr&#x60;, &#x60;it&#x60;, &#x60;ja&#x60;, &#x60;nl&#x60;, &#x60;pl&#x60;, &#x60;pt&#x60; | [optional] 
**Price** | Pointer to [**Money**](Money.md) | Additional cost of adding the shipping upgrade. | [optional] 
**SecondaryPrice** | Pointer to [**Money**](Money.md) | Additional cost of adding the shipping upgrade for each additional item. | [optional] 
**ShippingCarrierId** | Pointer to **NullableInt64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] 
**MailClass** | Pointer to **NullableString** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] 
**MinDeliveryDays** | Pointer to **NullableInt64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 
**MaxDeliveryDays** | Pointer to **NullableInt64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 

## Methods

### NewShopShippingProfileUpgrade

`func NewShopShippingProfileUpgrade() *ShopShippingProfileUpgrade`

NewShopShippingProfileUpgrade instantiates a new ShopShippingProfileUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopShippingProfileUpgradeWithDefaults

`func NewShopShippingProfileUpgradeWithDefaults() *ShopShippingProfileUpgrade`

NewShopShippingProfileUpgradeWithDefaults instantiates a new ShopShippingProfileUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingProfileId

`func (o *ShopShippingProfileUpgrade) GetShippingProfileId() int64`

GetShippingProfileId returns the ShippingProfileId field if non-nil, zero value otherwise.

### GetShippingProfileIdOk

`func (o *ShopShippingProfileUpgrade) GetShippingProfileIdOk() (*int64, bool)`

GetShippingProfileIdOk returns a tuple with the ShippingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileId

`func (o *ShopShippingProfileUpgrade) SetShippingProfileId(v int64)`

SetShippingProfileId sets ShippingProfileId field to given value.

### HasShippingProfileId

`func (o *ShopShippingProfileUpgrade) HasShippingProfileId() bool`

HasShippingProfileId returns a boolean if a field has been set.

### GetUpgradeId

`func (o *ShopShippingProfileUpgrade) GetUpgradeId() int64`

GetUpgradeId returns the UpgradeId field if non-nil, zero value otherwise.

### GetUpgradeIdOk

`func (o *ShopShippingProfileUpgrade) GetUpgradeIdOk() (*int64, bool)`

GetUpgradeIdOk returns a tuple with the UpgradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeId

`func (o *ShopShippingProfileUpgrade) SetUpgradeId(v int64)`

SetUpgradeId sets UpgradeId field to given value.

### HasUpgradeId

`func (o *ShopShippingProfileUpgrade) HasUpgradeId() bool`

HasUpgradeId returns a boolean if a field has been set.

### GetUpgradeName

`func (o *ShopShippingProfileUpgrade) GetUpgradeName() string`

GetUpgradeName returns the UpgradeName field if non-nil, zero value otherwise.

### GetUpgradeNameOk

`func (o *ShopShippingProfileUpgrade) GetUpgradeNameOk() (*string, bool)`

GetUpgradeNameOk returns a tuple with the UpgradeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeName

`func (o *ShopShippingProfileUpgrade) SetUpgradeName(v string)`

SetUpgradeName sets UpgradeName field to given value.

### HasUpgradeName

`func (o *ShopShippingProfileUpgrade) HasUpgradeName() bool`

HasUpgradeName returns a boolean if a field has been set.

### GetType

`func (o *ShopShippingProfileUpgrade) GetType() CreateShopShippingProfileUpgradeRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShopShippingProfileUpgrade) GetTypeOk() (*CreateShopShippingProfileUpgradeRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShopShippingProfileUpgrade) SetType(v CreateShopShippingProfileUpgradeRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *ShopShippingProfileUpgrade) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRank

`func (o *ShopShippingProfileUpgrade) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ShopShippingProfileUpgrade) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ShopShippingProfileUpgrade) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ShopShippingProfileUpgrade) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetLanguage

`func (o *ShopShippingProfileUpgrade) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ShopShippingProfileUpgrade) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ShopShippingProfileUpgrade) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ShopShippingProfileUpgrade) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPrice

`func (o *ShopShippingProfileUpgrade) GetPrice() Money`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ShopShippingProfileUpgrade) GetPriceOk() (*Money, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ShopShippingProfileUpgrade) SetPrice(v Money)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ShopShippingProfileUpgrade) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSecondaryPrice

`func (o *ShopShippingProfileUpgrade) GetSecondaryPrice() Money`

GetSecondaryPrice returns the SecondaryPrice field if non-nil, zero value otherwise.

### GetSecondaryPriceOk

`func (o *ShopShippingProfileUpgrade) GetSecondaryPriceOk() (*Money, bool)`

GetSecondaryPriceOk returns a tuple with the SecondaryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPrice

`func (o *ShopShippingProfileUpgrade) SetSecondaryPrice(v Money)`

SetSecondaryPrice sets SecondaryPrice field to given value.

### HasSecondaryPrice

`func (o *ShopShippingProfileUpgrade) HasSecondaryPrice() bool`

HasSecondaryPrice returns a boolean if a field has been set.

### GetShippingCarrierId

`func (o *ShopShippingProfileUpgrade) GetShippingCarrierId() int64`

GetShippingCarrierId returns the ShippingCarrierId field if non-nil, zero value otherwise.

### GetShippingCarrierIdOk

`func (o *ShopShippingProfileUpgrade) GetShippingCarrierIdOk() (*int64, bool)`

GetShippingCarrierIdOk returns a tuple with the ShippingCarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCarrierId

`func (o *ShopShippingProfileUpgrade) SetShippingCarrierId(v int64)`

SetShippingCarrierId sets ShippingCarrierId field to given value.

### HasShippingCarrierId

`func (o *ShopShippingProfileUpgrade) HasShippingCarrierId() bool`

HasShippingCarrierId returns a boolean if a field has been set.

### SetShippingCarrierIdNil

`func (o *ShopShippingProfileUpgrade) SetShippingCarrierIdNil(b bool)`

 SetShippingCarrierIdNil sets the value for ShippingCarrierId to be an explicit nil

### UnsetShippingCarrierId
`func (o *ShopShippingProfileUpgrade) UnsetShippingCarrierId()`

UnsetShippingCarrierId ensures that no value is present for ShippingCarrierId, not even an explicit nil
### GetMailClass

`func (o *ShopShippingProfileUpgrade) GetMailClass() string`

GetMailClass returns the MailClass field if non-nil, zero value otherwise.

### GetMailClassOk

`func (o *ShopShippingProfileUpgrade) GetMailClassOk() (*string, bool)`

GetMailClassOk returns a tuple with the MailClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClass

`func (o *ShopShippingProfileUpgrade) SetMailClass(v string)`

SetMailClass sets MailClass field to given value.

### HasMailClass

`func (o *ShopShippingProfileUpgrade) HasMailClass() bool`

HasMailClass returns a boolean if a field has been set.

### SetMailClassNil

`func (o *ShopShippingProfileUpgrade) SetMailClassNil(b bool)`

 SetMailClassNil sets the value for MailClass to be an explicit nil

### UnsetMailClass
`func (o *ShopShippingProfileUpgrade) UnsetMailClass()`

UnsetMailClass ensures that no value is present for MailClass, not even an explicit nil
### GetMinDeliveryDays

`func (o *ShopShippingProfileUpgrade) GetMinDeliveryDays() int64`

GetMinDeliveryDays returns the MinDeliveryDays field if non-nil, zero value otherwise.

### GetMinDeliveryDaysOk

`func (o *ShopShippingProfileUpgrade) GetMinDeliveryDaysOk() (*int64, bool)`

GetMinDeliveryDaysOk returns a tuple with the MinDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDeliveryDays

`func (o *ShopShippingProfileUpgrade) SetMinDeliveryDays(v int64)`

SetMinDeliveryDays sets MinDeliveryDays field to given value.

### HasMinDeliveryDays

`func (o *ShopShippingProfileUpgrade) HasMinDeliveryDays() bool`

HasMinDeliveryDays returns a boolean if a field has been set.

### SetMinDeliveryDaysNil

`func (o *ShopShippingProfileUpgrade) SetMinDeliveryDaysNil(b bool)`

 SetMinDeliveryDaysNil sets the value for MinDeliveryDays to be an explicit nil

### UnsetMinDeliveryDays
`func (o *ShopShippingProfileUpgrade) UnsetMinDeliveryDays()`

UnsetMinDeliveryDays ensures that no value is present for MinDeliveryDays, not even an explicit nil
### GetMaxDeliveryDays

`func (o *ShopShippingProfileUpgrade) GetMaxDeliveryDays() int64`

GetMaxDeliveryDays returns the MaxDeliveryDays field if non-nil, zero value otherwise.

### GetMaxDeliveryDaysOk

`func (o *ShopShippingProfileUpgrade) GetMaxDeliveryDaysOk() (*int64, bool)`

GetMaxDeliveryDaysOk returns a tuple with the MaxDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveryDays

`func (o *ShopShippingProfileUpgrade) SetMaxDeliveryDays(v int64)`

SetMaxDeliveryDays sets MaxDeliveryDays field to given value.

### HasMaxDeliveryDays

`func (o *ShopShippingProfileUpgrade) HasMaxDeliveryDays() bool`

HasMaxDeliveryDays returns a boolean if a field has been set.

### SetMaxDeliveryDaysNil

`func (o *ShopShippingProfileUpgrade) SetMaxDeliveryDaysNil(b bool)`

 SetMaxDeliveryDaysNil sets the value for MaxDeliveryDays to be an explicit nil

### UnsetMaxDeliveryDays
`func (o *ShopShippingProfileUpgrade) UnsetMaxDeliveryDays()`

UnsetMaxDeliveryDays ensures that no value is present for MaxDeliveryDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


