# UAR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UarId** | **string** | The ID of the UAR. | 
**Name** | **string** | The name of the UAR. | 
**ReviewerAssignmentPolicy** | [**UARReviewerAssignmentPolicyEnum**](UARReviewerAssignmentPolicyEnum.md) |  | 
**SendReviewerAssignmentNotification** | **bool** | A bool representing whether to send a notification to reviewers when they&#39;re assigned a new review. Default is False. | 
**Deadline** | **time.Time** | The last day for reviewers to complete their access reviews. | 
**TimeZone** | **string** | The time zone name (as defined by the IANA Time Zone database) used in the access review deadline and exported audit report. Default is America/Los_Angeles. | 
**SelfReviewAllowed** | **bool** | A bool representing whether to present a warning when a user is the only reviewer for themself. Default is False. | 
**UarScope** | Pointer to [**UARScope**](UARScope.md) |  | [optional] 

## Methods

### NewUAR

`func NewUAR(uarId string, name string, reviewerAssignmentPolicy UARReviewerAssignmentPolicyEnum, sendReviewerAssignmentNotification bool, deadline time.Time, timeZone string, selfReviewAllowed bool, ) *UAR`

NewUAR instantiates a new UAR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUARWithDefaults

`func NewUARWithDefaults() *UAR`

NewUARWithDefaults instantiates a new UAR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUarId

`func (o *UAR) GetUarId() string`

GetUarId returns the UarId field if non-nil, zero value otherwise.

### GetUarIdOk

`func (o *UAR) GetUarIdOk() (*string, bool)`

GetUarIdOk returns a tuple with the UarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUarId

`func (o *UAR) SetUarId(v string)`

SetUarId sets UarId field to given value.


### GetName

`func (o *UAR) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UAR) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UAR) SetName(v string)`

SetName sets Name field to given value.


### GetReviewerAssignmentPolicy

`func (o *UAR) GetReviewerAssignmentPolicy() UARReviewerAssignmentPolicyEnum`

GetReviewerAssignmentPolicy returns the ReviewerAssignmentPolicy field if non-nil, zero value otherwise.

### GetReviewerAssignmentPolicyOk

`func (o *UAR) GetReviewerAssignmentPolicyOk() (*UARReviewerAssignmentPolicyEnum, bool)`

GetReviewerAssignmentPolicyOk returns a tuple with the ReviewerAssignmentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerAssignmentPolicy

`func (o *UAR) SetReviewerAssignmentPolicy(v UARReviewerAssignmentPolicyEnum)`

SetReviewerAssignmentPolicy sets ReviewerAssignmentPolicy field to given value.


### GetSendReviewerAssignmentNotification

`func (o *UAR) GetSendReviewerAssignmentNotification() bool`

GetSendReviewerAssignmentNotification returns the SendReviewerAssignmentNotification field if non-nil, zero value otherwise.

### GetSendReviewerAssignmentNotificationOk

`func (o *UAR) GetSendReviewerAssignmentNotificationOk() (*bool, bool)`

GetSendReviewerAssignmentNotificationOk returns a tuple with the SendReviewerAssignmentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReviewerAssignmentNotification

`func (o *UAR) SetSendReviewerAssignmentNotification(v bool)`

SetSendReviewerAssignmentNotification sets SendReviewerAssignmentNotification field to given value.


### GetDeadline

`func (o *UAR) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *UAR) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *UAR) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### GetTimeZone

`func (o *UAR) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UAR) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UAR) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetSelfReviewAllowed

`func (o *UAR) GetSelfReviewAllowed() bool`

GetSelfReviewAllowed returns the SelfReviewAllowed field if non-nil, zero value otherwise.

### GetSelfReviewAllowedOk

`func (o *UAR) GetSelfReviewAllowedOk() (*bool, bool)`

GetSelfReviewAllowedOk returns a tuple with the SelfReviewAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfReviewAllowed

`func (o *UAR) SetSelfReviewAllowed(v bool)`

SetSelfReviewAllowed sets SelfReviewAllowed field to given value.


### GetUarScope

`func (o *UAR) GetUarScope() UARScope`

GetUarScope returns the UarScope field if non-nil, zero value otherwise.

### GetUarScopeOk

`func (o *UAR) GetUarScopeOk() (*UARScope, bool)`

GetUarScopeOk returns a tuple with the UarScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUarScope

`func (o *UAR) SetUarScope(v UARScope)`

SetUarScope sets UarScope field to given value.

### HasUarScope

`func (o *UAR) HasUarScope() bool`

HasUarScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


