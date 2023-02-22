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

// checks if the CreateOwnerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOwnerInfo{}

// CreateOwnerInfo # CreateOwnerInfo Object ### Description The `CreateOwnerInfo` object is used to store creation info for an owner.  ### Usage Example Use in the `POST Owners` endpoint.
type CreateOwnerInfo struct {
	// The name of the owner.
	Name string `json:"name"`
	// A description of the owner.
	Description *string `json:"description,omitempty"`
	// The amount of time (in minutes) before the next reviewer is notified. Use 0 to remove escalation policy.
	AccessRequestEscalationPeriod *int32 `json:"access_request_escalation_period,omitempty"`
	// Users to add to the created owner. If setting a source_group_id this list must be empty.
	UserIds []string `json:"user_ids"`
	// The message channel id for the reviewer channel.
	ReviewerMessageChannelId *string `json:"reviewer_message_channel_id,omitempty"`
	// Sync this owner's user list with a source group.
	SourceGroupId *string `json:"source_group_id,omitempty"`
}

// NewCreateOwnerInfo instantiates a new CreateOwnerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOwnerInfo(name string, userIds []string) *CreateOwnerInfo {
	this := CreateOwnerInfo{}
	this.Name = name
	this.UserIds = userIds
	return &this
}

// NewCreateOwnerInfoWithDefaults instantiates a new CreateOwnerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOwnerInfoWithDefaults() *CreateOwnerInfo {
	this := CreateOwnerInfo{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOwnerInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOwnerInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOwnerInfo) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOwnerInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOwnerInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOwnerInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOwnerInfo) SetDescription(v string) {
	o.Description = &v
}

// GetAccessRequestEscalationPeriod returns the AccessRequestEscalationPeriod field value if set, zero value otherwise.
func (o *CreateOwnerInfo) GetAccessRequestEscalationPeriod() int32 {
	if o == nil || IsNil(o.AccessRequestEscalationPeriod) {
		var ret int32
		return ret
	}
	return *o.AccessRequestEscalationPeriod
}

// GetAccessRequestEscalationPeriodOk returns a tuple with the AccessRequestEscalationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOwnerInfo) GetAccessRequestEscalationPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessRequestEscalationPeriod) {
		return nil, false
	}
	return o.AccessRequestEscalationPeriod, true
}

// HasAccessRequestEscalationPeriod returns a boolean if a field has been set.
func (o *CreateOwnerInfo) HasAccessRequestEscalationPeriod() bool {
	if o != nil && !IsNil(o.AccessRequestEscalationPeriod) {
		return true
	}

	return false
}

// SetAccessRequestEscalationPeriod gets a reference to the given int32 and assigns it to the AccessRequestEscalationPeriod field.
func (o *CreateOwnerInfo) SetAccessRequestEscalationPeriod(v int32) {
	o.AccessRequestEscalationPeriod = &v
}

// GetUserIds returns the UserIds field value
func (o *CreateOwnerInfo) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value
// and a boolean to check if the value has been set.
func (o *CreateOwnerInfo) GetUserIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserIds, true
}

// SetUserIds sets field value
func (o *CreateOwnerInfo) SetUserIds(v []string) {
	o.UserIds = v
}

// GetReviewerMessageChannelId returns the ReviewerMessageChannelId field value if set, zero value otherwise.
func (o *CreateOwnerInfo) GetReviewerMessageChannelId() string {
	if o == nil || IsNil(o.ReviewerMessageChannelId) {
		var ret string
		return ret
	}
	return *o.ReviewerMessageChannelId
}

// GetReviewerMessageChannelIdOk returns a tuple with the ReviewerMessageChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOwnerInfo) GetReviewerMessageChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerMessageChannelId) {
		return nil, false
	}
	return o.ReviewerMessageChannelId, true
}

// HasReviewerMessageChannelId returns a boolean if a field has been set.
func (o *CreateOwnerInfo) HasReviewerMessageChannelId() bool {
	if o != nil && !IsNil(o.ReviewerMessageChannelId) {
		return true
	}

	return false
}

// SetReviewerMessageChannelId gets a reference to the given string and assigns it to the ReviewerMessageChannelId field.
func (o *CreateOwnerInfo) SetReviewerMessageChannelId(v string) {
	o.ReviewerMessageChannelId = &v
}

// GetSourceGroupId returns the SourceGroupId field value if set, zero value otherwise.
func (o *CreateOwnerInfo) GetSourceGroupId() string {
	if o == nil || IsNil(o.SourceGroupId) {
		var ret string
		return ret
	}
	return *o.SourceGroupId
}

// GetSourceGroupIdOk returns a tuple with the SourceGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOwnerInfo) GetSourceGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceGroupId) {
		return nil, false
	}
	return o.SourceGroupId, true
}

// HasSourceGroupId returns a boolean if a field has been set.
func (o *CreateOwnerInfo) HasSourceGroupId() bool {
	if o != nil && !IsNil(o.SourceGroupId) {
		return true
	}

	return false
}

// SetSourceGroupId gets a reference to the given string and assigns it to the SourceGroupId field.
func (o *CreateOwnerInfo) SetSourceGroupId(v string) {
	o.SourceGroupId = &v
}

func (o CreateOwnerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOwnerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AccessRequestEscalationPeriod) {
		toSerialize["access_request_escalation_period"] = o.AccessRequestEscalationPeriod
	}
	toSerialize["user_ids"] = o.UserIds
	if !IsNil(o.ReviewerMessageChannelId) {
		toSerialize["reviewer_message_channel_id"] = o.ReviewerMessageChannelId
	}
	if !IsNil(o.SourceGroupId) {
		toSerialize["source_group_id"] = o.SourceGroupId
	}
	return toSerialize, nil
}

type NullableCreateOwnerInfo struct {
	value *CreateOwnerInfo
	isSet bool
}

func (v NullableCreateOwnerInfo) Get() *CreateOwnerInfo {
	return v.value
}

func (v *NullableCreateOwnerInfo) Set(val *CreateOwnerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOwnerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOwnerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOwnerInfo(val *CreateOwnerInfo) *NullableCreateOwnerInfo {
	return &NullableCreateOwnerInfo{value: val, isSet: true}
}

func (v NullableCreateOwnerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOwnerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


