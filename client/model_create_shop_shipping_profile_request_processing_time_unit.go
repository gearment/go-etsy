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

// CreateShopShippingProfileRequestProcessingTimeUnit The unit used to represent how long a processing time is. A week is equivalent to 5 business days. If none is provided, the unit is set to \"business_days\".
type CreateShopShippingProfileRequestProcessingTimeUnit string

// List of createShopShippingProfile_request_processing_time_unit
const (
	CREATESHOPSHIPPINGPROFILEREQUESTPROCESSINGTIMEUNIT_BUSINESS_DAYS CreateShopShippingProfileRequestProcessingTimeUnit = "business_days"
	CREATESHOPSHIPPINGPROFILEREQUESTPROCESSINGTIMEUNIT_WEEKS         CreateShopShippingProfileRequestProcessingTimeUnit = "weeks"
)

// All allowed values of CreateShopShippingProfileRequestProcessingTimeUnit enum
var AllowedCreateShopShippingProfileRequestProcessingTimeUnitEnumValues = []CreateShopShippingProfileRequestProcessingTimeUnit{
	"business_days",
	"weeks",
}

func (v *CreateShopShippingProfileRequestProcessingTimeUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateShopShippingProfileRequestProcessingTimeUnit(value)
	for _, existing := range AllowedCreateShopShippingProfileRequestProcessingTimeUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateShopShippingProfileRequestProcessingTimeUnit", value)
}

// NewCreateShopShippingProfileRequestProcessingTimeUnitFromValue returns a pointer to a valid CreateShopShippingProfileRequestProcessingTimeUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateShopShippingProfileRequestProcessingTimeUnitFromValue(v string) (*CreateShopShippingProfileRequestProcessingTimeUnit, error) {
	ev := CreateShopShippingProfileRequestProcessingTimeUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateShopShippingProfileRequestProcessingTimeUnit: valid values are %v", v, AllowedCreateShopShippingProfileRequestProcessingTimeUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateShopShippingProfileRequestProcessingTimeUnit) IsValid() bool {
	for _, existing := range AllowedCreateShopShippingProfileRequestProcessingTimeUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to createShopShippingProfile_request_processing_time_unit value
func (v CreateShopShippingProfileRequestProcessingTimeUnit) Ptr() *CreateShopShippingProfileRequestProcessingTimeUnit {
	return &v
}

type NullableCreateShopShippingProfileRequestProcessingTimeUnit struct {
	value *CreateShopShippingProfileRequestProcessingTimeUnit
	isSet bool
}

func (v NullableCreateShopShippingProfileRequestProcessingTimeUnit) Get() *CreateShopShippingProfileRequestProcessingTimeUnit {
	return v.value
}

func (v *NullableCreateShopShippingProfileRequestProcessingTimeUnit) Set(val *CreateShopShippingProfileRequestProcessingTimeUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShopShippingProfileRequestProcessingTimeUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShopShippingProfileRequestProcessingTimeUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShopShippingProfileRequestProcessingTimeUnit(val *CreateShopShippingProfileRequestProcessingTimeUnit) *NullableCreateShopShippingProfileRequestProcessingTimeUnit {
	return &NullableCreateShopShippingProfileRequestProcessingTimeUnit{value: val, isSet: true}
}

func (v NullableCreateShopShippingProfileRequestProcessingTimeUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShopShippingProfileRequestProcessingTimeUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
