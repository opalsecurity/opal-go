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

// checks if the ResourceRemoteInfoAwsAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfoAwsAccount{}

// ResourceRemoteInfoAwsAccount Remote info for AWS account.
type ResourceRemoteInfoAwsAccount struct {
	// The id of the AWS account.
	AccountId string `json:"account_id"`
}

type _ResourceRemoteInfoAwsAccount ResourceRemoteInfoAwsAccount

// NewResourceRemoteInfoAwsAccount instantiates a new ResourceRemoteInfoAwsAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfoAwsAccount(accountId string) *ResourceRemoteInfoAwsAccount {
	this := ResourceRemoteInfoAwsAccount{}
	this.AccountId = accountId
	return &this
}

// NewResourceRemoteInfoAwsAccountWithDefaults instantiates a new ResourceRemoteInfoAwsAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoAwsAccountWithDefaults() *ResourceRemoteInfoAwsAccount {
	this := ResourceRemoteInfoAwsAccount{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *ResourceRemoteInfoAwsAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfoAwsAccount) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ResourceRemoteInfoAwsAccount) SetAccountId(v string) {
	o.AccountId = v
}

func (o ResourceRemoteInfoAwsAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfoAwsAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	return toSerialize, nil
}

func (o *ResourceRemoteInfoAwsAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
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

	varResourceRemoteInfoAwsAccount := _ResourceRemoteInfoAwsAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceRemoteInfoAwsAccount)

	if err != nil {
		return err
	}

	*o = ResourceRemoteInfoAwsAccount(varResourceRemoteInfoAwsAccount)

	return err
}

type NullableResourceRemoteInfoAwsAccount struct {
	value *ResourceRemoteInfoAwsAccount
	isSet bool
}

func (v NullableResourceRemoteInfoAwsAccount) Get() *ResourceRemoteInfoAwsAccount {
	return v.value
}

func (v *NullableResourceRemoteInfoAwsAccount) Set(val *ResourceRemoteInfoAwsAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfoAwsAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfoAwsAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfoAwsAccount(val *ResourceRemoteInfoAwsAccount) *NullableResourceRemoteInfoAwsAccount {
	return &NullableResourceRemoteInfoAwsAccount{value: val, isSet: true}
}

func (v NullableResourceRemoteInfoAwsAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfoAwsAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


