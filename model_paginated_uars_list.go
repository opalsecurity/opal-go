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

// checks if the PaginatedUARsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedUARsList{}

// PaginatedUARsList A list of UARs.
type PaginatedUARsList struct {
	// The cursor with which to continue pagination if additional result pages exist.
	Next *string `json:"next,omitempty"`
	// The cursor used to obtain the current result page.
	Previous *string `json:"previous,omitempty"`
	Results []UAR `json:"results"`
}

type _PaginatedUARsList PaginatedUARsList

// NewPaginatedUARsList instantiates a new PaginatedUARsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUARsList(results []UAR) *PaginatedUARsList {
	this := PaginatedUARsList{}
	this.Results = results
	return &this
}

// NewPaginatedUARsListWithDefaults instantiates a new PaginatedUARsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUARsListWithDefaults() *PaginatedUARsList {
	this := PaginatedUARsList{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedUARsList) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedUARsList) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedUARsList) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedUARsList) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *PaginatedUARsList) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedUARsList) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedUARsList) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *PaginatedUARsList) SetPrevious(v string) {
	o.Previous = &v
}

// GetResults returns the Results field value
func (o *PaginatedUARsList) GetResults() []UAR {
	if o == nil {
		var ret []UAR
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUARsList) GetResultsOk() ([]UAR, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedUARsList) SetResults(v []UAR) {
	o.Results = v
}

func (o PaginatedUARsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedUARsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PaginatedUARsList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varPaginatedUARsList := _PaginatedUARsList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginatedUARsList)

	if err != nil {
		return err
	}

	*o = PaginatedUARsList(varPaginatedUARsList)

	return err
}

type NullablePaginatedUARsList struct {
	value *PaginatedUARsList
	isSet bool
}

func (v NullablePaginatedUARsList) Get() *PaginatedUARsList {
	return v.value
}

func (v *NullablePaginatedUARsList) Set(val *PaginatedUARsList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUARsList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUARsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUARsList(val *PaginatedUARsList) *NullablePaginatedUARsList {
	return &NullablePaginatedUARsList{value: val, isSet: true}
}

func (v NullablePaginatedUARsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUARsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


