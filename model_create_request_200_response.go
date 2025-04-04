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
)

// checks if the CreateRequest200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRequest200Response{}

// CreateRequest200Response struct for CreateRequest200Response
type CreateRequest200Response struct {
	Id *string `json:"id,omitempty"`
}

// NewCreateRequest200Response instantiates a new CreateRequest200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRequest200Response() *CreateRequest200Response {
	this := CreateRequest200Response{}
	return &this
}

// NewCreateRequest200ResponseWithDefaults instantiates a new CreateRequest200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRequest200ResponseWithDefaults() *CreateRequest200Response {
	this := CreateRequest200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateRequest200Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequest200Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateRequest200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateRequest200Response) SetId(v string) {
	o.Id = &v
}

func (o CreateRequest200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRequest200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateRequest200Response struct {
	value *CreateRequest200Response
	isSet bool
}

func (v NullableCreateRequest200Response) Get() *CreateRequest200Response {
	return v.value
}

func (v *NullableCreateRequest200Response) Set(val *CreateRequest200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRequest200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRequest200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRequest200Response(val *CreateRequest200Response) *NullableCreateRequest200Response {
	return &NullableCreateRequest200Response{value: val, isSet: true}
}

func (v NullableCreateRequest200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRequest200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


