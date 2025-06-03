# CreateDraftListingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **int64** | The positive non-zero number of products available for purchase in the listing. Note: The listing quantity is the sum of available offering quantities. You can request the quantities for individual offerings from the ListingInventory resource using the [getListingInventory](/documentation/reference#operation/getListingInventory) endpoint. | 
**Title** | **string** | The listing&#39;s title string. When creating or updating a listing, valid title strings contain only letters, numbers, punctuation marks, mathematical symbols, whitespace characters, ™, ©, and ®. (regex: /[^\\p{L}\\p{Nd}\\p{P}\\p{Sm}\\p{Zs}™©®]/u) You can only use the %, :, &amp; and + characters once each. | 
**Description** | **string** | A description string of the product for sale in the listing. | 
**Price** | **float32** | The positive non-zero price of the product. (Sold product listings are private) Note: The price is the minimum possible price. The [&#x60;getListingInventory&#x60;](/documentation/reference/#operation/getListingInventory) method requests exact prices for available offerings. | 
**WhoMade** | [**CreateDraftListingRequestWhoMade**](CreateDraftListingRequestWhoMade.md) |  | 
**WhenMade** | [**CreateDraftListingRequestWhenMade**](CreateDraftListingRequestWhenMade.md) |  | 
**TaxonomyId** | **int64** | The numerical taxonomy ID of the listing. See [SellerTaxonomy](/documentation/reference#tag/SellerTaxonomy) and [BuyerTaxonomy](/documentation/reference#tag/BuyerTaxonomy) for more information. | 
**ShippingProfileId** | Pointer to **NullableInt64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | [optional] 
**ReturnPolicyId** | Pointer to **NullableInt64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). | [optional] 
**Materials** | Pointer to **[]string** | A list of material strings for materials used in the product. Valid materials strings contain only letters, numbers, and whitespace characters. (regex: /[^\\p{L}\\p{Nd}\\p{Zs}]/u) Default value is null. | [optional] 
**ShopSectionId** | Pointer to **NullableInt64** | The numeric ID of the [shop section](/documentation/reference#tag/Shop-Section) for this listing. Default value is null. | [optional] 
**ProcessingMin** | Pointer to **NullableInt64** | The minimum number of days required to process this listing. Default value is null. | [optional] 
**ProcessingMax** | Pointer to **NullableInt64** | The maximum number of days required to process this listing. Default value is null. | [optional] 
**Tags** | Pointer to **[]string** | A comma-separated list of tag strings for the listing. When creating or updating a listing, valid tag strings contain only letters, numbers, whitespace characters, -, &#39;, ™, ©, and ®. (regex: /[^\\p{L}\\p{Nd}\\p{Zs}\\-&#39;™©®]/u) Default value is null. | [optional] 
**Styles** | Pointer to **[]string** | An array of style strings for this listing, each of which is free-form text string such as \&quot;Formal\&quot;, or \&quot;Steampunk\&quot;. When creating or updating a listing, the listing may have up to two styles. Valid style strings contain only letters, numbers, and whitespace characters. (regex: /[^\\p{L}\\p{Nd}\\p{Zs}]/u) Default value is null. | [optional] 
**ItemWeight** | Pointer to **NullableFloat32** | The numeric weight of the product measured in units set in &#39;item_weight_unit&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemLength** | Pointer to **NullableFloat32** | The numeric length of the product measured in units set in &#39;item_dimensions_unit&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemWidth** | Pointer to **NullableFloat32** | The numeric width of the product measured in units set in &#39;item_dimensions_unit&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemHeight** | Pointer to **NullableFloat32** | The numeric height of the product measured in units set in &#39;item_dimensions_unit&#39;. Default value is null. If set, the value must be greater than 0. | [optional] 
**ItemWeightUnit** | Pointer to [**NullableCreateDraftListingRequestItemWeightUnit**](CreateDraftListingRequestItemWeightUnit.md) |  | [optional] 
**ItemDimensionsUnit** | Pointer to [**NullableCreateDraftListingRequestItemDimensionsUnit**](CreateDraftListingRequestItemDimensionsUnit.md) |  | [optional] 
**IsPersonalizable** | Pointer to **bool** | When true, this listing is personalizable. The default value is null. | [optional] 
**PersonalizationIsRequired** | Pointer to **bool** | When true, this listing requires personalization. The default value is null. Will only change if is_personalizable is &#39;true&#39;. | [optional] 
**PersonalizationCharCountMax** | Pointer to **int64** | This is an integer value representing the maximum length for the personalization message entered by the buyer. Will only change if is_personalizable is &#39;true&#39;. | [optional] 
**PersonalizationInstructions** | Pointer to **string** | A string representing instructions for the buyer to enter the personalization. Will only change if is_personalizable is &#39;true&#39;. | [optional] 
**ProductionPartnerIds** | Pointer to **[]int64** | An array of unique IDs of production partner ids. | [optional] 
**ImageIds** | Pointer to **[]int64** | An array of numeric image IDs of the images in a listing, which can include up to 10 images. | [optional] 
**IsSupply** | Pointer to **bool** | When true, tags the listing as a supply product, else indicates that it&#39;s a finished product. Helps buyers locate the listing under the Supplies heading. Requires &#39;who_made&#39; and &#39;when_made&#39;. | [optional] 
**IsCustomizable** | Pointer to **bool** | When true, a buyer may contact the seller for a customized order. The default value is true when a shop accepts custom orders. Does not apply to shops that do not accept custom orders. | [optional] 
**ShouldAutoRenew** | Pointer to **bool** | When true, renews a listing for four months upon expiration. | [optional] 
**IsTaxable** | Pointer to **bool** | When true, applicable [shop](/documentation/reference#tag/Shop) tax rates apply to this listing at checkout. | [optional] 
**Type** | Pointer to [**CreateDraftListingRequestType**](CreateDraftListingRequestType.md) |  | [optional] 

## Methods

### NewCreateDraftListingRequest

`func NewCreateDraftListingRequest(quantity int64, title string, description string, price float32, whoMade CreateDraftListingRequestWhoMade, whenMade CreateDraftListingRequestWhenMade, taxonomyId int64, ) *CreateDraftListingRequest`

NewCreateDraftListingRequest instantiates a new CreateDraftListingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDraftListingRequestWithDefaults

`func NewCreateDraftListingRequestWithDefaults() *CreateDraftListingRequest`

NewCreateDraftListingRequestWithDefaults instantiates a new CreateDraftListingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *CreateDraftListingRequest) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateDraftListingRequest) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateDraftListingRequest) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetTitle

`func (o *CreateDraftListingRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateDraftListingRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateDraftListingRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CreateDraftListingRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDraftListingRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDraftListingRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPrice

`func (o *CreateDraftListingRequest) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateDraftListingRequest) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateDraftListingRequest) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetWhoMade

`func (o *CreateDraftListingRequest) GetWhoMade() CreateDraftListingRequestWhoMade`

GetWhoMade returns the WhoMade field if non-nil, zero value otherwise.

### GetWhoMadeOk

`func (o *CreateDraftListingRequest) GetWhoMadeOk() (*CreateDraftListingRequestWhoMade, bool)`

GetWhoMadeOk returns a tuple with the WhoMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoMade

`func (o *CreateDraftListingRequest) SetWhoMade(v CreateDraftListingRequestWhoMade)`

SetWhoMade sets WhoMade field to given value.


### GetWhenMade

`func (o *CreateDraftListingRequest) GetWhenMade() CreateDraftListingRequestWhenMade`

GetWhenMade returns the WhenMade field if non-nil, zero value otherwise.

### GetWhenMadeOk

`func (o *CreateDraftListingRequest) GetWhenMadeOk() (*CreateDraftListingRequestWhenMade, bool)`

GetWhenMadeOk returns a tuple with the WhenMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenMade

`func (o *CreateDraftListingRequest) SetWhenMade(v CreateDraftListingRequestWhenMade)`

SetWhenMade sets WhenMade field to given value.


### GetTaxonomyId

`func (o *CreateDraftListingRequest) GetTaxonomyId() int64`

GetTaxonomyId returns the TaxonomyId field if non-nil, zero value otherwise.

### GetTaxonomyIdOk

`func (o *CreateDraftListingRequest) GetTaxonomyIdOk() (*int64, bool)`

GetTaxonomyIdOk returns a tuple with the TaxonomyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyId

`func (o *CreateDraftListingRequest) SetTaxonomyId(v int64)`

SetTaxonomyId sets TaxonomyId field to given value.


### GetShippingProfileId

`func (o *CreateDraftListingRequest) GetShippingProfileId() int64`

GetShippingProfileId returns the ShippingProfileId field if non-nil, zero value otherwise.

### GetShippingProfileIdOk

`func (o *CreateDraftListingRequest) GetShippingProfileIdOk() (*int64, bool)`

GetShippingProfileIdOk returns a tuple with the ShippingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileId

`func (o *CreateDraftListingRequest) SetShippingProfileId(v int64)`

SetShippingProfileId sets ShippingProfileId field to given value.

### HasShippingProfileId

`func (o *CreateDraftListingRequest) HasShippingProfileId() bool`

HasShippingProfileId returns a boolean if a field has been set.

### SetShippingProfileIdNil

`func (o *CreateDraftListingRequest) SetShippingProfileIdNil(b bool)`

 SetShippingProfileIdNil sets the value for ShippingProfileId to be an explicit nil

### UnsetShippingProfileId
`func (o *CreateDraftListingRequest) UnsetShippingProfileId()`

UnsetShippingProfileId ensures that no value is present for ShippingProfileId, not even an explicit nil
### GetReturnPolicyId

`func (o *CreateDraftListingRequest) GetReturnPolicyId() int64`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *CreateDraftListingRequest) GetReturnPolicyIdOk() (*int64, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *CreateDraftListingRequest) SetReturnPolicyId(v int64)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *CreateDraftListingRequest) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### SetReturnPolicyIdNil

`func (o *CreateDraftListingRequest) SetReturnPolicyIdNil(b bool)`

 SetReturnPolicyIdNil sets the value for ReturnPolicyId to be an explicit nil

### UnsetReturnPolicyId
`func (o *CreateDraftListingRequest) UnsetReturnPolicyId()`

UnsetReturnPolicyId ensures that no value is present for ReturnPolicyId, not even an explicit nil
### GetMaterials

`func (o *CreateDraftListingRequest) GetMaterials() []string`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *CreateDraftListingRequest) GetMaterialsOk() (*[]string, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *CreateDraftListingRequest) SetMaterials(v []string)`

SetMaterials sets Materials field to given value.

### HasMaterials

`func (o *CreateDraftListingRequest) HasMaterials() bool`

HasMaterials returns a boolean if a field has been set.

### SetMaterialsNil

`func (o *CreateDraftListingRequest) SetMaterialsNil(b bool)`

 SetMaterialsNil sets the value for Materials to be an explicit nil

### UnsetMaterials
`func (o *CreateDraftListingRequest) UnsetMaterials()`

UnsetMaterials ensures that no value is present for Materials, not even an explicit nil
### GetShopSectionId

`func (o *CreateDraftListingRequest) GetShopSectionId() int64`

GetShopSectionId returns the ShopSectionId field if non-nil, zero value otherwise.

### GetShopSectionIdOk

`func (o *CreateDraftListingRequest) GetShopSectionIdOk() (*int64, bool)`

GetShopSectionIdOk returns a tuple with the ShopSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopSectionId

`func (o *CreateDraftListingRequest) SetShopSectionId(v int64)`

SetShopSectionId sets ShopSectionId field to given value.

### HasShopSectionId

`func (o *CreateDraftListingRequest) HasShopSectionId() bool`

HasShopSectionId returns a boolean if a field has been set.

### SetShopSectionIdNil

`func (o *CreateDraftListingRequest) SetShopSectionIdNil(b bool)`

 SetShopSectionIdNil sets the value for ShopSectionId to be an explicit nil

### UnsetShopSectionId
`func (o *CreateDraftListingRequest) UnsetShopSectionId()`

UnsetShopSectionId ensures that no value is present for ShopSectionId, not even an explicit nil
### GetProcessingMin

`func (o *CreateDraftListingRequest) GetProcessingMin() int64`

GetProcessingMin returns the ProcessingMin field if non-nil, zero value otherwise.

### GetProcessingMinOk

`func (o *CreateDraftListingRequest) GetProcessingMinOk() (*int64, bool)`

GetProcessingMinOk returns a tuple with the ProcessingMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMin

`func (o *CreateDraftListingRequest) SetProcessingMin(v int64)`

SetProcessingMin sets ProcessingMin field to given value.

### HasProcessingMin

`func (o *CreateDraftListingRequest) HasProcessingMin() bool`

HasProcessingMin returns a boolean if a field has been set.

### SetProcessingMinNil

`func (o *CreateDraftListingRequest) SetProcessingMinNil(b bool)`

 SetProcessingMinNil sets the value for ProcessingMin to be an explicit nil

### UnsetProcessingMin
`func (o *CreateDraftListingRequest) UnsetProcessingMin()`

UnsetProcessingMin ensures that no value is present for ProcessingMin, not even an explicit nil
### GetProcessingMax

`func (o *CreateDraftListingRequest) GetProcessingMax() int64`

GetProcessingMax returns the ProcessingMax field if non-nil, zero value otherwise.

### GetProcessingMaxOk

`func (o *CreateDraftListingRequest) GetProcessingMaxOk() (*int64, bool)`

GetProcessingMaxOk returns a tuple with the ProcessingMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMax

`func (o *CreateDraftListingRequest) SetProcessingMax(v int64)`

SetProcessingMax sets ProcessingMax field to given value.

### HasProcessingMax

`func (o *CreateDraftListingRequest) HasProcessingMax() bool`

HasProcessingMax returns a boolean if a field has been set.

### SetProcessingMaxNil

`func (o *CreateDraftListingRequest) SetProcessingMaxNil(b bool)`

 SetProcessingMaxNil sets the value for ProcessingMax to be an explicit nil

### UnsetProcessingMax
`func (o *CreateDraftListingRequest) UnsetProcessingMax()`

UnsetProcessingMax ensures that no value is present for ProcessingMax, not even an explicit nil
### GetTags

`func (o *CreateDraftListingRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDraftListingRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDraftListingRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDraftListingRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateDraftListingRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateDraftListingRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetStyles

`func (o *CreateDraftListingRequest) GetStyles() []string`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *CreateDraftListingRequest) GetStylesOk() (*[]string, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *CreateDraftListingRequest) SetStyles(v []string)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *CreateDraftListingRequest) HasStyles() bool`

HasStyles returns a boolean if a field has been set.

### SetStylesNil

`func (o *CreateDraftListingRequest) SetStylesNil(b bool)`

 SetStylesNil sets the value for Styles to be an explicit nil

### UnsetStyles
`func (o *CreateDraftListingRequest) UnsetStyles()`

UnsetStyles ensures that no value is present for Styles, not even an explicit nil
### GetItemWeight

`func (o *CreateDraftListingRequest) GetItemWeight() float32`

GetItemWeight returns the ItemWeight field if non-nil, zero value otherwise.

### GetItemWeightOk

`func (o *CreateDraftListingRequest) GetItemWeightOk() (*float32, bool)`

GetItemWeightOk returns a tuple with the ItemWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWeight

`func (o *CreateDraftListingRequest) SetItemWeight(v float32)`

SetItemWeight sets ItemWeight field to given value.

### HasItemWeight

`func (o *CreateDraftListingRequest) HasItemWeight() bool`

HasItemWeight returns a boolean if a field has been set.

### SetItemWeightNil

`func (o *CreateDraftListingRequest) SetItemWeightNil(b bool)`

 SetItemWeightNil sets the value for ItemWeight to be an explicit nil

### UnsetItemWeight
`func (o *CreateDraftListingRequest) UnsetItemWeight()`

UnsetItemWeight ensures that no value is present for ItemWeight, not even an explicit nil
### GetItemLength

`func (o *CreateDraftListingRequest) GetItemLength() float32`

GetItemLength returns the ItemLength field if non-nil, zero value otherwise.

### GetItemLengthOk

`func (o *CreateDraftListingRequest) GetItemLengthOk() (*float32, bool)`

GetItemLengthOk returns a tuple with the ItemLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLength

`func (o *CreateDraftListingRequest) SetItemLength(v float32)`

SetItemLength sets ItemLength field to given value.

### HasItemLength

`func (o *CreateDraftListingRequest) HasItemLength() bool`

HasItemLength returns a boolean if a field has been set.

### SetItemLengthNil

`func (o *CreateDraftListingRequest) SetItemLengthNil(b bool)`

 SetItemLengthNil sets the value for ItemLength to be an explicit nil

### UnsetItemLength
`func (o *CreateDraftListingRequest) UnsetItemLength()`

UnsetItemLength ensures that no value is present for ItemLength, not even an explicit nil
### GetItemWidth

`func (o *CreateDraftListingRequest) GetItemWidth() float32`

GetItemWidth returns the ItemWidth field if non-nil, zero value otherwise.

### GetItemWidthOk

`func (o *CreateDraftListingRequest) GetItemWidthOk() (*float32, bool)`

GetItemWidthOk returns a tuple with the ItemWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWidth

`func (o *CreateDraftListingRequest) SetItemWidth(v float32)`

SetItemWidth sets ItemWidth field to given value.

### HasItemWidth

`func (o *CreateDraftListingRequest) HasItemWidth() bool`

HasItemWidth returns a boolean if a field has been set.

### SetItemWidthNil

`func (o *CreateDraftListingRequest) SetItemWidthNil(b bool)`

 SetItemWidthNil sets the value for ItemWidth to be an explicit nil

### UnsetItemWidth
`func (o *CreateDraftListingRequest) UnsetItemWidth()`

UnsetItemWidth ensures that no value is present for ItemWidth, not even an explicit nil
### GetItemHeight

`func (o *CreateDraftListingRequest) GetItemHeight() float32`

GetItemHeight returns the ItemHeight field if non-nil, zero value otherwise.

### GetItemHeightOk

`func (o *CreateDraftListingRequest) GetItemHeightOk() (*float32, bool)`

GetItemHeightOk returns a tuple with the ItemHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemHeight

`func (o *CreateDraftListingRequest) SetItemHeight(v float32)`

SetItemHeight sets ItemHeight field to given value.

### HasItemHeight

`func (o *CreateDraftListingRequest) HasItemHeight() bool`

HasItemHeight returns a boolean if a field has been set.

### SetItemHeightNil

`func (o *CreateDraftListingRequest) SetItemHeightNil(b bool)`

 SetItemHeightNil sets the value for ItemHeight to be an explicit nil

### UnsetItemHeight
`func (o *CreateDraftListingRequest) UnsetItemHeight()`

UnsetItemHeight ensures that no value is present for ItemHeight, not even an explicit nil
### GetItemWeightUnit

`func (o *CreateDraftListingRequest) GetItemWeightUnit() CreateDraftListingRequestItemWeightUnit`

GetItemWeightUnit returns the ItemWeightUnit field if non-nil, zero value otherwise.

### GetItemWeightUnitOk

`func (o *CreateDraftListingRequest) GetItemWeightUnitOk() (*CreateDraftListingRequestItemWeightUnit, bool)`

GetItemWeightUnitOk returns a tuple with the ItemWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWeightUnit

`func (o *CreateDraftListingRequest) SetItemWeightUnit(v CreateDraftListingRequestItemWeightUnit)`

SetItemWeightUnit sets ItemWeightUnit field to given value.

### HasItemWeightUnit

`func (o *CreateDraftListingRequest) HasItemWeightUnit() bool`

HasItemWeightUnit returns a boolean if a field has been set.

### SetItemWeightUnitNil

`func (o *CreateDraftListingRequest) SetItemWeightUnitNil(b bool)`

 SetItemWeightUnitNil sets the value for ItemWeightUnit to be an explicit nil

### UnsetItemWeightUnit
`func (o *CreateDraftListingRequest) UnsetItemWeightUnit()`

UnsetItemWeightUnit ensures that no value is present for ItemWeightUnit, not even an explicit nil
### GetItemDimensionsUnit

`func (o *CreateDraftListingRequest) GetItemDimensionsUnit() CreateDraftListingRequestItemDimensionsUnit`

GetItemDimensionsUnit returns the ItemDimensionsUnit field if non-nil, zero value otherwise.

### GetItemDimensionsUnitOk

`func (o *CreateDraftListingRequest) GetItemDimensionsUnitOk() (*CreateDraftListingRequestItemDimensionsUnit, bool)`

GetItemDimensionsUnitOk returns a tuple with the ItemDimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDimensionsUnit

`func (o *CreateDraftListingRequest) SetItemDimensionsUnit(v CreateDraftListingRequestItemDimensionsUnit)`

SetItemDimensionsUnit sets ItemDimensionsUnit field to given value.

### HasItemDimensionsUnit

`func (o *CreateDraftListingRequest) HasItemDimensionsUnit() bool`

HasItemDimensionsUnit returns a boolean if a field has been set.

### SetItemDimensionsUnitNil

`func (o *CreateDraftListingRequest) SetItemDimensionsUnitNil(b bool)`

 SetItemDimensionsUnitNil sets the value for ItemDimensionsUnit to be an explicit nil

### UnsetItemDimensionsUnit
`func (o *CreateDraftListingRequest) UnsetItemDimensionsUnit()`

UnsetItemDimensionsUnit ensures that no value is present for ItemDimensionsUnit, not even an explicit nil
### GetIsPersonalizable

`func (o *CreateDraftListingRequest) GetIsPersonalizable() bool`

GetIsPersonalizable returns the IsPersonalizable field if non-nil, zero value otherwise.

### GetIsPersonalizableOk

`func (o *CreateDraftListingRequest) GetIsPersonalizableOk() (*bool, bool)`

GetIsPersonalizableOk returns a tuple with the IsPersonalizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonalizable

`func (o *CreateDraftListingRequest) SetIsPersonalizable(v bool)`

SetIsPersonalizable sets IsPersonalizable field to given value.

### HasIsPersonalizable

`func (o *CreateDraftListingRequest) HasIsPersonalizable() bool`

HasIsPersonalizable returns a boolean if a field has been set.

### GetPersonalizationIsRequired

`func (o *CreateDraftListingRequest) GetPersonalizationIsRequired() bool`

GetPersonalizationIsRequired returns the PersonalizationIsRequired field if non-nil, zero value otherwise.

### GetPersonalizationIsRequiredOk

`func (o *CreateDraftListingRequest) GetPersonalizationIsRequiredOk() (*bool, bool)`

GetPersonalizationIsRequiredOk returns a tuple with the PersonalizationIsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationIsRequired

`func (o *CreateDraftListingRequest) SetPersonalizationIsRequired(v bool)`

SetPersonalizationIsRequired sets PersonalizationIsRequired field to given value.

### HasPersonalizationIsRequired

`func (o *CreateDraftListingRequest) HasPersonalizationIsRequired() bool`

HasPersonalizationIsRequired returns a boolean if a field has been set.

### GetPersonalizationCharCountMax

`func (o *CreateDraftListingRequest) GetPersonalizationCharCountMax() int64`

GetPersonalizationCharCountMax returns the PersonalizationCharCountMax field if non-nil, zero value otherwise.

### GetPersonalizationCharCountMaxOk

`func (o *CreateDraftListingRequest) GetPersonalizationCharCountMaxOk() (*int64, bool)`

GetPersonalizationCharCountMaxOk returns a tuple with the PersonalizationCharCountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationCharCountMax

`func (o *CreateDraftListingRequest) SetPersonalizationCharCountMax(v int64)`

SetPersonalizationCharCountMax sets PersonalizationCharCountMax field to given value.

### HasPersonalizationCharCountMax

`func (o *CreateDraftListingRequest) HasPersonalizationCharCountMax() bool`

HasPersonalizationCharCountMax returns a boolean if a field has been set.

### GetPersonalizationInstructions

`func (o *CreateDraftListingRequest) GetPersonalizationInstructions() string`

GetPersonalizationInstructions returns the PersonalizationInstructions field if non-nil, zero value otherwise.

### GetPersonalizationInstructionsOk

`func (o *CreateDraftListingRequest) GetPersonalizationInstructionsOk() (*string, bool)`

GetPersonalizationInstructionsOk returns a tuple with the PersonalizationInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationInstructions

`func (o *CreateDraftListingRequest) SetPersonalizationInstructions(v string)`

SetPersonalizationInstructions sets PersonalizationInstructions field to given value.

### HasPersonalizationInstructions

`func (o *CreateDraftListingRequest) HasPersonalizationInstructions() bool`

HasPersonalizationInstructions returns a boolean if a field has been set.

### GetProductionPartnerIds

`func (o *CreateDraftListingRequest) GetProductionPartnerIds() []int64`

GetProductionPartnerIds returns the ProductionPartnerIds field if non-nil, zero value otherwise.

### GetProductionPartnerIdsOk

`func (o *CreateDraftListingRequest) GetProductionPartnerIdsOk() (*[]int64, bool)`

GetProductionPartnerIdsOk returns a tuple with the ProductionPartnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionPartnerIds

`func (o *CreateDraftListingRequest) SetProductionPartnerIds(v []int64)`

SetProductionPartnerIds sets ProductionPartnerIds field to given value.

### HasProductionPartnerIds

`func (o *CreateDraftListingRequest) HasProductionPartnerIds() bool`

HasProductionPartnerIds returns a boolean if a field has been set.

### SetProductionPartnerIdsNil

`func (o *CreateDraftListingRequest) SetProductionPartnerIdsNil(b bool)`

 SetProductionPartnerIdsNil sets the value for ProductionPartnerIds to be an explicit nil

### UnsetProductionPartnerIds
`func (o *CreateDraftListingRequest) UnsetProductionPartnerIds()`

UnsetProductionPartnerIds ensures that no value is present for ProductionPartnerIds, not even an explicit nil
### GetImageIds

`func (o *CreateDraftListingRequest) GetImageIds() []int64`

GetImageIds returns the ImageIds field if non-nil, zero value otherwise.

### GetImageIdsOk

`func (o *CreateDraftListingRequest) GetImageIdsOk() (*[]int64, bool)`

GetImageIdsOk returns a tuple with the ImageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIds

`func (o *CreateDraftListingRequest) SetImageIds(v []int64)`

SetImageIds sets ImageIds field to given value.

### HasImageIds

`func (o *CreateDraftListingRequest) HasImageIds() bool`

HasImageIds returns a boolean if a field has been set.

### SetImageIdsNil

`func (o *CreateDraftListingRequest) SetImageIdsNil(b bool)`

 SetImageIdsNil sets the value for ImageIds to be an explicit nil

### UnsetImageIds
`func (o *CreateDraftListingRequest) UnsetImageIds()`

UnsetImageIds ensures that no value is present for ImageIds, not even an explicit nil
### GetIsSupply

`func (o *CreateDraftListingRequest) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *CreateDraftListingRequest) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *CreateDraftListingRequest) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.

### HasIsSupply

`func (o *CreateDraftListingRequest) HasIsSupply() bool`

HasIsSupply returns a boolean if a field has been set.

### GetIsCustomizable

`func (o *CreateDraftListingRequest) GetIsCustomizable() bool`

GetIsCustomizable returns the IsCustomizable field if non-nil, zero value otherwise.

### GetIsCustomizableOk

`func (o *CreateDraftListingRequest) GetIsCustomizableOk() (*bool, bool)`

GetIsCustomizableOk returns a tuple with the IsCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomizable

`func (o *CreateDraftListingRequest) SetIsCustomizable(v bool)`

SetIsCustomizable sets IsCustomizable field to given value.

### HasIsCustomizable

`func (o *CreateDraftListingRequest) HasIsCustomizable() bool`

HasIsCustomizable returns a boolean if a field has been set.

### GetShouldAutoRenew

`func (o *CreateDraftListingRequest) GetShouldAutoRenew() bool`

GetShouldAutoRenew returns the ShouldAutoRenew field if non-nil, zero value otherwise.

### GetShouldAutoRenewOk

`func (o *CreateDraftListingRequest) GetShouldAutoRenewOk() (*bool, bool)`

GetShouldAutoRenewOk returns a tuple with the ShouldAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAutoRenew

`func (o *CreateDraftListingRequest) SetShouldAutoRenew(v bool)`

SetShouldAutoRenew sets ShouldAutoRenew field to given value.

### HasShouldAutoRenew

`func (o *CreateDraftListingRequest) HasShouldAutoRenew() bool`

HasShouldAutoRenew returns a boolean if a field has been set.

### GetIsTaxable

`func (o *CreateDraftListingRequest) GetIsTaxable() bool`

GetIsTaxable returns the IsTaxable field if non-nil, zero value otherwise.

### GetIsTaxableOk

`func (o *CreateDraftListingRequest) GetIsTaxableOk() (*bool, bool)`

GetIsTaxableOk returns a tuple with the IsTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxable

`func (o *CreateDraftListingRequest) SetIsTaxable(v bool)`

SetIsTaxable sets IsTaxable field to given value.

### HasIsTaxable

`func (o *CreateDraftListingRequest) HasIsTaxable() bool`

HasIsTaxable returns a boolean if a field has been set.

### GetType

`func (o *CreateDraftListingRequest) GetType() CreateDraftListingRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDraftListingRequest) GetTypeOk() (*CreateDraftListingRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDraftListingRequest) SetType(v CreateDraftListingRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateDraftListingRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


