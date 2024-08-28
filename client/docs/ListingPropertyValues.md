# ListingPropertyValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]ListingPropertyValue**](ListingPropertyValue.md) |  | [optional] 

## Methods

### NewListingPropertyValues

`func NewListingPropertyValues() *ListingPropertyValues`

NewListingPropertyValues instantiates a new ListingPropertyValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingPropertyValuesWithDefaults

`func NewListingPropertyValuesWithDefaults() *ListingPropertyValues`

NewListingPropertyValuesWithDefaults instantiates a new ListingPropertyValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListingPropertyValues) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListingPropertyValues) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListingPropertyValues) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListingPropertyValues) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *ListingPropertyValues) GetResults() []ListingPropertyValue`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListingPropertyValues) GetResultsOk() (*[]ListingPropertyValue, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListingPropertyValues) SetResults(v []ListingPropertyValue)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListingPropertyValues) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


