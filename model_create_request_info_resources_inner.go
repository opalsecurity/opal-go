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
)

// checks if the CreateRequestInfoResourcesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRequestInfoResourcesInner{}

// CreateRequestInfoResourcesInner struct for CreateRequestInfoResourcesInner
type CreateRequestInfoResourcesInner struct {
	// The ID of the resource requested. Should not be specified if group_id is specified.
	Id *string `json:"id,omitempty"`
	// The ID of the access level requested on the remote system.
	AccessLevelRemoteId *string `json:"access_level_remote_id,omitempty"`
	// The ID of the access level requested on the remote system.
	AccessLevelName *string `json:"access_level_name,omitempty"`
}

// NewCreateRequestInfoResourcesInner instantiates a new CreateRequestInfoResourcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRequestInfoResourcesInner() *CreateRequestInfoResourcesInner {
	this := CreateRequestInfoResourcesInner{}
	return &this
}

// NewCreateRequestInfoResourcesInnerWithDefaults instantiates a new CreateRequestInfoResourcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRequestInfoResourcesInnerWithDefaults() *CreateRequestInfoResourcesInner {
	this := CreateRequestInfoResourcesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateRequestInfoResourcesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestInfoResourcesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateRequestInfoResourcesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateRequestInfoResourcesInner) SetId(v string) {
	o.Id = &v
}

// GetAccessLevelRemoteId returns the AccessLevelRemoteId field value if set, zero value otherwise.
func (o *CreateRequestInfoResourcesInner) GetAccessLevelRemoteId() string {
	if o == nil || IsNil(o.AccessLevelRemoteId) {
		var ret string
		return ret
	}
	return *o.AccessLevelRemoteId
}

// GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestInfoResourcesInner) GetAccessLevelRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessLevelRemoteId) {
		return nil, false
	}
	return o.AccessLevelRemoteId, true
}

// HasAccessLevelRemoteId returns a boolean if a field has been set.
func (o *CreateRequestInfoResourcesInner) HasAccessLevelRemoteId() bool {
	if o != nil && !IsNil(o.AccessLevelRemoteId) {
		return true
	}

	return false
}

// SetAccessLevelRemoteId gets a reference to the given string and assigns it to the AccessLevelRemoteId field.
func (o *CreateRequestInfoResourcesInner) SetAccessLevelRemoteId(v string) {
	o.AccessLevelRemoteId = &v
}

// GetAccessLevelName returns the AccessLevelName field value if set, zero value otherwise.
func (o *CreateRequestInfoResourcesInner) GetAccessLevelName() string {
	if o == nil || IsNil(o.AccessLevelName) {
		var ret string
		return ret
	}
	return *o.AccessLevelName
}

// GetAccessLevelNameOk returns a tuple with the AccessLevelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestInfoResourcesInner) GetAccessLevelNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccessLevelName) {
		return nil, false
	}
	return o.AccessLevelName, true
}

// HasAccessLevelName returns a boolean if a field has been set.
func (o *CreateRequestInfoResourcesInner) HasAccessLevelName() bool {
	if o != nil && !IsNil(o.AccessLevelName) {
		return true
	}

	return false
}

// SetAccessLevelName gets a reference to the given string and assigns it to the AccessLevelName field.
func (o *CreateRequestInfoResourcesInner) SetAccessLevelName(v string) {
	o.AccessLevelName = &v
}

func (o CreateRequestInfoResourcesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRequestInfoResourcesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AccessLevelRemoteId) {
		toSerialize["access_level_remote_id"] = o.AccessLevelRemoteId
	}
	if !IsNil(o.AccessLevelName) {
		toSerialize["access_level_name"] = o.AccessLevelName
	}
	return toSerialize, nil
}

type NullableCreateRequestInfoResourcesInner struct {
	value *CreateRequestInfoResourcesInner
	isSet bool
}

func (v NullableCreateRequestInfoResourcesInner) Get() *CreateRequestInfoResourcesInner {
	return v.value
}

func (v *NullableCreateRequestInfoResourcesInner) Set(val *CreateRequestInfoResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRequestInfoResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRequestInfoResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRequestInfoResourcesInner(val *CreateRequestInfoResourcesInner) *NullableCreateRequestInfoResourcesInner {
	return &NullableCreateRequestInfoResourcesInner{value: val, isSet: true}
}

func (v NullableCreateRequestInfoResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRequestInfoResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


