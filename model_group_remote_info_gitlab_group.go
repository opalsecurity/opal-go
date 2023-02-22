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

// checks if the GroupRemoteInfoGitlabGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRemoteInfoGitlabGroup{}

// GroupRemoteInfoGitlabGroup Remote info for Gitlab group.
type GroupRemoteInfoGitlabGroup struct {
	// The id of the Gitlab group.
	GroupId string `json:"group_id"`
}

// NewGroupRemoteInfoGitlabGroup instantiates a new GroupRemoteInfoGitlabGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRemoteInfoGitlabGroup(groupId string) *GroupRemoteInfoGitlabGroup {
	this := GroupRemoteInfoGitlabGroup{}
	this.GroupId = groupId
	return &this
}

// NewGroupRemoteInfoGitlabGroupWithDefaults instantiates a new GroupRemoteInfoGitlabGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRemoteInfoGitlabGroupWithDefaults() *GroupRemoteInfoGitlabGroup {
	this := GroupRemoteInfoGitlabGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *GroupRemoteInfoGitlabGroup) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfoGitlabGroup) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *GroupRemoteInfoGitlabGroup) SetGroupId(v string) {
	o.GroupId = v
}

func (o GroupRemoteInfoGitlabGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRemoteInfoGitlabGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	return toSerialize, nil
}

type NullableGroupRemoteInfoGitlabGroup struct {
	value *GroupRemoteInfoGitlabGroup
	isSet bool
}

func (v NullableGroupRemoteInfoGitlabGroup) Get() *GroupRemoteInfoGitlabGroup {
	return v.value
}

func (v *NullableGroupRemoteInfoGitlabGroup) Set(val *GroupRemoteInfoGitlabGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRemoteInfoGitlabGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRemoteInfoGitlabGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRemoteInfoGitlabGroup(val *GroupRemoteInfoGitlabGroup) *NullableGroupRemoteInfoGitlabGroup {
	return &NullableGroupRemoteInfoGitlabGroup{value: val, isSet: true}
}

func (v NullableGroupRemoteInfoGitlabGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRemoteInfoGitlabGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


