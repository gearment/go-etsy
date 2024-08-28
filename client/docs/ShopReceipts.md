# ShopReceipts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of Shop Receipts found. | [optional] 
**Results** | Pointer to [**[]ShopReceipt**](ShopReceipt.md) | List of Shop Receipt resources found, with all Shop Receipt fields for each resource. | [optional] 

## Methods

### NewShopReceipts

`func NewShopReceipts() *ShopReceipts`

NewShopReceipts instantiates a new ShopReceipts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopReceiptsWithDefaults

`func NewShopReceiptsWithDefaults() *ShopReceipts`

NewShopReceiptsWithDefaults instantiates a new ShopReceipts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ShopReceipts) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ShopReceipts) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ShopReceipts) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ShopReceipts) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *ShopReceipts) GetResults() []ShopReceipt`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ShopReceipts) GetResultsOk() (*[]ShopReceipt, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ShopReceipts) SetResults(v []ShopReceipt)`

SetResults sets Results field to given value.

### HasResults

`func (o *ShopReceipts) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


