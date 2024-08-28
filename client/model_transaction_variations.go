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

// checks if the TransactionVariations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionVariations{}

// TransactionVariations A list of variations chosen by the buyer during checkout.
type TransactionVariations struct {
	// The variation property ID.
	PropertyId *int64 `json:"property_id,omitempty"`
	// The ID of the variation value selected.
	ValueId NullableInt64 `json:"value_id,omitempty"`
	// Formatted name of the variation.
	FormattedName *string `json:"formatted_name,omitempty"`
	// Value of the variation entered by the buyer.
	FormattedValue       *string `json:"formatted_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransactionVariations TransactionVariations

// NewTransactionVariations instantiates a new TransactionVariations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionVariations() *TransactionVariations {
	this := TransactionVariations{}
	return &this
}

// NewTransactionVariationsWithDefaults instantiates a new TransactionVariations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionVariationsWithDefaults() *TransactionVariations {
	this := TransactionVariations{}
	return &this
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *TransactionVariations) GetPropertyId() int64 {
	if o == nil || IsNil(o.PropertyId) {
		var ret int64
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionVariations) GetPropertyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PropertyId) {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *TransactionVariations) HasPropertyId() bool {
	if o != nil && !IsNil(o.PropertyId) {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given int64 and assigns it to the PropertyId field.
func (o *TransactionVariations) SetPropertyId(v int64) {
	o.PropertyId = &v
}

// GetValueId returns the ValueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionVariations) GetValueId() int64 {
	if o == nil || IsNil(o.ValueId.Get()) {
		var ret int64
		return ret
	}
	return *o.ValueId.Get()
}

// GetValueIdOk returns a tuple with the ValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionVariations) GetValueIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueId.Get(), o.ValueId.IsSet()
}

// HasValueId returns a boolean if a field has been set.
func (o *TransactionVariations) HasValueId() bool {
	if o != nil && o.ValueId.IsSet() {
		return true
	}

	return false
}

// SetValueId gets a reference to the given NullableInt64 and assigns it to the ValueId field.
func (o *TransactionVariations) SetValueId(v int64) {
	o.ValueId.Set(&v)
}

// SetValueIdNil sets the value for ValueId to be an explicit nil
func (o *TransactionVariations) SetValueIdNil() {
	o.ValueId.Set(nil)
}

// UnsetValueId ensures that no value is present for ValueId, not even an explicit nil
func (o *TransactionVariations) UnsetValueId() {
	o.ValueId.Unset()
}

// GetFormattedName returns the FormattedName field value if set, zero value otherwise.
func (o *TransactionVariations) GetFormattedName() string {
	if o == nil || IsNil(o.FormattedName) {
		var ret string
		return ret
	}
	return *o.FormattedName
}

// GetFormattedNameOk returns a tuple with the FormattedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionVariations) GetFormattedNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormattedName) {
		return nil, false
	}
	return o.FormattedName, true
}

// HasFormattedName returns a boolean if a field has been set.
func (o *TransactionVariations) HasFormattedName() bool {
	if o != nil && !IsNil(o.FormattedName) {
		return true
	}

	return false
}

// SetFormattedName gets a reference to the given string and assigns it to the FormattedName field.
func (o *TransactionVariations) SetFormattedName(v string) {
	o.FormattedName = &v
}

// GetFormattedValue returns the FormattedValue field value if set, zero value otherwise.
func (o *TransactionVariations) GetFormattedValue() string {
	if o == nil || IsNil(o.FormattedValue) {
		var ret string
		return ret
	}
	return *o.FormattedValue
}

// GetFormattedValueOk returns a tuple with the FormattedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionVariations) GetFormattedValueOk() (*string, bool) {
	if o == nil || IsNil(o.FormattedValue) {
		return nil, false
	}
	return o.FormattedValue, true
}

// HasFormattedValue returns a boolean if a field has been set.
func (o *TransactionVariations) HasFormattedValue() bool {
	if o != nil && !IsNil(o.FormattedValue) {
		return true
	}

	return false
}

// SetFormattedValue gets a reference to the given string and assigns it to the FormattedValue field.
func (o *TransactionVariations) SetFormattedValue(v string) {
	o.FormattedValue = &v
}

func (o TransactionVariations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionVariations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyId) {
		toSerialize["property_id"] = o.PropertyId
	}
	if o.ValueId.IsSet() {
		toSerialize["value_id"] = o.ValueId.Get()
	}
	if !IsNil(o.FormattedName) {
		toSerialize["formatted_name"] = o.FormattedName
	}
	if !IsNil(o.FormattedValue) {
		toSerialize["formatted_value"] = o.FormattedValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableTransactionVariations struct {
	value *TransactionVariations
	isSet bool
}

func (v NullableTransactionVariations) Get() *TransactionVariations {
	return v.value
}

func (v *NullableTransactionVariations) Set(val *TransactionVariations) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionVariations) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionVariations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionVariations(val *TransactionVariations) *NullableTransactionVariations {
	return &NullableTransactionVariations{value: val, isSet: true}
}

func (v NullableTransactionVariations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionVariations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
