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
)

// OnCallSchedule # OnCallSchedule Object ### Description The `OnCallSchedule` object is used to represent an on call schedule.  ### Usage Example Update a groups on call schedule from the `UPDATE Groups` endpoint.
type OnCallSchedule struct {
	// The ID of the on-call schedule.
	OnCallScheduleId *string `json:"on_call_schedule_id,omitempty"`
	ThirdPartyProvider *OnCallScheduleProviderEnum `json:"third_party_provider,omitempty"`
	// The remote ID of the on call schedule
	RemoteId *string `json:"remote_id,omitempty"`
	// The name of the on call schedule.
	Name *string `json:"name,omitempty"`
}

// NewOnCallSchedule instantiates a new OnCallSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnCallSchedule() *OnCallSchedule {
	this := OnCallSchedule{}
	return &this
}

// NewOnCallScheduleWithDefaults instantiates a new OnCallSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnCallScheduleWithDefaults() *OnCallSchedule {
	this := OnCallSchedule{}
	return &this
}

// GetOnCallScheduleId returns the OnCallScheduleId field value if set, zero value otherwise.
func (o *OnCallSchedule) GetOnCallScheduleId() string {
	if o == nil || isNil(o.OnCallScheduleId) {
		var ret string
		return ret
	}
	return *o.OnCallScheduleId
}

// GetOnCallScheduleIdOk returns a tuple with the OnCallScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallSchedule) GetOnCallScheduleIdOk() (*string, bool) {
	if o == nil || isNil(o.OnCallScheduleId) {
    return nil, false
	}
	return o.OnCallScheduleId, true
}

// HasOnCallScheduleId returns a boolean if a field has been set.
func (o *OnCallSchedule) HasOnCallScheduleId() bool {
	if o != nil && !isNil(o.OnCallScheduleId) {
		return true
	}

	return false
}

// SetOnCallScheduleId gets a reference to the given string and assigns it to the OnCallScheduleId field.
func (o *OnCallSchedule) SetOnCallScheduleId(v string) {
	o.OnCallScheduleId = &v
}

// GetThirdPartyProvider returns the ThirdPartyProvider field value if set, zero value otherwise.
func (o *OnCallSchedule) GetThirdPartyProvider() OnCallScheduleProviderEnum {
	if o == nil || isNil(o.ThirdPartyProvider) {
		var ret OnCallScheduleProviderEnum
		return ret
	}
	return *o.ThirdPartyProvider
}

// GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallSchedule) GetThirdPartyProviderOk() (*OnCallScheduleProviderEnum, bool) {
	if o == nil || isNil(o.ThirdPartyProvider) {
    return nil, false
	}
	return o.ThirdPartyProvider, true
}

// HasThirdPartyProvider returns a boolean if a field has been set.
func (o *OnCallSchedule) HasThirdPartyProvider() bool {
	if o != nil && !isNil(o.ThirdPartyProvider) {
		return true
	}

	return false
}

// SetThirdPartyProvider gets a reference to the given OnCallScheduleProviderEnum and assigns it to the ThirdPartyProvider field.
func (o *OnCallSchedule) SetThirdPartyProvider(v OnCallScheduleProviderEnum) {
	o.ThirdPartyProvider = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *OnCallSchedule) GetRemoteId() string {
	if o == nil || isNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallSchedule) GetRemoteIdOk() (*string, bool) {
	if o == nil || isNil(o.RemoteId) {
    return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *OnCallSchedule) HasRemoteId() bool {
	if o != nil && !isNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *OnCallSchedule) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OnCallSchedule) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallSchedule) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OnCallSchedule) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OnCallSchedule) SetName(v string) {
	o.Name = &v
}

func (o OnCallSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OnCallScheduleId) {
		toSerialize["on_call_schedule_id"] = o.OnCallScheduleId
	}
	if !isNil(o.ThirdPartyProvider) {
		toSerialize["third_party_provider"] = o.ThirdPartyProvider
	}
	if !isNil(o.RemoteId) {
		toSerialize["remote_id"] = o.RemoteId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOnCallSchedule struct {
	value *OnCallSchedule
	isSet bool
}

func (v NullableOnCallSchedule) Get() *OnCallSchedule {
	return v.value
}

func (v *NullableOnCallSchedule) Set(val *OnCallSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableOnCallSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableOnCallSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnCallSchedule(val *OnCallSchedule) *NullableOnCallSchedule {
	return &NullableOnCallSchedule{value: val, isSet: true}
}

func (v NullableOnCallSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnCallSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


