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

// checks if the GroupRemoteInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRemoteInfo{}

// GroupRemoteInfo Remote info that is required for the creation of remote groups. Providing this will allow you to omit deprecated remote_id and metadata fields.
type GroupRemoteInfo struct {
	ActiveDirectoryGroup *GroupRemoteInfoActiveDirectoryGroup `json:"active_directory_group,omitempty"`
	GithubTeam *GroupRemoteInfoGithubTeam `json:"github_team,omitempty"`
	GitlabGroup *GroupRemoteInfoGitlabGroup `json:"gitlab_group,omitempty"`
	GoogleGroup *GroupRemoteInfoGoogleGroup `json:"google_group,omitempty"`
	LdapGroup *GroupRemoteInfoLdapGroup `json:"ldap_group,omitempty"`
	OktaGroup *GroupRemoteInfoOktaGroup `json:"okta_group,omitempty"`
	DuoGroup *GroupRemoteInfoDuoGroup `json:"duo_group,omitempty"`
}

// NewGroupRemoteInfo instantiates a new GroupRemoteInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRemoteInfo() *GroupRemoteInfo {
	this := GroupRemoteInfo{}
	return &this
}

// NewGroupRemoteInfoWithDefaults instantiates a new GroupRemoteInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRemoteInfoWithDefaults() *GroupRemoteInfo {
	this := GroupRemoteInfo{}
	return &this
}

// GetActiveDirectoryGroup returns the ActiveDirectoryGroup field value if set, zero value otherwise.
func (o *GroupRemoteInfo) GetActiveDirectoryGroup() GroupRemoteInfoActiveDirectoryGroup {
	if o == nil || IsNil(o.ActiveDirectoryGroup) {
		var ret GroupRemoteInfoActiveDirectoryGroup
		return ret
	}
	return *o.ActiveDirectoryGroup
}

// GetActiveDirectoryGroupOk returns a tuple with the ActiveDirectoryGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfo) GetActiveDirectoryGroupOk() (*GroupRemoteInfoActiveDirectoryGroup, bool) {
	if o == nil || IsNil(o.ActiveDirectoryGroup) {
		return nil, false
	}
	return o.ActiveDirectoryGroup, true
}

// HasActiveDirectoryGroup returns a boolean if a field has been set.
func (o *GroupRemoteInfo) HasActiveDirectoryGroup() bool {
	if o != nil && !IsNil(o.ActiveDirectoryGroup) {
		return true
	}

	return false
}

// SetActiveDirectoryGroup gets a reference to the given GroupRemoteInfoActiveDirectoryGroup and assigns it to the ActiveDirectoryGroup field.
func (o *GroupRemoteInfo) SetActiveDirectoryGroup(v GroupRemoteInfoActiveDirectoryGroup) {
	o.ActiveDirectoryGroup = &v
}

// GetGithubTeam returns the GithubTeam field value if set, zero value otherwise.
func (o *GroupRemoteInfo) GetGithubTeam() GroupRemoteInfoGithubTeam {
	if o == nil || IsNil(o.GithubTeam) {
		var ret GroupRemoteInfoGithubTeam
		return ret
	}
	return *o.GithubTeam
}

// GetGithubTeamOk returns a tuple with the GithubTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfo) GetGithubTeamOk() (*GroupRemoteInfoGithubTeam, bool) {
	if o == nil || IsNil(o.GithubTeam) {
		return nil, false
	}
	return o.GithubTeam, true
}

// HasGithubTeam returns a boolean if a field has been set.
func (o *GroupRemoteInfo) HasGithubTeam() bool {
	if o != nil && !IsNil(o.GithubTeam) {
		return true
	}

	return false
}

// SetGithubTeam gets a reference to the given GroupRemoteInfoGithubTeam and assigns it to the GithubTeam field.
func (o *GroupRemoteInfo) SetGithubTeam(v GroupRemoteInfoGithubTeam) {
	o.GithubTeam = &v
}

// GetGitlabGroup returns the GitlabGroup field value if set, zero value otherwise.
func (o *GroupRemoteInfo) GetGitlabGroup() GroupRemoteInfoGitlabGroup {
	if o == nil || IsNil(o.GitlabGroup) {
		var ret GroupRemoteInfoGitlabGroup
		return ret
	}
	return *o.GitlabGroup
}

// GetGitlabGroupOk returns a tuple with the GitlabGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfo) GetGitlabGroupOk() (*GroupRemoteInfoGitlabGroup, bool) {
	if o == nil || IsNil(o.GitlabGroup) {
		return nil, false
	}
	return o.GitlabGroup, true
}

// HasGitlabGroup returns a boolean if a field has been set.
func (o *GroupRemoteInfo) HasGitlabGroup() bool {
	if o != nil && !IsNil(o.GitlabGroup) {
		return true
	}

	return false
}

// SetGitlabGroup gets a reference to the given GroupRemoteInfoGitlabGroup and assigns it to the GitlabGroup field.
func (o *GroupRemoteInfo) SetGitlabGroup(v GroupRemoteInfoGitlabGroup) {
	o.GitlabGroup = &v
}

// GetGoogleGroup returns the GoogleGroup field value if set, zero value otherwise.
func (o *GroupRemoteInfo) GetGoogleGroup() GroupRemoteInfoGoogleGroup {
	if o == nil || IsNil(o.GoogleGroup) {
		var ret GroupRemoteInfoGoogleGroup
		return ret
	}
	return *o.GoogleGroup
}

// GetGoogleGroupOk returns a tuple with the GoogleGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfo) GetGoogleGroupOk() (*GroupRemoteInfoGoogleGroup, bool) {
	if o == nil || IsNil(o.GoogleGroup) {
		return nil, false
	}
	return o.GoogleGroup, true
}

// HasGoogleGroup returns a boolean if a field has been set.
func (o *GroupRemoteInfo) HasGoogleGroup() bool {
	if o != nil && !IsNil(o.GoogleGroup) {
		return true
	}

	return false
}

// SetGoogleGroup gets a reference to the given GroupRemoteInfoGoogleGroup and assigns it to the GoogleGroup field.
func (o *GroupRemoteInfo) SetGoogleGroup(v GroupRemoteInfoGoogleGroup) {
	o.GoogleGroup = &v
}

// GetLdapGroup returns the LdapGroup field value if set, zero value otherwise.
func (o *GroupRemoteInfo) GetLdapGroup() GroupRemoteInfoLdapGroup {
	if o == nil || IsNil(o.LdapGroup) {
		var ret GroupRemoteInfoLdapGroup
		return ret
	}
	return *o.LdapGroup
}

// GetLdapGroupOk returns a tuple with the LdapGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfo) GetLdapGroupOk() (*GroupRemoteInfoLdapGroup, bool) {
	if o == nil || IsNil(o.LdapGroup) {
		return nil, false
	}
	return o.LdapGroup, true
}

// HasLdapGroup returns a boolean if a field has been set.
func (o *GroupRemoteInfo) HasLdapGroup() bool {
	if o != nil && !IsNil(o.LdapGroup) {
		return true
	}

	return false
}

// SetLdapGroup gets a reference to the given GroupRemoteInfoLdapGroup and assigns it to the LdapGroup field.
func (o *GroupRemoteInfo) SetLdapGroup(v GroupRemoteInfoLdapGroup) {
	o.LdapGroup = &v
}

// GetOktaGroup returns the OktaGroup field value if set, zero value otherwise.
func (o *GroupRemoteInfo) GetOktaGroup() GroupRemoteInfoOktaGroup {
	if o == nil || IsNil(o.OktaGroup) {
		var ret GroupRemoteInfoOktaGroup
		return ret
	}
	return *o.OktaGroup
}

// GetOktaGroupOk returns a tuple with the OktaGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfo) GetOktaGroupOk() (*GroupRemoteInfoOktaGroup, bool) {
	if o == nil || IsNil(o.OktaGroup) {
		return nil, false
	}
	return o.OktaGroup, true
}

// HasOktaGroup returns a boolean if a field has been set.
func (o *GroupRemoteInfo) HasOktaGroup() bool {
	if o != nil && !IsNil(o.OktaGroup) {
		return true
	}

	return false
}

// SetOktaGroup gets a reference to the given GroupRemoteInfoOktaGroup and assigns it to the OktaGroup field.
func (o *GroupRemoteInfo) SetOktaGroup(v GroupRemoteInfoOktaGroup) {
	o.OktaGroup = &v
}

// GetDuoGroup returns the DuoGroup field value if set, zero value otherwise.
func (o *GroupRemoteInfo) GetDuoGroup() GroupRemoteInfoDuoGroup {
	if o == nil || IsNil(o.DuoGroup) {
		var ret GroupRemoteInfoDuoGroup
		return ret
	}
	return *o.DuoGroup
}

// GetDuoGroupOk returns a tuple with the DuoGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfo) GetDuoGroupOk() (*GroupRemoteInfoDuoGroup, bool) {
	if o == nil || IsNil(o.DuoGroup) {
		return nil, false
	}
	return o.DuoGroup, true
}

// HasDuoGroup returns a boolean if a field has been set.
func (o *GroupRemoteInfo) HasDuoGroup() bool {
	if o != nil && !IsNil(o.DuoGroup) {
		return true
	}

	return false
}

// SetDuoGroup gets a reference to the given GroupRemoteInfoDuoGroup and assigns it to the DuoGroup field.
func (o *GroupRemoteInfo) SetDuoGroup(v GroupRemoteInfoDuoGroup) {
	o.DuoGroup = &v
}

func (o GroupRemoteInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRemoteInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveDirectoryGroup) {
		toSerialize["active_directory_group"] = o.ActiveDirectoryGroup
	}
	if !IsNil(o.GithubTeam) {
		toSerialize["github_team"] = o.GithubTeam
	}
	if !IsNil(o.GitlabGroup) {
		toSerialize["gitlab_group"] = o.GitlabGroup
	}
	if !IsNil(o.GoogleGroup) {
		toSerialize["google_group"] = o.GoogleGroup
	}
	if !IsNil(o.LdapGroup) {
		toSerialize["ldap_group"] = o.LdapGroup
	}
	if !IsNil(o.OktaGroup) {
		toSerialize["okta_group"] = o.OktaGroup
	}
	if !IsNil(o.DuoGroup) {
		toSerialize["duo_group"] = o.DuoGroup
	}
	return toSerialize, nil
}

type NullableGroupRemoteInfo struct {
	value *GroupRemoteInfo
	isSet bool
}

func (v NullableGroupRemoteInfo) Get() *GroupRemoteInfo {
	return v.value
}

func (v *NullableGroupRemoteInfo) Set(val *GroupRemoteInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRemoteInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRemoteInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRemoteInfo(val *GroupRemoteInfo) *NullableGroupRemoteInfo {
	return &NullableGroupRemoteInfo{value: val, isSet: true}
}

func (v NullableGroupRemoteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRemoteInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


