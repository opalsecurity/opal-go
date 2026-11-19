# CampaignItemUnassignedPresetMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignItemId** | **string** | The ID of the campaign item. | 
**AssignmentSource** | [**CampaignItemReviewerAssignmentSourceEnum**](CampaignItemReviewerAssignmentSourceEnum.md) |  | 
**Message** | **string** | Human-readable explanation of why no reviewer was assigned. | 

## Methods

### NewCampaignItemUnassignedPresetMessage

`func NewCampaignItemUnassignedPresetMessage(campaignItemId string, assignmentSource CampaignItemReviewerAssignmentSourceEnum, message string, ) *CampaignItemUnassignedPresetMessage`

NewCampaignItemUnassignedPresetMessage instantiates a new CampaignItemUnassignedPresetMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignItemUnassignedPresetMessageWithDefaults

`func NewCampaignItemUnassignedPresetMessageWithDefaults() *CampaignItemUnassignedPresetMessage`

NewCampaignItemUnassignedPresetMessageWithDefaults instantiates a new CampaignItemUnassignedPresetMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignItemId

`func (o *CampaignItemUnassignedPresetMessage) GetCampaignItemId() string`

GetCampaignItemId returns the CampaignItemId field if non-nil, zero value otherwise.

### GetCampaignItemIdOk

`func (o *CampaignItemUnassignedPresetMessage) GetCampaignItemIdOk() (*string, bool)`

GetCampaignItemIdOk returns a tuple with the CampaignItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignItemId

`func (o *CampaignItemUnassignedPresetMessage) SetCampaignItemId(v string)`

SetCampaignItemId sets CampaignItemId field to given value.


### GetAssignmentSource

`func (o *CampaignItemUnassignedPresetMessage) GetAssignmentSource() CampaignItemReviewerAssignmentSourceEnum`

GetAssignmentSource returns the AssignmentSource field if non-nil, zero value otherwise.

### GetAssignmentSourceOk

`func (o *CampaignItemUnassignedPresetMessage) GetAssignmentSourceOk() (*CampaignItemReviewerAssignmentSourceEnum, bool)`

GetAssignmentSourceOk returns a tuple with the AssignmentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSource

`func (o *CampaignItemUnassignedPresetMessage) SetAssignmentSource(v CampaignItemReviewerAssignmentSourceEnum)`

SetAssignmentSource sets AssignmentSource field to given value.


### GetMessage

`func (o *CampaignItemUnassignedPresetMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CampaignItemUnassignedPresetMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CampaignItemUnassignedPresetMessage) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


