# TransactionReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | Pointer to **int64** | The shop&#39;s numeric ID. | [optional] 
**ListingId** | Pointer to **int64** | The ID of the ShopListing that the TransactionReview belongs to. | [optional] 
**TransactionId** | Pointer to **int64** | The ID of the ShopReceipt Transaction that the TransactionReview belongs to. | [optional] 
**BuyerUserId** | Pointer to **NullableInt64** | The numeric ID of the user who was the buyer in this transaction. Note: This field may be absent, depending on the buyer&#39;s privacy settings. | [optional] 
**Rating** | Pointer to **int64** | Rating value on scale from 1 to 5 | [optional] 
**Review** | Pointer to **string** | A message left by the author, explaining the feedback, if provided. | [optional] [default to ""]
**Language** | Pointer to **string** | The language of the TransactionReview | [optional] 
**ImageUrlFullxfull** | Pointer to **NullableString** | The url to a photo provided with the feedback, dimensions fullxfull. Note: This field may be absent, depending on the buyer&#39;s privacy settings. | [optional] 
**CreateTimestamp** | Pointer to **int64** | The date and time the TransactionReview was created in epoch seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The date and time the TransactionReview was created in epoch seconds. | [optional] 
**UpdateTimestamp** | Pointer to **int64** | The date and time the TransactionReview was updated in epoch seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The date and time the TransactionReview was updated in epoch seconds. | [optional] 

## Methods

### NewTransactionReview

`func NewTransactionReview() *TransactionReview`

NewTransactionReview instantiates a new TransactionReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReviewWithDefaults

`func NewTransactionReviewWithDefaults() *TransactionReview`

NewTransactionReviewWithDefaults instantiates a new TransactionReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *TransactionReview) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *TransactionReview) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *TransactionReview) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *TransactionReview) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetListingId

`func (o *TransactionReview) GetListingId() int64`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *TransactionReview) GetListingIdOk() (*int64, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *TransactionReview) SetListingId(v int64)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *TransactionReview) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### GetTransactionId

`func (o *TransactionReview) GetTransactionId() int64`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionReview) GetTransactionIdOk() (*int64, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionReview) SetTransactionId(v int64)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TransactionReview) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetBuyerUserId

`func (o *TransactionReview) GetBuyerUserId() int64`

GetBuyerUserId returns the BuyerUserId field if non-nil, zero value otherwise.

### GetBuyerUserIdOk

`func (o *TransactionReview) GetBuyerUserIdOk() (*int64, bool)`

GetBuyerUserIdOk returns a tuple with the BuyerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUserId

`func (o *TransactionReview) SetBuyerUserId(v int64)`

SetBuyerUserId sets BuyerUserId field to given value.

### HasBuyerUserId

`func (o *TransactionReview) HasBuyerUserId() bool`

HasBuyerUserId returns a boolean if a field has been set.

### SetBuyerUserIdNil

`func (o *TransactionReview) SetBuyerUserIdNil(b bool)`

 SetBuyerUserIdNil sets the value for BuyerUserId to be an explicit nil

### UnsetBuyerUserId
`func (o *TransactionReview) UnsetBuyerUserId()`

UnsetBuyerUserId ensures that no value is present for BuyerUserId, not even an explicit nil
### GetRating

`func (o *TransactionReview) GetRating() int64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *TransactionReview) GetRatingOk() (*int64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *TransactionReview) SetRating(v int64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *TransactionReview) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetReview

`func (o *TransactionReview) GetReview() string`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *TransactionReview) GetReviewOk() (*string, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *TransactionReview) SetReview(v string)`

SetReview sets Review field to given value.

### HasReview

`func (o *TransactionReview) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetLanguage

`func (o *TransactionReview) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TransactionReview) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TransactionReview) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *TransactionReview) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetImageUrlFullxfull

`func (o *TransactionReview) GetImageUrlFullxfull() string`

GetImageUrlFullxfull returns the ImageUrlFullxfull field if non-nil, zero value otherwise.

### GetImageUrlFullxfullOk

`func (o *TransactionReview) GetImageUrlFullxfullOk() (*string, bool)`

GetImageUrlFullxfullOk returns a tuple with the ImageUrlFullxfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrlFullxfull

`func (o *TransactionReview) SetImageUrlFullxfull(v string)`

SetImageUrlFullxfull sets ImageUrlFullxfull field to given value.

### HasImageUrlFullxfull

`func (o *TransactionReview) HasImageUrlFullxfull() bool`

HasImageUrlFullxfull returns a boolean if a field has been set.

### SetImageUrlFullxfullNil

`func (o *TransactionReview) SetImageUrlFullxfullNil(b bool)`

 SetImageUrlFullxfullNil sets the value for ImageUrlFullxfull to be an explicit nil

### UnsetImageUrlFullxfull
`func (o *TransactionReview) UnsetImageUrlFullxfull()`

UnsetImageUrlFullxfull ensures that no value is present for ImageUrlFullxfull, not even an explicit nil
### GetCreateTimestamp

`func (o *TransactionReview) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *TransactionReview) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *TransactionReview) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *TransactionReview) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *TransactionReview) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TransactionReview) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TransactionReview) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TransactionReview) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *TransactionReview) GetUpdateTimestamp() int64`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *TransactionReview) GetUpdateTimestampOk() (*int64, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *TransactionReview) SetUpdateTimestamp(v int64)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *TransactionReview) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *TransactionReview) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *TransactionReview) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *TransactionReview) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *TransactionReview) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


