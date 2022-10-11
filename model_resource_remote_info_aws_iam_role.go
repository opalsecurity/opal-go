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

// ResourceRemoteInfoAwsIamRole Remote info for AWS IAM role.
type ResourceRemoteInfoAwsIamRole struct {
	// The ARN of the IAM role.
	Arn string `json:"arn"`
}

// NewResourceRemoteInfoAwsIamRole instantiates a new ResourceRemoteInfoAwsIamRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoAwsIamRole(arn string) *ResourceRemoteInfoAwsIamRole {
	this := ResourceRemoteInfoAwsIamRole{}
	this.Arn = arn
	return &this
}

// NewResourceRemoteInfoAwsIamRoleWithDefaults instantiates a new ResourceRemoteInfoAwsIamRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoAwsIamRoleWithDefaults() *ResourceRemoteInfoAwsIamRole {
	this := ResourceRemoteInfoAwsIamRole{}
	return &this
}

// GetArn returns the Arn field value
func (o *ResourceRemoteInfoAwsIamRole) GetArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arn
}

// GetArnOk returns a tuple with the Arn field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoAwsIamRole) GetArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arn, true
}

// SetArn sets field value
func (o *ResourceRemoteInfoAwsIamRole) SetArn(v string) {
	o.Arn = v
}

func (o ResourceRemoteInfoAwsIamRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["arn"] = o.Arn
	}
	return json.Marshal(toSerialize)
}

type NullableResourceRemoteInfoAwsIamRole struct {
	value *ResourceRemoteInfoAwsIamRole
	isSet bool
}

func (v NullableResourceRemoteInfoAwsIamRole) Get() *ResourceRemoteInfoAwsIamRole {
	return v.value
}

func (v *NullableResourceRemoteInfoAwsIamRole) Set(val *ResourceRemoteInfoAwsIamRole) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoAwsIamRole) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoAwsIamRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoAwsIamRole(val *ResourceRemoteInfoAwsIamRole) *NullableResourceRemoteInfoAwsIamRole {
	return &NullableResourceRemoteInfoAwsIamRole{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoAwsIamRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoAwsIamRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

