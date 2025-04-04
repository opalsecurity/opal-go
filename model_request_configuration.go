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

// checks if the RequestConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestConfiguration{}

// RequestConfiguration # Request Configuration Object ### Description The `RequestConfiguration` object is used to represent a request configuration.  ### Usage Example Returned from the `GET Request Configurations` endpoint.
type RequestConfiguration struct {
	Condition *Condition `json:"condition,omitempty"`
	// A bool representing whether or not to allow requests for this resource.
	AllowRequests bool `json:"allow_requests"`
	// A bool representing whether or not to automatically approve requests for this resource.
	AutoApproval bool `json:"auto_approval"`
	// A bool representing whether or not to require MFA for requesting access to this resource.
	RequireMfaToRequest bool `json:"require_mfa_to_request"`
	// The maximum duration for which the resource can be requested (in minutes).
	MaxDurationMinutes *int32 `json:"max_duration_minutes,omitempty"`
	// The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration.
	RecommendedDurationMinutes *int32 `json:"recommended_duration_minutes,omitempty"`
	// A bool representing whether or not access requests to the resource require an access ticket.
	RequireSupportTicket bool `json:"require_support_ticket"`
	// The ID of the associated request template.
	RequestTemplateId *string `json:"request_template_id,omitempty"`
	// The list of reviewer stages for the request configuration.
	ReviewerStages []ReviewerStage `json:"reviewer_stages,omitempty"`
	// The priority of the request configuration.
	Priority int32 `json:"priority"`
}

type _RequestConfiguration RequestConfiguration

// NewRequestConfiguration instantiates a new RequestConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConfiguration(allowRequests bool, autoApproval bool, requireMfaToRequest bool, requireSupportTicket bool, priority int32) *RequestConfiguration {
	this := RequestConfiguration{}
	this.AllowRequests = allowRequests
	this.AutoApproval = autoApproval
	this.RequireMfaToRequest = requireMfaToRequest
	this.RequireSupportTicket = requireSupportTicket
	this.Priority = priority
	return &this
}

// NewRequestConfigurationWithDefaults instantiates a new RequestConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConfigurationWithDefaults() *RequestConfiguration {
	this := RequestConfiguration{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *RequestConfiguration) GetCondition() Condition {
	if o == nil || IsNil(o.Condition) {
		var ret Condition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetConditionOk() (*Condition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *RequestConfiguration) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given Condition and assigns it to the Condition field.
func (o *RequestConfiguration) SetCondition(v Condition) {
	o.Condition = &v
}

// GetAllowRequests returns the AllowRequests field value
func (o *RequestConfiguration) GetAllowRequests() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowRequests
}

// GetAllowRequestsOk returns a tuple with the AllowRequests field value
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetAllowRequestsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowRequests, true
}

// SetAllowRequests sets field value
func (o *RequestConfiguration) SetAllowRequests(v bool) {
	o.AllowRequests = v
}

// GetAutoApproval returns the AutoApproval field value
func (o *RequestConfiguration) GetAutoApproval() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoApproval
}

// GetAutoApprovalOk returns a tuple with the AutoApproval field value
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetAutoApprovalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoApproval, true
}

// SetAutoApproval sets field value
func (o *RequestConfiguration) SetAutoApproval(v bool) {
	o.AutoApproval = v
}

// GetRequireMfaToRequest returns the RequireMfaToRequest field value
func (o *RequestConfiguration) GetRequireMfaToRequest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireMfaToRequest
}

// GetRequireMfaToRequestOk returns a tuple with the RequireMfaToRequest field value
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetRequireMfaToRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireMfaToRequest, true
}

// SetRequireMfaToRequest sets field value
func (o *RequestConfiguration) SetRequireMfaToRequest(v bool) {
	o.RequireMfaToRequest = v
}

// GetMaxDurationMinutes returns the MaxDurationMinutes field value if set, zero value otherwise.
func (o *RequestConfiguration) GetMaxDurationMinutes() int32 {
	if o == nil || IsNil(o.MaxDurationMinutes) {
		var ret int32
		return ret
	}
	return *o.MaxDurationMinutes
}

// GetMaxDurationMinutesOk returns a tuple with the MaxDurationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetMaxDurationMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDurationMinutes) {
		return nil, false
	}
	return o.MaxDurationMinutes, true
}

// HasMaxDurationMinutes returns a boolean if a field has been set.
func (o *RequestConfiguration) HasMaxDurationMinutes() bool {
	if o != nil && !IsNil(o.MaxDurationMinutes) {
		return true
	}

	return false
}

// SetMaxDurationMinutes gets a reference to the given int32 and assigns it to the MaxDurationMinutes field.
func (o *RequestConfiguration) SetMaxDurationMinutes(v int32) {
	o.MaxDurationMinutes = &v
}

// GetRecommendedDurationMinutes returns the RecommendedDurationMinutes field value if set, zero value otherwise.
func (o *RequestConfiguration) GetRecommendedDurationMinutes() int32 {
	if o == nil || IsNil(o.RecommendedDurationMinutes) {
		var ret int32
		return ret
	}
	return *o.RecommendedDurationMinutes
}

// GetRecommendedDurationMinutesOk returns a tuple with the RecommendedDurationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetRecommendedDurationMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.RecommendedDurationMinutes) {
		return nil, false
	}
	return o.RecommendedDurationMinutes, true
}

// HasRecommendedDurationMinutes returns a boolean if a field has been set.
func (o *RequestConfiguration) HasRecommendedDurationMinutes() bool {
	if o != nil && !IsNil(o.RecommendedDurationMinutes) {
		return true
	}

	return false
}

// SetRecommendedDurationMinutes gets a reference to the given int32 and assigns it to the RecommendedDurationMinutes field.
func (o *RequestConfiguration) SetRecommendedDurationMinutes(v int32) {
	o.RecommendedDurationMinutes = &v
}

// GetRequireSupportTicket returns the RequireSupportTicket field value
func (o *RequestConfiguration) GetRequireSupportTicket() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireSupportTicket
}

// GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field value
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetRequireSupportTicketOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireSupportTicket, true
}

// SetRequireSupportTicket sets field value
func (o *RequestConfiguration) SetRequireSupportTicket(v bool) {
	o.RequireSupportTicket = v
}

// GetRequestTemplateId returns the RequestTemplateId field value if set, zero value otherwise.
func (o *RequestConfiguration) GetRequestTemplateId() string {
	if o == nil || IsNil(o.RequestTemplateId) {
		var ret string
		return ret
	}
	return *o.RequestTemplateId
}

// GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetRequestTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestTemplateId) {
		return nil, false
	}
	return o.RequestTemplateId, true
}

// HasRequestTemplateId returns a boolean if a field has been set.
func (o *RequestConfiguration) HasRequestTemplateId() bool {
	if o != nil && !IsNil(o.RequestTemplateId) {
		return true
	}

	return false
}

// SetRequestTemplateId gets a reference to the given string and assigns it to the RequestTemplateId field.
func (o *RequestConfiguration) SetRequestTemplateId(v string) {
	o.RequestTemplateId = &v
}

// GetReviewerStages returns the ReviewerStages field value if set, zero value otherwise.
func (o *RequestConfiguration) GetReviewerStages() []ReviewerStage {
	if o == nil || IsNil(o.ReviewerStages) {
		var ret []ReviewerStage
		return ret
	}
	return o.ReviewerStages
}

// GetReviewerStagesOk returns a tuple with the ReviewerStages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetReviewerStagesOk() ([]ReviewerStage, bool) {
	if o == nil || IsNil(o.ReviewerStages) {
		return nil, false
	}
	return o.ReviewerStages, true
}

// HasReviewerStages returns a boolean if a field has been set.
func (o *RequestConfiguration) HasReviewerStages() bool {
	if o != nil && !IsNil(o.ReviewerStages) {
		return true
	}

	return false
}

// SetReviewerStages gets a reference to the given []ReviewerStage and assigns it to the ReviewerStages field.
func (o *RequestConfiguration) SetReviewerStages(v []ReviewerStage) {
	o.ReviewerStages = v
}

// GetPriority returns the Priority field value
func (o *RequestConfiguration) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *RequestConfiguration) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *RequestConfiguration) SetPriority(v int32) {
	o.Priority = v
}

func (o RequestConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	toSerialize["allow_requests"] = o.AllowRequests
	toSerialize["auto_approval"] = o.AutoApproval
	toSerialize["require_mfa_to_request"] = o.RequireMfaToRequest
	if !IsNil(o.MaxDurationMinutes) {
		toSerialize["max_duration_minutes"] = o.MaxDurationMinutes
	}
	if !IsNil(o.RecommendedDurationMinutes) {
		toSerialize["recommended_duration_minutes"] = o.RecommendedDurationMinutes
	}
	toSerialize["require_support_ticket"] = o.RequireSupportTicket
	if !IsNil(o.RequestTemplateId) {
		toSerialize["request_template_id"] = o.RequestTemplateId
	}
	if !IsNil(o.ReviewerStages) {
		toSerialize["reviewer_stages"] = o.ReviewerStages
	}
	toSerialize["priority"] = o.Priority
	return toSerialize, nil
}

func (o *RequestConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allow_requests",
		"auto_approval",
		"require_mfa_to_request",
		"require_support_ticket",
		"priority",
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

	varRequestConfiguration := _RequestConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequestConfiguration)

	if err != nil {
		return err
	}

	*o = RequestConfiguration(varRequestConfiguration)

	return err
}

type NullableRequestConfiguration struct {
	value *RequestConfiguration
	isSet bool
}

func (v NullableRequestConfiguration) Get() *RequestConfiguration {
	return v.value
}

func (v *NullableRequestConfiguration) Set(val *RequestConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConfiguration(val *RequestConfiguration) *NullableRequestConfiguration {
	return &NullableRequestConfiguration{value: val, isSet: true}
}

func (v NullableRequestConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


