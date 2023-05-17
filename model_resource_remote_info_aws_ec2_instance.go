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

// checks if the ResourceRemoteInfoAwsEc2Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoAwsEc2Instance{}

// ResourceRemoteInfoAwsEc2Instance Remote info for AWS EC2 instance.
type ResourceRemoteInfoAwsEc2Instance struct {
	// The instanceId of the EC2 instance.
	InstanceId string `json:"instance_id"`
	// The region of the EC2 instance.
	Region string `json:"region"`
	// The id of the AWS account. Required for AWS Organizations.
	AccountId *string `json:"account_id,omitempty"`
}

// NewResourceRemoteInfoAwsEc2Instance instantiates a new ResourceRemoteInfoAwsEc2Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoAwsEc2Instance(instanceId string, region string) *ResourceRemoteInfoAwsEc2Instance {
	this := ResourceRemoteInfoAwsEc2Instance{}
	this.InstanceId = instanceId
	this.Region = region
	return &this
}

// NewResourceRemoteInfoAwsEc2InstanceWithDefaults instantiates a new ResourceRemoteInfoAwsEc2Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoAwsEc2InstanceWithDefaults() *ResourceRemoteInfoAwsEc2Instance {
	this := ResourceRemoteInfoAwsEc2Instance{}
	return &this
}

// GetInstanceId returns the InstanceId field value
func (o *ResourceRemoteInfoAwsEc2Instance) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoAwsEc2Instance) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *ResourceRemoteInfoAwsEc2Instance) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetRegion returns the Region field value
func (o *ResourceRemoteInfoAwsEc2Instance) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoAwsEc2Instance) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ResourceRemoteInfoAwsEc2Instance) SetRegion(v string) {
	o.Region = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ResourceRemoteInfoAwsEc2Instance) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoAwsEc2Instance) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ResourceRemoteInfoAwsEc2Instance) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ResourceRemoteInfoAwsEc2Instance) SetAccountId(v string) {
	o.AccountId = &v
}

func (o ResourceRemoteInfoAwsEc2Instance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoAwsEc2Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instance_id"] = o.InstanceId
	toSerialize["region"] = o.Region
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	return toSerialize, nil
}

type NullableResourceRemoteInfoAwsEc2Instance struct {
	value *ResourceRemoteInfoAwsEc2Instance
	isSet bool
}

func (v NullableResourceRemoteInfoAwsEc2Instance) Get() *ResourceRemoteInfoAwsEc2Instance {
	return v.value
}

func (v *NullableResourceRemoteInfoAwsEc2Instance) Set(val *ResourceRemoteInfoAwsEc2Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoAwsEc2Instance) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoAwsEc2Instance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoAwsEc2Instance(val *ResourceRemoteInfoAwsEc2Instance) *NullableResourceRemoteInfoAwsEc2Instance {
	return &NullableResourceRemoteInfoAwsEc2Instance{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoAwsEc2Instance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoAwsEc2Instance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


