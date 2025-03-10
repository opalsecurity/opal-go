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

// checks if the GroupResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupResource{}

// GroupResource # GroupResource Object ### Description The `GroupResource` object is used to represent a relationship between a group and a resource.
type GroupResource struct {
	// The ID of the group.
	GroupId string `json:"group_id"`
	// The ID of the resource.
	ResourceId string `json:"resource_id"`
	AccessLevel ResourceAccessLevel `json:"access_level"`
}

type _GroupResource GroupResource

// NewGroupResource instantiates a new GroupResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupResource(groupId string, resourceId string, accessLevel ResourceAccessLevel) *GroupResource {
	this := GroupResource{}
	this.GroupId = groupId
	this.ResourceId = resourceId
	this.AccessLevel = accessLevel
	return &this
}

// NewGroupResourceWithDefaults instantiates a new GroupResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupResourceWithDefaults() *GroupResource {
	this := GroupResource{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *GroupResource) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *GroupResource) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *GroupResource) SetGroupId(v string) {
	o.GroupId = v
}

// GetResourceId returns the ResourceId field value
func (o *GroupResource) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *GroupResource) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *GroupResource) SetResourceId(v string) {
	o.ResourceId = v
}

// GetAccessLevel returns the AccessLevel field value
func (o *GroupResource) GetAccessLevel() ResourceAccessLevel {
	if o == nil {
		var ret ResourceAccessLevel
		return ret
	}

	return o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value
// and a boolean to check if the value has been set.
func (o *GroupResource) GetAccessLevelOk() (*ResourceAccessLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevel, true
}

// SetAccessLevel sets field value
func (o *GroupResource) SetAccessLevel(v ResourceAccessLevel) {
	o.AccessLevel = v
}

func (o GroupResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["access_level"] = o.AccessLevel
	return toSerialize, nil
}

func (o *GroupResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group_id",
		"resource_id",
		"access_level",
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

	varGroupResource := _GroupResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupResource)

	if err != nil {
		return err
	}

	*o = GroupResource(varGroupResource)

	return err
}

type NullableGroupResource struct {
	value *GroupResource
	isSet bool
}

func (v NullableGroupResource) Get() *GroupResource {
	return v.value
}

func (v *NullableGroupResource) Set(val *GroupResource) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupResource) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupResource(val *GroupResource) *NullableGroupResource {
	return &NullableGroupResource{value: val, isSet: true}
}

func (v NullableGroupResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


