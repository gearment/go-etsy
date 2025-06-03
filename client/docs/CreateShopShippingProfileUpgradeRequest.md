# CreateShopShippingProfileUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CreateShopShippingProfileUpgradeRequestType**](CreateShopShippingProfileUpgradeRequestType.md) |  | 
**UpgradeName** | **string** | Name for the shipping upgrade shown to shoppers at checkout, e.g. USPS Priority. | 
**Price** | **float32** | Additional cost of adding the shipping upgrade. | 
**SecondaryPrice** | **float32** | Additional cost of adding the shipping upgrade for each additional item. | 
**ShippingCarrierId** | Pointer to **int64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] [default to 0]
**MailClass** | Pointer to **string** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [optional] 
**MinDeliveryDays** | Pointer to **int64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 
**MaxDeliveryDays** | Pointer to **int64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | [optional] 

## Methods

### NewCreateShopShippingProfileUpgradeRequest

`func NewCreateShopShippingProfileUpgradeRequest(type_ CreateShopShippingProfileUpgradeRequestType, upgradeName string, price float32, secondaryPrice float32, ) *CreateShopShippingProfileUpgradeRequest`

NewCreateShopShippingProfileUpgradeRequest instantiates a new CreateShopShippingProfileUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShopShippingProfileUpgradeRequestWithDefaults

`func NewCreateShopShippingProfileUpgradeRequestWithDefaults() *CreateShopShippingProfileUpgradeRequest`

NewCreateShopShippingProfileUpgradeRequestWithDefaults instantiates a new CreateShopShippingProfileUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateShopShippingProfileUpgradeRequest) GetType() CreateShopShippingProfileUpgradeRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateShopShippingProfileUpgradeRequest) GetTypeOk() (*CreateShopShippingProfileUpgradeRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateShopShippingProfileUpgradeRequest) SetType(v CreateShopShippingProfileUpgradeRequestType)`

SetType sets Type field to given value.


### GetUpgradeName

`func (o *CreateShopShippingProfileUpgradeRequest) GetUpgradeName() string`

GetUpgradeName returns the UpgradeName field if non-nil, zero value otherwise.

### GetUpgradeNameOk

`func (o *CreateShopShippingProfileUpgradeRequest) GetUpgradeNameOk() (*string, bool)`

GetUpgradeNameOk returns a tuple with the UpgradeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeName

`func (o *CreateShopShippingProfileUpgradeRequest) SetUpgradeName(v string)`

SetUpgradeName sets UpgradeName field to given value.


### GetPrice

`func (o *CreateShopShippingProfileUpgradeRequest) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateShopShippingProfileUpgradeRequest) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateShopShippingProfileUpgradeRequest) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetSecondaryPrice

`func (o *CreateShopShippingProfileUpgradeRequest) GetSecondaryPrice() float32`

GetSecondaryPrice returns the SecondaryPrice field if non-nil, zero value otherwise.

### GetSecondaryPriceOk

`func (o *CreateShopShippingProfileUpgradeRequest) GetSecondaryPriceOk() (*float32, bool)`

GetSecondaryPriceOk returns a tuple with the SecondaryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPrice

`func (o *CreateShopShippingProfileUpgradeRequest) SetSecondaryPrice(v float32)`

SetSecondaryPrice sets SecondaryPrice field to given value.


### GetShippingCarrierId

`func (o *CreateShopShippingProfileUpgradeRequest) GetShippingCarrierId() int64`

GetShippingCarrierId returns the ShippingCarrierId field if non-nil, zero value otherwise.

### GetShippingCarrierIdOk

`func (o *CreateShopShippingProfileUpgradeRequest) GetShippingCarrierIdOk() (*int64, bool)`

GetShippingCarrierIdOk returns a tuple with the ShippingCarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCarrierId

`func (o *CreateShopShippingProfileUpgradeRequest) SetShippingCarrierId(v int64)`

SetShippingCarrierId sets ShippingCarrierId field to given value.

### HasShippingCarrierId

`func (o *CreateShopShippingProfileUpgradeRequest) HasShippingCarrierId() bool`

HasShippingCarrierId returns a boolean if a field has been set.

### GetMailClass

`func (o *CreateShopShippingProfileUpgradeRequest) GetMailClass() string`

GetMailClass returns the MailClass field if non-nil, zero value otherwise.

### GetMailClassOk

`func (o *CreateShopShippingProfileUpgradeRequest) GetMailClassOk() (*string, bool)`

GetMailClassOk returns a tuple with the MailClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClass

`func (o *CreateShopShippingProfileUpgradeRequest) SetMailClass(v string)`

SetMailClass sets MailClass field to given value.

### HasMailClass

`func (o *CreateShopShippingProfileUpgradeRequest) HasMailClass() bool`

HasMailClass returns a boolean if a field has been set.

### GetMinDeliveryDays

`func (o *CreateShopShippingProfileUpgradeRequest) GetMinDeliveryDays() int64`

GetMinDeliveryDays returns the MinDeliveryDays field if non-nil, zero value otherwise.

### GetMinDeliveryDaysOk

`func (o *CreateShopShippingProfileUpgradeRequest) GetMinDeliveryDaysOk() (*int64, bool)`

GetMinDeliveryDaysOk returns a tuple with the MinDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDeliveryDays

`func (o *CreateShopShippingProfileUpgradeRequest) SetMinDeliveryDays(v int64)`

SetMinDeliveryDays sets MinDeliveryDays field to given value.

### HasMinDeliveryDays

`func (o *CreateShopShippingProfileUpgradeRequest) HasMinDeliveryDays() bool`

HasMinDeliveryDays returns a boolean if a field has been set.

### GetMaxDeliveryDays

`func (o *CreateShopShippingProfileUpgradeRequest) GetMaxDeliveryDays() int64`

GetMaxDeliveryDays returns the MaxDeliveryDays field if non-nil, zero value otherwise.

### GetMaxDeliveryDaysOk

`func (o *CreateShopShippingProfileUpgradeRequest) GetMaxDeliveryDaysOk() (*int64, bool)`

GetMaxDeliveryDaysOk returns a tuple with the MaxDeliveryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveryDays

`func (o *CreateShopShippingProfileUpgradeRequest) SetMaxDeliveryDays(v int64)`

SetMaxDeliveryDays sets MaxDeliveryDays field to given value.

### HasMaxDeliveryDays

`func (o *CreateShopShippingProfileUpgradeRequest) HasMaxDeliveryDays() bool`

HasMaxDeliveryDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


