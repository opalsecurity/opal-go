# CampaignItemReviewResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignItemReviewId** | **string** | The ID of the campaign item review. | 
**Decision** | Pointer to [**CampaignItemReviewDecisionEnum**](CampaignItemReviewDecisionEnum.md) |  | [optional] 
**Note** | Pointer to **string** | Note recorded with the decision, if any. | [optional] 
**DecidedAt** | Pointer to **time.Time** | When the decision was submitted. | [optional] 
**UpdatedAccessLevelRemoteId** | Pointer to **string** | When &#x60;decision&#x60; is &#x60;CHANGE_ROLE&#x60;, the remote ID of the access level the reviewer selected.  | [optional] 
**ReassignedToReviewerIds** | Pointer to **[]string** | When &#x60;decision&#x60; is &#x60;REASSIGNED&#x60;, the user IDs the review was reassigned to.  | [optional] 

## Methods

### NewCampaignItemReviewResult

`func NewCampaignItemReviewResult(campaignItemReviewId string, ) *CampaignItemReviewResult`

NewCampaignItemReviewResult instantiates a new CampaignItemReviewResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignItemReviewResultWithDefaults

`func NewCampaignItemReviewResultWithDefaults() *CampaignItemReviewResult`

NewCampaignItemReviewResultWithDefaults instantiates a new CampaignItemReviewResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignItemReviewId

`func (o *CampaignItemReviewResult) GetCampaignItemReviewId() string`

GetCampaignItemReviewId returns the CampaignItemReviewId field if non-nil, zero value otherwise.

### GetCampaignItemReviewIdOk

`func (o *CampaignItemReviewResult) GetCampaignItemReviewIdOk() (*string, bool)`

GetCampaignItemReviewIdOk returns a tuple with the CampaignItemReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignItemReviewId

`func (o *CampaignItemReviewResult) SetCampaignItemReviewId(v string)`

SetCampaignItemReviewId sets CampaignItemReviewId field to given value.


### GetDecision

`func (o *CampaignItemReviewResult) GetDecision() CampaignItemReviewDecisionEnum`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *CampaignItemReviewResult) GetDecisionOk() (*CampaignItemReviewDecisionEnum, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *CampaignItemReviewResult) SetDecision(v CampaignItemReviewDecisionEnum)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *CampaignItemReviewResult) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetNote

`func (o *CampaignItemReviewResult) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CampaignItemReviewResult) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CampaignItemReviewResult) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CampaignItemReviewResult) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetDecidedAt

`func (o *CampaignItemReviewResult) GetDecidedAt() time.Time`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *CampaignItemReviewResult) GetDecidedAtOk() (*time.Time, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *CampaignItemReviewResult) SetDecidedAt(v time.Time)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *CampaignItemReviewResult) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetUpdatedAccessLevelRemoteId

`func (o *CampaignItemReviewResult) GetUpdatedAccessLevelRemoteId() string`

GetUpdatedAccessLevelRemoteId returns the UpdatedAccessLevelRemoteId field if non-nil, zero value otherwise.

### GetUpdatedAccessLevelRemoteIdOk

`func (o *CampaignItemReviewResult) GetUpdatedAccessLevelRemoteIdOk() (*string, bool)`

GetUpdatedAccessLevelRemoteIdOk returns a tuple with the UpdatedAccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAccessLevelRemoteId

`func (o *CampaignItemReviewResult) SetUpdatedAccessLevelRemoteId(v string)`

SetUpdatedAccessLevelRemoteId sets UpdatedAccessLevelRemoteId field to given value.

### HasUpdatedAccessLevelRemoteId

`func (o *CampaignItemReviewResult) HasUpdatedAccessLevelRemoteId() bool`

HasUpdatedAccessLevelRemoteId returns a boolean if a field has been set.

### GetReassignedToReviewerIds

`func (o *CampaignItemReviewResult) GetReassignedToReviewerIds() []string`

GetReassignedToReviewerIds returns the ReassignedToReviewerIds field if non-nil, zero value otherwise.

### GetReassignedToReviewerIdsOk

`func (o *CampaignItemReviewResult) GetReassignedToReviewerIdsOk() (*[]string, bool)`

GetReassignedToReviewerIdsOk returns a tuple with the ReassignedToReviewerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignedToReviewerIds

`func (o *CampaignItemReviewResult) SetReassignedToReviewerIds(v []string)`

SetReassignedToReviewerIds sets ReassignedToReviewerIds field to given value.

### HasReassignedToReviewerIds

`func (o *CampaignItemReviewResult) HasReassignedToReviewerIds() bool`

HasReassignedToReviewerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


