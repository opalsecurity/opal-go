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

// checks if the TicketPropagationConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TicketPropagationConfiguration{}

// TicketPropagationConfiguration Configuration for ticket propagation, when enabled, a ticket will be created for access changes related to the users in this resource.
type TicketPropagationConfiguration struct {
	EnabledOnGrant bool `json:"enabled_on_grant"`
	EnabledOnRevocation bool `json:"enabled_on_revocation"`
	TicketProvider *TicketingProviderEnum `json:"ticket_provider,omitempty"`
	TicketProjectId *string `json:"ticket_project_id,omitempty"`
}

type _TicketPropagationConfiguration TicketPropagationConfiguration

// NewTicketPropagationConfiguration instantiates a new TicketPropagationConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketPropagationConfiguration(enabledOnGrant bool, enabledOnRevocation bool) *TicketPropagationConfiguration {
	this := TicketPropagationConfiguration{}
	this.EnabledOnGrant = enabledOnGrant
	this.EnabledOnRevocation = enabledOnRevocation
	return &this
}

// NewTicketPropagationConfigurationWithDefaults instantiates a new TicketPropagationConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketPropagationConfigurationWithDefaults() *TicketPropagationConfiguration {
	this := TicketPropagationConfiguration{}
	return &this
}

// GetEnabledOnGrant returns the EnabledOnGrant field value
func (o *TicketPropagationConfiguration) GetEnabledOnGrant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnabledOnGrant
}

// GetEnabledOnGrantOk returns a tuple with the EnabledOnGrant field value
// and a boolean to check if the value has been set.
func (o *TicketPropagationConfiguration) GetEnabledOnGrantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledOnGrant, true
}

// SetEnabledOnGrant sets field value
func (o *TicketPropagationConfiguration) SetEnabledOnGrant(v bool) {
	o.EnabledOnGrant = v
}

// GetEnabledOnRevocation returns the EnabledOnRevocation field value
func (o *TicketPropagationConfiguration) GetEnabledOnRevocation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnabledOnRevocation
}

// GetEnabledOnRevocationOk returns a tuple with the EnabledOnRevocation field value
// and a boolean to check if the value has been set.
func (o *TicketPropagationConfiguration) GetEnabledOnRevocationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledOnRevocation, true
}

// SetEnabledOnRevocation sets field value
func (o *TicketPropagationConfiguration) SetEnabledOnRevocation(v bool) {
	o.EnabledOnRevocation = v
}

// GetTicketProvider returns the TicketProvider field value if set, zero value otherwise.
func (o *TicketPropagationConfiguration) GetTicketProvider() TicketingProviderEnum {
	if o == nil || IsNil(o.TicketProvider) {
		var ret TicketingProviderEnum
		return ret
	}
	return *o.TicketProvider
}

// GetTicketProviderOk returns a tuple with the TicketProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketPropagationConfiguration) GetTicketProviderOk() (*TicketingProviderEnum, bool) {
	if o == nil || IsNil(o.TicketProvider) {
		return nil, false
	}
	return o.TicketProvider, true
}

// HasTicketProvider returns a boolean if a field has been set.
func (o *TicketPropagationConfiguration) HasTicketProvider() bool {
	if o != nil && !IsNil(o.TicketProvider) {
		return true
	}

	return false
}

// SetTicketProvider gets a reference to the given TicketingProviderEnum and assigns it to the TicketProvider field.
func (o *TicketPropagationConfiguration) SetTicketProvider(v TicketingProviderEnum) {
	o.TicketProvider = &v
}

// GetTicketProjectId returns the TicketProjectId field value if set, zero value otherwise.
func (o *TicketPropagationConfiguration) GetTicketProjectId() string {
	if o == nil || IsNil(o.TicketProjectId) {
		var ret string
		return ret
	}
	return *o.TicketProjectId
}

// GetTicketProjectIdOk returns a tuple with the TicketProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketPropagationConfiguration) GetTicketProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.TicketProjectId) {
		return nil, false
	}
	return o.TicketProjectId, true
}

// HasTicketProjectId returns a boolean if a field has been set.
func (o *TicketPropagationConfiguration) HasTicketProjectId() bool {
	if o != nil && !IsNil(o.TicketProjectId) {
		return true
	}

	return false
}

// SetTicketProjectId gets a reference to the given string and assigns it to the TicketProjectId field.
func (o *TicketPropagationConfiguration) SetTicketProjectId(v string) {
	o.TicketProjectId = &v
}

func (o TicketPropagationConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TicketPropagationConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled_on_grant"] = o.EnabledOnGrant
	toSerialize["enabled_on_revocation"] = o.EnabledOnRevocation
	if !IsNil(o.TicketProvider) {
		toSerialize["ticket_provider"] = o.TicketProvider
	}
	if !IsNil(o.TicketProjectId) {
		toSerialize["ticket_project_id"] = o.TicketProjectId
	}
	return toSerialize, nil
}

func (o *TicketPropagationConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled_on_grant",
		"enabled_on_revocation",
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

	varTicketPropagationConfiguration := _TicketPropagationConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTicketPropagationConfiguration)

	if err != nil {
		return err
	}

	*o = TicketPropagationConfiguration(varTicketPropagationConfiguration)

	return err
}

type NullableTicketPropagationConfiguration struct {
	value *TicketPropagationConfiguration
	isSet bool
}

func (v NullableTicketPropagationConfiguration) Get() *TicketPropagationConfiguration {
	return v.value
}

func (v *NullableTicketPropagationConfiguration) Set(val *TicketPropagationConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketPropagationConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketPropagationConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketPropagationConfiguration(val *TicketPropagationConfiguration) *NullableTicketPropagationConfiguration {
	return &NullableTicketPropagationConfiguration{value: val, isSet: true}
}

func (v NullableTicketPropagationConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketPropagationConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


