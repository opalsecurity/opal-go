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

// UpdateGroupResourcesInfo struct for UpdateGroupResourcesInfo
type UpdateGroupResourcesInfo struct {
	Resources []ResourceWithAccessLevel `json:"resources"`
}

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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
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


