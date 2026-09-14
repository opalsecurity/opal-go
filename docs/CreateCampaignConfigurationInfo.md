# CreateCampaignConfigurationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | [**OpalAccessPathQueryBody**](OpalAccessPathQueryBody.md) | Access-path query defining the scope of access to review. Required. Uses the same principalFilter / entitlementFilter shape as ACCESS_PATH OpalQuery. Must include at least one of principalFilter or entitlementFilter.  Campaign scope only supports direct access edges: &#x60;edgeFilter.directOnly&#x60; defaults to &#x60;true&#x60;, is always stored as &#x60;true&#x60;, and passing &#x60;false&#x60; returns 400.  | 
**ReviewerAssignmentPolicy** | Pointer to [**UARReviewerAssignmentPolicyEnum**](UARReviewerAssignmentPolicyEnum.md) |  | [optional] 
**AllowSelfReview** | Pointer to **bool** | Whether reviewers can review their own access. | [optional] 
**SendReviewerAssignmentNotification** | Pointer to **bool** | Whether to notify reviewers upon assignment. | [optional] 
**AllowReviewerReassignment** | Pointer to **bool** | Whether reviewers may reassign their reviews to another user. | [optional] 
**StartDate** | Pointer to **time.Time** | Scheduled start date of the campaign. | [optional] 
**EndDate** | Pointer to **time.Time** | Scheduled end date of the campaign. | [optional] 
**Timezone** | Pointer to **string** | IANA timezone used to interpret campaign deadlines (e.g. America/Los_Angeles). | [optional] 
**RevokeOn** | Pointer to [**CampaignRevokeOnEnum**](CampaignRevokeOnEnum.md) |  | [optional] 
**ReminderSchedule** | Pointer to **[]int32** | Days before end date to send reminder notifications. | [optional] 
**ReminderIncludeManager** | Pointer to **bool** | Whether to include the reviewer&#39;s manager in reminders. | [optional] 
**RequireReasonOnDenial** | Pointer to **bool** | Whether reviewers must provide a reason when denying (revoking) access. | [optional] 
**HideAiSuggestions** | Pointer to **bool** | Whether AI suggestions are hidden from reviewers. | [optional] 
**CustomStartMessage** | Pointer to **string** | Optional custom message included when notifying reviewers that the campaign started. | [optional] 
**GroupAssetVisibilityPolicy** | Pointer to [**CampaignGroupAssetVisibilityPolicyEnum**](CampaignGroupAssetVisibilityPolicyEnum.md) |  | [optional] 
**IsTemplate** | Pointer to **bool** | Whether this configuration is a recurring schedule template. | [optional] 
**CronExpression** | Pointer to **string** | Cron expression driving the recurring schedule. Null for one-off campaigns. | [optional] 
**RecurringDurationDays** | Pointer to **int32** | Deadline window in days applied to each draft generated from this template. | [optional] 
**ExcludedRoleAssignmentIds** | Pointer to **[]string** | Role assignment IDs to exclude from the campaign scope during population. | [optional] 

## Methods

### NewCreateCampaignConfigurationInfo

`func NewCreateCampaignConfigurationInfo(query OpalAccessPathQueryBody, ) *CreateCampaignConfigurationInfo`

NewCreateCampaignConfigurationInfo instantiates a new CreateCampaignConfigurationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCampaignConfigurationInfoWithDefaults

`func NewCreateCampaignConfigurationInfoWithDefaults() *CreateCampaignConfigurationInfo`

NewCreateCampaignConfigurationInfoWithDefaults instantiates a new CreateCampaignConfigurationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CreateCampaignConfigurationInfo) GetQuery() OpalAccessPathQueryBody`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateCampaignConfigurationInfo) GetQueryOk() (*OpalAccessPathQueryBody, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateCampaignConfigurationInfo) SetQuery(v OpalAccessPathQueryBody)`

SetQuery sets Query field to given value.


### GetReviewerAssignmentPolicy

`func (o *CreateCampaignConfigurationInfo) GetReviewerAssignmentPolicy() UARReviewerAssignmentPolicyEnum`

GetReviewerAssignmentPolicy returns the ReviewerAssignmentPolicy field if non-nil, zero value otherwise.

### GetReviewerAssignmentPolicyOk

`func (o *CreateCampaignConfigurationInfo) GetReviewerAssignmentPolicyOk() (*UARReviewerAssignmentPolicyEnum, bool)`

GetReviewerAssignmentPolicyOk returns a tuple with the ReviewerAssignmentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerAssignmentPolicy

`func (o *CreateCampaignConfigurationInfo) SetReviewerAssignmentPolicy(v UARReviewerAssignmentPolicyEnum)`

SetReviewerAssignmentPolicy sets ReviewerAssignmentPolicy field to given value.

### HasReviewerAssignmentPolicy

`func (o *CreateCampaignConfigurationInfo) HasReviewerAssignmentPolicy() bool`

HasReviewerAssignmentPolicy returns a boolean if a field has been set.

### GetAllowSelfReview

`func (o *CreateCampaignConfigurationInfo) GetAllowSelfReview() bool`

GetAllowSelfReview returns the AllowSelfReview field if non-nil, zero value otherwise.

### GetAllowSelfReviewOk

`func (o *CreateCampaignConfigurationInfo) GetAllowSelfReviewOk() (*bool, bool)`

GetAllowSelfReviewOk returns a tuple with the AllowSelfReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSelfReview

`func (o *CreateCampaignConfigurationInfo) SetAllowSelfReview(v bool)`

SetAllowSelfReview sets AllowSelfReview field to given value.

### HasAllowSelfReview

`func (o *CreateCampaignConfigurationInfo) HasAllowSelfReview() bool`

HasAllowSelfReview returns a boolean if a field has been set.

### GetSendReviewerAssignmentNotification

`func (o *CreateCampaignConfigurationInfo) GetSendReviewerAssignmentNotification() bool`

GetSendReviewerAssignmentNotification returns the SendReviewerAssignmentNotification field if non-nil, zero value otherwise.

### GetSendReviewerAssignmentNotificationOk

`func (o *CreateCampaignConfigurationInfo) GetSendReviewerAssignmentNotificationOk() (*bool, bool)`

GetSendReviewerAssignmentNotificationOk returns a tuple with the SendReviewerAssignmentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReviewerAssignmentNotification

`func (o *CreateCampaignConfigurationInfo) SetSendReviewerAssignmentNotification(v bool)`

SetSendReviewerAssignmentNotification sets SendReviewerAssignmentNotification field to given value.

### HasSendReviewerAssignmentNotification

`func (o *CreateCampaignConfigurationInfo) HasSendReviewerAssignmentNotification() bool`

HasSendReviewerAssignmentNotification returns a boolean if a field has been set.

### GetAllowReviewerReassignment

`func (o *CreateCampaignConfigurationInfo) GetAllowReviewerReassignment() bool`

GetAllowReviewerReassignment returns the AllowReviewerReassignment field if non-nil, zero value otherwise.

### GetAllowReviewerReassignmentOk

`func (o *CreateCampaignConfigurationInfo) GetAllowReviewerReassignmentOk() (*bool, bool)`

GetAllowReviewerReassignmentOk returns a tuple with the AllowReviewerReassignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReviewerReassignment

`func (o *CreateCampaignConfigurationInfo) SetAllowReviewerReassignment(v bool)`

SetAllowReviewerReassignment sets AllowReviewerReassignment field to given value.

### HasAllowReviewerReassignment

`func (o *CreateCampaignConfigurationInfo) HasAllowReviewerReassignment() bool`

HasAllowReviewerReassignment returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateCampaignConfigurationInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateCampaignConfigurationInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateCampaignConfigurationInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateCampaignConfigurationInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateCampaignConfigurationInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateCampaignConfigurationInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateCampaignConfigurationInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateCampaignConfigurationInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTimezone

`func (o *CreateCampaignConfigurationInfo) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CreateCampaignConfigurationInfo) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CreateCampaignConfigurationInfo) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CreateCampaignConfigurationInfo) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetRevokeOn

`func (o *CreateCampaignConfigurationInfo) GetRevokeOn() CampaignRevokeOnEnum`

GetRevokeOn returns the RevokeOn field if non-nil, zero value otherwise.

### GetRevokeOnOk

`func (o *CreateCampaignConfigurationInfo) GetRevokeOnOk() (*CampaignRevokeOnEnum, bool)`

GetRevokeOnOk returns a tuple with the RevokeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeOn

`func (o *CreateCampaignConfigurationInfo) SetRevokeOn(v CampaignRevokeOnEnum)`

SetRevokeOn sets RevokeOn field to given value.

### HasRevokeOn

`func (o *CreateCampaignConfigurationInfo) HasRevokeOn() bool`

HasRevokeOn returns a boolean if a field has been set.

### GetReminderSchedule

`func (o *CreateCampaignConfigurationInfo) GetReminderSchedule() []int32`

GetReminderSchedule returns the ReminderSchedule field if non-nil, zero value otherwise.

### GetReminderScheduleOk

`func (o *CreateCampaignConfigurationInfo) GetReminderScheduleOk() (*[]int32, bool)`

GetReminderScheduleOk returns a tuple with the ReminderSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderSchedule

`func (o *CreateCampaignConfigurationInfo) SetReminderSchedule(v []int32)`

SetReminderSchedule sets ReminderSchedule field to given value.

### HasReminderSchedule

`func (o *CreateCampaignConfigurationInfo) HasReminderSchedule() bool`

HasReminderSchedule returns a boolean if a field has been set.

### GetReminderIncludeManager

`func (o *CreateCampaignConfigurationInfo) GetReminderIncludeManager() bool`

GetReminderIncludeManager returns the ReminderIncludeManager field if non-nil, zero value otherwise.

### GetReminderIncludeManagerOk

`func (o *CreateCampaignConfigurationInfo) GetReminderIncludeManagerOk() (*bool, bool)`

GetReminderIncludeManagerOk returns a tuple with the ReminderIncludeManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderIncludeManager

`func (o *CreateCampaignConfigurationInfo) SetReminderIncludeManager(v bool)`

SetReminderIncludeManager sets ReminderIncludeManager field to given value.

### HasReminderIncludeManager

`func (o *CreateCampaignConfigurationInfo) HasReminderIncludeManager() bool`

HasReminderIncludeManager returns a boolean if a field has been set.

### GetRequireReasonOnDenial

`func (o *CreateCampaignConfigurationInfo) GetRequireReasonOnDenial() bool`

GetRequireReasonOnDenial returns the RequireReasonOnDenial field if non-nil, zero value otherwise.

### GetRequireReasonOnDenialOk

`func (o *CreateCampaignConfigurationInfo) GetRequireReasonOnDenialOk() (*bool, bool)`

GetRequireReasonOnDenialOk returns a tuple with the RequireReasonOnDenial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireReasonOnDenial

`func (o *CreateCampaignConfigurationInfo) SetRequireReasonOnDenial(v bool)`

SetRequireReasonOnDenial sets RequireReasonOnDenial field to given value.

### HasRequireReasonOnDenial

`func (o *CreateCampaignConfigurationInfo) HasRequireReasonOnDenial() bool`

HasRequireReasonOnDenial returns a boolean if a field has been set.

### GetHideAiSuggestions

`func (o *CreateCampaignConfigurationInfo) GetHideAiSuggestions() bool`

GetHideAiSuggestions returns the HideAiSuggestions field if non-nil, zero value otherwise.

### GetHideAiSuggestionsOk

`func (o *CreateCampaignConfigurationInfo) GetHideAiSuggestionsOk() (*bool, bool)`

GetHideAiSuggestionsOk returns a tuple with the HideAiSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAiSuggestions

`func (o *CreateCampaignConfigurationInfo) SetHideAiSuggestions(v bool)`

SetHideAiSuggestions sets HideAiSuggestions field to given value.

### HasHideAiSuggestions

`func (o *CreateCampaignConfigurationInfo) HasHideAiSuggestions() bool`

HasHideAiSuggestions returns a boolean if a field has been set.

### GetCustomStartMessage

`func (o *CreateCampaignConfigurationInfo) GetCustomStartMessage() string`

GetCustomStartMessage returns the CustomStartMessage field if non-nil, zero value otherwise.

### GetCustomStartMessageOk

`func (o *CreateCampaignConfigurationInfo) GetCustomStartMessageOk() (*string, bool)`

GetCustomStartMessageOk returns a tuple with the CustomStartMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStartMessage

`func (o *CreateCampaignConfigurationInfo) SetCustomStartMessage(v string)`

SetCustomStartMessage sets CustomStartMessage field to given value.

### HasCustomStartMessage

`func (o *CreateCampaignConfigurationInfo) HasCustomStartMessage() bool`

HasCustomStartMessage returns a boolean if a field has been set.

### GetGroupAssetVisibilityPolicy

`func (o *CreateCampaignConfigurationInfo) GetGroupAssetVisibilityPolicy() CampaignGroupAssetVisibilityPolicyEnum`

GetGroupAssetVisibilityPolicy returns the GroupAssetVisibilityPolicy field if non-nil, zero value otherwise.

### GetGroupAssetVisibilityPolicyOk

`func (o *CreateCampaignConfigurationInfo) GetGroupAssetVisibilityPolicyOk() (*CampaignGroupAssetVisibilityPolicyEnum, bool)`

GetGroupAssetVisibilityPolicyOk returns a tuple with the GroupAssetVisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAssetVisibilityPolicy

`func (o *CreateCampaignConfigurationInfo) SetGroupAssetVisibilityPolicy(v CampaignGroupAssetVisibilityPolicyEnum)`

SetGroupAssetVisibilityPolicy sets GroupAssetVisibilityPolicy field to given value.

### HasGroupAssetVisibilityPolicy

`func (o *CreateCampaignConfigurationInfo) HasGroupAssetVisibilityPolicy() bool`

HasGroupAssetVisibilityPolicy returns a boolean if a field has been set.

### GetIsTemplate

`func (o *CreateCampaignConfigurationInfo) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *CreateCampaignConfigurationInfo) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *CreateCampaignConfigurationInfo) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *CreateCampaignConfigurationInfo) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetCronExpression

`func (o *CreateCampaignConfigurationInfo) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *CreateCampaignConfigurationInfo) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *CreateCampaignConfigurationInfo) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *CreateCampaignConfigurationInfo) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetRecurringDurationDays

`func (o *CreateCampaignConfigurationInfo) GetRecurringDurationDays() int32`

GetRecurringDurationDays returns the RecurringDurationDays field if non-nil, zero value otherwise.

### GetRecurringDurationDaysOk

`func (o *CreateCampaignConfigurationInfo) GetRecurringDurationDaysOk() (*int32, bool)`

GetRecurringDurationDaysOk returns a tuple with the RecurringDurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDurationDays

`func (o *CreateCampaignConfigurationInfo) SetRecurringDurationDays(v int32)`

SetRecurringDurationDays sets RecurringDurationDays field to given value.

### HasRecurringDurationDays

`func (o *CreateCampaignConfigurationInfo) HasRecurringDurationDays() bool`

HasRecurringDurationDays returns a boolean if a field has been set.

### GetExcludedRoleAssignmentIds

`func (o *CreateCampaignConfigurationInfo) GetExcludedRoleAssignmentIds() []string`

GetExcludedRoleAssignmentIds returns the ExcludedRoleAssignmentIds field if non-nil, zero value otherwise.

### GetExcludedRoleAssignmentIdsOk

`func (o *CreateCampaignConfigurationInfo) GetExcludedRoleAssignmentIdsOk() (*[]string, bool)`

GetExcludedRoleAssignmentIdsOk returns a tuple with the ExcludedRoleAssignmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedRoleAssignmentIds

`func (o *CreateCampaignConfigurationInfo) SetExcludedRoleAssignmentIds(v []string)`

SetExcludedRoleAssignmentIds sets ExcludedRoleAssignmentIds field to given value.

### HasExcludedRoleAssignmentIds

`func (o *CreateCampaignConfigurationInfo) HasExcludedRoleAssignmentIds() bool`

HasExcludedRoleAssignmentIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


