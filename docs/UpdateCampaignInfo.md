# UpdateCampaignInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the campaign. | [optional] 
**Configuration** | Pointer to [**UpdateCampaignConfigurationInfo**](UpdateCampaignConfigurationInfo.md) | Configuration fields to create or update. | [optional] 

## Methods

### NewUpdateCampaignInfo

`func NewUpdateCampaignInfo() *UpdateCampaignInfo`

NewUpdateCampaignInfo instantiates a new UpdateCampaignInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignInfoWithDefaults

`func NewUpdateCampaignInfoWithDefaults() *UpdateCampaignInfo`

NewUpdateCampaignInfoWithDefaults instantiates a new UpdateCampaignInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCampaignInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCampaignInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCampaignInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCampaignInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateCampaignInfo) GetConfiguration() UpdateCampaignConfigurationInfo`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateCampaignInfo) GetConfigurationOk() (*UpdateCampaignConfigurationInfo, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateCampaignInfo) SetConfiguration(v UpdateCampaignConfigurationInfo)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateCampaignInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


