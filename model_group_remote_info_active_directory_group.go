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

// checks if the GroupRemoteInfoActiveDirectoryGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRemoteInfoActiveDirectoryGroup{}

// GroupRemoteInfoActiveDirectoryGroup Remote info for Active Directory group.
type GroupRemoteInfoActiveDirectoryGroup struct {
	// The id of the Google group.
	GroupId string `json:"group_id"`
}

type _GroupRemoteInfoActiveDirectoryGroup GroupRemoteInfoActiveDirectoryGroup

// NewGroupRemoteInfoActiveDirectoryGroup instantiates a new GroupRemoteInfoActiveDirectoryGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRemoteInfoActiveDirectoryGroup(groupId string) *GroupRemoteInfoActiveDirectoryGroup {
	this := GroupRemoteInfoActiveDirectoryGroup{}
	this.GroupId = groupId
	return &this
}

// NewGroupRemoteInfoActiveDirectoryGroupWithDefaults instantiates a new GroupRemoteInfoActiveDirectoryGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRemoteInfoActiveDirectoryGroupWithDefaults() *GroupRemoteInfoActiveDirectoryGroup {
	this := GroupRemoteInfoActiveDirectoryGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *GroupRemoteInfoActiveDirectoryGroup) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfoActiveDirectoryGroup) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *GroupRemoteInfoActiveDirectoryGroup) SetGroupId(v string) {
	o.GroupId = v
}

func (o GroupRemoteInfoActiveDirectoryGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRemoteInfoActiveDirectoryGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	return toSerialize, nil
}

func (o *GroupRemoteInfoActiveDirectoryGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group_id",
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

	varGroupRemoteInfoActiveDirectoryGroup := _GroupRemoteInfoActiveDirectoryGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupRemoteInfoActiveDirectoryGroup)

	if err != nil {
		return err
	}

	*o = GroupRemoteInfoActiveDirectoryGroup(varGroupRemoteInfoActiveDirectoryGroup)

	return err
}

type NullableGroupRemoteInfoActiveDirectoryGroup struct {
	value *GroupRemoteInfoActiveDirectoryGroup
	isSet bool
}

func (v NullableGroupRemoteInfoActiveDirectoryGroup) Get() *GroupRemoteInfoActiveDirectoryGroup {
	return v.value
}

func (v *NullableGroupRemoteInfoActiveDirectoryGroup) Set(val *GroupRemoteInfoActiveDirectoryGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRemoteInfoActiveDirectoryGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRemoteInfoActiveDirectoryGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRemoteInfoActiveDirectoryGroup(val *GroupRemoteInfoActiveDirectoryGroup) *NullableGroupRemoteInfoActiveDirectoryGroup {
	return &NullableGroupRemoteInfoActiveDirectoryGroup{value: val, isSet: true}
}

func (v NullableGroupRemoteInfoActiveDirectoryGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRemoteInfoActiveDirectoryGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


