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

// checks if the AccessList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessList{}

// AccessList struct for AccessList
type AccessList struct {
	Results []Access `json:"results,omitempty"`
}

// NewAccessList instantiates a new AccessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessList() *AccessList {
	this := AccessList{}
	return &this
}

// NewAccessListWithDefaults instantiates a new AccessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessListWithDefaults() *AccessList {
	this := AccessList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AccessList) GetResults() []Access {
	if o == nil || IsNil(o.Results) {
		var ret []Access
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessList) GetResultsOk() ([]Access, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AccessList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Access and assigns it to the Results field.
func (o *AccessList) SetResults(v []Access) {
	o.Results = v
}

func (o AccessList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableAccessList struct {
	value *AccessList
	isSet bool
}

func (v NullableAccessList) Get() *AccessList {
	return v.value
}

func (v *NullableAccessList) Set(val *AccessList) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessList) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessList(val *AccessList) *NullableAccessList {
	return &NullableAccessList{value: val, isSet: true}
}

func (v NullableAccessList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


