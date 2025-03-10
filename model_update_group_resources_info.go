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

// checks if the UpdateGroupResourcesInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGroupResourcesInfo{}

// UpdateGroupResourcesInfo struct for UpdateGroupResourcesInfo
type UpdateGroupResourcesInfo struct {
	Resources []ResourceWithAccessLevel `json:"resources"`
}

type _UpdateGroupResourcesInfo UpdateGroupResourcesInfo

// NewUpdateGroupResourcesInfo instantiates a new UpdateGroupResourcesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGroupResourcesInfo(resources []ResourceWithAccessLevel) *UpdateGroupResourcesInfo {
	this := UpdateGroupResourcesInfo{}
	this.Resources = resources
	return &this
}

// NewUpdateGroupResourcesInfoWithDefaults instantiates a new UpdateGroupResourcesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGroupResourcesInfoWithDefaults() *UpdateGroupResourcesInfo {
	this := UpdateGroupResourcesInfo{}
	return &this
}

// GetResources returns the Resources field value
func (o *UpdateGroupResourcesInfo) GetResources() []ResourceWithAccessLevel {
	if o == nil {
		var ret []ResourceWithAccessLevel
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupResourcesInfo) GetResourcesOk() ([]ResourceWithAccessLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *UpdateGroupResourcesInfo) SetResources(v []ResourceWithAccessLevel) {
	o.Resources = v
}

func (o UpdateGroupResourcesInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGroupResourcesInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resources"] = o.Resources
	return toSerialize, nil
}

func (o *UpdateGroupResourcesInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resources",
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

	varUpdateGroupResourcesInfo := _UpdateGroupResourcesInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateGroupResourcesInfo)

	if err != nil {
		return err
	}

	*o = UpdateGroupResourcesInfo(varUpdateGroupResourcesInfo)

	return err
}

type NullableUpdateGroupResourcesInfo struct {
	value *UpdateGroupResourcesInfo
	isSet bool
}

func (v NullableUpdateGroupResourcesInfo) Get() *UpdateGroupResourcesInfo {
	return v.value
}

func (v *NullableUpdateGroupResourcesInfo) Set(val *UpdateGroupResourcesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupResourcesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupResourcesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupResourcesInfo(val *UpdateGroupResourcesInfo) *NullableUpdateGroupResourcesInfo {
	return &NullableUpdateGroupResourcesInfo{value: val, isSet: true}
}

func (v NullableUpdateGroupResourcesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupResourcesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


