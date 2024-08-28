# ShopReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceiptId** | Pointer to **int64** | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction. | [optional] 
**ReceiptType** | Pointer to **int64** | The numeric value for the Etsy channel that serviced the purchase: 0 for Etsy.com, 1 for a Pattern shop. | [optional] 
**SellerUserId** | Pointer to **int64** | The numeric ID for the [user](/documentation/reference#tag/User) (seller) fulfilling the purchase. | [optional] 
**SellerEmail** | Pointer to **NullableString** | The email address string for the seller of the listing. | [optional] 
**BuyerUserId** | Pointer to **int64** | The numeric ID for the [user](/documentation/reference#tag/User) making the purchase. | [optional] 
**BuyerEmail** | Pointer to **NullableString** | The email address string for the buyer of the listing. It will be null if access hasn&#39;t been granted. Access is case-by-case and subject to approval. | [optional] 
**Name** | Pointer to **string** | The name string for the recipient in the shipping address. | [optional] 
**FirstLine** | Pointer to **NullableString** | The first address line string for the recipient in the shipping address. | [optional] 
**SecondLine** | Pointer to **NullableString** | The optional second address line string for the recipient in the shipping address. | [optional] 
**City** | Pointer to **NullableString** | The city string for the recipient in the shipping address. | [optional] 
**State** | Pointer to **NullableString** | The state string for the recipient in the shipping address. | [optional] 
**Zip** | Pointer to **NullableString** | The zip code string (not necessarily a number) for the recipient in the shipping address. | [optional] 
**Status** | Pointer to [**ShopReceiptStatus**](ShopReceiptStatus.md) |  | [optional] 
**FormattedAddress** | Pointer to **NullableString** | The formatted shipping address string for the recipient in the shipping address. | [optional] 
**CountryIso** | Pointer to **NullableString** | The ISO-3166 alpha-2 country code string for the recipient in the shipping address. | [optional] 
**PaymentMethod** | Pointer to **string** | The payment method string identifying purchaser&#39;s payment method, which must be one of: \\&#39;cc\\&#39; (credit card), \\&#39;paypal\\&#39;, \\&#39;check\\&#39;, \\&#39;mo\\&#39; (money order), \\&#39;bt\\&#39; (bank transfer), \\&#39;other\\&#39;, \\&#39;ideal\\&#39;, \\&#39;sofort\\&#39;, \\&#39;apple_pay\\&#39;, \\&#39;google\\&#39;, \\&#39;android_pay\\&#39;, \\&#39;google_pay\\&#39;, \\&#39;klarna\\&#39;, \\&#39;k_pay_in_4\\&#39; (klarna), \\&#39;k_pay_in_3\\&#39; (klarna), or \\&#39;k_financing\\&#39; (klarna). | [optional] 
**PaymentEmail** | Pointer to **NullableString** | The email address string for the email address to which to send payment confirmation | [optional] 
**MessageFromSeller** | Pointer to **NullableString** | An optional message string from the seller. | [optional] 
**MessageFromBuyer** | Pointer to **NullableString** | An optional message string from the buyer. | [optional] 
**MessageFromPayment** | Pointer to **NullableString** | The machine-generated acknowledgement string from the payment system. | [optional] 
**IsPaid** | Pointer to **bool** | When true, buyer paid for this purchase. | [optional] 
**IsShipped** | Pointer to **bool** | When true, seller shipped the products. | [optional] 
**CreateTimestamp** | Pointer to **int64** | The receipt\\&#39;s creation time, in epoch seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The receipt\\&#39;s creation time, in epoch seconds. | [optional] 
**UpdateTimestamp** | Pointer to **int64** | The time of the last update to the receipt, in epoch seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The time of the last update to the receipt, in epoch seconds. | [optional] 
**IsGift** | Pointer to **bool** | When true, the buyer indicated this purchase is a gift. | [optional] 
**GiftMessage** | Pointer to **string** | A gift message string the buyer requests delivered with the product. | [optional] 
**GiftSender** | Pointer to **string** | The name of the person who sent the gift. | [optional] 
**Grandtotal** | Pointer to [**Money**](Money.md) | A number equal to the total_price minus the coupon discount plus tax and shipping costs. | [optional] 
**Subtotal** | Pointer to [**Money**](Money.md) | A number equal to the total_price minus coupon discounts. Does not included tax or shipping costs. | [optional] 
**TotalPrice** | Pointer to [**Money**](Money.md) | A number equal to the sum of the individual listings\\&#39; (price * quantity). Does not included tax or shipping costs. | [optional] 
**TotalShippingCost** | Pointer to [**Money**](Money.md) | A number equal to the total shipping cost of the receipt. | [optional] 
**TotalTaxCost** | Pointer to [**Money**](Money.md) | The total sales tax of the receipt. | [optional] 
**TotalVatCost** | Pointer to [**Money**](Money.md) | A number equal to the total value-added tax (VAT) of the receipt. | [optional] 
**DiscountAmt** | Pointer to [**Money**](Money.md) | The numeric total discounted price for the receipt when using a discount (percent or fixed) coupon. Free shipping coupons are not included in this discount amount. | [optional] 
**GiftWrapPrice** | Pointer to [**Money**](Money.md) | The numeric price of gift wrap for this receipt. | [optional] 
**Shipments** | Pointer to [**[]ShopReceiptShipment**](ShopReceiptShipment.md) | A list of shipment statements for this receipt. | [optional] 
**Transactions** | Pointer to [**[]ShopReceiptTransaction**](ShopReceiptTransaction.md) | Array of transactions for the receipt. | [optional] 
**Refunds** | Pointer to [**[]ShopRefund**](ShopRefund.md) | Refunds for a given receipt. | [optional] 

## Methods

### NewShopReceipt

`func NewShopReceipt() *ShopReceipt`

NewShopReceipt instantiates a new ShopReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopReceiptWithDefaults

`func NewShopReceiptWithDefaults() *ShopReceipt`

NewShopReceiptWithDefaults instantiates a new ShopReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiptId

`func (o *ShopReceipt) GetReceiptId() int64`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ShopReceipt) GetReceiptIdOk() (*int64, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ShopReceipt) SetReceiptId(v int64)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ShopReceipt) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### GetReceiptType

`func (o *ShopReceipt) GetReceiptType() int64`

GetReceiptType returns the ReceiptType field if non-nil, zero value otherwise.

### GetReceiptTypeOk

`func (o *ShopReceipt) GetReceiptTypeOk() (*int64, bool)`

GetReceiptTypeOk returns a tuple with the ReceiptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptType

`func (o *ShopReceipt) SetReceiptType(v int64)`

SetReceiptType sets ReceiptType field to given value.

### HasReceiptType

`func (o *ShopReceipt) HasReceiptType() bool`

HasReceiptType returns a boolean if a field has been set.

### GetSellerUserId

`func (o *ShopReceipt) GetSellerUserId() int64`

GetSellerUserId returns the SellerUserId field if non-nil, zero value otherwise.

### GetSellerUserIdOk

`func (o *ShopReceipt) GetSellerUserIdOk() (*int64, bool)`

GetSellerUserIdOk returns a tuple with the SellerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerUserId

`func (o *ShopReceipt) SetSellerUserId(v int64)`

SetSellerUserId sets SellerUserId field to given value.

### HasSellerUserId

`func (o *ShopReceipt) HasSellerUserId() bool`

HasSellerUserId returns a boolean if a field has been set.

### GetSellerEmail

`func (o *ShopReceipt) GetSellerEmail() string`

GetSellerEmail returns the SellerEmail field if non-nil, zero value otherwise.

### GetSellerEmailOk

`func (o *ShopReceipt) GetSellerEmailOk() (*string, bool)`

GetSellerEmailOk returns a tuple with the SellerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerEmail

`func (o *ShopReceipt) SetSellerEmail(v string)`

SetSellerEmail sets SellerEmail field to given value.

### HasSellerEmail

`func (o *ShopReceipt) HasSellerEmail() bool`

HasSellerEmail returns a boolean if a field has been set.

### SetSellerEmailNil

`func (o *ShopReceipt) SetSellerEmailNil(b bool)`

 SetSellerEmailNil sets the value for SellerEmail to be an explicit nil

### UnsetSellerEmail
`func (o *ShopReceipt) UnsetSellerEmail()`

UnsetSellerEmail ensures that no value is present for SellerEmail, not even an explicit nil
### GetBuyerUserId

`func (o *ShopReceipt) GetBuyerUserId() int64`

GetBuyerUserId returns the BuyerUserId field if non-nil, zero value otherwise.

### GetBuyerUserIdOk

`func (o *ShopReceipt) GetBuyerUserIdOk() (*int64, bool)`

GetBuyerUserIdOk returns a tuple with the BuyerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUserId

`func (o *ShopReceipt) SetBuyerUserId(v int64)`

SetBuyerUserId sets BuyerUserId field to given value.

### HasBuyerUserId

`func (o *ShopReceipt) HasBuyerUserId() bool`

HasBuyerUserId returns a boolean if a field has been set.

### GetBuyerEmail

`func (o *ShopReceipt) GetBuyerEmail() string`

GetBuyerEmail returns the BuyerEmail field if non-nil, zero value otherwise.

### GetBuyerEmailOk

`func (o *ShopReceipt) GetBuyerEmailOk() (*string, bool)`

GetBuyerEmailOk returns a tuple with the BuyerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerEmail

`func (o *ShopReceipt) SetBuyerEmail(v string)`

SetBuyerEmail sets BuyerEmail field to given value.

### HasBuyerEmail

`func (o *ShopReceipt) HasBuyerEmail() bool`

HasBuyerEmail returns a boolean if a field has been set.

### SetBuyerEmailNil

`func (o *ShopReceipt) SetBuyerEmailNil(b bool)`

 SetBuyerEmailNil sets the value for BuyerEmail to be an explicit nil

### UnsetBuyerEmail
`func (o *ShopReceipt) UnsetBuyerEmail()`

UnsetBuyerEmail ensures that no value is present for BuyerEmail, not even an explicit nil
### GetName

`func (o *ShopReceipt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShopReceipt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShopReceipt) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShopReceipt) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFirstLine

`func (o *ShopReceipt) GetFirstLine() string`

GetFirstLine returns the FirstLine field if non-nil, zero value otherwise.

### GetFirstLineOk

`func (o *ShopReceipt) GetFirstLineOk() (*string, bool)`

GetFirstLineOk returns a tuple with the FirstLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLine

`func (o *ShopReceipt) SetFirstLine(v string)`

SetFirstLine sets FirstLine field to given value.

### HasFirstLine

`func (o *ShopReceipt) HasFirstLine() bool`

HasFirstLine returns a boolean if a field has been set.

### SetFirstLineNil

`func (o *ShopReceipt) SetFirstLineNil(b bool)`

 SetFirstLineNil sets the value for FirstLine to be an explicit nil

### UnsetFirstLine
`func (o *ShopReceipt) UnsetFirstLine()`

UnsetFirstLine ensures that no value is present for FirstLine, not even an explicit nil
### GetSecondLine

`func (o *ShopReceipt) GetSecondLine() string`

GetSecondLine returns the SecondLine field if non-nil, zero value otherwise.

### GetSecondLineOk

`func (o *ShopReceipt) GetSecondLineOk() (*string, bool)`

GetSecondLineOk returns a tuple with the SecondLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondLine

`func (o *ShopReceipt) SetSecondLine(v string)`

SetSecondLine sets SecondLine field to given value.

### HasSecondLine

`func (o *ShopReceipt) HasSecondLine() bool`

HasSecondLine returns a boolean if a field has been set.

### SetSecondLineNil

`func (o *ShopReceipt) SetSecondLineNil(b bool)`

 SetSecondLineNil sets the value for SecondLine to be an explicit nil

### UnsetSecondLine
`func (o *ShopReceipt) UnsetSecondLine()`

UnsetSecondLine ensures that no value is present for SecondLine, not even an explicit nil
### GetCity

`func (o *ShopReceipt) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ShopReceipt) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ShopReceipt) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ShopReceipt) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ShopReceipt) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ShopReceipt) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *ShopReceipt) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ShopReceipt) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ShopReceipt) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ShopReceipt) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ShopReceipt) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ShopReceipt) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetZip

`func (o *ShopReceipt) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ShopReceipt) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ShopReceipt) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ShopReceipt) HasZip() bool`

HasZip returns a boolean if a field has been set.

### SetZipNil

`func (o *ShopReceipt) SetZipNil(b bool)`

 SetZipNil sets the value for Zip to be an explicit nil

### UnsetZip
`func (o *ShopReceipt) UnsetZip()`

UnsetZip ensures that no value is present for Zip, not even an explicit nil
### GetStatus

`func (o *ShopReceipt) GetStatus() ShopReceiptStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShopReceipt) GetStatusOk() (*ShopReceiptStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShopReceipt) SetStatus(v ShopReceiptStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShopReceipt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFormattedAddress

`func (o *ShopReceipt) GetFormattedAddress() string`

GetFormattedAddress returns the FormattedAddress field if non-nil, zero value otherwise.

### GetFormattedAddressOk

`func (o *ShopReceipt) GetFormattedAddressOk() (*string, bool)`

GetFormattedAddressOk returns a tuple with the FormattedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAddress

`func (o *ShopReceipt) SetFormattedAddress(v string)`

SetFormattedAddress sets FormattedAddress field to given value.

### HasFormattedAddress

`func (o *ShopReceipt) HasFormattedAddress() bool`

HasFormattedAddress returns a boolean if a field has been set.

### SetFormattedAddressNil

`func (o *ShopReceipt) SetFormattedAddressNil(b bool)`

 SetFormattedAddressNil sets the value for FormattedAddress to be an explicit nil

### UnsetFormattedAddress
`func (o *ShopReceipt) UnsetFormattedAddress()`

UnsetFormattedAddress ensures that no value is present for FormattedAddress, not even an explicit nil
### GetCountryIso

`func (o *ShopReceipt) GetCountryIso() string`

GetCountryIso returns the CountryIso field if non-nil, zero value otherwise.

### GetCountryIsoOk

`func (o *ShopReceipt) GetCountryIsoOk() (*string, bool)`

GetCountryIsoOk returns a tuple with the CountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso

`func (o *ShopReceipt) SetCountryIso(v string)`

SetCountryIso sets CountryIso field to given value.

### HasCountryIso

`func (o *ShopReceipt) HasCountryIso() bool`

HasCountryIso returns a boolean if a field has been set.

### SetCountryIsoNil

`func (o *ShopReceipt) SetCountryIsoNil(b bool)`

 SetCountryIsoNil sets the value for CountryIso to be an explicit nil

### UnsetCountryIso
`func (o *ShopReceipt) UnsetCountryIso()`

UnsetCountryIso ensures that no value is present for CountryIso, not even an explicit nil
### GetPaymentMethod

`func (o *ShopReceipt) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ShopReceipt) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ShopReceipt) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ShopReceipt) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentEmail

`func (o *ShopReceipt) GetPaymentEmail() string`

GetPaymentEmail returns the PaymentEmail field if non-nil, zero value otherwise.

### GetPaymentEmailOk

`func (o *ShopReceipt) GetPaymentEmailOk() (*string, bool)`

GetPaymentEmailOk returns a tuple with the PaymentEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentEmail

`func (o *ShopReceipt) SetPaymentEmail(v string)`

SetPaymentEmail sets PaymentEmail field to given value.

### HasPaymentEmail

`func (o *ShopReceipt) HasPaymentEmail() bool`

HasPaymentEmail returns a boolean if a field has been set.

### SetPaymentEmailNil

`func (o *ShopReceipt) SetPaymentEmailNil(b bool)`

 SetPaymentEmailNil sets the value for PaymentEmail to be an explicit nil

### UnsetPaymentEmail
`func (o *ShopReceipt) UnsetPaymentEmail()`

UnsetPaymentEmail ensures that no value is present for PaymentEmail, not even an explicit nil
### GetMessageFromSeller

`func (o *ShopReceipt) GetMessageFromSeller() string`

GetMessageFromSeller returns the MessageFromSeller field if non-nil, zero value otherwise.

### GetMessageFromSellerOk

`func (o *ShopReceipt) GetMessageFromSellerOk() (*string, bool)`

GetMessageFromSellerOk returns a tuple with the MessageFromSeller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFromSeller

`func (o *ShopReceipt) SetMessageFromSeller(v string)`

SetMessageFromSeller sets MessageFromSeller field to given value.

### HasMessageFromSeller

`func (o *ShopReceipt) HasMessageFromSeller() bool`

HasMessageFromSeller returns a boolean if a field has been set.

### SetMessageFromSellerNil

`func (o *ShopReceipt) SetMessageFromSellerNil(b bool)`

 SetMessageFromSellerNil sets the value for MessageFromSeller to be an explicit nil

### UnsetMessageFromSeller
`func (o *ShopReceipt) UnsetMessageFromSeller()`

UnsetMessageFromSeller ensures that no value is present for MessageFromSeller, not even an explicit nil
### GetMessageFromBuyer

`func (o *ShopReceipt) GetMessageFromBuyer() string`

GetMessageFromBuyer returns the MessageFromBuyer field if non-nil, zero value otherwise.

### GetMessageFromBuyerOk

`func (o *ShopReceipt) GetMessageFromBuyerOk() (*string, bool)`

GetMessageFromBuyerOk returns a tuple with the MessageFromBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFromBuyer

`func (o *ShopReceipt) SetMessageFromBuyer(v string)`

SetMessageFromBuyer sets MessageFromBuyer field to given value.

### HasMessageFromBuyer

`func (o *ShopReceipt) HasMessageFromBuyer() bool`

HasMessageFromBuyer returns a boolean if a field has been set.

### SetMessageFromBuyerNil

`func (o *ShopReceipt) SetMessageFromBuyerNil(b bool)`

 SetMessageFromBuyerNil sets the value for MessageFromBuyer to be an explicit nil

### UnsetMessageFromBuyer
`func (o *ShopReceipt) UnsetMessageFromBuyer()`

UnsetMessageFromBuyer ensures that no value is present for MessageFromBuyer, not even an explicit nil
### GetMessageFromPayment

`func (o *ShopReceipt) GetMessageFromPayment() string`

GetMessageFromPayment returns the MessageFromPayment field if non-nil, zero value otherwise.

### GetMessageFromPaymentOk

`func (o *ShopReceipt) GetMessageFromPaymentOk() (*string, bool)`

GetMessageFromPaymentOk returns a tuple with the MessageFromPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFromPayment

`func (o *ShopReceipt) SetMessageFromPayment(v string)`

SetMessageFromPayment sets MessageFromPayment field to given value.

### HasMessageFromPayment

`func (o *ShopReceipt) HasMessageFromPayment() bool`

HasMessageFromPayment returns a boolean if a field has been set.

### SetMessageFromPaymentNil

`func (o *ShopReceipt) SetMessageFromPaymentNil(b bool)`

 SetMessageFromPaymentNil sets the value for MessageFromPayment to be an explicit nil

### UnsetMessageFromPayment
`func (o *ShopReceipt) UnsetMessageFromPayment()`

UnsetMessageFromPayment ensures that no value is present for MessageFromPayment, not even an explicit nil
### GetIsPaid

`func (o *ShopReceipt) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *ShopReceipt) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *ShopReceipt) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *ShopReceipt) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetIsShipped

`func (o *ShopReceipt) GetIsShipped() bool`

GetIsShipped returns the IsShipped field if non-nil, zero value otherwise.

### GetIsShippedOk

`func (o *ShopReceipt) GetIsShippedOk() (*bool, bool)`

GetIsShippedOk returns a tuple with the IsShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShipped

`func (o *ShopReceipt) SetIsShipped(v bool)`

SetIsShipped sets IsShipped field to given value.

### HasIsShipped

`func (o *ShopReceipt) HasIsShipped() bool`

HasIsShipped returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *ShopReceipt) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *ShopReceipt) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *ShopReceipt) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *ShopReceipt) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ShopReceipt) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ShopReceipt) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ShopReceipt) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ShopReceipt) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *ShopReceipt) GetUpdateTimestamp() int64`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *ShopReceipt) GetUpdateTimestampOk() (*int64, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *ShopReceipt) SetUpdateTimestamp(v int64)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *ShopReceipt) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ShopReceipt) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ShopReceipt) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ShopReceipt) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *ShopReceipt) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetIsGift

`func (o *ShopReceipt) GetIsGift() bool`

GetIsGift returns the IsGift field if non-nil, zero value otherwise.

### GetIsGiftOk

`func (o *ShopReceipt) GetIsGiftOk() (*bool, bool)`

GetIsGiftOk returns a tuple with the IsGift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGift

`func (o *ShopReceipt) SetIsGift(v bool)`

SetIsGift sets IsGift field to given value.

### HasIsGift

`func (o *ShopReceipt) HasIsGift() bool`

HasIsGift returns a boolean if a field has been set.

### GetGiftMessage

`func (o *ShopReceipt) GetGiftMessage() string`

GetGiftMessage returns the GiftMessage field if non-nil, zero value otherwise.

### GetGiftMessageOk

`func (o *ShopReceipt) GetGiftMessageOk() (*string, bool)`

GetGiftMessageOk returns a tuple with the GiftMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftMessage

`func (o *ShopReceipt) SetGiftMessage(v string)`

SetGiftMessage sets GiftMessage field to given value.

### HasGiftMessage

`func (o *ShopReceipt) HasGiftMessage() bool`

HasGiftMessage returns a boolean if a field has been set.

### GetGiftSender

`func (o *ShopReceipt) GetGiftSender() string`

GetGiftSender returns the GiftSender field if non-nil, zero value otherwise.

### GetGiftSenderOk

`func (o *ShopReceipt) GetGiftSenderOk() (*string, bool)`

GetGiftSenderOk returns a tuple with the GiftSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftSender

`func (o *ShopReceipt) SetGiftSender(v string)`

SetGiftSender sets GiftSender field to given value.

### HasGiftSender

`func (o *ShopReceipt) HasGiftSender() bool`

HasGiftSender returns a boolean if a field has been set.

### GetGrandtotal

`func (o *ShopReceipt) GetGrandtotal() Money`

GetGrandtotal returns the Grandtotal field if non-nil, zero value otherwise.

### GetGrandtotalOk

`func (o *ShopReceipt) GetGrandtotalOk() (*Money, bool)`

GetGrandtotalOk returns a tuple with the Grandtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandtotal

`func (o *ShopReceipt) SetGrandtotal(v Money)`

SetGrandtotal sets Grandtotal field to given value.

### HasGrandtotal

`func (o *ShopReceipt) HasGrandtotal() bool`

HasGrandtotal returns a boolean if a field has been set.

### GetSubtotal

`func (o *ShopReceipt) GetSubtotal() Money`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *ShopReceipt) GetSubtotalOk() (*Money, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *ShopReceipt) SetSubtotal(v Money)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *ShopReceipt) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTotalPrice

`func (o *ShopReceipt) GetTotalPrice() Money`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *ShopReceipt) GetTotalPriceOk() (*Money, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *ShopReceipt) SetTotalPrice(v Money)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *ShopReceipt) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetTotalShippingCost

`func (o *ShopReceipt) GetTotalShippingCost() Money`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *ShopReceipt) GetTotalShippingCostOk() (*Money, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *ShopReceipt) SetTotalShippingCost(v Money)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *ShopReceipt) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalTaxCost

`func (o *ShopReceipt) GetTotalTaxCost() Money`

GetTotalTaxCost returns the TotalTaxCost field if non-nil, zero value otherwise.

### GetTotalTaxCostOk

`func (o *ShopReceipt) GetTotalTaxCostOk() (*Money, bool)`

GetTotalTaxCostOk returns a tuple with the TotalTaxCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxCost

`func (o *ShopReceipt) SetTotalTaxCost(v Money)`

SetTotalTaxCost sets TotalTaxCost field to given value.

### HasTotalTaxCost

`func (o *ShopReceipt) HasTotalTaxCost() bool`

HasTotalTaxCost returns a boolean if a field has been set.

### GetTotalVatCost

`func (o *ShopReceipt) GetTotalVatCost() Money`

GetTotalVatCost returns the TotalVatCost field if non-nil, zero value otherwise.

### GetTotalVatCostOk

`func (o *ShopReceipt) GetTotalVatCostOk() (*Money, bool)`

GetTotalVatCostOk returns a tuple with the TotalVatCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVatCost

`func (o *ShopReceipt) SetTotalVatCost(v Money)`

SetTotalVatCost sets TotalVatCost field to given value.

### HasTotalVatCost

`func (o *ShopReceipt) HasTotalVatCost() bool`

HasTotalVatCost returns a boolean if a field has been set.

### GetDiscountAmt

`func (o *ShopReceipt) GetDiscountAmt() Money`

GetDiscountAmt returns the DiscountAmt field if non-nil, zero value otherwise.

### GetDiscountAmtOk

`func (o *ShopReceipt) GetDiscountAmtOk() (*Money, bool)`

GetDiscountAmtOk returns a tuple with the DiscountAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmt

`func (o *ShopReceipt) SetDiscountAmt(v Money)`

SetDiscountAmt sets DiscountAmt field to given value.

### HasDiscountAmt

`func (o *ShopReceipt) HasDiscountAmt() bool`

HasDiscountAmt returns a boolean if a field has been set.

### GetGiftWrapPrice

`func (o *ShopReceipt) GetGiftWrapPrice() Money`

GetGiftWrapPrice returns the GiftWrapPrice field if non-nil, zero value otherwise.

### GetGiftWrapPriceOk

`func (o *ShopReceipt) GetGiftWrapPriceOk() (*Money, bool)`

GetGiftWrapPriceOk returns a tuple with the GiftWrapPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftWrapPrice

`func (o *ShopReceipt) SetGiftWrapPrice(v Money)`

SetGiftWrapPrice sets GiftWrapPrice field to given value.

### HasGiftWrapPrice

`func (o *ShopReceipt) HasGiftWrapPrice() bool`

HasGiftWrapPrice returns a boolean if a field has been set.

### GetShipments

`func (o *ShopReceipt) GetShipments() []ShopReceiptShipment`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *ShopReceipt) GetShipmentsOk() (*[]ShopReceiptShipment, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *ShopReceipt) SetShipments(v []ShopReceiptShipment)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *ShopReceipt) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetTransactions

`func (o *ShopReceipt) GetTransactions() []ShopReceiptTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ShopReceipt) GetTransactionsOk() (*[]ShopReceiptTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ShopReceipt) SetTransactions(v []ShopReceiptTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ShopReceipt) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetRefunds

`func (o *ShopReceipt) GetRefunds() []ShopRefund`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *ShopReceipt) GetRefundsOk() (*[]ShopRefund, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *ShopReceipt) SetRefunds(v []ShopRefund)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *ShopReceipt) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


