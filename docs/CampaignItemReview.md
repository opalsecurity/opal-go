# CampaignItemReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignItemReviewId** | **string** | The ID of the campaign item review. | 
**CampaignItemId** | **string** | The ID of the campaign item. | 
**ReviewerUserId** | **string** | The ID of the assigned reviewer. | 
**AssignmentSource** | [**CampaignItemReviewerAssignmentSourceEnum**](CampaignItemReviewerAssignmentSourceEnum.md) |  | 
**Decision** | Pointer to [**CampaignItemReviewDecisionEnum**](CampaignItemReviewDecisionEnum.md) | The reviewer&#39;s decision. Null when the reviewer has not yet submitted.  | [optional] 
**Note** | Pointer to **string** | Optional note from the reviewer. | [optional] 
**UpdatedAccessLevelRemoteId** | Pointer to **string** | Target access level remote ID when decision is CHANGE_ROLE. | [optional] 
**DecidedAt** | Pointer to **time.Time** | When the decision was finalized. | [optional] 

## Methods

### NewCampaignItemReview

`func NewCampaignItemReview(campaignItemReviewId string, campaignItemId string, reviewerUserId string, assignmentSource CampaignItemReviewerAssignmentSourceEnum, ) *CampaignItemReview`

NewCampaignItemReview instantiates a new CampaignItemReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignItemReviewWithDefaults

`func NewCampaignItemReviewWithDefaults() *CampaignItemReview`

NewCampaignItemReviewWithDefaults instantiates a new CampaignItemReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignItemReviewId

`func (o *CampaignItemReview) GetCampaignItemReviewId() string`

GetCampaignItemReviewId returns the CampaignItemReviewId field if non-nil, zero value otherwise.

### GetCampaignItemReviewIdOk

`func (o *CampaignItemReview) GetCampaignItemReviewIdOk() (*string, bool)`

GetCampaignItemReviewIdOk returns a tuple with the CampaignItemReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignItemReviewId

`func (o *CampaignItemReview) SetCampaignItemReviewId(v string)`

SetCampaignItemReviewId sets CampaignItemReviewId field to given value.


### GetCampaignItemId

`func (o *CampaignItemReview) GetCampaignItemId() string`

GetCampaignItemId returns the CampaignItemId field if non-nil, zero value otherwise.

### GetCampaignItemIdOk

`func (o *CampaignItemReview) GetCampaignItemIdOk() (*string, bool)`

GetCampaignItemIdOk returns a tuple with the CampaignItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignItemId

`func (o *CampaignItemReview) SetCampaignItemId(v string)`

SetCampaignItemId sets CampaignItemId field to given value.


### GetReviewerUserId

`func (o *CampaignItemReview) GetReviewerUserId() string`

GetReviewerUserId returns the ReviewerUserId field if non-nil, zero value otherwise.

### GetReviewerUserIdOk

`func (o *CampaignItemReview) GetReviewerUserIdOk() (*string, bool)`

GetReviewerUserIdOk returns a tuple with the ReviewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerUserId

`func (o *CampaignItemReview) SetReviewerUserId(v string)`

SetReviewerUserId sets ReviewerUserId field to given value.


### GetAssignmentSource

`func (o *CampaignItemReview) GetAssignmentSource() CampaignItemReviewerAssignmentSourceEnum`

GetAssignmentSource returns the AssignmentSource field if non-nil, zero value otherwise.

### GetAssignmentSourceOk

`func (o *CampaignItemReview) GetAssignmentSourceOk() (*CampaignItemReviewerAssignmentSourceEnum, bool)`

GetAssignmentSourceOk returns a tuple with the AssignmentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSource

`func (o *CampaignItemReview) SetAssignmentSource(v CampaignItemReviewerAssignmentSourceEnum)`

SetAssignmentSource sets AssignmentSource field to given value.


### GetDecision

`func (o *CampaignItemReview) GetDecision() CampaignItemReviewDecisionEnum`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *CampaignItemReview) GetDecisionOk() (*CampaignItemReviewDecisionEnum, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *CampaignItemReview) SetDecision(v CampaignItemReviewDecisionEnum)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *CampaignItemReview) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetNote

`func (o *CampaignItemReview) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CampaignItemReview) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CampaignItemReview) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CampaignItemReview) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetUpdatedAccessLevelRemoteId

`func (o *CampaignItemReview) GetUpdatedAccessLevelRemoteId() string`

GetUpdatedAccessLevelRemoteId returns the UpdatedAccessLevelRemoteId field if non-nil, zero value otherwise.

### GetUpdatedAccessLevelRemoteIdOk

`func (o *CampaignItemReview) GetUpdatedAccessLevelRemoteIdOk() (*string, bool)`

GetUpdatedAccessLevelRemoteIdOk returns a tuple with the UpdatedAccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAccessLevelRemoteId

`func (o *CampaignItemReview) SetUpdatedAccessLevelRemoteId(v string)`

SetUpdatedAccessLevelRemoteId sets UpdatedAccessLevelRemoteId field to given value.

### HasUpdatedAccessLevelRemoteId

`func (o *CampaignItemReview) HasUpdatedAccessLevelRemoteId() bool`

HasUpdatedAccessLevelRemoteId returns a boolean if a field has been set.

### GetDecidedAt

`func (o *CampaignItemReview) GetDecidedAt() time.Time`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *CampaignItemReview) GetDecidedAtOk() (*time.Time, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *CampaignItemReview) SetDecidedAt(v time.Time)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *CampaignItemReview) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


