# UpdateIpZoneInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new name of the IP zone. | [optional] 
**Description** | Pointer to **string** | The new description of the IP zone. | [optional] 
**IpAddresses** | Pointer to **[]string** | The IP addresses or CIDR ranges to allow, replacing the current list. At least one is required. | [optional] 

## Methods

### NewUpdateIpZoneInfo

`func NewUpdateIpZoneInfo() *UpdateIpZoneInfo`

NewUpdateIpZoneInfo instantiates a new UpdateIpZoneInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIpZoneInfoWithDefaults

`func NewUpdateIpZoneInfoWithDefaults() *UpdateIpZoneInfo`

NewUpdateIpZoneInfoWithDefaults instantiates a new UpdateIpZoneInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateIpZoneInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIpZoneInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIpZoneInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIpZoneInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIpZoneInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIpZoneInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIpZoneInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIpZoneInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddresses

`func (o *UpdateIpZoneInfo) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *UpdateIpZoneInfo) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *UpdateIpZoneInfo) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *UpdateIpZoneInfo) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


