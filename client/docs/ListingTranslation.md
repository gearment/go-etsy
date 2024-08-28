# ListingTranslation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | Pointer to **int64** | The numeric ID for the Listing. | [optional] 
**Language** | Pointer to **string** | The IETF language tag (e.g. &#39;fr&#39;) for the language of this translation. | [optional] 
**Title** | Pointer to **NullableString** | The title of the Listing of this Translation. | [optional] 
**Description** | Pointer to **NullableString** | The description of the Listing of this Translation. | [optional] 
**Tags** | Pointer to **[]string** | The tags of the Listing of this Translation. | [optional] 

## Methods

### NewListingTranslation

`func NewListingTranslation() *ListingTranslation`

NewListingTranslation instantiates a new ListingTranslation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingTranslationWithDefaults

`func NewListingTranslationWithDefaults() *ListingTranslation`

NewListingTranslationWithDefaults instantiates a new ListingTranslation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *ListingTranslation) GetListingId() int64`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ListingTranslation) GetListingIdOk() (*int64, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ListingTranslation) SetListingId(v int64)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *ListingTranslation) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### GetLanguage

`func (o *ListingTranslation) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ListingTranslation) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ListingTranslation) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ListingTranslation) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetTitle

`func (o *ListingTranslation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListingTranslation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListingTranslation) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListingTranslation) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ListingTranslation) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListingTranslation) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ListingTranslation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListingTranslation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListingTranslation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListingTranslation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListingTranslation) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListingTranslation) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *ListingTranslation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListingTranslation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListingTranslation) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListingTranslation) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


