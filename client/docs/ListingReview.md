# ListingReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | Pointer to **int64** | The shop&#39;s numeric ID. | [optional] 
**ListingId** | Pointer to **int64** | The ID of the ShopListing that the TransactionReview belongs to. | [optional] 
**Rating** | Pointer to **int64** | Rating value on scale from 1 to 5 | [optional] 
**Review** | Pointer to **NullableString** | A message left by the author, explaining the feedback, if provided. | [optional] 
**Language** | Pointer to **string** | The language of the TransactionReview | [optional] 
**ImageUrlFullxfull** | Pointer to **NullableString** | The url to a photo provided with the feedback, dimensions fullxfull. Note: This field may be absent, depending on the buyer&#39;s privacy settings. | [optional] 
**CreateTimestamp** | Pointer to **int64** | The date and time the TransactionReview was created in epoch seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The date and time the TransactionReview was created in epoch seconds. | [optional] 
**UpdateTimestamp** | Pointer to **int64** | The date and time the TransactionReview was updated in epoch seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The date and time the TransactionReview was updated in epoch seconds. | [optional] 

## Methods

### NewListingReview

`func NewListingReview() *ListingReview`

NewListingReview instantiates a new ListingReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingReviewWithDefaults

`func NewListingReviewWithDefaults() *ListingReview`

NewListingReviewWithDefaults instantiates a new ListingReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *ListingReview) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ListingReview) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ListingReview) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ListingReview) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetListingId

`func (o *ListingReview) GetListingId() int64`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ListingReview) GetListingIdOk() (*int64, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ListingReview) SetListingId(v int64)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *ListingReview) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### GetRating

`func (o *ListingReview) GetRating() int64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ListingReview) GetRatingOk() (*int64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ListingReview) SetRating(v int64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ListingReview) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetReview

`func (o *ListingReview) GetReview() string`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *ListingReview) GetReviewOk() (*string, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *ListingReview) SetReview(v string)`

SetReview sets Review field to given value.

### HasReview

`func (o *ListingReview) HasReview() bool`

HasReview returns a boolean if a field has been set.

### SetReviewNil

`func (o *ListingReview) SetReviewNil(b bool)`

 SetReviewNil sets the value for Review to be an explicit nil

### UnsetReview
`func (o *ListingReview) UnsetReview()`

UnsetReview ensures that no value is present for Review, not even an explicit nil
### GetLanguage

`func (o *ListingReview) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ListingReview) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ListingReview) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ListingReview) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetImageUrlFullxfull

`func (o *ListingReview) GetImageUrlFullxfull() string`

GetImageUrlFullxfull returns the ImageUrlFullxfull field if non-nil, zero value otherwise.

### GetImageUrlFullxfullOk

`func (o *ListingReview) GetImageUrlFullxfullOk() (*string, bool)`

GetImageUrlFullxfullOk returns a tuple with the ImageUrlFullxfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrlFullxfull

`func (o *ListingReview) SetImageUrlFullxfull(v string)`

SetImageUrlFullxfull sets ImageUrlFullxfull field to given value.

### HasImageUrlFullxfull

`func (o *ListingReview) HasImageUrlFullxfull() bool`

HasImageUrlFullxfull returns a boolean if a field has been set.

### SetImageUrlFullxfullNil

`func (o *ListingReview) SetImageUrlFullxfullNil(b bool)`

 SetImageUrlFullxfullNil sets the value for ImageUrlFullxfull to be an explicit nil

### UnsetImageUrlFullxfull
`func (o *ListingReview) UnsetImageUrlFullxfull()`

UnsetImageUrlFullxfull ensures that no value is present for ImageUrlFullxfull, not even an explicit nil
### GetCreateTimestamp

`func (o *ListingReview) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *ListingReview) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *ListingReview) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *ListingReview) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ListingReview) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ListingReview) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ListingReview) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ListingReview) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *ListingReview) GetUpdateTimestamp() int64`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *ListingReview) GetUpdateTimestampOk() (*int64, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *ListingReview) SetUpdateTimestamp(v int64)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *ListingReview) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ListingReview) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ListingReview) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ListingReview) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *ListingReview) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


