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

// GroupRemoteInfoGoogleGroup Remote info for Google group.
type GroupRemoteInfoGoogleGroup struct {
	// The id of the Google group.
	GroupId string `json:"group_id"`
}

// NewGroupRemoteInfoGoogleGroup instantiates a new GroupRemoteInfoGoogleGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRemoteInfoGoogleGroup(groupId string) *GroupRemoteInfoGoogleGroup {
	this := GroupRemoteInfoGoogleGroup{}
	this.GroupId = groupId
	return &this
}

// NewGroupRemoteInfoGoogleGroupWithDefaults instantiates a new GroupRemoteInfoGoogleGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRemoteInfoGoogleGroupWithDefaults() *GroupRemoteInfoGoogleGroup {
	this := GroupRemoteInfoGoogleGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *GroupRemoteInfoGoogleGroup) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfoGoogleGroup) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *GroupRemoteInfoGoogleGroup) SetGroupId(v string) {
	o.GroupId = v
}

func (o GroupRemoteInfoGoogleGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group_id"] = o.GroupId
	}
	return json.Marshal(toSerialize)
}

type NullableGroupRemoteInfoGoogleGroup struct {
	value *GroupRemoteInfoGoogleGroup
	isSet bool
}

func (v NullableGroupRemoteInfoGoogleGroup) Get() *GroupRemoteInfoGoogleGroup {
	return v.value
}

func (v *NullableGroupRemoteInfoGoogleGroup) Set(val *GroupRemoteInfoGoogleGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRemoteInfoGoogleGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRemoteInfoGoogleGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRemoteInfoGoogleGroup(val *GroupRemoteInfoGoogleGroup) *NullableGroupRemoteInfoGoogleGroup {
	return &NullableGroupRemoteInfoGoogleGroup{value: val, isSet: true}
}

func (v NullableGroupRemoteInfoGoogleGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRemoteInfoGoogleGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

