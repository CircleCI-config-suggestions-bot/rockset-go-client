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

// PatchOperation struct for PatchOperation
type PatchOperation struct {
	// [JSON Pointer](https://datatracker.ietf.org/doc/html/rfc6901) referencing a location in the target document. Required for `COPY` and `MOVE` operations.
	From *string `json:"from,omitempty"`
	// [JSON Patch operation](https://datatracker.ietf.org/doc/html/rfc6902#page-4) to be performed in this patch. Case insensitive.
	Op string `json:"op"`
	// [JSON Pointer](https://datatracker.ietf.org/doc/html/rfc6901) referencing a location in the target document where the operation is performed
	Path string `json:"path"`
	// Value used in the patch operation. Required for `ADD`, `REPLACE`, `TEST`, and `INCREMENT` operations.
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewPatchOperation instantiates a new PatchOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchOperation(op string, path string) *PatchOperation {
	this := PatchOperation{}
	this.Op = op
	this.Path = path
	return &this
}

// NewPatchOperationWithDefaults instantiates a new PatchOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchOperationWithDefaults() *PatchOperation {
	this := PatchOperation{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PatchOperation) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOperation) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PatchOperation) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PatchOperation) SetFrom(v string) {
	o.From = &v
}

// GetOp returns the Op field value
func (o *PatchOperation) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchOperation) GetOpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchOperation) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchOperation) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchOperation) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOperation) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchOperation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *PatchOperation) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o PatchOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePatchOperation struct {
	value *PatchOperation
	isSet bool
}

func (v NullablePatchOperation) Get() *PatchOperation {
	return v.value
}

func (v *NullablePatchOperation) Set(val *PatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOperation(val *PatchOperation) *NullablePatchOperation {
	return &NullablePatchOperation{value: val, isSet: true}
}

func (v NullablePatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


