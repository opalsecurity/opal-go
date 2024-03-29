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

// checks if the Owner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Owner{}

// Owner # Owner Object ### Description The `Owner` object is used to represent an owner.
type Owner struct {
	// The ID of the owner.
	OwnerId string `json:"owner_id"`
	// The name of the owner.
	Name *string `json:"name,omitempty"`
	// A description of the owner.
	Description *string `json:"description,omitempty"`
	// The amount of time (in minutes) before the next reviewer is notified. Use 0 to remove escalation policy.
	AccessRequestEscalationPeriod *int32 `json:"access_request_escalation_period,omitempty"`
	ReviewerMessageChannelId NullableString `json:"reviewer_message_channel_id,omitempty"`
	SourceGroupId NullableString `json:"source_group_id,omitempty"`
}

// NewOwner instantiates a new Owner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwner(ownerId string) *Owner {
	this := Owner{}
	this.OwnerId = ownerId
	return &this
}

// NewOwnerWithDefaults instantiates a new Owner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerWithDefaults() *Owner {
	this := Owner{}
	return &this
}

// GetOwnerId returns the OwnerId field value
func (o *Owner) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *Owner) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *Owner) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Owner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Owner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Owner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Owner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Owner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Owner) SetDescription(v string) {
	o.Description = &v
}

// GetAccessRequestEscalationPeriod returns the AccessRequestEscalationPeriod field value if set, zero value otherwise.
func (o *Owner) GetAccessRequestEscalationPeriod() int32 {
	if o == nil || IsNil(o.AccessRequestEscalationPeriod) {
		var ret int32
		return ret
	}
	return *o.AccessRequestEscalationPeriod
}

// GetAccessRequestEscalationPeriodOk returns a tuple with the AccessRequestEscalationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owner) GetAccessRequestEscalationPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessRequestEscalationPeriod) {
		return nil, false
	}
	return o.AccessRequestEscalationPeriod, true
}

// HasAccessRequestEscalationPeriod returns a boolean if a field has been set.
func (o *Owner) HasAccessRequestEscalationPeriod() bool {
	if o != nil && !IsNil(o.AccessRequestEscalationPeriod) {
		return true
	}

	return false
}

// SetAccessRequestEscalationPeriod gets a reference to the given int32 and assigns it to the AccessRequestEscalationPeriod field.
func (o *Owner) SetAccessRequestEscalationPeriod(v int32) {
	o.AccessRequestEscalationPeriod = &v
}

// GetReviewerMessageChannelId returns the ReviewerMessageChannelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Owner) GetReviewerMessageChannelId() string {
	if o == nil || IsNil(o.ReviewerMessageChannelId.Get()) {
		var ret string
		return ret
	}
	return *o.ReviewerMessageChannelId.Get()
}

// GetReviewerMessageChannelIdOk returns a tuple with the ReviewerMessageChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Owner) GetReviewerMessageChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewerMessageChannelId.Get(), o.ReviewerMessageChannelId.IsSet()
}

// HasReviewerMessageChannelId returns a boolean if a field has been set.
func (o *Owner) HasReviewerMessageChannelId() bool {
	if o != nil && o.ReviewerMessageChannelId.IsSet() {
		return true
	}

	return false
}

// SetReviewerMessageChannelId gets a reference to the given NullableString and assigns it to the ReviewerMessageChannelId field.
func (o *Owner) SetReviewerMessageChannelId(v string) {
	o.ReviewerMessageChannelId.Set(&v)
}
// SetReviewerMessageChannelIdNil sets the value for ReviewerMessageChannelId to be an explicit nil
func (o *Owner) SetReviewerMessageChannelIdNil() {
	o.ReviewerMessageChannelId.Set(nil)
}

// UnsetReviewerMessageChannelId ensures that no value is present for ReviewerMessageChannelId, not even an explicit nil
func (o *Owner) UnsetReviewerMessageChannelId() {
	o.ReviewerMessageChannelId.Unset()
}

// GetSourceGroupId returns the SourceGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Owner) GetSourceGroupId() string {
	if o == nil || IsNil(o.SourceGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.SourceGroupId.Get()
}

// GetSourceGroupIdOk returns a tuple with the SourceGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Owner) GetSourceGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceGroupId.Get(), o.SourceGroupId.IsSet()
}

// HasSourceGroupId returns a boolean if a field has been set.
func (o *Owner) HasSourceGroupId() bool {
	if o != nil && o.SourceGroupId.IsSet() {
		return true
	}

	return false
}

// SetSourceGroupId gets a reference to the given NullableString and assigns it to the SourceGroupId field.
func (o *Owner) SetSourceGroupId(v string) {
	o.SourceGroupId.Set(&v)
}
// SetSourceGroupIdNil sets the value for SourceGroupId to be an explicit nil
func (o *Owner) SetSourceGroupIdNil() {
	o.SourceGroupId.Set(nil)
}

// UnsetSourceGroupId ensures that no value is present for SourceGroupId, not even an explicit nil
func (o *Owner) UnsetSourceGroupId() {
	o.SourceGroupId.Unset()
}

func (o Owner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Owner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["owner_id"] = o.OwnerId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AccessRequestEscalationPeriod) {
		toSerialize["access_request_escalation_period"] = o.AccessRequestEscalationPeriod
	}
	if o.ReviewerMessageChannelId.IsSet() {
		toSerialize["reviewer_message_channel_id"] = o.ReviewerMessageChannelId.Get()
	}
	if o.SourceGroupId.IsSet() {
		toSerialize["source_group_id"] = o.SourceGroupId.Get()
	}
	return toSerialize, nil
}

type NullableOwner struct {
	value *Owner
	isSet bool
}

func (v NullableOwner) Get() *Owner {
	return v.value
}

func (v *NullableOwner) Set(val *Owner) {
	v.value = val
	v.isSet = true
}

func (v NullableOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwner(val *Owner) *NullableOwner {
	return &NullableOwner{value: val, isSet: true}
}

func (v NullableOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


