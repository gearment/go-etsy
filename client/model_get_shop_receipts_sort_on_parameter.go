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

// GetShopReceiptsSortOnParameter The value to sort a search result of listings on.
type GetShopReceiptsSortOnParameter string

// List of getShopReceipts_sort_on_parameter
const (
	GETSHOPRECEIPTSSORTONPARAMETER_CREATED    GetShopReceiptsSortOnParameter = "created"
	GETSHOPRECEIPTSSORTONPARAMETER_UPDATED    GetShopReceiptsSortOnParameter = "updated"
	GETSHOPRECEIPTSSORTONPARAMETER_RECEIPT_ID GetShopReceiptsSortOnParameter = "receipt_id"
)

// All allowed values of GetShopReceiptsSortOnParameter enum
var AllowedGetShopReceiptsSortOnParameterEnumValues = []GetShopReceiptsSortOnParameter{
	"created",
	"updated",
	"receipt_id",
}

func (v *GetShopReceiptsSortOnParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetShopReceiptsSortOnParameter(value)
	for _, existing := range AllowedGetShopReceiptsSortOnParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetShopReceiptsSortOnParameter", value)
}

// NewGetShopReceiptsSortOnParameterFromValue returns a pointer to a valid GetShopReceiptsSortOnParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetShopReceiptsSortOnParameterFromValue(v string) (*GetShopReceiptsSortOnParameter, error) {
	ev := GetShopReceiptsSortOnParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetShopReceiptsSortOnParameter: valid values are %v", v, AllowedGetShopReceiptsSortOnParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetShopReceiptsSortOnParameter) IsValid() bool {
	for _, existing := range AllowedGetShopReceiptsSortOnParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getShopReceipts_sort_on_parameter value
func (v GetShopReceiptsSortOnParameter) Ptr() *GetShopReceiptsSortOnParameter {
	return &v
}

type NullableGetShopReceiptsSortOnParameter struct {
	value *GetShopReceiptsSortOnParameter
	isSet bool
}

func (v NullableGetShopReceiptsSortOnParameter) Get() *GetShopReceiptsSortOnParameter {
	return v.value
}

func (v *NullableGetShopReceiptsSortOnParameter) Set(val *GetShopReceiptsSortOnParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShopReceiptsSortOnParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShopReceiptsSortOnParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShopReceiptsSortOnParameter(val *GetShopReceiptsSortOnParameter) *NullableGetShopReceiptsSortOnParameter {
	return &NullableGetShopReceiptsSortOnParameter{value: val, isSet: true}
}

func (v NullableGetShopReceiptsSortOnParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShopReceiptsSortOnParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
