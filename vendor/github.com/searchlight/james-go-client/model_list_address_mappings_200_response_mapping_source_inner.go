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

// checks if the ListAddressMappings200ResponseMappingSourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAddressMappings200ResponseMappingSourceInner{}

// ListAddressMappings200ResponseMappingSourceInner struct for ListAddressMappings200ResponseMappingSourceInner
type ListAddressMappings200ResponseMappingSourceInner struct {
	Mapping *string `json:"mapping,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewListAddressMappings200ResponseMappingSourceInner instantiates a new ListAddressMappings200ResponseMappingSourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAddressMappings200ResponseMappingSourceInner() *ListAddressMappings200ResponseMappingSourceInner {
	this := ListAddressMappings200ResponseMappingSourceInner{}
	return &this
}

// NewListAddressMappings200ResponseMappingSourceInnerWithDefaults instantiates a new ListAddressMappings200ResponseMappingSourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAddressMappings200ResponseMappingSourceInnerWithDefaults() *ListAddressMappings200ResponseMappingSourceInner {
	this := ListAddressMappings200ResponseMappingSourceInner{}
	return &this
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *ListAddressMappings200ResponseMappingSourceInner) GetMapping() string {
	if o == nil || IsNil(o.Mapping) {
		var ret string
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAddressMappings200ResponseMappingSourceInner) GetMappingOk() (*string, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *ListAddressMappings200ResponseMappingSourceInner) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given string and assigns it to the Mapping field.
func (o *ListAddressMappings200ResponseMappingSourceInner) SetMapping(v string) {
	o.Mapping = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListAddressMappings200ResponseMappingSourceInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAddressMappings200ResponseMappingSourceInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListAddressMappings200ResponseMappingSourceInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListAddressMappings200ResponseMappingSourceInner) SetType(v string) {
	o.Type = &v
}

func (o ListAddressMappings200ResponseMappingSourceInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAddressMappings200ResponseMappingSourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableListAddressMappings200ResponseMappingSourceInner struct {
	value *ListAddressMappings200ResponseMappingSourceInner
	isSet bool
}

func (v NullableListAddressMappings200ResponseMappingSourceInner) Get() *ListAddressMappings200ResponseMappingSourceInner {
	return v.value
}

func (v *NullableListAddressMappings200ResponseMappingSourceInner) Set(val *ListAddressMappings200ResponseMappingSourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListAddressMappings200ResponseMappingSourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListAddressMappings200ResponseMappingSourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAddressMappings200ResponseMappingSourceInner(val *ListAddressMappings200ResponseMappingSourceInner) *NullableListAddressMappings200ResponseMappingSourceInner {
	return &NullableListAddressMappings200ResponseMappingSourceInner{value: val, isSet: true}
}

func (v NullableListAddressMappings200ResponseMappingSourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAddressMappings200ResponseMappingSourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

