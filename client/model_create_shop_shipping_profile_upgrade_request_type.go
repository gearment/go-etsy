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

// CreateShopShippingProfileUpgradeRequestType The type of the shipping upgrade. Domestic (0) or international (1).
type CreateShopShippingProfileUpgradeRequestType string

// List of createShopShippingProfileUpgrade_request_type
const (
	CREATESHOPSHIPPINGPROFILEUPGRADEREQUESTTYPE__0 CreateShopShippingProfileUpgradeRequestType = "0"
	CREATESHOPSHIPPINGPROFILEUPGRADEREQUESTTYPE__1 CreateShopShippingProfileUpgradeRequestType = "1"
)

// All allowed values of CreateShopShippingProfileUpgradeRequestType enum
var AllowedCreateShopShippingProfileUpgradeRequestTypeEnumValues = []CreateShopShippingProfileUpgradeRequestType{
	"0",
	"1",
}

func (v *CreateShopShippingProfileUpgradeRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateShopShippingProfileUpgradeRequestType(value)
	for _, existing := range AllowedCreateShopShippingProfileUpgradeRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateShopShippingProfileUpgradeRequestType", value)
}

// NewCreateShopShippingProfileUpgradeRequestTypeFromValue returns a pointer to a valid CreateShopShippingProfileUpgradeRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateShopShippingProfileUpgradeRequestTypeFromValue(v string) (*CreateShopShippingProfileUpgradeRequestType, error) {
	ev := CreateShopShippingProfileUpgradeRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateShopShippingProfileUpgradeRequestType: valid values are %v", v, AllowedCreateShopShippingProfileUpgradeRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateShopShippingProfileUpgradeRequestType) IsValid() bool {
	for _, existing := range AllowedCreateShopShippingProfileUpgradeRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to createShopShippingProfileUpgrade_request_type value
func (v CreateShopShippingProfileUpgradeRequestType) Ptr() *CreateShopShippingProfileUpgradeRequestType {
	return &v
}

type NullableCreateShopShippingProfileUpgradeRequestType struct {
	value *CreateShopShippingProfileUpgradeRequestType
	isSet bool
}

func (v NullableCreateShopShippingProfileUpgradeRequestType) Get() *CreateShopShippingProfileUpgradeRequestType {
	return v.value
}

func (v *NullableCreateShopShippingProfileUpgradeRequestType) Set(val *CreateShopShippingProfileUpgradeRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShopShippingProfileUpgradeRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShopShippingProfileUpgradeRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShopShippingProfileUpgradeRequestType(val *CreateShopShippingProfileUpgradeRequestType) *NullableCreateShopShippingProfileUpgradeRequestType {
	return &NullableCreateShopShippingProfileUpgradeRequestType{value: val, isSet: true}
}

func (v NullableCreateShopShippingProfileUpgradeRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShopShippingProfileUpgradeRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
