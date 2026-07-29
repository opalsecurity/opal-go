# CreateCampaignInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the campaign. | 
**Configuration** | [**CreateCampaignConfigurationInfo**](CreateCampaignConfigurationInfo.md) | Configuration for the campaign. Required; use an empty object to apply defaults. | 

## Methods

### NewCreateCampaignInfo

`func NewCreateCampaignInfo(name string, configuration CreateCampaignConfigurationInfo, ) *CreateCampaignInfo`

NewCreateCampaignInfo instantiates a new CreateCampaignInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCampaignInfoWithDefaults

`func NewCreateCampaignInfoWithDefaults() *CreateCampaignInfo`

NewCreateCampaignInfoWithDefaults instantiates a new CreateCampaignInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCampaignInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCampaignInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCampaignInfo) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *CreateCampaignInfo) GetConfiguration() CreateCampaignConfigurationInfo`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateCampaignInfo) GetConfigurationOk() (*CreateCampaignConfigurationInfo, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateCampaignInfo) SetConfiguration(v CreateCampaignConfigurationInfo)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


