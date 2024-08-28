# Self

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** | The numeric ID of a user. This number is also a valid shop ID for the user\\&#39;s shop. | [optional] 
**ShopId** | Pointer to **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | [optional] 

## Methods

### NewSelf

`func NewSelf() *Self`

NewSelf instantiates a new Self object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfWithDefaults

`func NewSelfWithDefaults() *Self`

NewSelfWithDefaults instantiates a new Self object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *Self) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Self) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Self) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Self) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetShopId

`func (o *Self) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Self) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Self) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Self) HasShopId() bool`

HasShopId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


