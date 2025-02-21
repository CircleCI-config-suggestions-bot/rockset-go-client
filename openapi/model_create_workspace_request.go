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

// CreateWorkspaceRequest struct for CreateWorkspaceRequest
type CreateWorkspaceRequest struct {
	// Longer explanation for the workspace.
	Description *string `json:"description,omitempty"`
	// Descriptive label and unique identifier.
	Name string `json:"name"`
}

// NewCreateWorkspaceRequest instantiates a new CreateWorkspaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceRequest(name string) *CreateWorkspaceRequest {
	this := CreateWorkspaceRequest{}
	this.Name = name
	return &this
}

// NewCreateWorkspaceRequestWithDefaults instantiates a new CreateWorkspaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceRequestWithDefaults() *CreateWorkspaceRequest {
	this := CreateWorkspaceRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateWorkspaceRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateWorkspaceRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateWorkspaceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *CreateWorkspaceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWorkspaceRequest) SetName(v string) {
	o.Name = v
}

func (o CreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWorkspaceRequest struct {
	value *CreateWorkspaceRequest
	isSet bool
}

func (v NullableCreateWorkspaceRequest) Get() *CreateWorkspaceRequest {
	return v.value
}

func (v *NullableCreateWorkspaceRequest) Set(val *CreateWorkspaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceRequest(val *CreateWorkspaceRequest) *NullableCreateWorkspaceRequest {
	return &NullableCreateWorkspaceRequest{value: val, isSet: true}
}

func (v NullableCreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


