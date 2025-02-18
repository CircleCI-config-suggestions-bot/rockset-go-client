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

// SourceKinesis struct for SourceKinesis
type SourceKinesis struct {
	// AWS region name of Kinesis stream, by default us-west-2 is used.
	AwsRegion *string `json:"aws_region,omitempty"`
	// Set of fields that correspond to a DMS primary key.
	DmsPrimaryKey []string `json:"dms_primary_key,omitempty"`
	// For non-DMS streams, Rockset can tail from the earliest end or latest end of kinesis source.
	OffsetResetPolicy *string `json:"offset_reset_policy,omitempty"`
	// Name of kinesis stream.
	StreamName string `json:"stream_name"`
}

// NewSourceKinesis instantiates a new SourceKinesis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceKinesis(streamName string) *SourceKinesis {
	this := SourceKinesis{}
	this.StreamName = streamName
	return &this
}

// NewSourceKinesisWithDefaults instantiates a new SourceKinesis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceKinesisWithDefaults() *SourceKinesis {
	this := SourceKinesis{}
	return &this
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *SourceKinesis) GetAwsRegion() string {
	if o == nil || o.AwsRegion == nil {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKinesis) GetAwsRegionOk() (*string, bool) {
	if o == nil || o.AwsRegion == nil {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *SourceKinesis) HasAwsRegion() bool {
	if o != nil && o.AwsRegion != nil {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *SourceKinesis) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetDmsPrimaryKey returns the DmsPrimaryKey field value if set, zero value otherwise.
func (o *SourceKinesis) GetDmsPrimaryKey() []string {
	if o == nil || o.DmsPrimaryKey == nil {
		var ret []string
		return ret
	}
	return o.DmsPrimaryKey
}

// GetDmsPrimaryKeyOk returns a tuple with the DmsPrimaryKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKinesis) GetDmsPrimaryKeyOk() ([]string, bool) {
	if o == nil || o.DmsPrimaryKey == nil {
		return nil, false
	}
	return o.DmsPrimaryKey, true
}

// HasDmsPrimaryKey returns a boolean if a field has been set.
func (o *SourceKinesis) HasDmsPrimaryKey() bool {
	if o != nil && o.DmsPrimaryKey != nil {
		return true
	}

	return false
}

// SetDmsPrimaryKey gets a reference to the given []string and assigns it to the DmsPrimaryKey field.
func (o *SourceKinesis) SetDmsPrimaryKey(v []string) {
	o.DmsPrimaryKey = v
}

// GetOffsetResetPolicy returns the OffsetResetPolicy field value if set, zero value otherwise.
func (o *SourceKinesis) GetOffsetResetPolicy() string {
	if o == nil || o.OffsetResetPolicy == nil {
		var ret string
		return ret
	}
	return *o.OffsetResetPolicy
}

// GetOffsetResetPolicyOk returns a tuple with the OffsetResetPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKinesis) GetOffsetResetPolicyOk() (*string, bool) {
	if o == nil || o.OffsetResetPolicy == nil {
		return nil, false
	}
	return o.OffsetResetPolicy, true
}

// HasOffsetResetPolicy returns a boolean if a field has been set.
func (o *SourceKinesis) HasOffsetResetPolicy() bool {
	if o != nil && o.OffsetResetPolicy != nil {
		return true
	}

	return false
}

// SetOffsetResetPolicy gets a reference to the given string and assigns it to the OffsetResetPolicy field.
func (o *SourceKinesis) SetOffsetResetPolicy(v string) {
	o.OffsetResetPolicy = &v
}

// GetStreamName returns the StreamName field value
func (o *SourceKinesis) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *SourceKinesis) GetStreamNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *SourceKinesis) SetStreamName(v string) {
	o.StreamName = v
}

func (o SourceKinesis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsRegion != nil {
		toSerialize["aws_region"] = o.AwsRegion
	}
	if o.DmsPrimaryKey != nil {
		toSerialize["dms_primary_key"] = o.DmsPrimaryKey
	}
	if o.OffsetResetPolicy != nil {
		toSerialize["offset_reset_policy"] = o.OffsetResetPolicy
	}
	if true {
		toSerialize["stream_name"] = o.StreamName
	}
	return json.Marshal(toSerialize)
}

type NullableSourceKinesis struct {
	value *SourceKinesis
	isSet bool
}

func (v NullableSourceKinesis) Get() *SourceKinesis {
	return v.value
}

func (v *NullableSourceKinesis) Set(val *SourceKinesis) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceKinesis) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceKinesis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceKinesis(val *SourceKinesis) *NullableSourceKinesis {
	return &NullableSourceKinesis{value: val, isSet: true}
}

func (v NullableSourceKinesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceKinesis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


