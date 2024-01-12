/*
Apache JAMES Web Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HealthCheckInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckInfo{}

// HealthCheckInfo struct for HealthCheckInfo
type HealthCheckInfo struct {
	ComponentName *string `json:"componentName,omitempty"`
	EscapedComponentName *string `json:"escapedComponentName,omitempty"`
}

// NewHealthCheckInfo instantiates a new HealthCheckInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckInfo() *HealthCheckInfo {
	this := HealthCheckInfo{}
	return &this
}

// NewHealthCheckInfoWithDefaults instantiates a new HealthCheckInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckInfoWithDefaults() *HealthCheckInfo {
	this := HealthCheckInfo{}
	return &this
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *HealthCheckInfo) GetComponentName() string {
	if o == nil || IsNil(o.ComponentName) {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckInfo) GetComponentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentName) {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *HealthCheckInfo) HasComponentName() bool {
	if o != nil && !IsNil(o.ComponentName) {
		return true
	}

	return false
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *HealthCheckInfo) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetEscapedComponentName returns the EscapedComponentName field value if set, zero value otherwise.
func (o *HealthCheckInfo) GetEscapedComponentName() string {
	if o == nil || IsNil(o.EscapedComponentName) {
		var ret string
		return ret
	}
	return *o.EscapedComponentName
}

// GetEscapedComponentNameOk returns a tuple with the EscapedComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckInfo) GetEscapedComponentNameOk() (*string, bool) {
	if o == nil || IsNil(o.EscapedComponentName) {
		return nil, false
	}
	return o.EscapedComponentName, true
}

// HasEscapedComponentName returns a boolean if a field has been set.
func (o *HealthCheckInfo) HasEscapedComponentName() bool {
	if o != nil && !IsNil(o.EscapedComponentName) {
		return true
	}

	return false
}

// SetEscapedComponentName gets a reference to the given string and assigns it to the EscapedComponentName field.
func (o *HealthCheckInfo) SetEscapedComponentName(v string) {
	o.EscapedComponentName = &v
}

func (o HealthCheckInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentName) {
		toSerialize["componentName"] = o.ComponentName
	}
	if !IsNil(o.EscapedComponentName) {
		toSerialize["escapedComponentName"] = o.EscapedComponentName
	}
	return toSerialize, nil
}

type NullableHealthCheckInfo struct {
	value *HealthCheckInfo
	isSet bool
}

func (v NullableHealthCheckInfo) Get() *HealthCheckInfo {
	return v.value
}

func (v *NullableHealthCheckInfo) Set(val *HealthCheckInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckInfo(val *HealthCheckInfo) *NullableHealthCheckInfo {
	return &NullableHealthCheckInfo{value: val, isSet: true}
}

func (v NullableHealthCheckInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

