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

// checks if the CreateGroupInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupInfo{}

// CreateGroupInfo # CreateGroupInfo Object ### Description The `CreateGroupInfo` object is used to store creation info for a group.  ### Usage Example Use in the `POST Groups` endpoint.
type CreateGroupInfo struct {
	// The name of the remote group.
	Name string `json:"name"`
	// A description of the remote group.
	Description *string `json:"description,omitempty"`
	GroupType GroupTypeEnum `json:"group_type"`
	// The ID of the app for the group.
	AppId string `json:"app_id"`
	RemoteInfo *GroupRemoteInfo `json:"remote_info,omitempty"`
	// Deprecated - use remote_info instead. The ID of the group on the remote system. Include only for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details on how to specify this field.
	// Deprecated
	RemoteGroupId *string `json:"remote_group_id,omitempty"`
	// Deprecated - use remote_info instead.  JSON metadata about the remote group. Include only for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details on how to specify this field. The required format is dependent on group_type and should have the following schema: <style type=\"text/css\"> code {max-height:300px !important} </style> ```json {   \"$schema\": \"http://json-schema.org/draft-04/schema#\",   \"title\": \"Group Metadata\",   \"properties\": {     \"ad_group\": {       \"properties\": {         \"object_guid\": {           \"type\": \"string\"         }       },       \"required\": [\"object_guid\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"Active Directory Group\"     },     \"duo_group\": {       \"properties\": {         \"group_id\": {           \"type\": \"string\"         }       },       \"required\": [\"group_id\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"Duo Group\"     },     \"git_hub_team\": {       \"properties\": {         \"org_name\": {           \"type\": \"string\"         },         \"team_slug\": {           \"type\": \"string\"         }       },       \"required\": [\"org_name\", \"team_slug\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"GitHub Team\"     },     \"google_groups_group\": {       \"properties\": {         \"group_id\": {           \"type\": \"string\"         }       },       \"required\": [\"group_id\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"Google Groups Group\"     },     \"ldap_group\": {       \"properties\": {         \"group_uid\": {           \"type\": \"string\"         }       },       \"required\": [\"group_uid\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"LDAP Group\"     },     \"okta_directory_group\": {       \"properties\": {         \"group_id\": {           \"type\": \"string\"         }       },       \"required\": [\"group_id\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"Okta Directory Group\"     }   },   \"additionalProperties\": false,   \"minProperties\": 1,   \"maxProperties\": 1,   \"type\": \"object\" } ```
	// Deprecated
	Metadata *string `json:"metadata,omitempty"`
	// Custom request notification sent upon request approval.
	CustomRequestNotification *string `json:"custom_request_notification,omitempty"`
	RiskSensitivityOverride *RiskSensitivityEnum `json:"risk_sensitivity_override,omitempty"`
}

type _CreateGroupInfo CreateGroupInfo

// NewCreateGroupInfo instantiates a new CreateGroupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupInfo(name string, groupType GroupTypeEnum, appId string) *CreateGroupInfo {
	this := CreateGroupInfo{}
	this.Name = name
	this.GroupType = groupType
	this.AppId = appId
	return &this
}

// NewCreateGroupInfoWithDefaults instantiates a new CreateGroupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupInfoWithDefaults() *CreateGroupInfo {
	this := CreateGroupInfo{}
	return &this
}

// GetName returns the Name field value
func (o *CreateGroupInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGroupInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGroupInfo) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateGroupInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGroupInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateGroupInfo) SetDescription(v string) {
	o.Description = &v
}

// GetGroupType returns the GroupType field value
func (o *CreateGroupInfo) GetGroupType() GroupTypeEnum {
	if o == nil {
		var ret GroupTypeEnum
		return ret
	}

	return o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value
// and a boolean to check if the value has been set.
func (o *CreateGroupInfo) GetGroupTypeOk() (*GroupTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupType, true
}

// SetGroupType sets field value
func (o *CreateGroupInfo) SetGroupType(v GroupTypeEnum) {
	o.GroupType = v
}

// GetAppId returns the AppId field value
func (o *CreateGroupInfo) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *CreateGroupInfo) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *CreateGroupInfo) SetAppId(v string) {
	o.AppId = v
}

// GetRemoteInfo returns the RemoteInfo field value if set, zero value otherwise.
func (o *CreateGroupInfo) GetRemoteInfo() GroupRemoteInfo {
	if o == nil || IsNil(o.RemoteInfo) {
		var ret GroupRemoteInfo
		return ret
	}
	return *o.RemoteInfo
}

// GetRemoteInfoOk returns a tuple with the RemoteInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupInfo) GetRemoteInfoOk() (*GroupRemoteInfo, bool) {
	if o == nil || IsNil(o.RemoteInfo) {
		return nil, false
	}
	return o.RemoteInfo, true
}

// HasRemoteInfo returns a boolean if a field has been set.
func (o *CreateGroupInfo) HasRemoteInfo() bool {
	if o != nil && !IsNil(o.RemoteInfo) {
		return true
	}

	return false
}

// SetRemoteInfo gets a reference to the given GroupRemoteInfo and assigns it to the RemoteInfo field.
func (o *CreateGroupInfo) SetRemoteInfo(v GroupRemoteInfo) {
	o.RemoteInfo = &v
}

// GetRemoteGroupId returns the RemoteGroupId field value if set, zero value otherwise.
// Deprecated
func (o *CreateGroupInfo) GetRemoteGroupId() string {
	if o == nil || IsNil(o.RemoteGroupId) {
		var ret string
		return ret
	}
	return *o.RemoteGroupId
}

// GetRemoteGroupIdOk returns a tuple with the RemoteGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateGroupInfo) GetRemoteGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteGroupId) {
		return nil, false
	}
	return o.RemoteGroupId, true
}

// HasRemoteGroupId returns a boolean if a field has been set.
func (o *CreateGroupInfo) HasRemoteGroupId() bool {
	if o != nil && !IsNil(o.RemoteGroupId) {
		return true
	}

	return false
}

// SetRemoteGroupId gets a reference to the given string and assigns it to the RemoteGroupId field.
// Deprecated
func (o *CreateGroupInfo) SetRemoteGroupId(v string) {
	o.RemoteGroupId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
// Deprecated
func (o *CreateGroupInfo) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateGroupInfo) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateGroupInfo) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
// Deprecated
func (o *CreateGroupInfo) SetMetadata(v string) {
	o.Metadata = &v
}

// GetCustomRequestNotification returns the CustomRequestNotification field value if set, zero value otherwise.
func (o *CreateGroupInfo) GetCustomRequestNotification() string {
	if o == nil || IsNil(o.CustomRequestNotification) {
		var ret string
		return ret
	}
	return *o.CustomRequestNotification
}

// GetCustomRequestNotificationOk returns a tuple with the CustomRequestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupInfo) GetCustomRequestNotificationOk() (*string, bool) {
	if o == nil || IsNil(o.CustomRequestNotification) {
		return nil, false
	}
	return o.CustomRequestNotification, true
}

// HasCustomRequestNotification returns a boolean if a field has been set.
func (o *CreateGroupInfo) HasCustomRequestNotification() bool {
	if o != nil && !IsNil(o.CustomRequestNotification) {
		return true
	}

	return false
}

// SetCustomRequestNotification gets a reference to the given string and assigns it to the CustomRequestNotification field.
func (o *CreateGroupInfo) SetCustomRequestNotification(v string) {
	o.CustomRequestNotification = &v
}

// GetRiskSensitivityOverride returns the RiskSensitivityOverride field value if set, zero value otherwise.
func (o *CreateGroupInfo) GetRiskSensitivityOverride() RiskSensitivityEnum {
	if o == nil || IsNil(o.RiskSensitivityOverride) {
		var ret RiskSensitivityEnum
		return ret
	}
	return *o.RiskSensitivityOverride
}

// GetRiskSensitivityOverrideOk returns a tuple with the RiskSensitivityOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupInfo) GetRiskSensitivityOverrideOk() (*RiskSensitivityEnum, bool) {
	if o == nil || IsNil(o.RiskSensitivityOverride) {
		return nil, false
	}
	return o.RiskSensitivityOverride, true
}

// HasRiskSensitivityOverride returns a boolean if a field has been set.
func (o *CreateGroupInfo) HasRiskSensitivityOverride() bool {
	if o != nil && !IsNil(o.RiskSensitivityOverride) {
		return true
	}

	return false
}

// SetRiskSensitivityOverride gets a reference to the given RiskSensitivityEnum and assigns it to the RiskSensitivityOverride field.
func (o *CreateGroupInfo) SetRiskSensitivityOverride(v RiskSensitivityEnum) {
	o.RiskSensitivityOverride = &v
}

func (o CreateGroupInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["group_type"] = o.GroupType
	toSerialize["app_id"] = o.AppId
	if !IsNil(o.RemoteInfo) {
		toSerialize["remote_info"] = o.RemoteInfo
	}
	if !IsNil(o.RemoteGroupId) {
		toSerialize["remote_group_id"] = o.RemoteGroupId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.CustomRequestNotification) {
		toSerialize["custom_request_notification"] = o.CustomRequestNotification
	}
	if !IsNil(o.RiskSensitivityOverride) {
		toSerialize["risk_sensitivity_override"] = o.RiskSensitivityOverride
	}
	return toSerialize, nil
}

func (o *CreateGroupInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"group_type",
		"app_id",
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

	varCreateGroupInfo := _CreateGroupInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGroupInfo)

	if err != nil {
		return err
	}

	*o = CreateGroupInfo(varCreateGroupInfo)

	return err
}

type NullableCreateGroupInfo struct {
	value *CreateGroupInfo
	isSet bool
}

func (v NullableCreateGroupInfo) Get() *CreateGroupInfo {
	return v.value
}

func (v *NullableCreateGroupInfo) Set(val *CreateGroupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupInfo(val *CreateGroupInfo) *NullableCreateGroupInfo {
	return &NullableCreateGroupInfo{value: val, isSet: true}
}

func (v NullableCreateGroupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


