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

// checks if the ResourceRemoteInfoGcpBigQueryDataset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoGcpBigQueryDataset{}

// ResourceRemoteInfoGcpBigQueryDataset Remote info for GCP BigQuery Dataset.
type ResourceRemoteInfoGcpBigQueryDataset struct {
	// The id of the project the dataset is in.
	ProjectId string `json:"project_id"`
	// The id of the dataset.
	DatasetId string `json:"dataset_id"`
}

type _ResourceRemoteInfoGcpBigQueryDataset ResourceRemoteInfoGcpBigQueryDataset

// NewResourceRemoteInfoGcpBigQueryDataset instantiates a new ResourceRemoteInfoGcpBigQueryDataset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoGcpBigQueryDataset(projectId string, datasetId string) *ResourceRemoteInfoGcpBigQueryDataset {
	this := ResourceRemoteInfoGcpBigQueryDataset{}
	this.ProjectId = projectId
	this.DatasetId = datasetId
	return &this
}

// NewResourceRemoteInfoGcpBigQueryDatasetWithDefaults instantiates a new ResourceRemoteInfoGcpBigQueryDataset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoGcpBigQueryDatasetWithDefaults() *ResourceRemoteInfoGcpBigQueryDataset {
	this := ResourceRemoteInfoGcpBigQueryDataset{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ResourceRemoteInfoGcpBigQueryDataset) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoGcpBigQueryDataset) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ResourceRemoteInfoGcpBigQueryDataset) SetProjectId(v string) {
	o.ProjectId = v
}

// GetDatasetId returns the DatasetId field value
func (o *ResourceRemoteInfoGcpBigQueryDataset) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoGcpBigQueryDataset) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value
func (o *ResourceRemoteInfoGcpBigQueryDataset) SetDatasetId(v string) {
	o.DatasetId = v
}

func (o ResourceRemoteInfoGcpBigQueryDataset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoGcpBigQueryDataset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["dataset_id"] = o.DatasetId
	return toSerialize, nil
}

func (o *ResourceRemoteInfoGcpBigQueryDataset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
		"dataset_id",
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

	varResourceRemoteInfoGcpBigQueryDataset := _ResourceRemoteInfoGcpBigQueryDataset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceRemoteInfoGcpBigQueryDataset)

	if err != nil {
		return err
	}

	*o = ResourceRemoteInfoGcpBigQueryDataset(varResourceRemoteInfoGcpBigQueryDataset)

	return err
}

type NullableResourceRemoteInfoGcpBigQueryDataset struct {
	value *ResourceRemoteInfoGcpBigQueryDataset
	isSet bool
}

func (v NullableResourceRemoteInfoGcpBigQueryDataset) Get() *ResourceRemoteInfoGcpBigQueryDataset {
	return v.value
}

func (v *NullableResourceRemoteInfoGcpBigQueryDataset) Set(val *ResourceRemoteInfoGcpBigQueryDataset) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoGcpBigQueryDataset) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoGcpBigQueryDataset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoGcpBigQueryDataset(val *ResourceRemoteInfoGcpBigQueryDataset) *NullableResourceRemoteInfoGcpBigQueryDataset {
	return &NullableResourceRemoteInfoGcpBigQueryDataset{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoGcpBigQueryDataset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoGcpBigQueryDataset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


