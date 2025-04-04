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

// checks if the ResourceRemoteInfoOktaCustomRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoOktaCustomRole{}

// ResourceRemoteInfoOktaCustomRole Remote info for Okta directory custom role.
type ResourceRemoteInfoOktaCustomRole struct {
	// The id of the custom role.
	RoleId string `json:"role_id"`
}

type _ResourceRemoteInfoOktaCustomRole ResourceRemoteInfoOktaCustomRole

// NewResourceRemoteInfoOktaCustomRole instantiates a new ResourceRemoteInfoOktaCustomRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoOktaCustomRole(roleId string) *ResourceRemoteInfoOktaCustomRole {
	this := ResourceRemoteInfoOktaCustomRole{}
	this.RoleId = roleId
	return &this
}

// NewResourceRemoteInfoOktaCustomRoleWithDefaults instantiates a new ResourceRemoteInfoOktaCustomRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoOktaCustomRoleWithDefaults() *ResourceRemoteInfoOktaCustomRole {
	this := ResourceRemoteInfoOktaCustomRole{}
	return &this
}

// GetRoleId returns the RoleId field value
func (o *ResourceRemoteInfoOktaCustomRole) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoOktaCustomRole) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *ResourceRemoteInfoOktaCustomRole) SetRoleId(v string) {
	o.RoleId = v
}

func (o ResourceRemoteInfoOktaCustomRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoOktaCustomRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role_id"] = o.RoleId
	return toSerialize, nil
}

func (o *ResourceRemoteInfoOktaCustomRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role_id",
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

	varResourceRemoteInfoOktaCustomRole := _ResourceRemoteInfoOktaCustomRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceRemoteInfoOktaCustomRole)

	if err != nil {
		return err
	}

	*o = ResourceRemoteInfoOktaCustomRole(varResourceRemoteInfoOktaCustomRole)

	return err
}

type NullableResourceRemoteInfoOktaCustomRole struct {
	value *ResourceRemoteInfoOktaCustomRole
	isSet bool
}

func (v NullableResourceRemoteInfoOktaCustomRole) Get() *ResourceRemoteInfoOktaCustomRole {
	return v.value
}

func (v *NullableResourceRemoteInfoOktaCustomRole) Set(val *ResourceRemoteInfoOktaCustomRole) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoOktaCustomRole) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoOktaCustomRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoOktaCustomRole(val *ResourceRemoteInfoOktaCustomRole) *NullableResourceRemoteInfoOktaCustomRole {
	return &NullableResourceRemoteInfoOktaCustomRole{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoOktaCustomRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoOktaCustomRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


