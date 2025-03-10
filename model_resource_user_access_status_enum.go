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
	"fmt"
)

// ResourceUserAccessStatusEnum The status of the user's access to the resource.
type ResourceUserAccessStatusEnum string

// List of ResourceUserAccessStatusEnum
const (
	RESOURCEUSERACCESSSTATUSENUM_AUTHORIZED ResourceUserAccessStatusEnum = "AUTHORIZED"
	RESOURCEUSERACCESSSTATUSENUM_REQUESTED ResourceUserAccessStatusEnum = "REQUESTED"
	RESOURCEUSERACCESSSTATUSENUM_UNAUTHORIZED ResourceUserAccessStatusEnum = "UNAUTHORIZED"
)

// All allowed values of ResourceUserAccessStatusEnum enum
var AllowedResourceUserAccessStatusEnumEnumValues = []ResourceUserAccessStatusEnum{
	"AUTHORIZED",
	"REQUESTED",
	"UNAUTHORIZED",
}

func (v *ResourceUserAccessStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceUserAccessStatusEnum(value)
	for _, existing := range AllowedResourceUserAccessStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceUserAccessStatusEnum", value)
}

// NewResourceUserAccessStatusEnumFromValue returns a pointer to a valid ResourceUserAccessStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceUserAccessStatusEnumFromValue(v string) (*ResourceUserAccessStatusEnum, error) {
	ev := ResourceUserAccessStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceUserAccessStatusEnum: valid values are %v", v, AllowedResourceUserAccessStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceUserAccessStatusEnum) IsValid() bool {
	for _, existing := range AllowedResourceUserAccessStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResourceUserAccessStatusEnum value
func (v ResourceUserAccessStatusEnum) Ptr() *ResourceUserAccessStatusEnum {
	return &v
}

type NullableResourceUserAccessStatusEnum struct {
	value *ResourceUserAccessStatusEnum
	isSet bool
}

func (v NullableResourceUserAccessStatusEnum) Get() *ResourceUserAccessStatusEnum {
	return v.value
}

func (v *NullableResourceUserAccessStatusEnum) Set(val *ResourceUserAccessStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceUserAccessStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceUserAccessStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceUserAccessStatusEnum(val *ResourceUserAccessStatusEnum) *NullableResourceUserAccessStatusEnum {
	return &NullableResourceUserAccessStatusEnum{value: val, isSet: true}
}

func (v NullableResourceUserAccessStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceUserAccessStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

