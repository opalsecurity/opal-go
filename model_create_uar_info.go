/*
Opal API

The Opal API is a RESTful API that allows you to interact with the Opal Security platform programmatically.

API version: 1.0
Contact: hello@opal.dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opal

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CreateUARInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUARInfo{}

// CreateUARInfo Information needed to start a user access review.
type CreateUARInfo struct {
	// The name of the UAR.
	Name string `json:"name"`
	ReviewerAssignmentPolicy UARReviewerAssignmentPolicyEnum `json:"reviewer_assignment_policy"`
	// A bool representing whether to send a notification to reviewers when they're assigned a new review. Default is False.
	SendReviewerAssignmentNotification bool `json:"send_reviewer_assignment_notification"`
	// The last day for reviewers to complete their access reviews.
	Deadline time.Time `json:"deadline"`
	// The time zone name (as defined by the IANA Time Zone database) used in the access review deadline and exported audit report. Default is America/Los_Angeles.
	TimeZone string `json:"time_zone"`
	// A bool representing whether to present a warning when a user is the only reviewer for themself. Default is False.
	SelfReviewAllowed bool `json:"self_review_allowed"`
	ReminderSchedule []int32 `json:"reminder_schedule,omitempty"`
	ReminderIncludeManager *bool `json:"reminder_include_manager,omitempty"`
	UarScope *UARScope `json:"uar_scope,omitempty"`
}

type _CreateUARInfo CreateUARInfo

// NewCreateUARInfo instantiates a new CreateUARInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUARInfo(name string, reviewerAssignmentPolicy UARReviewerAssignmentPolicyEnum, sendReviewerAssignmentNotification bool, deadline time.Time, timeZone string, selfReviewAllowed bool) *CreateUARInfo {
	this := CreateUARInfo{}
	this.Name = name
	this.ReviewerAssignmentPolicy = reviewerAssignmentPolicy
	this.SendReviewerAssignmentNotification = sendReviewerAssignmentNotification
	this.Deadline = deadline
	this.TimeZone = timeZone
	this.SelfReviewAllowed = selfReviewAllowed
	return &this
}

// NewCreateUARInfoWithDefaults instantiates a new CreateUARInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUARInfoWithDefaults() *CreateUARInfo {
	this := CreateUARInfo{}
	return &this
}

// GetName returns the Name field value
func (o *CreateUARInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateUARInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateUARInfo) SetName(v string) {
	o.Name = v
}

// GetReviewerAssignmentPolicy returns the ReviewerAssignmentPolicy field value
func (o *CreateUARInfo) GetReviewerAssignmentPolicy() UARReviewerAssignmentPolicyEnum {
	if o == nil {
		var ret UARReviewerAssignmentPolicyEnum
		return ret
	}

	return o.ReviewerAssignmentPolicy
}

// GetReviewerAssignmentPolicyOk returns a tuple with the ReviewerAssignmentPolicy field value
// and a boolean to check if the value has been set.
func (o *CreateUARInfo) GetReviewerAssignmentPolicyOk() (*UARReviewerAssignmentPolicyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerAssignmentPolicy, true
}

// SetReviewerAssignmentPolicy sets field value
func (o *CreateUARInfo) SetReviewerAssignmentPolicy(v UARReviewerAssignmentPolicyEnum) {
	o.ReviewerAssignmentPolicy = v
}

// GetSendReviewerAssignmentNotification returns the SendReviewerAssignmentNotification field value
func (o *CreateUARInfo) GetSendReviewerAssignmentNotification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SendReviewerAssignmentNotification
}

// GetSendReviewerAssignmentNotificationOk returns a tuple with the SendReviewerAssignmentNotification field value
// and a boolean to check if the value has been set.
func (o *CreateUARInfo) GetSendReviewerAssignmentNotificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendReviewerAssignmentNotification, true
}

// SetSendReviewerAssignmentNotification sets field value
func (o *CreateUARInfo) SetSendReviewerAssignmentNotification(v bool) {
	o.SendReviewerAssignmentNotification = v
}

// GetDeadline returns the Deadline field value
func (o *CreateUARInfo) GetDeadline() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value
// and a boolean to check if the value has been set.
func (o *CreateUARInfo) GetDeadlineOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deadline, true
}

// SetDeadline sets field value
func (o *CreateUARInfo) SetDeadline(v time.Time) {
	o.Deadline = v
}

// GetTimeZone returns the TimeZone field value
func (o *CreateUARInfo) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *CreateUARInfo) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *CreateUARInfo) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetSelfReviewAllowed returns the SelfReviewAllowed field value
func (o *CreateUARInfo) GetSelfReviewAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelfReviewAllowed
}

// GetSelfReviewAllowedOk returns a tuple with the SelfReviewAllowed field value
// and a boolean to check if the value has been set.
func (o *CreateUARInfo) GetSelfReviewAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfReviewAllowed, true
}

// SetSelfReviewAllowed sets field value
func (o *CreateUARInfo) SetSelfReviewAllowed(v bool) {
	o.SelfReviewAllowed = v
}

// GetReminderSchedule returns the ReminderSchedule field value if set, zero value otherwise.
func (o *CreateUARInfo) GetReminderSchedule() []int32 {
	if o == nil || IsNil(o.ReminderSchedule) {
		var ret []int32
		return ret
	}
	return o.ReminderSchedule
}

// GetReminderScheduleOk returns a tuple with the ReminderSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUARInfo) GetReminderScheduleOk() ([]int32, bool) {
	if o == nil || IsNil(o.ReminderSchedule) {
		return nil, false
	}
	return o.ReminderSchedule, true
}

// HasReminderSchedule returns a boolean if a field has been set.
func (o *CreateUARInfo) HasReminderSchedule() bool {
	if o != nil && !IsNil(o.ReminderSchedule) {
		return true
	}

	return false
}

// SetReminderSchedule gets a reference to the given []int32 and assigns it to the ReminderSchedule field.
func (o *CreateUARInfo) SetReminderSchedule(v []int32) {
	o.ReminderSchedule = v
}

// GetReminderIncludeManager returns the ReminderIncludeManager field value if set, zero value otherwise.
func (o *CreateUARInfo) GetReminderIncludeManager() bool {
	if o == nil || IsNil(o.ReminderIncludeManager) {
		var ret bool
		return ret
	}
	return *o.ReminderIncludeManager
}

// GetReminderIncludeManagerOk returns a tuple with the ReminderIncludeManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUARInfo) GetReminderIncludeManagerOk() (*bool, bool) {
	if o == nil || IsNil(o.ReminderIncludeManager) {
		return nil, false
	}
	return o.ReminderIncludeManager, true
}

// HasReminderIncludeManager returns a boolean if a field has been set.
func (o *CreateUARInfo) HasReminderIncludeManager() bool {
	if o != nil && !IsNil(o.ReminderIncludeManager) {
		return true
	}

	return false
}

// SetReminderIncludeManager gets a reference to the given bool and assigns it to the ReminderIncludeManager field.
func (o *CreateUARInfo) SetReminderIncludeManager(v bool) {
	o.ReminderIncludeManager = &v
}

// GetUarScope returns the UarScope field value if set, zero value otherwise.
func (o *CreateUARInfo) GetUarScope() UARScope {
	if o == nil || IsNil(o.UarScope) {
		var ret UARScope
		return ret
	}
	return *o.UarScope
}

// GetUarScopeOk returns a tuple with the UarScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUARInfo) GetUarScopeOk() (*UARScope, bool) {
	if o == nil || IsNil(o.UarScope) {
		return nil, false
	}
	return o.UarScope, true
}

// HasUarScope returns a boolean if a field has been set.
func (o *CreateUARInfo) HasUarScope() bool {
	if o != nil && !IsNil(o.UarScope) {
		return true
	}

	return false
}

// SetUarScope gets a reference to the given UARScope and assigns it to the UarScope field.
func (o *CreateUARInfo) SetUarScope(v UARScope) {
	o.UarScope = &v
}

func (o CreateUARInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUARInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["reviewer_assignment_policy"] = o.ReviewerAssignmentPolicy
	toSerialize["send_reviewer_assignment_notification"] = o.SendReviewerAssignmentNotification
	toSerialize["deadline"] = o.Deadline
	toSerialize["time_zone"] = o.TimeZone
	toSerialize["self_review_allowed"] = o.SelfReviewAllowed
	if !IsNil(o.ReminderSchedule) {
		toSerialize["reminder_schedule"] = o.ReminderSchedule
	}
	if !IsNil(o.ReminderIncludeManager) {
		toSerialize["reminder_include_manager"] = o.ReminderIncludeManager
	}
	if !IsNil(o.UarScope) {
		toSerialize["uar_scope"] = o.UarScope
	}
	return toSerialize, nil
}

func (o *CreateUARInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"reviewer_assignment_policy",
		"send_reviewer_assignment_notification",
		"deadline",
		"time_zone",
		"self_review_allowed",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateUARInfo := _CreateUARInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUARInfo)

	if err != nil {
		return err
	}

	*o = CreateUARInfo(varCreateUARInfo)

	return err
}

type NullableCreateUARInfo struct {
	value *CreateUARInfo
	isSet bool
}

func (v NullableCreateUARInfo) Get() *CreateUARInfo {
	return v.value
}

func (v *NullableCreateUARInfo) Set(val *CreateUARInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUARInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUARInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUARInfo(val *CreateUARInfo) *NullableCreateUARInfo {
	return &NullableCreateUARInfo{value: val, isSet: true}
}

func (v NullableCreateUARInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUARInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


