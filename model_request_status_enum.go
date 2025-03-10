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

// RequestStatusEnum # Request Status ### Description The `RequestStatus` enum is used to represent the status of a request.  ### Usage Example Returned from the `GET Requests` endpoint.
type RequestStatusEnum string

// List of RequestStatusEnum
const (
	REQUESTSTATUSENUM_PENDING RequestStatusEnum = "pending"
	REQUESTSTATUSENUM_APPROVED RequestStatusEnum = "approved"
	REQUESTSTATUSENUM_DENIED RequestStatusEnum = "denied"
	REQUESTSTATUSENUM_CANCELED RequestStatusEnum = "canceled"
)

// All allowed values of RequestStatusEnum enum
var AllowedRequestStatusEnumEnumValues = []RequestStatusEnum{
	"pending",
	"approved",
	"denied",
	"canceled",
}

func (v *RequestStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestStatusEnum(value)
	for _, existing := range AllowedRequestStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestStatusEnum", value)
}

// NewRequestStatusEnumFromValue returns a pointer to a valid RequestStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestStatusEnumFromValue(v string) (*RequestStatusEnum, error) {
	ev := RequestStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestStatusEnum: valid values are %v", v, AllowedRequestStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestStatusEnum) IsValid() bool {
	for _, existing := range AllowedRequestStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestStatusEnum value
func (v RequestStatusEnum) Ptr() *RequestStatusEnum {
	return &v
}

type NullableRequestStatusEnum struct {
	value *RequestStatusEnum
	isSet bool
}

func (v NullableRequestStatusEnum) Get() *RequestStatusEnum {
	return v.value
}

func (v *NullableRequestStatusEnum) Set(val *RequestStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestStatusEnum(val *RequestStatusEnum) *NullableRequestStatusEnum {
	return &NullableRequestStatusEnum{value: val, isSet: true}
}

func (v NullableRequestStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

