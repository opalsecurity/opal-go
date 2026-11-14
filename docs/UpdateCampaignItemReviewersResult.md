# UpdateCampaignItemReviewersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviews** | [**[]CampaignItemReview**](CampaignItemReview.md) | Full review set for every referenced item after the update. | 
**UnassignedPresetMessages** | [**[]CampaignItemUnassignedPresetMessage**](CampaignItemUnassignedPresetMessage.md) | Preset assignment messages for items that did not produce a review row. | 

## Methods

### NewUpdateCampaignItemReviewersResult

`func NewUpdateCampaignItemReviewersResult(reviews []CampaignItemReview, unassignedPresetMessages []CampaignItemUnassignedPresetMessage, ) *UpdateCampaignItemReviewersResult`

NewUpdateCampaignItemReviewersResult instantiates a new UpdateCampaignItemReviewersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignItemReviewersResultWithDefaults

`func NewUpdateCampaignItemReviewersResultWithDefaults() *UpdateCampaignItemReviewersResult`

NewUpdateCampaignItemReviewersResultWithDefaults instantiates a new UpdateCampaignItemReviewersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviews

`func (o *UpdateCampaignItemReviewersResult) GetReviews() []CampaignItemReview`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *UpdateCampaignItemReviewersResult) GetReviewsOk() (*[]CampaignItemReview, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *UpdateCampaignItemReviewersResult) SetReviews(v []CampaignItemReview)`

SetReviews sets Reviews field to given value.


### GetUnassignedPresetMessages

`func (o *UpdateCampaignItemReviewersResult) GetUnassignedPresetMessages() []CampaignItemUnassignedPresetMessage`

GetUnassignedPresetMessages returns the UnassignedPresetMessages field if non-nil, zero value otherwise.

### GetUnassignedPresetMessagesOk

`func (o *UpdateCampaignItemReviewersResult) GetUnassignedPresetMessagesOk() (*[]CampaignItemUnassignedPresetMessage, bool)`

GetUnassignedPresetMessagesOk returns a tuple with the UnassignedPresetMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedPresetMessages

`func (o *UpdateCampaignItemReviewersResult) SetUnassignedPresetMessages(v []CampaignItemUnassignedPresetMessage)`

SetUnassignedPresetMessages sets UnassignedPresetMessages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


