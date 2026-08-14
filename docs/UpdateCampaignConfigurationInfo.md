# UpdateCampaignConfigurationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSelfReview** | Pointer to **bool** | Whether reviewers can review their own access. | [optional] 
**SendReviewerAssignmentNotification** | Pointer to **bool** | Whether to notify reviewers upon assignment. | [optional] 
**AllowReviewerReassignment** | Pointer to **bool** | Whether reviewers may reassign their reviews to another user. | [optional] 
**StartDate** | Pointer to **time.Time** | Scheduled start date of the campaign. May only be updated while the campaign has not started (started_at is null). When set, the date&#39;s calendar day in the campaign timezone must be at least tomorrow.  | [optional] 
**EndDate** | Pointer to **time.Time** | Scheduled end date of the campaign. When set, the date&#39;s calendar day in the campaign timezone must be at least tomorrow.  | [optional] 
**Timezone** | Pointer to **string** | IANA timezone used to interpret campaign deadlines (e.g. America/Los_Angeles). | [optional] 
**RevokeOn** | Pointer to [**CampaignRevokeOnEnum**](CampaignRevokeOnEnum.md) |  | [optional] 
**ReminderSchedule** | Pointer to **[]int32** | Days before end date to send reminder notifications. | [optional] 
**ReminderIncludeManager** | Pointer to **bool** | Whether to include the reviewer&#39;s manager in reminders. | [optional] 
**RequireReasonOnDenial** | Pointer to **bool** | Whether reviewers must provide a reason when denying (revoking) access. | [optional] 
**HideAiSuggestions** | Pointer to **bool** | Whether AI suggestions are hidden from reviewers. | [optional] 
**CustomStartMessage** | Pointer to **string** | Optional custom message included when notifying reviewers that the campaign started. | [optional] 
**GroupAssetVisibilityPolicy** | Pointer to [**CampaignGroupAssetVisibilityPolicyEnum**](CampaignGroupAssetVisibilityPolicyEnum.md) |  | [optional] 
**CronExpression** | Pointer to **string** | Cron expression driving the recurring schedule. Only valid on template campaigns. Pass an empty string to clear the active months (next_scheduled_run is cleared); the campaign remains a template. | [optional] 
**RecurringDurationDays** | Pointer to **int32** | Deadline window in days applied to each draft generated from this template. Only valid on template campaigns. | [optional] 

## Methods

### NewUpdateCampaignConfigurationInfo

`func NewUpdateCampaignConfigurationInfo() *UpdateCampaignConfigurationInfo`

NewUpdateCampaignConfigurationInfo instantiates a new UpdateCampaignConfigurationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignConfigurationInfoWithDefaults

`func NewUpdateCampaignConfigurationInfoWithDefaults() *UpdateCampaignConfigurationInfo`

NewUpdateCampaignConfigurationInfoWithDefaults instantiates a new UpdateCampaignConfigurationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSelfReview

`func (o *UpdateCampaignConfigurationInfo) GetAllowSelfReview() bool`

GetAllowSelfReview returns the AllowSelfReview field if non-nil, zero value otherwise.

### GetAllowSelfReviewOk

`func (o *UpdateCampaignConfigurationInfo) GetAllowSelfReviewOk() (*bool, bool)`

GetAllowSelfReviewOk returns a tuple with the AllowSelfReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSelfReview

`func (o *UpdateCampaignConfigurationInfo) SetAllowSelfReview(v bool)`

SetAllowSelfReview sets AllowSelfReview field to given value.

### HasAllowSelfReview

`func (o *UpdateCampaignConfigurationInfo) HasAllowSelfReview() bool`

HasAllowSelfReview returns a boolean if a field has been set.

### GetSendReviewerAssignmentNotification

`func (o *UpdateCampaignConfigurationInfo) GetSendReviewerAssignmentNotification() bool`

GetSendReviewerAssignmentNotification returns the SendReviewerAssignmentNotification field if non-nil, zero value otherwise.

### GetSendReviewerAssignmentNotificationOk

`func (o *UpdateCampaignConfigurationInfo) GetSendReviewerAssignmentNotificationOk() (*bool, bool)`

GetSendReviewerAssignmentNotificationOk returns a tuple with the SendReviewerAssignmentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReviewerAssignmentNotification

`func (o *UpdateCampaignConfigurationInfo) SetSendReviewerAssignmentNotification(v bool)`

SetSendReviewerAssignmentNotification sets SendReviewerAssignmentNotification field to given value.

### HasSendReviewerAssignmentNotification

`func (o *UpdateCampaignConfigurationInfo) HasSendReviewerAssignmentNotification() bool`

HasSendReviewerAssignmentNotification returns a boolean if a field has been set.

### GetAllowReviewerReassignment

`func (o *UpdateCampaignConfigurationInfo) GetAllowReviewerReassignment() bool`

GetAllowReviewerReassignment returns the AllowReviewerReassignment field if non-nil, zero value otherwise.

### GetAllowReviewerReassignmentOk

`func (o *UpdateCampaignConfigurationInfo) GetAllowReviewerReassignmentOk() (*bool, bool)`

GetAllowReviewerReassignmentOk returns a tuple with the AllowReviewerReassignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReviewerReassignment

`func (o *UpdateCampaignConfigurationInfo) SetAllowReviewerReassignment(v bool)`

SetAllowReviewerReassignment sets AllowReviewerReassignment field to given value.

### HasAllowReviewerReassignment

`func (o *UpdateCampaignConfigurationInfo) HasAllowReviewerReassignment() bool`

HasAllowReviewerReassignment returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateCampaignConfigurationInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateCampaignConfigurationInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateCampaignConfigurationInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateCampaignConfigurationInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateCampaignConfigurationInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateCampaignConfigurationInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateCampaignConfigurationInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateCampaignConfigurationInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdateCampaignConfigurationInfo) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateCampaignConfigurationInfo) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateCampaignConfigurationInfo) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateCampaignConfigurationInfo) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetRevokeOn

`func (o *UpdateCampaignConfigurationInfo) GetRevokeOn() CampaignRevokeOnEnum`

GetRevokeOn returns the RevokeOn field if non-nil, zero value otherwise.

### GetRevokeOnOk

`func (o *UpdateCampaignConfigurationInfo) GetRevokeOnOk() (*CampaignRevokeOnEnum, bool)`

GetRevokeOnOk returns a tuple with the RevokeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeOn

`func (o *UpdateCampaignConfigurationInfo) SetRevokeOn(v CampaignRevokeOnEnum)`

SetRevokeOn sets RevokeOn field to given value.

### HasRevokeOn

`func (o *UpdateCampaignConfigurationInfo) HasRevokeOn() bool`

HasRevokeOn returns a boolean if a field has been set.

### GetReminderSchedule

`func (o *UpdateCampaignConfigurationInfo) GetReminderSchedule() []int32`

GetReminderSchedule returns the ReminderSchedule field if non-nil, zero value otherwise.

### GetReminderScheduleOk

`func (o *UpdateCampaignConfigurationInfo) GetReminderScheduleOk() (*[]int32, bool)`

GetReminderScheduleOk returns a tuple with the ReminderSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderSchedule

`func (o *UpdateCampaignConfigurationInfo) SetReminderSchedule(v []int32)`

SetReminderSchedule sets ReminderSchedule field to given value.

### HasReminderSchedule

`func (o *UpdateCampaignConfigurationInfo) HasReminderSchedule() bool`

HasReminderSchedule returns a boolean if a field has been set.

### GetReminderIncludeManager

`func (o *UpdateCampaignConfigurationInfo) GetReminderIncludeManager() bool`

GetReminderIncludeManager returns the ReminderIncludeManager field if non-nil, zero value otherwise.

### GetReminderIncludeManagerOk

`func (o *UpdateCampaignConfigurationInfo) GetReminderIncludeManagerOk() (*bool, bool)`

GetReminderIncludeManagerOk returns a tuple with the ReminderIncludeManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderIncludeManager

`func (o *UpdateCampaignConfigurationInfo) SetReminderIncludeManager(v bool)`

SetReminderIncludeManager sets ReminderIncludeManager field to given value.

### HasReminderIncludeManager

`func (o *UpdateCampaignConfigurationInfo) HasReminderIncludeManager() bool`

HasReminderIncludeManager returns a boolean if a field has been set.

### GetRequireReasonOnDenial

`func (o *UpdateCampaignConfigurationInfo) GetRequireReasonOnDenial() bool`

GetRequireReasonOnDenial returns the RequireReasonOnDenial field if non-nil, zero value otherwise.

### GetRequireReasonOnDenialOk

`func (o *UpdateCampaignConfigurationInfo) GetRequireReasonOnDenialOk() (*bool, bool)`

GetRequireReasonOnDenialOk returns a tuple with the RequireReasonOnDenial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireReasonOnDenial

`func (o *UpdateCampaignConfigurationInfo) SetRequireReasonOnDenial(v bool)`

SetRequireReasonOnDenial sets RequireReasonOnDenial field to given value.

### HasRequireReasonOnDenial

`func (o *UpdateCampaignConfigurationInfo) HasRequireReasonOnDenial() bool`

HasRequireReasonOnDenial returns a boolean if a field has been set.

### GetHideAiSuggestions

`func (o *UpdateCampaignConfigurationInfo) GetHideAiSuggestions() bool`

GetHideAiSuggestions returns the HideAiSuggestions field if non-nil, zero value otherwise.

### GetHideAiSuggestionsOk

`func (o *UpdateCampaignConfigurationInfo) GetHideAiSuggestionsOk() (*bool, bool)`

GetHideAiSuggestionsOk returns a tuple with the HideAiSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAiSuggestions

`func (o *UpdateCampaignConfigurationInfo) SetHideAiSuggestions(v bool)`

SetHideAiSuggestions sets HideAiSuggestions field to given value.

### HasHideAiSuggestions

`func (o *UpdateCampaignConfigurationInfo) HasHideAiSuggestions() bool`

HasHideAiSuggestions returns a boolean if a field has been set.

### GetCustomStartMessage

`func (o *UpdateCampaignConfigurationInfo) GetCustomStartMessage() string`

GetCustomStartMessage returns the CustomStartMessage field if non-nil, zero value otherwise.

### GetCustomStartMessageOk

`func (o *UpdateCampaignConfigurationInfo) GetCustomStartMessageOk() (*string, bool)`

GetCustomStartMessageOk returns a tuple with the CustomStartMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStartMessage

`func (o *UpdateCampaignConfigurationInfo) SetCustomStartMessage(v string)`

SetCustomStartMessage sets CustomStartMessage field to given value.

### HasCustomStartMessage

`func (o *UpdateCampaignConfigurationInfo) HasCustomStartMessage() bool`

HasCustomStartMessage returns a boolean if a field has been set.

### GetGroupAssetVisibilityPolicy

`func (o *UpdateCampaignConfigurationInfo) GetGroupAssetVisibilityPolicy() CampaignGroupAssetVisibilityPolicyEnum`

GetGroupAssetVisibilityPolicy returns the GroupAssetVisibilityPolicy field if non-nil, zero value otherwise.

### GetGroupAssetVisibilityPolicyOk

`func (o *UpdateCampaignConfigurationInfo) GetGroupAssetVisibilityPolicyOk() (*CampaignGroupAssetVisibilityPolicyEnum, bool)`

GetGroupAssetVisibilityPolicyOk returns a tuple with the GroupAssetVisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAssetVisibilityPolicy

`func (o *UpdateCampaignConfigurationInfo) SetGroupAssetVisibilityPolicy(v CampaignGroupAssetVisibilityPolicyEnum)`

SetGroupAssetVisibilityPolicy sets GroupAssetVisibilityPolicy field to given value.

### HasGroupAssetVisibilityPolicy

`func (o *UpdateCampaignConfigurationInfo) HasGroupAssetVisibilityPolicy() bool`

HasGroupAssetVisibilityPolicy returns a boolean if a field has been set.

### GetCronExpression

`func (o *UpdateCampaignConfigurationInfo) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *UpdateCampaignConfigurationInfo) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *UpdateCampaignConfigurationInfo) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *UpdateCampaignConfigurationInfo) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetRecurringDurationDays

`func (o *UpdateCampaignConfigurationInfo) GetRecurringDurationDays() int32`

GetRecurringDurationDays returns the RecurringDurationDays field if non-nil, zero value otherwise.

### GetRecurringDurationDaysOk

`func (o *UpdateCampaignConfigurationInfo) GetRecurringDurationDaysOk() (*int32, bool)`

GetRecurringDurationDaysOk returns a tuple with the RecurringDurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDurationDays

`func (o *UpdateCampaignConfigurationInfo) SetRecurringDurationDays(v int32)`

SetRecurringDurationDays sets RecurringDurationDays field to given value.

### HasRecurringDurationDays

`func (o *UpdateCampaignConfigurationInfo) HasRecurringDurationDays() bool`

HasRecurringDurationDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


