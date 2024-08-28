# ShopListingFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingFileId** | Pointer to **int64** | The unique numeric ID of a file associated with a digital listing. | [optional] 
**ListingId** | Pointer to **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | [optional] 
**Rank** | Pointer to **int64** | The numeric index of the display order position of this file in the listing, starting at 1. | [optional] 
**Filename** | Pointer to **string** | The file name string for a file associated with a digital listing. | [optional] 
**Filesize** | Pointer to **string** | A human-readable format size string for the size of a file. | [optional] 
**SizeBytes** | Pointer to **int64** | A number indicating the size of a file, measured in bytes. | [optional] 
**Filetype** | Pointer to **string** | A type string indicating a file&#39;s MIME type. | [optional] 
**CreateTimestamp** | Pointer to **int64** | The unique numeric ID of a file associated with a digital listing. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The unique numeric ID of a file associated with a digital listing. | [optional] 

## Methods

### NewShopListingFile

`func NewShopListingFile() *ShopListingFile`

NewShopListingFile instantiates a new ShopListingFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopListingFileWithDefaults

`func NewShopListingFileWithDefaults() *ShopListingFile`

NewShopListingFileWithDefaults instantiates a new ShopListingFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingFileId

`func (o *ShopListingFile) GetListingFileId() int64`

GetListingFileId returns the ListingFileId field if non-nil, zero value otherwise.

### GetListingFileIdOk

`func (o *ShopListingFile) GetListingFileIdOk() (*int64, bool)`

GetListingFileIdOk returns a tuple with the ListingFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingFileId

`func (o *ShopListingFile) SetListingFileId(v int64)`

SetListingFileId sets ListingFileId field to given value.

### HasListingFileId

`func (o *ShopListingFile) HasListingFileId() bool`

HasListingFileId returns a boolean if a field has been set.

### GetListingId

`func (o *ShopListingFile) GetListingId() int64`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ShopListingFile) GetListingIdOk() (*int64, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ShopListingFile) SetListingId(v int64)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *ShopListingFile) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### GetRank

`func (o *ShopListingFile) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ShopListingFile) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ShopListingFile) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ShopListingFile) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetFilename

`func (o *ShopListingFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ShopListingFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ShopListingFile) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ShopListingFile) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFilesize

`func (o *ShopListingFile) GetFilesize() string`

GetFilesize returns the Filesize field if non-nil, zero value otherwise.

### GetFilesizeOk

`func (o *ShopListingFile) GetFilesizeOk() (*string, bool)`

GetFilesizeOk returns a tuple with the Filesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesize

`func (o *ShopListingFile) SetFilesize(v string)`

SetFilesize sets Filesize field to given value.

### HasFilesize

`func (o *ShopListingFile) HasFilesize() bool`

HasFilesize returns a boolean if a field has been set.

### GetSizeBytes

`func (o *ShopListingFile) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ShopListingFile) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ShopListingFile) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *ShopListingFile) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetFiletype

`func (o *ShopListingFile) GetFiletype() string`

GetFiletype returns the Filetype field if non-nil, zero value otherwise.

### GetFiletypeOk

`func (o *ShopListingFile) GetFiletypeOk() (*string, bool)`

GetFiletypeOk returns a tuple with the Filetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiletype

`func (o *ShopListingFile) SetFiletype(v string)`

SetFiletype sets Filetype field to given value.

### HasFiletype

`func (o *ShopListingFile) HasFiletype() bool`

HasFiletype returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *ShopListingFile) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *ShopListingFile) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *ShopListingFile) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *ShopListingFile) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ShopListingFile) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ShopListingFile) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ShopListingFile) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ShopListingFile) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


