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

// checks if the CreateRequestConfigurationInfoList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRequestConfigurationInfoList{}

// CreateRequestConfigurationInfoList # CreateRequestConfigurationInfoList Object ### Description The `CreateRequestConfigurationInfoList` object is used as an input to the CreateRequestConfigurations API.  ### Formatting Requirements The `CreateRequestConfigurationInfoList` object must contain a list of `RequestConfiguration` objects. Exactly one default `RequestConfiguration` must be provided.  A default `RequestConfiguration` is one with a `condition` of `null` and a `priority` of `0`.  The default `RequestConfiguration` will be used when no other `RequestConfiguration` matches the request.  Only one `RequestConfiguration` may be provided for each priority, and the priorities must be contiguous.  For example, if there are two `RequestConfigurations` with priorities 0 and 2, there must be a `RequestConfiguration` with priority 1.  To use the `condition` field, the `condition` must be a valid JSON object.  The `condition` must be a JSON object with the key `group_ids` (more options may be added in the future), whose value is a list of group IDs. The `condition` will match if the user requesting access is a member of any of the groups in the list. Currently, we only support using a single group as a condition.
type CreateRequestConfigurationInfoList struct {
	// A list of request configurations to create.
	RequestConfigurations []RequestConfiguration `json:"request_configurations"`
}

// NewCreateRequestConfigurationInfoList instantiates a new CreateRequestConfigurationInfoList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRequestConfigurationInfoList(requestConfigurations []RequestConfiguration) *CreateRequestConfigurationInfoList {
	this := CreateRequestConfigurationInfoList{}
	this.RequestConfigurations = requestConfigurations
	return &this
}

// NewCreateRequestConfigurationInfoListWithDefaults instantiates a new CreateRequestConfigurationInfoList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRequestConfigurationInfoListWithDefaults() *CreateRequestConfigurationInfoList {
	this := CreateRequestConfigurationInfoList{}
	return &this
}

// GetRequestConfigurations returns the RequestConfigurations field value
func (o *CreateRequestConfigurationInfoList) GetRequestConfigurations() []RequestConfiguration {
	if o == nil {
		var ret []RequestConfiguration
		return ret
	}

	return o.RequestConfigurations
}

// GetRequestConfigurationsOk returns a tuple with the RequestConfigurations field value
// and a boolean to check if the value has been set.
func (o *CreateRequestConfigurationInfoList) GetRequestConfigurationsOk() ([]RequestConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestConfigurations, true
}

// SetRequestConfigurations sets field value
func (o *CreateRequestConfigurationInfoList) SetRequestConfigurations(v []RequestConfiguration) {
	o.RequestConfigurations = v
}

func (o CreateRequestConfigurationInfoList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRequestConfigurationInfoList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_configurations"] = o.RequestConfigurations
	return toSerialize, nil
}

type NullableCreateRequestConfigurationInfoList struct {
	value *CreateRequestConfigurationInfoList
	isSet bool
}

func (v NullableCreateRequestConfigurationInfoList) Get() *CreateRequestConfigurationInfoList {
	return v.value
}

func (v *NullableCreateRequestConfigurationInfoList) Set(val *CreateRequestConfigurationInfoList) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRequestConfigurationInfoList) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRequestConfigurationInfoList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRequestConfigurationInfoList(val *CreateRequestConfigurationInfoList) *NullableCreateRequestConfigurationInfoList {
	return &NullableCreateRequestConfigurationInfoList{value: val, isSet: true}
}

func (v NullableCreateRequestConfigurationInfoList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRequestConfigurationInfoList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

