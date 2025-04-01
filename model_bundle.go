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
	"time"
)

// checks if the Bundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bundle{}

// Bundle struct for Bundle
type Bundle struct {
	// The ID of the bundle.
	BundleId *string `json:"bundle_id,omitempty"`
	// The name of the bundle.
	Name *string `json:"name,omitempty"`
	// The description of the bundle.
	Description *string `json:"description,omitempty"`
	// The creation timestamp of the bundle, in ISO 8601 format
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last updated timestamp of the bundle, in ISO 8601 format
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The ID of the owner of the bundle.
	AdminOwnerId *string `json:"admin_owner_id,omitempty"`
	// The total number of items in the bundle.
	TotalNumItems *int32 `json:"total_num_items,omitempty"`
	// The total number of resources in the bundle.
	TotalNumResources *int32 `json:"total_num_resources,omitempty"`
	// The total number of groups in the bundle.
	TotalNumGroups *int32 `json:"total_num_groups,omitempty"`
}

// NewBundle instantiates a new Bundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundle() *Bundle {
	this := Bundle{}
	return &this
}

// NewBundleWithDefaults instantiates a new Bundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleWithDefaults() *Bundle {
	this := Bundle{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *Bundle) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *Bundle) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *Bundle) SetBundleId(v string) {
	o.BundleId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Bundle) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Bundle) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Bundle) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Bundle) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Bundle) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Bundle) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Bundle) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Bundle) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Bundle) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Bundle) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Bundle) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Bundle) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAdminOwnerId returns the AdminOwnerId field value if set, zero value otherwise.
func (o *Bundle) GetAdminOwnerId() string {
	if o == nil || IsNil(o.AdminOwnerId) {
		var ret string
		return ret
	}
	return *o.AdminOwnerId
}

// GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetAdminOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdminOwnerId) {
		return nil, false
	}
	return o.AdminOwnerId, true
}

// HasAdminOwnerId returns a boolean if a field has been set.
func (o *Bundle) HasAdminOwnerId() bool {
	if o != nil && !IsNil(o.AdminOwnerId) {
		return true
	}

	return false
}

// SetAdminOwnerId gets a reference to the given string and assigns it to the AdminOwnerId field.
func (o *Bundle) SetAdminOwnerId(v string) {
	o.AdminOwnerId = &v
}

// GetTotalNumItems returns the TotalNumItems field value if set, zero value otherwise.
func (o *Bundle) GetTotalNumItems() int32 {
	if o == nil || IsNil(o.TotalNumItems) {
		var ret int32
		return ret
	}
	return *o.TotalNumItems
}

// GetTotalNumItemsOk returns a tuple with the TotalNumItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetTotalNumItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNumItems) {
		return nil, false
	}
	return o.TotalNumItems, true
}

// HasTotalNumItems returns a boolean if a field has been set.
func (o *Bundle) HasTotalNumItems() bool {
	if o != nil && !IsNil(o.TotalNumItems) {
		return true
	}

	return false
}

// SetTotalNumItems gets a reference to the given int32 and assigns it to the TotalNumItems field.
func (o *Bundle) SetTotalNumItems(v int32) {
	o.TotalNumItems = &v
}

// GetTotalNumResources returns the TotalNumResources field value if set, zero value otherwise.
func (o *Bundle) GetTotalNumResources() int32 {
	if o == nil || IsNil(o.TotalNumResources) {
		var ret int32
		return ret
	}
	return *o.TotalNumResources
}

// GetTotalNumResourcesOk returns a tuple with the TotalNumResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetTotalNumResourcesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNumResources) {
		return nil, false
	}
	return o.TotalNumResources, true
}

// HasTotalNumResources returns a boolean if a field has been set.
func (o *Bundle) HasTotalNumResources() bool {
	if o != nil && !IsNil(o.TotalNumResources) {
		return true
	}

	return false
}

// SetTotalNumResources gets a reference to the given int32 and assigns it to the TotalNumResources field.
func (o *Bundle) SetTotalNumResources(v int32) {
	o.TotalNumResources = &v
}

// GetTotalNumGroups returns the TotalNumGroups field value if set, zero value otherwise.
func (o *Bundle) GetTotalNumGroups() int32 {
	if o == nil || IsNil(o.TotalNumGroups) {
		var ret int32
		return ret
	}
	return *o.TotalNumGroups
}

// GetTotalNumGroupsOk returns a tuple with the TotalNumGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetTotalNumGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNumGroups) {
		return nil, false
	}
	return o.TotalNumGroups, true
}

// HasTotalNumGroups returns a boolean if a field has been set.
func (o *Bundle) HasTotalNumGroups() bool {
	if o != nil && !IsNil(o.TotalNumGroups) {
		return true
	}

	return false
}

// SetTotalNumGroups gets a reference to the given int32 and assigns it to the TotalNumGroups field.
func (o *Bundle) SetTotalNumGroups(v int32) {
	o.TotalNumGroups = &v
}

func (o Bundle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.AdminOwnerId) {
		toSerialize["admin_owner_id"] = o.AdminOwnerId
	}
	if !IsNil(o.TotalNumItems) {
		toSerialize["total_num_items"] = o.TotalNumItems
	}
	if !IsNil(o.TotalNumResources) {
		toSerialize["total_num_resources"] = o.TotalNumResources
	}
	if !IsNil(o.TotalNumGroups) {
		toSerialize["total_num_groups"] = o.TotalNumGroups
	}
	return toSerialize, nil
}

type NullableBundle struct {
	value *Bundle
	isSet bool
}

func (v NullableBundle) Get() *Bundle {
	return v.value
}

func (v *NullableBundle) Set(val *Bundle) {
	v.value = val
	v.isSet = true
}

func (v NullableBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundle(val *Bundle) *NullableBundle {
	return &NullableBundle{value: val, isSet: true}
}

func (v NullableBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


