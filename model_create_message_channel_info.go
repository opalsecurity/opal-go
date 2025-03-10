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

// checks if the CreateMessageChannelInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessageChannelInfo{}

// CreateMessageChannelInfo # CreateMessageChannelInfo Object ### Description The `CreateMessageChannelInfo` object is used to describe the message channel object to be created.
type CreateMessageChannelInfo struct {
	ThirdPartyProvider MessageChannelProviderEnum `json:"third_party_provider"`
	// The remote ID of the message channel
	RemoteId string `json:"remote_id"`
}

type _CreateMessageChannelInfo CreateMessageChannelInfo

// NewCreateMessageChannelInfo instantiates a new CreateMessageChannelInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessageChannelInfo(thirdPartyProvider MessageChannelProviderEnum, remoteId string) *CreateMessageChannelInfo {
	this := CreateMessageChannelInfo{}
	this.ThirdPartyProvider = thirdPartyProvider
	this.RemoteId = remoteId
	return &this
}

// NewCreateMessageChannelInfoWithDefaults instantiates a new CreateMessageChannelInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessageChannelInfoWithDefaults() *CreateMessageChannelInfo {
	this := CreateMessageChannelInfo{}
	return &this
}

// GetThirdPartyProvider returns the ThirdPartyProvider field value
func (o *CreateMessageChannelInfo) GetThirdPartyProvider() MessageChannelProviderEnum {
	if o == nil {
		var ret MessageChannelProviderEnum
		return ret
	}

	return o.ThirdPartyProvider
}

// GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field value
// and a boolean to check if the value has been set.
func (o *CreateMessageChannelInfo) GetThirdPartyProviderOk() (*MessageChannelProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThirdPartyProvider, true
}

// SetThirdPartyProvider sets field value
func (o *CreateMessageChannelInfo) SetThirdPartyProvider(v MessageChannelProviderEnum) {
	o.ThirdPartyProvider = v
}

// GetRemoteId returns the RemoteId field value
func (o *CreateMessageChannelInfo) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *CreateMessageChannelInfo) GetRemoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *CreateMessageChannelInfo) SetRemoteId(v string) {
	o.RemoteId = v
}

func (o CreateMessageChannelInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMessageChannelInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["third_party_provider"] = o.ThirdPartyProvider
	toSerialize["remote_id"] = o.RemoteId
	return toSerialize, nil
}

func (o *CreateMessageChannelInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"third_party_provider",
		"remote_id",
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

	varCreateMessageChannelInfo := _CreateMessageChannelInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMessageChannelInfo)

	if err != nil {
		return err
	}

	*o = CreateMessageChannelInfo(varCreateMessageChannelInfo)

	return err
}

type NullableCreateMessageChannelInfo struct {
	value *CreateMessageChannelInfo
	isSet bool
}

func (v NullableCreateMessageChannelInfo) Get() *CreateMessageChannelInfo {
	return v.value
}

func (v *NullableCreateMessageChannelInfo) Set(val *CreateMessageChannelInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessageChannelInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessageChannelInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessageChannelInfo(val *CreateMessageChannelInfo) *NullableCreateMessageChannelInfo {
	return &NullableCreateMessageChannelInfo{value: val, isSet: true}
}

func (v NullableCreateMessageChannelInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessageChannelInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


