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

// KinesisIntegration struct for KinesisIntegration
type KinesisIntegration struct {
	AwsAccessKey *AwsAccessKey `json:"aws_access_key,omitempty"`
	AwsRole *AwsRole `json:"aws_role,omitempty"`
}

// NewKinesisIntegration instantiates a new KinesisIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKinesisIntegration() *KinesisIntegration {
	this := KinesisIntegration{}
	return &this
}

// NewKinesisIntegrationWithDefaults instantiates a new KinesisIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKinesisIntegrationWithDefaults() *KinesisIntegration {
	this := KinesisIntegration{}
	return &this
}

// GetAwsAccessKey returns the AwsAccessKey field value if set, zero value otherwise.
func (o *KinesisIntegration) GetAwsAccessKey() AwsAccessKey {
	if o == nil || o.AwsAccessKey == nil {
		var ret AwsAccessKey
		return ret
	}
	return *o.AwsAccessKey
}

// GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KinesisIntegration) GetAwsAccessKeyOk() (*AwsAccessKey, bool) {
	if o == nil || o.AwsAccessKey == nil {
		return nil, false
	}
	return o.AwsAccessKey, true
}

// HasAwsAccessKey returns a boolean if a field has been set.
func (o *KinesisIntegration) HasAwsAccessKey() bool {
	if o != nil && o.AwsAccessKey != nil {
		return true
	}

	return false
}

// SetAwsAccessKey gets a reference to the given AwsAccessKey and assigns it to the AwsAccessKey field.
func (o *KinesisIntegration) SetAwsAccessKey(v AwsAccessKey) {
	o.AwsAccessKey = &v
}

// GetAwsRole returns the AwsRole field value if set, zero value otherwise.
func (o *KinesisIntegration) GetAwsRole() AwsRole {
	if o == nil || o.AwsRole == nil {
		var ret AwsRole
		return ret
	}
	return *o.AwsRole
}

// GetAwsRoleOk returns a tuple with the AwsRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KinesisIntegration) GetAwsRoleOk() (*AwsRole, bool) {
	if o == nil || o.AwsRole == nil {
		return nil, false
	}
	return o.AwsRole, true
}

// HasAwsRole returns a boolean if a field has been set.
func (o *KinesisIntegration) HasAwsRole() bool {
	if o != nil && o.AwsRole != nil {
		return true
	}

	return false
}

// SetAwsRole gets a reference to the given AwsRole and assigns it to the AwsRole field.
func (o *KinesisIntegration) SetAwsRole(v AwsRole) {
	o.AwsRole = &v
}

func (o KinesisIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsAccessKey != nil {
		toSerialize["aws_access_key"] = o.AwsAccessKey
	}
	if o.AwsRole != nil {
		toSerialize["aws_role"] = o.AwsRole
	}
	return json.Marshal(toSerialize)
}

type NullableKinesisIntegration struct {
	value *KinesisIntegration
	isSet bool
}

func (v NullableKinesisIntegration) Get() *KinesisIntegration {
	return v.value
}

func (v *NullableKinesisIntegration) Set(val *KinesisIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableKinesisIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableKinesisIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKinesisIntegration(val *KinesisIntegration) *NullableKinesisIntegration {
	return &NullableKinesisIntegration{value: val, isSet: true}
}

func (v NullableKinesisIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKinesisIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


