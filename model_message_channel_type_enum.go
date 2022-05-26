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
	"fmt"
)

// MessageChannelTypeEnum The type of the message channel.
type MessageChannelTypeEnum string

// List of MessageChannelTypeEnum
const (
	MESSAGECHANNELTYPEENUM_AUDIT MessageChannelTypeEnum = "AUDIT"
	MESSAGECHANNELTYPEENUM_REVIEWER MessageChannelTypeEnum = "REVIEWER"
	MESSAGECHANNELTYPEENUM_MONITOR MessageChannelTypeEnum = "MONITOR"
)

// All allowed values of MessageChannelTypeEnum enum
var AllowedMessageChannelTypeEnumEnumValues = []MessageChannelTypeEnum{
	"AUDIT",
	"REVIEWER",
	"MONITOR",
}

func (v *MessageChannelTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageChannelTypeEnum(value)
	for _, existing := range AllowedMessageChannelTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageChannelTypeEnum", value)
}

// NewMessageChannelTypeEnumFromValue returns a pointer to a valid MessageChannelTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageChannelTypeEnumFromValue(v string) (*MessageChannelTypeEnum, error) {
	ev := MessageChannelTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageChannelTypeEnum: valid values are %v", v, AllowedMessageChannelTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageChannelTypeEnum) IsValid() bool {
	for _, existing := range AllowedMessageChannelTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageChannelTypeEnum value
func (v MessageChannelTypeEnum) Ptr() *MessageChannelTypeEnum {
	return &v
}

type NullableMessageChannelTypeEnum struct {
	value *MessageChannelTypeEnum
	isSet bool
}

func (v NullableMessageChannelTypeEnum) Get() *MessageChannelTypeEnum {
	return v.value
}

func (v *NullableMessageChannelTypeEnum) Set(val *MessageChannelTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageChannelTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageChannelTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageChannelTypeEnum(val *MessageChannelTypeEnum) *NullableMessageChannelTypeEnum {
	return &NullableMessageChannelTypeEnum{value: val, isSet: true}
}

func (v NullableMessageChannelTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageChannelTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

