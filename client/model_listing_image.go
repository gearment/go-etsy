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

// checks if the ListingImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingImage{}

// ListingImage Reference urls and metadata for an image associated with a specific listing. The `url_fullxfull` parameter contains the URL for full-sized binary image file.
type ListingImage struct {
	// The numeric ID for the [listing](/documentation/reference#tag/ShopListing) associated to this transaction.
	ListingId *int64 `json:"listing_id,omitempty"`
	// The numeric ID of the primary [listing image](/documentation/reference#tag/ShopListing-Image) for this transaction.
	ListingImageId *int64 `json:"listing_image_id,omitempty"`
	// The webhex string for the image's average color, in webhex notation.
	HexCode NullableString `json:"hex_code,omitempty"`
	// The numeric red value equal to the image's average red value, from 0-255 (RGB color).
	Red NullableInt64 `json:"red,omitempty"`
	// The numeric red value equal to the image's average red value, from 0-255 (RGB color).
	Green NullableInt64 `json:"green,omitempty"`
	// The numeric red value equal to the image's average red value, from 0-255 (RGB color).
	Blue NullableInt64 `json:"blue,omitempty"`
	// The numeric hue equal to the image's average hue, from 0-360 (HSV color).
	Hue NullableInt64 `json:"hue,omitempty"`
	// The numeric saturation equal to the image's average saturation, from 0-100 (HSV color).
	Saturation NullableInt64 `json:"saturation,omitempty"`
	// The numeric brightness equal to the image's average brightness, from 0-100 (HSV color).
	Brightness NullableInt64 `json:"brightness,omitempty"`
	// When true, the image is in black & white.
	IsBlackAndWhite NullableBool `json:"is_black_and_white,omitempty"`
	// The listing image\\'s creation time, in epoch seconds.
	CreationTsz *int64 `json:"creation_tsz,omitempty"`
	// The listing image\\'s creation time, in epoch seconds.
	CreatedTimestamp *int64 `json:"created_timestamp,omitempty"`
	// The positive non-zero numeric position in the images displayed in a listing, with rank 1 images appearing in the left-most position in a listing.
	Rank *int64 `json:"rank,omitempty"`
	// The url string for a 75x75 pixel thumbnail of the image.
	Url75x75 *string `json:"url_75x75,omitempty"`
	// The url string for a 170x135 pixel thumbnail of the image.
	Url170x135 *string `json:"url_170x135,omitempty"`
	// The url string for a thumbnail of the image, no more than 570 pixels wide with variable height.
	Url570xN *string `json:"url_570xN,omitempty"`
	// The url string for the full-size image, up to 3000 pixels in each dimension.
	UrlFullxfull *string `json:"url_fullxfull,omitempty"`
	// The numeric height, measured in pixels, of the full-sized image referenced in url_fullxfull.
	FullHeight NullableInt64 `json:"full_height,omitempty"`
	// The numeric width, measured in pixels, of the full-sized image referenced in url_fullxfull.
	FullWidth NullableInt64 `json:"full_width,omitempty"`
	// Alt text for the listing image. Max length 250 characters.
	AltText              NullableString `json:"alt_text,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListingImage ListingImage

// NewListingImage instantiates a new ListingImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingImage() *ListingImage {
	this := ListingImage{}
	return &this
}

// NewListingImageWithDefaults instantiates a new ListingImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingImageWithDefaults() *ListingImage {
	this := ListingImage{}
	return &this
}

// GetListingId returns the ListingId field value if set, zero value otherwise.
func (o *ListingImage) GetListingId() int64 {
	if o == nil || IsNil(o.ListingId) {
		var ret int64
		return ret
	}
	return *o.ListingId
}

// GetListingIdOk returns a tuple with the ListingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingImage) GetListingIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ListingId) {
		return nil, false
	}
	return o.ListingId, true
}

// HasListingId returns a boolean if a field has been set.
func (o *ListingImage) HasListingId() bool {
	if o != nil && !IsNil(o.ListingId) {
		return true
	}

	return false
}

// SetListingId gets a reference to the given int64 and assigns it to the ListingId field.
func (o *ListingImage) SetListingId(v int64) {
	o.ListingId = &v
}

// GetListingImageId returns the ListingImageId field value if set, zero value otherwise.
func (o *ListingImage) GetListingImageId() int64 {
	if o == nil || IsNil(o.ListingImageId) {
		var ret int64
		return ret
	}
	return *o.ListingImageId
}

// GetListingImageIdOk returns a tuple with the ListingImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingImage) GetListingImageIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ListingImageId) {
		return nil, false
	}
	return o.ListingImageId, true
}

// HasListingImageId returns a boolean if a field has been set.
func (o *ListingImage) HasListingImageId() bool {
	if o != nil && !IsNil(o.ListingImageId) {
		return true
	}

	return false
}

// SetListingImageId gets a reference to the given int64 and assigns it to the ListingImageId field.
func (o *ListingImage) SetListingImageId(v int64) {
	o.ListingImageId = &v
}

// GetHexCode returns the HexCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetHexCode() string {
	if o == nil || IsNil(o.HexCode.Get()) {
		var ret string
		return ret
	}
	return *o.HexCode.Get()
}

// GetHexCodeOk returns a tuple with the HexCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetHexCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HexCode.Get(), o.HexCode.IsSet()
}

// HasHexCode returns a boolean if a field has been set.
func (o *ListingImage) HasHexCode() bool {
	if o != nil && o.HexCode.IsSet() {
		return true
	}

	return false
}

// SetHexCode gets a reference to the given NullableString and assigns it to the HexCode field.
func (o *ListingImage) SetHexCode(v string) {
	o.HexCode.Set(&v)
}

// SetHexCodeNil sets the value for HexCode to be an explicit nil
func (o *ListingImage) SetHexCodeNil() {
	o.HexCode.Set(nil)
}

// UnsetHexCode ensures that no value is present for HexCode, not even an explicit nil
func (o *ListingImage) UnsetHexCode() {
	o.HexCode.Unset()
}

// GetRed returns the Red field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetRed() int64 {
	if o == nil || IsNil(o.Red.Get()) {
		var ret int64
		return ret
	}
	return *o.Red.Get()
}

// GetRedOk returns a tuple with the Red field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetRedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Red.Get(), o.Red.IsSet()
}

// HasRed returns a boolean if a field has been set.
func (o *ListingImage) HasRed() bool {
	if o != nil && o.Red.IsSet() {
		return true
	}

	return false
}

// SetRed gets a reference to the given NullableInt64 and assigns it to the Red field.
func (o *ListingImage) SetRed(v int64) {
	o.Red.Set(&v)
}

// SetRedNil sets the value for Red to be an explicit nil
func (o *ListingImage) SetRedNil() {
	o.Red.Set(nil)
}

// UnsetRed ensures that no value is present for Red, not even an explicit nil
func (o *ListingImage) UnsetRed() {
	o.Red.Unset()
}

// GetGreen returns the Green field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetGreen() int64 {
	if o == nil || IsNil(o.Green.Get()) {
		var ret int64
		return ret
	}
	return *o.Green.Get()
}

// GetGreenOk returns a tuple with the Green field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetGreenOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Green.Get(), o.Green.IsSet()
}

// HasGreen returns a boolean if a field has been set.
func (o *ListingImage) HasGreen() bool {
	if o != nil && o.Green.IsSet() {
		return true
	}

	return false
}

// SetGreen gets a reference to the given NullableInt64 and assigns it to the Green field.
func (o *ListingImage) SetGreen(v int64) {
	o.Green.Set(&v)
}

// SetGreenNil sets the value for Green to be an explicit nil
func (o *ListingImage) SetGreenNil() {
	o.Green.Set(nil)
}

// UnsetGreen ensures that no value is present for Green, not even an explicit nil
func (o *ListingImage) UnsetGreen() {
	o.Green.Unset()
}

// GetBlue returns the Blue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetBlue() int64 {
	if o == nil || IsNil(o.Blue.Get()) {
		var ret int64
		return ret
	}
	return *o.Blue.Get()
}

// GetBlueOk returns a tuple with the Blue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetBlueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blue.Get(), o.Blue.IsSet()
}

// HasBlue returns a boolean if a field has been set.
func (o *ListingImage) HasBlue() bool {
	if o != nil && o.Blue.IsSet() {
		return true
	}

	return false
}

// SetBlue gets a reference to the given NullableInt64 and assigns it to the Blue field.
func (o *ListingImage) SetBlue(v int64) {
	o.Blue.Set(&v)
}

// SetBlueNil sets the value for Blue to be an explicit nil
func (o *ListingImage) SetBlueNil() {
	o.Blue.Set(nil)
}

// UnsetBlue ensures that no value is present for Blue, not even an explicit nil
func (o *ListingImage) UnsetBlue() {
	o.Blue.Unset()
}

// GetHue returns the Hue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetHue() int64 {
	if o == nil || IsNil(o.Hue.Get()) {
		var ret int64
		return ret
	}
	return *o.Hue.Get()
}

// GetHueOk returns a tuple with the Hue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetHueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hue.Get(), o.Hue.IsSet()
}

// HasHue returns a boolean if a field has been set.
func (o *ListingImage) HasHue() bool {
	if o != nil && o.Hue.IsSet() {
		return true
	}

	return false
}

// SetHue gets a reference to the given NullableInt64 and assigns it to the Hue field.
func (o *ListingImage) SetHue(v int64) {
	o.Hue.Set(&v)
}

// SetHueNil sets the value for Hue to be an explicit nil
func (o *ListingImage) SetHueNil() {
	o.Hue.Set(nil)
}

// UnsetHue ensures that no value is present for Hue, not even an explicit nil
func (o *ListingImage) UnsetHue() {
	o.Hue.Unset()
}

// GetSaturation returns the Saturation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetSaturation() int64 {
	if o == nil || IsNil(o.Saturation.Get()) {
		var ret int64
		return ret
	}
	return *o.Saturation.Get()
}

// GetSaturationOk returns a tuple with the Saturation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetSaturationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saturation.Get(), o.Saturation.IsSet()
}

// HasSaturation returns a boolean if a field has been set.
func (o *ListingImage) HasSaturation() bool {
	if o != nil && o.Saturation.IsSet() {
		return true
	}

	return false
}

// SetSaturation gets a reference to the given NullableInt64 and assigns it to the Saturation field.
func (o *ListingImage) SetSaturation(v int64) {
	o.Saturation.Set(&v)
}

// SetSaturationNil sets the value for Saturation to be an explicit nil
func (o *ListingImage) SetSaturationNil() {
	o.Saturation.Set(nil)
}

// UnsetSaturation ensures that no value is present for Saturation, not even an explicit nil
func (o *ListingImage) UnsetSaturation() {
	o.Saturation.Unset()
}

// GetBrightness returns the Brightness field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetBrightness() int64 {
	if o == nil || IsNil(o.Brightness.Get()) {
		var ret int64
		return ret
	}
	return *o.Brightness.Get()
}

// GetBrightnessOk returns a tuple with the Brightness field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetBrightnessOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Brightness.Get(), o.Brightness.IsSet()
}

// HasBrightness returns a boolean if a field has been set.
func (o *ListingImage) HasBrightness() bool {
	if o != nil && o.Brightness.IsSet() {
		return true
	}

	return false
}

// SetBrightness gets a reference to the given NullableInt64 and assigns it to the Brightness field.
func (o *ListingImage) SetBrightness(v int64) {
	o.Brightness.Set(&v)
}

// SetBrightnessNil sets the value for Brightness to be an explicit nil
func (o *ListingImage) SetBrightnessNil() {
	o.Brightness.Set(nil)
}

// UnsetBrightness ensures that no value is present for Brightness, not even an explicit nil
func (o *ListingImage) UnsetBrightness() {
	o.Brightness.Unset()
}

// GetIsBlackAndWhite returns the IsBlackAndWhite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetIsBlackAndWhite() bool {
	if o == nil || IsNil(o.IsBlackAndWhite.Get()) {
		var ret bool
		return ret
	}
	return *o.IsBlackAndWhite.Get()
}

// GetIsBlackAndWhiteOk returns a tuple with the IsBlackAndWhite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetIsBlackAndWhiteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsBlackAndWhite.Get(), o.IsBlackAndWhite.IsSet()
}

// HasIsBlackAndWhite returns a boolean if a field has been set.
func (o *ListingImage) HasIsBlackAndWhite() bool {
	if o != nil && o.IsBlackAndWhite.IsSet() {
		return true
	}

	return false
}

// SetIsBlackAndWhite gets a reference to the given NullableBool and assigns it to the IsBlackAndWhite field.
func (o *ListingImage) SetIsBlackAndWhite(v bool) {
	o.IsBlackAndWhite.Set(&v)
}

// SetIsBlackAndWhiteNil sets the value for IsBlackAndWhite to be an explicit nil
func (o *ListingImage) SetIsBlackAndWhiteNil() {
	o.IsBlackAndWhite.Set(nil)
}

// UnsetIsBlackAndWhite ensures that no value is present for IsBlackAndWhite, not even an explicit nil
func (o *ListingImage) UnsetIsBlackAndWhite() {
	o.IsBlackAndWhite.Unset()
}

// GetCreationTsz returns the CreationTsz field value if set, zero value otherwise.
func (o *ListingImage) GetCreationTsz() int64 {
	if o == nil || IsNil(o.CreationTsz) {
		var ret int64
		return ret
	}
	return *o.CreationTsz
}

// GetCreationTszOk returns a tuple with the CreationTsz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingImage) GetCreationTszOk() (*int64, bool) {
	if o == nil || IsNil(o.CreationTsz) {
		return nil, false
	}
	return o.CreationTsz, true
}

// HasCreationTsz returns a boolean if a field has been set.
func (o *ListingImage) HasCreationTsz() bool {
	if o != nil && !IsNil(o.CreationTsz) {
		return true
	}

	return false
}

// SetCreationTsz gets a reference to the given int64 and assigns it to the CreationTsz field.
func (o *ListingImage) SetCreationTsz(v int64) {
	o.CreationTsz = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *ListingImage) GetCreatedTimestamp() int64 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingImage) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *ListingImage) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int64 and assigns it to the CreatedTimestamp field.
func (o *ListingImage) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *ListingImage) GetRank() int64 {
	if o == nil || IsNil(o.Rank) {
		var ret int64
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingImage) GetRankOk() (*int64, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *ListingImage) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int64 and assigns it to the Rank field.
func (o *ListingImage) SetRank(v int64) {
	o.Rank = &v
}

// GetUrl75x75 returns the Url75x75 field value if set, zero value otherwise.
func (o *ListingImage) GetUrl75x75() string {
	if o == nil || IsNil(o.Url75x75) {
		var ret string
		return ret
	}
	return *o.Url75x75
}

// GetUrl75x75Ok returns a tuple with the Url75x75 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingImage) GetUrl75x75Ok() (*string, bool) {
	if o == nil || IsNil(o.Url75x75) {
		return nil, false
	}
	return o.Url75x75, true
}

// HasUrl75x75 returns a boolean if a field has been set.
func (o *ListingImage) HasUrl75x75() bool {
	if o != nil && !IsNil(o.Url75x75) {
		return true
	}

	return false
}

// SetUrl75x75 gets a reference to the given string and assigns it to the Url75x75 field.
func (o *ListingImage) SetUrl75x75(v string) {
	o.Url75x75 = &v
}

// GetUrl170x135 returns the Url170x135 field value if set, zero value otherwise.
func (o *ListingImage) GetUrl170x135() string {
	if o == nil || IsNil(o.Url170x135) {
		var ret string
		return ret
	}
	return *o.Url170x135
}

// GetUrl170x135Ok returns a tuple with the Url170x135 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingImage) GetUrl170x135Ok() (*string, bool) {
	if o == nil || IsNil(o.Url170x135) {
		return nil, false
	}
	return o.Url170x135, true
}

// HasUrl170x135 returns a boolean if a field has been set.
func (o *ListingImage) HasUrl170x135() bool {
	if o != nil && !IsNil(o.Url170x135) {
		return true
	}

	return false
}

// SetUrl170x135 gets a reference to the given string and assigns it to the Url170x135 field.
func (o *ListingImage) SetUrl170x135(v string) {
	o.Url170x135 = &v
}

// GetUrl570xN returns the Url570xN field value if set, zero value otherwise.
func (o *ListingImage) GetUrl570xN() string {
	if o == nil || IsNil(o.Url570xN) {
		var ret string
		return ret
	}
	return *o.Url570xN
}

// GetUrl570xNOk returns a tuple with the Url570xN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingImage) GetUrl570xNOk() (*string, bool) {
	if o == nil || IsNil(o.Url570xN) {
		return nil, false
	}
	return o.Url570xN, true
}

// HasUrl570xN returns a boolean if a field has been set.
func (o *ListingImage) HasUrl570xN() bool {
	if o != nil && !IsNil(o.Url570xN) {
		return true
	}

	return false
}

// SetUrl570xN gets a reference to the given string and assigns it to the Url570xN field.
func (o *ListingImage) SetUrl570xN(v string) {
	o.Url570xN = &v
}

// GetUrlFullxfull returns the UrlFullxfull field value if set, zero value otherwise.
func (o *ListingImage) GetUrlFullxfull() string {
	if o == nil || IsNil(o.UrlFullxfull) {
		var ret string
		return ret
	}
	return *o.UrlFullxfull
}

// GetUrlFullxfullOk returns a tuple with the UrlFullxfull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingImage) GetUrlFullxfullOk() (*string, bool) {
	if o == nil || IsNil(o.UrlFullxfull) {
		return nil, false
	}
	return o.UrlFullxfull, true
}

// HasUrlFullxfull returns a boolean if a field has been set.
func (o *ListingImage) HasUrlFullxfull() bool {
	if o != nil && !IsNil(o.UrlFullxfull) {
		return true
	}

	return false
}

// SetUrlFullxfull gets a reference to the given string and assigns it to the UrlFullxfull field.
func (o *ListingImage) SetUrlFullxfull(v string) {
	o.UrlFullxfull = &v
}

// GetFullHeight returns the FullHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetFullHeight() int64 {
	if o == nil || IsNil(o.FullHeight.Get()) {
		var ret int64
		return ret
	}
	return *o.FullHeight.Get()
}

// GetFullHeightOk returns a tuple with the FullHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetFullHeightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullHeight.Get(), o.FullHeight.IsSet()
}

// HasFullHeight returns a boolean if a field has been set.
func (o *ListingImage) HasFullHeight() bool {
	if o != nil && o.FullHeight.IsSet() {
		return true
	}

	return false
}

// SetFullHeight gets a reference to the given NullableInt64 and assigns it to the FullHeight field.
func (o *ListingImage) SetFullHeight(v int64) {
	o.FullHeight.Set(&v)
}

// SetFullHeightNil sets the value for FullHeight to be an explicit nil
func (o *ListingImage) SetFullHeightNil() {
	o.FullHeight.Set(nil)
}

// UnsetFullHeight ensures that no value is present for FullHeight, not even an explicit nil
func (o *ListingImage) UnsetFullHeight() {
	o.FullHeight.Unset()
}

// GetFullWidth returns the FullWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetFullWidth() int64 {
	if o == nil || IsNil(o.FullWidth.Get()) {
		var ret int64
		return ret
	}
	return *o.FullWidth.Get()
}

// GetFullWidthOk returns a tuple with the FullWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetFullWidthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullWidth.Get(), o.FullWidth.IsSet()
}

// HasFullWidth returns a boolean if a field has been set.
func (o *ListingImage) HasFullWidth() bool {
	if o != nil && o.FullWidth.IsSet() {
		return true
	}

	return false
}

// SetFullWidth gets a reference to the given NullableInt64 and assigns it to the FullWidth field.
func (o *ListingImage) SetFullWidth(v int64) {
	o.FullWidth.Set(&v)
}

// SetFullWidthNil sets the value for FullWidth to be an explicit nil
func (o *ListingImage) SetFullWidthNil() {
	o.FullWidth.Set(nil)
}

// UnsetFullWidth ensures that no value is present for FullWidth, not even an explicit nil
func (o *ListingImage) UnsetFullWidth() {
	o.FullWidth.Unset()
}

// GetAltText returns the AltText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingImage) GetAltText() string {
	if o == nil || IsNil(o.AltText.Get()) {
		var ret string
		return ret
	}
	return *o.AltText.Get()
}

// GetAltTextOk returns a tuple with the AltText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingImage) GetAltTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AltText.Get(), o.AltText.IsSet()
}

// HasAltText returns a boolean if a field has been set.
func (o *ListingImage) HasAltText() bool {
	if o != nil && o.AltText.IsSet() {
		return true
	}

	return false
}

// SetAltText gets a reference to the given NullableString and assigns it to the AltText field.
func (o *ListingImage) SetAltText(v string) {
	o.AltText.Set(&v)
}

// SetAltTextNil sets the value for AltText to be an explicit nil
func (o *ListingImage) SetAltTextNil() {
	o.AltText.Set(nil)
}

// UnsetAltText ensures that no value is present for AltText, not even an explicit nil
func (o *ListingImage) UnsetAltText() {
	o.AltText.Unset()
}

func (o ListingImage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListingImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListingId) {
		toSerialize["listing_id"] = o.ListingId
	}
	if !IsNil(o.ListingImageId) {
		toSerialize["listing_image_id"] = o.ListingImageId
	}
	if o.HexCode.IsSet() {
		toSerialize["hex_code"] = o.HexCode.Get()
	}
	if o.Red.IsSet() {
		toSerialize["red"] = o.Red.Get()
	}
	if o.Green.IsSet() {
		toSerialize["green"] = o.Green.Get()
	}
	if o.Blue.IsSet() {
		toSerialize["blue"] = o.Blue.Get()
	}
	if o.Hue.IsSet() {
		toSerialize["hue"] = o.Hue.Get()
	}
	if o.Saturation.IsSet() {
		toSerialize["saturation"] = o.Saturation.Get()
	}
	if o.Brightness.IsSet() {
		toSerialize["brightness"] = o.Brightness.Get()
	}
	if o.IsBlackAndWhite.IsSet() {
		toSerialize["is_black_and_white"] = o.IsBlackAndWhite.Get()
	}
	if !IsNil(o.CreationTsz) {
		toSerialize["creation_tsz"] = o.CreationTsz
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	if !IsNil(o.Url75x75) {
		toSerialize["url_75x75"] = o.Url75x75
	}
	if !IsNil(o.Url170x135) {
		toSerialize["url_170x135"] = o.Url170x135
	}
	if !IsNil(o.Url570xN) {
		toSerialize["url_570xN"] = o.Url570xN
	}
	if !IsNil(o.UrlFullxfull) {
		toSerialize["url_fullxfull"] = o.UrlFullxfull
	}
	if o.FullHeight.IsSet() {
		toSerialize["full_height"] = o.FullHeight.Get()
	}
	if o.FullWidth.IsSet() {
		toSerialize["full_width"] = o.FullWidth.Get()
	}
	if o.AltText.IsSet() {
		toSerialize["alt_text"] = o.AltText.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableListingImage struct {
	value *ListingImage
	isSet bool
}

func (v NullableListingImage) Get() *ListingImage {
	return v.value
}

func (v *NullableListingImage) Set(val *ListingImage) {
	v.value = val
	v.isSet = true
}

func (v NullableListingImage) IsSet() bool {
	return v.isSet
}

func (v *NullableListingImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingImage(val *ListingImage) *NullableListingImage {
	return &NullableListingImage{value: val, isSet: true}
}

func (v NullableListingImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
