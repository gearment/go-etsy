# ListingVideos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of results. | [optional] 
**Results** | Pointer to [**[]ListingVideo**](ListingVideo.md) | The list of requested resources. | [optional] 

## Methods

### NewListingVideos

`func NewListingVideos() *ListingVideos`

NewListingVideos instantiates a new ListingVideos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingVideosWithDefaults

`func NewListingVideosWithDefaults() *ListingVideos`

NewListingVideosWithDefaults instantiates a new ListingVideos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListingVideos) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListingVideos) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListingVideos) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListingVideos) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *ListingVideos) GetResults() []ListingVideo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListingVideos) GetResultsOk() (*[]ListingVideo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListingVideos) SetResults(v []ListingVideo)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListingVideos) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


