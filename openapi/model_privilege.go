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

// Privilege struct for Privilege
type Privilege struct {
	// The action allowed by this privilege.
	Action *string `json:"action,omitempty"`
	// Cluster ID (`usw2a1` for us-west-2, `use1a1` for us-east-1, `euc1a1` for eu-central-1) for which the action is allowed. Defaults to '*All*' if not specified.
	Cluster *string `json:"cluster,omitempty"`
	// The resources on which the action is allowed. Defaults to '*All*' if not specified.
	ResourceName *string `json:"resource_name,omitempty"`
}

// NewPrivilege instantiates a new Privilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilege() *Privilege {
	this := Privilege{}
	return &this
}

// NewPrivilegeWithDefaults instantiates a new Privilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegeWithDefaults() *Privilege {
	this := Privilege{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Privilege) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Privilege) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Privilege) SetAction(v string) {
	o.Action = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *Privilege) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *Privilege) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *Privilege) SetCluster(v string) {
	o.Cluster = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *Privilege) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *Privilege) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *Privilege) SetResourceName(v string) {
	o.ResourceName = &v
}

func (o Privilege) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	return json.Marshal(toSerialize)
}

type NullablePrivilege struct {
	value *Privilege
	isSet bool
}

func (v NullablePrivilege) Get() *Privilege {
	return v.value
}

func (v *NullablePrivilege) Set(val *Privilege) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilege(val *Privilege) *NullablePrivilege {
	return &NullablePrivilege{value: val, isSet: true}
}

func (v NullablePrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


