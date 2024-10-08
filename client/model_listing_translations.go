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

// checks if the ListingTranslations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingTranslations{}

// ListingTranslations Container for all current supported translations of a listing. Note that Etsy periodically adds/removes languages, so this list may change in the future.
type ListingTranslations struct {
	De                   NullableListingTranslation `json:"de,omitempty"`
	EnGB                 NullableListingTranslation `json:"en-GB,omitempty"`
	EnIN                 NullableListingTranslation `json:"en-IN,omitempty"`
	EnUS                 NullableListingTranslation `json:"en-US,omitempty"`
	Es                   NullableListingTranslation `json:"es,omitempty"`
	Fr                   NullableListingTranslation `json:"fr,omitempty"`
	It                   NullableListingTranslation `json:"it,omitempty"`
	Ja                   NullableListingTranslation `json:"ja,omitempty"`
	Nl                   NullableListingTranslation `json:"nl,omitempty"`
	Pl                   NullableListingTranslation `json:"pl,omitempty"`
	Pt                   NullableListingTranslation `json:"pt,omitempty"`
	Ru                   NullableListingTranslation `json:"ru,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListingTranslations ListingTranslations

// NewListingTranslations instantiates a new ListingTranslations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingTranslations() *ListingTranslations {
	this := ListingTranslations{}
	return &this
}

// NewListingTranslationsWithDefaults instantiates a new ListingTranslations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingTranslationsWithDefaults() *ListingTranslations {
	this := ListingTranslations{}
	return &this
}

// GetDe returns the De field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetDe() ListingTranslation {
	if o == nil || IsNil(o.De.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.De.Get()
}

// GetDeOk returns a tuple with the De field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetDeOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.De.Get(), o.De.IsSet()
}

// HasDe returns a boolean if a field has been set.
func (o *ListingTranslations) HasDe() bool {
	if o != nil && o.De.IsSet() {
		return true
	}

	return false
}

// SetDe gets a reference to the given NullableListingTranslation and assigns it to the De field.
func (o *ListingTranslations) SetDe(v ListingTranslation) {
	o.De.Set(&v)
}

// SetDeNil sets the value for De to be an explicit nil
func (o *ListingTranslations) SetDeNil() {
	o.De.Set(nil)
}

// UnsetDe ensures that no value is present for De, not even an explicit nil
func (o *ListingTranslations) UnsetDe() {
	o.De.Unset()
}

// GetEnGB returns the EnGB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetEnGB() ListingTranslation {
	if o == nil || IsNil(o.EnGB.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.EnGB.Get()
}

// GetEnGBOk returns a tuple with the EnGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetEnGBOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnGB.Get(), o.EnGB.IsSet()
}

// HasEnGB returns a boolean if a field has been set.
func (o *ListingTranslations) HasEnGB() bool {
	if o != nil && o.EnGB.IsSet() {
		return true
	}

	return false
}

// SetEnGB gets a reference to the given NullableListingTranslation and assigns it to the EnGB field.
func (o *ListingTranslations) SetEnGB(v ListingTranslation) {
	o.EnGB.Set(&v)
}

// SetEnGBNil sets the value for EnGB to be an explicit nil
func (o *ListingTranslations) SetEnGBNil() {
	o.EnGB.Set(nil)
}

// UnsetEnGB ensures that no value is present for EnGB, not even an explicit nil
func (o *ListingTranslations) UnsetEnGB() {
	o.EnGB.Unset()
}

// GetEnIN returns the EnIN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetEnIN() ListingTranslation {
	if o == nil || IsNil(o.EnIN.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.EnIN.Get()
}

// GetEnINOk returns a tuple with the EnIN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetEnINOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnIN.Get(), o.EnIN.IsSet()
}

// HasEnIN returns a boolean if a field has been set.
func (o *ListingTranslations) HasEnIN() bool {
	if o != nil && o.EnIN.IsSet() {
		return true
	}

	return false
}

// SetEnIN gets a reference to the given NullableListingTranslation and assigns it to the EnIN field.
func (o *ListingTranslations) SetEnIN(v ListingTranslation) {
	o.EnIN.Set(&v)
}

// SetEnINNil sets the value for EnIN to be an explicit nil
func (o *ListingTranslations) SetEnINNil() {
	o.EnIN.Set(nil)
}

// UnsetEnIN ensures that no value is present for EnIN, not even an explicit nil
func (o *ListingTranslations) UnsetEnIN() {
	o.EnIN.Unset()
}

// GetEnUS returns the EnUS field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetEnUS() ListingTranslation {
	if o == nil || IsNil(o.EnUS.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.EnUS.Get()
}

// GetEnUSOk returns a tuple with the EnUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetEnUSOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnUS.Get(), o.EnUS.IsSet()
}

// HasEnUS returns a boolean if a field has been set.
func (o *ListingTranslations) HasEnUS() bool {
	if o != nil && o.EnUS.IsSet() {
		return true
	}

	return false
}

// SetEnUS gets a reference to the given NullableListingTranslation and assigns it to the EnUS field.
func (o *ListingTranslations) SetEnUS(v ListingTranslation) {
	o.EnUS.Set(&v)
}

// SetEnUSNil sets the value for EnUS to be an explicit nil
func (o *ListingTranslations) SetEnUSNil() {
	o.EnUS.Set(nil)
}

// UnsetEnUS ensures that no value is present for EnUS, not even an explicit nil
func (o *ListingTranslations) UnsetEnUS() {
	o.EnUS.Unset()
}

// GetEs returns the Es field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetEs() ListingTranslation {
	if o == nil || IsNil(o.Es.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.Es.Get()
}

// GetEsOk returns a tuple with the Es field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetEsOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Es.Get(), o.Es.IsSet()
}

// HasEs returns a boolean if a field has been set.
func (o *ListingTranslations) HasEs() bool {
	if o != nil && o.Es.IsSet() {
		return true
	}

	return false
}

// SetEs gets a reference to the given NullableListingTranslation and assigns it to the Es field.
func (o *ListingTranslations) SetEs(v ListingTranslation) {
	o.Es.Set(&v)
}

// SetEsNil sets the value for Es to be an explicit nil
func (o *ListingTranslations) SetEsNil() {
	o.Es.Set(nil)
}

// UnsetEs ensures that no value is present for Es, not even an explicit nil
func (o *ListingTranslations) UnsetEs() {
	o.Es.Unset()
}

// GetFr returns the Fr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetFr() ListingTranslation {
	if o == nil || IsNil(o.Fr.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.Fr.Get()
}

// GetFrOk returns a tuple with the Fr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetFrOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fr.Get(), o.Fr.IsSet()
}

// HasFr returns a boolean if a field has been set.
func (o *ListingTranslations) HasFr() bool {
	if o != nil && o.Fr.IsSet() {
		return true
	}

	return false
}

// SetFr gets a reference to the given NullableListingTranslation and assigns it to the Fr field.
func (o *ListingTranslations) SetFr(v ListingTranslation) {
	o.Fr.Set(&v)
}

// SetFrNil sets the value for Fr to be an explicit nil
func (o *ListingTranslations) SetFrNil() {
	o.Fr.Set(nil)
}

// UnsetFr ensures that no value is present for Fr, not even an explicit nil
func (o *ListingTranslations) UnsetFr() {
	o.Fr.Unset()
}

// GetIt returns the It field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetIt() ListingTranslation {
	if o == nil || IsNil(o.It.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.It.Get()
}

// GetItOk returns a tuple with the It field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetItOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.It.Get(), o.It.IsSet()
}

// HasIt returns a boolean if a field has been set.
func (o *ListingTranslations) HasIt() bool {
	if o != nil && o.It.IsSet() {
		return true
	}

	return false
}

// SetIt gets a reference to the given NullableListingTranslation and assigns it to the It field.
func (o *ListingTranslations) SetIt(v ListingTranslation) {
	o.It.Set(&v)
}

// SetItNil sets the value for It to be an explicit nil
func (o *ListingTranslations) SetItNil() {
	o.It.Set(nil)
}

// UnsetIt ensures that no value is present for It, not even an explicit nil
func (o *ListingTranslations) UnsetIt() {
	o.It.Unset()
}

// GetJa returns the Ja field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetJa() ListingTranslation {
	if o == nil || IsNil(o.Ja.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.Ja.Get()
}

// GetJaOk returns a tuple with the Ja field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetJaOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ja.Get(), o.Ja.IsSet()
}

// HasJa returns a boolean if a field has been set.
func (o *ListingTranslations) HasJa() bool {
	if o != nil && o.Ja.IsSet() {
		return true
	}

	return false
}

// SetJa gets a reference to the given NullableListingTranslation and assigns it to the Ja field.
func (o *ListingTranslations) SetJa(v ListingTranslation) {
	o.Ja.Set(&v)
}

// SetJaNil sets the value for Ja to be an explicit nil
func (o *ListingTranslations) SetJaNil() {
	o.Ja.Set(nil)
}

// UnsetJa ensures that no value is present for Ja, not even an explicit nil
func (o *ListingTranslations) UnsetJa() {
	o.Ja.Unset()
}

// GetNl returns the Nl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetNl() ListingTranslation {
	if o == nil || IsNil(o.Nl.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.Nl.Get()
}

// GetNlOk returns a tuple with the Nl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetNlOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nl.Get(), o.Nl.IsSet()
}

// HasNl returns a boolean if a field has been set.
func (o *ListingTranslations) HasNl() bool {
	if o != nil && o.Nl.IsSet() {
		return true
	}

	return false
}

// SetNl gets a reference to the given NullableListingTranslation and assigns it to the Nl field.
func (o *ListingTranslations) SetNl(v ListingTranslation) {
	o.Nl.Set(&v)
}

// SetNlNil sets the value for Nl to be an explicit nil
func (o *ListingTranslations) SetNlNil() {
	o.Nl.Set(nil)
}

// UnsetNl ensures that no value is present for Nl, not even an explicit nil
func (o *ListingTranslations) UnsetNl() {
	o.Nl.Unset()
}

// GetPl returns the Pl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetPl() ListingTranslation {
	if o == nil || IsNil(o.Pl.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.Pl.Get()
}

// GetPlOk returns a tuple with the Pl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetPlOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pl.Get(), o.Pl.IsSet()
}

// HasPl returns a boolean if a field has been set.
func (o *ListingTranslations) HasPl() bool {
	if o != nil && o.Pl.IsSet() {
		return true
	}

	return false
}

// SetPl gets a reference to the given NullableListingTranslation and assigns it to the Pl field.
func (o *ListingTranslations) SetPl(v ListingTranslation) {
	o.Pl.Set(&v)
}

// SetPlNil sets the value for Pl to be an explicit nil
func (o *ListingTranslations) SetPlNil() {
	o.Pl.Set(nil)
}

// UnsetPl ensures that no value is present for Pl, not even an explicit nil
func (o *ListingTranslations) UnsetPl() {
	o.Pl.Unset()
}

// GetPt returns the Pt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetPt() ListingTranslation {
	if o == nil || IsNil(o.Pt.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.Pt.Get()
}

// GetPtOk returns a tuple with the Pt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetPtOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pt.Get(), o.Pt.IsSet()
}

// HasPt returns a boolean if a field has been set.
func (o *ListingTranslations) HasPt() bool {
	if o != nil && o.Pt.IsSet() {
		return true
	}

	return false
}

// SetPt gets a reference to the given NullableListingTranslation and assigns it to the Pt field.
func (o *ListingTranslations) SetPt(v ListingTranslation) {
	o.Pt.Set(&v)
}

// SetPtNil sets the value for Pt to be an explicit nil
func (o *ListingTranslations) SetPtNil() {
	o.Pt.Set(nil)
}

// UnsetPt ensures that no value is present for Pt, not even an explicit nil
func (o *ListingTranslations) UnsetPt() {
	o.Pt.Unset()
}

// GetRu returns the Ru field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingTranslations) GetRu() ListingTranslation {
	if o == nil || IsNil(o.Ru.Get()) {
		var ret ListingTranslation
		return ret
	}
	return *o.Ru.Get()
}

// GetRuOk returns a tuple with the Ru field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingTranslations) GetRuOk() (*ListingTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ru.Get(), o.Ru.IsSet()
}

// HasRu returns a boolean if a field has been set.
func (o *ListingTranslations) HasRu() bool {
	if o != nil && o.Ru.IsSet() {
		return true
	}

	return false
}

// SetRu gets a reference to the given NullableListingTranslation and assigns it to the Ru field.
func (o *ListingTranslations) SetRu(v ListingTranslation) {
	o.Ru.Set(&v)
}

// SetRuNil sets the value for Ru to be an explicit nil
func (o *ListingTranslations) SetRuNil() {
	o.Ru.Set(nil)
}

// UnsetRu ensures that no value is present for Ru, not even an explicit nil
func (o *ListingTranslations) UnsetRu() {
	o.Ru.Unset()
}

func (o ListingTranslations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListingTranslations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.De.IsSet() {
		toSerialize["de"] = o.De.Get()
	}
	if o.EnGB.IsSet() {
		toSerialize["en-GB"] = o.EnGB.Get()
	}
	if o.EnIN.IsSet() {
		toSerialize["en-IN"] = o.EnIN.Get()
	}
	if o.EnUS.IsSet() {
		toSerialize["en-US"] = o.EnUS.Get()
	}
	if o.Es.IsSet() {
		toSerialize["es"] = o.Es.Get()
	}
	if o.Fr.IsSet() {
		toSerialize["fr"] = o.Fr.Get()
	}
	if o.It.IsSet() {
		toSerialize["it"] = o.It.Get()
	}
	if o.Ja.IsSet() {
		toSerialize["ja"] = o.Ja.Get()
	}
	if o.Nl.IsSet() {
		toSerialize["nl"] = o.Nl.Get()
	}
	if o.Pl.IsSet() {
		toSerialize["pl"] = o.Pl.Get()
	}
	if o.Pt.IsSet() {
		toSerialize["pt"] = o.Pt.Get()
	}
	if o.Ru.IsSet() {
		toSerialize["ru"] = o.Ru.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableListingTranslations struct {
	value *ListingTranslations
	isSet bool
}

func (v NullableListingTranslations) Get() *ListingTranslations {
	return v.value
}

func (v *NullableListingTranslations) Set(val *ListingTranslations) {
	v.value = val
	v.isSet = true
}

func (v NullableListingTranslations) IsSet() bool {
	return v.isSet
}

func (v *NullableListingTranslations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingTranslations(val *ListingTranslations) *NullableListingTranslations {
	return &NullableListingTranslations{value: val, isSet: true}
}

func (v NullableListingTranslations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingTranslations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
