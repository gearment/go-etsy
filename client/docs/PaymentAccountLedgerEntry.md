# PaymentAccountLedgerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | Pointer to **int64** | The ledger entry&#39;s numeric ID. | [optional] 
**LedgerId** | Pointer to **int64** | The ledger&#39;s numeric ID. | [optional] 
**SequenceNumber** | Pointer to **int64** | The sequence allows ledger entries to be sorted chronologically. The higher the sequence, the more recent the entry. | [optional] 
**Amount** | Pointer to **int64** | The amount of money credited to the ledger. | [optional] 
**Currency** | Pointer to **string** | The currency of the entry on the ledger. | [optional] 
**Description** | Pointer to **string** | Details what kind of ledger entry this is: a payment, refund, reversal of a failed refund, disbursement, returned disbursement, recoupment, miscellaneous credit, miscellaneous debit, or bill payment. | [optional] 
**Balance** | Pointer to **int64** | The amount of money in the shop&#39;s ledger the moment after this entry was applied. | [optional] 
**CreateDate** | Pointer to **int64** | The date and time the ledger entry was created in Epoch seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The date and time the ledger entry was created in Epoch seconds. | [optional] 
**LedgerType** | Pointer to **string** | The original reference type for the ledger entry. | [optional] 
**ReferenceType** | Pointer to **string** | The object type the ledger entry refers to. | [optional] 
**ReferenceId** | Pointer to **NullableString** | The object id the ledger entry refers to. | [optional] 
**PaymentAdjustments** | Pointer to [**[]PaymentAdjustment**](PaymentAdjustment.md) | List of refund objects on an Etsy Payments transaction. All monetary amounts are in USD pennies unless otherwise specified. | [optional] 

## Methods

### NewPaymentAccountLedgerEntry

`func NewPaymentAccountLedgerEntry() *PaymentAccountLedgerEntry`

NewPaymentAccountLedgerEntry instantiates a new PaymentAccountLedgerEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAccountLedgerEntryWithDefaults

`func NewPaymentAccountLedgerEntryWithDefaults() *PaymentAccountLedgerEntry`

NewPaymentAccountLedgerEntryWithDefaults instantiates a new PaymentAccountLedgerEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *PaymentAccountLedgerEntry) GetEntryId() int64`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *PaymentAccountLedgerEntry) GetEntryIdOk() (*int64, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *PaymentAccountLedgerEntry) SetEntryId(v int64)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *PaymentAccountLedgerEntry) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetLedgerId

`func (o *PaymentAccountLedgerEntry) GetLedgerId() int64`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *PaymentAccountLedgerEntry) GetLedgerIdOk() (*int64, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *PaymentAccountLedgerEntry) SetLedgerId(v int64)`

SetLedgerId sets LedgerId field to given value.

### HasLedgerId

`func (o *PaymentAccountLedgerEntry) HasLedgerId() bool`

HasLedgerId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *PaymentAccountLedgerEntry) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *PaymentAccountLedgerEntry) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *PaymentAccountLedgerEntry) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *PaymentAccountLedgerEntry) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentAccountLedgerEntry) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentAccountLedgerEntry) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentAccountLedgerEntry) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentAccountLedgerEntry) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentAccountLedgerEntry) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentAccountLedgerEntry) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentAccountLedgerEntry) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentAccountLedgerEntry) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentAccountLedgerEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentAccountLedgerEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentAccountLedgerEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentAccountLedgerEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBalance

`func (o *PaymentAccountLedgerEntry) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PaymentAccountLedgerEntry) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PaymentAccountLedgerEntry) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *PaymentAccountLedgerEntry) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCreateDate

`func (o *PaymentAccountLedgerEntry) GetCreateDate() int64`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *PaymentAccountLedgerEntry) GetCreateDateOk() (*int64, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *PaymentAccountLedgerEntry) SetCreateDate(v int64)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *PaymentAccountLedgerEntry) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *PaymentAccountLedgerEntry) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentAccountLedgerEntry) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentAccountLedgerEntry) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *PaymentAccountLedgerEntry) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetLedgerType

`func (o *PaymentAccountLedgerEntry) GetLedgerType() string`

GetLedgerType returns the LedgerType field if non-nil, zero value otherwise.

### GetLedgerTypeOk

`func (o *PaymentAccountLedgerEntry) GetLedgerTypeOk() (*string, bool)`

GetLedgerTypeOk returns a tuple with the LedgerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerType

`func (o *PaymentAccountLedgerEntry) SetLedgerType(v string)`

SetLedgerType sets LedgerType field to given value.

### HasLedgerType

`func (o *PaymentAccountLedgerEntry) HasLedgerType() bool`

HasLedgerType returns a boolean if a field has been set.

### GetReferenceType

`func (o *PaymentAccountLedgerEntry) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *PaymentAccountLedgerEntry) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *PaymentAccountLedgerEntry) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *PaymentAccountLedgerEntry) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetReferenceId

`func (o *PaymentAccountLedgerEntry) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentAccountLedgerEntry) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentAccountLedgerEntry) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PaymentAccountLedgerEntry) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *PaymentAccountLedgerEntry) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *PaymentAccountLedgerEntry) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetPaymentAdjustments

`func (o *PaymentAccountLedgerEntry) GetPaymentAdjustments() []PaymentAdjustment`

GetPaymentAdjustments returns the PaymentAdjustments field if non-nil, zero value otherwise.

### GetPaymentAdjustmentsOk

`func (o *PaymentAccountLedgerEntry) GetPaymentAdjustmentsOk() (*[]PaymentAdjustment, bool)`

GetPaymentAdjustmentsOk returns a tuple with the PaymentAdjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAdjustments

`func (o *PaymentAccountLedgerEntry) SetPaymentAdjustments(v []PaymentAdjustment)`

SetPaymentAdjustments sets PaymentAdjustments field to given value.

### HasPaymentAdjustments

`func (o *PaymentAccountLedgerEntry) HasPaymentAdjustments() bool`

HasPaymentAdjustments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


