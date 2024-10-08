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

// checks if the UpdateListingInventoryRequestProductsInnerOfferingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateListingInventoryRequestProductsInnerOfferingsInner{}

// UpdateListingInventoryRequestProductsInnerOfferingsInner struct for UpdateListingInventoryRequestProductsInnerOfferingsInner
type UpdateListingInventoryRequestProductsInnerOfferingsInner struct {
	// The price of the product.
	Price float32 `json:"price"`
	// How many of this product are available?
	Quantity int64 `json:"quantity"`
	// True if the offering is shown to buyers
	IsEnabled            bool `json:"is_enabled"`
	AdditionalProperties map[string]interface{}
}

type _UpdateListingInventoryRequestProductsInnerOfferingsInner UpdateListingInventoryRequestProductsInnerOfferingsInner

// NewUpdateListingInventoryRequestProductsInnerOfferingsInner instantiates a new UpdateListingInventoryRequestProductsInnerOfferingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateListingInventoryRequestProductsInnerOfferingsInner(price float32, quantity int64, isEnabled bool) *UpdateListingInventoryRequestProductsInnerOfferingsInner {
	this := UpdateListingInventoryRequestProductsInnerOfferingsInner{}
	this.Price = price
	this.Quantity = quantity
	this.IsEnabled = isEnabled
	return &this
}

// NewUpdateListingInventoryRequestProductsInnerOfferingsInnerWithDefaults instantiates a new UpdateListingInventoryRequestProductsInnerOfferingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateListingInventoryRequestProductsInnerOfferingsInnerWithDefaults() *UpdateListingInventoryRequestProductsInnerOfferingsInner {
	this := UpdateListingInventoryRequestProductsInnerOfferingsInner{}
	return &this
}

// GetPrice returns the Price field value
func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) SetPrice(v float32) {
	o.Price = v
}

// GetQuantity returns the Quantity field value
func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetQuantity() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetQuantityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) SetQuantity(v int64) {
	o.Quantity = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *UpdateListingInventoryRequestProductsInnerOfferingsInner) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

func (o UpdateListingInventoryRequestProductsInnerOfferingsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateListingInventoryRequestProductsInnerOfferingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["price"] = o.Price
	toSerialize["quantity"] = o.Quantity
	toSerialize["is_enabled"] = o.IsEnabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableUpdateListingInventoryRequestProductsInnerOfferingsInner struct {
	value *UpdateListingInventoryRequestProductsInnerOfferingsInner
	isSet bool
}

func (v NullableUpdateListingInventoryRequestProductsInnerOfferingsInner) Get() *UpdateListingInventoryRequestProductsInnerOfferingsInner {
	return v.value
}

func (v *NullableUpdateListingInventoryRequestProductsInnerOfferingsInner) Set(val *UpdateListingInventoryRequestProductsInnerOfferingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateListingInventoryRequestProductsInnerOfferingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateListingInventoryRequestProductsInnerOfferingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateListingInventoryRequestProductsInnerOfferingsInner(val *UpdateListingInventoryRequestProductsInnerOfferingsInner) *NullableUpdateListingInventoryRequestProductsInnerOfferingsInner {
	return &NullableUpdateListingInventoryRequestProductsInnerOfferingsInner{value: val, isSet: true}
}

func (v NullableUpdateListingInventoryRequestProductsInnerOfferingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateListingInventoryRequestProductsInnerOfferingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
