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

// checks if the GroupRemoteInfoOktaGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRemoteInfoOktaGroup{}

// GroupRemoteInfoOktaGroup Remote info for Okta Directory group.
type GroupRemoteInfoOktaGroup struct {
	// The id of the Okta Directory group.
	GroupId string `json:"group_id"`
}

// NewGroupRemoteInfoOktaGroup instantiates a new GroupRemoteInfoOktaGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRemoteInfoOktaGroup(groupId string) *GroupRemoteInfoOktaGroup {
	this := GroupRemoteInfoOktaGroup{}
	this.GroupId = groupId
	return &this
}

// NewGroupRemoteInfoOktaGroupWithDefaults instantiates a new GroupRemoteInfoOktaGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRemoteInfoOktaGroupWithDefaults() *GroupRemoteInfoOktaGroup {
	this := GroupRemoteInfoOktaGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *GroupRemoteInfoOktaGroup) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfoOktaGroup) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *GroupRemoteInfoOktaGroup) SetGroupId(v string) {
	o.GroupId = v
}

func (o GroupRemoteInfoOktaGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRemoteInfoOktaGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	return toSerialize, nil
}

type NullableGroupRemoteInfoOktaGroup struct {
	value *GroupRemoteInfoOktaGroup
	isSet bool
}

func (v NullableGroupRemoteInfoOktaGroup) Get() *GroupRemoteInfoOktaGroup {
	return v.value
}

func (v *NullableGroupRemoteInfoOktaGroup) Set(val *GroupRemoteInfoOktaGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRemoteInfoOktaGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRemoteInfoOktaGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRemoteInfoOktaGroup(val *GroupRemoteInfoOktaGroup) *NullableGroupRemoteInfoOktaGroup {
	return &NullableGroupRemoteInfoOktaGroup{value: val, isSet: true}
}

func (v NullableGroupRemoteInfoOktaGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRemoteInfoOktaGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


