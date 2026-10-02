# ReviewerStageEscalation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayMinutes** | **int32** | How long to wait for a response before escalating, in minutes. Between 1 and 1440 (24 hours). | 
**OwnerIds** | Pointer to **[]string** | The owners to escalate to. The stage&#39;s own owner_ids are added automatically and must not be repeated here. | [optional] 
**UserIds** | Pointer to **[]string** | The users to escalate to. The stage&#39;s own service_user_ids are added automatically and must not be repeated here. | [optional] 

## Methods

### NewReviewerStageEscalation

`func NewReviewerStageEscalation(delayMinutes int32, ) *ReviewerStageEscalation`

NewReviewerStageEscalation instantiates a new ReviewerStageEscalation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerStageEscalationWithDefaults

`func NewReviewerStageEscalationWithDefaults() *ReviewerStageEscalation`

NewReviewerStageEscalationWithDefaults instantiates a new ReviewerStageEscalation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayMinutes

`func (o *ReviewerStageEscalation) GetDelayMinutes() int32`

GetDelayMinutes returns the DelayMinutes field if non-nil, zero value otherwise.

### GetDelayMinutesOk

`func (o *ReviewerStageEscalation) GetDelayMinutesOk() (*int32, bool)`

GetDelayMinutesOk returns a tuple with the DelayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMinutes

`func (o *ReviewerStageEscalation) SetDelayMinutes(v int32)`

SetDelayMinutes sets DelayMinutes field to given value.


### GetOwnerIds

`func (o *ReviewerStageEscalation) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *ReviewerStageEscalation) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *ReviewerStageEscalation) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *ReviewerStageEscalation) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetUserIds

`func (o *ReviewerStageEscalation) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *ReviewerStageEscalation) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *ReviewerStageEscalation) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *ReviewerStageEscalation) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


