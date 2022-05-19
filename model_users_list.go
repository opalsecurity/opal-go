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

// UsersList struct for UsersList
type UsersList struct {
	Results []User `json:"results,omitempty"`
}

// NewUsersList instantiates a new UsersList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersList() *UsersList {
	this := UsersList{}
	return &this
}

// NewUsersListWithDefaults instantiates a new UsersList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersListWithDefaults() *UsersList {
	this := UsersList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *UsersList) GetResults() []User {
	if o == nil || o.Results == nil {
		var ret []User
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersList) GetResultsOk() ([]User, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *UsersList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []User and assigns it to the Results field.
func (o *UsersList) SetResults(v []User) {
	o.Results = v
}

func (o UsersList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableUsersList struct {
	value *UsersList
	isSet bool
}

func (v NullableUsersList) Get() *UsersList {
	return v.value
}

func (v *NullableUsersList) Set(val *UsersList) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersList) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersList(val *UsersList) *NullableUsersList {
	return &NullableUsersList{value: val, isSet: true}
}

func (v NullableUsersList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


