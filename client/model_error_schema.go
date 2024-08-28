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

// checks if the ErrorSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorSchema{}

// ErrorSchema struct for ErrorSchema
type ErrorSchema struct {
	Error                string `json:"error"`
	AdditionalProperties map[string]interface{}
}

type _ErrorSchema ErrorSchema

// NewErrorSchema instantiates a new ErrorSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorSchema(error_ string) *ErrorSchema {
	this := ErrorSchema{}
	this.Error = error_
	return &this
}

// NewErrorSchemaWithDefaults instantiates a new ErrorSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorSchemaWithDefaults() *ErrorSchema {
	this := ErrorSchema{}
	return &this
}

// GetError returns the Error field value
func (o *ErrorSchema) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ErrorSchema) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ErrorSchema) SetError(v string) {
	o.Error = v
}

func (o ErrorSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableErrorSchema struct {
	value *ErrorSchema
	isSet bool
}

func (v NullableErrorSchema) Get() *ErrorSchema {
	return v.value
}

func (v *NullableErrorSchema) Set(val *ErrorSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSchema(val *ErrorSchema) *NullableErrorSchema {
	return &NullableErrorSchema{value: val, isSet: true}
}

func (v NullableErrorSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
