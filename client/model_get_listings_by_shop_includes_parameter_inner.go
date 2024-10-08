/*
Etsy Open API v3

<div class=\"wt-text-body-01\"><p class=\"wt-pt-xs-2 wt-pb-xs-2\">Etsy's Open API provides a simple RESTful interface for various Etsy.com features. The API endpoints are meant to replace Etsy's Open API v2, which is scheduled to end service in 2022.</p><p class=\"wt-pb-xs-2\">All of the endpoints are callable and the majority of the API endpoints are now in a beta phase. This means we do not expect to make any breaking changes before our general release. A handful of endpoints are currently interface stubs (labeled “Feedback Only”) and returns a \"501 Not Implemented\" response code when called.</p><p class=\"wt-pb-xs-2\">If you'd like to report an issue or provide feedback on the API design, <a target=\"_blank\" class=\"wt-text-link wt-p-xs-0\" href=\"https://github.com/etsy/open-api/discussions\">please add an issue in Github</a>.</p></div>&copy; 2021-2024 Etsy, Inc. All Rights Reserved. Use of this code is subject to Etsy's <a class='wt-text-link wt-p-xs-0' target='_blank' href='https://www.etsy.com/legal/api'>API Developer Terms of Use</a>.

API version: 3.0.0
Contact: developers@etsy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goEtsy

import (
	"encoding/json"
	"fmt"
)

// GetListingsByShopIncludesParameterInner the model 'GetListingsByShopIncludesParameterInner'
type GetListingsByShopIncludesParameterInner string

// List of getListingsByShop_includes_parameter_inner
const (
	GETLISTINGSBYSHOPINCLUDESPARAMETERINNER_SHIPPING     GetListingsByShopIncludesParameterInner = "Shipping"
	GETLISTINGSBYSHOPINCLUDESPARAMETERINNER_IMAGES       GetListingsByShopIncludesParameterInner = "Images"
	GETLISTINGSBYSHOPINCLUDESPARAMETERINNER_SHOP         GetListingsByShopIncludesParameterInner = "Shop"
	GETLISTINGSBYSHOPINCLUDESPARAMETERINNER_USER         GetListingsByShopIncludesParameterInner = "User"
	GETLISTINGSBYSHOPINCLUDESPARAMETERINNER_TRANSLATIONS GetListingsByShopIncludesParameterInner = "Translations"
	GETLISTINGSBYSHOPINCLUDESPARAMETERINNER_INVENTORY    GetListingsByShopIncludesParameterInner = "Inventory"
	GETLISTINGSBYSHOPINCLUDESPARAMETERINNER_VIDEOS       GetListingsByShopIncludesParameterInner = "Videos"
)

// All allowed values of GetListingsByShopIncludesParameterInner enum
var AllowedGetListingsByShopIncludesParameterInnerEnumValues = []GetListingsByShopIncludesParameterInner{
	"Shipping",
	"Images",
	"Shop",
	"User",
	"Translations",
	"Inventory",
	"Videos",
}

func (v *GetListingsByShopIncludesParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetListingsByShopIncludesParameterInner(value)
	for _, existing := range AllowedGetListingsByShopIncludesParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetListingsByShopIncludesParameterInner", value)
}

// NewGetListingsByShopIncludesParameterInnerFromValue returns a pointer to a valid GetListingsByShopIncludesParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetListingsByShopIncludesParameterInnerFromValue(v string) (*GetListingsByShopIncludesParameterInner, error) {
	ev := GetListingsByShopIncludesParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetListingsByShopIncludesParameterInner: valid values are %v", v, AllowedGetListingsByShopIncludesParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetListingsByShopIncludesParameterInner) IsValid() bool {
	for _, existing := range AllowedGetListingsByShopIncludesParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getListingsByShop_includes_parameter_inner value
func (v GetListingsByShopIncludesParameterInner) Ptr() *GetListingsByShopIncludesParameterInner {
	return &v
}

type NullableGetListingsByShopIncludesParameterInner struct {
	value *GetListingsByShopIncludesParameterInner
	isSet bool
}

func (v NullableGetListingsByShopIncludesParameterInner) Get() *GetListingsByShopIncludesParameterInner {
	return v.value
}

func (v *NullableGetListingsByShopIncludesParameterInner) Set(val *GetListingsByShopIncludesParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListingsByShopIncludesParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListingsByShopIncludesParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetListingsByShopIncludesParameterInner(val *GetListingsByShopIncludesParameterInner) *NullableGetListingsByShopIncludesParameterInner {
	return &NullableGetListingsByShopIncludesParameterInner{value: val, isSet: true}
}

func (v NullableGetListingsByShopIncludesParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetListingsByShopIncludesParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
