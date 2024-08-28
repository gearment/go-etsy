# ShopListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | Pointer to **int64** | The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction. | [optional] 
**UserId** | Pointer to **int64** | The numeric ID for the [user](/documentation/reference#tag/User) posting the listing. | [optional] 
**ShopId** | Pointer to **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | [optional] 
**Title** | Pointer to **string** | The listing&#39;s title string. When creating or updating a listing, valid title strings contain only letters, numbers, punctuation marks, mathematical symbols, whitespace characters, ™, ©, and ®. (regex: /[^\\\\p{L}\\\\p{Nd}\\\\p{P}\\\\p{Sm}\\\\p{Zs}™©®]/u) You can only use the %, :, &amp; and + characters once each. | [optional] 
**Description** | Pointer to **string** | A description string of the product for sale in the listing. | [optional] 
**State** | Pointer to [**ShopListingState**](ShopListingState.md) |  | [optional] 
**CreationTimestamp** | Pointer to **int64** | The listing\\&#39;s creation time, in epoch seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The listing\\&#39;s creation time, in epoch seconds. | [optional] 
**EndingTimestamp** | Pointer to **int64** | The listing\\&#39;s expiration time, in epoch seconds. | [optional] 
**OriginalCreationTimestamp** | Pointer to **int64** | The listing\\&#39;s creation time, in epoch seconds. | [optional] 
**LastModifiedTimestamp** | Pointer to **int64** | The time of the last update to the listing, in epoch seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int64** | The time of the last update to the listing, in epoch seconds. | [optional] 
**StateTimestamp** | Pointer to **NullableInt64** | The date and time of the last state change of this listing. | [optional] 
**Quantity** | Pointer to **int64** | The positive non-zero number of products available for purchase in the listing. Note: The listing quantity is the sum of available offering quantities. You can request the quantities for individual offerings from the ListingInventory resource using the [getListingInventory](/documentation/reference#operation/getListingInventory) endpoint. | [optional] 
**ShopSectionId** | Pointer to **NullableInt64** | The numeric ID of a section in a specific Etsy shop. | [optional] 
**FeaturedRank** | Pointer to **int64** | The positive non-zero numeric position in the featured listings of the shop, with rank 1 listings appearing in the left-most position in featured listing on a shop’s home page. | [optional] 
**Url** | Pointer to **string** | The full URL to the listing&#39;s page on Etsy. | [optional] 
**NumFavorers** | Pointer to **int64** | The number of users who marked this Listing a favorite. | [optional] 
**NonTaxable** | Pointer to **bool** | When true, applicable [shop](/documentation/reference#tag/Shop) tax rates do not apply to this listing at checkout. | [optional] 
**IsTaxable** | Pointer to **bool** | When true, applicable [shop](/documentation/reference#tag/Shop) tax rates apply to this listing at checkout. | [optional] 
**IsCustomizable** | Pointer to **bool** | When true, a buyer may contact the seller for a customized order. The default value is true when a shop accepts custom orders. Does not apply to shops that do not accept custom orders. | [optional] 
**IsPersonalizable** | Pointer to **bool** | When true, this listing is personalizable. The default value is null. | [optional] 
**PersonalizationIsRequired** | Pointer to **bool** | When true, this listing requires personalization. The default value is null. Will only change if is_personalizable is &#39;true&#39;. | [optional] 
**PersonalizationCharCountMax** | Pointer to **NullableInt64** | This is an integer value representing the maximum length for the personalization message entered by the buyer. Will only change if is_personalizable is &#39;true&#39;. | [optional] 
**PersonalizationInstructions** | Pointer to **NullableString** | When true, this listing requires personalization. The default value is null. Will only change if is_personalizable is &#39;true&#39;. | [optional] 
**ListingType** | Pointer to [**CreateDraftListingRequestType**](CreateDraftListingRequestType.md) |  | [optional] 
**Tags** | Pointer to **[]string** | A comma-separated list of tag strings for the listing. When creating or updating a listing, valid tag strings contain only letters, numbers, whitespace characters, -, &#39;, ™, ©, and ®. (regex: /[^\\\\p{L}\\\\p{Nd}\\\\p{Zs}\\\\-&#39;™©®]/u) Default value is null. | [optional] 
**Materials** | Pointer to **[]string** | A list of material strings for materials used in the product. Valid materials strings contain only letters, numbers, and whitespace characters. (regex: /[^\\\\p{L}\\\\p{Nd}\\\\p{Zs}]/u) Default value is null. | [optional] 
**ShippingProfileId** | Pointer to **NullableInt64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | [optional] 
**ReturnPolicyId** | Pointer to **NullableInt64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). | [optional] 
**ProcessingMin** | Pointer to **NullableInt64** | The minimum number of days required to process this listing. Default value is null. | [optional] 
**ProcessingMax** | Pointer to **NullableInt64** | The maximum number of days required to process this listing. Default value is null. | [optional] 
**WhoMade** | Pointer to [**NullableShopListingWhoMade**](ShopListingWhoMade.md) |  | [optional] 
**WhenMade** | Pointer to [**NullableShopListingWhenMade**](ShopListingWhenMade.md) |  | [optional] 
**IsSupply** | Pointer to **NullableBool** | When true, tags the listing as a supply product, else indicates that it&#39;s a finished product. Helps buyers locate the listing under the Supplies heading. Requires &#39;who_made&#39; and &#39;when_made&#39;. | [optional] 
**ItemWeight** | Pointer to **NullableFloat32** | The numeric weight of the product measured in units set in \\&#39;item_weight_unit\\&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemWeightUnit** | Pointer to [**NullableCreateDraftListingRequestItemWeightUnit**](CreateDraftListingRequestItemWeightUnit.md) |  | [optional] 
**ItemLength** | Pointer to **NullableFloat32** | The numeric length of the product measured in units set in \\&#39;item_dimensions_unit\\&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemWidth** | Pointer to **NullableFloat32** | The numeric width of the product measured in units set in \\&#39;item_dimensions_unit\\&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemHeight** | Pointer to **NullableFloat32** | The numeric length of the product measured in units set in \\&#39;item_dimensions_unit\\&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemDimensionsUnit** | Pointer to [**NullableCreateDraftListingRequestItemDimensionsUnit**](CreateDraftListingRequestItemDimensionsUnit.md) |  | [optional] 
**IsPrivate** | Pointer to **bool** | When true, this is a private listing intended for a specific buyer and hidden from shop view. | [optional] 
**Style** | Pointer to **[]string** | An array of style strings for this listing, each of which is free-form text string such as \\\&quot;Formal\\\&quot;, or \\\&quot;Steampunk\\\&quot;. When creating or updating a listing, the listing may have up to two styles. Valid style strings contain only letters, numbers, and whitespace characters. (regex: /[^\\\\p{L}\\\\p{Nd}\\\\p{Zs}]/u) Default value is null. | [optional] 
**FileData** | Pointer to **string** | A string describing the files attached to a digital listing. | [optional] 
**HasVariations** | Pointer to **bool** | When true, the listing has variations. | [optional] 
**ShouldAutoRenew** | Pointer to **bool** | When true, renews a listing for four months upon expiration. | [optional] 
**Language** | Pointer to **NullableString** | The IETF language tag for the default language of the listing. Ex: &#x60;de&#x60;, &#x60;en&#x60;, &#x60;es&#x60;, &#x60;fr&#x60;, &#x60;it&#x60;, &#x60;ja&#x60;, &#x60;nl&#x60;, &#x60;pl&#x60;, &#x60;pt&#x60;, &#x60;ru&#x60;. | [optional] 
**Price** | Pointer to [**Money**](Money.md) | The positive non-zero price of the product. (Sold product listings are private) Note: The price is the minimum possible price. The [&#x60;getListingInventory&#x60;](/documentation/reference/#operation/getListingInventory) method requests exact prices for available offerings. | [optional] 
**TaxonomyId** | Pointer to **NullableInt64** | The numerical taxonomy ID of the listing. See [SellerTaxonomy](/documentation/reference#tag/SellerTaxonomy) and [BuyerTaxonomy](/documentation/reference#tag/BuyerTaxonomy) for more information. | [optional] 

## Methods

### NewShopListing

`func NewShopListing() *ShopListing`

NewShopListing instantiates a new ShopListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopListingWithDefaults

`func NewShopListingWithDefaults() *ShopListing`

NewShopListingWithDefaults instantiates a new ShopListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *ShopListing) GetListingId() int64`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ShopListing) GetListingIdOk() (*int64, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ShopListing) SetListingId(v int64)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *ShopListing) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### GetUserId

`func (o *ShopListing) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ShopListing) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ShopListing) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ShopListing) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetShopId

`func (o *ShopListing) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ShopListing) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ShopListing) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ShopListing) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetTitle

`func (o *ShopListing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShopListing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShopListing) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ShopListing) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ShopListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShopListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShopListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShopListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *ShopListing) GetState() ShopListingState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ShopListing) GetStateOk() (*ShopListingState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ShopListing) SetState(v ShopListingState)`

SetState sets State field to given value.

### HasState

`func (o *ShopListing) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *ShopListing) GetCreationTimestamp() int64`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ShopListing) GetCreationTimestampOk() (*int64, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ShopListing) SetCreationTimestamp(v int64)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ShopListing) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ShopListing) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ShopListing) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ShopListing) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ShopListing) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetEndingTimestamp

`func (o *ShopListing) GetEndingTimestamp() int64`

GetEndingTimestamp returns the EndingTimestamp field if non-nil, zero value otherwise.

### GetEndingTimestampOk

`func (o *ShopListing) GetEndingTimestampOk() (*int64, bool)`

GetEndingTimestampOk returns a tuple with the EndingTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingTimestamp

`func (o *ShopListing) SetEndingTimestamp(v int64)`

SetEndingTimestamp sets EndingTimestamp field to given value.

### HasEndingTimestamp

`func (o *ShopListing) HasEndingTimestamp() bool`

HasEndingTimestamp returns a boolean if a field has been set.

### GetOriginalCreationTimestamp

`func (o *ShopListing) GetOriginalCreationTimestamp() int64`

GetOriginalCreationTimestamp returns the OriginalCreationTimestamp field if non-nil, zero value otherwise.

### GetOriginalCreationTimestampOk

`func (o *ShopListing) GetOriginalCreationTimestampOk() (*int64, bool)`

GetOriginalCreationTimestampOk returns a tuple with the OriginalCreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCreationTimestamp

`func (o *ShopListing) SetOriginalCreationTimestamp(v int64)`

SetOriginalCreationTimestamp sets OriginalCreationTimestamp field to given value.

### HasOriginalCreationTimestamp

`func (o *ShopListing) HasOriginalCreationTimestamp() bool`

HasOriginalCreationTimestamp returns a boolean if a field has been set.

### GetLastModifiedTimestamp

`func (o *ShopListing) GetLastModifiedTimestamp() int64`

GetLastModifiedTimestamp returns the LastModifiedTimestamp field if non-nil, zero value otherwise.

### GetLastModifiedTimestampOk

`func (o *ShopListing) GetLastModifiedTimestampOk() (*int64, bool)`

GetLastModifiedTimestampOk returns a tuple with the LastModifiedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimestamp

`func (o *ShopListing) SetLastModifiedTimestamp(v int64)`

SetLastModifiedTimestamp sets LastModifiedTimestamp field to given value.

### HasLastModifiedTimestamp

`func (o *ShopListing) HasLastModifiedTimestamp() bool`

HasLastModifiedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ShopListing) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ShopListing) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ShopListing) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *ShopListing) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetStateTimestamp

`func (o *ShopListing) GetStateTimestamp() int64`

GetStateTimestamp returns the StateTimestamp field if non-nil, zero value otherwise.

### GetStateTimestampOk

`func (o *ShopListing) GetStateTimestampOk() (*int64, bool)`

GetStateTimestampOk returns a tuple with the StateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTimestamp

`func (o *ShopListing) SetStateTimestamp(v int64)`

SetStateTimestamp sets StateTimestamp field to given value.

### HasStateTimestamp

`func (o *ShopListing) HasStateTimestamp() bool`

HasStateTimestamp returns a boolean if a field has been set.

### SetStateTimestampNil

`func (o *ShopListing) SetStateTimestampNil(b bool)`

 SetStateTimestampNil sets the value for StateTimestamp to be an explicit nil

### UnsetStateTimestamp
`func (o *ShopListing) UnsetStateTimestamp()`

UnsetStateTimestamp ensures that no value is present for StateTimestamp, not even an explicit nil
### GetQuantity

`func (o *ShopListing) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShopListing) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShopListing) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ShopListing) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetShopSectionId

`func (o *ShopListing) GetShopSectionId() int64`

GetShopSectionId returns the ShopSectionId field if non-nil, zero value otherwise.

### GetShopSectionIdOk

`func (o *ShopListing) GetShopSectionIdOk() (*int64, bool)`

GetShopSectionIdOk returns a tuple with the ShopSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopSectionId

`func (o *ShopListing) SetShopSectionId(v int64)`

SetShopSectionId sets ShopSectionId field to given value.

### HasShopSectionId

`func (o *ShopListing) HasShopSectionId() bool`

HasShopSectionId returns a boolean if a field has been set.

### SetShopSectionIdNil

`func (o *ShopListing) SetShopSectionIdNil(b bool)`

 SetShopSectionIdNil sets the value for ShopSectionId to be an explicit nil

### UnsetShopSectionId
`func (o *ShopListing) UnsetShopSectionId()`

UnsetShopSectionId ensures that no value is present for ShopSectionId, not even an explicit nil
### GetFeaturedRank

`func (o *ShopListing) GetFeaturedRank() int64`

GetFeaturedRank returns the FeaturedRank field if non-nil, zero value otherwise.

### GetFeaturedRankOk

`func (o *ShopListing) GetFeaturedRankOk() (*int64, bool)`

GetFeaturedRankOk returns a tuple with the FeaturedRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedRank

`func (o *ShopListing) SetFeaturedRank(v int64)`

SetFeaturedRank sets FeaturedRank field to given value.

### HasFeaturedRank

`func (o *ShopListing) HasFeaturedRank() bool`

HasFeaturedRank returns a boolean if a field has been set.

### GetUrl

`func (o *ShopListing) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ShopListing) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ShopListing) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ShopListing) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNumFavorers

`func (o *ShopListing) GetNumFavorers() int64`

GetNumFavorers returns the NumFavorers field if non-nil, zero value otherwise.

### GetNumFavorersOk

`func (o *ShopListing) GetNumFavorersOk() (*int64, bool)`

GetNumFavorersOk returns a tuple with the NumFavorers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFavorers

`func (o *ShopListing) SetNumFavorers(v int64)`

SetNumFavorers sets NumFavorers field to given value.

### HasNumFavorers

`func (o *ShopListing) HasNumFavorers() bool`

HasNumFavorers returns a boolean if a field has been set.

### GetNonTaxable

`func (o *ShopListing) GetNonTaxable() bool`

GetNonTaxable returns the NonTaxable field if non-nil, zero value otherwise.

### GetNonTaxableOk

`func (o *ShopListing) GetNonTaxableOk() (*bool, bool)`

GetNonTaxableOk returns a tuple with the NonTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTaxable

`func (o *ShopListing) SetNonTaxable(v bool)`

SetNonTaxable sets NonTaxable field to given value.

### HasNonTaxable

`func (o *ShopListing) HasNonTaxable() bool`

HasNonTaxable returns a boolean if a field has been set.

### GetIsTaxable

`func (o *ShopListing) GetIsTaxable() bool`

GetIsTaxable returns the IsTaxable field if non-nil, zero value otherwise.

### GetIsTaxableOk

`func (o *ShopListing) GetIsTaxableOk() (*bool, bool)`

GetIsTaxableOk returns a tuple with the IsTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxable

`func (o *ShopListing) SetIsTaxable(v bool)`

SetIsTaxable sets IsTaxable field to given value.

### HasIsTaxable

`func (o *ShopListing) HasIsTaxable() bool`

HasIsTaxable returns a boolean if a field has been set.

### GetIsCustomizable

`func (o *ShopListing) GetIsCustomizable() bool`

GetIsCustomizable returns the IsCustomizable field if non-nil, zero value otherwise.

### GetIsCustomizableOk

`func (o *ShopListing) GetIsCustomizableOk() (*bool, bool)`

GetIsCustomizableOk returns a tuple with the IsCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomizable

`func (o *ShopListing) SetIsCustomizable(v bool)`

SetIsCustomizable sets IsCustomizable field to given value.

### HasIsCustomizable

`func (o *ShopListing) HasIsCustomizable() bool`

HasIsCustomizable returns a boolean if a field has been set.

### GetIsPersonalizable

`func (o *ShopListing) GetIsPersonalizable() bool`

GetIsPersonalizable returns the IsPersonalizable field if non-nil, zero value otherwise.

### GetIsPersonalizableOk

`func (o *ShopListing) GetIsPersonalizableOk() (*bool, bool)`

GetIsPersonalizableOk returns a tuple with the IsPersonalizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonalizable

`func (o *ShopListing) SetIsPersonalizable(v bool)`

SetIsPersonalizable sets IsPersonalizable field to given value.

### HasIsPersonalizable

`func (o *ShopListing) HasIsPersonalizable() bool`

HasIsPersonalizable returns a boolean if a field has been set.

### GetPersonalizationIsRequired

`func (o *ShopListing) GetPersonalizationIsRequired() bool`

GetPersonalizationIsRequired returns the PersonalizationIsRequired field if non-nil, zero value otherwise.

### GetPersonalizationIsRequiredOk

`func (o *ShopListing) GetPersonalizationIsRequiredOk() (*bool, bool)`

GetPersonalizationIsRequiredOk returns a tuple with the PersonalizationIsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationIsRequired

`func (o *ShopListing) SetPersonalizationIsRequired(v bool)`

SetPersonalizationIsRequired sets PersonalizationIsRequired field to given value.

### HasPersonalizationIsRequired

`func (o *ShopListing) HasPersonalizationIsRequired() bool`

HasPersonalizationIsRequired returns a boolean if a field has been set.

### GetPersonalizationCharCountMax

`func (o *ShopListing) GetPersonalizationCharCountMax() int64`

GetPersonalizationCharCountMax returns the PersonalizationCharCountMax field if non-nil, zero value otherwise.

### GetPersonalizationCharCountMaxOk

`func (o *ShopListing) GetPersonalizationCharCountMaxOk() (*int64, bool)`

GetPersonalizationCharCountMaxOk returns a tuple with the PersonalizationCharCountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationCharCountMax

`func (o *ShopListing) SetPersonalizationCharCountMax(v int64)`

SetPersonalizationCharCountMax sets PersonalizationCharCountMax field to given value.

### HasPersonalizationCharCountMax

`func (o *ShopListing) HasPersonalizationCharCountMax() bool`

HasPersonalizationCharCountMax returns a boolean if a field has been set.

### SetPersonalizationCharCountMaxNil

`func (o *ShopListing) SetPersonalizationCharCountMaxNil(b bool)`

 SetPersonalizationCharCountMaxNil sets the value for PersonalizationCharCountMax to be an explicit nil

### UnsetPersonalizationCharCountMax
`func (o *ShopListing) UnsetPersonalizationCharCountMax()`

UnsetPersonalizationCharCountMax ensures that no value is present for PersonalizationCharCountMax, not even an explicit nil
### GetPersonalizationInstructions

`func (o *ShopListing) GetPersonalizationInstructions() string`

GetPersonalizationInstructions returns the PersonalizationInstructions field if non-nil, zero value otherwise.

### GetPersonalizationInstructionsOk

`func (o *ShopListing) GetPersonalizationInstructionsOk() (*string, bool)`

GetPersonalizationInstructionsOk returns a tuple with the PersonalizationInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationInstructions

`func (o *ShopListing) SetPersonalizationInstructions(v string)`

SetPersonalizationInstructions sets PersonalizationInstructions field to given value.

### HasPersonalizationInstructions

`func (o *ShopListing) HasPersonalizationInstructions() bool`

HasPersonalizationInstructions returns a boolean if a field has been set.

### SetPersonalizationInstructionsNil

`func (o *ShopListing) SetPersonalizationInstructionsNil(b bool)`

 SetPersonalizationInstructionsNil sets the value for PersonalizationInstructions to be an explicit nil

### UnsetPersonalizationInstructions
`func (o *ShopListing) UnsetPersonalizationInstructions()`

UnsetPersonalizationInstructions ensures that no value is present for PersonalizationInstructions, not even an explicit nil
### GetListingType

`func (o *ShopListing) GetListingType() CreateDraftListingRequestType`

GetListingType returns the ListingType field if non-nil, zero value otherwise.

### GetListingTypeOk

`func (o *ShopListing) GetListingTypeOk() (*CreateDraftListingRequestType, bool)`

GetListingTypeOk returns a tuple with the ListingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingType

`func (o *ShopListing) SetListingType(v CreateDraftListingRequestType)`

SetListingType sets ListingType field to given value.

### HasListingType

`func (o *ShopListing) HasListingType() bool`

HasListingType returns a boolean if a field has been set.

### GetTags

`func (o *ShopListing) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ShopListing) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ShopListing) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ShopListing) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMaterials

`func (o *ShopListing) GetMaterials() []string`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *ShopListing) GetMaterialsOk() (*[]string, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *ShopListing) SetMaterials(v []string)`

SetMaterials sets Materials field to given value.

### HasMaterials

`func (o *ShopListing) HasMaterials() bool`

HasMaterials returns a boolean if a field has been set.

### GetShippingProfileId

`func (o *ShopListing) GetShippingProfileId() int64`

GetShippingProfileId returns the ShippingProfileId field if non-nil, zero value otherwise.

### GetShippingProfileIdOk

`func (o *ShopListing) GetShippingProfileIdOk() (*int64, bool)`

GetShippingProfileIdOk returns a tuple with the ShippingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileId

`func (o *ShopListing) SetShippingProfileId(v int64)`

SetShippingProfileId sets ShippingProfileId field to given value.

### HasShippingProfileId

`func (o *ShopListing) HasShippingProfileId() bool`

HasShippingProfileId returns a boolean if a field has been set.

### SetShippingProfileIdNil

`func (o *ShopListing) SetShippingProfileIdNil(b bool)`

 SetShippingProfileIdNil sets the value for ShippingProfileId to be an explicit nil

### UnsetShippingProfileId
`func (o *ShopListing) UnsetShippingProfileId()`

UnsetShippingProfileId ensures that no value is present for ShippingProfileId, not even an explicit nil
### GetReturnPolicyId

`func (o *ShopListing) GetReturnPolicyId() int64`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *ShopListing) GetReturnPolicyIdOk() (*int64, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *ShopListing) SetReturnPolicyId(v int64)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *ShopListing) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### SetReturnPolicyIdNil

`func (o *ShopListing) SetReturnPolicyIdNil(b bool)`

 SetReturnPolicyIdNil sets the value for ReturnPolicyId to be an explicit nil

### UnsetReturnPolicyId
`func (o *ShopListing) UnsetReturnPolicyId()`

UnsetReturnPolicyId ensures that no value is present for ReturnPolicyId, not even an explicit nil
### GetProcessingMin

`func (o *ShopListing) GetProcessingMin() int64`

GetProcessingMin returns the ProcessingMin field if non-nil, zero value otherwise.

### GetProcessingMinOk

`func (o *ShopListing) GetProcessingMinOk() (*int64, bool)`

GetProcessingMinOk returns a tuple with the ProcessingMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMin

`func (o *ShopListing) SetProcessingMin(v int64)`

SetProcessingMin sets ProcessingMin field to given value.

### HasProcessingMin

`func (o *ShopListing) HasProcessingMin() bool`

HasProcessingMin returns a boolean if a field has been set.

### SetProcessingMinNil

`func (o *ShopListing) SetProcessingMinNil(b bool)`

 SetProcessingMinNil sets the value for ProcessingMin to be an explicit nil

### UnsetProcessingMin
`func (o *ShopListing) UnsetProcessingMin()`

UnsetProcessingMin ensures that no value is present for ProcessingMin, not even an explicit nil
### GetProcessingMax

`func (o *ShopListing) GetProcessingMax() int64`

GetProcessingMax returns the ProcessingMax field if non-nil, zero value otherwise.

### GetProcessingMaxOk

`func (o *ShopListing) GetProcessingMaxOk() (*int64, bool)`

GetProcessingMaxOk returns a tuple with the ProcessingMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMax

`func (o *ShopListing) SetProcessingMax(v int64)`

SetProcessingMax sets ProcessingMax field to given value.

### HasProcessingMax

`func (o *ShopListing) HasProcessingMax() bool`

HasProcessingMax returns a boolean if a field has been set.

### SetProcessingMaxNil

`func (o *ShopListing) SetProcessingMaxNil(b bool)`

 SetProcessingMaxNil sets the value for ProcessingMax to be an explicit nil

### UnsetProcessingMax
`func (o *ShopListing) UnsetProcessingMax()`

UnsetProcessingMax ensures that no value is present for ProcessingMax, not even an explicit nil
### GetWhoMade

`func (o *ShopListing) GetWhoMade() ShopListingWhoMade`

GetWhoMade returns the WhoMade field if non-nil, zero value otherwise.

### GetWhoMadeOk

`func (o *ShopListing) GetWhoMadeOk() (*ShopListingWhoMade, bool)`

GetWhoMadeOk returns a tuple with the WhoMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoMade

`func (o *ShopListing) SetWhoMade(v ShopListingWhoMade)`

SetWhoMade sets WhoMade field to given value.

### HasWhoMade

`func (o *ShopListing) HasWhoMade() bool`

HasWhoMade returns a boolean if a field has been set.

### SetWhoMadeNil

`func (o *ShopListing) SetWhoMadeNil(b bool)`

 SetWhoMadeNil sets the value for WhoMade to be an explicit nil

### UnsetWhoMade
`func (o *ShopListing) UnsetWhoMade()`

UnsetWhoMade ensures that no value is present for WhoMade, not even an explicit nil
### GetWhenMade

`func (o *ShopListing) GetWhenMade() ShopListingWhenMade`

GetWhenMade returns the WhenMade field if non-nil, zero value otherwise.

### GetWhenMadeOk

`func (o *ShopListing) GetWhenMadeOk() (*ShopListingWhenMade, bool)`

GetWhenMadeOk returns a tuple with the WhenMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenMade

`func (o *ShopListing) SetWhenMade(v ShopListingWhenMade)`

SetWhenMade sets WhenMade field to given value.

### HasWhenMade

`func (o *ShopListing) HasWhenMade() bool`

HasWhenMade returns a boolean if a field has been set.

### SetWhenMadeNil

`func (o *ShopListing) SetWhenMadeNil(b bool)`

 SetWhenMadeNil sets the value for WhenMade to be an explicit nil

### UnsetWhenMade
`func (o *ShopListing) UnsetWhenMade()`

UnsetWhenMade ensures that no value is present for WhenMade, not even an explicit nil
### GetIsSupply

`func (o *ShopListing) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *ShopListing) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *ShopListing) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.

### HasIsSupply

`func (o *ShopListing) HasIsSupply() bool`

HasIsSupply returns a boolean if a field has been set.

### SetIsSupplyNil

`func (o *ShopListing) SetIsSupplyNil(b bool)`

 SetIsSupplyNil sets the value for IsSupply to be an explicit nil

### UnsetIsSupply
`func (o *ShopListing) UnsetIsSupply()`

UnsetIsSupply ensures that no value is present for IsSupply, not even an explicit nil
### GetItemWeight

`func (o *ShopListing) GetItemWeight() float32`

GetItemWeight returns the ItemWeight field if non-nil, zero value otherwise.

### GetItemWeightOk

`func (o *ShopListing) GetItemWeightOk() (*float32, bool)`

GetItemWeightOk returns a tuple with the ItemWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWeight

`func (o *ShopListing) SetItemWeight(v float32)`

SetItemWeight sets ItemWeight field to given value.

### HasItemWeight

`func (o *ShopListing) HasItemWeight() bool`

HasItemWeight returns a boolean if a field has been set.

### SetItemWeightNil

`func (o *ShopListing) SetItemWeightNil(b bool)`

 SetItemWeightNil sets the value for ItemWeight to be an explicit nil

### UnsetItemWeight
`func (o *ShopListing) UnsetItemWeight()`

UnsetItemWeight ensures that no value is present for ItemWeight, not even an explicit nil
### GetItemWeightUnit

`func (o *ShopListing) GetItemWeightUnit() CreateDraftListingRequestItemWeightUnit`

GetItemWeightUnit returns the ItemWeightUnit field if non-nil, zero value otherwise.

### GetItemWeightUnitOk

`func (o *ShopListing) GetItemWeightUnitOk() (*CreateDraftListingRequestItemWeightUnit, bool)`

GetItemWeightUnitOk returns a tuple with the ItemWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWeightUnit

`func (o *ShopListing) SetItemWeightUnit(v CreateDraftListingRequestItemWeightUnit)`

SetItemWeightUnit sets ItemWeightUnit field to given value.

### HasItemWeightUnit

`func (o *ShopListing) HasItemWeightUnit() bool`

HasItemWeightUnit returns a boolean if a field has been set.

### SetItemWeightUnitNil

`func (o *ShopListing) SetItemWeightUnitNil(b bool)`

 SetItemWeightUnitNil sets the value for ItemWeightUnit to be an explicit nil

### UnsetItemWeightUnit
`func (o *ShopListing) UnsetItemWeightUnit()`

UnsetItemWeightUnit ensures that no value is present for ItemWeightUnit, not even an explicit nil
### GetItemLength

`func (o *ShopListing) GetItemLength() float32`

GetItemLength returns the ItemLength field if non-nil, zero value otherwise.

### GetItemLengthOk

`func (o *ShopListing) GetItemLengthOk() (*float32, bool)`

GetItemLengthOk returns a tuple with the ItemLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLength

`func (o *ShopListing) SetItemLength(v float32)`

SetItemLength sets ItemLength field to given value.

### HasItemLength

`func (o *ShopListing) HasItemLength() bool`

HasItemLength returns a boolean if a field has been set.

### SetItemLengthNil

`func (o *ShopListing) SetItemLengthNil(b bool)`

 SetItemLengthNil sets the value for ItemLength to be an explicit nil

### UnsetItemLength
`func (o *ShopListing) UnsetItemLength()`

UnsetItemLength ensures that no value is present for ItemLength, not even an explicit nil
### GetItemWidth

`func (o *ShopListing) GetItemWidth() float32`

GetItemWidth returns the ItemWidth field if non-nil, zero value otherwise.

### GetItemWidthOk

`func (o *ShopListing) GetItemWidthOk() (*float32, bool)`

GetItemWidthOk returns a tuple with the ItemWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWidth

`func (o *ShopListing) SetItemWidth(v float32)`

SetItemWidth sets ItemWidth field to given value.

### HasItemWidth

`func (o *ShopListing) HasItemWidth() bool`

HasItemWidth returns a boolean if a field has been set.

### SetItemWidthNil

`func (o *ShopListing) SetItemWidthNil(b bool)`

 SetItemWidthNil sets the value for ItemWidth to be an explicit nil

### UnsetItemWidth
`func (o *ShopListing) UnsetItemWidth()`

UnsetItemWidth ensures that no value is present for ItemWidth, not even an explicit nil
### GetItemHeight

`func (o *ShopListing) GetItemHeight() float32`

GetItemHeight returns the ItemHeight field if non-nil, zero value otherwise.

### GetItemHeightOk

`func (o *ShopListing) GetItemHeightOk() (*float32, bool)`

GetItemHeightOk returns a tuple with the ItemHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemHeight

`func (o *ShopListing) SetItemHeight(v float32)`

SetItemHeight sets ItemHeight field to given value.

### HasItemHeight

`func (o *ShopListing) HasItemHeight() bool`

HasItemHeight returns a boolean if a field has been set.

### SetItemHeightNil

`func (o *ShopListing) SetItemHeightNil(b bool)`

 SetItemHeightNil sets the value for ItemHeight to be an explicit nil

### UnsetItemHeight
`func (o *ShopListing) UnsetItemHeight()`

UnsetItemHeight ensures that no value is present for ItemHeight, not even an explicit nil
### GetItemDimensionsUnit

`func (o *ShopListing) GetItemDimensionsUnit() CreateDraftListingRequestItemDimensionsUnit`

GetItemDimensionsUnit returns the ItemDimensionsUnit field if non-nil, zero value otherwise.

### GetItemDimensionsUnitOk

`func (o *ShopListing) GetItemDimensionsUnitOk() (*CreateDraftListingRequestItemDimensionsUnit, bool)`

GetItemDimensionsUnitOk returns a tuple with the ItemDimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDimensionsUnit

`func (o *ShopListing) SetItemDimensionsUnit(v CreateDraftListingRequestItemDimensionsUnit)`

SetItemDimensionsUnit sets ItemDimensionsUnit field to given value.

### HasItemDimensionsUnit

`func (o *ShopListing) HasItemDimensionsUnit() bool`

HasItemDimensionsUnit returns a boolean if a field has been set.

### SetItemDimensionsUnitNil

`func (o *ShopListing) SetItemDimensionsUnitNil(b bool)`

 SetItemDimensionsUnitNil sets the value for ItemDimensionsUnit to be an explicit nil

### UnsetItemDimensionsUnit
`func (o *ShopListing) UnsetItemDimensionsUnit()`

UnsetItemDimensionsUnit ensures that no value is present for ItemDimensionsUnit, not even an explicit nil
### GetIsPrivate

`func (o *ShopListing) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ShopListing) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ShopListing) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ShopListing) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetStyle

`func (o *ShopListing) GetStyle() []string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ShopListing) GetStyleOk() (*[]string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ShopListing) SetStyle(v []string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ShopListing) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetFileData

`func (o *ShopListing) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *ShopListing) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *ShopListing) SetFileData(v string)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *ShopListing) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### GetHasVariations

`func (o *ShopListing) GetHasVariations() bool`

GetHasVariations returns the HasVariations field if non-nil, zero value otherwise.

### GetHasVariationsOk

`func (o *ShopListing) GetHasVariationsOk() (*bool, bool)`

GetHasVariationsOk returns a tuple with the HasVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVariations

`func (o *ShopListing) SetHasVariations(v bool)`

SetHasVariations sets HasVariations field to given value.

### HasHasVariations

`func (o *ShopListing) HasHasVariations() bool`

HasHasVariations returns a boolean if a field has been set.

### GetShouldAutoRenew

`func (o *ShopListing) GetShouldAutoRenew() bool`

GetShouldAutoRenew returns the ShouldAutoRenew field if non-nil, zero value otherwise.

### GetShouldAutoRenewOk

`func (o *ShopListing) GetShouldAutoRenewOk() (*bool, bool)`

GetShouldAutoRenewOk returns a tuple with the ShouldAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAutoRenew

`func (o *ShopListing) SetShouldAutoRenew(v bool)`

SetShouldAutoRenew sets ShouldAutoRenew field to given value.

### HasShouldAutoRenew

`func (o *ShopListing) HasShouldAutoRenew() bool`

HasShouldAutoRenew returns a boolean if a field has been set.

### GetLanguage

`func (o *ShopListing) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ShopListing) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ShopListing) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ShopListing) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *ShopListing) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *ShopListing) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetPrice

`func (o *ShopListing) GetPrice() Money`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ShopListing) GetPriceOk() (*Money, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ShopListing) SetPrice(v Money)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ShopListing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxonomyId

`func (o *ShopListing) GetTaxonomyId() int64`

GetTaxonomyId returns the TaxonomyId field if non-nil, zero value otherwise.

### GetTaxonomyIdOk

`func (o *ShopListing) GetTaxonomyIdOk() (*int64, bool)`

GetTaxonomyIdOk returns a tuple with the TaxonomyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyId

`func (o *ShopListing) SetTaxonomyId(v int64)`

SetTaxonomyId sets TaxonomyId field to given value.

### HasTaxonomyId

`func (o *ShopListing) HasTaxonomyId() bool`

HasTaxonomyId returns a boolean if a field has been set.

### SetTaxonomyIdNil

`func (o *ShopListing) SetTaxonomyIdNil(b bool)`

 SetTaxonomyIdNil sets the value for TaxonomyId to be an explicit nil

### UnsetTaxonomyId
`func (o *ShopListing) UnsetTaxonomyId()`

UnsetTaxonomyId ensures that no value is present for TaxonomyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


