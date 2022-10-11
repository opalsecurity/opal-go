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

// SubEvent # Sub event Object ### Description The `SubEvent` object is used to represent a subevent.  ### Usage Example Fetch from the `LIST Events` endpoint.
type SubEvent struct {
	// The subevent type.
	SubEventType string `json:"sub_event_type"`
	AdditionalProperties map[string]interface{}
}

type _SubEvent SubEvent

// NewSubEvent instantiates a new SubEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubEvent(subEventType string) *SubEvent {
	this := SubEvent{}
	this.SubEventType = subEventType
	return &this
}

// NewSubEventWithDefaults instantiates a new SubEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubEventWithDefaults() *SubEvent {
	this := SubEvent{}
	return &this
}

// GetSubEventType returns the SubEventType field value
func (o *SubEvent) GetSubEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubEventType
}

// GetSubEventTypeOk returns a tuple with the SubEventType field value
// and a boolean to check if the value has been set.
func (o *SubEvent) GetSubEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubEventType, true
}

// SetSubEventType sets field value
func (o *SubEvent) SetSubEventType(v string) {
	o.SubEventType = v
}

func (o SubEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sub_event_type"] = o.SubEventType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SubEvent) UnmarshalJSON(bytes []byte) (err error) {
	varSubEvent := _SubEvent{}

	if err = json.Unmarshal(bytes, &varSubEvent); err == nil {
		*o = SubEvent(varSubEvent)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sub_event_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubEvent struct {
	value *SubEvent
	isSet bool
}

func (v NullableSubEvent) Get() *SubEvent {
	return v.value
}

func (v *NullableSubEvent) Set(val *SubEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSubEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSubEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubEvent(val *SubEvent) *NullableSubEvent {
	return &NullableSubEvent{value: val, isSet: true}
}

func (v NullableSubEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


