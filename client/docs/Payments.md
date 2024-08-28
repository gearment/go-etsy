# Payments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of payments in the response. | [optional] 
**Results** | Pointer to [**[]Payment**](Payment.md) | A list of payments. | [optional] 

## Methods

### NewPayments

`func NewPayments() *Payments`

NewPayments instantiates a new Payments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentsWithDefaults

`func NewPaymentsWithDefaults() *Payments`

NewPaymentsWithDefaults instantiates a new Payments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Payments) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Payments) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Payments) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Payments) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *Payments) GetResults() []Payment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Payments) GetResultsOk() (*[]Payment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Payments) SetResults(v []Payment)`

SetResults sets Results field to given value.

### HasResults

`func (o *Payments) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


