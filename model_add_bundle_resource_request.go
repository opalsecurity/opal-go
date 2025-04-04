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

// checks if the AddBundleResourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBundleResourceRequest{}

// AddBundleResourceRequest struct for AddBundleResourceRequest
type AddBundleResourceRequest struct {
	// The ID of the resource to add.
	ResourceId string `json:"resource_id"`
	// The remote ID of the access level to grant to this user. Required if the resource being added requires an access level. If omitted, the default access level remote ID value (empty string) is used.
	AccessLevelRemoteId *string `json:"access_level_remote_id,omitempty"`
	// The name of the access level to grant to this user. If omitted, the default access level name value (empty string) is used.
	AccessLevelName *string `json:"access_level_name,omitempty"`
}

type _AddBundleResourceRequest AddBundleResourceRequest

// NewAddBundleResourceRequest instantiates a new AddBundleResourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBundleResourceRequest(resourceId string) *AddBundleResourceRequest {
	this := AddBundleResourceRequest{}
	this.ResourceId = resourceId
	return &this
}

// NewAddBundleResourceRequestWithDefaults instantiates a new AddBundleResourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBundleResourceRequestWithDefaults() *AddBundleResourceRequest {
	this := AddBundleResourceRequest{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *AddBundleResourceRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *AddBundleResourceRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *AddBundleResourceRequest) SetResourceId(v string) {
	o.ResourceId = v
}

// GetAccessLevelRemoteId returns the AccessLevelRemoteId field value if set, zero value otherwise.
func (o *AddBundleResourceRequest) GetAccessLevelRemoteId() string {
	if o == nil || IsNil(o.AccessLevelRemoteId) {
		var ret string
		return ret
	}
	return *o.AccessLevelRemoteId
}

// GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBundleResourceRequest) GetAccessLevelRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessLevelRemoteId) {
		return nil, false
	}
	return o.AccessLevelRemoteId, true
}

// HasAccessLevelRemoteId returns a boolean if a field has been set.
func (o *AddBundleResourceRequest) HasAccessLevelRemoteId() bool {
	if o != nil && !IsNil(o.AccessLevelRemoteId) {
		return true
	}

	return false
}

// SetAccessLevelRemoteId gets a reference to the given string and assigns it to the AccessLevelRemoteId field.
func (o *AddBundleResourceRequest) SetAccessLevelRemoteId(v string) {
	o.AccessLevelRemoteId = &v
}

// GetAccessLevelName returns the AccessLevelName field value if set, zero value otherwise.
func (o *AddBundleResourceRequest) GetAccessLevelName() string {
	if o == nil || IsNil(o.AccessLevelName) {
		var ret string
		return ret
	}
	return *o.AccessLevelName
}

// GetAccessLevelNameOk returns a tuple with the AccessLevelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBundleResourceRequest) GetAccessLevelNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccessLevelName) {
		return nil, false
	}
	return o.AccessLevelName, true
}

// HasAccessLevelName returns a boolean if a field has been set.
func (o *AddBundleResourceRequest) HasAccessLevelName() bool {
	if o != nil && !IsNil(o.AccessLevelName) {
		return true
	}

	return false
}

// SetAccessLevelName gets a reference to the given string and assigns it to the AccessLevelName field.
func (o *AddBundleResourceRequest) SetAccessLevelName(v string) {
	o.AccessLevelName = &v
}

func (o AddBundleResourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBundleResourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_id"] = o.ResourceId
	if !IsNil(o.AccessLevelRemoteId) {
		toSerialize["access_level_remote_id"] = o.AccessLevelRemoteId
	}
	if !IsNil(o.AccessLevelName) {
		toSerialize["access_level_name"] = o.AccessLevelName
	}
	return toSerialize, nil
}

func (o *AddBundleResourceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource_id",
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

	varAddBundleResourceRequest := _AddBundleResourceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddBundleResourceRequest)

	if err != nil {
		return err
	}

	*o = AddBundleResourceRequest(varAddBundleResourceRequest)

	return err
}

type NullableAddBundleResourceRequest struct {
	value *AddBundleResourceRequest
	isSet bool
}

func (v NullableAddBundleResourceRequest) Get() *AddBundleResourceRequest {
	return v.value
}

func (v *NullableAddBundleResourceRequest) Set(val *AddBundleResourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddBundleResourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddBundleResourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddBundleResourceRequest(val *AddBundleResourceRequest) *NullableAddBundleResourceRequest {
	return &NullableAddBundleResourceRequest{value: val, isSet: true}
}

func (v NullableAddBundleResourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddBundleResourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


