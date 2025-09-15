# RequestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**Condition**](Condition.md) | The condition for the request configuration. | [optional] 
**AllowRequests** | **bool** | A bool representing whether or not to allow requests for this resource. | 
**AutoApproval** | **bool** | A bool representing whether or not to automatically approve requests for this resource. | 
**RequireMfaToRequest** | **bool** | A bool representing whether or not to require MFA for requesting access to this resource. | 
**MaxDurationMinutes** | Pointer to **int32** | The maximum duration for which the resource can be requested (in minutes). | [optional] 
**RecommendedDurationMinutes** | Pointer to **int32** | The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration. | [optional] 
**RequireSupportTicket** | **bool** | A bool representing whether or not access requests to the resource require an access ticket. | 
**ExtensionsDurationInMinutes** | Pointer to **int32** | The duration for which access can be extended (in minutes). Set to 0 to disable extensions. When &gt; 0, extensions are enabled for the specified duration. | [optional] 
**RequestTemplateId** | Pointer to **string** | The ID of the associated request template. | [optional] 
**ReviewerStages** | Pointer to [**[]ReviewerStage**](ReviewerStage.md) | The list of reviewer stages for the request configuration. | [optional] 
**Priority** | **int32** | The priority of the request configuration. | 

## Methods

### NewRequestConfiguration

`func NewRequestConfiguration(allowRequests bool, autoApproval bool, requireMfaToRequest bool, requireSupportTicket bool, priority int32, ) *RequestConfiguration`

NewRequestConfiguration instantiates a new RequestConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestConfigurationWithDefaults

`func NewRequestConfigurationWithDefaults() *RequestConfiguration`

NewRequestConfigurationWithDefaults instantiates a new RequestConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *RequestConfiguration) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RequestConfiguration) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RequestConfiguration) SetCondition(v Condition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RequestConfiguration) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetAllowRequests

`func (o *RequestConfiguration) GetAllowRequests() bool`

GetAllowRequests returns the AllowRequests field if non-nil, zero value otherwise.

### GetAllowRequestsOk

`func (o *RequestConfiguration) GetAllowRequestsOk() (*bool, bool)`

GetAllowRequestsOk returns a tuple with the AllowRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRequests

`func (o *RequestConfiguration) SetAllowRequests(v bool)`

SetAllowRequests sets AllowRequests field to given value.


### GetAutoApproval

`func (o *RequestConfiguration) GetAutoApproval() bool`

GetAutoApproval returns the AutoApproval field if non-nil, zero value otherwise.

### GetAutoApprovalOk

`func (o *RequestConfiguration) GetAutoApprovalOk() (*bool, bool)`

GetAutoApprovalOk returns a tuple with the AutoApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproval

`func (o *RequestConfiguration) SetAutoApproval(v bool)`

SetAutoApproval sets AutoApproval field to given value.


### GetRequireMfaToRequest

`func (o *RequestConfiguration) GetRequireMfaToRequest() bool`

GetRequireMfaToRequest returns the RequireMfaToRequest field if non-nil, zero value otherwise.

### GetRequireMfaToRequestOk

`func (o *RequestConfiguration) GetRequireMfaToRequestOk() (*bool, bool)`

GetRequireMfaToRequestOk returns a tuple with the RequireMfaToRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToRequest

`func (o *RequestConfiguration) SetRequireMfaToRequest(v bool)`

SetRequireMfaToRequest sets RequireMfaToRequest field to given value.


### GetMaxDurationMinutes

`func (o *RequestConfiguration) GetMaxDurationMinutes() int32`

GetMaxDurationMinutes returns the MaxDurationMinutes field if non-nil, zero value otherwise.

### GetMaxDurationMinutesOk

`func (o *RequestConfiguration) GetMaxDurationMinutesOk() (*int32, bool)`

GetMaxDurationMinutesOk returns a tuple with the MaxDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationMinutes

`func (o *RequestConfiguration) SetMaxDurationMinutes(v int32)`

SetMaxDurationMinutes sets MaxDurationMinutes field to given value.

### HasMaxDurationMinutes

`func (o *RequestConfiguration) HasMaxDurationMinutes() bool`

HasMaxDurationMinutes returns a boolean if a field has been set.

### GetRecommendedDurationMinutes

`func (o *RequestConfiguration) GetRecommendedDurationMinutes() int32`

GetRecommendedDurationMinutes returns the RecommendedDurationMinutes field if non-nil, zero value otherwise.

### GetRecommendedDurationMinutesOk

`func (o *RequestConfiguration) GetRecommendedDurationMinutesOk() (*int32, bool)`

GetRecommendedDurationMinutesOk returns a tuple with the RecommendedDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedDurationMinutes

`func (o *RequestConfiguration) SetRecommendedDurationMinutes(v int32)`

SetRecommendedDurationMinutes sets RecommendedDurationMinutes field to given value.

### HasRecommendedDurationMinutes

`func (o *RequestConfiguration) HasRecommendedDurationMinutes() bool`

HasRecommendedDurationMinutes returns a boolean if a field has been set.

### GetRequireSupportTicket

`func (o *RequestConfiguration) GetRequireSupportTicket() bool`

GetRequireSupportTicket returns the RequireSupportTicket field if non-nil, zero value otherwise.

### GetRequireSupportTicketOk

`func (o *RequestConfiguration) GetRequireSupportTicketOk() (*bool, bool)`

GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSupportTicket

`func (o *RequestConfiguration) SetRequireSupportTicket(v bool)`

SetRequireSupportTicket sets RequireSupportTicket field to given value.


### GetExtensionsDurationInMinutes

`func (o *RequestConfiguration) GetExtensionsDurationInMinutes() int32`

GetExtensionsDurationInMinutes returns the ExtensionsDurationInMinutes field if non-nil, zero value otherwise.

### GetExtensionsDurationInMinutesOk

`func (o *RequestConfiguration) GetExtensionsDurationInMinutesOk() (*int32, bool)`

GetExtensionsDurationInMinutesOk returns a tuple with the ExtensionsDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionsDurationInMinutes

`func (o *RequestConfiguration) SetExtensionsDurationInMinutes(v int32)`

SetExtensionsDurationInMinutes sets ExtensionsDurationInMinutes field to given value.

### HasExtensionsDurationInMinutes

`func (o *RequestConfiguration) HasExtensionsDurationInMinutes() bool`

HasExtensionsDurationInMinutes returns a boolean if a field has been set.

### GetRequestTemplateId

`func (o *RequestConfiguration) GetRequestTemplateId() string`

GetRequestTemplateId returns the RequestTemplateId field if non-nil, zero value otherwise.

### GetRequestTemplateIdOk

`func (o *RequestConfiguration) GetRequestTemplateIdOk() (*string, bool)`

GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplateId

`func (o *RequestConfiguration) SetRequestTemplateId(v string)`

SetRequestTemplateId sets RequestTemplateId field to given value.

### HasRequestTemplateId

`func (o *RequestConfiguration) HasRequestTemplateId() bool`

HasRequestTemplateId returns a boolean if a field has been set.

### GetReviewerStages

`func (o *RequestConfiguration) GetReviewerStages() []ReviewerStage`

GetReviewerStages returns the ReviewerStages field if non-nil, zero value otherwise.

### GetReviewerStagesOk

`func (o *RequestConfiguration) GetReviewerStagesOk() (*[]ReviewerStage, bool)`

GetReviewerStagesOk returns a tuple with the ReviewerStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerStages

`func (o *RequestConfiguration) SetReviewerStages(v []ReviewerStage)`

SetReviewerStages sets ReviewerStages field to given value.

### HasReviewerStages

`func (o *RequestConfiguration) HasReviewerStages() bool`

HasReviewerStages returns a boolean if a field has been set.

### GetPriority

`func (o *RequestConfiguration) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RequestConfiguration) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RequestConfiguration) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


