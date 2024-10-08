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

// checks if the TransactionReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionReview{}

// TransactionReview A transaction review record left by a User.
type TransactionReview struct {
	// The shop's numeric ID.
	ShopId *int64 `json:"shop_id,omitempty"`
	// The ID of the ShopListing that the TransactionReview belongs to.
	ListingId *int64 `json:"listing_id,omitempty"`
	// The ID of the ShopReceipt Transaction that the TransactionReview belongs to.
	TransactionId *int64 `json:"transaction_id,omitempty"`
	// The numeric ID of the user who was the buyer in this transaction. Note: This field may be absent, depending on the buyer's privacy settings.
	BuyerUserId NullableInt64 `json:"buyer_user_id,omitempty"`
	// Rating value on scale from 1 to 5
	Rating *int64 `json:"rating,omitempty"`
	// A message left by the author, explaining the feedback, if provided.
	Review *string `json:"review,omitempty"`
	// The language of the TransactionReview
	Language *string `json:"language,omitempty"`
	// The url to a photo provided with the feedback, dimensions fullxfull. Note: This field may be absent, depending on the buyer's privacy settings.
	ImageUrlFullxfull NullableString `json:"image_url_fullxfull,omitempty"`
	// The date and time the TransactionReview was created in epoch seconds.
	CreateTimestamp *int64 `json:"create_timestamp,omitempty"`
	// The date and time the TransactionReview was created in epoch seconds.
	CreatedTimestamp *int64 `json:"created_timestamp,omitempty"`
	// The date and time the TransactionReview was updated in epoch seconds.
	UpdateTimestamp *int64 `json:"update_timestamp,omitempty"`
	// The date and time the TransactionReview was updated in epoch seconds.
	UpdatedTimestamp     *int64 `json:"updated_timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransactionReview TransactionReview

// NewTransactionReview instantiates a new TransactionReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionReview() *TransactionReview {
	this := TransactionReview{}
	var review string = ""
	this.Review = &review
	return &this
}

// NewTransactionReviewWithDefaults instantiates a new TransactionReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionReviewWithDefaults() *TransactionReview {
	this := TransactionReview{}
	var review string = ""
	this.Review = &review
	return &this
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *TransactionReview) GetShopId() int64 {
	if o == nil || IsNil(o.ShopId) {
		var ret int64
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetShopIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *TransactionReview) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int64 and assigns it to the ShopId field.
func (o *TransactionReview) SetShopId(v int64) {
	o.ShopId = &v
}

// GetListingId returns the ListingId field value if set, zero value otherwise.
func (o *TransactionReview) GetListingId() int64 {
	if o == nil || IsNil(o.ListingId) {
		var ret int64
		return ret
	}
	return *o.ListingId
}

// GetListingIdOk returns a tuple with the ListingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetListingIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ListingId) {
		return nil, false
	}
	return o.ListingId, true
}

// HasListingId returns a boolean if a field has been set.
func (o *TransactionReview) HasListingId() bool {
	if o != nil && !IsNil(o.ListingId) {
		return true
	}

	return false
}

// SetListingId gets a reference to the given int64 and assigns it to the ListingId field.
func (o *TransactionReview) SetListingId(v int64) {
	o.ListingId = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *TransactionReview) GetTransactionId() int64 {
	if o == nil || IsNil(o.TransactionId) {
		var ret int64
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetTransactionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *TransactionReview) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given int64 and assigns it to the TransactionId field.
func (o *TransactionReview) SetTransactionId(v int64) {
	o.TransactionId = &v
}

// GetBuyerUserId returns the BuyerUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionReview) GetBuyerUserId() int64 {
	if o == nil || IsNil(o.BuyerUserId.Get()) {
		var ret int64
		return ret
	}
	return *o.BuyerUserId.Get()
}

// GetBuyerUserIdOk returns a tuple with the BuyerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionReview) GetBuyerUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerUserId.Get(), o.BuyerUserId.IsSet()
}

// HasBuyerUserId returns a boolean if a field has been set.
func (o *TransactionReview) HasBuyerUserId() bool {
	if o != nil && o.BuyerUserId.IsSet() {
		return true
	}

	return false
}

// SetBuyerUserId gets a reference to the given NullableInt64 and assigns it to the BuyerUserId field.
func (o *TransactionReview) SetBuyerUserId(v int64) {
	o.BuyerUserId.Set(&v)
}

// SetBuyerUserIdNil sets the value for BuyerUserId to be an explicit nil
func (o *TransactionReview) SetBuyerUserIdNil() {
	o.BuyerUserId.Set(nil)
}

// UnsetBuyerUserId ensures that no value is present for BuyerUserId, not even an explicit nil
func (o *TransactionReview) UnsetBuyerUserId() {
	o.BuyerUserId.Unset()
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *TransactionReview) GetRating() int64 {
	if o == nil || IsNil(o.Rating) {
		var ret int64
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *TransactionReview) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int64 and assigns it to the Rating field.
func (o *TransactionReview) SetRating(v int64) {
	o.Rating = &v
}

// GetReview returns the Review field value if set, zero value otherwise.
func (o *TransactionReview) GetReview() string {
	if o == nil || IsNil(o.Review) {
		var ret string
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetReviewOk() (*string, bool) {
	if o == nil || IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *TransactionReview) HasReview() bool {
	if o != nil && !IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given string and assigns it to the Review field.
func (o *TransactionReview) SetReview(v string) {
	o.Review = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *TransactionReview) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *TransactionReview) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *TransactionReview) SetLanguage(v string) {
	o.Language = &v
}

// GetImageUrlFullxfull returns the ImageUrlFullxfull field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionReview) GetImageUrlFullxfull() string {
	if o == nil || IsNil(o.ImageUrlFullxfull.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrlFullxfull.Get()
}

// GetImageUrlFullxfullOk returns a tuple with the ImageUrlFullxfull field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionReview) GetImageUrlFullxfullOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrlFullxfull.Get(), o.ImageUrlFullxfull.IsSet()
}

// HasImageUrlFullxfull returns a boolean if a field has been set.
func (o *TransactionReview) HasImageUrlFullxfull() bool {
	if o != nil && o.ImageUrlFullxfull.IsSet() {
		return true
	}

	return false
}

// SetImageUrlFullxfull gets a reference to the given NullableString and assigns it to the ImageUrlFullxfull field.
func (o *TransactionReview) SetImageUrlFullxfull(v string) {
	o.ImageUrlFullxfull.Set(&v)
}

// SetImageUrlFullxfullNil sets the value for ImageUrlFullxfull to be an explicit nil
func (o *TransactionReview) SetImageUrlFullxfullNil() {
	o.ImageUrlFullxfull.Set(nil)
}

// UnsetImageUrlFullxfull ensures that no value is present for ImageUrlFullxfull, not even an explicit nil
func (o *TransactionReview) UnsetImageUrlFullxfull() {
	o.ImageUrlFullxfull.Unset()
}

// GetCreateTimestamp returns the CreateTimestamp field value if set, zero value otherwise.
func (o *TransactionReview) GetCreateTimestamp() int64 {
	if o == nil || IsNil(o.CreateTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreateTimestamp
}

// GetCreateTimestampOk returns a tuple with the CreateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetCreateTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimestamp) {
		return nil, false
	}
	return o.CreateTimestamp, true
}

// HasCreateTimestamp returns a boolean if a field has been set.
func (o *TransactionReview) HasCreateTimestamp() bool {
	if o != nil && !IsNil(o.CreateTimestamp) {
		return true
	}

	return false
}

// SetCreateTimestamp gets a reference to the given int64 and assigns it to the CreateTimestamp field.
func (o *TransactionReview) SetCreateTimestamp(v int64) {
	o.CreateTimestamp = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *TransactionReview) GetCreatedTimestamp() int64 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *TransactionReview) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int64 and assigns it to the CreatedTimestamp field.
func (o *TransactionReview) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = &v
}

// GetUpdateTimestamp returns the UpdateTimestamp field value if set, zero value otherwise.
func (o *TransactionReview) GetUpdateTimestamp() int64 {
	if o == nil || IsNil(o.UpdateTimestamp) {
		var ret int64
		return ret
	}
	return *o.UpdateTimestamp
}

// GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetUpdateTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTimestamp) {
		return nil, false
	}
	return o.UpdateTimestamp, true
}

// HasUpdateTimestamp returns a boolean if a field has been set.
func (o *TransactionReview) HasUpdateTimestamp() bool {
	if o != nil && !IsNil(o.UpdateTimestamp) {
		return true
	}

	return false
}

// SetUpdateTimestamp gets a reference to the given int64 and assigns it to the UpdateTimestamp field.
func (o *TransactionReview) SetUpdateTimestamp(v int64) {
	o.UpdateTimestamp = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value if set, zero value otherwise.
func (o *TransactionReview) GetUpdatedTimestamp() int64 {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReview) GetUpdatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		return nil, false
	}
	return o.UpdatedTimestamp, true
}

// HasUpdatedTimestamp returns a boolean if a field has been set.
func (o *TransactionReview) HasUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UpdatedTimestamp) {
		return true
	}

	return false
}

// SetUpdatedTimestamp gets a reference to the given int64 and assigns it to the UpdatedTimestamp field.
func (o *TransactionReview) SetUpdatedTimestamp(v int64) {
	o.UpdatedTimestamp = &v
}

func (o TransactionReview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.ListingId) {
		toSerialize["listing_id"] = o.ListingId
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if o.BuyerUserId.IsSet() {
		toSerialize["buyer_user_id"] = o.BuyerUserId.Get()
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.Review) {
		toSerialize["review"] = o.Review
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if o.ImageUrlFullxfull.IsSet() {
		toSerialize["image_url_fullxfull"] = o.ImageUrlFullxfull.Get()
	}
	if !IsNil(o.CreateTimestamp) {
		toSerialize["create_timestamp"] = o.CreateTimestamp
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.UpdateTimestamp) {
		toSerialize["update_timestamp"] = o.UpdateTimestamp
	}
	if !IsNil(o.UpdatedTimestamp) {
		toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableTransactionReview struct {
	value *TransactionReview
	isSet bool
}

func (v NullableTransactionReview) Get() *TransactionReview {
	return v.value
}

func (v *NullableTransactionReview) Set(val *TransactionReview) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionReview) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionReview(val *TransactionReview) *NullableTransactionReview {
	return &NullableTransactionReview{value: val, isSet: true}
}

func (v NullableTransactionReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
