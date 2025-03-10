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

// checks if the ResourceAccessLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAccessLevel{}

// ResourceAccessLevel # Access Level Object ### Description The `AccessLevel` object is used to represent the level of access that a principal has. The \"default\" access level is a `AccessLevel` object whose fields are all empty strings.  ### Usage Example View the `AccessLevel` of a resource/user or resource/group pair to see the level of access granted to the resource.
type ResourceAccessLevel struct {
	// The human-readable name of the access level.
	AccessLevelName string `json:"access_level_name"`
	// The machine-readable identifier of the access level.
	AccessLevelRemoteId string `json:"access_level_remote_id"`
}

type _ResourceAccessLevel ResourceAccessLevel

// NewResourceAccessLevel instantiates a new ResourceAccessLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAccessLevel(accessLevelName string, accessLevelRemoteId string) *ResourceAccessLevel {
	this := ResourceAccessLevel{}
	this.AccessLevelName = accessLevelName
	this.AccessLevelRemoteId = accessLevelRemoteId
	return &this
}

// NewResourceAccessLevelWithDefaults instantiates a new ResourceAccessLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAccessLevelWithDefaults() *ResourceAccessLevel {
	this := ResourceAccessLevel{}
	return &this
}

// GetAccessLevelName returns the AccessLevelName field value
func (o *ResourceAccessLevel) GetAccessLevelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessLevelName
}

// GetAccessLevelNameOk returns a tuple with the AccessLevelName field value
// and a boolean to check if the value has been set.
func (o *ResourceAccessLevel) GetAccessLevelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevelName, true
}

// SetAccessLevelName sets field value
func (o *ResourceAccessLevel) SetAccessLevelName(v string) {
	o.AccessLevelName = v
}

// GetAccessLevelRemoteId returns the AccessLevelRemoteId field value
func (o *ResourceAccessLevel) GetAccessLevelRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessLevelRemoteId
}

// GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field value
// and a boolean to check if the value has been set.
func (o *ResourceAccessLevel) GetAccessLevelRemoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevelRemoteId, true
}

// SetAccessLevelRemoteId sets field value
func (o *ResourceAccessLevel) SetAccessLevelRemoteId(v string) {
	o.AccessLevelRemoteId = v
}

func (o ResourceAccessLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAccessLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_level_name"] = o.AccessLevelName
	toSerialize["access_level_remote_id"] = o.AccessLevelRemoteId
	return toSerialize, nil
}

func (o *ResourceAccessLevel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_level_name",
		"access_level_remote_id",
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

	varResourceAccessLevel := _ResourceAccessLevel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceAccessLevel)

	if err != nil {
		return err
	}

	*o = ResourceAccessLevel(varResourceAccessLevel)

	return err
}

type NullableResourceAccessLevel struct {
	value *ResourceAccessLevel
	isSet bool
}

func (v NullableResourceAccessLevel) Get() *ResourceAccessLevel {
	return v.value
}

func (v *NullableResourceAccessLevel) Set(val *ResourceAccessLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAccessLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAccessLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAccessLevel(val *ResourceAccessLevel) *NullableResourceAccessLevel {
	return &NullableResourceAccessLevel{value: val, isSet: true}
}

func (v NullableResourceAccessLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAccessLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


