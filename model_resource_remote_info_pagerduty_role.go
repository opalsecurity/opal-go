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

// checks if the ResourceRemoteInfoPagerdutyRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoPagerdutyRole{}

// ResourceRemoteInfoPagerdutyRole Remote info for Pagerduty role.
type ResourceRemoteInfoPagerdutyRole struct {
	// The name of the role.
	RoleName string `json:"role_name"`
}

type _ResourceRemoteInfoPagerdutyRole ResourceRemoteInfoPagerdutyRole

// NewResourceRemoteInfoPagerdutyRole instantiates a new ResourceRemoteInfoPagerdutyRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoPagerdutyRole(roleName string) *ResourceRemoteInfoPagerdutyRole {
	this := ResourceRemoteInfoPagerdutyRole{}
	this.RoleName = roleName
	return &this
}

// NewResourceRemoteInfoPagerdutyRoleWithDefaults instantiates a new ResourceRemoteInfoPagerdutyRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoPagerdutyRoleWithDefaults() *ResourceRemoteInfoPagerdutyRole {
	this := ResourceRemoteInfoPagerdutyRole{}
	return &this
}

// GetRoleName returns the RoleName field value
func (o *ResourceRemoteInfoPagerdutyRole) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoPagerdutyRole) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *ResourceRemoteInfoPagerdutyRole) SetRoleName(v string) {
	o.RoleName = v
}

func (o ResourceRemoteInfoPagerdutyRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoPagerdutyRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role_name"] = o.RoleName
	return toSerialize, nil
}

func (o *ResourceRemoteInfoPagerdutyRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role_name",
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

	varResourceRemoteInfoPagerdutyRole := _ResourceRemoteInfoPagerdutyRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceRemoteInfoPagerdutyRole)

	if err != nil {
		return err
	}

	*o = ResourceRemoteInfoPagerdutyRole(varResourceRemoteInfoPagerdutyRole)

	return err
}

type NullableResourceRemoteInfoPagerdutyRole struct {
	value *ResourceRemoteInfoPagerdutyRole
	isSet bool
}

func (v NullableResourceRemoteInfoPagerdutyRole) Get() *ResourceRemoteInfoPagerdutyRole {
	return v.value
}

func (v *NullableResourceRemoteInfoPagerdutyRole) Set(val *ResourceRemoteInfoPagerdutyRole) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoPagerdutyRole) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoPagerdutyRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoPagerdutyRole(val *ResourceRemoteInfoPagerdutyRole) *NullableResourceRemoteInfoPagerdutyRole {
	return &NullableResourceRemoteInfoPagerdutyRole{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoPagerdutyRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoPagerdutyRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


