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

// checks if the Pong type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pong{}

// Pong A confirmation that the current application has access to the Open API
type Pong struct {
	// The authenticated application's ID
	ApplicationId        *int64 `json:"application_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Pong Pong

// NewPong instantiates a new Pong object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPong() *Pong {
	this := Pong{}
	return &this
}

// NewPongWithDefaults instantiates a new Pong object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPongWithDefaults() *Pong {
	this := Pong{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *Pong) GetApplicationId() int64 {
	if o == nil || IsNil(o.ApplicationId) {
		var ret int64
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pong) GetApplicationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *Pong) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given int64 and assigns it to the ApplicationId field.
func (o *Pong) SetApplicationId(v int64) {
	o.ApplicationId = &v
}

func (o Pong) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pong) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullablePong struct {
	value *Pong
	isSet bool
}

func (v NullablePong) Get() *Pong {
	return v.value
}

func (v *NullablePong) Set(val *Pong) {
	v.value = val
	v.isSet = true
}

func (v NullablePong) IsSet() bool {
	return v.isSet
}

func (v *NullablePong) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePong(val *Pong) *NullablePong {
	return &NullablePong{value: val, isSet: true}
}

func (v NullablePong) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePong) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
