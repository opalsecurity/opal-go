# ReviewCampaignItemsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decisions** | [**[]ReviewCampaignItemDecisionInfo**](ReviewCampaignItemDecisionInfo.md) | The decisions to record. Each targets a single campaign item review. | 

## Methods

### NewReviewCampaignItemsInfo

`func NewReviewCampaignItemsInfo(decisions []ReviewCampaignItemDecisionInfo, ) *ReviewCampaignItemsInfo`

NewReviewCampaignItemsInfo instantiates a new ReviewCampaignItemsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewCampaignItemsInfoWithDefaults

`func NewReviewCampaignItemsInfoWithDefaults() *ReviewCampaignItemsInfo`

NewReviewCampaignItemsInfoWithDefaults instantiates a new ReviewCampaignItemsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecisions

`func (o *ReviewCampaignItemsInfo) GetDecisions() []ReviewCampaignItemDecisionInfo`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *ReviewCampaignItemsInfo) GetDecisionsOk() (*[]ReviewCampaignItemDecisionInfo, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *ReviewCampaignItemsInfo) SetDecisions(v []ReviewCampaignItemDecisionInfo)`

SetDecisions sets Decisions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


