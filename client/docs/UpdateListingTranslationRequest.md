# UpdateListingTranslationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the Listing of this Translation. | 
**Description** | **string** | The description of the Listing of this Translation. | 
**Tags** | Pointer to **[]string** | The tags of the Listing of this Translation. | [optional] 

## Methods

### NewUpdateListingTranslationRequest

`func NewUpdateListingTranslationRequest(title string, description string, ) *UpdateListingTranslationRequest`

NewUpdateListingTranslationRequest instantiates a new UpdateListingTranslationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListingTranslationRequestWithDefaults

`func NewUpdateListingTranslationRequestWithDefaults() *UpdateListingTranslationRequest`

NewUpdateListingTranslationRequestWithDefaults instantiates a new UpdateListingTranslationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateListingTranslationRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateListingTranslationRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateListingTranslationRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *UpdateListingTranslationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateListingTranslationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateListingTranslationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *UpdateListingTranslationRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateListingTranslationRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateListingTranslationRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateListingTranslationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


