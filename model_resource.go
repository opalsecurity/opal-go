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

// Resource # Resource Object ### Description The `Resource` object is used to represent a resource.  ### Usage Example Update from the `UPDATE Resources` endpoint.
type Resource struct {
	// The ID of the resource.
	ResourceId string `json:"resource_id"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// A description of the resource.
	Description *string `json:"description,omitempty"`
	// The ID of the owning team of the resource.
	OwnerTeamId *string `json:"owner_team_id,omitempty"`
	Visibility *VisibilityEnum `json:"visibility,omitempty"`
	ResourceType *ResourceTypeEnum `json:"resource_type,omitempty"`
	// The maximum duration access to the resource can be requested for (in minutes).
	MaxDuration *int32 `json:"max_duration,omitempty"`
	// A bool representing whether or not access requests to the resource require manager approval.
	RequireManagerApproval *bool `json:"require_manager_approval,omitempty"`
	// A bool representing whether or not access requests to the resource require a support ticket.
	RequireSupportTicket *bool `json:"require_support_ticket,omitempty"`
	// The ID of the folder that the resource is located in.
	FolderId *string `json:"folder_id,omitempty"`
}

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource(resourceId string) *Resource {
	this := Resource{}
	this.ResourceId = resourceId
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *Resource) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *Resource) SetResourceId(v string) {
	o.ResourceId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Resource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Resource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Resource) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Resource) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Resource) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Resource) SetDescription(v string) {
	o.Description = &v
}

// GetOwnerTeamId returns the OwnerTeamId field value if set, zero value otherwise.
func (o *Resource) GetOwnerTeamId() string {
	if o == nil || o.OwnerTeamId == nil {
		var ret string
		return ret
	}
	return *o.OwnerTeamId
}

// GetOwnerTeamIdOk returns a tuple with the OwnerTeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetOwnerTeamIdOk() (*string, bool) {
	if o == nil || o.OwnerTeamId == nil {
		return nil, false
	}
	return o.OwnerTeamId, true
}

// HasOwnerTeamId returns a boolean if a field has been set.
func (o *Resource) HasOwnerTeamId() bool {
	if o != nil && o.OwnerTeamId != nil {
		return true
	}

	return false
}

// SetOwnerTeamId gets a reference to the given string and assigns it to the OwnerTeamId field.
func (o *Resource) SetOwnerTeamId(v string) {
	o.OwnerTeamId = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *Resource) GetVisibility() VisibilityEnum {
	if o == nil || o.Visibility == nil {
		var ret VisibilityEnum
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetVisibilityOk() (*VisibilityEnum, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *Resource) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given VisibilityEnum and assigns it to the Visibility field.
func (o *Resource) SetVisibility(v VisibilityEnum) {
	o.Visibility = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *Resource) GetResourceType() ResourceTypeEnum {
	if o == nil || o.ResourceType == nil {
		var ret ResourceTypeEnum
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceTypeOk() (*ResourceTypeEnum, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *Resource) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given ResourceTypeEnum and assigns it to the ResourceType field.
func (o *Resource) SetResourceType(v ResourceTypeEnum) {
	o.ResourceType = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
func (o *Resource) GetMaxDuration() int32 {
	if o == nil || o.MaxDuration == nil {
		var ret int32
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetMaxDurationOk() (*int32, bool) {
	if o == nil || o.MaxDuration == nil {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *Resource) HasMaxDuration() bool {
	if o != nil && o.MaxDuration != nil {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given int32 and assigns it to the MaxDuration field.
func (o *Resource) SetMaxDuration(v int32) {
	o.MaxDuration = &v
}

// GetRequireManagerApproval returns the RequireManagerApproval field value if set, zero value otherwise.
func (o *Resource) GetRequireManagerApproval() bool {
	if o == nil || o.RequireManagerApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireManagerApproval
}

// GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetRequireManagerApprovalOk() (*bool, bool) {
	if o == nil || o.RequireManagerApproval == nil {
		return nil, false
	}
	return o.RequireManagerApproval, true
}

// HasRequireManagerApproval returns a boolean if a field has been set.
func (o *Resource) HasRequireManagerApproval() bool {
	if o != nil && o.RequireManagerApproval != nil {
		return true
	}

	return false
}

// SetRequireManagerApproval gets a reference to the given bool and assigns it to the RequireManagerApproval field.
func (o *Resource) SetRequireManagerApproval(v bool) {
	o.RequireManagerApproval = &v
}

// GetRequireSupportTicket returns the RequireSupportTicket field value if set, zero value otherwise.
func (o *Resource) GetRequireSupportTicket() bool {
	if o == nil || o.RequireSupportTicket == nil {
		var ret bool
		return ret
	}
	return *o.RequireSupportTicket
}

// GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetRequireSupportTicketOk() (*bool, bool) {
	if o == nil || o.RequireSupportTicket == nil {
		return nil, false
	}
	return o.RequireSupportTicket, true
}

// HasRequireSupportTicket returns a boolean if a field has been set.
func (o *Resource) HasRequireSupportTicket() bool {
	if o != nil && o.RequireSupportTicket != nil {
		return true
	}

	return false
}

// SetRequireSupportTicket gets a reference to the given bool and assigns it to the RequireSupportTicket field.
func (o *Resource) SetRequireSupportTicket(v bool) {
	o.RequireSupportTicket = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *Resource) GetFolderId() string {
	if o == nil || o.FolderId == nil {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetFolderIdOk() (*string, bool) {
	if o == nil || o.FolderId == nil {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *Resource) HasFolderId() bool {
	if o != nil && o.FolderId != nil {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *Resource) SetFolderId(v string) {
	o.FolderId = &v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.OwnerTeamId != nil {
		toSerialize["owner_team_id"] = o.OwnerTeamId
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.MaxDuration != nil {
		toSerialize["max_duration"] = o.MaxDuration
	}
	if o.RequireManagerApproval != nil {
		toSerialize["require_manager_approval"] = o.RequireManagerApproval
	}
	if o.RequireSupportTicket != nil {
		toSerialize["require_support_ticket"] = o.RequireSupportTicket
	}
	if o.FolderId != nil {
		toSerialize["folder_id"] = o.FolderId
	}
	return json.Marshal(toSerialize)
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


