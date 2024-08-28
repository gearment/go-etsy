# ShopReceiptTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **int64** | The unique numeric ID for a transaction. | [optional] 
**Title** | Pointer to **NullableString** | The title string of the [listing](/documentation/reference#tag/ShopListing) purchased in this transaction. | [optional] 
**Description** | Pointer to **NullableString** | The description string of the [listing](/documentation/reference#tag/ShopListing) purchased in this transaction. | [optional] 
**SellerUserId** | Pointer to **int64** | The numeric user ID for the seller in this transaction. | [optional] 
**BuyerUserId** | Pointer to **int64** | The numeric user ID for the buyer in this transaction. | [optional] 
**CreateTimestamp** | Pointer to **int64** | The transaction\\&#39;s creation date and time, in epoch seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The transaction\\&#39;s creation date and time, in epoch seconds. | [optional] 
**PaidTimestamp** | Pointer to **NullableInt64** | The transaction\\&#39;s paid date and time, in epoch seconds. | [optional] 
**ShippedTimestamp** | Pointer to **NullableInt64** | The transaction\\&#39;s shipping date and time, in epoch seconds. | [optional] 
**Quantity** | Pointer to **int64** | The numeric quantity of products purchased in this transaction. | [optional] 
**ListingImageId** | Pointer to **NullableInt64** | The numeric ID of the primary [listing image](/documentation/reference#tag/ShopListing-Image) for this transaction. | [optional] 
**ReceiptId** | Pointer to **int64** | The numeric ID for the [receipt](/documentation/reference#tag/Shop-Receipt) associated to this transaction. | [optional] 
**IsDigital** | Pointer to **bool** | When true, the transaction recorded the purchase of a digital listing. | [optional] 
**FileData** | Pointer to **string** | A string describing the files purchased in this transaction. | [optional] 
**ListingId** | Pointer to **NullableInt64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | [optional] 
**TransactionType** | Pointer to **string** | The type string for the transaction, usually \&quot;listing\&quot;. | [optional] 
**ProductId** | Pointer to **NullableInt64** | The numeric ID for a specific [product](/documentation/reference#tag/ShopListing-Product) purchased from a listing. | [optional] 
**Sku** | Pointer to **NullableString** | The SKU string for the product | [optional] 
**Price** | Pointer to [**Money**](Money.md) | A money object representing the price recorded the transaction. | [optional] 
**ShippingCost** | Pointer to [**Money**](Money.md) | A money object representing the shipping cost for this transaction. | [optional] 
**Variations** | Pointer to [**[]TransactionVariations**](TransactionVariations.md) | Array of variations and personalizations the buyer chose. | [optional] 
**ProductData** | Pointer to [**[]ListingPropertyValue**](ListingPropertyValue.md) | A list of property value entries for this product. Note: parenthesis characters (&#x60;(&#x60; and &#x60;)&#x60;) are not allowed. | [optional] 
**ShippingProfileId** | Pointer to **NullableInt64** | The ID of the shipping profile selected for this listing. | [optional] 
**MinProcessingDays** | Pointer to **NullableInt64** | The minimum number of days for processing the listing. | [optional] 
**MaxProcessingDays** | Pointer to **NullableInt64** | The maximum number of days for processing the listing. | [optional] 
**ShippingMethod** | Pointer to **NullableString** | Name of the selected shipping method. | [optional] 
**ShippingUpgrade** | Pointer to **NullableString** | The name of the shipping upgrade selected for this listing. Default value is null. | [optional] 
**ExpectedShipDate** | Pointer to **NullableInt64** | The date &amp; time of the expected ship date, in epoch seconds. | [optional] 
**BuyerCoupon** | Pointer to **float32** | The amount of the buyer coupon that was discounted in the shop&#39;s currency. | [optional] [default to 0]
**ShopCoupon** | Pointer to **float32** | The amount of the shop coupon that was discounted in the shop&#39;s currency. | [optional] [default to 0]

## Methods

### NewShopReceiptTransaction

`func NewShopReceiptTransaction() *ShopReceiptTransaction`

NewShopReceiptTransaction instantiates a new ShopReceiptTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopReceiptTransactionWithDefaults

`func NewShopReceiptTransactionWithDefaults() *ShopReceiptTransaction`

NewShopReceiptTransactionWithDefaults instantiates a new ShopReceiptTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *ShopReceiptTransaction) GetTransactionId() int64`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ShopReceiptTransaction) GetTransactionIdOk() (*int64, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ShopReceiptTransaction) SetTransactionId(v int64)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ShopReceiptTransaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTitle

`func (o *ShopReceiptTransaction) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShopReceiptTransaction) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShopReceiptTransaction) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ShopReceiptTransaction) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ShopReceiptTransaction) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ShopReceiptTransaction) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ShopReceiptTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShopReceiptTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShopReceiptTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShopReceiptTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShopReceiptTransaction) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShopReceiptTransaction) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSellerUserId

`func (o *ShopReceiptTransaction) GetSellerUserId() int64`

GetSellerUserId returns the SellerUserId field if non-nil, zero value otherwise.

### GetSellerUserIdOk

`func (o *ShopReceiptTransaction) GetSellerUserIdOk() (*int64, bool)`

GetSellerUserIdOk returns a tuple with the SellerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerUserId

`func (o *ShopReceiptTransaction) SetSellerUserId(v int64)`

SetSellerUserId sets SellerUserId field to given value.

### HasSellerUserId

`func (o *ShopReceiptTransaction) HasSellerUserId() bool`

HasSellerUserId returns a boolean if a field has been set.

### GetBuyerUserId

`func (o *ShopReceiptTransaction) GetBuyerUserId() int64`

GetBuyerUserId returns the BuyerUserId field if non-nil, zero value otherwise.

### GetBuyerUserIdOk

`func (o *ShopReceiptTransaction) GetBuyerUserIdOk() (*int64, bool)`

GetBuyerUserIdOk returns a tuple with the BuyerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUserId

`func (o *ShopReceiptTransaction) SetBuyerUserId(v int64)`

SetBuyerUserId sets BuyerUserId field to given value.

### HasBuyerUserId

`func (o *ShopReceiptTransaction) HasBuyerUserId() bool`

HasBuyerUserId returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *ShopReceiptTransaction) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *ShopReceiptTransaction) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *ShopReceiptTransaction) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *ShopReceiptTransaction) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ShopReceiptTransaction) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ShopReceiptTransaction) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ShopReceiptTransaction) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ShopReceiptTransaction) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetPaidTimestamp

`func (o *ShopReceiptTransaction) GetPaidTimestamp() int64`

GetPaidTimestamp returns the PaidTimestamp field if non-nil, zero value otherwise.

### GetPaidTimestampOk

`func (o *ShopReceiptTransaction) GetPaidTimestampOk() (*int64, bool)`

GetPaidTimestampOk returns a tuple with the PaidTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTimestamp

`func (o *ShopReceiptTransaction) SetPaidTimestamp(v int64)`

SetPaidTimestamp sets PaidTimestamp field to given value.

### HasPaidTimestamp

`func (o *ShopReceiptTransaction) HasPaidTimestamp() bool`

HasPaidTimestamp returns a boolean if a field has been set.

### SetPaidTimestampNil

`func (o *ShopReceiptTransaction) SetPaidTimestampNil(b bool)`

 SetPaidTimestampNil sets the value for PaidTimestamp to be an explicit nil

### UnsetPaidTimestamp
`func (o *ShopReceiptTransaction) UnsetPaidTimestamp()`

UnsetPaidTimestamp ensures that no value is present for PaidTimestamp, not even an explicit nil
### GetShippedTimestamp

`func (o *ShopReceiptTransaction) GetShippedTimestamp() int64`

GetShippedTimestamp returns the ShippedTimestamp field if non-nil, zero value otherwise.

### GetShippedTimestampOk

`func (o *ShopReceiptTransaction) GetShippedTimestampOk() (*int64, bool)`

GetShippedTimestampOk returns a tuple with the ShippedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedTimestamp

`func (o *ShopReceiptTransaction) SetShippedTimestamp(v int64)`

SetShippedTimestamp sets ShippedTimestamp field to given value.

### HasShippedTimestamp

`func (o *ShopReceiptTransaction) HasShippedTimestamp() bool`

HasShippedTimestamp returns a boolean if a field has been set.

### SetShippedTimestampNil

`func (o *ShopReceiptTransaction) SetShippedTimestampNil(b bool)`

 SetShippedTimestampNil sets the value for ShippedTimestamp to be an explicit nil

### UnsetShippedTimestamp
`func (o *ShopReceiptTransaction) UnsetShippedTimestamp()`

UnsetShippedTimestamp ensures that no value is present for ShippedTimestamp, not even an explicit nil
### GetQuantity

`func (o *ShopReceiptTransaction) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShopReceiptTransaction) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShopReceiptTransaction) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ShopReceiptTransaction) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetListingImageId

`func (o *ShopReceiptTransaction) GetListingImageId() int64`

GetListingImageId returns the ListingImageId field if non-nil, zero value otherwise.

### GetListingImageIdOk

`func (o *ShopReceiptTransaction) GetListingImageIdOk() (*int64, bool)`

GetListingImageIdOk returns a tuple with the ListingImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingImageId

`func (o *ShopReceiptTransaction) SetListingImageId(v int64)`

SetListingImageId sets ListingImageId field to given value.

### HasListingImageId

`func (o *ShopReceiptTransaction) HasListingImageId() bool`

HasListingImageId returns a boolean if a field has been set.

### SetListingImageIdNil

`func (o *ShopReceiptTransaction) SetListingImageIdNil(b bool)`

 SetListingImageIdNil sets the value for ListingImageId to be an explicit nil

### UnsetListingImageId
`func (o *ShopReceiptTransaction) UnsetListingImageId()`

UnsetListingImageId ensures that no value is present for ListingImageId, not even an explicit nil
### GetReceiptId

`func (o *ShopReceiptTransaction) GetReceiptId() int64`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ShopReceiptTransaction) GetReceiptIdOk() (*int64, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ShopReceiptTransaction) SetReceiptId(v int64)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ShopReceiptTransaction) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### GetIsDigital

`func (o *ShopReceiptTransaction) GetIsDigital() bool`

GetIsDigital returns the IsDigital field if non-nil, zero value otherwise.

### GetIsDigitalOk

`func (o *ShopReceiptTransaction) GetIsDigitalOk() (*bool, bool)`

GetIsDigitalOk returns a tuple with the IsDigital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDigital

`func (o *ShopReceiptTransaction) SetIsDigital(v bool)`

SetIsDigital sets IsDigital field to given value.

### HasIsDigital

`func (o *ShopReceiptTransaction) HasIsDigital() bool`

HasIsDigital returns a boolean if a field has been set.

### GetFileData

`func (o *ShopReceiptTransaction) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *ShopReceiptTransaction) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *ShopReceiptTransaction) SetFileData(v string)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *ShopReceiptTransaction) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### GetListingId

`func (o *ShopReceiptTransaction) GetListingId() int64`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ShopReceiptTransaction) GetListingIdOk() (*int64, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ShopReceiptTransaction) SetListingId(v int64)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *ShopReceiptTransaction) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *ShopReceiptTransaction) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *ShopReceiptTransaction) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetTransactionType

`func (o *ShopReceiptTransaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ShopReceiptTransaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ShopReceiptTransaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ShopReceiptTransaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetProductId

`func (o *ShopReceiptTransaction) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ShopReceiptTransaction) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ShopReceiptTransaction) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ShopReceiptTransaction) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *ShopReceiptTransaction) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *ShopReceiptTransaction) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetSku

`func (o *ShopReceiptTransaction) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ShopReceiptTransaction) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ShopReceiptTransaction) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ShopReceiptTransaction) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *ShopReceiptTransaction) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *ShopReceiptTransaction) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetPrice

`func (o *ShopReceiptTransaction) GetPrice() Money`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ShopReceiptTransaction) GetPriceOk() (*Money, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ShopReceiptTransaction) SetPrice(v Money)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ShopReceiptTransaction) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetShippingCost

`func (o *ShopReceiptTransaction) GetShippingCost() Money`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *ShopReceiptTransaction) GetShippingCostOk() (*Money, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *ShopReceiptTransaction) SetShippingCost(v Money)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *ShopReceiptTransaction) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetVariations

`func (o *ShopReceiptTransaction) GetVariations() []TransactionVariations`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *ShopReceiptTransaction) GetVariationsOk() (*[]TransactionVariations, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *ShopReceiptTransaction) SetVariations(v []TransactionVariations)`

SetVariations sets Variations field to given value.

### HasVariations

`func (o *ShopReceiptTransaction) HasVariations() bool`

HasVariations returns a boolean if a field has been set.

### GetProductData

`func (o *ShopReceiptTransaction) GetProductData() []ListingPropertyValue`

GetProductData returns the ProductData field if non-nil, zero value otherwise.

### GetProductDataOk

`func (o *ShopReceiptTransaction) GetProductDataOk() (*[]ListingPropertyValue, bool)`

GetProductDataOk returns a tuple with the ProductData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductData

`func (o *ShopReceiptTransaction) SetProductData(v []ListingPropertyValue)`

SetProductData sets ProductData field to given value.

### HasProductData

`func (o *ShopReceiptTransaction) HasProductData() bool`

HasProductData returns a boolean if a field has been set.

### GetShippingProfileId

`func (o *ShopReceiptTransaction) GetShippingProfileId() int64`

GetShippingProfileId returns the ShippingProfileId field if non-nil, zero value otherwise.

### GetShippingProfileIdOk

`func (o *ShopReceiptTransaction) GetShippingProfileIdOk() (*int64, bool)`

GetShippingProfileIdOk returns a tuple with the ShippingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileId

`func (o *ShopReceiptTransaction) SetShippingProfileId(v int64)`

SetShippingProfileId sets ShippingProfileId field to given value.

### HasShippingProfileId

`func (o *ShopReceiptTransaction) HasShippingProfileId() bool`

HasShippingProfileId returns a boolean if a field has been set.

### SetShippingProfileIdNil

`func (o *ShopReceiptTransaction) SetShippingProfileIdNil(b bool)`

 SetShippingProfileIdNil sets the value for ShippingProfileId to be an explicit nil

### UnsetShippingProfileId
`func (o *ShopReceiptTransaction) UnsetShippingProfileId()`

UnsetShippingProfileId ensures that no value is present for ShippingProfileId, not even an explicit nil
### GetMinProcessingDays

`func (o *ShopReceiptTransaction) GetMinProcessingDays() int64`

GetMinProcessingDays returns the MinProcessingDays field if non-nil, zero value otherwise.

### GetMinProcessingDaysOk

`func (o *ShopReceiptTransaction) GetMinProcessingDaysOk() (*int64, bool)`

GetMinProcessingDaysOk returns a tuple with the MinProcessingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProcessingDays

`func (o *ShopReceiptTransaction) SetMinProcessingDays(v int64)`

SetMinProcessingDays sets MinProcessingDays field to given value.

### HasMinProcessingDays

`func (o *ShopReceiptTransaction) HasMinProcessingDays() bool`

HasMinProcessingDays returns a boolean if a field has been set.

### SetMinProcessingDaysNil

`func (o *ShopReceiptTransaction) SetMinProcessingDaysNil(b bool)`

 SetMinProcessingDaysNil sets the value for MinProcessingDays to be an explicit nil

### UnsetMinProcessingDays
`func (o *ShopReceiptTransaction) UnsetMinProcessingDays()`

UnsetMinProcessingDays ensures that no value is present for MinProcessingDays, not even an explicit nil
### GetMaxProcessingDays

`func (o *ShopReceiptTransaction) GetMaxProcessingDays() int64`

GetMaxProcessingDays returns the MaxProcessingDays field if non-nil, zero value otherwise.

### GetMaxProcessingDaysOk

`func (o *ShopReceiptTransaction) GetMaxProcessingDaysOk() (*int64, bool)`

GetMaxProcessingDaysOk returns a tuple with the MaxProcessingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProcessingDays

`func (o *ShopReceiptTransaction) SetMaxProcessingDays(v int64)`

SetMaxProcessingDays sets MaxProcessingDays field to given value.

### HasMaxProcessingDays

`func (o *ShopReceiptTransaction) HasMaxProcessingDays() bool`

HasMaxProcessingDays returns a boolean if a field has been set.

### SetMaxProcessingDaysNil

`func (o *ShopReceiptTransaction) SetMaxProcessingDaysNil(b bool)`

 SetMaxProcessingDaysNil sets the value for MaxProcessingDays to be an explicit nil

### UnsetMaxProcessingDays
`func (o *ShopReceiptTransaction) UnsetMaxProcessingDays()`

UnsetMaxProcessingDays ensures that no value is present for MaxProcessingDays, not even an explicit nil
### GetShippingMethod

`func (o *ShopReceiptTransaction) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *ShopReceiptTransaction) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *ShopReceiptTransaction) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *ShopReceiptTransaction) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### SetShippingMethodNil

`func (o *ShopReceiptTransaction) SetShippingMethodNil(b bool)`

 SetShippingMethodNil sets the value for ShippingMethod to be an explicit nil

### UnsetShippingMethod
`func (o *ShopReceiptTransaction) UnsetShippingMethod()`

UnsetShippingMethod ensures that no value is present for ShippingMethod, not even an explicit nil
### GetShippingUpgrade

`func (o *ShopReceiptTransaction) GetShippingUpgrade() string`

GetShippingUpgrade returns the ShippingUpgrade field if non-nil, zero value otherwise.

### GetShippingUpgradeOk

`func (o *ShopReceiptTransaction) GetShippingUpgradeOk() (*string, bool)`

GetShippingUpgradeOk returns a tuple with the ShippingUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingUpgrade

`func (o *ShopReceiptTransaction) SetShippingUpgrade(v string)`

SetShippingUpgrade sets ShippingUpgrade field to given value.

### HasShippingUpgrade

`func (o *ShopReceiptTransaction) HasShippingUpgrade() bool`

HasShippingUpgrade returns a boolean if a field has been set.

### SetShippingUpgradeNil

`func (o *ShopReceiptTransaction) SetShippingUpgradeNil(b bool)`

 SetShippingUpgradeNil sets the value for ShippingUpgrade to be an explicit nil

### UnsetShippingUpgrade
`func (o *ShopReceiptTransaction) UnsetShippingUpgrade()`

UnsetShippingUpgrade ensures that no value is present for ShippingUpgrade, not even an explicit nil
### GetExpectedShipDate

`func (o *ShopReceiptTransaction) GetExpectedShipDate() int64`

GetExpectedShipDate returns the ExpectedShipDate field if non-nil, zero value otherwise.

### GetExpectedShipDateOk

`func (o *ShopReceiptTransaction) GetExpectedShipDateOk() (*int64, bool)`

GetExpectedShipDateOk returns a tuple with the ExpectedShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedShipDate

`func (o *ShopReceiptTransaction) SetExpectedShipDate(v int64)`

SetExpectedShipDate sets ExpectedShipDate field to given value.

### HasExpectedShipDate

`func (o *ShopReceiptTransaction) HasExpectedShipDate() bool`

HasExpectedShipDate returns a boolean if a field has been set.

### SetExpectedShipDateNil

`func (o *ShopReceiptTransaction) SetExpectedShipDateNil(b bool)`

 SetExpectedShipDateNil sets the value for ExpectedShipDate to be an explicit nil

### UnsetExpectedShipDate
`func (o *ShopReceiptTransaction) UnsetExpectedShipDate()`

UnsetExpectedShipDate ensures that no value is present for ExpectedShipDate, not even an explicit nil
### GetBuyerCoupon

`func (o *ShopReceiptTransaction) GetBuyerCoupon() float32`

GetBuyerCoupon returns the BuyerCoupon field if non-nil, zero value otherwise.

### GetBuyerCouponOk

`func (o *ShopReceiptTransaction) GetBuyerCouponOk() (*float32, bool)`

GetBuyerCouponOk returns a tuple with the BuyerCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCoupon

`func (o *ShopReceiptTransaction) SetBuyerCoupon(v float32)`

SetBuyerCoupon sets BuyerCoupon field to given value.

### HasBuyerCoupon

`func (o *ShopReceiptTransaction) HasBuyerCoupon() bool`

HasBuyerCoupon returns a boolean if a field has been set.

### GetShopCoupon

`func (o *ShopReceiptTransaction) GetShopCoupon() float32`

GetShopCoupon returns the ShopCoupon field if non-nil, zero value otherwise.

### GetShopCouponOk

`func (o *ShopReceiptTransaction) GetShopCouponOk() (*float32, bool)`

GetShopCouponOk returns a tuple with the ShopCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopCoupon

`func (o *ShopReceiptTransaction) SetShopCoupon(v float32)`

SetShopCoupon sets ShopCoupon field to given value.

### HasShopCoupon

`func (o *ShopReceiptTransaction) HasShopCoupon() bool`

HasShopCoupon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


