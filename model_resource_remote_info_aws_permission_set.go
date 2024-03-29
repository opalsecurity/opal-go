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

// checks if the ResourceRemoteInfoAwsPermissionSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoAwsPermissionSet{}

// ResourceRemoteInfoAwsPermissionSet Remote info for AWS Identity Center permission set.
type ResourceRemoteInfoAwsPermissionSet struct {
	// The ARN of the permission set.
	Arn string `json:"arn"`
	// The ID of an AWS account to which this permission set is provisioned.
	AccountId string `json:"account_id"`
}

// NewResourceRemoteInfoAwsPermissionSet instantiates a new ResourceRemoteInfoAwsPermissionSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoAwsPermissionSet(arn string, accountId string) *ResourceRemoteInfoAwsPermissionSet {
	this := ResourceRemoteInfoAwsPermissionSet{}
	this.Arn = arn
	this.AccountId = accountId
	return &this
}

// NewResourceRemoteInfoAwsPermissionSetWithDefaults instantiates a new ResourceRemoteInfoAwsPermissionSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoAwsPermissionSetWithDefaults() *ResourceRemoteInfoAwsPermissionSet {
	this := ResourceRemoteInfoAwsPermissionSet{}
	return &this
}

// GetArn returns the Arn field value
func (o *ResourceRemoteInfoAwsPermissionSet) GetArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arn
}

// GetArnOk returns a tuple with the Arn field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoAwsPermissionSet) GetArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arn, true
}

// SetArn sets field value
func (o *ResourceRemoteInfoAwsPermissionSet) SetArn(v string) {
	o.Arn = v
}

// GetAccountId returns the AccountId field value
func (o *ResourceRemoteInfoAwsPermissionSet) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoAwsPermissionSet) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ResourceRemoteInfoAwsPermissionSet) SetAccountId(v string) {
	o.AccountId = v
}

func (o ResourceRemoteInfoAwsPermissionSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoAwsPermissionSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["arn"] = o.Arn
	toSerialize["account_id"] = o.AccountId
	return toSerialize, nil
}

type NullableResourceRemoteInfoAwsPermissionSet struct {
	value *ResourceRemoteInfoAwsPermissionSet
	isSet bool
}

func (v NullableResourceRemoteInfoAwsPermissionSet) Get() *ResourceRemoteInfoAwsPermissionSet {
	return v.value
}

func (v *NullableResourceRemoteInfoAwsPermissionSet) Set(val *ResourceRemoteInfoAwsPermissionSet) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoAwsPermissionSet) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoAwsPermissionSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoAwsPermissionSet(val *ResourceRemoteInfoAwsPermissionSet) *NullableResourceRemoteInfoAwsPermissionSet {
	return &NullableResourceRemoteInfoAwsPermissionSet{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoAwsPermissionSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoAwsPermissionSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


