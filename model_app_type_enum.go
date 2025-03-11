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
	"fmt"
)

// AppTypeEnum The type of an app.
type AppTypeEnum string

// List of AppTypeEnum
const (
	APPTYPEENUM_ACTIVE_DIRECTORY AppTypeEnum = "ACTIVE_DIRECTORY"
	APPTYPEENUM_AZURE_AD AppTypeEnum = "AZURE_AD"
	APPTYPEENUM_AWS AppTypeEnum = "AWS"
	APPTYPEENUM_AWS_SSO AppTypeEnum = "AWS_SSO"
	APPTYPEENUM_CUSTOM AppTypeEnum = "CUSTOM"
	APPTYPEENUM_DUO AppTypeEnum = "DUO"
	APPTYPEENUM_GCP AppTypeEnum = "GCP"
	APPTYPEENUM_GIT_HUB AppTypeEnum = "GIT_HUB"
	APPTYPEENUM_GIT_LAB AppTypeEnum = "GIT_LAB"
	APPTYPEENUM_GOOGLE_GROUPS AppTypeEnum = "GOOGLE_GROUPS"
	APPTYPEENUM_GOOGLE_WORKSPACE AppTypeEnum = "GOOGLE_WORKSPACE"
	APPTYPEENUM_LDAP AppTypeEnum = "LDAP"
	APPTYPEENUM_MARIADB AppTypeEnum = "MARIADB"
	APPTYPEENUM_MONGO AppTypeEnum = "MONGO"
	APPTYPEENUM_MONGO_ATLAS AppTypeEnum = "MONGO_ATLAS"
	APPTYPEENUM_MYSQL AppTypeEnum = "MYSQL"
	APPTYPEENUM_OKTA_DIRECTORY AppTypeEnum = "OKTA_DIRECTORY"
	APPTYPEENUM_OPAL AppTypeEnum = "OPAL"
	APPTYPEENUM_PAGERDUTY AppTypeEnum = "PAGERDUTY"
	APPTYPEENUM_SALESFORCE AppTypeEnum = "SALESFORCE"
	APPTYPEENUM_TAILSCALE AppTypeEnum = "TAILSCALE"
	APPTYPEENUM_TELEPORT AppTypeEnum = "TELEPORT"
	APPTYPEENUM_WORKDAY AppTypeEnum = "WORKDAY"
)

// All allowed values of AppTypeEnum enum
var AllowedAppTypeEnumEnumValues = []AppTypeEnum{
	"ACTIVE_DIRECTORY",
	"AZURE_AD",
	"AWS",
	"AWS_SSO",
	"CUSTOM",
	"DUO",
	"GCP",
	"GIT_HUB",
	"GIT_LAB",
	"GOOGLE_GROUPS",
	"GOOGLE_WORKSPACE",
	"LDAP",
	"MARIADB",
	"MONGO",
	"MONGO_ATLAS",
	"MYSQL",
	"OKTA_DIRECTORY",
	"OPAL",
	"PAGERDUTY",
	"SALESFORCE",
	"TAILSCALE",
	"TELEPORT",
	"WORKDAY",
}

func (v *AppTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppTypeEnum(value)
	for _, existing := range AllowedAppTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppTypeEnum", value)
}

// NewAppTypeEnumFromValue returns a pointer to a valid AppTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppTypeEnumFromValue(v string) (*AppTypeEnum, error) {
	ev := AppTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppTypeEnum: valid values are %v", v, AllowedAppTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppTypeEnum) IsValid() bool {
	for _, existing := range AllowedAppTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppTypeEnum value
func (v AppTypeEnum) Ptr() *AppTypeEnum {
	return &v
}

type NullableAppTypeEnum struct {
	value *AppTypeEnum
	isSet bool
}

func (v NullableAppTypeEnum) Get() *AppTypeEnum {
	return v.value
}

func (v *NullableAppTypeEnum) Set(val *AppTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAppTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAppTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppTypeEnum(val *AppTypeEnum) *NullableAppTypeEnum {
	return &NullableAppTypeEnum{value: val, isSet: true}
}

func (v NullableAppTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

