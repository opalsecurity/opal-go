# UpdateOwnerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | **string** | The ID of the owner. | 
**Name** | Pointer to **string** | The name of the owner. | [optional] 
**Description** | Pointer to **string** | A description of the owner. | [optional] 
**AccessRequestEscalationPeriod** | Pointer to **int32** | The amount of time (in minutes) before the next reviewer is notified. Use 0 to remove escalation policy. | [optional] 
**ReviewerMessageChannelId** | Pointer to **string** | The message channel id for the reviewer channel. Use \&quot;\&quot; to remove an existing message channel. | [optional] 

## Methods

### NewUpdateOwnerInfo

`func NewUpdateOwnerInfo(ownerId string, ) *UpdateOwnerInfo`

NewUpdateOwnerInfo instantiates a new UpdateOwnerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOwnerInfoWithDefaults

`func NewUpdateOwnerInfoWithDefaults() *UpdateOwnerInfo`

NewUpdateOwnerInfoWithDefaults instantiates a new UpdateOwnerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *UpdateOwnerInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UpdateOwnerInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UpdateOwnerInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetName

`func (o *UpdateOwnerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOwnerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOwnerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOwnerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOwnerInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOwnerInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOwnerInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOwnerInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessRequestEscalationPeriod

`func (o *UpdateOwnerInfo) GetAccessRequestEscalationPeriod() int32`

GetAccessRequestEscalationPeriod returns the AccessRequestEscalationPeriod field if non-nil, zero value otherwise.

### GetAccessRequestEscalationPeriodOk

`func (o *UpdateOwnerInfo) GetAccessRequestEscalationPeriodOk() (*int32, bool)`

GetAccessRequestEscalationPeriodOk returns a tuple with the AccessRequestEscalationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestEscalationPeriod

`func (o *UpdateOwnerInfo) SetAccessRequestEscalationPeriod(v int32)`

SetAccessRequestEscalationPeriod sets AccessRequestEscalationPeriod field to given value.

### HasAccessRequestEscalationPeriod

`func (o *UpdateOwnerInfo) HasAccessRequestEscalationPeriod() bool`

HasAccessRequestEscalationPeriod returns a boolean if a field has been set.

### GetReviewerMessageChannelId

`func (o *UpdateOwnerInfo) GetReviewerMessageChannelId() string`

GetReviewerMessageChannelId returns the ReviewerMessageChannelId field if non-nil, zero value otherwise.

### GetReviewerMessageChannelIdOk

`func (o *UpdateOwnerInfo) GetReviewerMessageChannelIdOk() (*string, bool)`

GetReviewerMessageChannelIdOk returns a tuple with the ReviewerMessageChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerMessageChannelId

`func (o *UpdateOwnerInfo) SetReviewerMessageChannelId(v string)`

SetReviewerMessageChannelId sets ReviewerMessageChannelId field to given value.

### HasReviewerMessageChannelId

`func (o *UpdateOwnerInfo) HasReviewerMessageChannelId() bool`

HasReviewerMessageChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


