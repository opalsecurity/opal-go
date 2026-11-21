# CampaignItemAdminOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorUserId** | **string** | The admin who applied the override. | 
**Decision** | [**CampaignItemReviewDecisionEnum**](CampaignItemReviewDecisionEnum.md) |  | 
**DecidedAt** | **time.Time** | When the override was applied. | 

## Methods

### NewCampaignItemAdminOverride

`func NewCampaignItemAdminOverride(actorUserId string, decision CampaignItemReviewDecisionEnum, decidedAt time.Time, ) *CampaignItemAdminOverride`

NewCampaignItemAdminOverride instantiates a new CampaignItemAdminOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignItemAdminOverrideWithDefaults

`func NewCampaignItemAdminOverrideWithDefaults() *CampaignItemAdminOverride`

NewCampaignItemAdminOverrideWithDefaults instantiates a new CampaignItemAdminOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorUserId

`func (o *CampaignItemAdminOverride) GetActorUserId() string`

GetActorUserId returns the ActorUserId field if non-nil, zero value otherwise.

### GetActorUserIdOk

`func (o *CampaignItemAdminOverride) GetActorUserIdOk() (*string, bool)`

GetActorUserIdOk returns a tuple with the ActorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUserId

`func (o *CampaignItemAdminOverride) SetActorUserId(v string)`

SetActorUserId sets ActorUserId field to given value.


### GetDecision

`func (o *CampaignItemAdminOverride) GetDecision() CampaignItemReviewDecisionEnum`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *CampaignItemAdminOverride) GetDecisionOk() (*CampaignItemReviewDecisionEnum, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *CampaignItemAdminOverride) SetDecision(v CampaignItemReviewDecisionEnum)`

SetDecision sets Decision field to given value.


### GetDecidedAt

`func (o *CampaignItemAdminOverride) GetDecidedAt() time.Time`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *CampaignItemAdminOverride) GetDecidedAtOk() (*time.Time, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *CampaignItemAdminOverride) SetDecidedAt(v time.Time)`

SetDecidedAt sets DecidedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


