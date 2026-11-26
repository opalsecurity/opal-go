# StopCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevokeUnreviewed** | Pointer to **bool** | Revoke all unreviewed access grants. Access grants with no reviewer decision will be immediately revoked. | [optional] [default to false]

## Methods

### NewStopCampaignRequest

`func NewStopCampaignRequest() *StopCampaignRequest`

NewStopCampaignRequest instantiates a new StopCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopCampaignRequestWithDefaults

`func NewStopCampaignRequestWithDefaults() *StopCampaignRequest`

NewStopCampaignRequestWithDefaults instantiates a new StopCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevokeUnreviewed

`func (o *StopCampaignRequest) GetRevokeUnreviewed() bool`

GetRevokeUnreviewed returns the RevokeUnreviewed field if non-nil, zero value otherwise.

### GetRevokeUnreviewedOk

`func (o *StopCampaignRequest) GetRevokeUnreviewedOk() (*bool, bool)`

GetRevokeUnreviewedOk returns a tuple with the RevokeUnreviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeUnreviewed

`func (o *StopCampaignRequest) SetRevokeUnreviewed(v bool)`

SetRevokeUnreviewed sets RevokeUnreviewed field to given value.

### HasRevokeUnreviewed

`func (o *StopCampaignRequest) HasRevokeUnreviewed() bool`

HasRevokeUnreviewed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


