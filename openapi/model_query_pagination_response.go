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

// QueryPaginationResponse struct for QueryPaginationResponse
type QueryPaginationResponse struct {
	Pagination *PaginationInfo `json:"pagination,omitempty"`
	// List of documents returned by the query.
	Results []map[string]interface{} `json:"results,omitempty"`
	// Total documents returned by the query.
	ResultsTotalDocCount *int64 `json:"results_total_doc_count,omitempty"`
}

// NewQueryPaginationResponse instantiates a new QueryPaginationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPaginationResponse() *QueryPaginationResponse {
	this := QueryPaginationResponse{}
	return &this
}

// NewQueryPaginationResponseWithDefaults instantiates a new QueryPaginationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPaginationResponseWithDefaults() *QueryPaginationResponse {
	this := QueryPaginationResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *QueryPaginationResponse) GetPagination() PaginationInfo {
	if o == nil || o.Pagination == nil {
		var ret PaginationInfo
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPaginationResponse) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *QueryPaginationResponse) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationInfo and assigns it to the Pagination field.
func (o *QueryPaginationResponse) SetPagination(v PaginationInfo) {
	o.Pagination = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *QueryPaginationResponse) GetResults() []map[string]interface{} {
	if o == nil || o.Results == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPaginationResponse) GetResultsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *QueryPaginationResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []map[string]interface{} and assigns it to the Results field.
func (o *QueryPaginationResponse) SetResults(v []map[string]interface{}) {
	o.Results = v
}

// GetResultsTotalDocCount returns the ResultsTotalDocCount field value if set, zero value otherwise.
func (o *QueryPaginationResponse) GetResultsTotalDocCount() int64 {
	if o == nil || o.ResultsTotalDocCount == nil {
		var ret int64
		return ret
	}
	return *o.ResultsTotalDocCount
}

// GetResultsTotalDocCountOk returns a tuple with the ResultsTotalDocCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPaginationResponse) GetResultsTotalDocCountOk() (*int64, bool) {
	if o == nil || o.ResultsTotalDocCount == nil {
		return nil, false
	}
	return o.ResultsTotalDocCount, true
}

// HasResultsTotalDocCount returns a boolean if a field has been set.
func (o *QueryPaginationResponse) HasResultsTotalDocCount() bool {
	if o != nil && o.ResultsTotalDocCount != nil {
		return true
	}

	return false
}

// SetResultsTotalDocCount gets a reference to the given int64 and assigns it to the ResultsTotalDocCount field.
func (o *QueryPaginationResponse) SetResultsTotalDocCount(v int64) {
	o.ResultsTotalDocCount = &v
}

func (o QueryPaginationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.ResultsTotalDocCount != nil {
		toSerialize["results_total_doc_count"] = o.ResultsTotalDocCount
	}
	return json.Marshal(toSerialize)
}

type NullableQueryPaginationResponse struct {
	value *QueryPaginationResponse
	isSet bool
}

func (v NullableQueryPaginationResponse) Get() *QueryPaginationResponse {
	return v.value
}

func (v *NullableQueryPaginationResponse) Set(val *QueryPaginationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPaginationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPaginationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPaginationResponse(val *QueryPaginationResponse) *NullableQueryPaginationResponse {
	return &NullableQueryPaginationResponse{value: val, isSet: true}
}

func (v NullableQueryPaginationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPaginationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


