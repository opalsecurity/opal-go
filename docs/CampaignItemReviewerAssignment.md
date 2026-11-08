# CampaignItemReviewerAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignItemId** | **string** | The ID of the campaign item. | 
**ReviewerId** | Pointer to **string** | The ID of the reviewer. Required when &#x60;assignment_source&#x60; is &#x60;MANUAL&#x60; (or omitted), and required for every unassign entry.  | [optional] 
**AssignmentSource** | Pointer to [**CampaignItemReviewerAssignmentSourceEnum**](CampaignItemReviewerAssignmentSourceEnum.md) | Defaults to &#x60;MANUAL&#x60;. When &#x60;PRINCIPAL_MANAGER&#x60; or &#x60;ASSET_ADMIN_OWNER&#x60;, the backend resolves reviewer ID(s) and &#x60;reviewer_id&#x60; may be omitted.  | [optional] 

## Methods

### NewCampaignItemReviewerAssignment

`func NewCampaignItemReviewerAssignment(campaignItemId string, ) *CampaignItemReviewerAssignment`

NewCampaignItemReviewerAssignment instantiates a new CampaignItemReviewerAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignItemReviewerAssignmentWithDefaults

`func NewCampaignItemReviewerAssignmentWithDefaults() *CampaignItemReviewerAssignment`

NewCampaignItemReviewerAssignmentWithDefaults instantiates a new CampaignItemReviewerAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignItemId

`func (o *CampaignItemReviewerAssignment) GetCampaignItemId() string`

GetCampaignItemId returns the CampaignItemId field if non-nil, zero value otherwise.

### GetCampaignItemIdOk

`func (o *CampaignItemReviewerAssignment) GetCampaignItemIdOk() (*string, bool)`

GetCampaignItemIdOk returns a tuple with the CampaignItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignItemId

`func (o *CampaignItemReviewerAssignment) SetCampaignItemId(v string)`

SetCampaignItemId sets CampaignItemId field to given value.


### GetReviewerId

`func (o *CampaignItemReviewerAssignment) GetReviewerId() string`

GetReviewerId returns the ReviewerId field if non-nil, zero value otherwise.

### GetReviewerIdOk

`func (o *CampaignItemReviewerAssignment) GetReviewerIdOk() (*string, bool)`

GetReviewerIdOk returns a tuple with the ReviewerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerId

`func (o *CampaignItemReviewerAssignment) SetReviewerId(v string)`

SetReviewerId sets ReviewerId field to given value.

### HasReviewerId

`func (o *CampaignItemReviewerAssignment) HasReviewerId() bool`

HasReviewerId returns a boolean if a field has been set.

### GetAssignmentSource

`func (o *CampaignItemReviewerAssignment) GetAssignmentSource() CampaignItemReviewerAssignmentSourceEnum`

GetAssignmentSource returns the AssignmentSource field if non-nil, zero value otherwise.

### GetAssignmentSourceOk

`func (o *CampaignItemReviewerAssignment) GetAssignmentSourceOk() (*CampaignItemReviewerAssignmentSourceEnum, bool)`

GetAssignmentSourceOk returns a tuple with the AssignmentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSource

`func (o *CampaignItemReviewerAssignment) SetAssignmentSource(v CampaignItemReviewerAssignmentSourceEnum)`

SetAssignmentSource sets AssignmentSource field to given value.

### HasAssignmentSource

`func (o *CampaignItemReviewerAssignment) HasAssignmentSource() bool`

HasAssignmentSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


