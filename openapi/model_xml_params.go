/*
 * REST API
 *
 * Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// XmlParams struct for XmlParams
type XmlParams struct {
	// tag until which xml is ignored
	RootTag *string `json:"root_tag,omitempty"`
	// encoding in which data source is encoded
	Encoding *string `json:"encoding,omitempty"`
	// tags with which documents are identified
	DocTag *string `json:"doc_tag,omitempty"`
	// tag used for the value when there are attributes in the element having no child
	ValueTag *string `json:"value_tag,omitempty"`
	// tag to differentiate between attributes and elements
	AttributePrefix *string `json:"attribute_prefix,omitempty"`
}

// NewXmlParams instantiates a new XmlParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlParams() *XmlParams {
	this := XmlParams{}
	return &this
}

// NewXmlParamsWithDefaults instantiates a new XmlParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlParamsWithDefaults() *XmlParams {
	this := XmlParams{}
	return &this
}

// GetRootTag returns the RootTag field value if set, zero value otherwise.
func (o *XmlParams) GetRootTag() string {
	if o == nil || o.RootTag == nil {
		var ret string
		return ret
	}
	return *o.RootTag
}

// GetRootTagOk returns a tuple with the RootTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlParams) GetRootTagOk() (*string, bool) {
	if o == nil || o.RootTag == nil {
		return nil, false
	}
	return o.RootTag, true
}

// HasRootTag returns a boolean if a field has been set.
func (o *XmlParams) HasRootTag() bool {
	if o != nil && o.RootTag != nil {
		return true
	}

	return false
}

// SetRootTag gets a reference to the given string and assigns it to the RootTag field.
func (o *XmlParams) SetRootTag(v string) {
	o.RootTag = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *XmlParams) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlParams) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *XmlParams) HasEncoding() bool {
	if o != nil && o.Encoding != nil {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *XmlParams) SetEncoding(v string) {
	o.Encoding = &v
}

// GetDocTag returns the DocTag field value if set, zero value otherwise.
func (o *XmlParams) GetDocTag() string {
	if o == nil || o.DocTag == nil {
		var ret string
		return ret
	}
	return *o.DocTag
}

// GetDocTagOk returns a tuple with the DocTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlParams) GetDocTagOk() (*string, bool) {
	if o == nil || o.DocTag == nil {
		return nil, false
	}
	return o.DocTag, true
}

// HasDocTag returns a boolean if a field has been set.
func (o *XmlParams) HasDocTag() bool {
	if o != nil && o.DocTag != nil {
		return true
	}

	return false
}

// SetDocTag gets a reference to the given string and assigns it to the DocTag field.
func (o *XmlParams) SetDocTag(v string) {
	o.DocTag = &v
}

// GetValueTag returns the ValueTag field value if set, zero value otherwise.
func (o *XmlParams) GetValueTag() string {
	if o == nil || o.ValueTag == nil {
		var ret string
		return ret
	}
	return *o.ValueTag
}

// GetValueTagOk returns a tuple with the ValueTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlParams) GetValueTagOk() (*string, bool) {
	if o == nil || o.ValueTag == nil {
		return nil, false
	}
	return o.ValueTag, true
}

// HasValueTag returns a boolean if a field has been set.
func (o *XmlParams) HasValueTag() bool {
	if o != nil && o.ValueTag != nil {
		return true
	}

	return false
}

// SetValueTag gets a reference to the given string and assigns it to the ValueTag field.
func (o *XmlParams) SetValueTag(v string) {
	o.ValueTag = &v
}

// GetAttributePrefix returns the AttributePrefix field value if set, zero value otherwise.
func (o *XmlParams) GetAttributePrefix() string {
	if o == nil || o.AttributePrefix == nil {
		var ret string
		return ret
	}
	return *o.AttributePrefix
}

// GetAttributePrefixOk returns a tuple with the AttributePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlParams) GetAttributePrefixOk() (*string, bool) {
	if o == nil || o.AttributePrefix == nil {
		return nil, false
	}
	return o.AttributePrefix, true
}

// HasAttributePrefix returns a boolean if a field has been set.
func (o *XmlParams) HasAttributePrefix() bool {
	if o != nil && o.AttributePrefix != nil {
		return true
	}

	return false
}

// SetAttributePrefix gets a reference to the given string and assigns it to the AttributePrefix field.
func (o *XmlParams) SetAttributePrefix(v string) {
	o.AttributePrefix = &v
}

func (o XmlParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RootTag != nil {
		toSerialize["root_tag"] = o.RootTag
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.DocTag != nil {
		toSerialize["doc_tag"] = o.DocTag
	}
	if o.ValueTag != nil {
		toSerialize["value_tag"] = o.ValueTag
	}
	if o.AttributePrefix != nil {
		toSerialize["attribute_prefix"] = o.AttributePrefix
	}
	return json.Marshal(toSerialize)
}

type NullableXmlParams struct {
	value *XmlParams
	isSet bool
}

func (v NullableXmlParams) Get() *XmlParams {
	return v.value
}

func (v *NullableXmlParams) Set(val *XmlParams) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlParams) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlParams(val *XmlParams) *NullableXmlParams {
	return &NullableXmlParams{value: val, isSet: true}
}

func (v NullableXmlParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


