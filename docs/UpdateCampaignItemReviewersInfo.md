# UpdateCampaignItemReviewersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToAssign** | Pointer to [**[]CampaignItemReviewerAssignment**](CampaignItemReviewerAssignment.md) | Pairs to assign. Existing pairs are ignored. | [optional] 
**ToUnassign** | Pointer to [**[]CampaignItemReviewerAssignment**](CampaignItemReviewerAssignment.md) | Pairs to unassign. Only pending &#x60;MANUAL&#x60; pairs are removed; missing pairs are skipped. &#x60;reviewer_id&#x60; is required on each entry.  | [optional] 
**ToRemoveSources** | Pointer to [**[]CampaignItemReviewerSourceRemoval**](CampaignItemReviewerSourceRemoval.md) | Remove all pending rows for the given assignment sources. | [optional] 

## Methods

### NewUpdateCampaignItemReviewersInfo

`func NewUpdateCampaignItemReviewersInfo() *UpdateCampaignItemReviewersInfo`

NewUpdateCampaignItemReviewersInfo instantiates a new UpdateCampaignItemReviewersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignItemReviewersInfoWithDefaults

`func NewUpdateCampaignItemReviewersInfoWithDefaults() *UpdateCampaignItemReviewersInfo`

NewUpdateCampaignItemReviewersInfoWithDefaults instantiates a new UpdateCampaignItemReviewersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToAssign

`func (o *UpdateCampaignItemReviewersInfo) GetToAssign() []CampaignItemReviewerAssignment`

GetToAssign returns the ToAssign field if non-nil, zero value otherwise.

### GetToAssignOk

`func (o *UpdateCampaignItemReviewersInfo) GetToAssignOk() (*[]CampaignItemReviewerAssignment, bool)`

GetToAssignOk returns a tuple with the ToAssign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAssign

`func (o *UpdateCampaignItemReviewersInfo) SetToAssign(v []CampaignItemReviewerAssignment)`

SetToAssign sets ToAssign field to given value.

### HasToAssign

`func (o *UpdateCampaignItemReviewersInfo) HasToAssign() bool`

HasToAssign returns a boolean if a field has been set.

### GetToUnassign

`func (o *UpdateCampaignItemReviewersInfo) GetToUnassign() []CampaignItemReviewerAssignment`

GetToUnassign returns the ToUnassign field if non-nil, zero value otherwise.

### GetToUnassignOk

`func (o *UpdateCampaignItemReviewersInfo) GetToUnassignOk() (*[]CampaignItemReviewerAssignment, bool)`

GetToUnassignOk returns a tuple with the ToUnassign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUnassign

`func (o *UpdateCampaignItemReviewersInfo) SetToUnassign(v []CampaignItemReviewerAssignment)`

SetToUnassign sets ToUnassign field to given value.

### HasToUnassign

`func (o *UpdateCampaignItemReviewersInfo) HasToUnassign() bool`

HasToUnassign returns a boolean if a field has been set.

### GetToRemoveSources

`func (o *UpdateCampaignItemReviewersInfo) GetToRemoveSources() []CampaignItemReviewerSourceRemoval`

GetToRemoveSources returns the ToRemoveSources field if non-nil, zero value otherwise.

### GetToRemoveSourcesOk

`func (o *UpdateCampaignItemReviewersInfo) GetToRemoveSourcesOk() (*[]CampaignItemReviewerSourceRemoval, bool)`

GetToRemoveSourcesOk returns a tuple with the ToRemoveSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRemoveSources

`func (o *UpdateCampaignItemReviewersInfo) SetToRemoveSources(v []CampaignItemReviewerSourceRemoval)`

SetToRemoveSources sets ToRemoveSources field to given value.

### HasToRemoveSources

`func (o *UpdateCampaignItemReviewersInfo) HasToRemoveSources() bool`

HasToRemoveSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


