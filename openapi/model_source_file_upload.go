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

// SourceFileUpload struct for SourceFileUpload
type SourceFileUpload struct {
	// Name of the file.
	FileName string `json:"file_name"`
	// Size of the file in bytes.
	FileSize int64 `json:"file_size"`
	// Time of file upload.
	FileUploadTime string `json:"file_upload_time"`
}

// NewSourceFileUpload instantiates a new SourceFileUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceFileUpload(fileName string, fileSize int64, fileUploadTime string) *SourceFileUpload {
	this := SourceFileUpload{}
	this.FileName = fileName
	this.FileSize = fileSize
	this.FileUploadTime = fileUploadTime
	return &this
}

// NewSourceFileUploadWithDefaults instantiates a new SourceFileUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceFileUploadWithDefaults() *SourceFileUpload {
	this := SourceFileUpload{}
	return &this
}

// GetFileName returns the FileName field value
func (o *SourceFileUpload) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *SourceFileUpload) GetFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *SourceFileUpload) SetFileName(v string) {
	o.FileName = v
}

// GetFileSize returns the FileSize field value
func (o *SourceFileUpload) GetFileSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *SourceFileUpload) GetFileSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *SourceFileUpload) SetFileSize(v int64) {
	o.FileSize = v
}

// GetFileUploadTime returns the FileUploadTime field value
func (o *SourceFileUpload) GetFileUploadTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileUploadTime
}

// GetFileUploadTimeOk returns a tuple with the FileUploadTime field value
// and a boolean to check if the value has been set.
func (o *SourceFileUpload) GetFileUploadTimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FileUploadTime, true
}

// SetFileUploadTime sets field value
func (o *SourceFileUpload) SetFileUploadTime(v string) {
	o.FileUploadTime = v
}

func (o SourceFileUpload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["file_name"] = o.FileName
	}
	if true {
		toSerialize["file_size"] = o.FileSize
	}
	if true {
		toSerialize["file_upload_time"] = o.FileUploadTime
	}
	return json.Marshal(toSerialize)
}

type NullableSourceFileUpload struct {
	value *SourceFileUpload
	isSet bool
}

func (v NullableSourceFileUpload) Get() *SourceFileUpload {
	return v.value
}

func (v *NullableSourceFileUpload) Set(val *SourceFileUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceFileUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceFileUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceFileUpload(val *SourceFileUpload) *NullableSourceFileUpload {
	return &NullableSourceFileUpload{value: val, isSet: true}
}

func (v NullableSourceFileUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceFileUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


