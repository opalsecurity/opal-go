# CreateUARInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the UAR. | 
**ReviewerAssignmentPolicy** | [**UARReviewerAssignmentPolicyEnum**](UARReviewerAssignmentPolicyEnum.md) |  | 
**SendReviewerAssignmentNotification** | **bool** | A bool representing whether to send a notification to reviewers when they&#39;re assigned a new review. Default is False. | 
**Deadline** | **time.Time** | The last day for reviewers to complete their access reviews. | 
**TimeZone** | **string** | The time zone name (as defined by the IANA Time Zone database) used in the access review deadline and exported audit report. Default is America/Los_Angeles. | 
**SelfReviewAllowed** | **bool** | A bool representing whether to present a warning when a user is the only reviewer for themself. Default is False. | 
**ReminderSchedule** | Pointer to **[]int32** |  | [optional] 
**ReminderIncludeManager** | Pointer to **bool** |  | [optional] 
**UarScope** | Pointer to [**UARScope**](UARScope.md) |  | [optional] 

## Methods

### NewCreateUARInfo

`func NewCreateUARInfo(name string, reviewerAssignmentPolicy UARReviewerAssignmentPolicyEnum, sendReviewerAssignmentNotification bool, deadline time.Time, timeZone string, selfReviewAllowed bool, ) *CreateUARInfo`

NewCreateUARInfo instantiates a new CreateUARInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUARInfoWithDefaults

`func NewCreateUARInfoWithDefaults() *CreateUARInfo`

NewCreateUARInfoWithDefaults instantiates a new CreateUARInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUARInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUARInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUARInfo) SetName(v string)`

SetName sets Name field to given value.


### GetReviewerAssignmentPolicy

`func (o *CreateUARInfo) GetReviewerAssignmentPolicy() UARReviewerAssignmentPolicyEnum`

GetReviewerAssignmentPolicy returns the ReviewerAssignmentPolicy field if non-nil, zero value otherwise.

### GetReviewerAssignmentPolicyOk

`func (o *CreateUARInfo) GetReviewerAssignmentPolicyOk() (*UARReviewerAssignmentPolicyEnum, bool)`

GetReviewerAssignmentPolicyOk returns a tuple with the ReviewerAssignmentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerAssignmentPolicy

`func (o *CreateUARInfo) SetReviewerAssignmentPolicy(v UARReviewerAssignmentPolicyEnum)`

SetReviewerAssignmentPolicy sets ReviewerAssignmentPolicy field to given value.


### GetSendReviewerAssignmentNotification

`func (o *CreateUARInfo) GetSendReviewerAssignmentNotification() bool`

GetSendReviewerAssignmentNotification returns the SendReviewerAssignmentNotification field if non-nil, zero value otherwise.

### GetSendReviewerAssignmentNotificationOk

`func (o *CreateUARInfo) GetSendReviewerAssignmentNotificationOk() (*bool, bool)`

GetSendReviewerAssignmentNotificationOk returns a tuple with the SendReviewerAssignmentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReviewerAssignmentNotification

`func (o *CreateUARInfo) SetSendReviewerAssignmentNotification(v bool)`

SetSendReviewerAssignmentNotification sets SendReviewerAssignmentNotification field to given value.


### GetDeadline

`func (o *CreateUARInfo) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *CreateUARInfo) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *CreateUARInfo) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### GetTimeZone

`func (o *CreateUARInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CreateUARInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CreateUARInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetSelfReviewAllowed

`func (o *CreateUARInfo) GetSelfReviewAllowed() bool`

GetSelfReviewAllowed returns the SelfReviewAllowed field if non-nil, zero value otherwise.

### GetSelfReviewAllowedOk

`func (o *CreateUARInfo) GetSelfReviewAllowedOk() (*bool, bool)`

GetSelfReviewAllowedOk returns a tuple with the SelfReviewAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfReviewAllowed

`func (o *CreateUARInfo) SetSelfReviewAllowed(v bool)`

SetSelfReviewAllowed sets SelfReviewAllowed field to given value.


### GetReminderSchedule

`func (o *CreateUARInfo) GetReminderSchedule() []int32`

GetReminderSchedule returns the ReminderSchedule field if non-nil, zero value otherwise.

### GetReminderScheduleOk

`func (o *CreateUARInfo) GetReminderScheduleOk() (*[]int32, bool)`

GetReminderScheduleOk returns a tuple with the ReminderSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderSchedule

`func (o *CreateUARInfo) SetReminderSchedule(v []int32)`

SetReminderSchedule sets ReminderSchedule field to given value.

### HasReminderSchedule

`func (o *CreateUARInfo) HasReminderSchedule() bool`

HasReminderSchedule returns a boolean if a field has been set.

### GetReminderIncludeManager

`func (o *CreateUARInfo) GetReminderIncludeManager() bool`

GetReminderIncludeManager returns the ReminderIncludeManager field if non-nil, zero value otherwise.

### GetReminderIncludeManagerOk

`func (o *CreateUARInfo) GetReminderIncludeManagerOk() (*bool, bool)`

GetReminderIncludeManagerOk returns a tuple with the ReminderIncludeManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderIncludeManager

`func (o *CreateUARInfo) SetReminderIncludeManager(v bool)`

SetReminderIncludeManager sets ReminderIncludeManager field to given value.

### HasReminderIncludeManager

`func (o *CreateUARInfo) HasReminderIncludeManager() bool`

HasReminderIncludeManager returns a boolean if a field has been set.

### GetUarScope

`func (o *CreateUARInfo) GetUarScope() UARScope`

GetUarScope returns the UarScope field if non-nil, zero value otherwise.

### GetUarScopeOk

`func (o *CreateUARInfo) GetUarScopeOk() (*UARScope, bool)`

GetUarScopeOk returns a tuple with the UarScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUarScope

`func (o *CreateUARInfo) SetUarScope(v UARScope)`

SetUarScope sets UarScope field to given value.

### HasUarScope

`func (o *CreateUARInfo) HasUarScope() bool`

HasUarScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


