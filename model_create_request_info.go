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

// checks if the CreateRequestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRequestInfo{}

// CreateRequestInfo All the information needed for creating a request
type CreateRequestInfo struct {
	Resources []CreateRequestInfoResourcesInner `json:"resources"`
	Groups []CreateRequestInfoGroupsInner `json:"groups"`
	// The ID of the user to be granted access. Should not be specified if target_group_id is specified.
	TargetUserId *string `json:"target_user_id,omitempty"`
	// The ID of the group the request is for.  Should not be specified if target_user_id is specified.
	TargetGroupId *string `json:"target_group_id,omitempty"`
	Reason string `json:"reason"`
	SupportTicket *CreateRequestInfoSupportTicket `json:"support_ticket,omitempty"`
	// The duration of the request in minutes. -1 represents an indefinite duration
	DurationMinutes int32 `json:"duration_minutes"`
	CustomMetadata []CreateRequestInfoCustomMetadataInner `json:"custom_metadata,omitempty"`
}

type _CreateRequestInfo CreateRequestInfo

// NewCreateRequestInfo instantiates a new CreateRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRequestInfo(resources []CreateRequestInfoResourcesInner, groups []CreateRequestInfoGroupsInner, reason string, durationMinutes int32) *CreateRequestInfo {
	this := CreateRequestInfo{}
	this.Resources = resources
	this.Groups = groups
	this.Reason = reason
	this.DurationMinutes = durationMinutes
	return &this
}

// NewCreateRequestInfoWithDefaults instantiates a new CreateRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRequestInfoWithDefaults() *CreateRequestInfo {
	this := CreateRequestInfo{}
	return &this
}

// GetResources returns the Resources field value
func (o *CreateRequestInfo) GetResources() []CreateRequestInfoResourcesInner {
	if o == nil {
		var ret []CreateRequestInfoResourcesInner
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *CreateRequestInfo) GetResourcesOk() ([]CreateRequestInfoResourcesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *CreateRequestInfo) SetResources(v []CreateRequestInfoResourcesInner) {
	o.Resources = v
}

// GetGroups returns the Groups field value
func (o *CreateRequestInfo) GetGroups() []CreateRequestInfoGroupsInner {
	if o == nil {
		var ret []CreateRequestInfoGroupsInner
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *CreateRequestInfo) GetGroupsOk() ([]CreateRequestInfoGroupsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *CreateRequestInfo) SetGroups(v []CreateRequestInfoGroupsInner) {
	o.Groups = v
}

// GetTargetUserId returns the TargetUserId field value if set, zero value otherwise.
func (o *CreateRequestInfo) GetTargetUserId() string {
	if o == nil || IsNil(o.TargetUserId) {
		var ret string
		return ret
	}
	return *o.TargetUserId
}

// GetTargetUserIdOk returns a tuple with the TargetUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestInfo) GetTargetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUserId) {
		return nil, false
	}
	return o.TargetUserId, true
}

// HasTargetUserId returns a boolean if a field has been set.
func (o *CreateRequestInfo) HasTargetUserId() bool {
	if o != nil && !IsNil(o.TargetUserId) {
		return true
	}

	return false
}

// SetTargetUserId gets a reference to the given string and assigns it to the TargetUserId field.
func (o *CreateRequestInfo) SetTargetUserId(v string) {
	o.TargetUserId = &v
}

// GetTargetGroupId returns the TargetGroupId field value if set, zero value otherwise.
func (o *CreateRequestInfo) GetTargetGroupId() string {
	if o == nil || IsNil(o.TargetGroupId) {
		var ret string
		return ret
	}
	return *o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestInfo) GetTargetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetGroupId) {
		return nil, false
	}
	return o.TargetGroupId, true
}

// HasTargetGroupId returns a boolean if a field has been set.
func (o *CreateRequestInfo) HasTargetGroupId() bool {
	if o != nil && !IsNil(o.TargetGroupId) {
		return true
	}

	return false
}

// SetTargetGroupId gets a reference to the given string and assigns it to the TargetGroupId field.
func (o *CreateRequestInfo) SetTargetGroupId(v string) {
	o.TargetGroupId = &v
}

// GetReason returns the Reason field value
func (o *CreateRequestInfo) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *CreateRequestInfo) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *CreateRequestInfo) SetReason(v string) {
	o.Reason = v
}

// GetSupportTicket returns the SupportTicket field value if set, zero value otherwise.
func (o *CreateRequestInfo) GetSupportTicket() CreateRequestInfoSupportTicket {
	if o == nil || IsNil(o.SupportTicket) {
		var ret CreateRequestInfoSupportTicket
		return ret
	}
	return *o.SupportTicket
}

// GetSupportTicketOk returns a tuple with the SupportTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestInfo) GetSupportTicketOk() (*CreateRequestInfoSupportTicket, bool) {
	if o == nil || IsNil(o.SupportTicket) {
		return nil, false
	}
	return o.SupportTicket, true
}

// HasSupportTicket returns a boolean if a field has been set.
func (o *CreateRequestInfo) HasSupportTicket() bool {
	if o != nil && !IsNil(o.SupportTicket) {
		return true
	}

	return false
}

// SetSupportTicket gets a reference to the given CreateRequestInfoSupportTicket and assigns it to the SupportTicket field.
func (o *CreateRequestInfo) SetSupportTicket(v CreateRequestInfoSupportTicket) {
	o.SupportTicket = &v
}

// GetDurationMinutes returns the DurationMinutes field value
func (o *CreateRequestInfo) GetDurationMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationMinutes
}

// GetDurationMinutesOk returns a tuple with the DurationMinutes field value
// and a boolean to check if the value has been set.
func (o *CreateRequestInfo) GetDurationMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMinutes, true
}

// SetDurationMinutes sets field value
func (o *CreateRequestInfo) SetDurationMinutes(v int32) {
	o.DurationMinutes = v
}

// GetCustomMetadata returns the CustomMetadata field value if set, zero value otherwise.
func (o *CreateRequestInfo) GetCustomMetadata() []CreateRequestInfoCustomMetadataInner {
	if o == nil || IsNil(o.CustomMetadata) {
		var ret []CreateRequestInfoCustomMetadataInner
		return ret
	}
	return o.CustomMetadata
}

// GetCustomMetadataOk returns a tuple with the CustomMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestInfo) GetCustomMetadataOk() ([]CreateRequestInfoCustomMetadataInner, bool) {
	if o == nil || IsNil(o.CustomMetadata) {
		return nil, false
	}
	return o.CustomMetadata, true
}

// HasCustomMetadata returns a boolean if a field has been set.
func (o *CreateRequestInfo) HasCustomMetadata() bool {
	if o != nil && !IsNil(o.CustomMetadata) {
		return true
	}

	return false
}

// SetCustomMetadata gets a reference to the given []CreateRequestInfoCustomMetadataInner and assigns it to the CustomMetadata field.
func (o *CreateRequestInfo) SetCustomMetadata(v []CreateRequestInfoCustomMetadataInner) {
	o.CustomMetadata = v
}

func (o CreateRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRequestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resources"] = o.Resources
	toSerialize["groups"] = o.Groups
	if !IsNil(o.TargetUserId) {
		toSerialize["target_user_id"] = o.TargetUserId
	}
	if !IsNil(o.TargetGroupId) {
		toSerialize["target_group_id"] = o.TargetGroupId
	}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.SupportTicket) {
		toSerialize["support_ticket"] = o.SupportTicket
	}
	toSerialize["duration_minutes"] = o.DurationMinutes
	if !IsNil(o.CustomMetadata) {
		toSerialize["custom_metadata"] = o.CustomMetadata
	}
	return toSerialize, nil
}

func (o *CreateRequestInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resources",
		"groups",
		"reason",
		"duration_minutes",
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

	varCreateRequestInfo := _CreateRequestInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRequestInfo)

	if err != nil {
		return err
	}

	*o = CreateRequestInfo(varCreateRequestInfo)

	return err
}

type NullableCreateRequestInfo struct {
	value *CreateRequestInfo
	isSet bool
}

func (v NullableCreateRequestInfo) Get() *CreateRequestInfo {
	return v.value
}

func (v *NullableCreateRequestInfo) Set(val *CreateRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRequestInfo(val *CreateRequestInfo) *NullableCreateRequestInfo {
	return &NullableCreateRequestInfo{value: val, isSet: true}
}

func (v NullableCreateRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


