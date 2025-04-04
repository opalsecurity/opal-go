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

// BundlesVisibilityTypeEnum The visibility level of the entity.
type BundlesVisibilityTypeEnum string

// List of BundlesVisibilityTypeEnum
const (
	BUNDLESVISIBILITYTYPEENUM_GLOBAL BundlesVisibilityTypeEnum = "GLOBAL"
	BUNDLESVISIBILITYTYPEENUM_TEAM BundlesVisibilityTypeEnum = "TEAM"
)

// All allowed values of BundlesVisibilityTypeEnum enum
var AllowedBundlesVisibilityTypeEnumEnumValues = []BundlesVisibilityTypeEnum{
	"GLOBAL",
	"TEAM",
}

func (v *BundlesVisibilityTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BundlesVisibilityTypeEnum(value)
	for _, existing := range AllowedBundlesVisibilityTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BundlesVisibilityTypeEnum", value)
}

// NewBundlesVisibilityTypeEnumFromValue returns a pointer to a valid BundlesVisibilityTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBundlesVisibilityTypeEnumFromValue(v string) (*BundlesVisibilityTypeEnum, error) {
	ev := BundlesVisibilityTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BundlesVisibilityTypeEnum: valid values are %v", v, AllowedBundlesVisibilityTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BundlesVisibilityTypeEnum) IsValid() bool {
	for _, existing := range AllowedBundlesVisibilityTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BundlesVisibilityTypeEnum value
func (v BundlesVisibilityTypeEnum) Ptr() *BundlesVisibilityTypeEnum {
	return &v
}

type NullableBundlesVisibilityTypeEnum struct {
	value *BundlesVisibilityTypeEnum
	isSet bool
}

func (v NullableBundlesVisibilityTypeEnum) Get() *BundlesVisibilityTypeEnum {
	return v.value
}

func (v *NullableBundlesVisibilityTypeEnum) Set(val *BundlesVisibilityTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBundlesVisibilityTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBundlesVisibilityTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundlesVisibilityTypeEnum(val *BundlesVisibilityTypeEnum) *NullableBundlesVisibilityTypeEnum {
	return &NullableBundlesVisibilityTypeEnum{value: val, isSet: true}
}

func (v NullableBundlesVisibilityTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundlesVisibilityTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

