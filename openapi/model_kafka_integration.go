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

// KafkaIntegration struct for KafkaIntegration
type KafkaIntegration struct {
	AwsRole *AwsRole `json:"aws_role,omitempty"`
	// The Kafka bootstrap server url(s). Required only for V3 integration.
	BootstrapServers *string `json:"bootstrap_servers,omitempty"`
	// Kafka connection string.
	ConnectionString *string `json:"connection_string,omitempty"`
	// The format of the Kafka topics being tailed.
	KafkaDataFormat *string `json:"kafka_data_format,omitempty"`
	// Kafka topics to tail.
	KafkaTopicNames []string `json:"kafka_topic_names,omitempty"`
	SchemaRegistryConfig *SchemaRegistryConfig `json:"schema_registry_config,omitempty"`
	SecurityConfig *KafkaV3SecurityConfig `json:"security_config,omitempty"`
	// The status of the Kafka source by topic.
	SourceStatusByTopic *map[string]StatusKafka `json:"source_status_by_topic,omitempty"`
	UseV3 *bool `json:"use_v3,omitempty"`
}

// NewKafkaIntegration instantiates a new KafkaIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaIntegration() *KafkaIntegration {
	this := KafkaIntegration{}
	return &this
}

// NewKafkaIntegrationWithDefaults instantiates a new KafkaIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaIntegrationWithDefaults() *KafkaIntegration {
	this := KafkaIntegration{}
	return &this
}

// GetAwsRole returns the AwsRole field value if set, zero value otherwise.
func (o *KafkaIntegration) GetAwsRole() AwsRole {
	if o == nil || o.AwsRole == nil {
		var ret AwsRole
		return ret
	}
	return *o.AwsRole
}

// GetAwsRoleOk returns a tuple with the AwsRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaIntegration) GetAwsRoleOk() (*AwsRole, bool) {
	if o == nil || o.AwsRole == nil {
		return nil, false
	}
	return o.AwsRole, true
}

// HasAwsRole returns a boolean if a field has been set.
func (o *KafkaIntegration) HasAwsRole() bool {
	if o != nil && o.AwsRole != nil {
		return true
	}

	return false
}

// SetAwsRole gets a reference to the given AwsRole and assigns it to the AwsRole field.
func (o *KafkaIntegration) SetAwsRole(v AwsRole) {
	o.AwsRole = &v
}

// GetBootstrapServers returns the BootstrapServers field value if set, zero value otherwise.
func (o *KafkaIntegration) GetBootstrapServers() string {
	if o == nil || o.BootstrapServers == nil {
		var ret string
		return ret
	}
	return *o.BootstrapServers
}

// GetBootstrapServersOk returns a tuple with the BootstrapServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaIntegration) GetBootstrapServersOk() (*string, bool) {
	if o == nil || o.BootstrapServers == nil {
		return nil, false
	}
	return o.BootstrapServers, true
}

// HasBootstrapServers returns a boolean if a field has been set.
func (o *KafkaIntegration) HasBootstrapServers() bool {
	if o != nil && o.BootstrapServers != nil {
		return true
	}

	return false
}

// SetBootstrapServers gets a reference to the given string and assigns it to the BootstrapServers field.
func (o *KafkaIntegration) SetBootstrapServers(v string) {
	o.BootstrapServers = &v
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise.
func (o *KafkaIntegration) GetConnectionString() string {
	if o == nil || o.ConnectionString == nil {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaIntegration) GetConnectionStringOk() (*string, bool) {
	if o == nil || o.ConnectionString == nil {
		return nil, false
	}
	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *KafkaIntegration) HasConnectionString() bool {
	if o != nil && o.ConnectionString != nil {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *KafkaIntegration) SetConnectionString(v string) {
	o.ConnectionString = &v
}

// GetKafkaDataFormat returns the KafkaDataFormat field value if set, zero value otherwise.
func (o *KafkaIntegration) GetKafkaDataFormat() string {
	if o == nil || o.KafkaDataFormat == nil {
		var ret string
		return ret
	}
	return *o.KafkaDataFormat
}

// GetKafkaDataFormatOk returns a tuple with the KafkaDataFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaIntegration) GetKafkaDataFormatOk() (*string, bool) {
	if o == nil || o.KafkaDataFormat == nil {
		return nil, false
	}
	return o.KafkaDataFormat, true
}

// HasKafkaDataFormat returns a boolean if a field has been set.
func (o *KafkaIntegration) HasKafkaDataFormat() bool {
	if o != nil && o.KafkaDataFormat != nil {
		return true
	}

	return false
}

// SetKafkaDataFormat gets a reference to the given string and assigns it to the KafkaDataFormat field.
func (o *KafkaIntegration) SetKafkaDataFormat(v string) {
	o.KafkaDataFormat = &v
}

// GetKafkaTopicNames returns the KafkaTopicNames field value if set, zero value otherwise.
func (o *KafkaIntegration) GetKafkaTopicNames() []string {
	if o == nil || o.KafkaTopicNames == nil {
		var ret []string
		return ret
	}
	return o.KafkaTopicNames
}

// GetKafkaTopicNamesOk returns a tuple with the KafkaTopicNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaIntegration) GetKafkaTopicNamesOk() ([]string, bool) {
	if o == nil || o.KafkaTopicNames == nil {
		return nil, false
	}
	return o.KafkaTopicNames, true
}

// HasKafkaTopicNames returns a boolean if a field has been set.
func (o *KafkaIntegration) HasKafkaTopicNames() bool {
	if o != nil && o.KafkaTopicNames != nil {
		return true
	}

	return false
}

// SetKafkaTopicNames gets a reference to the given []string and assigns it to the KafkaTopicNames field.
func (o *KafkaIntegration) SetKafkaTopicNames(v []string) {
	o.KafkaTopicNames = v
}

// GetSchemaRegistryConfig returns the SchemaRegistryConfig field value if set, zero value otherwise.
func (o *KafkaIntegration) GetSchemaRegistryConfig() SchemaRegistryConfig {
	if o == nil || o.SchemaRegistryConfig == nil {
		var ret SchemaRegistryConfig
		return ret
	}
	return *o.SchemaRegistryConfig
}

// GetSchemaRegistryConfigOk returns a tuple with the SchemaRegistryConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaIntegration) GetSchemaRegistryConfigOk() (*SchemaRegistryConfig, bool) {
	if o == nil || o.SchemaRegistryConfig == nil {
		return nil, false
	}
	return o.SchemaRegistryConfig, true
}

// HasSchemaRegistryConfig returns a boolean if a field has been set.
func (o *KafkaIntegration) HasSchemaRegistryConfig() bool {
	if o != nil && o.SchemaRegistryConfig != nil {
		return true
	}

	return false
}

// SetSchemaRegistryConfig gets a reference to the given SchemaRegistryConfig and assigns it to the SchemaRegistryConfig field.
func (o *KafkaIntegration) SetSchemaRegistryConfig(v SchemaRegistryConfig) {
	o.SchemaRegistryConfig = &v
}

// GetSecurityConfig returns the SecurityConfig field value if set, zero value otherwise.
func (o *KafkaIntegration) GetSecurityConfig() KafkaV3SecurityConfig {
	if o == nil || o.SecurityConfig == nil {
		var ret KafkaV3SecurityConfig
		return ret
	}
	return *o.SecurityConfig
}

// GetSecurityConfigOk returns a tuple with the SecurityConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaIntegration) GetSecurityConfigOk() (*KafkaV3SecurityConfig, bool) {
	if o == nil || o.SecurityConfig == nil {
		return nil, false
	}
	return o.SecurityConfig, true
}

// HasSecurityConfig returns a boolean if a field has been set.
func (o *KafkaIntegration) HasSecurityConfig() bool {
	if o != nil && o.SecurityConfig != nil {
		return true
	}

	return false
}

// SetSecurityConfig gets a reference to the given KafkaV3SecurityConfig and assigns it to the SecurityConfig field.
func (o *KafkaIntegration) SetSecurityConfig(v KafkaV3SecurityConfig) {
	o.SecurityConfig = &v
}

// GetSourceStatusByTopic returns the SourceStatusByTopic field value if set, zero value otherwise.
func (o *KafkaIntegration) GetSourceStatusByTopic() map[string]StatusKafka {
	if o == nil || o.SourceStatusByTopic == nil {
		var ret map[string]StatusKafka
		return ret
	}
	return *o.SourceStatusByTopic
}

// GetSourceStatusByTopicOk returns a tuple with the SourceStatusByTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaIntegration) GetSourceStatusByTopicOk() (*map[string]StatusKafka, bool) {
	if o == nil || o.SourceStatusByTopic == nil {
		return nil, false
	}
	return o.SourceStatusByTopic, true
}

// HasSourceStatusByTopic returns a boolean if a field has been set.
func (o *KafkaIntegration) HasSourceStatusByTopic() bool {
	if o != nil && o.SourceStatusByTopic != nil {
		return true
	}

	return false
}

// SetSourceStatusByTopic gets a reference to the given map[string]StatusKafka and assigns it to the SourceStatusByTopic field.
func (o *KafkaIntegration) SetSourceStatusByTopic(v map[string]StatusKafka) {
	o.SourceStatusByTopic = &v
}

// GetUseV3 returns the UseV3 field value if set, zero value otherwise.
func (o *KafkaIntegration) GetUseV3() bool {
	if o == nil || o.UseV3 == nil {
		var ret bool
		return ret
	}
	return *o.UseV3
}

// GetUseV3Ok returns a tuple with the UseV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaIntegration) GetUseV3Ok() (*bool, bool) {
	if o == nil || o.UseV3 == nil {
		return nil, false
	}
	return o.UseV3, true
}

// HasUseV3 returns a boolean if a field has been set.
func (o *KafkaIntegration) HasUseV3() bool {
	if o != nil && o.UseV3 != nil {
		return true
	}

	return false
}

// SetUseV3 gets a reference to the given bool and assigns it to the UseV3 field.
func (o *KafkaIntegration) SetUseV3(v bool) {
	o.UseV3 = &v
}

func (o KafkaIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsRole != nil {
		toSerialize["aws_role"] = o.AwsRole
	}
	if o.BootstrapServers != nil {
		toSerialize["bootstrap_servers"] = o.BootstrapServers
	}
	if o.ConnectionString != nil {
		toSerialize["connection_string"] = o.ConnectionString
	}
	if o.KafkaDataFormat != nil {
		toSerialize["kafka_data_format"] = o.KafkaDataFormat
	}
	if o.KafkaTopicNames != nil {
		toSerialize["kafka_topic_names"] = o.KafkaTopicNames
	}
	if o.SchemaRegistryConfig != nil {
		toSerialize["schema_registry_config"] = o.SchemaRegistryConfig
	}
	if o.SecurityConfig != nil {
		toSerialize["security_config"] = o.SecurityConfig
	}
	if o.SourceStatusByTopic != nil {
		toSerialize["source_status_by_topic"] = o.SourceStatusByTopic
	}
	if o.UseV3 != nil {
		toSerialize["use_v3"] = o.UseV3
	}
	return json.Marshal(toSerialize)
}

type NullableKafkaIntegration struct {
	value *KafkaIntegration
	isSet bool
}

func (v NullableKafkaIntegration) Get() *KafkaIntegration {
	return v.value
}

func (v *NullableKafkaIntegration) Set(val *KafkaIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaIntegration(val *KafkaIntegration) *NullableKafkaIntegration {
	return &NullableKafkaIntegration{value: val, isSet: true}
}

func (v NullableKafkaIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


