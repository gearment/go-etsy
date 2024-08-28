# ShopSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopSectionId** | Pointer to **int64** | The numeric ID of a section in a specific Etsy shop. | [optional] 
**Title** | Pointer to **string** | The title string for a shop section. | [optional] 
**Rank** | Pointer to **int64** | The positive non-zero numeric position of this section in the section display order for a shop, with rank 1 sections appearing first. | [optional] 
**UserId** | Pointer to **int64** | The numeric ID of the [user](/documentation/reference#tag/User) who owns this shop section. | [optional] 
**ActiveListingCount** | Pointer to **int64** | The number of active listings in one section of a specific Etsy shop. | [optional] 

## Methods

### NewShopSection

`func NewShopSection() *ShopSection`

NewShopSection instantiates a new ShopSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopSectionWithDefaults

`func NewShopSectionWithDefaults() *ShopSection`

NewShopSectionWithDefaults instantiates a new ShopSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopSectionId

`func (o *ShopSection) GetShopSectionId() int64`

GetShopSectionId returns the ShopSectionId field if non-nil, zero value otherwise.

### GetShopSectionIdOk

`func (o *ShopSection) GetShopSectionIdOk() (*int64, bool)`

GetShopSectionIdOk returns a tuple with the ShopSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopSectionId

`func (o *ShopSection) SetShopSectionId(v int64)`

SetShopSectionId sets ShopSectionId field to given value.

### HasShopSectionId

`func (o *ShopSection) HasShopSectionId() bool`

HasShopSectionId returns a boolean if a field has been set.

### GetTitle

`func (o *ShopSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShopSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShopSection) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ShopSection) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRank

`func (o *ShopSection) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ShopSection) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ShopSection) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ShopSection) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetUserId

`func (o *ShopSection) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ShopSection) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ShopSection) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ShopSection) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetActiveListingCount

`func (o *ShopSection) GetActiveListingCount() int64`

GetActiveListingCount returns the ActiveListingCount field if non-nil, zero value otherwise.

### GetActiveListingCountOk

`func (o *ShopSection) GetActiveListingCountOk() (*int64, bool)`

GetActiveListingCountOk returns a tuple with the ActiveListingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveListingCount

`func (o *ShopSection) SetActiveListingCount(v int64)`

SetActiveListingCount sets ActiveListingCount field to given value.

### HasActiveListingCount

`func (o *ShopSection) HasActiveListingCount() bool`

HasActiveListingCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


