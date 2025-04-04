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

// checks if the ResourceRemoteInfoOktaStandardRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoOktaStandardRole{}

// ResourceRemoteInfoOktaStandardRole Remote info for Okta directory standard role.
type ResourceRemoteInfoOktaStandardRole struct {
	// The type of the standard role.
	RoleType string `json:"role_type"`
}

type _ResourceRemoteInfoOktaStandardRole ResourceRemoteInfoOktaStandardRole

// NewResourceRemoteInfoOktaStandardRole instantiates a new ResourceRemoteInfoOktaStandardRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoOktaStandardRole(roleType string) *ResourceRemoteInfoOktaStandardRole {
	this := ResourceRemoteInfoOktaStandardRole{}
	this.RoleType = roleType
	return &this
}

// NewResourceRemoteInfoOktaStandardRoleWithDefaults instantiates a new ResourceRemoteInfoOktaStandardRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoOktaStandardRoleWithDefaults() *ResourceRemoteInfoOktaStandardRole {
	this := ResourceRemoteInfoOktaStandardRole{}
	return &this
}

// GetRoleType returns the RoleType field value
func (o *ResourceRemoteInfoOktaStandardRole) GetRoleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoOktaStandardRole) GetRoleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleType, true
}

// SetRoleType sets field value
func (o *ResourceRemoteInfoOktaStandardRole) SetRoleType(v string) {
	o.RoleType = v
}

func (o ResourceRemoteInfoOktaStandardRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoOktaStandardRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role_type"] = o.RoleType
	return toSerialize, nil
}

func (o *ResourceRemoteInfoOktaStandardRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role_type",
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

	varResourceRemoteInfoOktaStandardRole := _ResourceRemoteInfoOktaStandardRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceRemoteInfoOktaStandardRole)

	if err != nil {
		return err
	}

	*o = ResourceRemoteInfoOktaStandardRole(varResourceRemoteInfoOktaStandardRole)

	return err
}

type NullableResourceRemoteInfoOktaStandardRole struct {
	value *ResourceRemoteInfoOktaStandardRole
	isSet bool
}

func (v NullableResourceRemoteInfoOktaStandardRole) Get() *ResourceRemoteInfoOktaStandardRole {
	return v.value
}

func (v *NullableResourceRemoteInfoOktaStandardRole) Set(val *ResourceRemoteInfoOktaStandardRole) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoOktaStandardRole) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoOktaStandardRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoOktaStandardRole(val *ResourceRemoteInfoOktaStandardRole) *NullableResourceRemoteInfoOktaStandardRole {
	return &NullableResourceRemoteInfoOktaStandardRole{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoOktaStandardRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoOktaStandardRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


