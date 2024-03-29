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

// checks if the ResourceRemoteInfoSalesforceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoSalesforceProfile{}

// ResourceRemoteInfoSalesforceProfile Remote info for Salesforce profile.
type ResourceRemoteInfoSalesforceProfile struct {
	// The id of the permission set.
	ProfileId string `json:"profile_id"`
	// The id of the user license.
	UserLicenseId string `json:"user_license_id"`
}

// NewResourceRemoteInfoSalesforceProfile instantiates a new ResourceRemoteInfoSalesforceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoSalesforceProfile(profileId string, userLicenseId string) *ResourceRemoteInfoSalesforceProfile {
	this := ResourceRemoteInfoSalesforceProfile{}
	this.ProfileId = profileId
	this.UserLicenseId = userLicenseId
	return &this
}

// NewResourceRemoteInfoSalesforceProfileWithDefaults instantiates a new ResourceRemoteInfoSalesforceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoSalesforceProfileWithDefaults() *ResourceRemoteInfoSalesforceProfile {
	this := ResourceRemoteInfoSalesforceProfile{}
	return &this
}

// GetProfileId returns the ProfileId field value
func (o *ResourceRemoteInfoSalesforceProfile) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoSalesforceProfile) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *ResourceRemoteInfoSalesforceProfile) SetProfileId(v string) {
	o.ProfileId = v
}

// GetUserLicenseId returns the UserLicenseId field value
func (o *ResourceRemoteInfoSalesforceProfile) GetUserLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserLicenseId
}

// GetUserLicenseIdOk returns a tuple with the UserLicenseId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoSalesforceProfile) GetUserLicenseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserLicenseId, true
}

// SetUserLicenseId sets field value
func (o *ResourceRemoteInfoSalesforceProfile) SetUserLicenseId(v string) {
	o.UserLicenseId = v
}

func (o ResourceRemoteInfoSalesforceProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoSalesforceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profile_id"] = o.ProfileId
	toSerialize["user_license_id"] = o.UserLicenseId
	return toSerialize, nil
}

type NullableResourceRemoteInfoSalesforceProfile struct {
	value *ResourceRemoteInfoSalesforceProfile
	isSet bool
}

func (v NullableResourceRemoteInfoSalesforceProfile) Get() *ResourceRemoteInfoSalesforceProfile {
	return v.value
}

func (v *NullableResourceRemoteInfoSalesforceProfile) Set(val *ResourceRemoteInfoSalesforceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoSalesforceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoSalesforceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoSalesforceProfile(val *ResourceRemoteInfoSalesforceProfile) *NullableResourceRemoteInfoSalesforceProfile {
	return &NullableResourceRemoteInfoSalesforceProfile{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoSalesforceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoSalesforceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


