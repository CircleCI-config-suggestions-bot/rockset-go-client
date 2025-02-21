/*
REST API

Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FieldMappingV2 struct for FieldMappingV2
type FieldMappingV2 struct {
	// A List of InputField for this mapping.
	InputFields []InputField `json:"input_fields,omitempty"`
	// A boolean that determines whether to drop all fields in this document. If set, input and output fields should not be set
	IsDropAllFields *bool `json:"is_drop_all_fields,omitempty"`
	// A user specified string that is a name for this mapping.
	Name *string `json:"name,omitempty"`
	OutputField *OutputField `json:"output_field,omitempty"`
}

// NewFieldMappingV2 instantiates a new FieldMappingV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldMappingV2() *FieldMappingV2 {
	this := FieldMappingV2{}
	return &this
}

// NewFieldMappingV2WithDefaults instantiates a new FieldMappingV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldMappingV2WithDefaults() *FieldMappingV2 {
	this := FieldMappingV2{}
	return &this
}

// GetInputFields returns the InputFields field value if set, zero value otherwise.
func (o *FieldMappingV2) GetInputFields() []InputField {
	if o == nil || o.InputFields == nil {
		var ret []InputField
		return ret
	}
	return o.InputFields
}

// GetInputFieldsOk returns a tuple with the InputFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingV2) GetInputFieldsOk() ([]InputField, bool) {
	if o == nil || o.InputFields == nil {
		return nil, false
	}
	return o.InputFields, true
}

// HasInputFields returns a boolean if a field has been set.
func (o *FieldMappingV2) HasInputFields() bool {
	if o != nil && o.InputFields != nil {
		return true
	}

	return false
}

// SetInputFields gets a reference to the given []InputField and assigns it to the InputFields field.
func (o *FieldMappingV2) SetInputFields(v []InputField) {
	o.InputFields = v
}

// GetIsDropAllFields returns the IsDropAllFields field value if set, zero value otherwise.
func (o *FieldMappingV2) GetIsDropAllFields() bool {
	if o == nil || o.IsDropAllFields == nil {
		var ret bool
		return ret
	}
	return *o.IsDropAllFields
}

// GetIsDropAllFieldsOk returns a tuple with the IsDropAllFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingV2) GetIsDropAllFieldsOk() (*bool, bool) {
	if o == nil || o.IsDropAllFields == nil {
		return nil, false
	}
	return o.IsDropAllFields, true
}

// HasIsDropAllFields returns a boolean if a field has been set.
func (o *FieldMappingV2) HasIsDropAllFields() bool {
	if o != nil && o.IsDropAllFields != nil {
		return true
	}

	return false
}

// SetIsDropAllFields gets a reference to the given bool and assigns it to the IsDropAllFields field.
func (o *FieldMappingV2) SetIsDropAllFields(v bool) {
	o.IsDropAllFields = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FieldMappingV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FieldMappingV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FieldMappingV2) SetName(v string) {
	o.Name = &v
}

// GetOutputField returns the OutputField field value if set, zero value otherwise.
func (o *FieldMappingV2) GetOutputField() OutputField {
	if o == nil || o.OutputField == nil {
		var ret OutputField
		return ret
	}
	return *o.OutputField
}

// GetOutputFieldOk returns a tuple with the OutputField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingV2) GetOutputFieldOk() (*OutputField, bool) {
	if o == nil || o.OutputField == nil {
		return nil, false
	}
	return o.OutputField, true
}

// HasOutputField returns a boolean if a field has been set.
func (o *FieldMappingV2) HasOutputField() bool {
	if o != nil && o.OutputField != nil {
		return true
	}

	return false
}

// SetOutputField gets a reference to the given OutputField and assigns it to the OutputField field.
func (o *FieldMappingV2) SetOutputField(v OutputField) {
	o.OutputField = &v
}

func (o FieldMappingV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InputFields != nil {
		toSerialize["input_fields"] = o.InputFields
	}
	if o.IsDropAllFields != nil {
		toSerialize["is_drop_all_fields"] = o.IsDropAllFields
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OutputField != nil {
		toSerialize["output_field"] = o.OutputField
	}
	return json.Marshal(toSerialize)
}

type NullableFieldMappingV2 struct {
	value *FieldMappingV2
	isSet bool
}

func (v NullableFieldMappingV2) Get() *FieldMappingV2 {
	return v.value
}

func (v *NullableFieldMappingV2) Set(val *FieldMappingV2) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldMappingV2) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldMappingV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldMappingV2(val *FieldMappingV2) *NullableFieldMappingV2 {
	return &NullableFieldMappingV2{value: val, isSet: true}
}

func (v NullableFieldMappingV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldMappingV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


