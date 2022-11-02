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

// Group # Group Object ### Description The `Group` object is used to represent a group.  ### Usage Example Update from the `UPDATE Groups` endpoint.
type Group struct {
	// The ID of the group.
	GroupId string `json:"group_id"`
	// The ID of the group's app.
	AppId *string `json:"app_id,omitempty"`
	// The name of the group.
	Name *string `json:"name,omitempty"`
	// A description of the group.
	Description *string `json:"description,omitempty"`
	// The ID of the owner of the group.
	AdminOwnerId *string `json:"admin_owner_id,omitempty"`
	// The ID of the remote.
	RemoteId *string `json:"remote_id,omitempty"`
	// The name of the remote.
	RemoteName *string `json:"remote_name,omitempty"`
	GroupType *GroupTypeEnum `json:"group_type,omitempty"`
	// The maximum duration for which the group can be requested (in minutes).
	MaxDuration *int32 `json:"max_duration,omitempty"`
	// The recommended duration for which the group should be requested (in minutes). -1 represents an indefinite duration.
	RecommendedDuration *int32 `json:"recommended_duration,omitempty"`
	// A bool representing whether or not access requests to the group require manager approval.
	// Deprecated
	RequireManagerApproval *bool `json:"require_manager_approval,omitempty"`
	// A bool representing whether or not access requests to the group require an access ticket.
	RequireSupportTicket *bool `json:"require_support_ticket,omitempty"`
	// The ID of the folder that the group is located in.
	FolderId *string `json:"folder_id,omitempty"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this group.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
	// A bool representing whether or not to automatically approve requests to this group.
	AutoApproval *bool `json:"auto_approval,omitempty"`
	// The ID of the associated request template.
	RequestTemplateId *string `json:"request_template_id,omitempty"`
	// JSON metadata about the remote group. Only set for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details.
	Metadata *string `json:"metadata,omitempty"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup(groupId string) *Group {
	this := Group{}
	this.GroupId = groupId
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *Group) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *Group) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *Group) SetGroupId(v string) {
	o.GroupId = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *Group) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *Group) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *Group) SetAppId(v string) {
	o.AppId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Group) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Group) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Group) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Group) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Group) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Group) SetDescription(v string) {
	o.Description = &v
}

// GetAdminOwnerId returns the AdminOwnerId field value if set, zero value otherwise.
func (o *Group) GetAdminOwnerId() string {
	if o == nil || o.AdminOwnerId == nil {
		var ret string
		return ret
	}
	return *o.AdminOwnerId
}

// GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetAdminOwnerIdOk() (*string, bool) {
	if o == nil || o.AdminOwnerId == nil {
		return nil, false
	}
	return o.AdminOwnerId, true
}

// HasAdminOwnerId returns a boolean if a field has been set.
func (o *Group) HasAdminOwnerId() bool {
	if o != nil && o.AdminOwnerId != nil {
		return true
	}

	return false
}

// SetAdminOwnerId gets a reference to the given string and assigns it to the AdminOwnerId field.
func (o *Group) SetAdminOwnerId(v string) {
	o.AdminOwnerId = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *Group) GetRemoteId() string {
	if o == nil || o.RemoteId == nil {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetRemoteIdOk() (*string, bool) {
	if o == nil || o.RemoteId == nil {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Group) HasRemoteId() bool {
	if o != nil && o.RemoteId != nil {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *Group) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetRemoteName returns the RemoteName field value if set, zero value otherwise.
func (o *Group) GetRemoteName() string {
	if o == nil || o.RemoteName == nil {
		var ret string
		return ret
	}
	return *o.RemoteName
}

// GetRemoteNameOk returns a tuple with the RemoteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetRemoteNameOk() (*string, bool) {
	if o == nil || o.RemoteName == nil {
		return nil, false
	}
	return o.RemoteName, true
}

// HasRemoteName returns a boolean if a field has been set.
func (o *Group) HasRemoteName() bool {
	if o != nil && o.RemoteName != nil {
		return true
	}

	return false
}

// SetRemoteName gets a reference to the given string and assigns it to the RemoteName field.
func (o *Group) SetRemoteName(v string) {
	o.RemoteName = &v
}

// GetGroupType returns the GroupType field value if set, zero value otherwise.
func (o *Group) GetGroupType() GroupTypeEnum {
	if o == nil || o.GroupType == nil {
		var ret GroupTypeEnum
		return ret
	}
	return *o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetGroupTypeOk() (*GroupTypeEnum, bool) {
	if o == nil || o.GroupType == nil {
		return nil, false
	}
	return o.GroupType, true
}

// HasGroupType returns a boolean if a field has been set.
func (o *Group) HasGroupType() bool {
	if o != nil && o.GroupType != nil {
		return true
	}

	return false
}

// SetGroupType gets a reference to the given GroupTypeEnum and assigns it to the GroupType field.
func (o *Group) SetGroupType(v GroupTypeEnum) {
	o.GroupType = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
func (o *Group) GetMaxDuration() int32 {
	if o == nil || o.MaxDuration == nil {
		var ret int32
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMaxDurationOk() (*int32, bool) {
	if o == nil || o.MaxDuration == nil {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *Group) HasMaxDuration() bool {
	if o != nil && o.MaxDuration != nil {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given int32 and assigns it to the MaxDuration field.
func (o *Group) SetMaxDuration(v int32) {
	o.MaxDuration = &v
}

// GetRecommendedDuration returns the RecommendedDuration field value if set, zero value otherwise.
func (o *Group) GetRecommendedDuration() int32 {
	if o == nil || o.RecommendedDuration == nil {
		var ret int32
		return ret
	}
	return *o.RecommendedDuration
}

// GetRecommendedDurationOk returns a tuple with the RecommendedDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetRecommendedDurationOk() (*int32, bool) {
	if o == nil || o.RecommendedDuration == nil {
		return nil, false
	}
	return o.RecommendedDuration, true
}

// HasRecommendedDuration returns a boolean if a field has been set.
func (o *Group) HasRecommendedDuration() bool {
	if o != nil && o.RecommendedDuration != nil {
		return true
	}

	return false
}

// SetRecommendedDuration gets a reference to the given int32 and assigns it to the RecommendedDuration field.
func (o *Group) SetRecommendedDuration(v int32) {
	o.RecommendedDuration = &v
}

// GetRequireManagerApproval returns the RequireManagerApproval field value if set, zero value otherwise.
// Deprecated
func (o *Group) GetRequireManagerApproval() bool {
	if o == nil || o.RequireManagerApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireManagerApproval
}

// GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Group) GetRequireManagerApprovalOk() (*bool, bool) {
	if o == nil || o.RequireManagerApproval == nil {
		return nil, false
	}
	return o.RequireManagerApproval, true
}

// HasRequireManagerApproval returns a boolean if a field has been set.
func (o *Group) HasRequireManagerApproval() bool {
	if o != nil && o.RequireManagerApproval != nil {
		return true
	}

	return false
}

// SetRequireManagerApproval gets a reference to the given bool and assigns it to the RequireManagerApproval field.
// Deprecated
func (o *Group) SetRequireManagerApproval(v bool) {
	o.RequireManagerApproval = &v
}

// GetRequireSupportTicket returns the RequireSupportTicket field value if set, zero value otherwise.
func (o *Group) GetRequireSupportTicket() bool {
	if o == nil || o.RequireSupportTicket == nil {
		var ret bool
		return ret
	}
	return *o.RequireSupportTicket
}

// GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetRequireSupportTicketOk() (*bool, bool) {
	if o == nil || o.RequireSupportTicket == nil {
		return nil, false
	}
	return o.RequireSupportTicket, true
}

// HasRequireSupportTicket returns a boolean if a field has been set.
func (o *Group) HasRequireSupportTicket() bool {
	if o != nil && o.RequireSupportTicket != nil {
		return true
	}

	return false
}

// SetRequireSupportTicket gets a reference to the given bool and assigns it to the RequireSupportTicket field.
func (o *Group) SetRequireSupportTicket(v bool) {
	o.RequireSupportTicket = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *Group) GetFolderId() string {
	if o == nil || o.FolderId == nil {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetFolderIdOk() (*string, bool) {
	if o == nil || o.FolderId == nil {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *Group) HasFolderId() bool {
	if o != nil && o.FolderId != nil {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *Group) SetFolderId(v string) {
	o.FolderId = &v
}

// GetRequireMfaToApprove returns the RequireMfaToApprove field value if set, zero value otherwise.
func (o *Group) GetRequireMfaToApprove() bool {
	if o == nil || o.RequireMfaToApprove == nil {
		var ret bool
		return ret
	}
	return *o.RequireMfaToApprove
}

// GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetRequireMfaToApproveOk() (*bool, bool) {
	if o == nil || o.RequireMfaToApprove == nil {
		return nil, false
	}
	return o.RequireMfaToApprove, true
}

// HasRequireMfaToApprove returns a boolean if a field has been set.
func (o *Group) HasRequireMfaToApprove() bool {
	if o != nil && o.RequireMfaToApprove != nil {
		return true
	}

	return false
}

// SetRequireMfaToApprove gets a reference to the given bool and assigns it to the RequireMfaToApprove field.
func (o *Group) SetRequireMfaToApprove(v bool) {
	o.RequireMfaToApprove = &v
}

// GetAutoApproval returns the AutoApproval field value if set, zero value otherwise.
func (o *Group) GetAutoApproval() bool {
	if o == nil || o.AutoApproval == nil {
		var ret bool
		return ret
	}
	return *o.AutoApproval
}

// GetAutoApprovalOk returns a tuple with the AutoApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetAutoApprovalOk() (*bool, bool) {
	if o == nil || o.AutoApproval == nil {
		return nil, false
	}
	return o.AutoApproval, true
}

// HasAutoApproval returns a boolean if a field has been set.
func (o *Group) HasAutoApproval() bool {
	if o != nil && o.AutoApproval != nil {
		return true
	}

	return false
}

// SetAutoApproval gets a reference to the given bool and assigns it to the AutoApproval field.
func (o *Group) SetAutoApproval(v bool) {
	o.AutoApproval = &v
}

// GetRequestTemplateId returns the RequestTemplateId field value if set, zero value otherwise.
func (o *Group) GetRequestTemplateId() string {
	if o == nil || o.RequestTemplateId == nil {
		var ret string
		return ret
	}
	return *o.RequestTemplateId
}

// GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetRequestTemplateIdOk() (*string, bool) {
	if o == nil || o.RequestTemplateId == nil {
		return nil, false
	}
	return o.RequestTemplateId, true
}

// HasRequestTemplateId returns a boolean if a field has been set.
func (o *Group) HasRequestTemplateId() bool {
	if o != nil && o.RequestTemplateId != nil {
		return true
	}

	return false
}

// SetRequestTemplateId gets a reference to the given string and assigns it to the RequestTemplateId field.
func (o *Group) SetRequestTemplateId(v string) {
	o.RequestTemplateId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Group) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Group) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *Group) SetMetadata(v string) {
	o.Metadata = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group_id"] = o.GroupId
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.AdminOwnerId != nil {
		toSerialize["admin_owner_id"] = o.AdminOwnerId
	}
	if o.RemoteId != nil {
		toSerialize["remote_id"] = o.RemoteId
	}
	if o.RemoteName != nil {
		toSerialize["remote_name"] = o.RemoteName
	}
	if o.GroupType != nil {
		toSerialize["group_type"] = o.GroupType
	}
	if o.MaxDuration != nil {
		toSerialize["max_duration"] = o.MaxDuration
	}
	if o.RecommendedDuration != nil {
		toSerialize["recommended_duration"] = o.RecommendedDuration
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
	if o.RequireMfaToApprove != nil {
		toSerialize["require_mfa_to_approve"] = o.RequireMfaToApprove
	}
	if o.AutoApproval != nil {
		toSerialize["auto_approval"] = o.AutoApproval
	}
	if o.RequestTemplateId != nil {
		toSerialize["request_template_id"] = o.RequestTemplateId
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


