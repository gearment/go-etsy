# ShopReceiptTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of ShopReceiptTransaction resources found. | [optional] 
**Results** | Pointer to [**[]ShopReceiptTransaction**](ShopReceiptTransaction.md) | The ShopReceiptTransaction resources found. | [optional] 

## Methods

### NewShopReceiptTransactions

`func NewShopReceiptTransactions() *ShopReceiptTransactions`

NewShopReceiptTransactions instantiates a new ShopReceiptTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopReceiptTransactionsWithDefaults

`func NewShopReceiptTransactionsWithDefaults() *ShopReceiptTransactions`

NewShopReceiptTransactionsWithDefaults instantiates a new ShopReceiptTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ShopReceiptTransactions) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ShopReceiptTransactions) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ShopReceiptTransactions) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ShopReceiptTransactions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *ShopReceiptTransactions) GetResults() []ShopReceiptTransaction`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ShopReceiptTransactions) GetResultsOk() (*[]ShopReceiptTransaction, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ShopReceiptTransactions) SetResults(v []ShopReceiptTransaction)`

SetResults sets Results field to given value.

### HasResults

`func (o *ShopReceiptTransactions) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


