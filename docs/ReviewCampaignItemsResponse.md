# ReviewCampaignItemsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviews** | [**[]CampaignItemReviewResult**](CampaignItemReviewResult.md) | Updated reviews in the same order as the request decisions. | 

## Methods

### NewReviewCampaignItemsResponse

`func NewReviewCampaignItemsResponse(reviews []CampaignItemReviewResult, ) *ReviewCampaignItemsResponse`

NewReviewCampaignItemsResponse instantiates a new ReviewCampaignItemsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewCampaignItemsResponseWithDefaults

`func NewReviewCampaignItemsResponseWithDefaults() *ReviewCampaignItemsResponse`

NewReviewCampaignItemsResponseWithDefaults instantiates a new ReviewCampaignItemsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviews

`func (o *ReviewCampaignItemsResponse) GetReviews() []CampaignItemReviewResult`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *ReviewCampaignItemsResponse) GetReviewsOk() (*[]CampaignItemReviewResult, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *ReviewCampaignItemsResponse) SetReviews(v []CampaignItemReviewResult)`

SetReviews sets Reviews field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


