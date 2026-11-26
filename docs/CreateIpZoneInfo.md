# CreateIpZoneInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the IP zone. | 
**Description** | Pointer to **string** | A description of the IP zone. | [optional] 
**IpAddresses** | **[]string** | The IP addresses or CIDR ranges to allow. At least one is required. | 

## Methods

### NewCreateIpZoneInfo

`func NewCreateIpZoneInfo(name string, ipAddresses []string, ) *CreateIpZoneInfo`

NewCreateIpZoneInfo instantiates a new CreateIpZoneInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIpZoneInfoWithDefaults

`func NewCreateIpZoneInfoWithDefaults() *CreateIpZoneInfo`

NewCreateIpZoneInfoWithDefaults instantiates a new CreateIpZoneInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIpZoneInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIpZoneInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIpZoneInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateIpZoneInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIpZoneInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIpZoneInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIpZoneInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddresses

`func (o *CreateIpZoneInfo) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CreateIpZoneInfo) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CreateIpZoneInfo) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


