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

// UpdateQueryLambdaRequest struct for UpdateQueryLambdaRequest
type UpdateQueryLambdaRequest struct {
	// Optional description.
	Description *string `json:"description,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	Sql *QueryLambdaSql `json:"sql,omitempty"`
}

// NewUpdateQueryLambdaRequest instantiates a new UpdateQueryLambdaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateQueryLambdaRequest() *UpdateQueryLambdaRequest {
	this := UpdateQueryLambdaRequest{}
	return &this
}

// NewUpdateQueryLambdaRequestWithDefaults instantiates a new UpdateQueryLambdaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateQueryLambdaRequestWithDefaults() *UpdateQueryLambdaRequest {
	this := UpdateQueryLambdaRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateQueryLambdaRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQueryLambdaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateQueryLambdaRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateQueryLambdaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *UpdateQueryLambdaRequest) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQueryLambdaRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *UpdateQueryLambdaRequest) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *UpdateQueryLambdaRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *UpdateQueryLambdaRequest) GetSql() QueryLambdaSql {
	if o == nil || o.Sql == nil {
		var ret QueryLambdaSql
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQueryLambdaRequest) GetSqlOk() (*QueryLambdaSql, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *UpdateQueryLambdaRequest) HasSql() bool {
	if o != nil && o.Sql != nil {
		return true
	}

	return false
}

// SetSql gets a reference to the given QueryLambdaSql and assigns it to the Sql field.
func (o *UpdateQueryLambdaRequest) SetSql(v QueryLambdaSql) {
	o.Sql = &v
}

func (o UpdateQueryLambdaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IsPublic != nil {
		toSerialize["is_public"] = o.IsPublic
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateQueryLambdaRequest struct {
	value *UpdateQueryLambdaRequest
	isSet bool
}

func (v NullableUpdateQueryLambdaRequest) Get() *UpdateQueryLambdaRequest {
	return v.value
}

func (v *NullableUpdateQueryLambdaRequest) Set(val *UpdateQueryLambdaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateQueryLambdaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateQueryLambdaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateQueryLambdaRequest(val *UpdateQueryLambdaRequest) *NullableUpdateQueryLambdaRequest {
	return &NullableUpdateQueryLambdaRequest{value: val, isSet: true}
}

func (v NullableUpdateQueryLambdaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateQueryLambdaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


