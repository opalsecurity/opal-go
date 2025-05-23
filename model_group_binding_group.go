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

// checks if the GroupBindingGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupBindingGroup{}

// GroupBindingGroup # Group Binding Group Object ### Description The `GroupBindingGroup` object is used to represent a group binding group.  ### Usage Example Get group binding groups from the `GET Group Bindings` endpoint.
type GroupBindingGroup struct {
	// The ID of the group.
	GroupId string `json:"group_id"`
	GroupType GroupTypeEnum `json:"group_type"`
}

type _GroupBindingGroup GroupBindingGroup

// NewGroupBindingGroup instantiates a new GroupBindingGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupBindingGroup(groupId string, groupType GroupTypeEnum) *GroupBindingGroup {
	this := GroupBindingGroup{}
	this.GroupId = groupId
	this.GroupType = groupType
	return &this
}

// NewGroupBindingGroupWithDefaults instantiates a new GroupBindingGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupBindingGroupWithDefaults() *GroupBindingGroup {
	this := GroupBindingGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *GroupBindingGroup) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *GroupBindingGroup) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *GroupBindingGroup) SetGroupId(v string) {
	o.GroupId = v
}

// GetGroupType returns the GroupType field value
func (o *GroupBindingGroup) GetGroupType() GroupTypeEnum {
	if o == nil {
		var ret GroupTypeEnum
		return ret
	}

	return o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value
// and a boolean to check if the value has been set.
func (o *GroupBindingGroup) GetGroupTypeOk() (*GroupTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupType, true
}

// SetGroupType sets field value
func (o *GroupBindingGroup) SetGroupType(v GroupTypeEnum) {
	o.GroupType = v
}

func (o GroupBindingGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupBindingGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	toSerialize["group_type"] = o.GroupType
	return toSerialize, nil
}

func (o *GroupBindingGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group_id",
		"group_type",
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

	varGroupBindingGroup := _GroupBindingGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupBindingGroup)

	if err != nil {
		return err
	}

	*o = GroupBindingGroup(varGroupBindingGroup)

	return err
}

type NullableGroupBindingGroup struct {
	value *GroupBindingGroup
	isSet bool
}

func (v NullableGroupBindingGroup) Get() *GroupBindingGroup {
	return v.value
}

func (v *NullableGroupBindingGroup) Set(val *GroupBindingGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupBindingGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupBindingGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupBindingGroup(val *GroupBindingGroup) *NullableGroupBindingGroup {
	return &NullableGroupBindingGroup{value: val, isSet: true}
}

func (v NullableGroupBindingGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupBindingGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


