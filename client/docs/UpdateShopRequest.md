# UpdateShopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A brief heading string for the shop&#39;s main page. | [optional] 
**Announcement** | Pointer to **string** | An announcement string to buyers that displays on the shop&#39;s homepage. | [optional] 
**SaleMessage** | Pointer to **string** | A message string sent to users who complete a purchase from this shop. | [optional] 
**DigitalSaleMessage** | Pointer to **string** | A message string sent to users who purchase a digital item from this shop. | [optional] 
**PolicyAdditional** | Pointer to **string** | The shop&#39;s additional policies string (may be blank). | [optional] 

## Methods

### NewUpdateShopRequest

`func NewUpdateShopRequest() *UpdateShopRequest`

NewUpdateShopRequest instantiates a new UpdateShopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShopRequestWithDefaults

`func NewUpdateShopRequestWithDefaults() *UpdateShopRequest`

NewUpdateShopRequestWithDefaults instantiates a new UpdateShopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateShopRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateShopRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateShopRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateShopRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAnnouncement

`func (o *UpdateShopRequest) GetAnnouncement() string`

GetAnnouncement returns the Announcement field if non-nil, zero value otherwise.

### GetAnnouncementOk

`func (o *UpdateShopRequest) GetAnnouncementOk() (*string, bool)`

GetAnnouncementOk returns a tuple with the Announcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncement

`func (o *UpdateShopRequest) SetAnnouncement(v string)`

SetAnnouncement sets Announcement field to given value.

### HasAnnouncement

`func (o *UpdateShopRequest) HasAnnouncement() bool`

HasAnnouncement returns a boolean if a field has been set.

### GetSaleMessage

`func (o *UpdateShopRequest) GetSaleMessage() string`

GetSaleMessage returns the SaleMessage field if non-nil, zero value otherwise.

### GetSaleMessageOk

`func (o *UpdateShopRequest) GetSaleMessageOk() (*string, bool)`

GetSaleMessageOk returns a tuple with the SaleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleMessage

`func (o *UpdateShopRequest) SetSaleMessage(v string)`

SetSaleMessage sets SaleMessage field to given value.

### HasSaleMessage

`func (o *UpdateShopRequest) HasSaleMessage() bool`

HasSaleMessage returns a boolean if a field has been set.

### GetDigitalSaleMessage

`func (o *UpdateShopRequest) GetDigitalSaleMessage() string`

GetDigitalSaleMessage returns the DigitalSaleMessage field if non-nil, zero value otherwise.

### GetDigitalSaleMessageOk

`func (o *UpdateShopRequest) GetDigitalSaleMessageOk() (*string, bool)`

GetDigitalSaleMessageOk returns a tuple with the DigitalSaleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSaleMessage

`func (o *UpdateShopRequest) SetDigitalSaleMessage(v string)`

SetDigitalSaleMessage sets DigitalSaleMessage field to given value.

### HasDigitalSaleMessage

`func (o *UpdateShopRequest) HasDigitalSaleMessage() bool`

HasDigitalSaleMessage returns a boolean if a field has been set.

### GetPolicyAdditional

`func (o *UpdateShopRequest) GetPolicyAdditional() string`

GetPolicyAdditional returns the PolicyAdditional field if non-nil, zero value otherwise.

### GetPolicyAdditionalOk

`func (o *UpdateShopRequest) GetPolicyAdditionalOk() (*string, bool)`

GetPolicyAdditionalOk returns a tuple with the PolicyAdditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyAdditional

`func (o *UpdateShopRequest) SetPolicyAdditional(v string)`

SetPolicyAdditional sets PolicyAdditional field to given value.

### HasPolicyAdditional

`func (o *UpdateShopRequest) HasPolicyAdditional() bool`

HasPolicyAdditional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


