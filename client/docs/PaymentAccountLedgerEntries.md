# PaymentAccountLedgerEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of PaymentAccountLedgerEntry resources found. | [optional] 
**Results** | Pointer to [**[]PaymentAccountLedgerEntry**](PaymentAccountLedgerEntry.md) | The PaymentAccountLedgerEntry resources found. | [optional] 

## Methods

### NewPaymentAccountLedgerEntries

`func NewPaymentAccountLedgerEntries() *PaymentAccountLedgerEntries`

NewPaymentAccountLedgerEntries instantiates a new PaymentAccountLedgerEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAccountLedgerEntriesWithDefaults

`func NewPaymentAccountLedgerEntriesWithDefaults() *PaymentAccountLedgerEntries`

NewPaymentAccountLedgerEntriesWithDefaults instantiates a new PaymentAccountLedgerEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaymentAccountLedgerEntries) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaymentAccountLedgerEntries) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaymentAccountLedgerEntries) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaymentAccountLedgerEntries) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaymentAccountLedgerEntries) GetResults() []PaymentAccountLedgerEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaymentAccountLedgerEntries) GetResultsOk() (*[]PaymentAccountLedgerEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaymentAccountLedgerEntries) SetResults(v []PaymentAccountLedgerEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaymentAccountLedgerEntries) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


