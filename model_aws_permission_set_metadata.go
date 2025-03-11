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
	"bytes"
	"fmt"
)

// checks if the AwsPermissionSetMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsPermissionSetMetadata{}

// AwsPermissionSetMetadata Metadata for AWS Identity Center permission set.
type AwsPermissionSetMetadata struct {
	AwsPermissionSet AwsPermissionSetMetadataAwsPermissionSet `json:"aws_permission_set"`
}

type _AwsPermissionSetMetadata AwsPermissionSetMetadata

// NewAwsPermissionSetMetadata instantiates a new AwsPermissionSetMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsPermissionSetMetadata(awsPermissionSet AwsPermissionSetMetadataAwsPermissionSet) *AwsPermissionSetMetadata {
	this := AwsPermissionSetMetadata{}
	this.AwsPermissionSet = awsPermissionSet
	return &this
}

// NewAwsPermissionSetMetadataWithDefaults instantiates a new AwsPermissionSetMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsPermissionSetMetadataWithDefaults() *AwsPermissionSetMetadata {
	this := AwsPermissionSetMetadata{}
	return &this
}

// GetAwsPermissionSet returns the AwsPermissionSet field value
func (o *AwsPermissionSetMetadata) GetAwsPermissionSet() AwsPermissionSetMetadataAwsPermissionSet {
	if o == nil {
		var ret AwsPermissionSetMetadataAwsPermissionSet
		return ret
	}

	return o.AwsPermissionSet
}

// GetAwsPermissionSetOk returns a tuple with the AwsPermissionSet field value
// and a boolean to check if the value has been set.
func (o *AwsPermissionSetMetadata) GetAwsPermissionSetOk() (*AwsPermissionSetMetadataAwsPermissionSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsPermissionSet, true
}

// SetAwsPermissionSet sets field value
func (o *AwsPermissionSetMetadata) SetAwsPermissionSet(v AwsPermissionSetMetadataAwsPermissionSet) {
	o.AwsPermissionSet = v
}

func (o AwsPermissionSetMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsPermissionSetMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws_permission_set"] = o.AwsPermissionSet
	return toSerialize, nil
}

func (o *AwsPermissionSetMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aws_permission_set",
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

	varAwsPermissionSetMetadata := _AwsPermissionSetMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAwsPermissionSetMetadata)

	if err != nil {
		return err
	}

	*o = AwsPermissionSetMetadata(varAwsPermissionSetMetadata)

	return err
}

type NullableAwsPermissionSetMetadata struct {
	value *AwsPermissionSetMetadata
	isSet bool
}

func (v NullableAwsPermissionSetMetadata) Get() *AwsPermissionSetMetadata {
	return v.value
}

func (v *NullableAwsPermissionSetMetadata) Set(val *AwsPermissionSetMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsPermissionSetMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsPermissionSetMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsPermissionSetMetadata(val *AwsPermissionSetMetadata) *NullableAwsPermissionSetMetadata {
	return &NullableAwsPermissionSetMetadata{value: val, isSet: true}
}

func (v NullableAwsPermissionSetMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsPermissionSetMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


