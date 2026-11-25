# IpAclSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether API token requests are restricted to the IP addresses in the organization&#39;s IP zones. | 

## Methods

### NewIpAclSettings

`func NewIpAclSettings(enabled bool, ) *IpAclSettings`

NewIpAclSettings instantiates a new IpAclSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAclSettingsWithDefaults

`func NewIpAclSettingsWithDefaults() *IpAclSettings`

NewIpAclSettingsWithDefaults instantiates a new IpAclSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *IpAclSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IpAclSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IpAclSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


