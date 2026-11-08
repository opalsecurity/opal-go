# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **string** | The ID of the campaign. | 
**Name** | **string** | The name of the campaign. | 
**Status** | [**CampaignStatusEnum**](CampaignStatusEnum.md) |  | 
**IsTemplate** | **bool** | Whether this campaign is a recurring schedule template. Templates spawn draft campaigns on schedule rather than being reviewed directly. | 
**CreatedAt** | **time.Time** | The creation time of the campaign. | 
**UpdatedAt** | **time.Time** | The last updated time of the campaign. | 
**CreatedByUserId** | **string** | The ID of the user who created the campaign. | 
**Configuration** | Pointer to [**CampaignConfiguration**](CampaignConfiguration.md) | The campaign&#39;s configuration, if set. | [optional] 
**StartedAt** | Pointer to **time.Time** | The time the campaign was started, if started. | [optional] 
**StartedByUserId** | Pointer to **string** | The ID of the user who started the campaign, if started. | [optional] 
**StoppedAt** | Pointer to **time.Time** | The time the campaign was manually stopped, if stopped. | [optional] 
**StoppedByUserId** | Pointer to **string** | The ID of the user who stopped the campaign, if stopped. | [optional] 
**EndedAt** | Pointer to **time.Time** | The time the campaign reached its scheduled end, if ended. | [optional] 
**EndedByUserId** | Pointer to **string** | The ID of the user who ended the campaign, if ended. | [optional] 

## Methods

### NewCampaign

`func NewCampaign(campaignId string, name string, status CampaignStatusEnum, isTemplate bool, createdAt time.Time, updatedAt time.Time, createdByUserId string, ) *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *Campaign) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Campaign) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *Campaign) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Campaign) GetStatus() CampaignStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Campaign) GetStatusOk() (*CampaignStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Campaign) SetStatus(v CampaignStatusEnum)`

SetStatus sets Status field to given value.


### GetIsTemplate

`func (o *Campaign) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *Campaign) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *Campaign) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.


### GetCreatedAt

`func (o *Campaign) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Campaign) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Campaign) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Campaign) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Campaign) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Campaign) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedByUserId

`func (o *Campaign) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *Campaign) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *Campaign) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.


### GetConfiguration

`func (o *Campaign) GetConfiguration() CampaignConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Campaign) GetConfigurationOk() (*CampaignConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Campaign) SetConfiguration(v CampaignConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Campaign) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetStartedAt

`func (o *Campaign) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Campaign) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Campaign) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Campaign) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStartedByUserId

`func (o *Campaign) GetStartedByUserId() string`

GetStartedByUserId returns the StartedByUserId field if non-nil, zero value otherwise.

### GetStartedByUserIdOk

`func (o *Campaign) GetStartedByUserIdOk() (*string, bool)`

GetStartedByUserIdOk returns a tuple with the StartedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedByUserId

`func (o *Campaign) SetStartedByUserId(v string)`

SetStartedByUserId sets StartedByUserId field to given value.

### HasStartedByUserId

`func (o *Campaign) HasStartedByUserId() bool`

HasStartedByUserId returns a boolean if a field has been set.

### GetStoppedAt

`func (o *Campaign) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *Campaign) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *Campaign) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.

### HasStoppedAt

`func (o *Campaign) HasStoppedAt() bool`

HasStoppedAt returns a boolean if a field has been set.

### GetStoppedByUserId

`func (o *Campaign) GetStoppedByUserId() string`

GetStoppedByUserId returns the StoppedByUserId field if non-nil, zero value otherwise.

### GetStoppedByUserIdOk

`func (o *Campaign) GetStoppedByUserIdOk() (*string, bool)`

GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedByUserId

`func (o *Campaign) SetStoppedByUserId(v string)`

SetStoppedByUserId sets StoppedByUserId field to given value.

### HasStoppedByUserId

`func (o *Campaign) HasStoppedByUserId() bool`

HasStoppedByUserId returns a boolean if a field has been set.

### GetEndedAt

`func (o *Campaign) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Campaign) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Campaign) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Campaign) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEndedByUserId

`func (o *Campaign) GetEndedByUserId() string`

GetEndedByUserId returns the EndedByUserId field if non-nil, zero value otherwise.

### GetEndedByUserIdOk

`func (o *Campaign) GetEndedByUserIdOk() (*string, bool)`

GetEndedByUserIdOk returns a tuple with the EndedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedByUserId

`func (o *Campaign) SetEndedByUserId(v string)`

SetEndedByUserId sets EndedByUserId field to given value.

### HasEndedByUserId

`func (o *Campaign) HasEndedByUserId() bool`

HasEndedByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


