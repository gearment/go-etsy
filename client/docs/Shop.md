# Shop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | Pointer to **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | [optional] 
**UserId** | Pointer to **int64** | The numeric user ID of the [user](/documentation/reference#tag/User) who owns this shop. | [optional] 
**ShopName** | Pointer to **string** | The shop&#39;s name string. | [optional] 
**CreateDate** | Pointer to **int64** | The date and time this shop was created, in epoch seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The date and time this shop was created, in epoch seconds. | [optional] 
**Title** | Pointer to **NullableString** | A brief heading string for the shop\\&#39;s main page. | [optional] 
**Announcement** | Pointer to **NullableString** | An announcement string to buyers that displays on the shop&#39;s homepage. | [optional] 
**CurrencyCode** | Pointer to **string** | The ISO (alphabetic) code for the shop&#39;s currency. The shop displays all prices in this currency by default. | [optional] 
**IsVacation** | Pointer to **bool** | When true, this shop is not accepting purchases. | [optional] 
**VacationMessage** | Pointer to **NullableString** | The shop&#39;s message string displayed when &#x60;is_vacation&#x60; is true. | [optional] 
**SaleMessage** | Pointer to **NullableString** | A message string sent to users who complete a purchase from this shop. | [optional] 
**DigitalSaleMessage** | Pointer to **NullableString** | A message string sent to users who purchase a digital item from this shop. | [optional] 
**UpdateDate** | Pointer to **int64** | The date and time of the last update to the shop, in epoch seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The date and time of the last update to the shop, in epoch seconds. | [optional] 
**ListingActiveCount** | Pointer to **int64** | The number of active listings in the shop. | [optional] 
**DigitalListingCount** | Pointer to **int64** | The number of digital listings in the shop. | [optional] 
**LoginName** | Pointer to **string** | The shop owner\\&#39;s login name string. | [optional] 
**AcceptsCustomRequests** | Pointer to **bool** | When true, the shop accepts customization requests. | [optional] 
**PolicyWelcome** | Pointer to **NullableString** | The shop&#39;s policy welcome string (may be blank). | [optional] 
**PolicyPayment** | Pointer to **NullableString** | The shop&#39;s payment policy string (may be blank). | [optional] 
**PolicyShipping** | Pointer to **NullableString** | The shop&#39;s shipping policy string (may be blank). | [optional] 
**PolicyRefunds** | Pointer to **NullableString** | The shop&#39;s refund policy string (may be blank). | [optional] 
**PolicyAdditional** | Pointer to **NullableString** | The shop&#39;s additional policies string (may be blank). | [optional] 
**PolicySellerInfo** | Pointer to **NullableString** | The shop&#39;s seller information string (may be blank). | [optional] 
**PolicyUpdateDate** | Pointer to **int64** | The date and time of the last update to the shop&#39;s policies, in epoch seconds. | [optional] 
**PolicyHasPrivateReceiptInfo** | Pointer to **bool** | When true, EU receipts display private info. | [optional] 
**HasUnstructuredPolicies** | Pointer to **bool** | When true, the shop displays additional unstructured policy fields. | [optional] 
**PolicyPrivacy** | Pointer to **NullableString** | The shop&#39;s privacy policy string (may be blank). | [optional] 
**VacationAutoreply** | Pointer to **NullableString** | The shop&#39;s automatic reply string displayed in new conversations when &#x60;is_vacation&#x60; is true. | [optional] 
**Url** | Pointer to **string** | The URL string for this shop. | [optional] 
**ImageUrl760x100** | Pointer to **NullableString** | The URL string for this shop&#39;s banner image. | [optional] 
**NumFavorers** | Pointer to **int64** | The number of users who marked this shop a favorite. | [optional] 
**Languages** | Pointer to **[]string** | A list of language strings for the shop&#39;s enrolled languages where the default shop language is the first element in the array. | [optional] 
**IconUrlFullxfull** | Pointer to **NullableString** | The URL string for this shop&#39;s icon image. | [optional] 
**IsUsingStructuredPolicies** | Pointer to **bool** | When true, the shop accepted using structured policies. | [optional] 
**HasOnboardedStructuredPolicies** | Pointer to **bool** | When true, the shop accepted OR declined after viewing structured policies onboarding. | [optional] 
**IncludeDisputeFormLink** | Pointer to **bool** | When true, this shop\\&#39;s policies include a link to an EU online dispute form. | [optional] 
**IsDirectCheckoutOnboarded** | Pointer to **bool** | (**DEPRECATED: Replaced by _is_etsy_payments_onboarded_.) When true, the shop has onboarded onto Etsy Payments. | [optional] 
**IsEtsyPaymentsOnboarded** | Pointer to **bool** | When true, the shop has onboarded onto Etsy Payments. | [optional] 
**IsCalculatedEligible** | Pointer to **bool** | When true, the shop is eligible for calculated shipping profiles. (Only available in the US and Canada) | [optional] 
**IsOptedInToBuyerPromise** | Pointer to **bool** | When true, the shop opted in to buyer promise. | [optional] 
**IsShopUsBased** | Pointer to **bool** | When true, the shop is based in the US. | [optional] 
**TransactionSoldCount** | Pointer to **int64** | The total number of sales ([transactions](/documentation/reference#tag/Shop-Receipt-Transactions)) for this shop. | [optional] 
**ShippingFromCountryIso** | Pointer to **NullableString** | The country iso the shop is shipping from. | [optional] 
**ShopLocationCountryIso** | Pointer to **NullableString** | The country iso where the shop is located. | [optional] 
**ReviewCount** | Pointer to **NullableInt64** | Number of reviews of shop listings in the past year. | [optional] 
**ReviewAverage** | Pointer to **NullableFloat32** | Average rating based on reviews of shop listings in the past year. | [optional] 

## Methods

### NewShop

`func NewShop() *Shop`

NewShop instantiates a new Shop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopWithDefaults

`func NewShopWithDefaults() *Shop`

NewShopWithDefaults instantiates a new Shop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *Shop) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Shop) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Shop) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Shop) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetUserId

`func (o *Shop) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Shop) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Shop) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Shop) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetShopName

`func (o *Shop) GetShopName() string`

GetShopName returns the ShopName field if non-nil, zero value otherwise.

### GetShopNameOk

`func (o *Shop) GetShopNameOk() (*string, bool)`

GetShopNameOk returns a tuple with the ShopName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopName

`func (o *Shop) SetShopName(v string)`

SetShopName sets ShopName field to given value.

### HasShopName

`func (o *Shop) HasShopName() bool`

HasShopName returns a boolean if a field has been set.

### GetCreateDate

`func (o *Shop) GetCreateDate() int64`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *Shop) GetCreateDateOk() (*int64, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *Shop) SetCreateDate(v int64)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *Shop) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Shop) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Shop) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Shop) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Shop) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *Shop) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Shop) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Shop) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Shop) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Shop) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Shop) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAnnouncement

`func (o *Shop) GetAnnouncement() string`

GetAnnouncement returns the Announcement field if non-nil, zero value otherwise.

### GetAnnouncementOk

`func (o *Shop) GetAnnouncementOk() (*string, bool)`

GetAnnouncementOk returns a tuple with the Announcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncement

`func (o *Shop) SetAnnouncement(v string)`

SetAnnouncement sets Announcement field to given value.

### HasAnnouncement

`func (o *Shop) HasAnnouncement() bool`

HasAnnouncement returns a boolean if a field has been set.

### SetAnnouncementNil

`func (o *Shop) SetAnnouncementNil(b bool)`

 SetAnnouncementNil sets the value for Announcement to be an explicit nil

### UnsetAnnouncement
`func (o *Shop) UnsetAnnouncement()`

UnsetAnnouncement ensures that no value is present for Announcement, not even an explicit nil
### GetCurrencyCode

`func (o *Shop) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Shop) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Shop) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *Shop) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetIsVacation

`func (o *Shop) GetIsVacation() bool`

GetIsVacation returns the IsVacation field if non-nil, zero value otherwise.

### GetIsVacationOk

`func (o *Shop) GetIsVacationOk() (*bool, bool)`

GetIsVacationOk returns a tuple with the IsVacation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVacation

`func (o *Shop) SetIsVacation(v bool)`

SetIsVacation sets IsVacation field to given value.

### HasIsVacation

`func (o *Shop) HasIsVacation() bool`

HasIsVacation returns a boolean if a field has been set.

### GetVacationMessage

`func (o *Shop) GetVacationMessage() string`

GetVacationMessage returns the VacationMessage field if non-nil, zero value otherwise.

### GetVacationMessageOk

`func (o *Shop) GetVacationMessageOk() (*string, bool)`

GetVacationMessageOk returns a tuple with the VacationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationMessage

`func (o *Shop) SetVacationMessage(v string)`

SetVacationMessage sets VacationMessage field to given value.

### HasVacationMessage

`func (o *Shop) HasVacationMessage() bool`

HasVacationMessage returns a boolean if a field has been set.

### SetVacationMessageNil

`func (o *Shop) SetVacationMessageNil(b bool)`

 SetVacationMessageNil sets the value for VacationMessage to be an explicit nil

### UnsetVacationMessage
`func (o *Shop) UnsetVacationMessage()`

UnsetVacationMessage ensures that no value is present for VacationMessage, not even an explicit nil
### GetSaleMessage

`func (o *Shop) GetSaleMessage() string`

GetSaleMessage returns the SaleMessage field if non-nil, zero value otherwise.

### GetSaleMessageOk

`func (o *Shop) GetSaleMessageOk() (*string, bool)`

GetSaleMessageOk returns a tuple with the SaleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleMessage

`func (o *Shop) SetSaleMessage(v string)`

SetSaleMessage sets SaleMessage field to given value.

### HasSaleMessage

`func (o *Shop) HasSaleMessage() bool`

HasSaleMessage returns a boolean if a field has been set.

### SetSaleMessageNil

`func (o *Shop) SetSaleMessageNil(b bool)`

 SetSaleMessageNil sets the value for SaleMessage to be an explicit nil

### UnsetSaleMessage
`func (o *Shop) UnsetSaleMessage()`

UnsetSaleMessage ensures that no value is present for SaleMessage, not even an explicit nil
### GetDigitalSaleMessage

`func (o *Shop) GetDigitalSaleMessage() string`

GetDigitalSaleMessage returns the DigitalSaleMessage field if non-nil, zero value otherwise.

### GetDigitalSaleMessageOk

`func (o *Shop) GetDigitalSaleMessageOk() (*string, bool)`

GetDigitalSaleMessageOk returns a tuple with the DigitalSaleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSaleMessage

`func (o *Shop) SetDigitalSaleMessage(v string)`

SetDigitalSaleMessage sets DigitalSaleMessage field to given value.

### HasDigitalSaleMessage

`func (o *Shop) HasDigitalSaleMessage() bool`

HasDigitalSaleMessage returns a boolean if a field has been set.

### SetDigitalSaleMessageNil

`func (o *Shop) SetDigitalSaleMessageNil(b bool)`

 SetDigitalSaleMessageNil sets the value for DigitalSaleMessage to be an explicit nil

### UnsetDigitalSaleMessage
`func (o *Shop) UnsetDigitalSaleMessage()`

UnsetDigitalSaleMessage ensures that no value is present for DigitalSaleMessage, not even an explicit nil
### GetUpdateDate

`func (o *Shop) GetUpdateDate() int64`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *Shop) GetUpdateDateOk() (*int64, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *Shop) SetUpdateDate(v int64)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *Shop) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *Shop) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Shop) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Shop) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *Shop) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetListingActiveCount

`func (o *Shop) GetListingActiveCount() int64`

GetListingActiveCount returns the ListingActiveCount field if non-nil, zero value otherwise.

### GetListingActiveCountOk

`func (o *Shop) GetListingActiveCountOk() (*int64, bool)`

GetListingActiveCountOk returns a tuple with the ListingActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingActiveCount

`func (o *Shop) SetListingActiveCount(v int64)`

SetListingActiveCount sets ListingActiveCount field to given value.

### HasListingActiveCount

`func (o *Shop) HasListingActiveCount() bool`

HasListingActiveCount returns a boolean if a field has been set.

### GetDigitalListingCount

`func (o *Shop) GetDigitalListingCount() int64`

GetDigitalListingCount returns the DigitalListingCount field if non-nil, zero value otherwise.

### GetDigitalListingCountOk

`func (o *Shop) GetDigitalListingCountOk() (*int64, bool)`

GetDigitalListingCountOk returns a tuple with the DigitalListingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalListingCount

`func (o *Shop) SetDigitalListingCount(v int64)`

SetDigitalListingCount sets DigitalListingCount field to given value.

### HasDigitalListingCount

`func (o *Shop) HasDigitalListingCount() bool`

HasDigitalListingCount returns a boolean if a field has been set.

### GetLoginName

`func (o *Shop) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *Shop) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *Shop) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *Shop) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetAcceptsCustomRequests

`func (o *Shop) GetAcceptsCustomRequests() bool`

GetAcceptsCustomRequests returns the AcceptsCustomRequests field if non-nil, zero value otherwise.

### GetAcceptsCustomRequestsOk

`func (o *Shop) GetAcceptsCustomRequestsOk() (*bool, bool)`

GetAcceptsCustomRequestsOk returns a tuple with the AcceptsCustomRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptsCustomRequests

`func (o *Shop) SetAcceptsCustomRequests(v bool)`

SetAcceptsCustomRequests sets AcceptsCustomRequests field to given value.

### HasAcceptsCustomRequests

`func (o *Shop) HasAcceptsCustomRequests() bool`

HasAcceptsCustomRequests returns a boolean if a field has been set.

### GetPolicyWelcome

`func (o *Shop) GetPolicyWelcome() string`

GetPolicyWelcome returns the PolicyWelcome field if non-nil, zero value otherwise.

### GetPolicyWelcomeOk

`func (o *Shop) GetPolicyWelcomeOk() (*string, bool)`

GetPolicyWelcomeOk returns a tuple with the PolicyWelcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWelcome

`func (o *Shop) SetPolicyWelcome(v string)`

SetPolicyWelcome sets PolicyWelcome field to given value.

### HasPolicyWelcome

`func (o *Shop) HasPolicyWelcome() bool`

HasPolicyWelcome returns a boolean if a field has been set.

### SetPolicyWelcomeNil

`func (o *Shop) SetPolicyWelcomeNil(b bool)`

 SetPolicyWelcomeNil sets the value for PolicyWelcome to be an explicit nil

### UnsetPolicyWelcome
`func (o *Shop) UnsetPolicyWelcome()`

UnsetPolicyWelcome ensures that no value is present for PolicyWelcome, not even an explicit nil
### GetPolicyPayment

`func (o *Shop) GetPolicyPayment() string`

GetPolicyPayment returns the PolicyPayment field if non-nil, zero value otherwise.

### GetPolicyPaymentOk

`func (o *Shop) GetPolicyPaymentOk() (*string, bool)`

GetPolicyPaymentOk returns a tuple with the PolicyPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPayment

`func (o *Shop) SetPolicyPayment(v string)`

SetPolicyPayment sets PolicyPayment field to given value.

### HasPolicyPayment

`func (o *Shop) HasPolicyPayment() bool`

HasPolicyPayment returns a boolean if a field has been set.

### SetPolicyPaymentNil

`func (o *Shop) SetPolicyPaymentNil(b bool)`

 SetPolicyPaymentNil sets the value for PolicyPayment to be an explicit nil

### UnsetPolicyPayment
`func (o *Shop) UnsetPolicyPayment()`

UnsetPolicyPayment ensures that no value is present for PolicyPayment, not even an explicit nil
### GetPolicyShipping

`func (o *Shop) GetPolicyShipping() string`

GetPolicyShipping returns the PolicyShipping field if non-nil, zero value otherwise.

### GetPolicyShippingOk

`func (o *Shop) GetPolicyShippingOk() (*string, bool)`

GetPolicyShippingOk returns a tuple with the PolicyShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyShipping

`func (o *Shop) SetPolicyShipping(v string)`

SetPolicyShipping sets PolicyShipping field to given value.

### HasPolicyShipping

`func (o *Shop) HasPolicyShipping() bool`

HasPolicyShipping returns a boolean if a field has been set.

### SetPolicyShippingNil

`func (o *Shop) SetPolicyShippingNil(b bool)`

 SetPolicyShippingNil sets the value for PolicyShipping to be an explicit nil

### UnsetPolicyShipping
`func (o *Shop) UnsetPolicyShipping()`

UnsetPolicyShipping ensures that no value is present for PolicyShipping, not even an explicit nil
### GetPolicyRefunds

`func (o *Shop) GetPolicyRefunds() string`

GetPolicyRefunds returns the PolicyRefunds field if non-nil, zero value otherwise.

### GetPolicyRefundsOk

`func (o *Shop) GetPolicyRefundsOk() (*string, bool)`

GetPolicyRefundsOk returns a tuple with the PolicyRefunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRefunds

`func (o *Shop) SetPolicyRefunds(v string)`

SetPolicyRefunds sets PolicyRefunds field to given value.

### HasPolicyRefunds

`func (o *Shop) HasPolicyRefunds() bool`

HasPolicyRefunds returns a boolean if a field has been set.

### SetPolicyRefundsNil

`func (o *Shop) SetPolicyRefundsNil(b bool)`

 SetPolicyRefundsNil sets the value for PolicyRefunds to be an explicit nil

### UnsetPolicyRefunds
`func (o *Shop) UnsetPolicyRefunds()`

UnsetPolicyRefunds ensures that no value is present for PolicyRefunds, not even an explicit nil
### GetPolicyAdditional

`func (o *Shop) GetPolicyAdditional() string`

GetPolicyAdditional returns the PolicyAdditional field if non-nil, zero value otherwise.

### GetPolicyAdditionalOk

`func (o *Shop) GetPolicyAdditionalOk() (*string, bool)`

GetPolicyAdditionalOk returns a tuple with the PolicyAdditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyAdditional

`func (o *Shop) SetPolicyAdditional(v string)`

SetPolicyAdditional sets PolicyAdditional field to given value.

### HasPolicyAdditional

`func (o *Shop) HasPolicyAdditional() bool`

HasPolicyAdditional returns a boolean if a field has been set.

### SetPolicyAdditionalNil

`func (o *Shop) SetPolicyAdditionalNil(b bool)`

 SetPolicyAdditionalNil sets the value for PolicyAdditional to be an explicit nil

### UnsetPolicyAdditional
`func (o *Shop) UnsetPolicyAdditional()`

UnsetPolicyAdditional ensures that no value is present for PolicyAdditional, not even an explicit nil
### GetPolicySellerInfo

`func (o *Shop) GetPolicySellerInfo() string`

GetPolicySellerInfo returns the PolicySellerInfo field if non-nil, zero value otherwise.

### GetPolicySellerInfoOk

`func (o *Shop) GetPolicySellerInfoOk() (*string, bool)`

GetPolicySellerInfoOk returns a tuple with the PolicySellerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySellerInfo

`func (o *Shop) SetPolicySellerInfo(v string)`

SetPolicySellerInfo sets PolicySellerInfo field to given value.

### HasPolicySellerInfo

`func (o *Shop) HasPolicySellerInfo() bool`

HasPolicySellerInfo returns a boolean if a field has been set.

### SetPolicySellerInfoNil

`func (o *Shop) SetPolicySellerInfoNil(b bool)`

 SetPolicySellerInfoNil sets the value for PolicySellerInfo to be an explicit nil

### UnsetPolicySellerInfo
`func (o *Shop) UnsetPolicySellerInfo()`

UnsetPolicySellerInfo ensures that no value is present for PolicySellerInfo, not even an explicit nil
### GetPolicyUpdateDate

`func (o *Shop) GetPolicyUpdateDate() int64`

GetPolicyUpdateDate returns the PolicyUpdateDate field if non-nil, zero value otherwise.

### GetPolicyUpdateDateOk

`func (o *Shop) GetPolicyUpdateDateOk() (*int64, bool)`

GetPolicyUpdateDateOk returns a tuple with the PolicyUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUpdateDate

`func (o *Shop) SetPolicyUpdateDate(v int64)`

SetPolicyUpdateDate sets PolicyUpdateDate field to given value.

### HasPolicyUpdateDate

`func (o *Shop) HasPolicyUpdateDate() bool`

HasPolicyUpdateDate returns a boolean if a field has been set.

### GetPolicyHasPrivateReceiptInfo

`func (o *Shop) GetPolicyHasPrivateReceiptInfo() bool`

GetPolicyHasPrivateReceiptInfo returns the PolicyHasPrivateReceiptInfo field if non-nil, zero value otherwise.

### GetPolicyHasPrivateReceiptInfoOk

`func (o *Shop) GetPolicyHasPrivateReceiptInfoOk() (*bool, bool)`

GetPolicyHasPrivateReceiptInfoOk returns a tuple with the PolicyHasPrivateReceiptInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyHasPrivateReceiptInfo

`func (o *Shop) SetPolicyHasPrivateReceiptInfo(v bool)`

SetPolicyHasPrivateReceiptInfo sets PolicyHasPrivateReceiptInfo field to given value.

### HasPolicyHasPrivateReceiptInfo

`func (o *Shop) HasPolicyHasPrivateReceiptInfo() bool`

HasPolicyHasPrivateReceiptInfo returns a boolean if a field has been set.

### GetHasUnstructuredPolicies

`func (o *Shop) GetHasUnstructuredPolicies() bool`

GetHasUnstructuredPolicies returns the HasUnstructuredPolicies field if non-nil, zero value otherwise.

### GetHasUnstructuredPoliciesOk

`func (o *Shop) GetHasUnstructuredPoliciesOk() (*bool, bool)`

GetHasUnstructuredPoliciesOk returns a tuple with the HasUnstructuredPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnstructuredPolicies

`func (o *Shop) SetHasUnstructuredPolicies(v bool)`

SetHasUnstructuredPolicies sets HasUnstructuredPolicies field to given value.

### HasHasUnstructuredPolicies

`func (o *Shop) HasHasUnstructuredPolicies() bool`

HasHasUnstructuredPolicies returns a boolean if a field has been set.

### GetPolicyPrivacy

`func (o *Shop) GetPolicyPrivacy() string`

GetPolicyPrivacy returns the PolicyPrivacy field if non-nil, zero value otherwise.

### GetPolicyPrivacyOk

`func (o *Shop) GetPolicyPrivacyOk() (*string, bool)`

GetPolicyPrivacyOk returns a tuple with the PolicyPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPrivacy

`func (o *Shop) SetPolicyPrivacy(v string)`

SetPolicyPrivacy sets PolicyPrivacy field to given value.

### HasPolicyPrivacy

`func (o *Shop) HasPolicyPrivacy() bool`

HasPolicyPrivacy returns a boolean if a field has been set.

### SetPolicyPrivacyNil

`func (o *Shop) SetPolicyPrivacyNil(b bool)`

 SetPolicyPrivacyNil sets the value for PolicyPrivacy to be an explicit nil

### UnsetPolicyPrivacy
`func (o *Shop) UnsetPolicyPrivacy()`

UnsetPolicyPrivacy ensures that no value is present for PolicyPrivacy, not even an explicit nil
### GetVacationAutoreply

`func (o *Shop) GetVacationAutoreply() string`

GetVacationAutoreply returns the VacationAutoreply field if non-nil, zero value otherwise.

### GetVacationAutoreplyOk

`func (o *Shop) GetVacationAutoreplyOk() (*string, bool)`

GetVacationAutoreplyOk returns a tuple with the VacationAutoreply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationAutoreply

`func (o *Shop) SetVacationAutoreply(v string)`

SetVacationAutoreply sets VacationAutoreply field to given value.

### HasVacationAutoreply

`func (o *Shop) HasVacationAutoreply() bool`

HasVacationAutoreply returns a boolean if a field has been set.

### SetVacationAutoreplyNil

`func (o *Shop) SetVacationAutoreplyNil(b bool)`

 SetVacationAutoreplyNil sets the value for VacationAutoreply to be an explicit nil

### UnsetVacationAutoreply
`func (o *Shop) UnsetVacationAutoreply()`

UnsetVacationAutoreply ensures that no value is present for VacationAutoreply, not even an explicit nil
### GetUrl

`func (o *Shop) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Shop) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Shop) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Shop) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetImageUrl760x100

`func (o *Shop) GetImageUrl760x100() string`

GetImageUrl760x100 returns the ImageUrl760x100 field if non-nil, zero value otherwise.

### GetImageUrl760x100Ok

`func (o *Shop) GetImageUrl760x100Ok() (*string, bool)`

GetImageUrl760x100Ok returns a tuple with the ImageUrl760x100 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl760x100

`func (o *Shop) SetImageUrl760x100(v string)`

SetImageUrl760x100 sets ImageUrl760x100 field to given value.

### HasImageUrl760x100

`func (o *Shop) HasImageUrl760x100() bool`

HasImageUrl760x100 returns a boolean if a field has been set.

### SetImageUrl760x100Nil

`func (o *Shop) SetImageUrl760x100Nil(b bool)`

 SetImageUrl760x100Nil sets the value for ImageUrl760x100 to be an explicit nil

### UnsetImageUrl760x100
`func (o *Shop) UnsetImageUrl760x100()`

UnsetImageUrl760x100 ensures that no value is present for ImageUrl760x100, not even an explicit nil
### GetNumFavorers

`func (o *Shop) GetNumFavorers() int64`

GetNumFavorers returns the NumFavorers field if non-nil, zero value otherwise.

### GetNumFavorersOk

`func (o *Shop) GetNumFavorersOk() (*int64, bool)`

GetNumFavorersOk returns a tuple with the NumFavorers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFavorers

`func (o *Shop) SetNumFavorers(v int64)`

SetNumFavorers sets NumFavorers field to given value.

### HasNumFavorers

`func (o *Shop) HasNumFavorers() bool`

HasNumFavorers returns a boolean if a field has been set.

### GetLanguages

`func (o *Shop) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *Shop) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *Shop) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *Shop) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetIconUrlFullxfull

`func (o *Shop) GetIconUrlFullxfull() string`

GetIconUrlFullxfull returns the IconUrlFullxfull field if non-nil, zero value otherwise.

### GetIconUrlFullxfullOk

`func (o *Shop) GetIconUrlFullxfullOk() (*string, bool)`

GetIconUrlFullxfullOk returns a tuple with the IconUrlFullxfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrlFullxfull

`func (o *Shop) SetIconUrlFullxfull(v string)`

SetIconUrlFullxfull sets IconUrlFullxfull field to given value.

### HasIconUrlFullxfull

`func (o *Shop) HasIconUrlFullxfull() bool`

HasIconUrlFullxfull returns a boolean if a field has been set.

### SetIconUrlFullxfullNil

`func (o *Shop) SetIconUrlFullxfullNil(b bool)`

 SetIconUrlFullxfullNil sets the value for IconUrlFullxfull to be an explicit nil

### UnsetIconUrlFullxfull
`func (o *Shop) UnsetIconUrlFullxfull()`

UnsetIconUrlFullxfull ensures that no value is present for IconUrlFullxfull, not even an explicit nil
### GetIsUsingStructuredPolicies

`func (o *Shop) GetIsUsingStructuredPolicies() bool`

GetIsUsingStructuredPolicies returns the IsUsingStructuredPolicies field if non-nil, zero value otherwise.

### GetIsUsingStructuredPoliciesOk

`func (o *Shop) GetIsUsingStructuredPoliciesOk() (*bool, bool)`

GetIsUsingStructuredPoliciesOk returns a tuple with the IsUsingStructuredPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingStructuredPolicies

`func (o *Shop) SetIsUsingStructuredPolicies(v bool)`

SetIsUsingStructuredPolicies sets IsUsingStructuredPolicies field to given value.

### HasIsUsingStructuredPolicies

`func (o *Shop) HasIsUsingStructuredPolicies() bool`

HasIsUsingStructuredPolicies returns a boolean if a field has been set.

### GetHasOnboardedStructuredPolicies

`func (o *Shop) GetHasOnboardedStructuredPolicies() bool`

GetHasOnboardedStructuredPolicies returns the HasOnboardedStructuredPolicies field if non-nil, zero value otherwise.

### GetHasOnboardedStructuredPoliciesOk

`func (o *Shop) GetHasOnboardedStructuredPoliciesOk() (*bool, bool)`

GetHasOnboardedStructuredPoliciesOk returns a tuple with the HasOnboardedStructuredPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOnboardedStructuredPolicies

`func (o *Shop) SetHasOnboardedStructuredPolicies(v bool)`

SetHasOnboardedStructuredPolicies sets HasOnboardedStructuredPolicies field to given value.

### HasHasOnboardedStructuredPolicies

`func (o *Shop) HasHasOnboardedStructuredPolicies() bool`

HasHasOnboardedStructuredPolicies returns a boolean if a field has been set.

### GetIncludeDisputeFormLink

`func (o *Shop) GetIncludeDisputeFormLink() bool`

GetIncludeDisputeFormLink returns the IncludeDisputeFormLink field if non-nil, zero value otherwise.

### GetIncludeDisputeFormLinkOk

`func (o *Shop) GetIncludeDisputeFormLinkOk() (*bool, bool)`

GetIncludeDisputeFormLinkOk returns a tuple with the IncludeDisputeFormLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisputeFormLink

`func (o *Shop) SetIncludeDisputeFormLink(v bool)`

SetIncludeDisputeFormLink sets IncludeDisputeFormLink field to given value.

### HasIncludeDisputeFormLink

`func (o *Shop) HasIncludeDisputeFormLink() bool`

HasIncludeDisputeFormLink returns a boolean if a field has been set.

### GetIsDirectCheckoutOnboarded

`func (o *Shop) GetIsDirectCheckoutOnboarded() bool`

GetIsDirectCheckoutOnboarded returns the IsDirectCheckoutOnboarded field if non-nil, zero value otherwise.

### GetIsDirectCheckoutOnboardedOk

`func (o *Shop) GetIsDirectCheckoutOnboardedOk() (*bool, bool)`

GetIsDirectCheckoutOnboardedOk returns a tuple with the IsDirectCheckoutOnboarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectCheckoutOnboarded

`func (o *Shop) SetIsDirectCheckoutOnboarded(v bool)`

SetIsDirectCheckoutOnboarded sets IsDirectCheckoutOnboarded field to given value.

### HasIsDirectCheckoutOnboarded

`func (o *Shop) HasIsDirectCheckoutOnboarded() bool`

HasIsDirectCheckoutOnboarded returns a boolean if a field has been set.

### GetIsEtsyPaymentsOnboarded

`func (o *Shop) GetIsEtsyPaymentsOnboarded() bool`

GetIsEtsyPaymentsOnboarded returns the IsEtsyPaymentsOnboarded field if non-nil, zero value otherwise.

### GetIsEtsyPaymentsOnboardedOk

`func (o *Shop) GetIsEtsyPaymentsOnboardedOk() (*bool, bool)`

GetIsEtsyPaymentsOnboardedOk returns a tuple with the IsEtsyPaymentsOnboarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEtsyPaymentsOnboarded

`func (o *Shop) SetIsEtsyPaymentsOnboarded(v bool)`

SetIsEtsyPaymentsOnboarded sets IsEtsyPaymentsOnboarded field to given value.

### HasIsEtsyPaymentsOnboarded

`func (o *Shop) HasIsEtsyPaymentsOnboarded() bool`

HasIsEtsyPaymentsOnboarded returns a boolean if a field has been set.

### GetIsCalculatedEligible

`func (o *Shop) GetIsCalculatedEligible() bool`

GetIsCalculatedEligible returns the IsCalculatedEligible field if non-nil, zero value otherwise.

### GetIsCalculatedEligibleOk

`func (o *Shop) GetIsCalculatedEligibleOk() (*bool, bool)`

GetIsCalculatedEligibleOk returns a tuple with the IsCalculatedEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCalculatedEligible

`func (o *Shop) SetIsCalculatedEligible(v bool)`

SetIsCalculatedEligible sets IsCalculatedEligible field to given value.

### HasIsCalculatedEligible

`func (o *Shop) HasIsCalculatedEligible() bool`

HasIsCalculatedEligible returns a boolean if a field has been set.

### GetIsOptedInToBuyerPromise

`func (o *Shop) GetIsOptedInToBuyerPromise() bool`

GetIsOptedInToBuyerPromise returns the IsOptedInToBuyerPromise field if non-nil, zero value otherwise.

### GetIsOptedInToBuyerPromiseOk

`func (o *Shop) GetIsOptedInToBuyerPromiseOk() (*bool, bool)`

GetIsOptedInToBuyerPromiseOk returns a tuple with the IsOptedInToBuyerPromise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptedInToBuyerPromise

`func (o *Shop) SetIsOptedInToBuyerPromise(v bool)`

SetIsOptedInToBuyerPromise sets IsOptedInToBuyerPromise field to given value.

### HasIsOptedInToBuyerPromise

`func (o *Shop) HasIsOptedInToBuyerPromise() bool`

HasIsOptedInToBuyerPromise returns a boolean if a field has been set.

### GetIsShopUsBased

`func (o *Shop) GetIsShopUsBased() bool`

GetIsShopUsBased returns the IsShopUsBased field if non-nil, zero value otherwise.

### GetIsShopUsBasedOk

`func (o *Shop) GetIsShopUsBasedOk() (*bool, bool)`

GetIsShopUsBasedOk returns a tuple with the IsShopUsBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShopUsBased

`func (o *Shop) SetIsShopUsBased(v bool)`

SetIsShopUsBased sets IsShopUsBased field to given value.

### HasIsShopUsBased

`func (o *Shop) HasIsShopUsBased() bool`

HasIsShopUsBased returns a boolean if a field has been set.

### GetTransactionSoldCount

`func (o *Shop) GetTransactionSoldCount() int64`

GetTransactionSoldCount returns the TransactionSoldCount field if non-nil, zero value otherwise.

### GetTransactionSoldCountOk

`func (o *Shop) GetTransactionSoldCountOk() (*int64, bool)`

GetTransactionSoldCountOk returns a tuple with the TransactionSoldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSoldCount

`func (o *Shop) SetTransactionSoldCount(v int64)`

SetTransactionSoldCount sets TransactionSoldCount field to given value.

### HasTransactionSoldCount

`func (o *Shop) HasTransactionSoldCount() bool`

HasTransactionSoldCount returns a boolean if a field has been set.

### GetShippingFromCountryIso

`func (o *Shop) GetShippingFromCountryIso() string`

GetShippingFromCountryIso returns the ShippingFromCountryIso field if non-nil, zero value otherwise.

### GetShippingFromCountryIsoOk

`func (o *Shop) GetShippingFromCountryIsoOk() (*string, bool)`

GetShippingFromCountryIsoOk returns a tuple with the ShippingFromCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingFromCountryIso

`func (o *Shop) SetShippingFromCountryIso(v string)`

SetShippingFromCountryIso sets ShippingFromCountryIso field to given value.

### HasShippingFromCountryIso

`func (o *Shop) HasShippingFromCountryIso() bool`

HasShippingFromCountryIso returns a boolean if a field has been set.

### SetShippingFromCountryIsoNil

`func (o *Shop) SetShippingFromCountryIsoNil(b bool)`

 SetShippingFromCountryIsoNil sets the value for ShippingFromCountryIso to be an explicit nil

### UnsetShippingFromCountryIso
`func (o *Shop) UnsetShippingFromCountryIso()`

UnsetShippingFromCountryIso ensures that no value is present for ShippingFromCountryIso, not even an explicit nil
### GetShopLocationCountryIso

`func (o *Shop) GetShopLocationCountryIso() string`

GetShopLocationCountryIso returns the ShopLocationCountryIso field if non-nil, zero value otherwise.

### GetShopLocationCountryIsoOk

`func (o *Shop) GetShopLocationCountryIsoOk() (*string, bool)`

GetShopLocationCountryIsoOk returns a tuple with the ShopLocationCountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopLocationCountryIso

`func (o *Shop) SetShopLocationCountryIso(v string)`

SetShopLocationCountryIso sets ShopLocationCountryIso field to given value.

### HasShopLocationCountryIso

`func (o *Shop) HasShopLocationCountryIso() bool`

HasShopLocationCountryIso returns a boolean if a field has been set.

### SetShopLocationCountryIsoNil

`func (o *Shop) SetShopLocationCountryIsoNil(b bool)`

 SetShopLocationCountryIsoNil sets the value for ShopLocationCountryIso to be an explicit nil

### UnsetShopLocationCountryIso
`func (o *Shop) UnsetShopLocationCountryIso()`

UnsetShopLocationCountryIso ensures that no value is present for ShopLocationCountryIso, not even an explicit nil
### GetReviewCount

`func (o *Shop) GetReviewCount() int64`

GetReviewCount returns the ReviewCount field if non-nil, zero value otherwise.

### GetReviewCountOk

`func (o *Shop) GetReviewCountOk() (*int64, bool)`

GetReviewCountOk returns a tuple with the ReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewCount

`func (o *Shop) SetReviewCount(v int64)`

SetReviewCount sets ReviewCount field to given value.

### HasReviewCount

`func (o *Shop) HasReviewCount() bool`

HasReviewCount returns a boolean if a field has been set.

### SetReviewCountNil

`func (o *Shop) SetReviewCountNil(b bool)`

 SetReviewCountNil sets the value for ReviewCount to be an explicit nil

### UnsetReviewCount
`func (o *Shop) UnsetReviewCount()`

UnsetReviewCount ensures that no value is present for ReviewCount, not even an explicit nil
### GetReviewAverage

`func (o *Shop) GetReviewAverage() float32`

GetReviewAverage returns the ReviewAverage field if non-nil, zero value otherwise.

### GetReviewAverageOk

`func (o *Shop) GetReviewAverageOk() (*float32, bool)`

GetReviewAverageOk returns a tuple with the ReviewAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewAverage

`func (o *Shop) SetReviewAverage(v float32)`

SetReviewAverage sets ReviewAverage field to given value.

### HasReviewAverage

`func (o *Shop) HasReviewAverage() bool`

HasReviewAverage returns a boolean if a field has been set.

### SetReviewAverageNil

`func (o *Shop) SetReviewAverageNil(b bool)`

 SetReviewAverageNil sets the value for ReviewAverage to be an explicit nil

### UnsetReviewAverage
`func (o *Shop) UnsetReviewAverage()`

UnsetReviewAverage ensures that no value is present for ReviewAverage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


