# \ShopShippingProfileAPI

All URIs are relative to *https://openapi.etsy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShopShippingProfile**](ShopShippingProfileAPI.md#CreateShopShippingProfile) | **Post** /v3/application/shops/{shop_id}/shipping-profiles | 
[**CreateShopShippingProfileDestination**](ShopShippingProfileAPI.md#CreateShopShippingProfileDestination) | **Post** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations | 
[**CreateShopShippingProfileUpgrade**](ShopShippingProfileAPI.md#CreateShopShippingProfileUpgrade) | **Post** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades | 
[**DeleteShopShippingProfile**](ShopShippingProfileAPI.md#DeleteShopShippingProfile) | **Delete** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id} | 
[**DeleteShopShippingProfileDestination**](ShopShippingProfileAPI.md#DeleteShopShippingProfileDestination) | **Delete** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations/{shipping_profile_destination_id} | 
[**DeleteShopShippingProfileUpgrade**](ShopShippingProfileAPI.md#DeleteShopShippingProfileUpgrade) | **Delete** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades/{upgrade_id} | 
[**GetShippingCarriers**](ShopShippingProfileAPI.md#GetShippingCarriers) | **Get** /v3/application/shipping-carriers | 
[**GetShopShippingProfile**](ShopShippingProfileAPI.md#GetShopShippingProfile) | **Get** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id} | 
[**GetShopShippingProfileDestinationsByShippingProfile**](ShopShippingProfileAPI.md#GetShopShippingProfileDestinationsByShippingProfile) | **Get** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations | 
[**GetShopShippingProfileUpgrades**](ShopShippingProfileAPI.md#GetShopShippingProfileUpgrades) | **Get** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades | 
[**GetShopShippingProfiles**](ShopShippingProfileAPI.md#GetShopShippingProfiles) | **Get** /v3/application/shops/{shop_id}/shipping-profiles | 
[**UpdateShopShippingProfile**](ShopShippingProfileAPI.md#UpdateShopShippingProfile) | **Put** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id} | 
[**UpdateShopShippingProfileDestination**](ShopShippingProfileAPI.md#UpdateShopShippingProfileDestination) | **Put** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/destinations/{shipping_profile_destination_id} | 
[**UpdateShopShippingProfileUpgrade**](ShopShippingProfileAPI.md#UpdateShopShippingProfileUpgrade) | **Put** /v3/application/shops/{shop_id}/shipping-profiles/{shipping_profile_id}/upgrades/{upgrade_id} | 



## CreateShopShippingProfile

> ShopShippingProfile CreateShopShippingProfile(ctx, shopId).Title(title).OriginCountryIso(originCountryIso).PrimaryCost(primaryCost).SecondaryCost(secondaryCost).MinProcessingTime(minProcessingTime).MaxProcessingTime(maxProcessingTime).ProcessingTimeUnit(processingTimeUnit).DestinationCountryIso(destinationCountryIso).DestinationRegion(destinationRegion).OriginPostalCode(originPostalCode).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	title := "title_example" // string | The name string of this shipping profile.
	originCountryIso := "originCountryIso_example" // string | The ISO code of the country from which the listing ships.
	primaryCost := float32(3.4) // float32 | The cost of shipping to this country/region alone, measured in the store's default currency.
	secondaryCost := float32(3.4) // float32 | The cost of shipping to this country/region with another item, measured in the store's default currency.
	minProcessingTime := int64(56) // int64 | The minimum time required to process to ship listings with this shipping profile.
	maxProcessingTime := int64(56) // int64 | The maximum processing time the listing needs to ship.
	processingTimeUnit := openapiclient.createShopShippingProfile_request_processing_time_unit("business_days") // CreateShopShippingProfileRequestProcessingTimeUnit |  (optional) (default to "business_days")
	destinationCountryIso := "destinationCountryIso_example" // string | The ISO code of the country to which the listing ships. If null, request sets destination to destination_region. Required if destination_region is null or not provided. (optional)
	destinationRegion := openapiclient.createShopShippingProfile_request_destination_region("eu") // CreateShopShippingProfileRequestDestinationRegion |  (optional) (default to "none")
	originPostalCode := "originPostalCode_example" // string | The postal code string (not necessarily a number) for the location from which the listing ships. Required if the `origin_country_iso` supports postal codes. See the [Fulfillment Tutorial docs](https://developer.etsy.com/documentation/tutorials/fulfillment/#countries-requiring-postal-codes) for more info (optional) (default to "")
	shippingCarrierId := int64(56) // int64 | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with `mail_class`** if `min_delivery_days` and `max_delivery_days` are null. (optional) (default to 0)
	mailClass := "mailClass_example" // string | The unique ID string of a shipping carrier's mail class, which is used to calculate an estimated delivery date. **Required with `shipping_carrier_id`** if `min_delivery_days` and `max_delivery_days` are null. (optional)
	minDeliveryDays := int64(56) // int64 | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `max_delivery_days`** if `mail_class` is null. (optional)
	maxDeliveryDays := int64(56) // int64 | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `min_delivery_days`** if `mail_class` is null. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.CreateShopShippingProfile(context.Background(), shopId).Title(title).OriginCountryIso(originCountryIso).PrimaryCost(primaryCost).SecondaryCost(secondaryCost).MinProcessingTime(minProcessingTime).MaxProcessingTime(maxProcessingTime).ProcessingTimeUnit(processingTimeUnit).DestinationCountryIso(destinationCountryIso).DestinationRegion(destinationRegion).OriginPostalCode(originPostalCode).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.CreateShopShippingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShopShippingProfile`: ShopShippingProfile
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.CreateShopShippingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShopShippingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** | The name string of this shipping profile. | 
 **originCountryIso** | **string** | The ISO code of the country from which the listing ships. | 
 **primaryCost** | **float32** | The cost of shipping to this country/region alone, measured in the store&#39;s default currency. | 
 **secondaryCost** | **float32** | The cost of shipping to this country/region with another item, measured in the store&#39;s default currency. | 
 **minProcessingTime** | **int64** | The minimum time required to process to ship listings with this shipping profile. | 
 **maxProcessingTime** | **int64** | The maximum processing time the listing needs to ship. | 
 **processingTimeUnit** | [**CreateShopShippingProfileRequestProcessingTimeUnit**](CreateShopShippingProfileRequestProcessingTimeUnit.md) |  | [default to &quot;business_days&quot;]
 **destinationCountryIso** | **string** | The ISO code of the country to which the listing ships. If null, request sets destination to destination_region. Required if destination_region is null or not provided. | 
 **destinationRegion** | [**CreateShopShippingProfileRequestDestinationRegion**](CreateShopShippingProfileRequestDestinationRegion.md) |  | [default to &quot;none&quot;]
 **originPostalCode** | **string** | The postal code string (not necessarily a number) for the location from which the listing ships. Required if the &#x60;origin_country_iso&#x60; supports postal codes. See the [Fulfillment Tutorial docs](https://developer.etsy.com/documentation/tutorials/fulfillment/#countries-requiring-postal-codes) for more info | [default to &quot;&quot;]
 **shippingCarrierId** | **int64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [default to 0]
 **mailClass** | **string** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | 
 **minDeliveryDays** | **int64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 
 **maxDeliveryDays** | **int64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 

### Return type

[**ShopShippingProfile**](ShopShippingProfile.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShopShippingProfileDestination

> ShopShippingProfileDestination CreateShopShippingProfileDestination(ctx, shopId, shippingProfileId).PrimaryCost(primaryCost).SecondaryCost(secondaryCost).DestinationCountryIso(destinationCountryIso).DestinationRegion(destinationRegion).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.
	primaryCost := float32(3.4) // float32 | The cost of shipping to this country/region alone, measured in the store's default currency.
	secondaryCost := float32(3.4) // float32 | The cost of shipping to this country/region with another item, measured in the store's default currency.
	destinationCountryIso := "destinationCountryIso_example" // string | The ISO code of the country to which the listing ships. If null, request sets destination to destination_region. Required if destination_region is null or not provided. (optional)
	destinationRegion := openapiclient.createShopShippingProfile_request_destination_region("eu") // CreateShopShippingProfileRequestDestinationRegion |  (optional) (default to "none")
	shippingCarrierId := int64(56) // int64 | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with `mail_class`** if `min_delivery_days` and `max_delivery_days` are null. (optional) (default to 0)
	mailClass := "mailClass_example" // string | The unique ID string of a shipping carrier's mail class, which is used to calculate an estimated delivery date. **Required with `shipping_carrier_id`** if `min_delivery_days` and `max_delivery_days` are null. (optional)
	minDeliveryDays := int64(56) // int64 | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `max_delivery_days`** if `mail_class` is null. (optional)
	maxDeliveryDays := int64(56) // int64 | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `min_delivery_days`** if `mail_class` is null. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.CreateShopShippingProfileDestination(context.Background(), shopId, shippingProfileId).PrimaryCost(primaryCost).SecondaryCost(secondaryCost).DestinationCountryIso(destinationCountryIso).DestinationRegion(destinationRegion).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.CreateShopShippingProfileDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShopShippingProfileDestination`: ShopShippingProfileDestination
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.CreateShopShippingProfileDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShopShippingProfileDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **primaryCost** | **float32** | The cost of shipping to this country/region alone, measured in the store&#39;s default currency. | 
 **secondaryCost** | **float32** | The cost of shipping to this country/region with another item, measured in the store&#39;s default currency. | 
 **destinationCountryIso** | **string** | The ISO code of the country to which the listing ships. If null, request sets destination to destination_region. Required if destination_region is null or not provided. | 
 **destinationRegion** | [**CreateShopShippingProfileRequestDestinationRegion**](CreateShopShippingProfileRequestDestinationRegion.md) |  | [default to &quot;none&quot;]
 **shippingCarrierId** | **int64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [default to 0]
 **mailClass** | **string** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | 
 **minDeliveryDays** | **int64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 
 **maxDeliveryDays** | **int64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 

### Return type

[**ShopShippingProfileDestination**](ShopShippingProfileDestination.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShopShippingProfileUpgrade

> ShopShippingProfileUpgrade CreateShopShippingProfileUpgrade(ctx, shopId, shippingProfileId).Type_(type_).UpgradeName(upgradeName).Price(price).SecondaryPrice(secondaryPrice).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.
	type_ := openapiclient.createShopShippingProfileUpgrade_request_type("0") // CreateShopShippingProfileUpgradeRequestType | 
	upgradeName := "upgradeName_example" // string | Name for the shipping upgrade shown to shoppers at checkout, e.g. USPS Priority.
	price := float32(3.4) // float32 | Additional cost of adding the shipping upgrade.
	secondaryPrice := float32(3.4) // float32 | Additional cost of adding the shipping upgrade for each additional item.
	shippingCarrierId := int64(56) // int64 | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with `mail_class`** if `min_delivery_days` and `max_delivery_days` are null. (optional) (default to 0)
	mailClass := "mailClass_example" // string | The unique ID string of a shipping carrier's mail class, which is used to calculate an estimated delivery date. **Required with `shipping_carrier_id`** if `min_delivery_days` and `max_delivery_days` are null. (optional)
	minDeliveryDays := int64(56) // int64 | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `max_delivery_days`** if `mail_class` is null. (optional)
	maxDeliveryDays := int64(56) // int64 | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `min_delivery_days`** if `mail_class` is null. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.CreateShopShippingProfileUpgrade(context.Background(), shopId, shippingProfileId).Type_(type_).UpgradeName(upgradeName).Price(price).SecondaryPrice(secondaryPrice).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.CreateShopShippingProfileUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShopShippingProfileUpgrade`: ShopShippingProfileUpgrade
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.CreateShopShippingProfileUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShopShippingProfileUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | [**CreateShopShippingProfileUpgradeRequestType**](CreateShopShippingProfileUpgradeRequestType.md) |  | 
 **upgradeName** | **string** | Name for the shipping upgrade shown to shoppers at checkout, e.g. USPS Priority. | 
 **price** | **float32** | Additional cost of adding the shipping upgrade. | 
 **secondaryPrice** | **float32** | Additional cost of adding the shipping upgrade for each additional item. | 
 **shippingCarrierId** | **int64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | [default to 0]
 **mailClass** | **string** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | 
 **minDeliveryDays** | **int64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 
 **maxDeliveryDays** | **int64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 

### Return type

[**ShopShippingProfileUpgrade**](ShopShippingProfileUpgrade.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShopShippingProfile

> DeleteShopShippingProfile(ctx, shopId, shippingProfileId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShopShippingProfileAPI.DeleteShopShippingProfile(context.Background(), shopId, shippingProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.DeleteShopShippingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShopShippingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShopShippingProfileDestination

> DeleteShopShippingProfileDestination(ctx, shopId, shippingProfileId, shippingProfileDestinationId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.
	shippingProfileDestinationId := int64(56) // int64 | The numeric ID of the shipping profile destination in the [shipping profile](/documentation/reference#tag/Shop-ShippingProfile) associated with the listing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShopShippingProfileAPI.DeleteShopShippingProfileDestination(context.Background(), shopId, shippingProfileId, shippingProfileDestinationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.DeleteShopShippingProfileDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 
**shippingProfileDestinationId** | **int64** | The numeric ID of the shipping profile destination in the [shipping profile](/documentation/reference#tag/Shop-ShippingProfile) associated with the listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShopShippingProfileDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShopShippingProfileUpgrade

> DeleteShopShippingProfileUpgrade(ctx, shopId, shippingProfileId, upgradeId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the shipping profile.
	upgradeId := int64(56) // int64 | The numeric ID that is associated with a shipping upgrade

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShopShippingProfileAPI.DeleteShopShippingProfileUpgrade(context.Background(), shopId, shippingProfileId, upgradeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.DeleteShopShippingProfileUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the shipping profile. | 
**upgradeId** | **int64** | The numeric ID that is associated with a shipping upgrade | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShopShippingProfileUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShippingCarriers

> ShippingCarriers GetShippingCarriers(ctx).OriginCountryIso(originCountryIso).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	originCountryIso := "originCountryIso_example" // string | The ISO code of the country from which the listing ships.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.GetShippingCarriers(context.Background()).OriginCountryIso(originCountryIso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.GetShippingCarriers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShippingCarriers`: ShippingCarriers
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.GetShippingCarriers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetShippingCarriersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **originCountryIso** | **string** | The ISO code of the country from which the listing ships. | 

### Return type

[**ShippingCarriers**](ShippingCarriers.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopShippingProfile

> ShopShippingProfile GetShopShippingProfile(ctx, shopId, shippingProfileId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.GetShopShippingProfile(context.Background(), shopId, shippingProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.GetShopShippingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopShippingProfile`: ShopShippingProfile
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.GetShopShippingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopShippingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShopShippingProfile**](ShopShippingProfile.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopShippingProfileDestinationsByShippingProfile

> ShopShippingProfileDestinations GetShopShippingProfileDestinationsByShippingProfile(ctx, shopId, shippingProfileId).Limit(limit).Offset(offset).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.
	limit := int64(56) // int64 | The maximum number of results to return. (optional) (default to 25)
	offset := int64(56) // int64 | The number of records to skip before selecting the first result. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.GetShopShippingProfileDestinationsByShippingProfile(context.Background(), shopId, shippingProfileId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.GetShopShippingProfileDestinationsByShippingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopShippingProfileDestinationsByShippingProfile`: ShopShippingProfileDestinations
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.GetShopShippingProfileDestinationsByShippingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopShippingProfileDestinationsByShippingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int64** | The maximum number of results to return. | [default to 25]
 **offset** | **int64** | The number of records to skip before selecting the first result. | [default to 0]

### Return type

[**ShopShippingProfileDestinations**](ShopShippingProfileDestinations.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopShippingProfileUpgrades

> ShopShippingProfileUpgrades GetShopShippingProfileUpgrades(ctx, shopId, shippingProfileId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.GetShopShippingProfileUpgrades(context.Background(), shopId, shippingProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.GetShopShippingProfileUpgrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopShippingProfileUpgrades`: ShopShippingProfileUpgrades
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.GetShopShippingProfileUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopShippingProfileUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShopShippingProfileUpgrades**](ShopShippingProfileUpgrades.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopShippingProfiles

> ShopShippingProfiles GetShopShippingProfiles(ctx, shopId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.GetShopShippingProfiles(context.Background(), shopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.GetShopShippingProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShopShippingProfiles`: ShopShippingProfiles
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.GetShopShippingProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopShippingProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShopShippingProfiles**](ShopShippingProfiles.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShopShippingProfile

> ShopShippingProfile UpdateShopShippingProfile(ctx, shopId, shippingProfileId).Title(title).OriginCountryIso(originCountryIso).MinProcessingTime(minProcessingTime).MaxProcessingTime(maxProcessingTime).ProcessingTimeUnit(processingTimeUnit).OriginPostalCode(originPostalCode).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.
	title := "title_example" // string | The name string of this shipping profile. (optional)
	originCountryIso := "originCountryIso_example" // string | The ISO code of the country from which the listing ships. (optional)
	minProcessingTime := int64(56) // int64 | The minimum time required to process to ship listings with this shipping profile. (optional)
	maxProcessingTime := int64(56) // int64 | The maximum processing time the listing needs to ship. (optional)
	processingTimeUnit := openapiclient.createShopShippingProfile_request_processing_time_unit("business_days") // CreateShopShippingProfileRequestProcessingTimeUnit |  (optional) (default to "business_days")
	originPostalCode := "originPostalCode_example" // string | The postal code string (not necessarily a number) for the location from which the listing ships. Required if the `origin_country_iso` supports postal codes. See the [Fulfillment Tutorial docs](https://developer.etsy.com/documentation/tutorials/fulfillment/#countries-requiring-postal-codes) for more info (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.UpdateShopShippingProfile(context.Background(), shopId, shippingProfileId).Title(title).OriginCountryIso(originCountryIso).MinProcessingTime(minProcessingTime).MaxProcessingTime(maxProcessingTime).ProcessingTimeUnit(processingTimeUnit).OriginPostalCode(originPostalCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.UpdateShopShippingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShopShippingProfile`: ShopShippingProfile
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.UpdateShopShippingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShopShippingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **title** | **string** | The name string of this shipping profile. | 
 **originCountryIso** | **string** | The ISO code of the country from which the listing ships. | 
 **minProcessingTime** | **int64** | The minimum time required to process to ship listings with this shipping profile. | 
 **maxProcessingTime** | **int64** | The maximum processing time the listing needs to ship. | 
 **processingTimeUnit** | [**CreateShopShippingProfileRequestProcessingTimeUnit**](CreateShopShippingProfileRequestProcessingTimeUnit.md) |  | [default to &quot;business_days&quot;]
 **originPostalCode** | **string** | The postal code string (not necessarily a number) for the location from which the listing ships. Required if the &#x60;origin_country_iso&#x60; supports postal codes. See the [Fulfillment Tutorial docs](https://developer.etsy.com/documentation/tutorials/fulfillment/#countries-requiring-postal-codes) for more info | 

### Return type

[**ShopShippingProfile**](ShopShippingProfile.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShopShippingProfileDestination

> ShopShippingProfileDestination UpdateShopShippingProfileDestination(ctx, shopId, shippingProfileId, shippingProfileDestinationId).PrimaryCost(primaryCost).SecondaryCost(secondaryCost).DestinationCountryIso(destinationCountryIso).DestinationRegion(destinationRegion).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.
	shippingProfileDestinationId := int64(56) // int64 | The numeric ID of the shipping profile destination in the [shipping profile](/documentation/reference#tag/Shop-ShippingProfile) associated with the listing.
	primaryCost := float32(3.4) // float32 | The cost of shipping to this country/region alone, measured in the store's default currency. (optional)
	secondaryCost := float32(3.4) // float32 | The cost of shipping to this country/region with another item, measured in the store's default currency. (optional)
	destinationCountryIso := "destinationCountryIso_example" // string | The ISO code of the country to which the listing ships. If null, request sets destination to destination_region. Required if destination_region is null or not provided. (optional)
	destinationRegion := openapiclient.createShopShippingProfile_request_destination_region("eu") // CreateShopShippingProfileRequestDestinationRegion |  (optional) (default to "none")
	shippingCarrierId := int64(56) // int64 | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with `mail_class`** if `min_delivery_days` and `max_delivery_days` are null. (optional)
	mailClass := "mailClass_example" // string | The unique ID string of a shipping carrier's mail class, which is used to calculate an estimated delivery date. **Required with `shipping_carrier_id`** if `min_delivery_days` and `max_delivery_days` are null. (optional)
	minDeliveryDays := int64(56) // int64 | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `max_delivery_days`** if `mail_class` is null. (optional)
	maxDeliveryDays := int64(56) // int64 | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `min_delivery_days`** if `mail_class` is null. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.UpdateShopShippingProfileDestination(context.Background(), shopId, shippingProfileId, shippingProfileDestinationId).PrimaryCost(primaryCost).SecondaryCost(secondaryCost).DestinationCountryIso(destinationCountryIso).DestinationRegion(destinationRegion).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.UpdateShopShippingProfileDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShopShippingProfileDestination`: ShopShippingProfileDestination
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.UpdateShopShippingProfileDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 
**shippingProfileDestinationId** | **int64** | The numeric ID of the shipping profile destination in the [shipping profile](/documentation/reference#tag/Shop-ShippingProfile) associated with the listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShopShippingProfileDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **primaryCost** | **float32** | The cost of shipping to this country/region alone, measured in the store&#39;s default currency. | 
 **secondaryCost** | **float32** | The cost of shipping to this country/region with another item, measured in the store&#39;s default currency. | 
 **destinationCountryIso** | **string** | The ISO code of the country to which the listing ships. If null, request sets destination to destination_region. Required if destination_region is null or not provided. | 
 **destinationRegion** | [**CreateShopShippingProfileRequestDestinationRegion**](CreateShopShippingProfileRequestDestinationRegion.md) |  | [default to &quot;none&quot;]
 **shippingCarrierId** | **int64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | 
 **mailClass** | **string** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | 
 **minDeliveryDays** | **int64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 
 **maxDeliveryDays** | **int64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 

### Return type

[**ShopShippingProfileDestination**](ShopShippingProfileDestination.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShopShippingProfileUpgrade

> ShopShippingProfileUpgrade UpdateShopShippingProfileUpgrade(ctx, shopId, shippingProfileId, upgradeId).UpgradeName(upgradeName).Type_(type_).Price(price).SecondaryPrice(secondaryPrice).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gearment/go-etsy/client"
)

func main() {
	shopId := int64(56) // int64 | The unique positive non-zero numeric ID for an Etsy Shop.
	shippingProfileId := int64(56) // int64 | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is `physical`.
	upgradeId := int64(56) // int64 | The numeric ID that is associated with a shipping upgrade
	upgradeName := "upgradeName_example" // string | Name for the shipping upgrade shown to shoppers at checkout, e.g. USPS Priority. (optional)
	type_ := openapiclient.createShopShippingProfileUpgrade_request_type("0") // CreateShopShippingProfileUpgradeRequestType |  (optional)
	price := float32(3.4) // float32 | Additional cost of adding the shipping upgrade. (optional)
	secondaryPrice := float32(3.4) // float32 | Additional cost of adding the shipping upgrade for each additional item. (optional)
	shippingCarrierId := int64(56) // int64 | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with `mail_class`** if `min_delivery_days` and `max_delivery_days` are null. (optional)
	mailClass := "mailClass_example" // string | The unique ID string of a shipping carrier's mail class, which is used to calculate an estimated delivery date. **Required with `shipping_carrier_id`** if `min_delivery_days` and `max_delivery_days` are null. (optional)
	minDeliveryDays := int64(56) // int64 | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `max_delivery_days`** if `mail_class` is null. (optional)
	maxDeliveryDays := int64(56) // int64 | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with `min_delivery_days`** if `mail_class` is null. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShopShippingProfileAPI.UpdateShopShippingProfileUpgrade(context.Background(), shopId, shippingProfileId, upgradeId).UpgradeName(upgradeName).Type_(type_).Price(price).SecondaryPrice(secondaryPrice).ShippingCarrierId(shippingCarrierId).MailClass(mailClass).MinDeliveryDays(minDeliveryDays).MaxDeliveryDays(maxDeliveryDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShopShippingProfileAPI.UpdateShopShippingProfileUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShopShippingProfileUpgrade`: ShopShippingProfileUpgrade
	fmt.Fprintf(os.Stdout, "Response from `ShopShippingProfileAPI.UpdateShopShippingProfileUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **int64** | The unique positive non-zero numeric ID for an Etsy Shop. | 
**shippingProfileId** | **int64** | The numeric ID of the [shipping profile](/documentation/reference#operation/getShopShippingProfile) associated with the listing. Required when listing type is &#x60;physical&#x60;. | 
**upgradeId** | **int64** | The numeric ID that is associated with a shipping upgrade | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShopShippingProfileUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **upgradeName** | **string** | Name for the shipping upgrade shown to shoppers at checkout, e.g. USPS Priority. | 
 **type_** | [**CreateShopShippingProfileUpgradeRequestType**](CreateShopShippingProfileUpgradeRequestType.md) |  | 
 **price** | **float32** | Additional cost of adding the shipping upgrade. | 
 **secondaryPrice** | **float32** | Additional cost of adding the shipping upgrade for each additional item. | 
 **shippingCarrierId** | **int64** | The unique ID of a supported shipping carrier, which is used to calculate an Estimated Delivery Date. **Required with &#x60;mail_class&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | 
 **mailClass** | **string** | The unique ID string of a shipping carrier&#39;s mail class, which is used to calculate an estimated delivery date. **Required with &#x60;shipping_carrier_id&#x60;** if &#x60;min_delivery_days&#x60; and &#x60;max_delivery_days&#x60; are null. | 
 **minDeliveryDays** | **int64** | The minimum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;max_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 
 **maxDeliveryDays** | **int64** | The maximum number of business days a buyer can expect to wait to receive their purchased item once it has shipped. **Required with &#x60;min_delivery_days&#x60;** if &#x60;mail_class&#x60; is null. | 

### Return type

[**ShopShippingProfileUpgrade**](ShopShippingProfileUpgrade.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

