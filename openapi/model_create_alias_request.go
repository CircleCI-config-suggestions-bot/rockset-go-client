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

// CreateAliasRequest struct for CreateAliasRequest
type CreateAliasRequest struct {
	// List of fully qualified collection names referenced by alias.
	Collections []string `json:"collections"`
	// Optional description.
	Description *string `json:"description,omitempty"`
	// Alias name.
	Name string `json:"name"`
}

// NewCreateAliasRequest instantiates a new CreateAliasRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAliasRequest(collections []string, name string) *CreateAliasRequest {
	this := CreateAliasRequest{}
	this.Collections = collections
	this.Name = name
	return &this
}

// NewCreateAliasRequestWithDefaults instantiates a new CreateAliasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAliasRequestWithDefaults() *CreateAliasRequest {
	this := CreateAliasRequest{}
	return &this
}

// GetCollections returns the Collections field value
func (o *CreateAliasRequest) GetCollections() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetCollectionsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Collections, true
}

// SetCollections sets field value
func (o *CreateAliasRequest) SetCollections(v []string) {
	o.Collections = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAliasRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAliasRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAliasRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *CreateAliasRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAliasRequest) SetName(v string) {
	o.Name = v
}

func (o CreateAliasRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["collections"] = o.Collections
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAliasRequest struct {
	value *CreateAliasRequest
	isSet bool
}

func (v NullableCreateAliasRequest) Get() *CreateAliasRequest {
	return v.value
}

func (v *NullableCreateAliasRequest) Set(val *CreateAliasRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAliasRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAliasRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAliasRequest(val *CreateAliasRequest) *NullableCreateAliasRequest {
	return &NullableCreateAliasRequest{value: val, isSet: true}
}

func (v NullableCreateAliasRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAliasRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


