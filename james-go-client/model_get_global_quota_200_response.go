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

// checks if the GetGlobalQuota200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGlobalQuota200Response{}

// GetGlobalQuota200Response struct for GetGlobalQuota200Response
type GetGlobalQuota200Response struct {
	Count *int32 `json:"count,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewGetGlobalQuota200Response instantiates a new GetGlobalQuota200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGlobalQuota200Response() *GetGlobalQuota200Response {
	this := GetGlobalQuota200Response{}
	return &this
}

// NewGetGlobalQuota200ResponseWithDefaults instantiates a new GetGlobalQuota200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGlobalQuota200ResponseWithDefaults() *GetGlobalQuota200Response {
	this := GetGlobalQuota200Response{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetGlobalQuota200Response) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGlobalQuota200Response) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetGlobalQuota200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetGlobalQuota200Response) SetCount(v int32) {
	o.Count = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GetGlobalQuota200Response) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGlobalQuota200Response) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GetGlobalQuota200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *GetGlobalQuota200Response) SetSize(v int32) {
	o.Size = &v
}

func (o GetGlobalQuota200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGlobalQuota200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableGetGlobalQuota200Response struct {
	value *GetGlobalQuota200Response
	isSet bool
}

func (v NullableGetGlobalQuota200Response) Get() *GetGlobalQuota200Response {
	return v.value
}

func (v *NullableGetGlobalQuota200Response) Set(val *GetGlobalQuota200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGlobalQuota200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGlobalQuota200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGlobalQuota200Response(val *GetGlobalQuota200Response) *NullableGetGlobalQuota200Response {
	return &NullableGetGlobalQuota200Response{value: val, isSet: true}
}

func (v NullableGetGlobalQuota200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGlobalQuota200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

