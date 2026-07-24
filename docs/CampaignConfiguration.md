# CampaignConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationId** | **string** | The ID of the campaign configuration. | 
**CreatedAt** | **time.Time** | The creation time of the configuration. | 
**UpdatedAt** | **time.Time** | The last updated time of the configuration. | 
**Query** | Pointer to **map[string]interface{}** | Edge-based query defining the scope of access to review. Matches the stored Opal edge-query JSON shape. | [optional] 
**ReviewerAssignmentPolicy** | [**UARReviewerAssignmentPolicyEnum**](UARReviewerAssignmentPolicyEnum.md) |  | 
**AllowSelfReview** | **bool** | Whether reviewers can review their own access. | 
**SendReviewerAssignmentNotification** | **bool** | Whether to notify reviewers upon assignment. | 
**AllowReviewerReassignment** | **bool** | Whether reviewers may reassign their reviews to another user. | 
**StartDate** | Pointer to **time.Time** | Scheduled start date of the campaign. | [optional] 
**EndDate** | Pointer to **time.Time** | Scheduled end date of the campaign. | [optional] 
**Timezone** | **string** | IANA timezone used to interpret campaign deadlines (e.g. America/Los_Angeles). | 
**RevokeOn** | [**CampaignRevokeOnEnum**](CampaignRevokeOnEnum.md) |  | 
**ReminderSchedule** | Pointer to **[]int32** | Days before end date to send reminder notifications. | [optional] 
**ReminderIncludeManager** | **bool** | Whether to include the reviewer&#39;s manager in reminders. | 
**RequireReasonOnDenial** | **bool** | Whether reviewers must provide a reason when denying (revoking) access. | 
**HideAiSuggestions** | **bool** | Whether AI suggestions are hidden from reviewers. | 
**CustomStartMessage** | Pointer to **string** | Optional custom message included when notifying reviewers that the campaign started. | [optional] 
**GroupAssetVisibilityPolicy** | [**CampaignGroupAssetVisibilityPolicyEnum**](CampaignGroupAssetVisibilityPolicyEnum.md) |  | 
**IsTemplate** | **bool** | Whether this configuration is a recurring schedule template. | 
**CronExpression** | Pointer to **string** | Cron expression driving the recurring schedule. Null for one-off campaigns. | [optional] 
**NextScheduledRun** | Pointer to **time.Time** | Next time a draft will be generated from this template. | [optional] 
**LastScheduledRun** | Pointer to **time.Time** | Most recent time a draft was generated from this template. | [optional] 
**RecurringDurationDays** | Pointer to **int32** | Deadline window in days applied to each draft generated from this template. | [optional] 

## Methods

### NewCampaignConfiguration

`func NewCampaignConfiguration(configurationId string, createdAt time.Time, updatedAt time.Time, reviewerAssignmentPolicy UARReviewerAssignmentPolicyEnum, allowSelfReview bool, sendReviewerAssignmentNotification bool, allowReviewerReassignment bool, timezone string, revokeOn CampaignRevokeOnEnum, reminderIncludeManager bool, requireReasonOnDenial bool, hideAiSuggestions bool, groupAssetVisibilityPolicy CampaignGroupAssetVisibilityPolicyEnum, isTemplate bool, ) *CampaignConfiguration`

NewCampaignConfiguration instantiates a new CampaignConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignConfigurationWithDefaults

`func NewCampaignConfigurationWithDefaults() *CampaignConfiguration`

NewCampaignConfigurationWithDefaults instantiates a new CampaignConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationId

`func (o *CampaignConfiguration) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *CampaignConfiguration) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *CampaignConfiguration) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetCreatedAt

`func (o *CampaignConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CampaignConfiguration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CampaignConfiguration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CampaignConfiguration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetQuery

`func (o *CampaignConfiguration) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CampaignConfiguration) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CampaignConfiguration) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CampaignConfiguration) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetReviewerAssignmentPolicy

`func (o *CampaignConfiguration) GetReviewerAssignmentPolicy() UARReviewerAssignmentPolicyEnum`

GetReviewerAssignmentPolicy returns the ReviewerAssignmentPolicy field if non-nil, zero value otherwise.

### GetReviewerAssignmentPolicyOk

`func (o *CampaignConfiguration) GetReviewerAssignmentPolicyOk() (*UARReviewerAssignmentPolicyEnum, bool)`

GetReviewerAssignmentPolicyOk returns a tuple with the ReviewerAssignmentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerAssignmentPolicy

`func (o *CampaignConfiguration) SetReviewerAssignmentPolicy(v UARReviewerAssignmentPolicyEnum)`

SetReviewerAssignmentPolicy sets ReviewerAssignmentPolicy field to given value.


### GetAllowSelfReview

`func (o *CampaignConfiguration) GetAllowSelfReview() bool`

GetAllowSelfReview returns the AllowSelfReview field if non-nil, zero value otherwise.

### GetAllowSelfReviewOk

`func (o *CampaignConfiguration) GetAllowSelfReviewOk() (*bool, bool)`

GetAllowSelfReviewOk returns a tuple with the AllowSelfReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSelfReview

`func (o *CampaignConfiguration) SetAllowSelfReview(v bool)`

SetAllowSelfReview sets AllowSelfReview field to given value.


### GetSendReviewerAssignmentNotification

`func (o *CampaignConfiguration) GetSendReviewerAssignmentNotification() bool`

GetSendReviewerAssignmentNotification returns the SendReviewerAssignmentNotification field if non-nil, zero value otherwise.

### GetSendReviewerAssignmentNotificationOk

`func (o *CampaignConfiguration) GetSendReviewerAssignmentNotificationOk() (*bool, bool)`

GetSendReviewerAssignmentNotificationOk returns a tuple with the SendReviewerAssignmentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReviewerAssignmentNotification

`func (o *CampaignConfiguration) SetSendReviewerAssignmentNotification(v bool)`

SetSendReviewerAssignmentNotification sets SendReviewerAssignmentNotification field to given value.


### GetAllowReviewerReassignment

`func (o *CampaignConfiguration) GetAllowReviewerReassignment() bool`

GetAllowReviewerReassignment returns the AllowReviewerReassignment field if non-nil, zero value otherwise.

### GetAllowReviewerReassignmentOk

`func (o *CampaignConfiguration) GetAllowReviewerReassignmentOk() (*bool, bool)`

GetAllowReviewerReassignmentOk returns a tuple with the AllowReviewerReassignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReviewerReassignment

`func (o *CampaignConfiguration) SetAllowReviewerReassignment(v bool)`

SetAllowReviewerReassignment sets AllowReviewerReassignment field to given value.


### GetStartDate

`func (o *CampaignConfiguration) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CampaignConfiguration) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CampaignConfiguration) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CampaignConfiguration) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CampaignConfiguration) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CampaignConfiguration) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CampaignConfiguration) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CampaignConfiguration) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTimezone

`func (o *CampaignConfiguration) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CampaignConfiguration) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CampaignConfiguration) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetRevokeOn

`func (o *CampaignConfiguration) GetRevokeOn() CampaignRevokeOnEnum`

GetRevokeOn returns the RevokeOn field if non-nil, zero value otherwise.

### GetRevokeOnOk

`func (o *CampaignConfiguration) GetRevokeOnOk() (*CampaignRevokeOnEnum, bool)`

GetRevokeOnOk returns a tuple with the RevokeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeOn

`func (o *CampaignConfiguration) SetRevokeOn(v CampaignRevokeOnEnum)`

SetRevokeOn sets RevokeOn field to given value.


### GetReminderSchedule

`func (o *CampaignConfiguration) GetReminderSchedule() []int32`

GetReminderSchedule returns the ReminderSchedule field if non-nil, zero value otherwise.

### GetReminderScheduleOk

`func (o *CampaignConfiguration) GetReminderScheduleOk() (*[]int32, bool)`

GetReminderScheduleOk returns a tuple with the ReminderSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderSchedule

`func (o *CampaignConfiguration) SetReminderSchedule(v []int32)`

SetReminderSchedule sets ReminderSchedule field to given value.

### HasReminderSchedule

`func (o *CampaignConfiguration) HasReminderSchedule() bool`

HasReminderSchedule returns a boolean if a field has been set.

### GetReminderIncludeManager

`func (o *CampaignConfiguration) GetReminderIncludeManager() bool`

GetReminderIncludeManager returns the ReminderIncludeManager field if non-nil, zero value otherwise.

### GetReminderIncludeManagerOk

`func (o *CampaignConfiguration) GetReminderIncludeManagerOk() (*bool, bool)`

GetReminderIncludeManagerOk returns a tuple with the ReminderIncludeManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderIncludeManager

`func (o *CampaignConfiguration) SetReminderIncludeManager(v bool)`

SetReminderIncludeManager sets ReminderIncludeManager field to given value.


### GetRequireReasonOnDenial

`func (o *CampaignConfiguration) GetRequireReasonOnDenial() bool`

GetRequireReasonOnDenial returns the RequireReasonOnDenial field if non-nil, zero value otherwise.

### GetRequireReasonOnDenialOk

`func (o *CampaignConfiguration) GetRequireReasonOnDenialOk() (*bool, bool)`

GetRequireReasonOnDenialOk returns a tuple with the RequireReasonOnDenial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireReasonOnDenial

`func (o *CampaignConfiguration) SetRequireReasonOnDenial(v bool)`

SetRequireReasonOnDenial sets RequireReasonOnDenial field to given value.


### GetHideAiSuggestions

`func (o *CampaignConfiguration) GetHideAiSuggestions() bool`

GetHideAiSuggestions returns the HideAiSuggestions field if non-nil, zero value otherwise.

### GetHideAiSuggestionsOk

`func (o *CampaignConfiguration) GetHideAiSuggestionsOk() (*bool, bool)`

GetHideAiSuggestionsOk returns a tuple with the HideAiSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAiSuggestions

`func (o *CampaignConfiguration) SetHideAiSuggestions(v bool)`

SetHideAiSuggestions sets HideAiSuggestions field to given value.


### GetCustomStartMessage

`func (o *CampaignConfiguration) GetCustomStartMessage() string`

GetCustomStartMessage returns the CustomStartMessage field if non-nil, zero value otherwise.

### GetCustomStartMessageOk

`func (o *CampaignConfiguration) GetCustomStartMessageOk() (*string, bool)`

GetCustomStartMessageOk returns a tuple with the CustomStartMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStartMessage

`func (o *CampaignConfiguration) SetCustomStartMessage(v string)`

SetCustomStartMessage sets CustomStartMessage field to given value.

### HasCustomStartMessage

`func (o *CampaignConfiguration) HasCustomStartMessage() bool`

HasCustomStartMessage returns a boolean if a field has been set.

### GetGroupAssetVisibilityPolicy

`func (o *CampaignConfiguration) GetGroupAssetVisibilityPolicy() CampaignGroupAssetVisibilityPolicyEnum`

GetGroupAssetVisibilityPolicy returns the GroupAssetVisibilityPolicy field if non-nil, zero value otherwise.

### GetGroupAssetVisibilityPolicyOk

`func (o *CampaignConfiguration) GetGroupAssetVisibilityPolicyOk() (*CampaignGroupAssetVisibilityPolicyEnum, bool)`

GetGroupAssetVisibilityPolicyOk returns a tuple with the GroupAssetVisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAssetVisibilityPolicy

`func (o *CampaignConfiguration) SetGroupAssetVisibilityPolicy(v CampaignGroupAssetVisibilityPolicyEnum)`

SetGroupAssetVisibilityPolicy sets GroupAssetVisibilityPolicy field to given value.


### GetIsTemplate

`func (o *CampaignConfiguration) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *CampaignConfiguration) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *CampaignConfiguration) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.


### GetCronExpression

`func (o *CampaignConfiguration) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *CampaignConfiguration) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *CampaignConfiguration) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *CampaignConfiguration) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetNextScheduledRun

`func (o *CampaignConfiguration) GetNextScheduledRun() time.Time`

GetNextScheduledRun returns the NextScheduledRun field if non-nil, zero value otherwise.

### GetNextScheduledRunOk

`func (o *CampaignConfiguration) GetNextScheduledRunOk() (*time.Time, bool)`

GetNextScheduledRunOk returns a tuple with the NextScheduledRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduledRun

`func (o *CampaignConfiguration) SetNextScheduledRun(v time.Time)`

SetNextScheduledRun sets NextScheduledRun field to given value.

### HasNextScheduledRun

`func (o *CampaignConfiguration) HasNextScheduledRun() bool`

HasNextScheduledRun returns a boolean if a field has been set.

### GetLastScheduledRun

`func (o *CampaignConfiguration) GetLastScheduledRun() time.Time`

GetLastScheduledRun returns the LastScheduledRun field if non-nil, zero value otherwise.

### GetLastScheduledRunOk

`func (o *CampaignConfiguration) GetLastScheduledRunOk() (*time.Time, bool)`

GetLastScheduledRunOk returns a tuple with the LastScheduledRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScheduledRun

`func (o *CampaignConfiguration) SetLastScheduledRun(v time.Time)`

SetLastScheduledRun sets LastScheduledRun field to given value.

### HasLastScheduledRun

`func (o *CampaignConfiguration) HasLastScheduledRun() bool`

HasLastScheduledRun returns a boolean if a field has been set.

### GetRecurringDurationDays

`func (o *CampaignConfiguration) GetRecurringDurationDays() int32`

GetRecurringDurationDays returns the RecurringDurationDays field if non-nil, zero value otherwise.

### GetRecurringDurationDaysOk

`func (o *CampaignConfiguration) GetRecurringDurationDaysOk() (*int32, bool)`

GetRecurringDurationDaysOk returns a tuple with the RecurringDurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDurationDays

`func (o *CampaignConfiguration) SetRecurringDurationDays(v int32)`

SetRecurringDurationDays sets RecurringDurationDays field to given value.

### HasRecurringDurationDays

`func (o *CampaignConfiguration) HasRecurringDurationDays() bool`

HasRecurringDurationDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


