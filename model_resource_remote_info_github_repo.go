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

// checks if the ResourceRemoteInfoGithubRepo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoGithubRepo{}

// ResourceRemoteInfoGithubRepo Remote info for GitHub repository.
type ResourceRemoteInfoGithubRepo struct {
	// The id of the repository.
	// Deprecated
	RepoId *string `json:"repo_id,omitempty"`
	// The name of the repository.
	RepoName string `json:"repo_name"`
}

type _ResourceRemoteInfoGithubRepo ResourceRemoteInfoGithubRepo

// NewResourceRemoteInfoGithubRepo instantiates a new ResourceRemoteInfoGithubRepo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoGithubRepo(repoName string) *ResourceRemoteInfoGithubRepo {
	this := ResourceRemoteInfoGithubRepo{}
	this.RepoName = repoName
	return &this
}

// NewResourceRemoteInfoGithubRepoWithDefaults instantiates a new ResourceRemoteInfoGithubRepo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoGithubRepoWithDefaults() *ResourceRemoteInfoGithubRepo {
	this := ResourceRemoteInfoGithubRepo{}
	return &this
}

// GetRepoId returns the RepoId field value if set, zero value otherwise.
// Deprecated
func (o *ResourceRemoteInfoGithubRepo) GetRepoId() string {
	if o == nil || IsNil(o.RepoId) {
		var ret string
		return ret
	}
	return *o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ResourceRemoteInfoGithubRepo) GetRepoIdOk() (*string, bool) {
	if o == nil || IsNil(o.RepoId) {
		return nil, false
	}
	return o.RepoId, true
}

// HasRepoId returns a boolean if a field has been set.
func (o *ResourceRemoteInfoGithubRepo) HasRepoId() bool {
	if o != nil && !IsNil(o.RepoId) {
		return true
	}

	return false
}

// SetRepoId gets a reference to the given string and assigns it to the RepoId field.
// Deprecated
func (o *ResourceRemoteInfoGithubRepo) SetRepoId(v string) {
	o.RepoId = &v
}

// GetRepoName returns the RepoName field value
func (o *ResourceRemoteInfoGithubRepo) GetRepoName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepoName
}

// GetRepoNameOk returns a tuple with the RepoName field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoGithubRepo) GetRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepoName, true
}

// SetRepoName sets field value
func (o *ResourceRemoteInfoGithubRepo) SetRepoName(v string) {
	o.RepoName = v
}

func (o ResourceRemoteInfoGithubRepo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoGithubRepo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepoId) {
		toSerialize["repo_id"] = o.RepoId
	}
	toSerialize["repo_name"] = o.RepoName
	return toSerialize, nil
}

func (o *ResourceRemoteInfoGithubRepo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"repo_name",
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

	varResourceRemoteInfoGithubRepo := _ResourceRemoteInfoGithubRepo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceRemoteInfoGithubRepo)

	if err != nil {
		return err
	}

	*o = ResourceRemoteInfoGithubRepo(varResourceRemoteInfoGithubRepo)

	return err
}

type NullableResourceRemoteInfoGithubRepo struct {
	value *ResourceRemoteInfoGithubRepo
	isSet bool
}

func (v NullableResourceRemoteInfoGithubRepo) Get() *ResourceRemoteInfoGithubRepo {
	return v.value
}

func (v *NullableResourceRemoteInfoGithubRepo) Set(val *ResourceRemoteInfoGithubRepo) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoGithubRepo) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoGithubRepo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoGithubRepo(val *ResourceRemoteInfoGithubRepo) *NullableResourceRemoteInfoGithubRepo {
	return &NullableResourceRemoteInfoGithubRepo{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoGithubRepo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoGithubRepo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


