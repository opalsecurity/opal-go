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

// MessageChannel # MessageChannel Object ### Description The `MessageChannel` object is used to represent a message channel.  ### Usage Example Update a groups message channel from the `UPDATE Groups` endpoint.
type MessageChannel struct {
	// The ID of the message channel.
	MessageChannelId string `json:"message_channel_id"`
	ThirdPartyProvider *MessageChannelProviderEnum `json:"third_party_provider,omitempty"`
	// The remote ID of the message channel
	RemoteId *string `json:"remote_id,omitempty"`
	// The name of the message channel.
	Name *string `json:"name,omitempty"`
	MessageChannelType *MessageChannelTypeEnum `json:"message_channel_type,omitempty"`
	// A bool representing whether or not the message channel is private.
	IsPrivate *bool `json:"is_private,omitempty"`
}

// NewMessageChannel instantiates a new MessageChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageChannel(messageChannelId string) *MessageChannel {
	this := MessageChannel{}
	this.MessageChannelId = messageChannelId
	return &this
}

// NewMessageChannelWithDefaults instantiates a new MessageChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageChannelWithDefaults() *MessageChannel {
	this := MessageChannel{}
	return &this
}

// GetMessageChannelId returns the MessageChannelId field value
func (o *MessageChannel) GetMessageChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageChannelId
}

// GetMessageChannelIdOk returns a tuple with the MessageChannelId field value
// and a boolean to check if the value has been set.
func (o *MessageChannel) GetMessageChannelIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MessageChannelId, true
}

// SetMessageChannelId sets field value
func (o *MessageChannel) SetMessageChannelId(v string) {
	o.MessageChannelId = v
}

// GetThirdPartyProvider returns the ThirdPartyProvider field value if set, zero value otherwise.
func (o *MessageChannel) GetThirdPartyProvider() MessageChannelProviderEnum {
	if o == nil || o.ThirdPartyProvider == nil {
		var ret MessageChannelProviderEnum
		return ret
	}
	return *o.ThirdPartyProvider
}

// GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageChannel) GetThirdPartyProviderOk() (*MessageChannelProviderEnum, bool) {
	if o == nil || o.ThirdPartyProvider == nil {
		return nil, false
	}
	return o.ThirdPartyProvider, true
}

// HasThirdPartyProvider returns a boolean if a field has been set.
func (o *MessageChannel) HasThirdPartyProvider() bool {
	if o != nil && o.ThirdPartyProvider != nil {
		return true
	}

	return false
}

// SetThirdPartyProvider gets a reference to the given MessageChannelProviderEnum and assigns it to the ThirdPartyProvider field.
func (o *MessageChannel) SetThirdPartyProvider(v MessageChannelProviderEnum) {
	o.ThirdPartyProvider = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *MessageChannel) GetRemoteId() string {
	if o == nil || o.RemoteId == nil {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageChannel) GetRemoteIdOk() (*string, bool) {
	if o == nil || o.RemoteId == nil {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *MessageChannel) HasRemoteId() bool {
	if o != nil && o.RemoteId != nil {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *MessageChannel) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MessageChannel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageChannel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MessageChannel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MessageChannel) SetName(v string) {
	o.Name = &v
}

// GetMessageChannelType returns the MessageChannelType field value if set, zero value otherwise.
func (o *MessageChannel) GetMessageChannelType() MessageChannelTypeEnum {
	if o == nil || o.MessageChannelType == nil {
		var ret MessageChannelTypeEnum
		return ret
	}
	return *o.MessageChannelType
}

// GetMessageChannelTypeOk returns a tuple with the MessageChannelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageChannel) GetMessageChannelTypeOk() (*MessageChannelTypeEnum, bool) {
	if o == nil || o.MessageChannelType == nil {
		return nil, false
	}
	return o.MessageChannelType, true
}

// HasMessageChannelType returns a boolean if a field has been set.
func (o *MessageChannel) HasMessageChannelType() bool {
	if o != nil && o.MessageChannelType != nil {
		return true
	}

	return false
}

// SetMessageChannelType gets a reference to the given MessageChannelTypeEnum and assigns it to the MessageChannelType field.
func (o *MessageChannel) SetMessageChannelType(v MessageChannelTypeEnum) {
	o.MessageChannelType = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *MessageChannel) GetIsPrivate() bool {
	if o == nil || o.IsPrivate == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageChannel) GetIsPrivateOk() (*bool, bool) {
	if o == nil || o.IsPrivate == nil {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *MessageChannel) HasIsPrivate() bool {
	if o != nil && o.IsPrivate != nil {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *MessageChannel) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

func (o MessageChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message_channel_id"] = o.MessageChannelId
	}
	if o.ThirdPartyProvider != nil {
		toSerialize["third_party_provider"] = o.ThirdPartyProvider
	}
	if o.RemoteId != nil {
		toSerialize["remote_id"] = o.RemoteId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.MessageChannelType != nil {
		toSerialize["message_channel_type"] = o.MessageChannelType
	}
	if o.IsPrivate != nil {
		toSerialize["is_private"] = o.IsPrivate
	}
	return json.Marshal(toSerialize)
}

type NullableMessageChannel struct {
	value *MessageChannel
	isSet bool
}

func (v NullableMessageChannel) Get() *MessageChannel {
	return v.value
}

func (v *NullableMessageChannel) Set(val *MessageChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageChannel(val *MessageChannel) *NullableMessageChannel {
	return &NullableMessageChannel{value: val, isSet: true}
}

func (v NullableMessageChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

