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

// checks if the UpdateGroupInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGroupInfo{}

// UpdateGroupInfo # UpdateGroupInfo Object ### Description The `UpdateGroupInfo` object is used as an input to the UpdateGroup API.
type UpdateGroupInfo struct {
	// The ID of the group.
	GroupId string `json:"group_id"`
	// The name of the group.
	Name *string `json:"name,omitempty"`
	// A description of the group.
	Description *string `json:"description,omitempty"`
	// The ID of the owner of the group.
	AdminOwnerId *string `json:"admin_owner_id,omitempty"`
	// The maximum duration for which the group can be requested (in minutes). Use -1 to set to indefinite. Deprecated in favor of `request_configuration_list`.
	// Deprecated
	MaxDuration *int32 `json:"max_duration,omitempty"`
	// The recommended duration for which the group should be requested (in minutes). Will be the default value in a request. Use -1 to set to indefinite and 0 to unset. Deprecated in favor of `request_configuration_list`.
	// Deprecated
	RecommendedDuration *int32 `json:"recommended_duration,omitempty"`
	// A bool representing whether or not access requests to the group require manager approval. Deprecated in favor of `request_configuration_list`.
	// Deprecated
	RequireManagerApproval *bool `json:"require_manager_approval,omitempty"`
	// A bool representing whether or not access requests to the group require an access ticket. Deprecated in favor of `request_configuration_list`.
	// Deprecated
	RequireSupportTicket *bool `json:"require_support_ticket,omitempty"`
	// The ID of the folder that the group is located in.
	// Deprecated
	FolderId *string `json:"folder_id,omitempty"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this group.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
	// A bool representing whether or not to require MFA for requesting access to this group. Deprecated in favor of `request_configuration_list`.
	// Deprecated
	RequireMfaToRequest *bool `json:"require_mfa_to_request,omitempty"`
	// A bool representing whether or not to automatically approve requests to this group. Deprecated in favor of `request_configuration_list`.
	// Deprecated
	AutoApproval *bool `json:"auto_approval,omitempty"`
	// The ID of the associated configuration template.
	ConfigurationTemplateId *string `json:"configuration_template_id,omitempty"`
	// The ID of the associated request template. Deprecated in favor of `request_configuration_list`.
	// Deprecated
	RequestTemplateId *string `json:"request_template_id,omitempty"`
	// A bool representing whether or not to allow access requests to this group. Deprecated in favor of `request_configuration_list`.
	// Deprecated
	IsRequestable *bool `json:"is_requestable,omitempty"`
	RequestConfigurationList *CreateRequestConfigurationInfoList `json:"request_configuration_list,omitempty"`
}

// NewUpdateGroupInfo instantiates a new UpdateGroupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGroupInfo(groupId string) *UpdateGroupInfo {
	this := UpdateGroupInfo{}
	this.GroupId = groupId
	return &this
}

// NewUpdateGroupInfoWithDefaults instantiates a new UpdateGroupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGroupInfoWithDefaults() *UpdateGroupInfo {
	this := UpdateGroupInfo{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *UpdateGroupInfo) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupInfo) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *UpdateGroupInfo) SetGroupId(v string) {
	o.GroupId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateGroupInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateGroupInfo) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateGroupInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateGroupInfo) SetDescription(v string) {
	o.Description = &v
}

// GetAdminOwnerId returns the AdminOwnerId field value if set, zero value otherwise.
func (o *UpdateGroupInfo) GetAdminOwnerId() string {
	if o == nil || IsNil(o.AdminOwnerId) {
		var ret string
		return ret
	}
	return *o.AdminOwnerId
}

// GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupInfo) GetAdminOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdminOwnerId) {
		return nil, false
	}
	return o.AdminOwnerId, true
}

// HasAdminOwnerId returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasAdminOwnerId() bool {
	if o != nil && !IsNil(o.AdminOwnerId) {
		return true
	}

	return false
}

// SetAdminOwnerId gets a reference to the given string and assigns it to the AdminOwnerId field.
func (o *UpdateGroupInfo) SetAdminOwnerId(v string) {
	o.AdminOwnerId = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
// Deprecated
func (o *UpdateGroupInfo) GetMaxDuration() int32 {
	if o == nil || IsNil(o.MaxDuration) {
		var ret int32
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateGroupInfo) GetMaxDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDuration) {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasMaxDuration() bool {
	if o != nil && !IsNil(o.MaxDuration) {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given int32 and assigns it to the MaxDuration field.
// Deprecated
func (o *UpdateGroupInfo) SetMaxDuration(v int32) {
	o.MaxDuration = &v
}

// GetRecommendedDuration returns the RecommendedDuration field value if set, zero value otherwise.
// Deprecated
func (o *UpdateGroupInfo) GetRecommendedDuration() int32 {
	if o == nil || IsNil(o.RecommendedDuration) {
		var ret int32
		return ret
	}
	return *o.RecommendedDuration
}

// GetRecommendedDurationOk returns a tuple with the RecommendedDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateGroupInfo) GetRecommendedDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.RecommendedDuration) {
		return nil, false
	}
	return o.RecommendedDuration, true
}

// HasRecommendedDuration returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasRecommendedDuration() bool {
	if o != nil && !IsNil(o.RecommendedDuration) {
		return true
	}

	return false
}

// SetRecommendedDuration gets a reference to the given int32 and assigns it to the RecommendedDuration field.
// Deprecated
func (o *UpdateGroupInfo) SetRecommendedDuration(v int32) {
	o.RecommendedDuration = &v
}

// GetRequireManagerApproval returns the RequireManagerApproval field value if set, zero value otherwise.
// Deprecated
func (o *UpdateGroupInfo) GetRequireManagerApproval() bool {
	if o == nil || IsNil(o.RequireManagerApproval) {
		var ret bool
		return ret
	}
	return *o.RequireManagerApproval
}

// GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateGroupInfo) GetRequireManagerApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireManagerApproval) {
		return nil, false
	}
	return o.RequireManagerApproval, true
}

// HasRequireManagerApproval returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasRequireManagerApproval() bool {
	if o != nil && !IsNil(o.RequireManagerApproval) {
		return true
	}

	return false
}

// SetRequireManagerApproval gets a reference to the given bool and assigns it to the RequireManagerApproval field.
// Deprecated
func (o *UpdateGroupInfo) SetRequireManagerApproval(v bool) {
	o.RequireManagerApproval = &v
}

// GetRequireSupportTicket returns the RequireSupportTicket field value if set, zero value otherwise.
// Deprecated
func (o *UpdateGroupInfo) GetRequireSupportTicket() bool {
	if o == nil || IsNil(o.RequireSupportTicket) {
		var ret bool
		return ret
	}
	return *o.RequireSupportTicket
}

// GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateGroupInfo) GetRequireSupportTicketOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSupportTicket) {
		return nil, false
	}
	return o.RequireSupportTicket, true
}

// HasRequireSupportTicket returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasRequireSupportTicket() bool {
	if o != nil && !IsNil(o.RequireSupportTicket) {
		return true
	}

	return false
}

// SetRequireSupportTicket gets a reference to the given bool and assigns it to the RequireSupportTicket field.
// Deprecated
func (o *UpdateGroupInfo) SetRequireSupportTicket(v bool) {
	o.RequireSupportTicket = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
// Deprecated
func (o *UpdateGroupInfo) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateGroupInfo) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
// Deprecated
func (o *UpdateGroupInfo) SetFolderId(v string) {
	o.FolderId = &v
}

// GetRequireMfaToApprove returns the RequireMfaToApprove field value if set, zero value otherwise.
func (o *UpdateGroupInfo) GetRequireMfaToApprove() bool {
	if o == nil || IsNil(o.RequireMfaToApprove) {
		var ret bool
		return ret
	}
	return *o.RequireMfaToApprove
}

// GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupInfo) GetRequireMfaToApproveOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireMfaToApprove) {
		return nil, false
	}
	return o.RequireMfaToApprove, true
}

// HasRequireMfaToApprove returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasRequireMfaToApprove() bool {
	if o != nil && !IsNil(o.RequireMfaToApprove) {
		return true
	}

	return false
}

// SetRequireMfaToApprove gets a reference to the given bool and assigns it to the RequireMfaToApprove field.
func (o *UpdateGroupInfo) SetRequireMfaToApprove(v bool) {
	o.RequireMfaToApprove = &v
}

// GetRequireMfaToRequest returns the RequireMfaToRequest field value if set, zero value otherwise.
// Deprecated
func (o *UpdateGroupInfo) GetRequireMfaToRequest() bool {
	if o == nil || IsNil(o.RequireMfaToRequest) {
		var ret bool
		return ret
	}
	return *o.RequireMfaToRequest
}

// GetRequireMfaToRequestOk returns a tuple with the RequireMfaToRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateGroupInfo) GetRequireMfaToRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireMfaToRequest) {
		return nil, false
	}
	return o.RequireMfaToRequest, true
}

// HasRequireMfaToRequest returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasRequireMfaToRequest() bool {
	if o != nil && !IsNil(o.RequireMfaToRequest) {
		return true
	}

	return false
}

// SetRequireMfaToRequest gets a reference to the given bool and assigns it to the RequireMfaToRequest field.
// Deprecated
func (o *UpdateGroupInfo) SetRequireMfaToRequest(v bool) {
	o.RequireMfaToRequest = &v
}

// GetAutoApproval returns the AutoApproval field value if set, zero value otherwise.
// Deprecated
func (o *UpdateGroupInfo) GetAutoApproval() bool {
	if o == nil || IsNil(o.AutoApproval) {
		var ret bool
		return ret
	}
	return *o.AutoApproval
}

// GetAutoApprovalOk returns a tuple with the AutoApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateGroupInfo) GetAutoApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoApproval) {
		return nil, false
	}
	return o.AutoApproval, true
}

// HasAutoApproval returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasAutoApproval() bool {
	if o != nil && !IsNil(o.AutoApproval) {
		return true
	}

	return false
}

// SetAutoApproval gets a reference to the given bool and assigns it to the AutoApproval field.
// Deprecated
func (o *UpdateGroupInfo) SetAutoApproval(v bool) {
	o.AutoApproval = &v
}

// GetConfigurationTemplateId returns the ConfigurationTemplateId field value if set, zero value otherwise.
func (o *UpdateGroupInfo) GetConfigurationTemplateId() string {
	if o == nil || IsNil(o.ConfigurationTemplateId) {
		var ret string
		return ret
	}
	return *o.ConfigurationTemplateId
}

// GetConfigurationTemplateIdOk returns a tuple with the ConfigurationTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupInfo) GetConfigurationTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationTemplateId) {
		return nil, false
	}
	return o.ConfigurationTemplateId, true
}

// HasConfigurationTemplateId returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasConfigurationTemplateId() bool {
	if o != nil && !IsNil(o.ConfigurationTemplateId) {
		return true
	}

	return false
}

// SetConfigurationTemplateId gets a reference to the given string and assigns it to the ConfigurationTemplateId field.
func (o *UpdateGroupInfo) SetConfigurationTemplateId(v string) {
	o.ConfigurationTemplateId = &v
}

// GetRequestTemplateId returns the RequestTemplateId field value if set, zero value otherwise.
// Deprecated
func (o *UpdateGroupInfo) GetRequestTemplateId() string {
	if o == nil || IsNil(o.RequestTemplateId) {
		var ret string
		return ret
	}
	return *o.RequestTemplateId
}

// GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateGroupInfo) GetRequestTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestTemplateId) {
		return nil, false
	}
	return o.RequestTemplateId, true
}

// HasRequestTemplateId returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasRequestTemplateId() bool {
	if o != nil && !IsNil(o.RequestTemplateId) {
		return true
	}

	return false
}

// SetRequestTemplateId gets a reference to the given string and assigns it to the RequestTemplateId field.
// Deprecated
func (o *UpdateGroupInfo) SetRequestTemplateId(v string) {
	o.RequestTemplateId = &v
}

// GetIsRequestable returns the IsRequestable field value if set, zero value otherwise.
// Deprecated
func (o *UpdateGroupInfo) GetIsRequestable() bool {
	if o == nil || IsNil(o.IsRequestable) {
		var ret bool
		return ret
	}
	return *o.IsRequestable
}

// GetIsRequestableOk returns a tuple with the IsRequestable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateGroupInfo) GetIsRequestableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequestable) {
		return nil, false
	}
	return o.IsRequestable, true
}

// HasIsRequestable returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasIsRequestable() bool {
	if o != nil && !IsNil(o.IsRequestable) {
		return true
	}

	return false
}

// SetIsRequestable gets a reference to the given bool and assigns it to the IsRequestable field.
// Deprecated
func (o *UpdateGroupInfo) SetIsRequestable(v bool) {
	o.IsRequestable = &v
}

// GetRequestConfigurationList returns the RequestConfigurationList field value if set, zero value otherwise.
func (o *UpdateGroupInfo) GetRequestConfigurationList() CreateRequestConfigurationInfoList {
	if o == nil || IsNil(o.RequestConfigurationList) {
		var ret CreateRequestConfigurationInfoList
		return ret
	}
	return *o.RequestConfigurationList
}

// GetRequestConfigurationListOk returns a tuple with the RequestConfigurationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupInfo) GetRequestConfigurationListOk() (*CreateRequestConfigurationInfoList, bool) {
	if o == nil || IsNil(o.RequestConfigurationList) {
		return nil, false
	}
	return o.RequestConfigurationList, true
}

// HasRequestConfigurationList returns a boolean if a field has been set.
func (o *UpdateGroupInfo) HasRequestConfigurationList() bool {
	if o != nil && !IsNil(o.RequestConfigurationList) {
		return true
	}

	return false
}

// SetRequestConfigurationList gets a reference to the given CreateRequestConfigurationInfoList and assigns it to the RequestConfigurationList field.
func (o *UpdateGroupInfo) SetRequestConfigurationList(v CreateRequestConfigurationInfoList) {
	o.RequestConfigurationList = &v
}

func (o UpdateGroupInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGroupInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AdminOwnerId) {
		toSerialize["admin_owner_id"] = o.AdminOwnerId
	}
	if !IsNil(o.MaxDuration) {
		toSerialize["max_duration"] = o.MaxDuration
	}
	if !IsNil(o.RecommendedDuration) {
		toSerialize["recommended_duration"] = o.RecommendedDuration
	}
	if !IsNil(o.RequireManagerApproval) {
		toSerialize["require_manager_approval"] = o.RequireManagerApproval
	}
	if !IsNil(o.RequireSupportTicket) {
		toSerialize["require_support_ticket"] = o.RequireSupportTicket
	}
	if !IsNil(o.FolderId) {
		toSerialize["folder_id"] = o.FolderId
	}
	if !IsNil(o.RequireMfaToApprove) {
		toSerialize["require_mfa_to_approve"] = o.RequireMfaToApprove
	}
	if !IsNil(o.RequireMfaToRequest) {
		toSerialize["require_mfa_to_request"] = o.RequireMfaToRequest
	}
	if !IsNil(o.AutoApproval) {
		toSerialize["auto_approval"] = o.AutoApproval
	}
	if !IsNil(o.ConfigurationTemplateId) {
		toSerialize["configuration_template_id"] = o.ConfigurationTemplateId
	}
	if !IsNil(o.RequestTemplateId) {
		toSerialize["request_template_id"] = o.RequestTemplateId
	}
	if !IsNil(o.IsRequestable) {
		toSerialize["is_requestable"] = o.IsRequestable
	}
	if !IsNil(o.RequestConfigurationList) {
		toSerialize["request_configuration_list"] = o.RequestConfigurationList
	}
	return toSerialize, nil
}

type NullableUpdateGroupInfo struct {
	value *UpdateGroupInfo
	isSet bool
}

func (v NullableUpdateGroupInfo) Get() *UpdateGroupInfo {
	return v.value
}

func (v *NullableUpdateGroupInfo) Set(val *UpdateGroupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupInfo(val *UpdateGroupInfo) *NullableUpdateGroupInfo {
	return &NullableUpdateGroupInfo{value: val, isSet: true}
}

func (v NullableUpdateGroupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


