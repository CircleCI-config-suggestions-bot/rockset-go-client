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

// AzureBlobStorageIntegration struct for AzureBlobStorageIntegration
type AzureBlobStorageIntegration struct {
	// Credentials for the Azure Blob Service.
	ConnectionString string `json:"connection_string"`
}

// NewAzureBlobStorageIntegration instantiates a new AzureBlobStorageIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStorageIntegration(connectionString string) *AzureBlobStorageIntegration {
	this := AzureBlobStorageIntegration{}
	this.ConnectionString = connectionString
	return &this
}

// NewAzureBlobStorageIntegrationWithDefaults instantiates a new AzureBlobStorageIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStorageIntegrationWithDefaults() *AzureBlobStorageIntegration {
	this := AzureBlobStorageIntegration{}
	return &this
}

// GetConnectionString returns the ConnectionString field value
func (o *AzureBlobStorageIntegration) GetConnectionString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageIntegration) GetConnectionStringOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectionString, true
}

// SetConnectionString sets field value
func (o *AzureBlobStorageIntegration) SetConnectionString(v string) {
	o.ConnectionString = v
}

func (o AzureBlobStorageIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connection_string"] = o.ConnectionString
	}
	return json.Marshal(toSerialize)
}

type NullableAzureBlobStorageIntegration struct {
	value *AzureBlobStorageIntegration
	isSet bool
}

func (v NullableAzureBlobStorageIntegration) Get() *AzureBlobStorageIntegration {
	return v.value
}

func (v *NullableAzureBlobStorageIntegration) Set(val *AzureBlobStorageIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStorageIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStorageIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStorageIntegration(val *AzureBlobStorageIntegration) *NullableAzureBlobStorageIntegration {
	return &NullableAzureBlobStorageIntegration{value: val, isSet: true}
}

func (v NullableAzureBlobStorageIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStorageIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


