# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** | The numeric ID of a user. This number is also a valid shop ID for the user\\&#39;s shop. | [optional] 
**PrimaryEmail** | Pointer to **NullableString** | An email address string for the user\\&#39;s primary email address. Access to this field is granted on a case by case basis for third-party integrations that require full access | [optional] 
**FirstName** | Pointer to **NullableString** | The user\\&#39;s first name. | [optional] 
**LastName** | Pointer to **NullableString** | The user\\&#39;s last name. | [optional] 
**ImageUrl75x75** | Pointer to **NullableString** | The user\\&#39;s avatar URL. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *User) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *User) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *User) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *User) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPrimaryEmail

`func (o *User) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *User) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *User) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *User) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### SetPrimaryEmailNil

`func (o *User) SetPrimaryEmailNil(b bool)`

 SetPrimaryEmailNil sets the value for PrimaryEmail to be an explicit nil

### UnsetPrimaryEmail
`func (o *User) UnsetPrimaryEmail()`

UnsetPrimaryEmail ensures that no value is present for PrimaryEmail, not even an explicit nil
### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *User) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *User) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *User) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *User) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetImageUrl75x75

`func (o *User) GetImageUrl75x75() string`

GetImageUrl75x75 returns the ImageUrl75x75 field if non-nil, zero value otherwise.

### GetImageUrl75x75Ok

`func (o *User) GetImageUrl75x75Ok() (*string, bool)`

GetImageUrl75x75Ok returns a tuple with the ImageUrl75x75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl75x75

`func (o *User) SetImageUrl75x75(v string)`

SetImageUrl75x75 sets ImageUrl75x75 field to given value.

### HasImageUrl75x75

`func (o *User) HasImageUrl75x75() bool`

HasImageUrl75x75 returns a boolean if a field has been set.

### SetImageUrl75x75Nil

`func (o *User) SetImageUrl75x75Nil(b bool)`

 SetImageUrl75x75Nil sets the value for ImageUrl75x75 to be an explicit nil

### UnsetImageUrl75x75
`func (o *User) UnsetImageUrl75x75()`

UnsetImageUrl75x75 ensures that no value is present for ImageUrl75x75, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


