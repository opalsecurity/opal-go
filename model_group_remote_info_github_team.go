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

// GroupRemoteInfoGithubTeam Remote info for GitHub team.
type GroupRemoteInfoGithubTeam struct {
	// The id of the GitHub team.
	// Deprecated
	TeamId *string `json:"team_id,omitempty"`
	// The slug of the GitHub team.
	TeamSlug string `json:"team_slug"`
}

// NewGroupRemoteInfoGithubTeam instantiates a new GroupRemoteInfoGithubTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRemoteInfoGithubTeam(teamSlug string) *GroupRemoteInfoGithubTeam {
	this := GroupRemoteInfoGithubTeam{}
	this.TeamSlug = teamSlug
	return &this
}

// NewGroupRemoteInfoGithubTeamWithDefaults instantiates a new GroupRemoteInfoGithubTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRemoteInfoGithubTeamWithDefaults() *GroupRemoteInfoGithubTeam {
	this := GroupRemoteInfoGithubTeam{}
	return &this
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
// Deprecated
func (o *GroupRemoteInfoGithubTeam) GetTeamId() string {
	if o == nil || isNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GroupRemoteInfoGithubTeam) GetTeamIdOk() (*string, bool) {
	if o == nil || isNil(o.TeamId) {
    return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *GroupRemoteInfoGithubTeam) HasTeamId() bool {
	if o != nil && !isNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
// Deprecated
func (o *GroupRemoteInfoGithubTeam) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTeamSlug returns the TeamSlug field value
func (o *GroupRemoteInfoGithubTeam) GetTeamSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamSlug
}

// GetTeamSlugOk returns a tuple with the TeamSlug field value
// and a boolean to check if the value has been set.
func (o *GroupRemoteInfoGithubTeam) GetTeamSlugOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TeamSlug, true
}

// SetTeamSlug sets field value
func (o *GroupRemoteInfoGithubTeam) SetTeamSlug(v string) {
	o.TeamSlug = v
}

func (o GroupRemoteInfoGithubTeam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TeamId) {
		toSerialize["team_id"] = o.TeamId
	}
	if true {
		toSerialize["team_slug"] = o.TeamSlug
	}
	return json.Marshal(toSerialize)
}

type NullableGroupRemoteInfoGithubTeam struct {
	value *GroupRemoteInfoGithubTeam
	isSet bool
}

func (v NullableGroupRemoteInfoGithubTeam) Get() *GroupRemoteInfoGithubTeam {
	return v.value
}

func (v *NullableGroupRemoteInfoGithubTeam) Set(val *GroupRemoteInfoGithubTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRemoteInfoGithubTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRemoteInfoGithubTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRemoteInfoGithubTeam(val *GroupRemoteInfoGithubTeam) *NullableGroupRemoteInfoGithubTeam {
	return &NullableGroupRemoteInfoGithubTeam{value: val, isSet: true}
}

func (v NullableGroupRemoteInfoGithubTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRemoteInfoGithubTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


