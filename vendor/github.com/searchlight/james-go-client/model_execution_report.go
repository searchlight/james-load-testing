/*
Apache JAMES Web Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ExecutionReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionReport{}

// ExecutionReport struct for ExecutionReport
type ExecutionReport struct {
	// Additional information about the task
	AdditionalInformation map[string]interface{} `json:"additionalInformation,omitempty"`
	// The cancellation date of the task (null if not cancelled)
	CancelledDate NullableTime `json:"cancelledDate,omitempty"`
	// The completion date of the task execution
	CompletedDate *time.Time `json:"completedDate,omitempty"`
	// The failure date of the task (null if not failed)
	FailedDate NullableTime `json:"failedDate,omitempty"`
	// The start date of the task execution
	StartedDate *time.Time `json:"startedDate,omitempty"`
	// The status of the task
	Status *string `json:"status,omitempty"`
	// The submission date of the task
	SubmitDate *time.Time `json:"submitDate,omitempty"`
	// The ID of the task
	TaskId *string `json:"taskId,omitempty"`
	// The type of the task
	Type *string `json:"type,omitempty"`
}

// NewExecutionReport instantiates a new ExecutionReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionReport() *ExecutionReport {
	this := ExecutionReport{}
	return &this
}

// NewExecutionReportWithDefaults instantiates a new ExecutionReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionReportWithDefaults() *ExecutionReport {
	this := ExecutionReport{}
	return &this
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *ExecutionReport) GetAdditionalInformation() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalInformation) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetAdditionalInformationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalInformation) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *ExecutionReport) HasAdditionalInformation() bool {
	if o != nil && !IsNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given map[string]interface{} and assigns it to the AdditionalInformation field.
func (o *ExecutionReport) SetAdditionalInformation(v map[string]interface{}) {
	o.AdditionalInformation = v
}

// GetCancelledDate returns the CancelledDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecutionReport) GetCancelledDate() time.Time {
	if o == nil || IsNil(o.CancelledDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CancelledDate.Get()
}

// GetCancelledDateOk returns a tuple with the CancelledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecutionReport) GetCancelledDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelledDate.Get(), o.CancelledDate.IsSet()
}

// HasCancelledDate returns a boolean if a field has been set.
func (o *ExecutionReport) HasCancelledDate() bool {
	if o != nil && o.CancelledDate.IsSet() {
		return true
	}

	return false
}

// SetCancelledDate gets a reference to the given NullableTime and assigns it to the CancelledDate field.
func (o *ExecutionReport) SetCancelledDate(v time.Time) {
	o.CancelledDate.Set(&v)
}
// SetCancelledDateNil sets the value for CancelledDate to be an explicit nil
func (o *ExecutionReport) SetCancelledDateNil() {
	o.CancelledDate.Set(nil)
}

// UnsetCancelledDate ensures that no value is present for CancelledDate, not even an explicit nil
func (o *ExecutionReport) UnsetCancelledDate() {
	o.CancelledDate.Unset()
}

// GetCompletedDate returns the CompletedDate field value if set, zero value otherwise.
func (o *ExecutionReport) GetCompletedDate() time.Time {
	if o == nil || IsNil(o.CompletedDate) {
		var ret time.Time
		return ret
	}
	return *o.CompletedDate
}

// GetCompletedDateOk returns a tuple with the CompletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetCompletedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedDate) {
		return nil, false
	}
	return o.CompletedDate, true
}

// HasCompletedDate returns a boolean if a field has been set.
func (o *ExecutionReport) HasCompletedDate() bool {
	if o != nil && !IsNil(o.CompletedDate) {
		return true
	}

	return false
}

// SetCompletedDate gets a reference to the given time.Time and assigns it to the CompletedDate field.
func (o *ExecutionReport) SetCompletedDate(v time.Time) {
	o.CompletedDate = &v
}

// GetFailedDate returns the FailedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecutionReport) GetFailedDate() time.Time {
	if o == nil || IsNil(o.FailedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.FailedDate.Get()
}

// GetFailedDateOk returns a tuple with the FailedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecutionReport) GetFailedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedDate.Get(), o.FailedDate.IsSet()
}

// HasFailedDate returns a boolean if a field has been set.
func (o *ExecutionReport) HasFailedDate() bool {
	if o != nil && o.FailedDate.IsSet() {
		return true
	}

	return false
}

// SetFailedDate gets a reference to the given NullableTime and assigns it to the FailedDate field.
func (o *ExecutionReport) SetFailedDate(v time.Time) {
	o.FailedDate.Set(&v)
}
// SetFailedDateNil sets the value for FailedDate to be an explicit nil
func (o *ExecutionReport) SetFailedDateNil() {
	o.FailedDate.Set(nil)
}

// UnsetFailedDate ensures that no value is present for FailedDate, not even an explicit nil
func (o *ExecutionReport) UnsetFailedDate() {
	o.FailedDate.Unset()
}

// GetStartedDate returns the StartedDate field value if set, zero value otherwise.
func (o *ExecutionReport) GetStartedDate() time.Time {
	if o == nil || IsNil(o.StartedDate) {
		var ret time.Time
		return ret
	}
	return *o.StartedDate
}

// GetStartedDateOk returns a tuple with the StartedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetStartedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedDate) {
		return nil, false
	}
	return o.StartedDate, true
}

// HasStartedDate returns a boolean if a field has been set.
func (o *ExecutionReport) HasStartedDate() bool {
	if o != nil && !IsNil(o.StartedDate) {
		return true
	}

	return false
}

// SetStartedDate gets a reference to the given time.Time and assigns it to the StartedDate field.
func (o *ExecutionReport) SetStartedDate(v time.Time) {
	o.StartedDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExecutionReport) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExecutionReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExecutionReport) SetStatus(v string) {
	o.Status = &v
}

// GetSubmitDate returns the SubmitDate field value if set, zero value otherwise.
func (o *ExecutionReport) GetSubmitDate() time.Time {
	if o == nil || IsNil(o.SubmitDate) {
		var ret time.Time
		return ret
	}
	return *o.SubmitDate
}

// GetSubmitDateOk returns a tuple with the SubmitDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSubmitDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmitDate) {
		return nil, false
	}
	return o.SubmitDate, true
}

// HasSubmitDate returns a boolean if a field has been set.
func (o *ExecutionReport) HasSubmitDate() bool {
	if o != nil && !IsNil(o.SubmitDate) {
		return true
	}

	return false
}

// SetSubmitDate gets a reference to the given time.Time and assigns it to the SubmitDate field.
func (o *ExecutionReport) SetSubmitDate(v time.Time) {
	o.SubmitDate = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *ExecutionReport) GetTaskId() string {
	if o == nil || IsNil(o.TaskId) {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *ExecutionReport) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *ExecutionReport) SetTaskId(v string) {
	o.TaskId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExecutionReport) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExecutionReport) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExecutionReport) SetType(v string) {
	o.Type = &v
}

func (o ExecutionReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalInformation) {
		toSerialize["additionalInformation"] = o.AdditionalInformation
	}
	if o.CancelledDate.IsSet() {
		toSerialize["cancelledDate"] = o.CancelledDate.Get()
	}
	if !IsNil(o.CompletedDate) {
		toSerialize["completedDate"] = o.CompletedDate
	}
	if o.FailedDate.IsSet() {
		toSerialize["failedDate"] = o.FailedDate.Get()
	}
	if !IsNil(o.StartedDate) {
		toSerialize["startedDate"] = o.StartedDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubmitDate) {
		toSerialize["submitDate"] = o.SubmitDate
	}
	if !IsNil(o.TaskId) {
		toSerialize["taskId"] = o.TaskId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableExecutionReport struct {
	value *ExecutionReport
	isSet bool
}

func (v NullableExecutionReport) Get() *ExecutionReport {
	return v.value
}

func (v *NullableExecutionReport) Set(val *ExecutionReport) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionReport) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionReport(val *ExecutionReport) *NullableExecutionReport {
	return &NullableExecutionReport{value: val, isSet: true}
}

func (v NullableExecutionReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

