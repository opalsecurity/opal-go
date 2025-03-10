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

// checks if the ResourceRemoteInfoSalesforcePermissionSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoSalesforcePermissionSet{}

// ResourceRemoteInfoSalesforcePermissionSet Remote info for Salesforce permission set.
type ResourceRemoteInfoSalesforcePermissionSet struct {
	// The id of the permission set.
	PermissionSetId string `json:"permission_set_id"`
}

type _ResourceRemoteInfoSalesforcePermissionSet ResourceRemoteInfoSalesforcePermissionSet

// NewResourceRemoteInfoSalesforcePermissionSet instantiates a new ResourceRemoteInfoSalesforcePermissionSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoSalesforcePermissionSet(permissionSetId string) *ResourceRemoteInfoSalesforcePermissionSet {
	this := ResourceRemoteInfoSalesforcePermissionSet{}
	this.PermissionSetId = permissionSetId
	return &this
}

// NewResourceRemoteInfoSalesforcePermissionSetWithDefaults instantiates a new ResourceRemoteInfoSalesforcePermissionSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoSalesforcePermissionSetWithDefaults() *ResourceRemoteInfoSalesforcePermissionSet {
	this := ResourceRemoteInfoSalesforcePermissionSet{}
	return &this
}

// GetPermissionSetId returns the PermissionSetId field value
func (o *ResourceRemoteInfoSalesforcePermissionSet) GetPermissionSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionSetId
}

// GetPermissionSetIdOk returns a tuple with the PermissionSetId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoSalesforcePermissionSet) GetPermissionSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionSetId, true
}

// SetPermissionSetId sets field value
func (o *ResourceRemoteInfoSalesforcePermissionSet) SetPermissionSetId(v string) {
	o.PermissionSetId = v
}

func (o ResourceRemoteInfoSalesforcePermissionSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoSalesforcePermissionSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permission_set_id"] = o.PermissionSetId
	return toSerialize, nil
}

func (o *ResourceRemoteInfoSalesforcePermissionSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"permission_set_id",
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

	varResourceRemoteInfoSalesforcePermissionSet := _ResourceRemoteInfoSalesforcePermissionSet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceRemoteInfoSalesforcePermissionSet)

	if err != nil {
		return err
	}

	*o = ResourceRemoteInfoSalesforcePermissionSet(varResourceRemoteInfoSalesforcePermissionSet)

	return err
}

type NullableResourceRemoteInfoSalesforcePermissionSet struct {
	value *ResourceRemoteInfoSalesforcePermissionSet
	isSet bool
}

func (v NullableResourceRemoteInfoSalesforcePermissionSet) Get() *ResourceRemoteInfoSalesforcePermissionSet {
	return v.value
}

func (v *NullableResourceRemoteInfoSalesforcePermissionSet) Set(val *ResourceRemoteInfoSalesforcePermissionSet) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoSalesforcePermissionSet) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoSalesforcePermissionSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoSalesforcePermissionSet(val *ResourceRemoteInfoSalesforcePermissionSet) *NullableResourceRemoteInfoSalesforcePermissionSet {
	return &NullableResourceRemoteInfoSalesforcePermissionSet{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoSalesforcePermissionSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoSalesforcePermissionSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


