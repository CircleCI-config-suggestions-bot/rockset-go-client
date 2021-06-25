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

// QueryLambdaTagResponse struct for QueryLambdaTagResponse
type QueryLambdaTagResponse struct {
	Data *QueryLambdaTag `json:"data,omitempty"`
}

// NewQueryLambdaTagResponse instantiates a new QueryLambdaTagResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLambdaTagResponse() *QueryLambdaTagResponse {
	this := QueryLambdaTagResponse{}
	return &this
}

// NewQueryLambdaTagResponseWithDefaults instantiates a new QueryLambdaTagResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLambdaTagResponseWithDefaults() *QueryLambdaTagResponse {
	this := QueryLambdaTagResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *QueryLambdaTagResponse) GetData() QueryLambdaTag {
	if o == nil || o.Data == nil {
		var ret QueryLambdaTag
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaTagResponse) GetDataOk() (*QueryLambdaTag, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *QueryLambdaTagResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given QueryLambdaTag and assigns it to the Data field.
func (o *QueryLambdaTagResponse) SetData(v QueryLambdaTag) {
	o.Data = &v
}

func (o QueryLambdaTagResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableQueryLambdaTagResponse struct {
	value *QueryLambdaTagResponse
	isSet bool
}

func (v NullableQueryLambdaTagResponse) Get() *QueryLambdaTagResponse {
	return v.value
}

func (v *NullableQueryLambdaTagResponse) Set(val *QueryLambdaTagResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLambdaTagResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLambdaTagResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLambdaTagResponse(val *QueryLambdaTagResponse) *NullableQueryLambdaTagResponse {
	return &NullableQueryLambdaTagResponse{value: val, isSet: true}
}

func (v NullableQueryLambdaTagResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLambdaTagResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


