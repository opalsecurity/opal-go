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

// checks if the MessageChannelIDList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageChannelIDList{}

// MessageChannelIDList A list of message channel IDs.
type MessageChannelIDList struct {
	MessageChannelIds []string `json:"message_channel_ids"`
}

// NewMessageChannelIDList instantiates a new MessageChannelIDList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageChannelIDList(messageChannelIds []string) *MessageChannelIDList {
	this := MessageChannelIDList{}
	this.MessageChannelIds = messageChannelIds
	return &this
}

// NewMessageChannelIDListWithDefaults instantiates a new MessageChannelIDList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageChannelIDListWithDefaults() *MessageChannelIDList {
	this := MessageChannelIDList{}
	return &this
}

// GetMessageChannelIds returns the MessageChannelIds field value
func (o *MessageChannelIDList) GetMessageChannelIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MessageChannelIds
}

// GetMessageChannelIdsOk returns a tuple with the MessageChannelIds field value
// and a boolean to check if the value has been set.
func (o *MessageChannelIDList) GetMessageChannelIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageChannelIds, true
}

// SetMessageChannelIds sets field value
func (o *MessageChannelIDList) SetMessageChannelIds(v []string) {
	o.MessageChannelIds = v
}

func (o MessageChannelIDList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageChannelIDList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_channel_ids"] = o.MessageChannelIds
	return toSerialize, nil
}

type NullableMessageChannelIDList struct {
	value *MessageChannelIDList
	isSet bool
}

func (v NullableMessageChannelIDList) Get() *MessageChannelIDList {
	return v.value
}

func (v *NullableMessageChannelIDList) Set(val *MessageChannelIDList) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageChannelIDList) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageChannelIDList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageChannelIDList(val *MessageChannelIDList) *NullableMessageChannelIDList {
	return &NullableMessageChannelIDList{value: val, isSet: true}
}

func (v NullableMessageChannelIDList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageChannelIDList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


