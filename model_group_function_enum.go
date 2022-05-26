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

// GroupFunctionEnum The function type of the group.
type GroupFunctionEnum string

// List of GroupFunctionEnum
const (
	GROUPFUNCTIONENUM_REGULAR GroupFunctionEnum = "REGULAR"
	GROUPFUNCTIONENUM_ON_CALL GroupFunctionEnum = "ON_CALL"
	GROUPFUNCTIONENUM_TEAM GroupFunctionEnum = "TEAM"
	GROUPFUNCTIONENUM_UNKNOWN GroupFunctionEnum = "UNKNOWN"
)

// All allowed values of GroupFunctionEnum enum
var AllowedGroupFunctionEnumEnumValues = []GroupFunctionEnum{
	"REGULAR",
	"ON_CALL",
	"TEAM",
	"UNKNOWN",
}

func (v *GroupFunctionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupFunctionEnum(value)
	for _, existing := range AllowedGroupFunctionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupFunctionEnum", value)
}

// NewGroupFunctionEnumFromValue returns a pointer to a valid GroupFunctionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupFunctionEnumFromValue(v string) (*GroupFunctionEnum, error) {
	ev := GroupFunctionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupFunctionEnum: valid values are %v", v, AllowedGroupFunctionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupFunctionEnum) IsValid() bool {
	for _, existing := range AllowedGroupFunctionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupFunctionEnum value
func (v GroupFunctionEnum) Ptr() *GroupFunctionEnum {
	return &v
}

type NullableGroupFunctionEnum struct {
	value *GroupFunctionEnum
	isSet bool
}

func (v NullableGroupFunctionEnum) Get() *GroupFunctionEnum {
	return v.value
}

func (v *NullableGroupFunctionEnum) Set(val *GroupFunctionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupFunctionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupFunctionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupFunctionEnum(val *GroupFunctionEnum) *NullableGroupFunctionEnum {
	return &NullableGroupFunctionEnum{value: val, isSet: true}
}

func (v NullableGroupFunctionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupFunctionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

