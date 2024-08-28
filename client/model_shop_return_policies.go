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
)

// checks if the ShopReturnPolicies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShopReturnPolicies{}

// ShopReturnPolicies Represents a shop's listing-level return policies list.
type ShopReturnPolicies struct {
	Count                *int64             `json:"count,omitempty"`
	Results              []ShopReturnPolicy `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ShopReturnPolicies ShopReturnPolicies

// NewShopReturnPolicies instantiates a new ShopReturnPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopReturnPolicies() *ShopReturnPolicies {
	this := ShopReturnPolicies{}
	return &this
}

// NewShopReturnPoliciesWithDefaults instantiates a new ShopReturnPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopReturnPoliciesWithDefaults() *ShopReturnPolicies {
	this := ShopReturnPolicies{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ShopReturnPolicies) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopReturnPolicies) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ShopReturnPolicies) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ShopReturnPolicies) SetCount(v int64) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ShopReturnPolicies) GetResults() []ShopReturnPolicy {
	if o == nil || IsNil(o.Results) {
		var ret []ShopReturnPolicy
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopReturnPolicies) GetResultsOk() ([]ShopReturnPolicy, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ShopReturnPolicies) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ShopReturnPolicy and assigns it to the Results field.
func (o *ShopReturnPolicies) SetResults(v []ShopReturnPolicy) {
	o.Results = v
}

func (o ShopReturnPolicies) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopReturnPolicies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableShopReturnPolicies struct {
	value *ShopReturnPolicies
	isSet bool
}

func (v NullableShopReturnPolicies) Get() *ShopReturnPolicies {
	return v.value
}

func (v *NullableShopReturnPolicies) Set(val *ShopReturnPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableShopReturnPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableShopReturnPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopReturnPolicies(val *ShopReturnPolicies) *NullableShopReturnPolicies {
	return &NullableShopReturnPolicies{value: val, isSet: true}
}

func (v NullableShopReturnPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopReturnPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
