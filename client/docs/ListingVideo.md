# ListingVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoId** | Pointer to **int64** | The unique ID of a video associated with a listing. | [optional] 
**Height** | Pointer to **int64** | The video height dimension in pixels. | [optional] 
**Width** | Pointer to **int64** | The video width dimension in pixels. | [optional] 
**ThumbnailUrl** | Pointer to **string** | The url of the video thumbnail. | [optional] 
**VideoUrl** | Pointer to **string** | The url of the video file. | [optional] 
**VideoState** | Pointer to [**ListingVideoVideoState**](ListingVideoVideoState.md) |  | [optional] 

## Methods

### NewListingVideo

`func NewListingVideo() *ListingVideo`

NewListingVideo instantiates a new ListingVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingVideoWithDefaults

`func NewListingVideoWithDefaults() *ListingVideo`

NewListingVideoWithDefaults instantiates a new ListingVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoId

`func (o *ListingVideo) GetVideoId() int64`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *ListingVideo) GetVideoIdOk() (*int64, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *ListingVideo) SetVideoId(v int64)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *ListingVideo) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetHeight

`func (o *ListingVideo) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ListingVideo) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ListingVideo) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ListingVideo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *ListingVideo) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ListingVideo) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ListingVideo) SetWidth(v int64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ListingVideo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *ListingVideo) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *ListingVideo) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *ListingVideo) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *ListingVideo) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetVideoUrl

`func (o *ListingVideo) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *ListingVideo) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *ListingVideo) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *ListingVideo) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetVideoState

`func (o *ListingVideo) GetVideoState() ListingVideoVideoState`

GetVideoState returns the VideoState field if non-nil, zero value otherwise.

### GetVideoStateOk

`func (o *ListingVideo) GetVideoStateOk() (*ListingVideoVideoState, bool)`

GetVideoStateOk returns a tuple with the VideoState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoState

`func (o *ListingVideo) SetVideoState(v ListingVideoVideoState)`

SetVideoState sets VideoState field to given value.

### HasVideoState

`func (o *ListingVideo) HasVideoState() bool`

HasVideoState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


