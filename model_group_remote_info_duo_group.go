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

// GroupRemoteInfoDuoGroup Remote info for Duo Security group.
type GroupRemoteInfoDuoGroup struct {
	// The id of the Duo Security group.
	GroupId string `json:"group_id"`
}

// NewGroupRemoteInfoDuoGroup instantiates a new GroupRemoteInfoDuoGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRemoteInfoDuoGroup(groupId string) *GroupRemoteInfoDuoGroup {
	this := GroupRemoteInfoDuoGroup{}
	this.GroupId = groupId
	return &this
}

// NewGroupRemoteInfoDuoGroupWithDefaults instantiates a new GroupRemoteInfoDuoGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRemoteInfoDuoGroupWithDefaults() *GroupRemoteInfoDuoGroup {
	this := GroupRemoteInfoDuoGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *GroupRemoteInfoDuoGroup) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfoDuoGroup) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *GroupRemoteInfoDuoGroup) SetGroupId(v string) {
	o.GroupId = v
}

func (o GroupRemoteInfoDuoGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group_id"] = o.GroupId
	}
	return json.Marshal(toSerialize)
}

type NullableGroupRemoteInfoDuoGroup struct {
	value *GroupRemoteInfoDuoGroup
	isSet bool
}

func (v NullableGroupRemoteInfoDuoGroup) Get() *GroupRemoteInfoDuoGroup {
	return v.value
}

func (v *NullableGroupRemoteInfoDuoGroup) Set(val *GroupRemoteInfoDuoGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRemoteInfoDuoGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRemoteInfoDuoGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRemoteInfoDuoGroup(val *GroupRemoteInfoDuoGroup) *NullableGroupRemoteInfoDuoGroup {
	return &NullableGroupRemoteInfoDuoGroup{value: val, isSet: true}
}

func (v NullableGroupRemoteInfoDuoGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRemoteInfoDuoGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

