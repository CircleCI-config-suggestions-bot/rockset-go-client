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

// CsvParams struct for CsvParams
type CsvParams struct {
	// Names of columns.
	ColumnNames []string `json:"columnNames,omitempty"`
	// Names of columns.
	ColumnTypes []string `json:"columnTypes,omitempty"`
	// One of: UTF-8, ISO_8859_1, UTF-16.
	Encoding *string `json:"encoding,omitempty"`
	// escape character removes any special meaning from the following character,default is '\\'
	EscapeChar *string `json:"escapeChar,omitempty"`
	// If the first line in every object specifies the column names.
	FirstLineAsColumnNames *bool `json:"firstLineAsColumnNames,omitempty"`
	// character within which a cell value is enclosed,null character if no such character, default is '\"'
	QuoteChar *string `json:"quoteChar,omitempty"`
	// A single character that is the column separator.
	Separator *string `json:"separator,omitempty"`
}

// NewCsvParams instantiates a new CsvParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsvParams() *CsvParams {
	this := CsvParams{}
	return &this
}

// NewCsvParamsWithDefaults instantiates a new CsvParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsvParamsWithDefaults() *CsvParams {
	this := CsvParams{}
	return &this
}

// GetColumnNames returns the ColumnNames field value if set, zero value otherwise.
func (o *CsvParams) GetColumnNames() []string {
	if o == nil || o.ColumnNames == nil {
		var ret []string
		return ret
	}
	return o.ColumnNames
}

// GetColumnNamesOk returns a tuple with the ColumnNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvParams) GetColumnNamesOk() ([]string, bool) {
	if o == nil || o.ColumnNames == nil {
		return nil, false
	}
	return o.ColumnNames, true
}

// HasColumnNames returns a boolean if a field has been set.
func (o *CsvParams) HasColumnNames() bool {
	if o != nil && o.ColumnNames != nil {
		return true
	}

	return false
}

// SetColumnNames gets a reference to the given []string and assigns it to the ColumnNames field.
func (o *CsvParams) SetColumnNames(v []string) {
	o.ColumnNames = v
}

// GetColumnTypes returns the ColumnTypes field value if set, zero value otherwise.
func (o *CsvParams) GetColumnTypes() []string {
	if o == nil || o.ColumnTypes == nil {
		var ret []string
		return ret
	}
	return o.ColumnTypes
}

// GetColumnTypesOk returns a tuple with the ColumnTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvParams) GetColumnTypesOk() ([]string, bool) {
	if o == nil || o.ColumnTypes == nil {
		return nil, false
	}
	return o.ColumnTypes, true
}

// HasColumnTypes returns a boolean if a field has been set.
func (o *CsvParams) HasColumnTypes() bool {
	if o != nil && o.ColumnTypes != nil {
		return true
	}

	return false
}

// SetColumnTypes gets a reference to the given []string and assigns it to the ColumnTypes field.
func (o *CsvParams) SetColumnTypes(v []string) {
	o.ColumnTypes = v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *CsvParams) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvParams) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *CsvParams) HasEncoding() bool {
	if o != nil && o.Encoding != nil {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *CsvParams) SetEncoding(v string) {
	o.Encoding = &v
}

// GetEscapeChar returns the EscapeChar field value if set, zero value otherwise.
func (o *CsvParams) GetEscapeChar() string {
	if o == nil || o.EscapeChar == nil {
		var ret string
		return ret
	}
	return *o.EscapeChar
}

// GetEscapeCharOk returns a tuple with the EscapeChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvParams) GetEscapeCharOk() (*string, bool) {
	if o == nil || o.EscapeChar == nil {
		return nil, false
	}
	return o.EscapeChar, true
}

// HasEscapeChar returns a boolean if a field has been set.
func (o *CsvParams) HasEscapeChar() bool {
	if o != nil && o.EscapeChar != nil {
		return true
	}

	return false
}

// SetEscapeChar gets a reference to the given string and assigns it to the EscapeChar field.
func (o *CsvParams) SetEscapeChar(v string) {
	o.EscapeChar = &v
}

// GetFirstLineAsColumnNames returns the FirstLineAsColumnNames field value if set, zero value otherwise.
func (o *CsvParams) GetFirstLineAsColumnNames() bool {
	if o == nil || o.FirstLineAsColumnNames == nil {
		var ret bool
		return ret
	}
	return *o.FirstLineAsColumnNames
}

// GetFirstLineAsColumnNamesOk returns a tuple with the FirstLineAsColumnNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvParams) GetFirstLineAsColumnNamesOk() (*bool, bool) {
	if o == nil || o.FirstLineAsColumnNames == nil {
		return nil, false
	}
	return o.FirstLineAsColumnNames, true
}

// HasFirstLineAsColumnNames returns a boolean if a field has been set.
func (o *CsvParams) HasFirstLineAsColumnNames() bool {
	if o != nil && o.FirstLineAsColumnNames != nil {
		return true
	}

	return false
}

// SetFirstLineAsColumnNames gets a reference to the given bool and assigns it to the FirstLineAsColumnNames field.
func (o *CsvParams) SetFirstLineAsColumnNames(v bool) {
	o.FirstLineAsColumnNames = &v
}

// GetQuoteChar returns the QuoteChar field value if set, zero value otherwise.
func (o *CsvParams) GetQuoteChar() string {
	if o == nil || o.QuoteChar == nil {
		var ret string
		return ret
	}
	return *o.QuoteChar
}

// GetQuoteCharOk returns a tuple with the QuoteChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvParams) GetQuoteCharOk() (*string, bool) {
	if o == nil || o.QuoteChar == nil {
		return nil, false
	}
	return o.QuoteChar, true
}

// HasQuoteChar returns a boolean if a field has been set.
func (o *CsvParams) HasQuoteChar() bool {
	if o != nil && o.QuoteChar != nil {
		return true
	}

	return false
}

// SetQuoteChar gets a reference to the given string and assigns it to the QuoteChar field.
func (o *CsvParams) SetQuoteChar(v string) {
	o.QuoteChar = &v
}

// GetSeparator returns the Separator field value if set, zero value otherwise.
func (o *CsvParams) GetSeparator() string {
	if o == nil || o.Separator == nil {
		var ret string
		return ret
	}
	return *o.Separator
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvParams) GetSeparatorOk() (*string, bool) {
	if o == nil || o.Separator == nil {
		return nil, false
	}
	return o.Separator, true
}

// HasSeparator returns a boolean if a field has been set.
func (o *CsvParams) HasSeparator() bool {
	if o != nil && o.Separator != nil {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given string and assigns it to the Separator field.
func (o *CsvParams) SetSeparator(v string) {
	o.Separator = &v
}

func (o CsvParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColumnNames != nil {
		toSerialize["columnNames"] = o.ColumnNames
	}
	if o.ColumnTypes != nil {
		toSerialize["columnTypes"] = o.ColumnTypes
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.EscapeChar != nil {
		toSerialize["escapeChar"] = o.EscapeChar
	}
	if o.FirstLineAsColumnNames != nil {
		toSerialize["firstLineAsColumnNames"] = o.FirstLineAsColumnNames
	}
	if o.QuoteChar != nil {
		toSerialize["quoteChar"] = o.QuoteChar
	}
	if o.Separator != nil {
		toSerialize["separator"] = o.Separator
	}
	return json.Marshal(toSerialize)
}

type NullableCsvParams struct {
	value *CsvParams
	isSet bool
}

func (v NullableCsvParams) Get() *CsvParams {
	return v.value
}

func (v *NullableCsvParams) Set(val *CsvParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCsvParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCsvParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsvParams(val *CsvParams) *NullableCsvParams {
	return &NullableCsvParams{value: val, isSet: true}
}

func (v NullableCsvParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsvParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


