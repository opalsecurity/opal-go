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

// checks if the ResourceRemoteInfoGcpFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoGcpFolder{}

// ResourceRemoteInfoGcpFolder Remote info for GCP folder.
type ResourceRemoteInfoGcpFolder struct {
	// The id of the folder.
	FolderId string `json:"folder_id"`
}

type _ResourceRemoteInfoGcpFolder ResourceRemoteInfoGcpFolder

// NewResourceRemoteInfoGcpFolder instantiates a new ResourceRemoteInfoGcpFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoGcpFolder(folderId string) *ResourceRemoteInfoGcpFolder {
	this := ResourceRemoteInfoGcpFolder{}
	this.FolderId = folderId
	return &this
}

// NewResourceRemoteInfoGcpFolderWithDefaults instantiates a new ResourceRemoteInfoGcpFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoGcpFolderWithDefaults() *ResourceRemoteInfoGcpFolder {
	this := ResourceRemoteInfoGcpFolder{}
	return &this
}

// GetFolderId returns the FolderId field value
func (o *ResourceRemoteInfoGcpFolder) GetFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoGcpFolder) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *ResourceRemoteInfoGcpFolder) SetFolderId(v string) {
	o.FolderId = v
}

func (o ResourceRemoteInfoGcpFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoGcpFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folder_id"] = o.FolderId
	return toSerialize, nil
}

func (o *ResourceRemoteInfoGcpFolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"folder_id",
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

	varResourceRemoteInfoGcpFolder := _ResourceRemoteInfoGcpFolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceRemoteInfoGcpFolder)

	if err != nil {
		return err
	}

	*o = ResourceRemoteInfoGcpFolder(varResourceRemoteInfoGcpFolder)

	return err
}

type NullableResourceRemoteInfoGcpFolder struct {
	value *ResourceRemoteInfoGcpFolder
	isSet bool
}

func (v NullableResourceRemoteInfoGcpFolder) Get() *ResourceRemoteInfoGcpFolder {
	return v.value
}

func (v *NullableResourceRemoteInfoGcpFolder) Set(val *ResourceRemoteInfoGcpFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoGcpFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoGcpFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoGcpFolder(val *ResourceRemoteInfoGcpFolder) *NullableResourceRemoteInfoGcpFolder {
	return &NullableResourceRemoteInfoGcpFolder{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoGcpFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoGcpFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


