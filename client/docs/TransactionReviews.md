# TransactionReviews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of TransactionReview resources found. | [optional] 
**Results** | Pointer to [**[]TransactionReview**](TransactionReview.md) | The TransactionReview resources found. | [optional] 

## Methods

### NewTransactionReviews

`func NewTransactionReviews() *TransactionReviews`

NewTransactionReviews instantiates a new TransactionReviews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReviewsWithDefaults

`func NewTransactionReviewsWithDefaults() *TransactionReviews`

NewTransactionReviewsWithDefaults instantiates a new TransactionReviews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TransactionReviews) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransactionReviews) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransactionReviews) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TransactionReviews) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *TransactionReviews) GetResults() []TransactionReview`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TransactionReviews) GetResultsOk() (*[]TransactionReview, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TransactionReviews) SetResults(v []TransactionReview)`

SetResults sets Results field to given value.

### HasResults

`func (o *TransactionReviews) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


