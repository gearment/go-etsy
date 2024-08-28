# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **int64** | A unique numeric ID for a payment to a specific Etsy [shop](/documentation/reference#tag/Shop). | [optional] 
**BuyerUserId** | Pointer to **int64** | The numeric ID for the [user](/documentation/reference#tag/User) who paid the purchase. | [optional] 
**ShopId** | Pointer to **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | [optional] 
**ReceiptId** | Pointer to **int64** | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction. | [optional] 
**AmountGross** | Pointer to [**Money**](Money.md) | An integer equal to gross amount of the order, in pennies, including shipping and taxes. | [optional] 
**AmountFees** | Pointer to [**Money**](Money.md) | An integer equal to the original card processing fee of the order in pennies. | [optional] 
**AmountNet** | Pointer to [**Money**](Money.md) | An integer equal to the payment value, in pennies, less fees (&#x60;amount_gross&#x60; - &#x60;amount_fees&#x60;). | [optional] 
**PostedGross** | Pointer to [**NullableMoney**](Money.md) | The total gross value of the payment posted once the purchase ships. This is equal to the &#x60;amount_gross&#x60; UNLESS the seller issues a refund prior to shipping. We consider \&quot;shipping\&quot; to be the event which \&quot;posts\&quot; to the ledger. Therefore, if the seller refunds first, we reduce the &#x60;amount_gross&#x60; first and post then that amount. The seller never sees the refunded amount in their ledger. This is equal to the \&quot;Credit\&quot; amount in the ledger entry. | [optional] 
**PostedFees** | Pointer to [**NullableMoney**](Money.md) | The total value of the fees posted once the purchase ships. Etsy refunds a proportional amount of the fees when a seller refunds a buyer. When the seller issues a refund prior to shipping, the posted amount is less then the original. | [optional] 
**PostedNet** | Pointer to [**NullableMoney**](Money.md) | The total value of the payment at the time of posting, less fees. (&#x60;posted_gross&#x60; - &#x60;posted_fees&#x60;) | [optional] 
**AdjustedGross** | Pointer to [**NullableMoney**](Money.md) | The gross payment amount after the seller refunds a payment, partially or fully. | [optional] 
**AdjustedFees** | Pointer to [**NullableMoney**](Money.md) | The new fee amount after a seller refunds a payment, partially or fully. | [optional] 
**AdjustedNet** | Pointer to [**NullableMoney**](Money.md) | The total value of the payment after refunds, less fees (&#x60;adjusted_gross&#x60; - &#x60;adjusted_fees&#x60;). | [optional] 
**Currency** | Pointer to **string** | The ISO (alphabetic) code string for the payment&#39;s currency. | [optional] 
**ShopCurrency** | Pointer to **NullableString** | The ISO (alphabetic) code for the shop&#39;s currency. The shop displays all prices in this currency by default. | [optional] 
**BuyerCurrency** | Pointer to **NullableString** | The currency string of the buyer. | [optional] 
**ShippingUserId** | Pointer to **NullableInt64** | The numeric ID of the user to which the seller ships the order. | [optional] 
**ShippingAddressId** | Pointer to **int64** | The numeric id identifying the shipping address. | [optional] 
**BillingAddressId** | Pointer to **int64** | The numeric ID identifying the billing address of the buyer. | [optional] 
**Status** | Pointer to **string** | A string indicating the current status of the payment, most commonly \&quot;settled\&quot; or \&quot;authed\&quot;. | [optional] 
**ShippedTimestamp** | Pointer to **NullableInt64** | The transaction\\&#39;s shipping date and time, in epoch seconds. | [optional] 
**CreateTimestamp** | Pointer to **int64** | The transaction\\&#39;s creation date and time, in epoch seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The transaction\\&#39;s creation date and time, in epoch seconds. | [optional] 
**UpdateTimestamp** | Pointer to **int64** | The date and time of the last change to the payment adjustment in epoch seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The date and time of the last change to the payment adjustment in epoch seconds. | [optional] 
**PaymentAdjustments** | Pointer to [**[]PaymentAdjustment**](PaymentAdjustment.md) | List of refund objects on an Etsy Payments transaction. All monetary amounts are in USD pennies unless otherwise specified. | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *Payment) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *Payment) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *Payment) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *Payment) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetBuyerUserId

`func (o *Payment) GetBuyerUserId() int64`

GetBuyerUserId returns the BuyerUserId field if non-nil, zero value otherwise.

### GetBuyerUserIdOk

`func (o *Payment) GetBuyerUserIdOk() (*int64, bool)`

GetBuyerUserIdOk returns a tuple with the BuyerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUserId

`func (o *Payment) SetBuyerUserId(v int64)`

SetBuyerUserId sets BuyerUserId field to given value.

### HasBuyerUserId

`func (o *Payment) HasBuyerUserId() bool`

HasBuyerUserId returns a boolean if a field has been set.

### GetShopId

`func (o *Payment) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Payment) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Payment) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Payment) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetReceiptId

`func (o *Payment) GetReceiptId() int64`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *Payment) GetReceiptIdOk() (*int64, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *Payment) SetReceiptId(v int64)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *Payment) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### GetAmountGross

`func (o *Payment) GetAmountGross() Money`

GetAmountGross returns the AmountGross field if non-nil, zero value otherwise.

### GetAmountGrossOk

`func (o *Payment) GetAmountGrossOk() (*Money, bool)`

GetAmountGrossOk returns a tuple with the AmountGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGross

`func (o *Payment) SetAmountGross(v Money)`

SetAmountGross sets AmountGross field to given value.

### HasAmountGross

`func (o *Payment) HasAmountGross() bool`

HasAmountGross returns a boolean if a field has been set.

### GetAmountFees

`func (o *Payment) GetAmountFees() Money`

GetAmountFees returns the AmountFees field if non-nil, zero value otherwise.

### GetAmountFeesOk

`func (o *Payment) GetAmountFeesOk() (*Money, bool)`

GetAmountFeesOk returns a tuple with the AmountFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFees

`func (o *Payment) SetAmountFees(v Money)`

SetAmountFees sets AmountFees field to given value.

### HasAmountFees

`func (o *Payment) HasAmountFees() bool`

HasAmountFees returns a boolean if a field has been set.

### GetAmountNet

`func (o *Payment) GetAmountNet() Money`

GetAmountNet returns the AmountNet field if non-nil, zero value otherwise.

### GetAmountNetOk

`func (o *Payment) GetAmountNetOk() (*Money, bool)`

GetAmountNetOk returns a tuple with the AmountNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountNet

`func (o *Payment) SetAmountNet(v Money)`

SetAmountNet sets AmountNet field to given value.

### HasAmountNet

`func (o *Payment) HasAmountNet() bool`

HasAmountNet returns a boolean if a field has been set.

### GetPostedGross

`func (o *Payment) GetPostedGross() Money`

GetPostedGross returns the PostedGross field if non-nil, zero value otherwise.

### GetPostedGrossOk

`func (o *Payment) GetPostedGrossOk() (*Money, bool)`

GetPostedGrossOk returns a tuple with the PostedGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedGross

`func (o *Payment) SetPostedGross(v Money)`

SetPostedGross sets PostedGross field to given value.

### HasPostedGross

`func (o *Payment) HasPostedGross() bool`

HasPostedGross returns a boolean if a field has been set.

### SetPostedGrossNil

`func (o *Payment) SetPostedGrossNil(b bool)`

 SetPostedGrossNil sets the value for PostedGross to be an explicit nil

### UnsetPostedGross
`func (o *Payment) UnsetPostedGross()`

UnsetPostedGross ensures that no value is present for PostedGross, not even an explicit nil
### GetPostedFees

`func (o *Payment) GetPostedFees() Money`

GetPostedFees returns the PostedFees field if non-nil, zero value otherwise.

### GetPostedFeesOk

`func (o *Payment) GetPostedFeesOk() (*Money, bool)`

GetPostedFeesOk returns a tuple with the PostedFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedFees

`func (o *Payment) SetPostedFees(v Money)`

SetPostedFees sets PostedFees field to given value.

### HasPostedFees

`func (o *Payment) HasPostedFees() bool`

HasPostedFees returns a boolean if a field has been set.

### SetPostedFeesNil

`func (o *Payment) SetPostedFeesNil(b bool)`

 SetPostedFeesNil sets the value for PostedFees to be an explicit nil

### UnsetPostedFees
`func (o *Payment) UnsetPostedFees()`

UnsetPostedFees ensures that no value is present for PostedFees, not even an explicit nil
### GetPostedNet

`func (o *Payment) GetPostedNet() Money`

GetPostedNet returns the PostedNet field if non-nil, zero value otherwise.

### GetPostedNetOk

`func (o *Payment) GetPostedNetOk() (*Money, bool)`

GetPostedNetOk returns a tuple with the PostedNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedNet

`func (o *Payment) SetPostedNet(v Money)`

SetPostedNet sets PostedNet field to given value.

### HasPostedNet

`func (o *Payment) HasPostedNet() bool`

HasPostedNet returns a boolean if a field has been set.

### SetPostedNetNil

`func (o *Payment) SetPostedNetNil(b bool)`

 SetPostedNetNil sets the value for PostedNet to be an explicit nil

### UnsetPostedNet
`func (o *Payment) UnsetPostedNet()`

UnsetPostedNet ensures that no value is present for PostedNet, not even an explicit nil
### GetAdjustedGross

`func (o *Payment) GetAdjustedGross() Money`

GetAdjustedGross returns the AdjustedGross field if non-nil, zero value otherwise.

### GetAdjustedGrossOk

`func (o *Payment) GetAdjustedGrossOk() (*Money, bool)`

GetAdjustedGrossOk returns a tuple with the AdjustedGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedGross

`func (o *Payment) SetAdjustedGross(v Money)`

SetAdjustedGross sets AdjustedGross field to given value.

### HasAdjustedGross

`func (o *Payment) HasAdjustedGross() bool`

HasAdjustedGross returns a boolean if a field has been set.

### SetAdjustedGrossNil

`func (o *Payment) SetAdjustedGrossNil(b bool)`

 SetAdjustedGrossNil sets the value for AdjustedGross to be an explicit nil

### UnsetAdjustedGross
`func (o *Payment) UnsetAdjustedGross()`

UnsetAdjustedGross ensures that no value is present for AdjustedGross, not even an explicit nil
### GetAdjustedFees

`func (o *Payment) GetAdjustedFees() Money`

GetAdjustedFees returns the AdjustedFees field if non-nil, zero value otherwise.

### GetAdjustedFeesOk

`func (o *Payment) GetAdjustedFeesOk() (*Money, bool)`

GetAdjustedFeesOk returns a tuple with the AdjustedFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedFees

`func (o *Payment) SetAdjustedFees(v Money)`

SetAdjustedFees sets AdjustedFees field to given value.

### HasAdjustedFees

`func (o *Payment) HasAdjustedFees() bool`

HasAdjustedFees returns a boolean if a field has been set.

### SetAdjustedFeesNil

`func (o *Payment) SetAdjustedFeesNil(b bool)`

 SetAdjustedFeesNil sets the value for AdjustedFees to be an explicit nil

### UnsetAdjustedFees
`func (o *Payment) UnsetAdjustedFees()`

UnsetAdjustedFees ensures that no value is present for AdjustedFees, not even an explicit nil
### GetAdjustedNet

`func (o *Payment) GetAdjustedNet() Money`

GetAdjustedNet returns the AdjustedNet field if non-nil, zero value otherwise.

### GetAdjustedNetOk

`func (o *Payment) GetAdjustedNetOk() (*Money, bool)`

GetAdjustedNetOk returns a tuple with the AdjustedNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedNet

`func (o *Payment) SetAdjustedNet(v Money)`

SetAdjustedNet sets AdjustedNet field to given value.

### HasAdjustedNet

`func (o *Payment) HasAdjustedNet() bool`

HasAdjustedNet returns a boolean if a field has been set.

### SetAdjustedNetNil

`func (o *Payment) SetAdjustedNetNil(b bool)`

 SetAdjustedNetNil sets the value for AdjustedNet to be an explicit nil

### UnsetAdjustedNet
`func (o *Payment) UnsetAdjustedNet()`

UnsetAdjustedNet ensures that no value is present for AdjustedNet, not even an explicit nil
### GetCurrency

`func (o *Payment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Payment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetShopCurrency

`func (o *Payment) GetShopCurrency() string`

GetShopCurrency returns the ShopCurrency field if non-nil, zero value otherwise.

### GetShopCurrencyOk

`func (o *Payment) GetShopCurrencyOk() (*string, bool)`

GetShopCurrencyOk returns a tuple with the ShopCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopCurrency

`func (o *Payment) SetShopCurrency(v string)`

SetShopCurrency sets ShopCurrency field to given value.

### HasShopCurrency

`func (o *Payment) HasShopCurrency() bool`

HasShopCurrency returns a boolean if a field has been set.

### SetShopCurrencyNil

`func (o *Payment) SetShopCurrencyNil(b bool)`

 SetShopCurrencyNil sets the value for ShopCurrency to be an explicit nil

### UnsetShopCurrency
`func (o *Payment) UnsetShopCurrency()`

UnsetShopCurrency ensures that no value is present for ShopCurrency, not even an explicit nil
### GetBuyerCurrency

`func (o *Payment) GetBuyerCurrency() string`

GetBuyerCurrency returns the BuyerCurrency field if non-nil, zero value otherwise.

### GetBuyerCurrencyOk

`func (o *Payment) GetBuyerCurrencyOk() (*string, bool)`

GetBuyerCurrencyOk returns a tuple with the BuyerCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCurrency

`func (o *Payment) SetBuyerCurrency(v string)`

SetBuyerCurrency sets BuyerCurrency field to given value.

### HasBuyerCurrency

`func (o *Payment) HasBuyerCurrency() bool`

HasBuyerCurrency returns a boolean if a field has been set.

### SetBuyerCurrencyNil

`func (o *Payment) SetBuyerCurrencyNil(b bool)`

 SetBuyerCurrencyNil sets the value for BuyerCurrency to be an explicit nil

### UnsetBuyerCurrency
`func (o *Payment) UnsetBuyerCurrency()`

UnsetBuyerCurrency ensures that no value is present for BuyerCurrency, not even an explicit nil
### GetShippingUserId

`func (o *Payment) GetShippingUserId() int64`

GetShippingUserId returns the ShippingUserId field if non-nil, zero value otherwise.

### GetShippingUserIdOk

`func (o *Payment) GetShippingUserIdOk() (*int64, bool)`

GetShippingUserIdOk returns a tuple with the ShippingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingUserId

`func (o *Payment) SetShippingUserId(v int64)`

SetShippingUserId sets ShippingUserId field to given value.

### HasShippingUserId

`func (o *Payment) HasShippingUserId() bool`

HasShippingUserId returns a boolean if a field has been set.

### SetShippingUserIdNil

`func (o *Payment) SetShippingUserIdNil(b bool)`

 SetShippingUserIdNil sets the value for ShippingUserId to be an explicit nil

### UnsetShippingUserId
`func (o *Payment) UnsetShippingUserId()`

UnsetShippingUserId ensures that no value is present for ShippingUserId, not even an explicit nil
### GetShippingAddressId

`func (o *Payment) GetShippingAddressId() int64`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *Payment) GetShippingAddressIdOk() (*int64, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *Payment) SetShippingAddressId(v int64)`

SetShippingAddressId sets ShippingAddressId field to given value.

### HasShippingAddressId

`func (o *Payment) HasShippingAddressId() bool`

HasShippingAddressId returns a boolean if a field has been set.

### GetBillingAddressId

`func (o *Payment) GetBillingAddressId() int64`

GetBillingAddressId returns the BillingAddressId field if non-nil, zero value otherwise.

### GetBillingAddressIdOk

`func (o *Payment) GetBillingAddressIdOk() (*int64, bool)`

GetBillingAddressIdOk returns a tuple with the BillingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressId

`func (o *Payment) SetBillingAddressId(v int64)`

SetBillingAddressId sets BillingAddressId field to given value.

### HasBillingAddressId

`func (o *Payment) HasBillingAddressId() bool`

HasBillingAddressId returns a boolean if a field has been set.

### GetStatus

`func (o *Payment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Payment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetShippedTimestamp

`func (o *Payment) GetShippedTimestamp() int64`

GetShippedTimestamp returns the ShippedTimestamp field if non-nil, zero value otherwise.

### GetShippedTimestampOk

`func (o *Payment) GetShippedTimestampOk() (*int64, bool)`

GetShippedTimestampOk returns a tuple with the ShippedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedTimestamp

`func (o *Payment) SetShippedTimestamp(v int64)`

SetShippedTimestamp sets ShippedTimestamp field to given value.

### HasShippedTimestamp

`func (o *Payment) HasShippedTimestamp() bool`

HasShippedTimestamp returns a boolean if a field has been set.

### SetShippedTimestampNil

`func (o *Payment) SetShippedTimestampNil(b bool)`

 SetShippedTimestampNil sets the value for ShippedTimestamp to be an explicit nil

### UnsetShippedTimestamp
`func (o *Payment) UnsetShippedTimestamp()`

UnsetShippedTimestamp ensures that no value is present for ShippedTimestamp, not even an explicit nil
### GetCreateTimestamp

`func (o *Payment) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *Payment) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *Payment) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *Payment) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Payment) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Payment) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Payment) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Payment) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *Payment) GetUpdateTimestamp() int64`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *Payment) GetUpdateTimestampOk() (*int64, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *Payment) SetUpdateTimestamp(v int64)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *Payment) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *Payment) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Payment) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Payment) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *Payment) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetPaymentAdjustments

`func (o *Payment) GetPaymentAdjustments() []PaymentAdjustment`

GetPaymentAdjustments returns the PaymentAdjustments field if non-nil, zero value otherwise.

### GetPaymentAdjustmentsOk

`func (o *Payment) GetPaymentAdjustmentsOk() (*[]PaymentAdjustment, bool)`

GetPaymentAdjustmentsOk returns a tuple with the PaymentAdjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAdjustments

`func (o *Payment) SetPaymentAdjustments(v []PaymentAdjustment)`

SetPaymentAdjustments sets PaymentAdjustments field to given value.

### HasPaymentAdjustments

`func (o *Payment) HasPaymentAdjustments() bool`

HasPaymentAdjustments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


