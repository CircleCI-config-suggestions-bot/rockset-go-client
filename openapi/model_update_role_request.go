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

// UpdateRoleRequest struct for UpdateRoleRequest
type UpdateRoleRequest struct {
	// Description for the role.
	Description *string `json:"description,omitempty"`
	// List of privileges that will be associated with the role.
	Privileges []Privilege `json:"privileges,omitempty"`
}

// NewUpdateRoleRequest instantiates a new UpdateRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleRequest() *UpdateRoleRequest {
	this := UpdateRoleRequest{}
	return &this
}

// NewUpdateRoleRequestWithDefaults instantiates a new UpdateRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleRequestWithDefaults() *UpdateRoleRequest {
	this := UpdateRoleRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateRoleRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateRoleRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateRoleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *UpdateRoleRequest) GetPrivileges() []Privilege {
	if o == nil || o.Privileges == nil {
		var ret []Privilege
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequest) GetPrivilegesOk() ([]Privilege, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *UpdateRoleRequest) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []Privilege and assigns it to the Privileges field.
func (o *UpdateRoleRequest) SetPrivileges(v []Privilege) {
	o.Privileges = v
}

func (o UpdateRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRoleRequest struct {
	value *UpdateRoleRequest
	isSet bool
}

func (v NullableUpdateRoleRequest) Get() *UpdateRoleRequest {
	return v.value
}

func (v *NullableUpdateRoleRequest) Set(val *UpdateRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleRequest(val *UpdateRoleRequest) *NullableUpdateRoleRequest {
	return &NullableUpdateRoleRequest{value: val, isSet: true}
}

func (v NullableUpdateRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


