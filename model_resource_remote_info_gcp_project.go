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

// checks if the ResourceRemoteInfoGcpProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoGcpProject{}

// ResourceRemoteInfoGcpProject Remote info for GCP project.
type ResourceRemoteInfoGcpProject struct {
	// The id of the project.
	ProjectId string `json:"project_id"`
}

type _ResourceRemoteInfoGcpProject ResourceRemoteInfoGcpProject

// NewResourceRemoteInfoGcpProject instantiates a new ResourceRemoteInfoGcpProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoGcpProject(projectId string) *ResourceRemoteInfoGcpProject {
	this := ResourceRemoteInfoGcpProject{}
	this.ProjectId = projectId
	return &this
}

// NewResourceRemoteInfoGcpProjectWithDefaults instantiates a new ResourceRemoteInfoGcpProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoGcpProjectWithDefaults() *ResourceRemoteInfoGcpProject {
	this := ResourceRemoteInfoGcpProject{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ResourceRemoteInfoGcpProject) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoGcpProject) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ResourceRemoteInfoGcpProject) SetProjectId(v string) {
	o.ProjectId = v
}

func (o ResourceRemoteInfoGcpProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoGcpProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId
	return toSerialize, nil
}

func (o *ResourceRemoteInfoGcpProject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
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

	varResourceRemoteInfoGcpProject := _ResourceRemoteInfoGcpProject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceRemoteInfoGcpProject)

	if err != nil {
		return err
	}

	*o = ResourceRemoteInfoGcpProject(varResourceRemoteInfoGcpProject)

	return err
}

type NullableResourceRemoteInfoGcpProject struct {
	value *ResourceRemoteInfoGcpProject
	isSet bool
}

func (v NullableResourceRemoteInfoGcpProject) Get() *ResourceRemoteInfoGcpProject {
	return v.value
}

func (v *NullableResourceRemoteInfoGcpProject) Set(val *ResourceRemoteInfoGcpProject) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoGcpProject) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoGcpProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoGcpProject(val *ResourceRemoteInfoGcpProject) *NullableResourceRemoteInfoGcpProject {
	return &NullableResourceRemoteInfoGcpProject{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoGcpProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoGcpProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


