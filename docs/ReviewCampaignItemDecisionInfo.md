# ReviewCampaignItemDecisionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignItemReviewId** | **string** | The ID of the campaign item review to update. | 
**Decision** | [**CampaignItemReviewDecisionEnum**](CampaignItemReviewDecisionEnum.md) |  | 
**Note** | Pointer to **string** | Optional note from the reviewer for this decision. | [optional] 
**AccessLevelRemoteId** | Pointer to **string** | Required when &#x60;decision&#x60; is &#x60;CHANGE_ROLE&#x60;. The remote ID of the access level to change to. Ignored for other decisions.  | [optional] 
**NewReviewerUserIds** | Pointer to **[]string** | When &#x60;decision&#x60; is &#x60;REASSIGNED&#x60; and &#x60;assignment_source&#x60; is &#x60;MANUAL&#x60; or omitted, the user IDs to reassign this review to.  | [optional] 
**AssignmentSource** | Pointer to [**CampaignItemReviewerAssignmentSourceEnum**](CampaignItemReviewerAssignmentSourceEnum.md) |  | [optional] 

## Methods

### NewReviewCampaignItemDecisionInfo

`func NewReviewCampaignItemDecisionInfo(campaignItemReviewId string, decision CampaignItemReviewDecisionEnum, ) *ReviewCampaignItemDecisionInfo`

NewReviewCampaignItemDecisionInfo instantiates a new ReviewCampaignItemDecisionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewCampaignItemDecisionInfoWithDefaults

`func NewReviewCampaignItemDecisionInfoWithDefaults() *ReviewCampaignItemDecisionInfo`

NewReviewCampaignItemDecisionInfoWithDefaults instantiates a new ReviewCampaignItemDecisionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignItemReviewId

`func (o *ReviewCampaignItemDecisionInfo) GetCampaignItemReviewId() string`

GetCampaignItemReviewId returns the CampaignItemReviewId field if non-nil, zero value otherwise.

### GetCampaignItemReviewIdOk

`func (o *ReviewCampaignItemDecisionInfo) GetCampaignItemReviewIdOk() (*string, bool)`

GetCampaignItemReviewIdOk returns a tuple with the CampaignItemReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignItemReviewId

`func (o *ReviewCampaignItemDecisionInfo) SetCampaignItemReviewId(v string)`

SetCampaignItemReviewId sets CampaignItemReviewId field to given value.


### GetDecision

`func (o *ReviewCampaignItemDecisionInfo) GetDecision() CampaignItemReviewDecisionEnum`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ReviewCampaignItemDecisionInfo) GetDecisionOk() (*CampaignItemReviewDecisionEnum, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ReviewCampaignItemDecisionInfo) SetDecision(v CampaignItemReviewDecisionEnum)`

SetDecision sets Decision field to given value.


### GetNote

`func (o *ReviewCampaignItemDecisionInfo) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ReviewCampaignItemDecisionInfo) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ReviewCampaignItemDecisionInfo) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ReviewCampaignItemDecisionInfo) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAccessLevelRemoteId

`func (o *ReviewCampaignItemDecisionInfo) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *ReviewCampaignItemDecisionInfo) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *ReviewCampaignItemDecisionInfo) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *ReviewCampaignItemDecisionInfo) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetNewReviewerUserIds

`func (o *ReviewCampaignItemDecisionInfo) GetNewReviewerUserIds() []string`

GetNewReviewerUserIds returns the NewReviewerUserIds field if non-nil, zero value otherwise.

### GetNewReviewerUserIdsOk

`func (o *ReviewCampaignItemDecisionInfo) GetNewReviewerUserIdsOk() (*[]string, bool)`

GetNewReviewerUserIdsOk returns a tuple with the NewReviewerUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewReviewerUserIds

`func (o *ReviewCampaignItemDecisionInfo) SetNewReviewerUserIds(v []string)`

SetNewReviewerUserIds sets NewReviewerUserIds field to given value.

### HasNewReviewerUserIds

`func (o *ReviewCampaignItemDecisionInfo) HasNewReviewerUserIds() bool`

HasNewReviewerUserIds returns a boolean if a field has been set.

### GetAssignmentSource

`func (o *ReviewCampaignItemDecisionInfo) GetAssignmentSource() CampaignItemReviewerAssignmentSourceEnum`

GetAssignmentSource returns the AssignmentSource field if non-nil, zero value otherwise.

### GetAssignmentSourceOk

`func (o *ReviewCampaignItemDecisionInfo) GetAssignmentSourceOk() (*CampaignItemReviewerAssignmentSourceEnum, bool)`

GetAssignmentSourceOk returns a tuple with the AssignmentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSource

`func (o *ReviewCampaignItemDecisionInfo) SetAssignmentSource(v CampaignItemReviewerAssignmentSourceEnum)`

SetAssignmentSource sets AssignmentSource field to given value.

### HasAssignmentSource

`func (o *ReviewCampaignItemDecisionInfo) HasAssignmentSource() bool`

HasAssignmentSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


