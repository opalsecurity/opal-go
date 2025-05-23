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

// checks if the AddResourceNhiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddResourceNhiRequest{}

// AddResourceNhiRequest struct for AddResourceNhiRequest
type AddResourceNhiRequest struct {
	// The duration for which the resource can be accessed (in minutes). Use 0 to set to indefinite.
	DurationMinutes int32 `json:"duration_minutes"`
	// The remote ID of the access level to grant. If omitted, the default access level remote ID value (empty string) is used.
	AccessLevelRemoteId *string `json:"access_level_remote_id,omitempty"`
}

type _AddResourceNhiRequest AddResourceNhiRequest

// NewAddResourceNhiRequest instantiates a new AddResourceNhiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddResourceNhiRequest(durationMinutes int32) *AddResourceNhiRequest {
	this := AddResourceNhiRequest{}
	this.DurationMinutes = durationMinutes
	return &this
}

// NewAddResourceNhiRequestWithDefaults instantiates a new AddResourceNhiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddResourceNhiRequestWithDefaults() *AddResourceNhiRequest {
	this := AddResourceNhiRequest{}
	return &this
}

// GetDurationMinutes returns the DurationMinutes field value
func (o *AddResourceNhiRequest) GetDurationMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationMinutes
}

// GetDurationMinutesOk returns a tuple with the DurationMinutes field value
// and a boolean to check if the value has been set.
func (o *AddResourceNhiRequest) GetDurationMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMinutes, true
}

// SetDurationMinutes sets field value
func (o *AddResourceNhiRequest) SetDurationMinutes(v int32) {
	o.DurationMinutes = v
}

// GetAccessLevelRemoteId returns the AccessLevelRemoteId field value if set, zero value otherwise.
func (o *AddResourceNhiRequest) GetAccessLevelRemoteId() string {
	if o == nil || IsNil(o.AccessLevelRemoteId) {
		var ret string
		return ret
	}
	return *o.AccessLevelRemoteId
}

// GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddResourceNhiRequest) GetAccessLevelRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessLevelRemoteId) {
		return nil, false
	}
	return o.AccessLevelRemoteId, true
}

// HasAccessLevelRemoteId returns a boolean if a field has been set.
func (o *AddResourceNhiRequest) HasAccessLevelRemoteId() bool {
	if o != nil && !IsNil(o.AccessLevelRemoteId) {
		return true
	}

	return false
}

// SetAccessLevelRemoteId gets a reference to the given string and assigns it to the AccessLevelRemoteId field.
func (o *AddResourceNhiRequest) SetAccessLevelRemoteId(v string) {
	o.AccessLevelRemoteId = &v
}

func (o AddResourceNhiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddResourceNhiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration_minutes"] = o.DurationMinutes
	if !IsNil(o.AccessLevelRemoteId) {
		toSerialize["access_level_remote_id"] = o.AccessLevelRemoteId
	}
	return toSerialize, nil
}

func (o *AddResourceNhiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"duration_minutes",
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

	varAddResourceNhiRequest := _AddResourceNhiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddResourceNhiRequest)

	if err != nil {
		return err
	}

	*o = AddResourceNhiRequest(varAddResourceNhiRequest)

	return err
}

type NullableAddResourceNhiRequest struct {
	value *AddResourceNhiRequest
	isSet bool
}

func (v NullableAddResourceNhiRequest) Get() *AddResourceNhiRequest {
	return v.value
}

func (v *NullableAddResourceNhiRequest) Set(val *AddResourceNhiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddResourceNhiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddResourceNhiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddResourceNhiRequest(val *AddResourceNhiRequest) *NullableAddResourceNhiRequest {
	return &NullableAddResourceNhiRequest{value: val, isSet: true}
}

func (v NullableAddResourceNhiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddResourceNhiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


