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

// checks if the ListMailboxes200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMailboxes200ResponseInner{}

// ListMailboxes200ResponseInner struct for ListMailboxes200ResponseInner
type ListMailboxes200ResponseInner struct {
	MailboxName *string `json:"mailboxName,omitempty"`
}

// NewListMailboxes200ResponseInner instantiates a new ListMailboxes200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMailboxes200ResponseInner() *ListMailboxes200ResponseInner {
	this := ListMailboxes200ResponseInner{}
	return &this
}

// NewListMailboxes200ResponseInnerWithDefaults instantiates a new ListMailboxes200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMailboxes200ResponseInnerWithDefaults() *ListMailboxes200ResponseInner {
	this := ListMailboxes200ResponseInner{}
	return &this
}

// GetMailboxName returns the MailboxName field value if set, zero value otherwise.
func (o *ListMailboxes200ResponseInner) GetMailboxName() string {
	if o == nil || IsNil(o.MailboxName) {
		var ret string
		return ret
	}
	return *o.MailboxName
}

// GetMailboxNameOk returns a tuple with the MailboxName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMailboxes200ResponseInner) GetMailboxNameOk() (*string, bool) {
	if o == nil || IsNil(o.MailboxName) {
		return nil, false
	}
	return o.MailboxName, true
}

// HasMailboxName returns a boolean if a field has been set.
func (o *ListMailboxes200ResponseInner) HasMailboxName() bool {
	if o != nil && !IsNil(o.MailboxName) {
		return true
	}

	return false
}

// SetMailboxName gets a reference to the given string and assigns it to the MailboxName field.
func (o *ListMailboxes200ResponseInner) SetMailboxName(v string) {
	o.MailboxName = &v
}

func (o ListMailboxes200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMailboxes200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MailboxName) {
		toSerialize["mailboxName"] = o.MailboxName
	}
	return toSerialize, nil
}

type NullableListMailboxes200ResponseInner struct {
	value *ListMailboxes200ResponseInner
	isSet bool
}

func (v NullableListMailboxes200ResponseInner) Get() *ListMailboxes200ResponseInner {
	return v.value
}

func (v *NullableListMailboxes200ResponseInner) Set(val *ListMailboxes200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListMailboxes200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListMailboxes200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMailboxes200ResponseInner(val *ListMailboxes200ResponseInner) *NullableListMailboxes200ResponseInner {
	return &NullableListMailboxes200ResponseInner{value: val, isSet: true}
}

func (v NullableListMailboxes200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMailboxes200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

