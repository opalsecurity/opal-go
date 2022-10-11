# CreateOwnerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the owner. | 
**Description** | Pointer to **string** | A description of the owner. | [optional] 
**AccessRequestEscalationPeriod** | Pointer to **int32** | The amount of time (in minutes) before the next reviewer is notified. Use 0 to remove escalation policy. | [optional] 
**UserIds** | **[]string** | Users to add to the created owner. | 
**ReviewerMessageChannelId** | Pointer to **string** | The message channel id for the reviewer channel. | [optional] 

## Methods

### NewCreateOwnerInfo

`func NewCreateOwnerInfo(name string, userIds []string, ) *CreateOwnerInfo`

NewCreateOwnerInfo instantiates a new CreateOwnerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOwnerInfoWithDefaults

`func NewCreateOwnerInfoWithDefaults() *CreateOwnerInfo`

NewCreateOwnerInfoWithDefaults instantiates a new CreateOwnerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOwnerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOwnerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOwnerInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateOwnerInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOwnerInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOwnerInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOwnerInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessRequestEscalationPeriod

`func (o *CreateOwnerInfo) GetAccessRequestEscalationPeriod() int32`

GetAccessRequestEscalationPeriod returns the AccessRequestEscalationPeriod field if non-nil, zero value otherwise.

### GetAccessRequestEscalationPeriodOk

`func (o *CreateOwnerInfo) GetAccessRequestEscalationPeriodOk() (*int32, bool)`

GetAccessRequestEscalationPeriodOk returns a tuple with the AccessRequestEscalationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestEscalationPeriod

`func (o *CreateOwnerInfo) SetAccessRequestEscalationPeriod(v int32)`

SetAccessRequestEscalationPeriod sets AccessRequestEscalationPeriod field to given value.

### HasAccessRequestEscalationPeriod

`func (o *CreateOwnerInfo) HasAccessRequestEscalationPeriod() bool`

HasAccessRequestEscalationPeriod returns a boolean if a field has been set.

### GetUserIds

`func (o *CreateOwnerInfo) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *CreateOwnerInfo) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *CreateOwnerInfo) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.


### GetReviewerMessageChannelId

`func (o *CreateOwnerInfo) GetReviewerMessageChannelId() string`

GetReviewerMessageChannelId returns the ReviewerMessageChannelId field if non-nil, zero value otherwise.

### GetReviewerMessageChannelIdOk

`func (o *CreateOwnerInfo) GetReviewerMessageChannelIdOk() (*string, bool)`

GetReviewerMessageChannelIdOk returns a tuple with the ReviewerMessageChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerMessageChannelId

`func (o *CreateOwnerInfo) SetReviewerMessageChannelId(v string)`

SetReviewerMessageChannelId sets ReviewerMessageChannelId field to given value.

### HasReviewerMessageChannelId

`func (o *CreateOwnerInfo) HasReviewerMessageChannelId() bool`

HasReviewerMessageChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


