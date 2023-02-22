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

// checks if the ResourceWithAccessLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceWithAccessLevel{}

// ResourceWithAccessLevel Information about a resource and corresponding access level
type ResourceWithAccessLevel struct {
	// The ID of the resource.
	ResourceId string `json:"resource_id"`
	// The ID of the resource.
	AccessLevelRemoteId *string `json:"access_level_remote_id,omitempty"`
}

// NewResourceWithAccessLevel instantiates a new ResourceWithAccessLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceWithAccessLevel(resourceId string) *ResourceWithAccessLevel {
	this := ResourceWithAccessLevel{}
	this.ResourceId = resourceId
	return &this
}

// NewResourceWithAccessLevelWithDefaults instantiates a new ResourceWithAccessLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithAccessLevelWithDefaults() *ResourceWithAccessLevel {
	this := ResourceWithAccessLevel{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *ResourceWithAccessLevel) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ResourceWithAccessLevel) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ResourceWithAccessLevel) SetResourceId(v string) {
	o.ResourceId = v
}

// GetAccessLevelRemoteId returns the AccessLevelRemoteId field value if set, zero value otherwise.
func (o *ResourceWithAccessLevel) GetAccessLevelRemoteId() string {
	if o == nil || IsNil(o.AccessLevelRemoteId) {
		var ret string
		return ret
	}
	return *o.AccessLevelRemoteId
}

// GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceWithAccessLevel) GetAccessLevelRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessLevelRemoteId) {
		return nil, false
	}
	return o.AccessLevelRemoteId, true
}

// HasAccessLevelRemoteId returns a boolean if a field has been set.
func (o *ResourceWithAccessLevel) HasAccessLevelRemoteId() bool {
	if o != nil && !IsNil(o.AccessLevelRemoteId) {
		return true
	}

	return false
}

// SetAccessLevelRemoteId gets a reference to the given string and assigns it to the AccessLevelRemoteId field.
func (o *ResourceWithAccessLevel) SetAccessLevelRemoteId(v string) {
	o.AccessLevelRemoteId = &v
}

func (o ResourceWithAccessLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceWithAccessLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_id"] = o.ResourceId
	if !IsNil(o.AccessLevelRemoteId) {
		toSerialize["access_level_remote_id"] = o.AccessLevelRemoteId
	}
	return toSerialize, nil
}

type NullableResourceWithAccessLevel struct {
	value *ResourceWithAccessLevel
	isSet bool
}

func (v NullableResourceWithAccessLevel) Get() *ResourceWithAccessLevel {
	return v.value
}

func (v *NullableResourceWithAccessLevel) Set(val *ResourceWithAccessLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceWithAccessLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceWithAccessLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceWithAccessLevel(val *ResourceWithAccessLevel) *NullableResourceWithAccessLevel {
	return &NullableResourceWithAccessLevel{value: val, isSet: true}
}

func (v NullableResourceWithAccessLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceWithAccessLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


