/*
Opal API

Your Home For Developer Resources.

API version: 1.0
Contact: hello@opal.dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opal

import (
	"encoding/json"
	"time"
)

// UAR A user access review.
type UAR struct {
	// The ID of the UAR.
	UarId string `json:"uar_id"`
	// The name of the UAR.
	Name string `json:"name"`
	ReviewerAssignmentPolicy UARReviewerAssignmentPolicyEnum `json:"reviewer_assignment_policy"`
	// A bool representing whether to send a notification to reviewers when they're assigned a new review. Default is False.
	SendReviewerAssignmentNotification bool `json:"send_reviewer_assignment_notification"`
	// The last day for reviewers to complete their access reviews before access is frozen for incomplete reviews.
	Deadline time.Time `json:"deadline"`
	// The time zone name (as defined by the IANA Time Zone database) used in the access review deadline and exported audit report. Default is America/Los_Angeles.
	TimeZone string `json:"time_zone"`
	// A bool representing whether to present a warning when a user is the only reviewer for themself. Default is False.
	SelfReviewAllowed bool `json:"self_review_allowed"`
	UarScope *UARScope `json:"uar_scope,omitempty"`
}

// NewUAR instantiates a new UAR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUAR(uarId string, name string, reviewerAssignmentPolicy UARReviewerAssignmentPolicyEnum, sendReviewerAssignmentNotification bool, deadline time.Time, timeZone string, selfReviewAllowed bool) *UAR {
	this := UAR{}
	this.UarId = uarId
	this.Name = name
	this.ReviewerAssignmentPolicy = reviewerAssignmentPolicy
	this.SendReviewerAssignmentNotification = sendReviewerAssignmentNotification
	this.Deadline = deadline
	this.TimeZone = timeZone
	this.SelfReviewAllowed = selfReviewAllowed
	return &this
}

// NewUARWithDefaults instantiates a new UAR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUARWithDefaults() *UAR {
	this := UAR{}
	return &this
}

// GetUarId returns the UarId field value
func (o *UAR) GetUarId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UarId
}

// GetUarIdOk returns a tuple with the UarId field value
// and a boolean to check if the value has been set.
func (o *UAR) GetUarIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UarId, true
}

// SetUarId sets field value
func (o *UAR) SetUarId(v string) {
	o.UarId = v
}

// GetName returns the Name field value
func (o *UAR) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UAR) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UAR) SetName(v string) {
	o.Name = v
}

// GetReviewerAssignmentPolicy returns the ReviewerAssignmentPolicy field value
func (o *UAR) GetReviewerAssignmentPolicy() UARReviewerAssignmentPolicyEnum {
	if o == nil {
		var ret UARReviewerAssignmentPolicyEnum
		return ret
	}

	return o.ReviewerAssignmentPolicy
}

// GetReviewerAssignmentPolicyOk returns a tuple with the ReviewerAssignmentPolicy field value
// and a boolean to check if the value has been set.
func (o *UAR) GetReviewerAssignmentPolicyOk() (*UARReviewerAssignmentPolicyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerAssignmentPolicy, true
}

// SetReviewerAssignmentPolicy sets field value
func (o *UAR) SetReviewerAssignmentPolicy(v UARReviewerAssignmentPolicyEnum) {
	o.ReviewerAssignmentPolicy = v
}

// GetSendReviewerAssignmentNotification returns the SendReviewerAssignmentNotification field value
func (o *UAR) GetSendReviewerAssignmentNotification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SendReviewerAssignmentNotification
}

// GetSendReviewerAssignmentNotificationOk returns a tuple with the SendReviewerAssignmentNotification field value
// and a boolean to check if the value has been set.
func (o *UAR) GetSendReviewerAssignmentNotificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendReviewerAssignmentNotification, true
}

// SetSendReviewerAssignmentNotification sets field value
func (o *UAR) SetSendReviewerAssignmentNotification(v bool) {
	o.SendReviewerAssignmentNotification = v
}

// GetDeadline returns the Deadline field value
func (o *UAR) GetDeadline() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value
// and a boolean to check if the value has been set.
func (o *UAR) GetDeadlineOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deadline, true
}

// SetDeadline sets field value
func (o *UAR) SetDeadline(v time.Time) {
	o.Deadline = v
}

// GetTimeZone returns the TimeZone field value
func (o *UAR) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *UAR) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *UAR) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetSelfReviewAllowed returns the SelfReviewAllowed field value
func (o *UAR) GetSelfReviewAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelfReviewAllowed
}

// GetSelfReviewAllowedOk returns a tuple with the SelfReviewAllowed field value
// and a boolean to check if the value has been set.
func (o *UAR) GetSelfReviewAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfReviewAllowed, true
}

// SetSelfReviewAllowed sets field value
func (o *UAR) SetSelfReviewAllowed(v bool) {
	o.SelfReviewAllowed = v
}

// GetUarScope returns the UarScope field value if set, zero value otherwise.
func (o *UAR) GetUarScope() UARScope {
	if o == nil || o.UarScope == nil {
		var ret UARScope
		return ret
	}
	return *o.UarScope
}

// GetUarScopeOk returns a tuple with the UarScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAR) GetUarScopeOk() (*UARScope, bool) {
	if o == nil || o.UarScope == nil {
		return nil, false
	}
	return o.UarScope, true
}

// HasUarScope returns a boolean if a field has been set.
func (o *UAR) HasUarScope() bool {
	if o != nil && o.UarScope != nil {
		return true
	}

	return false
}

// SetUarScope gets a reference to the given UARScope and assigns it to the UarScope field.
func (o *UAR) SetUarScope(v UARScope) {
	o.UarScope = &v
}

func (o UAR) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uar_id"] = o.UarId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["reviewer_assignment_policy"] = o.ReviewerAssignmentPolicy
	}
	if true {
		toSerialize["send_reviewer_assignment_notification"] = o.SendReviewerAssignmentNotification
	}
	if true {
		toSerialize["deadline"] = o.Deadline
	}
	if true {
		toSerialize["time_zone"] = o.TimeZone
	}
	if true {
		toSerialize["self_review_allowed"] = o.SelfReviewAllowed
	}
	if o.UarScope != nil {
		toSerialize["uar_scope"] = o.UarScope
	}
	return json.Marshal(toSerialize)
}

type NullableUAR struct {
	value *UAR
	isSet bool
}

func (v NullableUAR) Get() *UAR {
	return v.value
}

func (v *NullableUAR) Set(val *UAR) {
	v.value = val
	v.isSet = true
}

func (v NullableUAR) IsSet() bool {
	return v.isSet
}

func (v *NullableUAR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUAR(val *UAR) *NullableUAR {
	return &NullableUAR{value: val, isSet: true}
}

func (v NullableUAR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUAR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

