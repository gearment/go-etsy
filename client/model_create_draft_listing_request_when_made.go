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

// CreateDraftListingRequestWhenMade An enumerated string for the era in which the maker made the product in this listing. Helps buyers locate the listing under the Vintage heading. Requires 'is_supply' and 'who_made'.
type CreateDraftListingRequestWhenMade string

// List of createDraftListing_request_when_made
const (
	CREATEDRAFTLISTINGREQUESTWHENMADE_MADE_TO_ORDER CreateDraftListingRequestWhenMade = "made_to_order"
	CREATEDRAFTLISTINGREQUESTWHENMADE__2020_2024    CreateDraftListingRequestWhenMade = "2020_2024"
	CREATEDRAFTLISTINGREQUESTWHENMADE__2010_2019    CreateDraftListingRequestWhenMade = "2010_2019"
	CREATEDRAFTLISTINGREQUESTWHENMADE__2005_2009    CreateDraftListingRequestWhenMade = "2005_2009"
	CREATEDRAFTLISTINGREQUESTWHENMADE_BEFORE_2005   CreateDraftListingRequestWhenMade = "before_2005"
	CREATEDRAFTLISTINGREQUESTWHENMADE__2000_2004    CreateDraftListingRequestWhenMade = "2000_2004"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1990S        CreateDraftListingRequestWhenMade = "1990s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1980S        CreateDraftListingRequestWhenMade = "1980s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1970S        CreateDraftListingRequestWhenMade = "1970s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1960S        CreateDraftListingRequestWhenMade = "1960s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1950S        CreateDraftListingRequestWhenMade = "1950s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1940S        CreateDraftListingRequestWhenMade = "1940s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1930S        CreateDraftListingRequestWhenMade = "1930s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1920S        CreateDraftListingRequestWhenMade = "1920s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1910S        CreateDraftListingRequestWhenMade = "1910s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1900S        CreateDraftListingRequestWhenMade = "1900s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1800S        CreateDraftListingRequestWhenMade = "1800s"
	CREATEDRAFTLISTINGREQUESTWHENMADE__1700S        CreateDraftListingRequestWhenMade = "1700s"
	CREATEDRAFTLISTINGREQUESTWHENMADE_BEFORE_1700   CreateDraftListingRequestWhenMade = "before_1700"
)

// All allowed values of CreateDraftListingRequestWhenMade enum
var AllowedCreateDraftListingRequestWhenMadeEnumValues = []CreateDraftListingRequestWhenMade{
	"made_to_order",
	"2020_2024",
	"2010_2019",
	"2005_2009",
	"before_2005",
	"2000_2004",
	"1990s",
	"1980s",
	"1970s",
	"1960s",
	"1950s",
	"1940s",
	"1930s",
	"1920s",
	"1910s",
	"1900s",
	"1800s",
	"1700s",
	"before_1700",
}

func (v *CreateDraftListingRequestWhenMade) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateDraftListingRequestWhenMade(value)
	for _, existing := range AllowedCreateDraftListingRequestWhenMadeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateDraftListingRequestWhenMade", value)
}

// NewCreateDraftListingRequestWhenMadeFromValue returns a pointer to a valid CreateDraftListingRequestWhenMade
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateDraftListingRequestWhenMadeFromValue(v string) (*CreateDraftListingRequestWhenMade, error) {
	ev := CreateDraftListingRequestWhenMade(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateDraftListingRequestWhenMade: valid values are %v", v, AllowedCreateDraftListingRequestWhenMadeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateDraftListingRequestWhenMade) IsValid() bool {
	for _, existing := range AllowedCreateDraftListingRequestWhenMadeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to createDraftListing_request_when_made value
func (v CreateDraftListingRequestWhenMade) Ptr() *CreateDraftListingRequestWhenMade {
	return &v
}

type NullableCreateDraftListingRequestWhenMade struct {
	value *CreateDraftListingRequestWhenMade
	isSet bool
}

func (v NullableCreateDraftListingRequestWhenMade) Get() *CreateDraftListingRequestWhenMade {
	return v.value
}

func (v *NullableCreateDraftListingRequestWhenMade) Set(val *CreateDraftListingRequestWhenMade) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDraftListingRequestWhenMade) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDraftListingRequestWhenMade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDraftListingRequestWhenMade(val *CreateDraftListingRequestWhenMade) *NullableCreateDraftListingRequestWhenMade {
	return &NullableCreateDraftListingRequestWhenMade{value: val, isSet: true}
}

func (v NullableCreateDraftListingRequestWhenMade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDraftListingRequestWhenMade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
