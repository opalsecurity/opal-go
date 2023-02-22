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
	"time"
)

// checks if the GroupUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUser{}

// GroupUser # Group Access User Object ### Description The `GroupAccessUser` object is used to represent a user with access to a group.  ### Usage Example Fetch from the `LIST GroupUsers` endpoint.
type GroupUser struct {
	// The ID of the group.
	GroupId string `json:"group_id"`
	// The ID of the user.
	UserId string `json:"user_id"`
	AccessLevel *GroupAccessLevel `json:"access_level,omitempty"`
	// The user's full name.
	FullName string `json:"full_name"`
	// The user's email.
	Email string `json:"email"`
	// The day and time the user's access will expire.
	ExpirationDate NullableTime `json:"expiration_date"`
}

// NewGroupUser instantiates a new GroupUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUser(groupId string, userId string, fullName string, email string, expirationDate NullableTime) *GroupUser {
	this := GroupUser{}
	this.GroupId = groupId
	this.UserId = userId
	this.FullName = fullName
	this.Email = email
	this.ExpirationDate = expirationDate
	return &this
}

// NewGroupUserWithDefaults instantiates a new GroupUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUserWithDefaults() *GroupUser {
	this := GroupUser{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *GroupUser) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *GroupUser) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *GroupUser) SetGroupId(v string) {
	o.GroupId = v
}

// GetUserId returns the UserId field value
func (o *GroupUser) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *GroupUser) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *GroupUser) SetUserId(v string) {
	o.UserId = v
}

// GetAccessLevel returns the AccessLevel field value if set, zero value otherwise.
func (o *GroupUser) GetAccessLevel() GroupAccessLevel {
	if o == nil || IsNil(o.AccessLevel) {
		var ret GroupAccessLevel
		return ret
	}
	return *o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUser) GetAccessLevelOk() (*GroupAccessLevel, bool) {
	if o == nil || IsNil(o.AccessLevel) {
		return nil, false
	}
	return o.AccessLevel, true
}

// HasAccessLevel returns a boolean if a field has been set.
func (o *GroupUser) HasAccessLevel() bool {
	if o != nil && !IsNil(o.AccessLevel) {
		return true
	}

	return false
}

// SetAccessLevel gets a reference to the given GroupAccessLevel and assigns it to the AccessLevel field.
func (o *GroupUser) SetAccessLevel(v GroupAccessLevel) {
	o.AccessLevel = &v
}

// GetFullName returns the FullName field value
func (o *GroupUser) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *GroupUser) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *GroupUser) SetFullName(v string) {
	o.FullName = v
}

// GetEmail returns the Email field value
func (o *GroupUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *GroupUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *GroupUser) SetEmail(v string) {
	o.Email = v
}

// GetExpirationDate returns the ExpirationDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *GroupUser) GetExpirationDate() time.Time {
	if o == nil || o.ExpirationDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUser) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// SetExpirationDate sets field value
func (o *GroupUser) SetExpirationDate(v time.Time) {
	o.ExpirationDate.Set(&v)
}

func (o GroupUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.AccessLevel) {
		toSerialize["access_level"] = o.AccessLevel
	}
	toSerialize["full_name"] = o.FullName
	toSerialize["email"] = o.Email
	toSerialize["expiration_date"] = o.ExpirationDate.Get()
	return toSerialize, nil
}

type NullableGroupUser struct {
	value *GroupUser
	isSet bool
}

func (v NullableGroupUser) Get() *GroupUser {
	return v.value
}

func (v *NullableGroupUser) Set(val *GroupUser) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUser) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUser(val *GroupUser) *NullableGroupUser {
	return &NullableGroupUser{value: val, isSet: true}
}

func (v NullableGroupUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


