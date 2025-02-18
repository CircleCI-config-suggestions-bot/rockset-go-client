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

// GcpServiceAccount struct for GcpServiceAccount
type GcpServiceAccount struct {
	// Contents of JSON Service Account key file.
	ServiceAccountKeyFileJson string `json:"service_account_key_file_json"`
}

// NewGcpServiceAccount instantiates a new GcpServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpServiceAccount(serviceAccountKeyFileJson string) *GcpServiceAccount {
	this := GcpServiceAccount{}
	this.ServiceAccountKeyFileJson = serviceAccountKeyFileJson
	return &this
}

// NewGcpServiceAccountWithDefaults instantiates a new GcpServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpServiceAccountWithDefaults() *GcpServiceAccount {
	this := GcpServiceAccount{}
	return &this
}

// GetServiceAccountKeyFileJson returns the ServiceAccountKeyFileJson field value
func (o *GcpServiceAccount) GetServiceAccountKeyFileJson() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAccountKeyFileJson
}

// GetServiceAccountKeyFileJsonOk returns a tuple with the ServiceAccountKeyFileJson field value
// and a boolean to check if the value has been set.
func (o *GcpServiceAccount) GetServiceAccountKeyFileJsonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceAccountKeyFileJson, true
}

// SetServiceAccountKeyFileJson sets field value
func (o *GcpServiceAccount) SetServiceAccountKeyFileJson(v string) {
	o.ServiceAccountKeyFileJson = v
}

func (o GcpServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["service_account_key_file_json"] = o.ServiceAccountKeyFileJson
	}
	return json.Marshal(toSerialize)
}

type NullableGcpServiceAccount struct {
	value *GcpServiceAccount
	isSet bool
}

func (v NullableGcpServiceAccount) Get() *GcpServiceAccount {
	return v.value
}

func (v *NullableGcpServiceAccount) Set(val *GcpServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpServiceAccount(val *GcpServiceAccount) *NullableGcpServiceAccount {
	return &NullableGcpServiceAccount{value: val, isSet: true}
}

func (v NullableGcpServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


