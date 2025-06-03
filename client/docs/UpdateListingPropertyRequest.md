# UpdateListingPropertyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueIds** | **[]int64** | An array of unique IDs of multiple Etsy [listing property](/documentation/reference#operation/getListingProperties) values. For example, if your listing offers different sizes of a product, then the value ID list contains value IDs for each size. | 
**Values** | **[]string** | An array of value strings for multiple Etsy [listing property](/documentation/reference#operation/getListingProperties) values. For example, if your listing offers different colored products, then the values array contains the color strings for each color. Note: parenthesis characters (&#x60;(&#x60; and &#x60;)&#x60;) are not allowed. | 
**ScaleId** | Pointer to **int64** | The numeric ID of a single Etsy.com measurement scale. For example, for shoe size, there are three &#x60;scale_id&#x60;s available - &#x60;UK&#x60;, &#x60;US/Canada&#x60;, and &#x60;EU&#x60;, where &#x60;US/Canada&#x60; has &#x60;scale_id&#x60; 19. | [optional] 

## Methods

### NewUpdateListingPropertyRequest

`func NewUpdateListingPropertyRequest(valueIds []int64, values []string, ) *UpdateListingPropertyRequest`

NewUpdateListingPropertyRequest instantiates a new UpdateListingPropertyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListingPropertyRequestWithDefaults

`func NewUpdateListingPropertyRequestWithDefaults() *UpdateListingPropertyRequest`

NewUpdateListingPropertyRequestWithDefaults instantiates a new UpdateListingPropertyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueIds

`func (o *UpdateListingPropertyRequest) GetValueIds() []int64`

GetValueIds returns the ValueIds field if non-nil, zero value otherwise.

### GetValueIdsOk

`func (o *UpdateListingPropertyRequest) GetValueIdsOk() (*[]int64, bool)`

GetValueIdsOk returns a tuple with the ValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueIds

`func (o *UpdateListingPropertyRequest) SetValueIds(v []int64)`

SetValueIds sets ValueIds field to given value.


### GetValues

`func (o *UpdateListingPropertyRequest) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *UpdateListingPropertyRequest) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *UpdateListingPropertyRequest) SetValues(v []string)`

SetValues sets Values field to given value.


### GetScaleId

`func (o *UpdateListingPropertyRequest) GetScaleId() int64`

GetScaleId returns the ScaleId field if non-nil, zero value otherwise.

### GetScaleIdOk

`func (o *UpdateListingPropertyRequest) GetScaleIdOk() (*int64, bool)`

GetScaleIdOk returns a tuple with the ScaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleId

`func (o *UpdateListingPropertyRequest) SetScaleId(v int64)`

SetScaleId sets ScaleId field to given value.

### HasScaleId

`func (o *UpdateListingPropertyRequest) HasScaleId() bool`

HasScaleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


