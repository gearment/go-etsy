# UpdateListingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageIds** | Pointer to **[]int64** | An array of numeric image IDs of the images in a listing, which can include up to 20 images. | [optional] 
**Title** | Pointer to **string** | The listing&#39;s title string. When creating or updating a listing, valid title strings contain only letters, numbers, punctuation marks, mathematical symbols, whitespace characters, ™, ©, and ®. (regex: /[^\\p{L}\\p{Nd}\\p{P}\\p{Sm}\\p{Zs}™©®]/u) You can only use the %, :, &amp; and + characters once each. | [optional] 
**Description** | Pointer to **string** | A description string of the product for sale in the listing. | [optional] 
**Materials** | Pointer to **[]string** | A list of material strings for materials used in the product. Valid materials strings contain only letters, numbers, and whitespace characters. (regex: /[^\\p{L}\\p{Nd}\\p{Zs}]/u) Default value is null. | [optional] 
**ShouldAutoRenew** | Pointer to **bool** | When true, renews a listing for four months upon expiration. | [optional] 
**ShippingProfileId** | Pointer to **NullableInt64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | [optional] 
**ReturnPolicyId** | Pointer to **NullableInt64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). Required for active physical listings. This requirement does not apply to listings of EU-based shops. | [optional] 
**ShopSectionId** | Pointer to **NullableInt64** | The numeric ID of the [shop section](/documentation/reference#tag/Shop-Section) for this listing. Default value is null. | [optional] 
**ItemWeight** | Pointer to **NullableFloat32** | The numeric weight of the product measured in units set in &#39;item_weight_unit&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemLength** | Pointer to **NullableFloat32** | The numeric length of the product measured in units set in &#39;item_dimensions_unit&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemWidth** | Pointer to **NullableFloat32** | The numeric width of the product measured in units set in &#39;item_dimensions_unit&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemHeight** | Pointer to **NullableFloat32** | The numeric height of the product measured in units set in &#39;item_dimensions_unit&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemWeightUnit** | Pointer to [**NullableUpdateListingRequestItemWeightUnit**](UpdateListingRequestItemWeightUnit.md) |  | [optional] 
**ItemDimensionsUnit** | Pointer to [**NullableUpdateListingRequestItemDimensionsUnit**](UpdateListingRequestItemDimensionsUnit.md) |  | [optional] 
**IsTaxable** | Pointer to **bool** | When true, applicable [shop](/documentation/reference#tag/Shop) tax rates apply to this listing at checkout. | [optional] 
**TaxonomyId** | Pointer to **int64** | The numerical taxonomy ID of the listing. See [SellerTaxonomy](/documentation/reference#tag/SellerTaxonomy) and [BuyerTaxonomy](/documentation/reference#tag/BuyerTaxonomy) for more information. | [optional] 
**Tags** | Pointer to **[]string** | A comma-separated list of tag strings for the listing. When creating or updating a listing, valid tag strings contain only letters, numbers, whitespace characters, -, &#39;, ™, ©, and ®. (regex: /[^\\p{L}\\p{Nd}\\p{Zs}\\-&#39;™©®]/u) Default value is null. | [optional] 
**WhoMade** | Pointer to [**CreateDraftListingRequestWhoMade**](CreateDraftListingRequestWhoMade.md) |  | [optional] 
**WhenMade** | Pointer to [**CreateDraftListingRequestWhenMade**](CreateDraftListingRequestWhenMade.md) |  | [optional] 
**FeaturedRank** | Pointer to **NullableInt64** | The positive non-zero numeric position in the featured listings of the shop, with rank 1 listings appearing in the left-most position in featured listing on a shop&#39;s home page. | [optional] 
**IsPersonalizable** | Pointer to **bool** | When true, this listing is personalizable. The default value is false. | [optional] 
**PersonalizationIsRequired** | Pointer to **bool** | [DEPRECATED] When true, this listing requires personalization. The default value is false. NOTE: This field will be removed on Apr. 9th, 2026. See https://developers.etsy.com/documentation/tutorials/personalization-migration for migration details. | [optional] 
**PersonalizationCharCountMax** | Pointer to **int64** | [DEPRECATED] This is an integer value representing the maximum length for the personalization message entered by the buyer. Will only change if is_personalizable is &#39;true&#39;. Note: This field will be removed on Apr. 9th, 2026. See https://developers.etsy.com/documentation/tutorials/personalization-migration for migration details. | [optional] 
**PersonalizationInstructions** | Pointer to **string** | [DEPRECATED] A string representing instructions for the buyer to enter the personalization. Will only change if is_personalizable is &#39;true&#39;. Note: This field will be removed on Apr. 9th, 2026. See https://developers.etsy.com/documentation/tutorials/personalization-migration for migration details. | [optional] 
**State** | Pointer to [**UpdateListingRequestState**](UpdateListingRequestState.md) |  | [optional] 
**IsSupply** | Pointer to **bool** | When true, tags the listing as a supply product, else indicates that it&#39;s a finished product. Helps buyers locate the listing under the Supplies heading. Requires &#39;who_made&#39; and &#39;when_made&#39;. | [optional] 
**ProductionPartnerIds** | Pointer to **[]int64** | An array of unique IDs of production partner ids. | [optional] 
**Type** | Pointer to [**NullableUpdateListingRequestType**](UpdateListingRequestType.md) |  | [optional] 

## Methods

### NewUpdateListingRequest

`func NewUpdateListingRequest() *UpdateListingRequest`

NewUpdateListingRequest instantiates a new UpdateListingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListingRequestWithDefaults

`func NewUpdateListingRequestWithDefaults() *UpdateListingRequest`

NewUpdateListingRequestWithDefaults instantiates a new UpdateListingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageIds

`func (o *UpdateListingRequest) GetImageIds() []int64`

GetImageIds returns the ImageIds field if non-nil, zero value otherwise.

### GetImageIdsOk

`func (o *UpdateListingRequest) GetImageIdsOk() (*[]int64, bool)`

GetImageIdsOk returns a tuple with the ImageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIds

`func (o *UpdateListingRequest) SetImageIds(v []int64)`

SetImageIds sets ImageIds field to given value.

### HasImageIds

`func (o *UpdateListingRequest) HasImageIds() bool`

HasImageIds returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateListingRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateListingRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateListingRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateListingRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateListingRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateListingRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateListingRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateListingRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaterials

`func (o *UpdateListingRequest) GetMaterials() []string`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *UpdateListingRequest) GetMaterialsOk() (*[]string, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *UpdateListingRequest) SetMaterials(v []string)`

SetMaterials sets Materials field to given value.

### HasMaterials

`func (o *UpdateListingRequest) HasMaterials() bool`

HasMaterials returns a boolean if a field has been set.

### SetMaterialsNil

`func (o *UpdateListingRequest) SetMaterialsNil(b bool)`

 SetMaterialsNil sets the value for Materials to be an explicit nil

### UnsetMaterials
`func (o *UpdateListingRequest) UnsetMaterials()`

UnsetMaterials ensures that no value is present for Materials, not even an explicit nil
### GetShouldAutoRenew

`func (o *UpdateListingRequest) GetShouldAutoRenew() bool`

GetShouldAutoRenew returns the ShouldAutoRenew field if non-nil, zero value otherwise.

### GetShouldAutoRenewOk

`func (o *UpdateListingRequest) GetShouldAutoRenewOk() (*bool, bool)`

GetShouldAutoRenewOk returns a tuple with the ShouldAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAutoRenew

`func (o *UpdateListingRequest) SetShouldAutoRenew(v bool)`

SetShouldAutoRenew sets ShouldAutoRenew field to given value.

### HasShouldAutoRenew

`func (o *UpdateListingRequest) HasShouldAutoRenew() bool`

HasShouldAutoRenew returns a boolean if a field has been set.

### GetShippingProfileId

`func (o *UpdateListingRequest) GetShippingProfileId() int64`

GetShippingProfileId returns the ShippingProfileId field if non-nil, zero value otherwise.

### GetShippingProfileIdOk

`func (o *UpdateListingRequest) GetShippingProfileIdOk() (*int64, bool)`

GetShippingProfileIdOk returns a tuple with the ShippingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileId

`func (o *UpdateListingRequest) SetShippingProfileId(v int64)`

SetShippingProfileId sets ShippingProfileId field to given value.

### HasShippingProfileId

`func (o *UpdateListingRequest) HasShippingProfileId() bool`

HasShippingProfileId returns a boolean if a field has been set.

### SetShippingProfileIdNil

`func (o *UpdateListingRequest) SetShippingProfileIdNil(b bool)`

 SetShippingProfileIdNil sets the value for ShippingProfileId to be an explicit nil

### UnsetShippingProfileId
`func (o *UpdateListingRequest) UnsetShippingProfileId()`

UnsetShippingProfileId ensures that no value is present for ShippingProfileId, not even an explicit nil
### GetReturnPolicyId

`func (o *UpdateListingRequest) GetReturnPolicyId() int64`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *UpdateListingRequest) GetReturnPolicyIdOk() (*int64, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *UpdateListingRequest) SetReturnPolicyId(v int64)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *UpdateListingRequest) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### SetReturnPolicyIdNil

`func (o *UpdateListingRequest) SetReturnPolicyIdNil(b bool)`

 SetReturnPolicyIdNil sets the value for ReturnPolicyId to be an explicit nil

### UnsetReturnPolicyId
`func (o *UpdateListingRequest) UnsetReturnPolicyId()`

UnsetReturnPolicyId ensures that no value is present for ReturnPolicyId, not even an explicit nil
### GetShopSectionId

`func (o *UpdateListingRequest) GetShopSectionId() int64`

GetShopSectionId returns the ShopSectionId field if non-nil, zero value otherwise.

### GetShopSectionIdOk

`func (o *UpdateListingRequest) GetShopSectionIdOk() (*int64, bool)`

GetShopSectionIdOk returns a tuple with the ShopSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopSectionId

`func (o *UpdateListingRequest) SetShopSectionId(v int64)`

SetShopSectionId sets ShopSectionId field to given value.

### HasShopSectionId

`func (o *UpdateListingRequest) HasShopSectionId() bool`

HasShopSectionId returns a boolean if a field has been set.

### SetShopSectionIdNil

`func (o *UpdateListingRequest) SetShopSectionIdNil(b bool)`

 SetShopSectionIdNil sets the value for ShopSectionId to be an explicit nil

### UnsetShopSectionId
`func (o *UpdateListingRequest) UnsetShopSectionId()`

UnsetShopSectionId ensures that no value is present for ShopSectionId, not even an explicit nil
### GetItemWeight

`func (o *UpdateListingRequest) GetItemWeight() float32`

GetItemWeight returns the ItemWeight field if non-nil, zero value otherwise.

### GetItemWeightOk

`func (o *UpdateListingRequest) GetItemWeightOk() (*float32, bool)`

GetItemWeightOk returns a tuple with the ItemWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWeight

`func (o *UpdateListingRequest) SetItemWeight(v float32)`

SetItemWeight sets ItemWeight field to given value.

### HasItemWeight

`func (o *UpdateListingRequest) HasItemWeight() bool`

HasItemWeight returns a boolean if a field has been set.

### SetItemWeightNil

`func (o *UpdateListingRequest) SetItemWeightNil(b bool)`

 SetItemWeightNil sets the value for ItemWeight to be an explicit nil

### UnsetItemWeight
`func (o *UpdateListingRequest) UnsetItemWeight()`

UnsetItemWeight ensures that no value is present for ItemWeight, not even an explicit nil
### GetItemLength

`func (o *UpdateListingRequest) GetItemLength() float32`

GetItemLength returns the ItemLength field if non-nil, zero value otherwise.

### GetItemLengthOk

`func (o *UpdateListingRequest) GetItemLengthOk() (*float32, bool)`

GetItemLengthOk returns a tuple with the ItemLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLength

`func (o *UpdateListingRequest) SetItemLength(v float32)`

SetItemLength sets ItemLength field to given value.

### HasItemLength

`func (o *UpdateListingRequest) HasItemLength() bool`

HasItemLength returns a boolean if a field has been set.

### SetItemLengthNil

`func (o *UpdateListingRequest) SetItemLengthNil(b bool)`

 SetItemLengthNil sets the value for ItemLength to be an explicit nil

### UnsetItemLength
`func (o *UpdateListingRequest) UnsetItemLength()`

UnsetItemLength ensures that no value is present for ItemLength, not even an explicit nil
### GetItemWidth

`func (o *UpdateListingRequest) GetItemWidth() float32`

GetItemWidth returns the ItemWidth field if non-nil, zero value otherwise.

### GetItemWidthOk

`func (o *UpdateListingRequest) GetItemWidthOk() (*float32, bool)`

GetItemWidthOk returns a tuple with the ItemWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWidth

`func (o *UpdateListingRequest) SetItemWidth(v float32)`

SetItemWidth sets ItemWidth field to given value.

### HasItemWidth

`func (o *UpdateListingRequest) HasItemWidth() bool`

HasItemWidth returns a boolean if a field has been set.

### SetItemWidthNil

`func (o *UpdateListingRequest) SetItemWidthNil(b bool)`

 SetItemWidthNil sets the value for ItemWidth to be an explicit nil

### UnsetItemWidth
`func (o *UpdateListingRequest) UnsetItemWidth()`

UnsetItemWidth ensures that no value is present for ItemWidth, not even an explicit nil
### GetItemHeight

`func (o *UpdateListingRequest) GetItemHeight() float32`

GetItemHeight returns the ItemHeight field if non-nil, zero value otherwise.

### GetItemHeightOk

`func (o *UpdateListingRequest) GetItemHeightOk() (*float32, bool)`

GetItemHeightOk returns a tuple with the ItemHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemHeight

`func (o *UpdateListingRequest) SetItemHeight(v float32)`

SetItemHeight sets ItemHeight field to given value.

### HasItemHeight

`func (o *UpdateListingRequest) HasItemHeight() bool`

HasItemHeight returns a boolean if a field has been set.

### SetItemHeightNil

`func (o *UpdateListingRequest) SetItemHeightNil(b bool)`

 SetItemHeightNil sets the value for ItemHeight to be an explicit nil

### UnsetItemHeight
`func (o *UpdateListingRequest) UnsetItemHeight()`

UnsetItemHeight ensures that no value is present for ItemHeight, not even an explicit nil
### GetItemWeightUnit

`func (o *UpdateListingRequest) GetItemWeightUnit() UpdateListingRequestItemWeightUnit`

GetItemWeightUnit returns the ItemWeightUnit field if non-nil, zero value otherwise.

### GetItemWeightUnitOk

`func (o *UpdateListingRequest) GetItemWeightUnitOk() (*UpdateListingRequestItemWeightUnit, bool)`

GetItemWeightUnitOk returns a tuple with the ItemWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWeightUnit

`func (o *UpdateListingRequest) SetItemWeightUnit(v UpdateListingRequestItemWeightUnit)`

SetItemWeightUnit sets ItemWeightUnit field to given value.

### HasItemWeightUnit

`func (o *UpdateListingRequest) HasItemWeightUnit() bool`

HasItemWeightUnit returns a boolean if a field has been set.

### SetItemWeightUnitNil

`func (o *UpdateListingRequest) SetItemWeightUnitNil(b bool)`

 SetItemWeightUnitNil sets the value for ItemWeightUnit to be an explicit nil

### UnsetItemWeightUnit
`func (o *UpdateListingRequest) UnsetItemWeightUnit()`

UnsetItemWeightUnit ensures that no value is present for ItemWeightUnit, not even an explicit nil
### GetItemDimensionsUnit

`func (o *UpdateListingRequest) GetItemDimensionsUnit() UpdateListingRequestItemDimensionsUnit`

GetItemDimensionsUnit returns the ItemDimensionsUnit field if non-nil, zero value otherwise.

### GetItemDimensionsUnitOk

`func (o *UpdateListingRequest) GetItemDimensionsUnitOk() (*UpdateListingRequestItemDimensionsUnit, bool)`

GetItemDimensionsUnitOk returns a tuple with the ItemDimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDimensionsUnit

`func (o *UpdateListingRequest) SetItemDimensionsUnit(v UpdateListingRequestItemDimensionsUnit)`

SetItemDimensionsUnit sets ItemDimensionsUnit field to given value.

### HasItemDimensionsUnit

`func (o *UpdateListingRequest) HasItemDimensionsUnit() bool`

HasItemDimensionsUnit returns a boolean if a field has been set.

### SetItemDimensionsUnitNil

`func (o *UpdateListingRequest) SetItemDimensionsUnitNil(b bool)`

 SetItemDimensionsUnitNil sets the value for ItemDimensionsUnit to be an explicit nil

### UnsetItemDimensionsUnit
`func (o *UpdateListingRequest) UnsetItemDimensionsUnit()`

UnsetItemDimensionsUnit ensures that no value is present for ItemDimensionsUnit, not even an explicit nil
### GetIsTaxable

`func (o *UpdateListingRequest) GetIsTaxable() bool`

GetIsTaxable returns the IsTaxable field if non-nil, zero value otherwise.

### GetIsTaxableOk

`func (o *UpdateListingRequest) GetIsTaxableOk() (*bool, bool)`

GetIsTaxableOk returns a tuple with the IsTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxable

`func (o *UpdateListingRequest) SetIsTaxable(v bool)`

SetIsTaxable sets IsTaxable field to given value.

### HasIsTaxable

`func (o *UpdateListingRequest) HasIsTaxable() bool`

HasIsTaxable returns a boolean if a field has been set.

### GetTaxonomyId

`func (o *UpdateListingRequest) GetTaxonomyId() int64`

GetTaxonomyId returns the TaxonomyId field if non-nil, zero value otherwise.

### GetTaxonomyIdOk

`func (o *UpdateListingRequest) GetTaxonomyIdOk() (*int64, bool)`

GetTaxonomyIdOk returns a tuple with the TaxonomyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyId

`func (o *UpdateListingRequest) SetTaxonomyId(v int64)`

SetTaxonomyId sets TaxonomyId field to given value.

### HasTaxonomyId

`func (o *UpdateListingRequest) HasTaxonomyId() bool`

HasTaxonomyId returns a boolean if a field has been set.

### GetTags

`func (o *UpdateListingRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateListingRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateListingRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateListingRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateListingRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateListingRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetWhoMade

`func (o *UpdateListingRequest) GetWhoMade() CreateDraftListingRequestWhoMade`

GetWhoMade returns the WhoMade field if non-nil, zero value otherwise.

### GetWhoMadeOk

`func (o *UpdateListingRequest) GetWhoMadeOk() (*CreateDraftListingRequestWhoMade, bool)`

GetWhoMadeOk returns a tuple with the WhoMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoMade

`func (o *UpdateListingRequest) SetWhoMade(v CreateDraftListingRequestWhoMade)`

SetWhoMade sets WhoMade field to given value.

### HasWhoMade

`func (o *UpdateListingRequest) HasWhoMade() bool`

HasWhoMade returns a boolean if a field has been set.

### GetWhenMade

`func (o *UpdateListingRequest) GetWhenMade() CreateDraftListingRequestWhenMade`

GetWhenMade returns the WhenMade field if non-nil, zero value otherwise.

### GetWhenMadeOk

`func (o *UpdateListingRequest) GetWhenMadeOk() (*CreateDraftListingRequestWhenMade, bool)`

GetWhenMadeOk returns a tuple with the WhenMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenMade

`func (o *UpdateListingRequest) SetWhenMade(v CreateDraftListingRequestWhenMade)`

SetWhenMade sets WhenMade field to given value.

### HasWhenMade

`func (o *UpdateListingRequest) HasWhenMade() bool`

HasWhenMade returns a boolean if a field has been set.

### GetFeaturedRank

`func (o *UpdateListingRequest) GetFeaturedRank() int64`

GetFeaturedRank returns the FeaturedRank field if non-nil, zero value otherwise.

### GetFeaturedRankOk

`func (o *UpdateListingRequest) GetFeaturedRankOk() (*int64, bool)`

GetFeaturedRankOk returns a tuple with the FeaturedRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedRank

`func (o *UpdateListingRequest) SetFeaturedRank(v int64)`

SetFeaturedRank sets FeaturedRank field to given value.

### HasFeaturedRank

`func (o *UpdateListingRequest) HasFeaturedRank() bool`

HasFeaturedRank returns a boolean if a field has been set.

### SetFeaturedRankNil

`func (o *UpdateListingRequest) SetFeaturedRankNil(b bool)`

 SetFeaturedRankNil sets the value for FeaturedRank to be an explicit nil

### UnsetFeaturedRank
`func (o *UpdateListingRequest) UnsetFeaturedRank()`

UnsetFeaturedRank ensures that no value is present for FeaturedRank, not even an explicit nil
### GetIsPersonalizable

`func (o *UpdateListingRequest) GetIsPersonalizable() bool`

GetIsPersonalizable returns the IsPersonalizable field if non-nil, zero value otherwise.

### GetIsPersonalizableOk

`func (o *UpdateListingRequest) GetIsPersonalizableOk() (*bool, bool)`

GetIsPersonalizableOk returns a tuple with the IsPersonalizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonalizable

`func (o *UpdateListingRequest) SetIsPersonalizable(v bool)`

SetIsPersonalizable sets IsPersonalizable field to given value.

### HasIsPersonalizable

`func (o *UpdateListingRequest) HasIsPersonalizable() bool`

HasIsPersonalizable returns a boolean if a field has been set.

### GetPersonalizationIsRequired

`func (o *UpdateListingRequest) GetPersonalizationIsRequired() bool`

GetPersonalizationIsRequired returns the PersonalizationIsRequired field if non-nil, zero value otherwise.

### GetPersonalizationIsRequiredOk

`func (o *UpdateListingRequest) GetPersonalizationIsRequiredOk() (*bool, bool)`

GetPersonalizationIsRequiredOk returns a tuple with the PersonalizationIsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationIsRequired

`func (o *UpdateListingRequest) SetPersonalizationIsRequired(v bool)`

SetPersonalizationIsRequired sets PersonalizationIsRequired field to given value.

### HasPersonalizationIsRequired

`func (o *UpdateListingRequest) HasPersonalizationIsRequired() bool`

HasPersonalizationIsRequired returns a boolean if a field has been set.

### GetPersonalizationCharCountMax

`func (o *UpdateListingRequest) GetPersonalizationCharCountMax() int64`

GetPersonalizationCharCountMax returns the PersonalizationCharCountMax field if non-nil, zero value otherwise.

### GetPersonalizationCharCountMaxOk

`func (o *UpdateListingRequest) GetPersonalizationCharCountMaxOk() (*int64, bool)`

GetPersonalizationCharCountMaxOk returns a tuple with the PersonalizationCharCountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationCharCountMax

`func (o *UpdateListingRequest) SetPersonalizationCharCountMax(v int64)`

SetPersonalizationCharCountMax sets PersonalizationCharCountMax field to given value.

### HasPersonalizationCharCountMax

`func (o *UpdateListingRequest) HasPersonalizationCharCountMax() bool`

HasPersonalizationCharCountMax returns a boolean if a field has been set.

### GetPersonalizationInstructions

`func (o *UpdateListingRequest) GetPersonalizationInstructions() string`

GetPersonalizationInstructions returns the PersonalizationInstructions field if non-nil, zero value otherwise.

### GetPersonalizationInstructionsOk

`func (o *UpdateListingRequest) GetPersonalizationInstructionsOk() (*string, bool)`

GetPersonalizationInstructionsOk returns a tuple with the PersonalizationInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationInstructions

`func (o *UpdateListingRequest) SetPersonalizationInstructions(v string)`

SetPersonalizationInstructions sets PersonalizationInstructions field to given value.

### HasPersonalizationInstructions

`func (o *UpdateListingRequest) HasPersonalizationInstructions() bool`

HasPersonalizationInstructions returns a boolean if a field has been set.

### GetState

`func (o *UpdateListingRequest) GetState() UpdateListingRequestState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateListingRequest) GetStateOk() (*UpdateListingRequestState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateListingRequest) SetState(v UpdateListingRequestState)`

SetState sets State field to given value.

### HasState

`func (o *UpdateListingRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIsSupply

`func (o *UpdateListingRequest) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *UpdateListingRequest) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *UpdateListingRequest) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.

### HasIsSupply

`func (o *UpdateListingRequest) HasIsSupply() bool`

HasIsSupply returns a boolean if a field has been set.

### GetProductionPartnerIds

`func (o *UpdateListingRequest) GetProductionPartnerIds() []int64`

GetProductionPartnerIds returns the ProductionPartnerIds field if non-nil, zero value otherwise.

### GetProductionPartnerIdsOk

`func (o *UpdateListingRequest) GetProductionPartnerIdsOk() (*[]int64, bool)`

GetProductionPartnerIdsOk returns a tuple with the ProductionPartnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionPartnerIds

`func (o *UpdateListingRequest) SetProductionPartnerIds(v []int64)`

SetProductionPartnerIds sets ProductionPartnerIds field to given value.

### HasProductionPartnerIds

`func (o *UpdateListingRequest) HasProductionPartnerIds() bool`

HasProductionPartnerIds returns a boolean if a field has been set.

### SetProductionPartnerIdsNil

`func (o *UpdateListingRequest) SetProductionPartnerIdsNil(b bool)`

 SetProductionPartnerIdsNil sets the value for ProductionPartnerIds to be an explicit nil

### UnsetProductionPartnerIds
`func (o *UpdateListingRequest) UnsetProductionPartnerIds()`

UnsetProductionPartnerIds ensures that no value is present for ProductionPartnerIds, not even an explicit nil
### GetType

`func (o *UpdateListingRequest) GetType() UpdateListingRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateListingRequest) GetTypeOk() (*UpdateListingRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateListingRequest) SetType(v UpdateListingRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateListingRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *UpdateListingRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UpdateListingRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


