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
	"gopkg.in/validator.v2"
)

// RequestCustomFieldResponseFieldValue - struct for RequestCustomFieldResponseFieldValue
type RequestCustomFieldResponseFieldValue struct {
	Bool *bool
	String *string
}

// boolAsRequestCustomFieldResponseFieldValue is a convenience function that returns bool wrapped in RequestCustomFieldResponseFieldValue
func BoolAsRequestCustomFieldResponseFieldValue(v *bool) RequestCustomFieldResponseFieldValue {
	return RequestCustomFieldResponseFieldValue{
		Bool: v,
	}
}

// stringAsRequestCustomFieldResponseFieldValue is a convenience function that returns string wrapped in RequestCustomFieldResponseFieldValue
func StringAsRequestCustomFieldResponseFieldValue(v *string) RequestCustomFieldResponseFieldValue {
	return RequestCustomFieldResponseFieldValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestCustomFieldResponseFieldValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RequestCustomFieldResponseFieldValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestCustomFieldResponseFieldValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestCustomFieldResponseFieldValue) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestCustomFieldResponseFieldValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RequestCustomFieldResponseFieldValue) GetActualInstanceValue() (interface{}) {
	if obj.Bool != nil {
		return *obj.Bool
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableRequestCustomFieldResponseFieldValue struct {
	value *RequestCustomFieldResponseFieldValue
	isSet bool
}

func (v NullableRequestCustomFieldResponseFieldValue) Get() *RequestCustomFieldResponseFieldValue {
	return v.value
}

func (v *NullableRequestCustomFieldResponseFieldValue) Set(val *RequestCustomFieldResponseFieldValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestCustomFieldResponseFieldValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestCustomFieldResponseFieldValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestCustomFieldResponseFieldValue(val *RequestCustomFieldResponseFieldValue) *NullableRequestCustomFieldResponseFieldValue {
	return &NullableRequestCustomFieldResponseFieldValue{value: val, isSet: true}
}

func (v NullableRequestCustomFieldResponseFieldValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestCustomFieldResponseFieldValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


