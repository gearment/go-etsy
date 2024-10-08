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

// UpdateListingDeprecatedRequestType An enumerated type string that indicates whether the listing is physical or a digital download.
type UpdateListingDeprecatedRequestType string

// List of updateListingDeprecated_request_type
const (
	UPDATELISTINGDEPRECATEDREQUESTTYPE_PHYSICAL UpdateListingDeprecatedRequestType = "physical"
	UPDATELISTINGDEPRECATEDREQUESTTYPE_DOWNLOAD UpdateListingDeprecatedRequestType = "download"
	UPDATELISTINGDEPRECATEDREQUESTTYPE_BOTH     UpdateListingDeprecatedRequestType = "both"
)

// All allowed values of UpdateListingDeprecatedRequestType enum
var AllowedUpdateListingDeprecatedRequestTypeEnumValues = []UpdateListingDeprecatedRequestType{
	"physical",
	"download",
	"both",
}

func (v *UpdateListingDeprecatedRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpdateListingDeprecatedRequestType(value)
	for _, existing := range AllowedUpdateListingDeprecatedRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpdateListingDeprecatedRequestType", value)
}

// NewUpdateListingDeprecatedRequestTypeFromValue returns a pointer to a valid UpdateListingDeprecatedRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpdateListingDeprecatedRequestTypeFromValue(v string) (*UpdateListingDeprecatedRequestType, error) {
	ev := UpdateListingDeprecatedRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpdateListingDeprecatedRequestType: valid values are %v", v, AllowedUpdateListingDeprecatedRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpdateListingDeprecatedRequestType) IsValid() bool {
	for _, existing := range AllowedUpdateListingDeprecatedRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to updateListingDeprecated_request_type value
func (v UpdateListingDeprecatedRequestType) Ptr() *UpdateListingDeprecatedRequestType {
	return &v
}

type NullableUpdateListingDeprecatedRequestType struct {
	value *UpdateListingDeprecatedRequestType
	isSet bool
}

func (v NullableUpdateListingDeprecatedRequestType) Get() *UpdateListingDeprecatedRequestType {
	return v.value
}

func (v *NullableUpdateListingDeprecatedRequestType) Set(val *UpdateListingDeprecatedRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateListingDeprecatedRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateListingDeprecatedRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateListingDeprecatedRequestType(val *UpdateListingDeprecatedRequestType) *NullableUpdateListingDeprecatedRequestType {
	return &NullableUpdateListingDeprecatedRequestType{value: val, isSet: true}
}

func (v NullableUpdateListingDeprecatedRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateListingDeprecatedRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
