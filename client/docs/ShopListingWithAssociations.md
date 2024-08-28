# ShopListingWithAssociations

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
**ShippingProfile** | Pointer to [**NullableShopShippingProfile**](ShopShippingProfile.md) | An array of data representing the shipping profile resource. | [optional] 
**User** | Pointer to [**NullableUser**](User.md) | Represents a single user of the site | [optional] 
**Shop** | Pointer to [**NullableShop**](Shop.md) | A shop created by an Etsy user. | [optional] 
**Images** | Pointer to [**[]ListingImage**](ListingImage.md) | Represents a list of listing image resources, each of which contains the reference URLs and metadata for an image | [optional] 
**Videos** | Pointer to [**[]ListingVideo**](ListingVideo.md) | The single video associated with a listing. | [optional] 
**Inventory** | Pointer to [**NullableListingInventory**](ListingInventory.md) | An enumerated string that attaches a valid association. Default value is null. | [optional] 
**ProductionPartners** | Pointer to [**[]ShopProductionPartner**](ShopProductionPartner.md) | Represents a list of production partners for a shop. | [optional] 
**Skus** | Pointer to **[]string** | A list of SKU strings for the listing. SKUs will only appear if the requesting user owns the shop and a valid matching OAuth 2 token is provided. When requested without the token it will be an empty array. | [optional] 
**Translations** | Pointer to [**NullableListingTranslations**](ListingTranslations.md) | A map of translations for the listing. Default value is a map of all supported languages keyed to null. | [optional] 
**Views** | Pointer to **int64** | The number of times the listing has been viewed. This value is tabulated once per day and **only for active listings**, so the value is not real-time. If &#x60;0&#x60;, the listing has either not been viewed, not yet tabulated, was not active during the last tabulation or there was an error fetching the value. If a value is expected, call &#x60;getListing&#x60; to confirm the value. | [optional] 

## Methods

### NewShopListingWithAssociations

`func NewShopListingWithAssociations() *ShopListingWithAssociations`

NewShopListingWithAssociations instantiates a new ShopListingWithAssociations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopListingWithAssociationsWithDefaults

`func NewShopListingWithAssociationsWithDefaults() *ShopListingWithAssociations`

NewShopListingWithAssociationsWithDefaults instantiates a new ShopListingWithAssociations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *ShopListingWithAssociations) GetListingId() int64`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ShopListingWithAssociations) GetListingIdOk() (*int64, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ShopListingWithAssociations) SetListingId(v int64)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *ShopListingWithAssociations) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### GetUserId

`func (o *ShopListingWithAssociations) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ShopListingWithAssociations) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ShopListingWithAssociations) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ShopListingWithAssociations) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetShopId

`func (o *ShopListingWithAssociations) GetShopId() int64`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ShopListingWithAssociations) GetShopIdOk() (*int64, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ShopListingWithAssociations) SetShopId(v int64)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ShopListingWithAssociations) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetTitle

`func (o *ShopListingWithAssociations) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShopListingWithAssociations) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShopListingWithAssociations) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ShopListingWithAssociations) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ShopListingWithAssociations) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShopListingWithAssociations) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShopListingWithAssociations) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShopListingWithAssociations) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *ShopListingWithAssociations) GetState() ShopListingState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ShopListingWithAssociations) GetStateOk() (*ShopListingState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ShopListingWithAssociations) SetState(v ShopListingState)`

SetState sets State field to given value.

### HasState

`func (o *ShopListingWithAssociations) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *ShopListingWithAssociations) GetCreationTimestamp() int64`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ShopListingWithAssociations) GetCreationTimestampOk() (*int64, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ShopListingWithAssociations) SetCreationTimestamp(v int64)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ShopListingWithAssociations) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ShopListingWithAssociations) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ShopListingWithAssociations) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ShopListingWithAssociations) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ShopListingWithAssociations) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetEndingTimestamp

`func (o *ShopListingWithAssociations) GetEndingTimestamp() int64`

GetEndingTimestamp returns the EndingTimestamp field if non-nil, zero value otherwise.

### GetEndingTimestampOk

`func (o *ShopListingWithAssociations) GetEndingTimestampOk() (*int64, bool)`

GetEndingTimestampOk returns a tuple with the EndingTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingTimestamp

`func (o *ShopListingWithAssociations) SetEndingTimestamp(v int64)`

SetEndingTimestamp sets EndingTimestamp field to given value.

### HasEndingTimestamp

`func (o *ShopListingWithAssociations) HasEndingTimestamp() bool`

HasEndingTimestamp returns a boolean if a field has been set.

### GetOriginalCreationTimestamp

`func (o *ShopListingWithAssociations) GetOriginalCreationTimestamp() int64`

GetOriginalCreationTimestamp returns the OriginalCreationTimestamp field if non-nil, zero value otherwise.

### GetOriginalCreationTimestampOk

`func (o *ShopListingWithAssociations) GetOriginalCreationTimestampOk() (*int64, bool)`

GetOriginalCreationTimestampOk returns a tuple with the OriginalCreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCreationTimestamp

`func (o *ShopListingWithAssociations) SetOriginalCreationTimestamp(v int64)`

SetOriginalCreationTimestamp sets OriginalCreationTimestamp field to given value.

### HasOriginalCreationTimestamp

`func (o *ShopListingWithAssociations) HasOriginalCreationTimestamp() bool`

HasOriginalCreationTimestamp returns a boolean if a field has been set.

### GetLastModifiedTimestamp

`func (o *ShopListingWithAssociations) GetLastModifiedTimestamp() int64`

GetLastModifiedTimestamp returns the LastModifiedTimestamp field if non-nil, zero value otherwise.

### GetLastModifiedTimestampOk

`func (o *ShopListingWithAssociations) GetLastModifiedTimestampOk() (*int64, bool)`

GetLastModifiedTimestampOk returns a tuple with the LastModifiedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimestamp

`func (o *ShopListingWithAssociations) SetLastModifiedTimestamp(v int64)`

SetLastModifiedTimestamp sets LastModifiedTimestamp field to given value.

### HasLastModifiedTimestamp

`func (o *ShopListingWithAssociations) HasLastModifiedTimestamp() bool`

HasLastModifiedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ShopListingWithAssociations) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ShopListingWithAssociations) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ShopListingWithAssociations) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *ShopListingWithAssociations) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetStateTimestamp

`func (o *ShopListingWithAssociations) GetStateTimestamp() int64`

GetStateTimestamp returns the StateTimestamp field if non-nil, zero value otherwise.

### GetStateTimestampOk

`func (o *ShopListingWithAssociations) GetStateTimestampOk() (*int64, bool)`

GetStateTimestampOk returns a tuple with the StateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTimestamp

`func (o *ShopListingWithAssociations) SetStateTimestamp(v int64)`

SetStateTimestamp sets StateTimestamp field to given value.

### HasStateTimestamp

`func (o *ShopListingWithAssociations) HasStateTimestamp() bool`

HasStateTimestamp returns a boolean if a field has been set.

### SetStateTimestampNil

`func (o *ShopListingWithAssociations) SetStateTimestampNil(b bool)`

 SetStateTimestampNil sets the value for StateTimestamp to be an explicit nil

### UnsetStateTimestamp
`func (o *ShopListingWithAssociations) UnsetStateTimestamp()`

UnsetStateTimestamp ensures that no value is present for StateTimestamp, not even an explicit nil
### GetQuantity

`func (o *ShopListingWithAssociations) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShopListingWithAssociations) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShopListingWithAssociations) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ShopListingWithAssociations) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetShopSectionId

`func (o *ShopListingWithAssociations) GetShopSectionId() int64`

GetShopSectionId returns the ShopSectionId field if non-nil, zero value otherwise.

### GetShopSectionIdOk

`func (o *ShopListingWithAssociations) GetShopSectionIdOk() (*int64, bool)`

GetShopSectionIdOk returns a tuple with the ShopSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopSectionId

`func (o *ShopListingWithAssociations) SetShopSectionId(v int64)`

SetShopSectionId sets ShopSectionId field to given value.

### HasShopSectionId

`func (o *ShopListingWithAssociations) HasShopSectionId() bool`

HasShopSectionId returns a boolean if a field has been set.

### SetShopSectionIdNil

`func (o *ShopListingWithAssociations) SetShopSectionIdNil(b bool)`

 SetShopSectionIdNil sets the value for ShopSectionId to be an explicit nil

### UnsetShopSectionId
`func (o *ShopListingWithAssociations) UnsetShopSectionId()`

UnsetShopSectionId ensures that no value is present for ShopSectionId, not even an explicit nil
### GetFeaturedRank

`func (o *ShopListingWithAssociations) GetFeaturedRank() int64`

GetFeaturedRank returns the FeaturedRank field if non-nil, zero value otherwise.

### GetFeaturedRankOk

`func (o *ShopListingWithAssociations) GetFeaturedRankOk() (*int64, bool)`

GetFeaturedRankOk returns a tuple with the FeaturedRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedRank

`func (o *ShopListingWithAssociations) SetFeaturedRank(v int64)`

SetFeaturedRank sets FeaturedRank field to given value.

### HasFeaturedRank

`func (o *ShopListingWithAssociations) HasFeaturedRank() bool`

HasFeaturedRank returns a boolean if a field has been set.

### GetUrl

`func (o *ShopListingWithAssociations) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ShopListingWithAssociations) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ShopListingWithAssociations) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ShopListingWithAssociations) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNumFavorers

`func (o *ShopListingWithAssociations) GetNumFavorers() int64`

GetNumFavorers returns the NumFavorers field if non-nil, zero value otherwise.

### GetNumFavorersOk

`func (o *ShopListingWithAssociations) GetNumFavorersOk() (*int64, bool)`

GetNumFavorersOk returns a tuple with the NumFavorers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFavorers

`func (o *ShopListingWithAssociations) SetNumFavorers(v int64)`

SetNumFavorers sets NumFavorers field to given value.

### HasNumFavorers

`func (o *ShopListingWithAssociations) HasNumFavorers() bool`

HasNumFavorers returns a boolean if a field has been set.

### GetNonTaxable

`func (o *ShopListingWithAssociations) GetNonTaxable() bool`

GetNonTaxable returns the NonTaxable field if non-nil, zero value otherwise.

### GetNonTaxableOk

`func (o *ShopListingWithAssociations) GetNonTaxableOk() (*bool, bool)`

GetNonTaxableOk returns a tuple with the NonTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTaxable

`func (o *ShopListingWithAssociations) SetNonTaxable(v bool)`

SetNonTaxable sets NonTaxable field to given value.

### HasNonTaxable

`func (o *ShopListingWithAssociations) HasNonTaxable() bool`

HasNonTaxable returns a boolean if a field has been set.

### GetIsTaxable

`func (o *ShopListingWithAssociations) GetIsTaxable() bool`

GetIsTaxable returns the IsTaxable field if non-nil, zero value otherwise.

### GetIsTaxableOk

`func (o *ShopListingWithAssociations) GetIsTaxableOk() (*bool, bool)`

GetIsTaxableOk returns a tuple with the IsTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxable

`func (o *ShopListingWithAssociations) SetIsTaxable(v bool)`

SetIsTaxable sets IsTaxable field to given value.

### HasIsTaxable

`func (o *ShopListingWithAssociations) HasIsTaxable() bool`

HasIsTaxable returns a boolean if a field has been set.

### GetIsCustomizable

`func (o *ShopListingWithAssociations) GetIsCustomizable() bool`

GetIsCustomizable returns the IsCustomizable field if non-nil, zero value otherwise.

### GetIsCustomizableOk

`func (o *ShopListingWithAssociations) GetIsCustomizableOk() (*bool, bool)`

GetIsCustomizableOk returns a tuple with the IsCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomizable

`func (o *ShopListingWithAssociations) SetIsCustomizable(v bool)`

SetIsCustomizable sets IsCustomizable field to given value.

### HasIsCustomizable

`func (o *ShopListingWithAssociations) HasIsCustomizable() bool`

HasIsCustomizable returns a boolean if a field has been set.

### GetIsPersonalizable

`func (o *ShopListingWithAssociations) GetIsPersonalizable() bool`

GetIsPersonalizable returns the IsPersonalizable field if non-nil, zero value otherwise.

### GetIsPersonalizableOk

`func (o *ShopListingWithAssociations) GetIsPersonalizableOk() (*bool, bool)`

GetIsPersonalizableOk returns a tuple with the IsPersonalizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonalizable

`func (o *ShopListingWithAssociations) SetIsPersonalizable(v bool)`

SetIsPersonalizable sets IsPersonalizable field to given value.

### HasIsPersonalizable

`func (o *ShopListingWithAssociations) HasIsPersonalizable() bool`

HasIsPersonalizable returns a boolean if a field has been set.

### GetPersonalizationIsRequired

`func (o *ShopListingWithAssociations) GetPersonalizationIsRequired() bool`

GetPersonalizationIsRequired returns the PersonalizationIsRequired field if non-nil, zero value otherwise.

### GetPersonalizationIsRequiredOk

`func (o *ShopListingWithAssociations) GetPersonalizationIsRequiredOk() (*bool, bool)`

GetPersonalizationIsRequiredOk returns a tuple with the PersonalizationIsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationIsRequired

`func (o *ShopListingWithAssociations) SetPersonalizationIsRequired(v bool)`

SetPersonalizationIsRequired sets PersonalizationIsRequired field to given value.

### HasPersonalizationIsRequired

`func (o *ShopListingWithAssociations) HasPersonalizationIsRequired() bool`

HasPersonalizationIsRequired returns a boolean if a field has been set.

### GetPersonalizationCharCountMax

`func (o *ShopListingWithAssociations) GetPersonalizationCharCountMax() int64`

GetPersonalizationCharCountMax returns the PersonalizationCharCountMax field if non-nil, zero value otherwise.

### GetPersonalizationCharCountMaxOk

`func (o *ShopListingWithAssociations) GetPersonalizationCharCountMaxOk() (*int64, bool)`

GetPersonalizationCharCountMaxOk returns a tuple with the PersonalizationCharCountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationCharCountMax

`func (o *ShopListingWithAssociations) SetPersonalizationCharCountMax(v int64)`

SetPersonalizationCharCountMax sets PersonalizationCharCountMax field to given value.

### HasPersonalizationCharCountMax

`func (o *ShopListingWithAssociations) HasPersonalizationCharCountMax() bool`

HasPersonalizationCharCountMax returns a boolean if a field has been set.

### SetPersonalizationCharCountMaxNil

`func (o *ShopListingWithAssociations) SetPersonalizationCharCountMaxNil(b bool)`

 SetPersonalizationCharCountMaxNil sets the value for PersonalizationCharCountMax to be an explicit nil

### UnsetPersonalizationCharCountMax
`func (o *ShopListingWithAssociations) UnsetPersonalizationCharCountMax()`

UnsetPersonalizationCharCountMax ensures that no value is present for PersonalizationCharCountMax, not even an explicit nil
### GetPersonalizationInstructions

`func (o *ShopListingWithAssociations) GetPersonalizationInstructions() string`

GetPersonalizationInstructions returns the PersonalizationInstructions field if non-nil, zero value otherwise.

### GetPersonalizationInstructionsOk

`func (o *ShopListingWithAssociations) GetPersonalizationInstructionsOk() (*string, bool)`

GetPersonalizationInstructionsOk returns a tuple with the PersonalizationInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationInstructions

`func (o *ShopListingWithAssociations) SetPersonalizationInstructions(v string)`

SetPersonalizationInstructions sets PersonalizationInstructions field to given value.

### HasPersonalizationInstructions

`func (o *ShopListingWithAssociations) HasPersonalizationInstructions() bool`

HasPersonalizationInstructions returns a boolean if a field has been set.

### SetPersonalizationInstructionsNil

`func (o *ShopListingWithAssociations) SetPersonalizationInstructionsNil(b bool)`

 SetPersonalizationInstructionsNil sets the value for PersonalizationInstructions to be an explicit nil

### UnsetPersonalizationInstructions
`func (o *ShopListingWithAssociations) UnsetPersonalizationInstructions()`

UnsetPersonalizationInstructions ensures that no value is present for PersonalizationInstructions, not even an explicit nil
### GetListingType

`func (o *ShopListingWithAssociations) GetListingType() CreateDraftListingRequestType`

GetListingType returns the ListingType field if non-nil, zero value otherwise.

### GetListingTypeOk

`func (o *ShopListingWithAssociations) GetListingTypeOk() (*CreateDraftListingRequestType, bool)`

GetListingTypeOk returns a tuple with the ListingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingType

`func (o *ShopListingWithAssociations) SetListingType(v CreateDraftListingRequestType)`

SetListingType sets ListingType field to given value.

### HasListingType

`func (o *ShopListingWithAssociations) HasListingType() bool`

HasListingType returns a boolean if a field has been set.

### GetTags

`func (o *ShopListingWithAssociations) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ShopListingWithAssociations) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ShopListingWithAssociations) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ShopListingWithAssociations) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMaterials

`func (o *ShopListingWithAssociations) GetMaterials() []string`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *ShopListingWithAssociations) GetMaterialsOk() (*[]string, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *ShopListingWithAssociations) SetMaterials(v []string)`

SetMaterials sets Materials field to given value.

### HasMaterials

`func (o *ShopListingWithAssociations) HasMaterials() bool`

HasMaterials returns a boolean if a field has been set.

### GetShippingProfileId

`func (o *ShopListingWithAssociations) GetShippingProfileId() int64`

GetShippingProfileId returns the ShippingProfileId field if non-nil, zero value otherwise.

### GetShippingProfileIdOk

`func (o *ShopListingWithAssociations) GetShippingProfileIdOk() (*int64, bool)`

GetShippingProfileIdOk returns a tuple with the ShippingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileId

`func (o *ShopListingWithAssociations) SetShippingProfileId(v int64)`

SetShippingProfileId sets ShippingProfileId field to given value.

### HasShippingProfileId

`func (o *ShopListingWithAssociations) HasShippingProfileId() bool`

HasShippingProfileId returns a boolean if a field has been set.

### SetShippingProfileIdNil

`func (o *ShopListingWithAssociations) SetShippingProfileIdNil(b bool)`

 SetShippingProfileIdNil sets the value for ShippingProfileId to be an explicit nil

### UnsetShippingProfileId
`func (o *ShopListingWithAssociations) UnsetShippingProfileId()`

UnsetShippingProfileId ensures that no value is present for ShippingProfileId, not even an explicit nil
### GetReturnPolicyId

`func (o *ShopListingWithAssociations) GetReturnPolicyId() int64`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *ShopListingWithAssociations) GetReturnPolicyIdOk() (*int64, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *ShopListingWithAssociations) SetReturnPolicyId(v int64)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *ShopListingWithAssociations) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### SetReturnPolicyIdNil

`func (o *ShopListingWithAssociations) SetReturnPolicyIdNil(b bool)`

 SetReturnPolicyIdNil sets the value for ReturnPolicyId to be an explicit nil

### UnsetReturnPolicyId
`func (o *ShopListingWithAssociations) UnsetReturnPolicyId()`

UnsetReturnPolicyId ensures that no value is present for ReturnPolicyId, not even an explicit nil
### GetProcessingMin

`func (o *ShopListingWithAssociations) GetProcessingMin() int64`

GetProcessingMin returns the ProcessingMin field if non-nil, zero value otherwise.

### GetProcessingMinOk

`func (o *ShopListingWithAssociations) GetProcessingMinOk() (*int64, bool)`

GetProcessingMinOk returns a tuple with the ProcessingMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMin

`func (o *ShopListingWithAssociations) SetProcessingMin(v int64)`

SetProcessingMin sets ProcessingMin field to given value.

### HasProcessingMin

`func (o *ShopListingWithAssociations) HasProcessingMin() bool`

HasProcessingMin returns a boolean if a field has been set.

### SetProcessingMinNil

`func (o *ShopListingWithAssociations) SetProcessingMinNil(b bool)`

 SetProcessingMinNil sets the value for ProcessingMin to be an explicit nil

### UnsetProcessingMin
`func (o *ShopListingWithAssociations) UnsetProcessingMin()`

UnsetProcessingMin ensures that no value is present for ProcessingMin, not even an explicit nil
### GetProcessingMax

`func (o *ShopListingWithAssociations) GetProcessingMax() int64`

GetProcessingMax returns the ProcessingMax field if non-nil, zero value otherwise.

### GetProcessingMaxOk

`func (o *ShopListingWithAssociations) GetProcessingMaxOk() (*int64, bool)`

GetProcessingMaxOk returns a tuple with the ProcessingMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMax

`func (o *ShopListingWithAssociations) SetProcessingMax(v int64)`

SetProcessingMax sets ProcessingMax field to given value.

### HasProcessingMax

`func (o *ShopListingWithAssociations) HasProcessingMax() bool`

HasProcessingMax returns a boolean if a field has been set.

### SetProcessingMaxNil

`func (o *ShopListingWithAssociations) SetProcessingMaxNil(b bool)`

 SetProcessingMaxNil sets the value for ProcessingMax to be an explicit nil

### UnsetProcessingMax
`func (o *ShopListingWithAssociations) UnsetProcessingMax()`

UnsetProcessingMax ensures that no value is present for ProcessingMax, not even an explicit nil
### GetWhoMade

`func (o *ShopListingWithAssociations) GetWhoMade() ShopListingWhoMade`

GetWhoMade returns the WhoMade field if non-nil, zero value otherwise.

### GetWhoMadeOk

`func (o *ShopListingWithAssociations) GetWhoMadeOk() (*ShopListingWhoMade, bool)`

GetWhoMadeOk returns a tuple with the WhoMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoMade

`func (o *ShopListingWithAssociations) SetWhoMade(v ShopListingWhoMade)`

SetWhoMade sets WhoMade field to given value.

### HasWhoMade

`func (o *ShopListingWithAssociations) HasWhoMade() bool`

HasWhoMade returns a boolean if a field has been set.

### SetWhoMadeNil

`func (o *ShopListingWithAssociations) SetWhoMadeNil(b bool)`

 SetWhoMadeNil sets the value for WhoMade to be an explicit nil

### UnsetWhoMade
`func (o *ShopListingWithAssociations) UnsetWhoMade()`

UnsetWhoMade ensures that no value is present for WhoMade, not even an explicit nil
### GetWhenMade

`func (o *ShopListingWithAssociations) GetWhenMade() ShopListingWhenMade`

GetWhenMade returns the WhenMade field if non-nil, zero value otherwise.

### GetWhenMadeOk

`func (o *ShopListingWithAssociations) GetWhenMadeOk() (*ShopListingWhenMade, bool)`

GetWhenMadeOk returns a tuple with the WhenMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenMade

`func (o *ShopListingWithAssociations) SetWhenMade(v ShopListingWhenMade)`

SetWhenMade sets WhenMade field to given value.

### HasWhenMade

`func (o *ShopListingWithAssociations) HasWhenMade() bool`

HasWhenMade returns a boolean if a field has been set.

### SetWhenMadeNil

`func (o *ShopListingWithAssociations) SetWhenMadeNil(b bool)`

 SetWhenMadeNil sets the value for WhenMade to be an explicit nil

### UnsetWhenMade
`func (o *ShopListingWithAssociations) UnsetWhenMade()`

UnsetWhenMade ensures that no value is present for WhenMade, not even an explicit nil
### GetIsSupply

`func (o *ShopListingWithAssociations) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *ShopListingWithAssociations) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *ShopListingWithAssociations) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.

### HasIsSupply

`func (o *ShopListingWithAssociations) HasIsSupply() bool`

HasIsSupply returns a boolean if a field has been set.

### SetIsSupplyNil

`func (o *ShopListingWithAssociations) SetIsSupplyNil(b bool)`

 SetIsSupplyNil sets the value for IsSupply to be an explicit nil

### UnsetIsSupply
`func (o *ShopListingWithAssociations) UnsetIsSupply()`

UnsetIsSupply ensures that no value is present for IsSupply, not even an explicit nil
### GetItemWeight

`func (o *ShopListingWithAssociations) GetItemWeight() float32`

GetItemWeight returns the ItemWeight field if non-nil, zero value otherwise.

### GetItemWeightOk

`func (o *ShopListingWithAssociations) GetItemWeightOk() (*float32, bool)`

GetItemWeightOk returns a tuple with the ItemWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWeight

`func (o *ShopListingWithAssociations) SetItemWeight(v float32)`

SetItemWeight sets ItemWeight field to given value.

### HasItemWeight

`func (o *ShopListingWithAssociations) HasItemWeight() bool`

HasItemWeight returns a boolean if a field has been set.

### SetItemWeightNil

`func (o *ShopListingWithAssociations) SetItemWeightNil(b bool)`

 SetItemWeightNil sets the value for ItemWeight to be an explicit nil

### UnsetItemWeight
`func (o *ShopListingWithAssociations) UnsetItemWeight()`

UnsetItemWeight ensures that no value is present for ItemWeight, not even an explicit nil
### GetItemWeightUnit

`func (o *ShopListingWithAssociations) GetItemWeightUnit() CreateDraftListingRequestItemWeightUnit`

GetItemWeightUnit returns the ItemWeightUnit field if non-nil, zero value otherwise.

### GetItemWeightUnitOk

`func (o *ShopListingWithAssociations) GetItemWeightUnitOk() (*CreateDraftListingRequestItemWeightUnit, bool)`

GetItemWeightUnitOk returns a tuple with the ItemWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWeightUnit

`func (o *ShopListingWithAssociations) SetItemWeightUnit(v CreateDraftListingRequestItemWeightUnit)`

SetItemWeightUnit sets ItemWeightUnit field to given value.

### HasItemWeightUnit

`func (o *ShopListingWithAssociations) HasItemWeightUnit() bool`

HasItemWeightUnit returns a boolean if a field has been set.

### SetItemWeightUnitNil

`func (o *ShopListingWithAssociations) SetItemWeightUnitNil(b bool)`

 SetItemWeightUnitNil sets the value for ItemWeightUnit to be an explicit nil

### UnsetItemWeightUnit
`func (o *ShopListingWithAssociations) UnsetItemWeightUnit()`

UnsetItemWeightUnit ensures that no value is present for ItemWeightUnit, not even an explicit nil
### GetItemLength

`func (o *ShopListingWithAssociations) GetItemLength() float32`

GetItemLength returns the ItemLength field if non-nil, zero value otherwise.

### GetItemLengthOk

`func (o *ShopListingWithAssociations) GetItemLengthOk() (*float32, bool)`

GetItemLengthOk returns a tuple with the ItemLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLength

`func (o *ShopListingWithAssociations) SetItemLength(v float32)`

SetItemLength sets ItemLength field to given value.

### HasItemLength

`func (o *ShopListingWithAssociations) HasItemLength() bool`

HasItemLength returns a boolean if a field has been set.

### SetItemLengthNil

`func (o *ShopListingWithAssociations) SetItemLengthNil(b bool)`

 SetItemLengthNil sets the value for ItemLength to be an explicit nil

### UnsetItemLength
`func (o *ShopListingWithAssociations) UnsetItemLength()`

UnsetItemLength ensures that no value is present for ItemLength, not even an explicit nil
### GetItemWidth

`func (o *ShopListingWithAssociations) GetItemWidth() float32`

GetItemWidth returns the ItemWidth field if non-nil, zero value otherwise.

### GetItemWidthOk

`func (o *ShopListingWithAssociations) GetItemWidthOk() (*float32, bool)`

GetItemWidthOk returns a tuple with the ItemWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWidth

`func (o *ShopListingWithAssociations) SetItemWidth(v float32)`

SetItemWidth sets ItemWidth field to given value.

### HasItemWidth

`func (o *ShopListingWithAssociations) HasItemWidth() bool`

HasItemWidth returns a boolean if a field has been set.

### SetItemWidthNil

`func (o *ShopListingWithAssociations) SetItemWidthNil(b bool)`

 SetItemWidthNil sets the value for ItemWidth to be an explicit nil

### UnsetItemWidth
`func (o *ShopListingWithAssociations) UnsetItemWidth()`

UnsetItemWidth ensures that no value is present for ItemWidth, not even an explicit nil
### GetItemHeight

`func (o *ShopListingWithAssociations) GetItemHeight() float32`

GetItemHeight returns the ItemHeight field if non-nil, zero value otherwise.

### GetItemHeightOk

`func (o *ShopListingWithAssociations) GetItemHeightOk() (*float32, bool)`

GetItemHeightOk returns a tuple with the ItemHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemHeight

`func (o *ShopListingWithAssociations) SetItemHeight(v float32)`

SetItemHeight sets ItemHeight field to given value.

### HasItemHeight

`func (o *ShopListingWithAssociations) HasItemHeight() bool`

HasItemHeight returns a boolean if a field has been set.

### SetItemHeightNil

`func (o *ShopListingWithAssociations) SetItemHeightNil(b bool)`

 SetItemHeightNil sets the value for ItemHeight to be an explicit nil

### UnsetItemHeight
`func (o *ShopListingWithAssociations) UnsetItemHeight()`

UnsetItemHeight ensures that no value is present for ItemHeight, not even an explicit nil
### GetItemDimensionsUnit

`func (o *ShopListingWithAssociations) GetItemDimensionsUnit() CreateDraftListingRequestItemDimensionsUnit`

GetItemDimensionsUnit returns the ItemDimensionsUnit field if non-nil, zero value otherwise.

### GetItemDimensionsUnitOk

`func (o *ShopListingWithAssociations) GetItemDimensionsUnitOk() (*CreateDraftListingRequestItemDimensionsUnit, bool)`

GetItemDimensionsUnitOk returns a tuple with the ItemDimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDimensionsUnit

`func (o *ShopListingWithAssociations) SetItemDimensionsUnit(v CreateDraftListingRequestItemDimensionsUnit)`

SetItemDimensionsUnit sets ItemDimensionsUnit field to given value.

### HasItemDimensionsUnit

`func (o *ShopListingWithAssociations) HasItemDimensionsUnit() bool`

HasItemDimensionsUnit returns a boolean if a field has been set.

### SetItemDimensionsUnitNil

`func (o *ShopListingWithAssociations) SetItemDimensionsUnitNil(b bool)`

 SetItemDimensionsUnitNil sets the value for ItemDimensionsUnit to be an explicit nil

### UnsetItemDimensionsUnit
`func (o *ShopListingWithAssociations) UnsetItemDimensionsUnit()`

UnsetItemDimensionsUnit ensures that no value is present for ItemDimensionsUnit, not even an explicit nil
### GetIsPrivate

`func (o *ShopListingWithAssociations) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ShopListingWithAssociations) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ShopListingWithAssociations) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ShopListingWithAssociations) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetStyle

`func (o *ShopListingWithAssociations) GetStyle() []string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ShopListingWithAssociations) GetStyleOk() (*[]string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ShopListingWithAssociations) SetStyle(v []string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ShopListingWithAssociations) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetFileData

`func (o *ShopListingWithAssociations) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *ShopListingWithAssociations) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *ShopListingWithAssociations) SetFileData(v string)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *ShopListingWithAssociations) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### GetHasVariations

`func (o *ShopListingWithAssociations) GetHasVariations() bool`

GetHasVariations returns the HasVariations field if non-nil, zero value otherwise.

### GetHasVariationsOk

`func (o *ShopListingWithAssociations) GetHasVariationsOk() (*bool, bool)`

GetHasVariationsOk returns a tuple with the HasVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVariations

`func (o *ShopListingWithAssociations) SetHasVariations(v bool)`

SetHasVariations sets HasVariations field to given value.

### HasHasVariations

`func (o *ShopListingWithAssociations) HasHasVariations() bool`

HasHasVariations returns a boolean if a field has been set.

### GetShouldAutoRenew

`func (o *ShopListingWithAssociations) GetShouldAutoRenew() bool`

GetShouldAutoRenew returns the ShouldAutoRenew field if non-nil, zero value otherwise.

### GetShouldAutoRenewOk

`func (o *ShopListingWithAssociations) GetShouldAutoRenewOk() (*bool, bool)`

GetShouldAutoRenewOk returns a tuple with the ShouldAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAutoRenew

`func (o *ShopListingWithAssociations) SetShouldAutoRenew(v bool)`

SetShouldAutoRenew sets ShouldAutoRenew field to given value.

### HasShouldAutoRenew

`func (o *ShopListingWithAssociations) HasShouldAutoRenew() bool`

HasShouldAutoRenew returns a boolean if a field has been set.

### GetLanguage

`func (o *ShopListingWithAssociations) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ShopListingWithAssociations) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ShopListingWithAssociations) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ShopListingWithAssociations) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *ShopListingWithAssociations) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *ShopListingWithAssociations) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetPrice

`func (o *ShopListingWithAssociations) GetPrice() Money`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ShopListingWithAssociations) GetPriceOk() (*Money, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ShopListingWithAssociations) SetPrice(v Money)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ShopListingWithAssociations) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxonomyId

`func (o *ShopListingWithAssociations) GetTaxonomyId() int64`

GetTaxonomyId returns the TaxonomyId field if non-nil, zero value otherwise.

### GetTaxonomyIdOk

`func (o *ShopListingWithAssociations) GetTaxonomyIdOk() (*int64, bool)`

GetTaxonomyIdOk returns a tuple with the TaxonomyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyId

`func (o *ShopListingWithAssociations) SetTaxonomyId(v int64)`

SetTaxonomyId sets TaxonomyId field to given value.

### HasTaxonomyId

`func (o *ShopListingWithAssociations) HasTaxonomyId() bool`

HasTaxonomyId returns a boolean if a field has been set.

### SetTaxonomyIdNil

`func (o *ShopListingWithAssociations) SetTaxonomyIdNil(b bool)`

 SetTaxonomyIdNil sets the value for TaxonomyId to be an explicit nil

### UnsetTaxonomyId
`func (o *ShopListingWithAssociations) UnsetTaxonomyId()`

UnsetTaxonomyId ensures that no value is present for TaxonomyId, not even an explicit nil
### GetShippingProfile

`func (o *ShopListingWithAssociations) GetShippingProfile() ShopShippingProfile`

GetShippingProfile returns the ShippingProfile field if non-nil, zero value otherwise.

### GetShippingProfileOk

`func (o *ShopListingWithAssociations) GetShippingProfileOk() (*ShopShippingProfile, bool)`

GetShippingProfileOk returns a tuple with the ShippingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfile

`func (o *ShopListingWithAssociations) SetShippingProfile(v ShopShippingProfile)`

SetShippingProfile sets ShippingProfile field to given value.

### HasShippingProfile

`func (o *ShopListingWithAssociations) HasShippingProfile() bool`

HasShippingProfile returns a boolean if a field has been set.

### SetShippingProfileNil

`func (o *ShopListingWithAssociations) SetShippingProfileNil(b bool)`

 SetShippingProfileNil sets the value for ShippingProfile to be an explicit nil

### UnsetShippingProfile
`func (o *ShopListingWithAssociations) UnsetShippingProfile()`

UnsetShippingProfile ensures that no value is present for ShippingProfile, not even an explicit nil
### GetUser

`func (o *ShopListingWithAssociations) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ShopListingWithAssociations) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ShopListingWithAssociations) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *ShopListingWithAssociations) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ShopListingWithAssociations) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ShopListingWithAssociations) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetShop

`func (o *ShopListingWithAssociations) GetShop() Shop`

GetShop returns the Shop field if non-nil, zero value otherwise.

### GetShopOk

`func (o *ShopListingWithAssociations) GetShopOk() (*Shop, bool)`

GetShopOk returns a tuple with the Shop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShop

`func (o *ShopListingWithAssociations) SetShop(v Shop)`

SetShop sets Shop field to given value.

### HasShop

`func (o *ShopListingWithAssociations) HasShop() bool`

HasShop returns a boolean if a field has been set.

### SetShopNil

`func (o *ShopListingWithAssociations) SetShopNil(b bool)`

 SetShopNil sets the value for Shop to be an explicit nil

### UnsetShop
`func (o *ShopListingWithAssociations) UnsetShop()`

UnsetShop ensures that no value is present for Shop, not even an explicit nil
### GetImages

`func (o *ShopListingWithAssociations) GetImages() []ListingImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ShopListingWithAssociations) GetImagesOk() (*[]ListingImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ShopListingWithAssociations) SetImages(v []ListingImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *ShopListingWithAssociations) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetVideos

`func (o *ShopListingWithAssociations) GetVideos() []ListingVideo`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *ShopListingWithAssociations) GetVideosOk() (*[]ListingVideo, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *ShopListingWithAssociations) SetVideos(v []ListingVideo)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *ShopListingWithAssociations) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetInventory

`func (o *ShopListingWithAssociations) GetInventory() ListingInventory`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ShopListingWithAssociations) GetInventoryOk() (*ListingInventory, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ShopListingWithAssociations) SetInventory(v ListingInventory)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *ShopListingWithAssociations) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *ShopListingWithAssociations) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *ShopListingWithAssociations) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
### GetProductionPartners

`func (o *ShopListingWithAssociations) GetProductionPartners() []ShopProductionPartner`

GetProductionPartners returns the ProductionPartners field if non-nil, zero value otherwise.

### GetProductionPartnersOk

`func (o *ShopListingWithAssociations) GetProductionPartnersOk() (*[]ShopProductionPartner, bool)`

GetProductionPartnersOk returns a tuple with the ProductionPartners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionPartners

`func (o *ShopListingWithAssociations) SetProductionPartners(v []ShopProductionPartner)`

SetProductionPartners sets ProductionPartners field to given value.

### HasProductionPartners

`func (o *ShopListingWithAssociations) HasProductionPartners() bool`

HasProductionPartners returns a boolean if a field has been set.

### GetSkus

`func (o *ShopListingWithAssociations) GetSkus() []string`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *ShopListingWithAssociations) GetSkusOk() (*[]string, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *ShopListingWithAssociations) SetSkus(v []string)`

SetSkus sets Skus field to given value.

### HasSkus

`func (o *ShopListingWithAssociations) HasSkus() bool`

HasSkus returns a boolean if a field has been set.

### GetTranslations

`func (o *ShopListingWithAssociations) GetTranslations() ListingTranslations`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *ShopListingWithAssociations) GetTranslationsOk() (*ListingTranslations, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *ShopListingWithAssociations) SetTranslations(v ListingTranslations)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *ShopListingWithAssociations) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### SetTranslationsNil

`func (o *ShopListingWithAssociations) SetTranslationsNil(b bool)`

 SetTranslationsNil sets the value for Translations to be an explicit nil

### UnsetTranslations
`func (o *ShopListingWithAssociations) UnsetTranslations()`

UnsetTranslations ensures that no value is present for Translations, not even an explicit nil
### GetViews

`func (o *ShopListingWithAssociations) GetViews() int64`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *ShopListingWithAssociations) GetViewsOk() (*int64, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *ShopListingWithAssociations) SetViews(v int64)`

SetViews sets Views field to given value.

### HasViews

`func (o *ShopListingWithAssociations) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


