# Go API client for goEtsy

<div class=\"wt-text-body-01\"><p class=\"wt-pt-xs-2 wt-pb-xs-2\">Etsy's Open API provides a simple RESTful interface for various Etsy.com features. The API endpoints are meant to replace Etsy's Open API v2, which is scheduled to end service in 2022.</p><p class=\"wt-pb-xs-2\">All of the endpoints are callable and the majority of the API endpoints are now in a beta phase. This means we do not expect to make any breaking changes before our general release. A handful of endpoints are currently interface stubs (labeled “Feedback Only”) and returns a \"501 Not Implemented\" response code when called.</p><p class=\"wt-pb-xs-2\">If you'd like to report an issue or provide feedback on the API design, <a target=\"_blank\" class=\"wt-text-link wt-p-xs-0\" href=\"https://github.com/etsy/open-api/discussions\">please add an issue in Github</a>.</p></div>&copy; 2021-2024 Etsy, Inc. All Rights Reserved. Use of this code is subject to Etsy's <a class='wt-text-link wt-p-xs-0' target='_blank' href='https://www.etsy.com/legal/api'>API Developer Terms of Use</a>.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 1.0.0
- Build date: 2024-08-28T16:45:20.627923016+07:00[Asia/Saigon]
- Generator version: 7.8.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import goEtsy "github.com/gearment/go-etsy/client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `goEtsy.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), goEtsy.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `goEtsy.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), goEtsy.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `goEtsy.ContextOperationServerIndices` and `goEtsy.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), goEtsy.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), goEtsy.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://openapi.etsy.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BuyerTaxonomyAPI* | [**GetBuyerTaxonomyNodes**](docs/BuyerTaxonomyAPI.md#getbuyertaxonomynodes) | **Get** /v3/application/buyer-taxonomy/nodes | 
*BuyerTaxonomyAPI* | [**GetPropertiesByBuyerTaxonomyId**](docs/BuyerTaxonomyAPI.md#getpropertiesbybuyertaxonomyid) | **Get** /v3/application/buyer-taxonomy/nodes/{taxonomy_id}/properties | 
*LedgerEntryAPI* | [**GetShopPaymentAccountLedgerEntries**](docs/LedgerEntryAPI.md#getshoppaymentaccountledgerentries) | **Get** /v3/application/shops/{shop_id}/payment-account/ledger-entries | 
*LedgerEntryAPI* | [**GetShopPaymentAccountLedgerEntry**](docs/LedgerEntryAPI.md#getshoppaymentaccountledgerentry) | **Get** /v3/application/shops/{shop_id}/payment-account/ledger-entries/{ledger_entry_id} | 
*OtherAPI* | [**Ping**](docs/OtherAPI.md#ping) | **Get** /v3/application/openapi-ping | 
*OtherAPI* | [**TokenScopes**](docs/OtherAPI.md#tokenscopes) | **Post** /v3/application/scopes | 
*PaymentAPI* | [**GetPaymentAccountLedgerEntryPayments**](docs/PaymentAPI.md#getpaymentaccountledgerentrypayments) | **Get** /v3/application/shops/{shop_id}/payment-account/ledger-entries/payments | 
*PaymentAPI* | [**GetPayments**](docs/PaymentAPI.md#getpayments) | **Get** /v3/application/shops/{shop_id}/payments | 
*PaymentAPI* | [**GetShopPaymentByReceiptId**](docs/PaymentAPI.md#getshoppaymentbyreceiptid) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id}/payments | 
*ReviewAPI* | [**GetReviewsByListing**](docs/ReviewAPI.md#getreviewsbylisting) | **Get** /v3/application/listings/{listing_id}/reviews | 
*ReviewAPI* | [**GetReviewsByShop**](docs/ReviewAPI.md#getreviewsbyshop) | **Get** /v3/application/shops/{shop_id}/reviews | 
*SellerTaxonomyAPI* | [**GetPropertiesByTaxonomyId**](docs/SellerTaxonomyAPI.md#getpropertiesbytaxonomyid) | **Get** /v3/application/seller-taxonomy/nodes/{taxonomy_id}/properties | 
*SellerTaxonomyAPI* | [**GetSellerTaxonomyNodes**](docs/SellerTaxonomyAPI.md#getsellertaxonomynodes) | **Get** /v3/application/seller-taxonomy/nodes | 
*ShopAPI* | [**FindShops**](docs/ShopAPI.md#findshops) | **Get** /v3/application/shops | 
*ShopAPI* | [**GetShop**](docs/ShopAPI.md#getshop) | **Get** /v3/application/shops/{shop_id} | 
*ShopAPI* | [**GetShopByOwnerUserId**](docs/ShopAPI.md#getshopbyowneruserid) | **Get** /v3/application/users/{user_id}/shops | 
*ShopAPI* | [**UpdateShop**](docs/ShopAPI.md#updateshop) | **Put** /v3/application/shops/{shop_id} | 
*ShopHolidayPreferencesAPI* | [**GetHolidayPreferences**](docs/ShopHolidayPreferencesAPI.md#getholidaypreferences) | **Get** /v3/application/shops/{shop_id}/holiday-preferences | 
*ShopHolidayPreferencesAPI* | [**UpdateHolidayPreferences**](docs/ShopHolidayPreferencesAPI.md#updateholidaypreferences) | **Put** /v3/application/shops/{shop_id}/holiday-preferences/{holiday_id} | 
*ShopListingAPI* | [**CreateDraftListing**](docs/ShopListingAPI.md#createdraftlisting) | **Post** /v3/application/shops/{shop_id}/listings | 
*ShopListingAPI* | [**DeleteListing**](docs/ShopListingAPI.md#deletelisting) | **Delete** /v3/application/listings/{listing_id} | 
*ShopListingAPI* | [**DeleteListingProperty**](docs/ShopListingAPI.md#deletelistingproperty) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/properties/{property_id} | 
*ShopListingAPI* | [**FindAllActiveListingsByShop**](docs/ShopListingAPI.md#findallactivelistingsbyshop) | **Get** /v3/application/shops/{shop_id}/listings/active | 
*ShopListingAPI* | [**FindAllListingsActive**](docs/ShopListingAPI.md#findalllistingsactive) | **Get** /v3/application/listings/active | 
*ShopListingAPI* | [**GetFeaturedListingsByShop**](docs/ShopListingAPI.md#getfeaturedlistingsbyshop) | **Get** /v3/application/shops/{shop_id}/listings/featured | 
*ShopListingAPI* | [**GetListing**](docs/ShopListingAPI.md#getlisting) | **Get** /v3/application/listings/{listing_id} | 
*ShopListingAPI* | [**GetListingProperties**](docs/ShopListingAPI.md#getlistingproperties) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/properties | 
*ShopListingAPI* | [**GetListingProperty**](docs/ShopListingAPI.md#getlistingproperty) | **Get** /v3/application/listings/{listing_id}/properties/{property_id} | 
*ShopListingAPI* | [**GetListingsByListingIds**](docs/ShopListingAPI.md#getlistingsbylistingids) | **Get** /v3/application/listings/batch | 
*ShopListingAPI* | [**GetListingsByShop**](docs/ShopListingAPI.md#getlistingsbyshop) | **Get** /v3/application/shops/{shop_id}/listings | 
*ShopListingAPI* | [**GetListingsByShopReceipt**](docs/ShopListingAPI.md#getlistingsbyshopreceipt) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id}/listings | 
*ShopListingAPI* | [**GetListingsByShopReturnPolicy**](docs/ShopListingAPI.md#getlistingsbyshopreturnpolicy) | **Get** /v3/application/shops/{shop_id}/policies/return/{return_policy_id}/listings | 
*ShopListingAPI* | [**GetListingsByShopSectionId**](docs/ShopListingAPI.md#getlistingsbyshopsectionid) | **Get** /v3/application/shops/{shop_id}/shop-sections/listings | 
*ShopListingAPI* | [**UpdateListing**](docs/ShopListingAPI.md#updatelisting) | **Patch** /v3/application/shops/{shop_id}/listings/{listing_id} | 
*ShopListingAPI* | [**UpdateListingDeprecated**](docs/ShopListingAPI.md#updatelistingdeprecated) | **Put** /v3/application/shops/{shop_id}/listings/{listing_id} | 
*ShopListingAPI* | [**UpdateListingProperty**](docs/ShopListingAPI.md#updatelistingproperty) | **Put** /v3/application/shops/{shop_id}/listings/{listing_id}/properties/{property_id} | 
*ShopListingFileAPI* | [**DeleteListingFile**](docs/ShopListingFileAPI.md#deletelistingfile) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/files/{listing_file_id} | 
*ShopListingFileAPI* | [**GetAllListingFiles**](docs/ShopListingFileAPI.md#getalllistingfiles) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/files | 
*ShopListingFileAPI* | [**GetListingFile**](docs/ShopListingFileAPI.md#getlistingfile) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/files/{listing_file_id} | 
*ShopListingFileAPI* | [**UploadListingFile**](docs/ShopListingFileAPI.md#uploadlistingfile) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/files | 
*ShopListingImageAPI* | [**DeleteListingImage**](docs/ShopListingImageAPI.md#deletelistingimage) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/images/{listing_image_id} | 
*ShopListingImageAPI* | [**GetListingImage**](docs/ShopListingImageAPI.md#getlistingimage) | **Get** /v3/application/listings/{listing_id}/images/{listing_image_id} | 
*ShopListingImageAPI* | [**GetListingImageDeprecated**](docs/ShopListingImageAPI.md#getlistingimagedeprecated) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/images/{listing_image_id} | 
*ShopListingImageAPI* | [**GetListingImages**](docs/ShopListingImageAPI.md#getlistingimages) | **Get** /v3/application/listings/{listing_id}/images | 
*ShopListingImageAPI* | [**GetListingImagesDeprecated**](docs/ShopListingImageAPI.md#getlistingimagesdeprecated) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/images | 
*ShopListingImageAPI* | [**UploadListingImage**](docs/ShopListingImageAPI.md#uploadlistingimage) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/images | 
*ShopListingInventoryAPI* | [**GetListingInventory**](docs/ShopListingInventoryAPI.md#getlistinginventory) | **Get** /v3/application/listings/{listing_id}/inventory | 
*ShopListingInventoryAPI* | [**UpdateListingInventory**](docs/ShopListingInventoryAPI.md#updatelistinginventory) | **Put** /v3/application/listings/{listing_id}/inventory | 
*ShopListingOfferingAPI* | [**GetListingOffering**](docs/ShopListingOfferingAPI.md#getlistingoffering) | **Get** /v3/application/listings/{listing_id}/products/{product_id}/offerings/{product_offering_id} | 
*ShopListingProductAPI* | [**GetListingProduct**](docs/ShopListingProductAPI.md#getlistingproduct) | **Get** /v3/application/listings/{listing_id}/inventory/products/{product_id} | 
*ShopListingTranslationAPI* | [**CreateListingTranslation**](docs/ShopListingTranslationAPI.md#createlistingtranslation) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/translations/{language} | 
*ShopListingTranslationAPI* | [**GetListingTranslation**](docs/ShopListingTranslationAPI.md#getlistingtranslation) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/translations/{language} | 
*ShopListingTranslationAPI* | [**UpdateListingTranslation**](docs/ShopListingTranslationAPI.md#updatelistingtranslation) | **Put** /v3/application/shops/{shop_id}/listings/{listing_id}/translations/{language} | 
*ShopListingVariationImageAPI* | [**GetListingVariationImages**](docs/ShopListingVariationImageAPI.md#getlistingvariationimages) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/variation-images | 
*ShopListingVariationImageAPI* | [**UpdateVariationImages**](docs/ShopListingVariationImageAPI.md#updatevariationimages) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/variation-images | 
*ShopListingVideoAPI* | [**DeleteListingVideo**](docs/ShopListingVideoAPI.md#deletelistingvideo) | **Delete** /v3/application/shops/{shop_id}/listings/{listing_id}/videos/{video_id} | 
*ShopListingVideoAPI* | [**GetListingVideo**](docs/ShopListingVideoAPI.md#getlistingvideo) | **Get** /v3/application/listings/{listing_id}/videos/{video_id} | 
*ShopListingVideoAPI* | [**GetListingVideos**](docs/ShopListingVideoAPI.md#getlistingvideos) | **Get** /v3/application/listings/{listing_id}/videos | 
*ShopListingVideoAPI* | [**UploadListingVideo**](docs/ShopListingVideoAPI.md#uploadlistingvideo) | **Post** /v3/application/shops/{shop_id}/listings/{listing_id}/videos | 
*ShopProductionPartnerAPI* | [**GetShopProductionPartners**](docs/ShopProductionPartnerAPI.md#getshopproductionpartners) | **Get** /v3/application/shops/{shop_id}/production-partners | 
*ShopReceiptAPI* | [**CreateReceiptShipment**](docs/ShopReceiptAPI.md#createreceiptshipment) | **Post** /v3/application/shops/{shop_id}/receipts/{receipt_id}/tracking | 
*ShopReceiptAPI* | [**GetShopReceipt**](docs/ShopReceiptAPI.md#getshopreceipt) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id} | 
*ShopReceiptAPI* | [**GetShopReceipts**](docs/ShopReceiptAPI.md#getshopreceipts) | **Get** /v3/application/shops/{shop_id}/receipts | 
*ShopReceiptAPI* | [**UpdateShopReceipt**](docs/ShopReceiptAPI.md#updateshopreceipt) | **Put** /v3/application/shops/{shop_id}/receipts/{receipt_id} | 
*ShopReceiptTransactionsAPI* | [**GetShopReceiptTransaction**](docs/ShopReceiptTransactionsAPI.md#getshopreceipttransaction) | **Get** /v3/application/shops/{shop_id}/transactions/{transaction_id} | 
*ShopReceiptTransactionsAPI* | [**GetShopReceiptTransactionsByListing**](docs/ShopReceiptTransactionsAPI.md#getshopreceipttransactionsbylisting) | **Get** /v3/application/shops/{shop_id}/listings/{listing_id}/transactions | 
*ShopReceiptTransactionsAPI* | [**GetShopReceiptTransactionsByReceipt**](docs/ShopReceiptTransactionsAPI.md#getshopreceipttransactionsbyreceipt) | **Get** /v3/application/shops/{shop_id}/receipts/{receipt_id}/transactions | 
*ShopReceiptTransactionsAPI* | [**GetShopReceiptTransactionsByShop**](docs/ShopReceiptTransactionsAPI.md#getshopreceipttransactionsbyshop) | **Get** /v3/application/shops/{shop_id}/transactions | 
*ShopReturnPolicyAPI* | [**ConsolidateShopReturnPolicies**](docs/ShopReturnPolicyAPI.md#consolidateshopreturnpolicies) | **Post** /v3/application/shops/{shop_id}/policies/return/consolidate | 
*ShopReturnPolicyAPI* | [**CreateShopReturnPolicy**](docs/ShopReturnPolicyAPI.md#createshopreturnpolicy) | **Post** /v3/application/shops/{shop_id}/policies/return | 
*ShopReturnPolicyAPI* | [**DeleteShopReturnPolicy**](docs/ShopReturnPolicyAPI.md#deleteshopreturnpolicy) | **Delete** /v3/application/shops/{shop_id}/policies/return/{return_policy_id} | 
*ShopReturnPolicyAPI* | [**GetShopReturnPolicies**](docs/ShopReturnPolicyAPI.md#getshopreturnpolicies) | **Get** /v3/application/shops/{shop_id}/policies/return | 
*ShopReturnPolicyAPI* | [**GetShopReturnPolicy**](docs/ShopReturnPolicyAPI.md#getshopreturnpolicy) | **Get** /v3/application/shops/{shop_id}/policies/return/{return_policy_id} | 
*ShopReturnPolicyAPI* | [**UpdateShopReturnPolicy**](docs/ShopReturnPolicyAPI.md#updateshopreturnpolicy) | **Put** /v3/application/shops/{shop_id}/policies/return/{return_policy_id} | 
*ShopSectionAPI* | [**CreateShopSection**](docs/ShopSectionAPI.md#createshopsection) | **Post** /v3/application/shops/{shop_id}/sections | 
*ShopSectionAPI* | [**DeleteShopSection**](docs/ShopSectionAPI.md#deleteshopsection) | **Delete** /v3/application/shops/{shop_id}/sections/{shop_section_id} | 
*ShopSectionAPI* | [**GetShopSection**](docs/ShopSectionAPI.md#getshopsection) | **Get** /v3/application/shops/{shop_id}/sections/{shop_section_id} | 
*ShopSectionAPI* | [**GetShopSections**](docs/ShopSectionAPI.md#getshopsections) | **Get** /v3/application/shops/{shop_id}/sections | 
*ShopSectionAPI* | [**UpdateShopSection**](docs/ShopSectionAPI.md#updateshopsection) | **Put** /v3/application/shops/{shop_id}/sections/{shop_section_id} | 
*ShopShippingProfileAPI* | [**CreateShopShippingProfile**](docs/ShopShippingProfileAPI.md#createshopshippingprofile) | **Post** /v3/application/shops/{shop_id}/shipping-profiles | 
*ShopShippingProfileAPI* | [**CreateShopShippingProfileDestination**](docs/ShopShippingProfileAPI.md#createshopshippingprofiledestination) | **Post** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations | 
*ShopShippingProfileAPI* | [**CreateShopShippingProfileUpgrade**](docs/ShopShippingProfileAPI.md#createshopshippingprofileupgrade) | **Post** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades | 
*ShopShippingProfileAPI* | [**DeleteShopShippingProfile**](docs/ShopShippingProfileAPI.md#deleteshopshippingprofile) | **Delete** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id} | 
*ShopShippingProfileAPI* | [**DeleteShopShippingProfileDestination**](docs/ShopShippingProfileAPI.md#deleteshopshippingprofiledestination) | **Delete** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations/{shipping_profile_destination_id} | 
*ShopShippingProfileAPI* | [**DeleteShopShippingProfileUpgrade**](docs/ShopShippingProfileAPI.md#deleteshopshippingprofileupgrade) | **Delete** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades/{upgrade_id} | 
*ShopShippingProfileAPI* | [**GetShippingCarriers**](docs/ShopShippingProfileAPI.md#getshippingcarriers) | **Get** /v3/application/shipping-carriers | 
*ShopShippingProfileAPI* | [**GetShopShippingProfile**](docs/ShopShippingProfileAPI.md#getshopshippingprofile) | **Get** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id} | 
*ShopShippingProfileAPI* | [**GetShopShippingProfileDestinationsByShippingProfile**](docs/ShopShippingProfileAPI.md#getshopshippingprofiledestinationsbyshippingprofile) | **Get** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations | 
*ShopShippingProfileAPI* | [**GetShopShippingProfileUpgrades**](docs/ShopShippingProfileAPI.md#getshopshippingprofileupgrades) | **Get** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades | 
*ShopShippingProfileAPI* | [**GetShopShippingProfiles**](docs/ShopShippingProfileAPI.md#getshopshippingprofiles) | **Get** /v3/application/shops/{shop_id}/shipping-profiles | 
*ShopShippingProfileAPI* | [**UpdateShopShippingProfile**](docs/ShopShippingProfileAPI.md#updateshopshippingprofile) | **Put** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id} | 
*ShopShippingProfileAPI* | [**UpdateShopShippingProfileDestination**](docs/ShopShippingProfileAPI.md#updateshopshippingprofiledestination) | **Put** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations/{shipping_profile_destination_id} | 
*ShopShippingProfileAPI* | [**UpdateShopShippingProfileUpgrade**](docs/ShopShippingProfileAPI.md#updateshopshippingprofileupgrade) | **Put** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades/{upgrade_id} | 
*UserAPI* | [**GetMe**](docs/UserAPI.md#getme) | **Get** /v3/application/users/me | 
*UserAPI* | [**GetUser**](docs/UserAPI.md#getuser) | **Get** /v3/application/users/{user_id} | 
*UserAddressAPI* | [**DeleteUserAddress**](docs/UserAddressAPI.md#deleteuseraddress) | **Delete** /v3/application/user/addresses/{user_address_id} | 
*UserAddressAPI* | [**GetUserAddress**](docs/UserAddressAPI.md#getuseraddress) | **Get** /v3/application/user/addresses/{user_address_id} | 
*UserAddressAPI* | [**GetUserAddresses**](docs/UserAddressAPI.md#getuseraddresses) | **Get** /v3/application/user/addresses | 


## Documentation For Models

 - [BuyerTaxonomyNode](docs/BuyerTaxonomyNode.md)
 - [BuyerTaxonomyNodeProperties](docs/BuyerTaxonomyNodeProperties.md)
 - [BuyerTaxonomyNodeProperty](docs/BuyerTaxonomyNodeProperty.md)
 - [BuyerTaxonomyNodes](docs/BuyerTaxonomyNodes.md)
 - [BuyerTaxonomyPropertyScale](docs/BuyerTaxonomyPropertyScale.md)
 - [BuyerTaxonomyPropertyValue](docs/BuyerTaxonomyPropertyValue.md)
 - [CreateDraftListingRequestItemDimensionsUnit](docs/CreateDraftListingRequestItemDimensionsUnit.md)
 - [CreateDraftListingRequestItemWeightUnit](docs/CreateDraftListingRequestItemWeightUnit.md)
 - [CreateDraftListingRequestType](docs/CreateDraftListingRequestType.md)
 - [CreateDraftListingRequestWhenMade](docs/CreateDraftListingRequestWhenMade.md)
 - [CreateDraftListingRequestWhoMade](docs/CreateDraftListingRequestWhoMade.md)
 - [CreateShopShippingProfileRequestDestinationRegion](docs/CreateShopShippingProfileRequestDestinationRegion.md)
 - [CreateShopShippingProfileRequestProcessingTimeUnit](docs/CreateShopShippingProfileRequestProcessingTimeUnit.md)
 - [CreateShopShippingProfileUpgradeRequestType](docs/CreateShopShippingProfileUpgradeRequestType.md)
 - [ErrorSchema](docs/ErrorSchema.md)
 - [GetListingInventoryIncludesParameter](docs/GetListingInventoryIncludesParameter.md)
 - [GetListingsByShopIncludesParameterInner](docs/GetListingsByShopIncludesParameterInner.md)
 - [GetListingsByShopSortOnParameter](docs/GetListingsByShopSortOnParameter.md)
 - [GetListingsByShopSortOrderParameter](docs/GetListingsByShopSortOrderParameter.md)
 - [GetListingsByShopStateParameter](docs/GetListingsByShopStateParameter.md)
 - [GetShopReceiptsSortOnParameter](docs/GetShopReceiptsSortOnParameter.md)
 - [GetShopReceiptsSortOrderParameter](docs/GetShopReceiptsSortOrderParameter.md)
 - [ListingImage](docs/ListingImage.md)
 - [ListingImages](docs/ListingImages.md)
 - [ListingInventory](docs/ListingInventory.md)
 - [ListingInventoryProduct](docs/ListingInventoryProduct.md)
 - [ListingInventoryProductOffering](docs/ListingInventoryProductOffering.md)
 - [ListingInventoryWithAssociations](docs/ListingInventoryWithAssociations.md)
 - [ListingPropertyValue](docs/ListingPropertyValue.md)
 - [ListingPropertyValues](docs/ListingPropertyValues.md)
 - [ListingReview](docs/ListingReview.md)
 - [ListingReviews](docs/ListingReviews.md)
 - [ListingTranslation](docs/ListingTranslation.md)
 - [ListingTranslations](docs/ListingTranslations.md)
 - [ListingVariationImage](docs/ListingVariationImage.md)
 - [ListingVariationImages](docs/ListingVariationImages.md)
 - [ListingVideo](docs/ListingVideo.md)
 - [ListingVideoVideoState](docs/ListingVideoVideoState.md)
 - [ListingVideos](docs/ListingVideos.md)
 - [Money](docs/Money.md)
 - [Payment](docs/Payment.md)
 - [PaymentAccountLedgerEntries](docs/PaymentAccountLedgerEntries.md)
 - [PaymentAccountLedgerEntry](docs/PaymentAccountLedgerEntry.md)
 - [PaymentAdjustment](docs/PaymentAdjustment.md)
 - [PaymentAdjustmentItem](docs/PaymentAdjustmentItem.md)
 - [Payments](docs/Payments.md)
 - [Pong](docs/Pong.md)
 - [Self](docs/Self.md)
 - [SellerTaxonomyNode](docs/SellerTaxonomyNode.md)
 - [SellerTaxonomyNodes](docs/SellerTaxonomyNodes.md)
 - [ShippingCarrier](docs/ShippingCarrier.md)
 - [ShippingCarrierMailClass](docs/ShippingCarrierMailClass.md)
 - [ShippingCarriers](docs/ShippingCarriers.md)
 - [Shop](docs/Shop.md)
 - [ShopHolidayPreference](docs/ShopHolidayPreference.md)
 - [ShopListing](docs/ShopListing.md)
 - [ShopListingFile](docs/ShopListingFile.md)
 - [ShopListingFiles](docs/ShopListingFiles.md)
 - [ShopListingState](docs/ShopListingState.md)
 - [ShopListingWhenMade](docs/ShopListingWhenMade.md)
 - [ShopListingWhoMade](docs/ShopListingWhoMade.md)
 - [ShopListingWithAssociations](docs/ShopListingWithAssociations.md)
 - [ShopListings](docs/ShopListings.md)
 - [ShopListingsWithAssociations](docs/ShopListingsWithAssociations.md)
 - [ShopProductionPartner](docs/ShopProductionPartner.md)
 - [ShopProductionPartners](docs/ShopProductionPartners.md)
 - [ShopReceipt](docs/ShopReceipt.md)
 - [ShopReceiptShipment](docs/ShopReceiptShipment.md)
 - [ShopReceiptStatus](docs/ShopReceiptStatus.md)
 - [ShopReceiptTransaction](docs/ShopReceiptTransaction.md)
 - [ShopReceiptTransactions](docs/ShopReceiptTransactions.md)
 - [ShopReceipts](docs/ShopReceipts.md)
 - [ShopRefund](docs/ShopRefund.md)
 - [ShopReturnPolicies](docs/ShopReturnPolicies.md)
 - [ShopReturnPolicy](docs/ShopReturnPolicy.md)
 - [ShopSection](docs/ShopSection.md)
 - [ShopSections](docs/ShopSections.md)
 - [ShopShippingProfile](docs/ShopShippingProfile.md)
 - [ShopShippingProfileDestination](docs/ShopShippingProfileDestination.md)
 - [ShopShippingProfileDestinationDestinationRegion](docs/ShopShippingProfileDestinationDestinationRegion.md)
 - [ShopShippingProfileDestinations](docs/ShopShippingProfileDestinations.md)
 - [ShopShippingProfileProfileType](docs/ShopShippingProfileProfileType.md)
 - [ShopShippingProfileUpgrade](docs/ShopShippingProfileUpgrade.md)
 - [ShopShippingProfileUpgrades](docs/ShopShippingProfileUpgrades.md)
 - [ShopShippingProfiles](docs/ShopShippingProfiles.md)
 - [Shops](docs/Shops.md)
 - [TaxonomyNodeProperties](docs/TaxonomyNodeProperties.md)
 - [TaxonomyNodeProperty](docs/TaxonomyNodeProperty.md)
 - [TaxonomyPropertyScale](docs/TaxonomyPropertyScale.md)
 - [TaxonomyPropertyValue](docs/TaxonomyPropertyValue.md)
 - [TransactionReview](docs/TransactionReview.md)
 - [TransactionReviews](docs/TransactionReviews.md)
 - [TransactionVariations](docs/TransactionVariations.md)
 - [TypeDiscriminator](docs/TypeDiscriminator.md)
 - [UpdateHolidayPreferencesHolidayIdParameter](docs/UpdateHolidayPreferencesHolidayIdParameter.md)
 - [UpdateListingDeprecatedRequestState](docs/UpdateListingDeprecatedRequestState.md)
 - [UpdateListingDeprecatedRequestType](docs/UpdateListingDeprecatedRequestType.md)
 - [UpdateListingInventoryRequest](docs/UpdateListingInventoryRequest.md)
 - [UpdateListingInventoryRequestProductsInner](docs/UpdateListingInventoryRequestProductsInner.md)
 - [UpdateListingInventoryRequestProductsInnerOfferingsInner](docs/UpdateListingInventoryRequestProductsInnerOfferingsInner.md)
 - [UpdateListingInventoryRequestProductsInnerPropertyValuesInner](docs/UpdateListingInventoryRequestProductsInnerPropertyValuesInner.md)
 - [UpdateListingRequestItemDimensionsUnit](docs/UpdateListingRequestItemDimensionsUnit.md)
 - [UpdateListingRequestItemWeightUnit](docs/UpdateListingRequestItemWeightUnit.md)
 - [UpdateVariationImagesRequest](docs/UpdateVariationImagesRequest.md)
 - [UpdateVariationImagesRequestVariationImagesInner](docs/UpdateVariationImagesRequestVariationImagesInner.md)
 - [User](docs/User.md)
 - [UserAddress](docs/UserAddress.md)
 - [UserAddresses](docs/UserAddresses.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### api_key

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: x-api-key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		goEtsy.ContextAPIKeys,
		map[string]goEtsy.APIKey{
			"x-api-key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://www.etsy.com/oauth/connect
- **Scopes**: 
 - **address_r**: see billing and shipping addresses
 - **address_w**: update billing and shipping addresses
 - **billing_r**: see all billing statement data
 - **cart_r**: read shopping carts
 - **cart_w**: add/remove from shopping carts
 - **email_r**: read a user profile
 - **favorites_r**: see private favorites
 - **favorites_w**: add/remove favorites
 - **feedback_r**: see purchase info in feedback
 - **listings_d**: delete listings
 - **listings_r**: see all listings (including expired etc)
 - **listings_w**: create/edit listings
 - **profile_r**: see all profile data
 - **profile_w**: update user profile, avatar, etc
 - **recommend_r**: see recommended listings
 - **recommend_w**: accept/reject recommended listings
 - **shops_r**: see private shop info
 - **shops_w**: update shop
 - **transactions_r**: see all checkout/payment data
 - **transactions_w**: update receipts

Example

```go
auth := context.WithValue(context.Background(), goEtsy.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, goEtsy.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

developers@etsy.com

