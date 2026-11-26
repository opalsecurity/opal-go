# IpZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpZoneId** | **string** | The ID of the IP zone. | 
**Name** | **string** | The name of the IP zone. | 
**Description** | Pointer to **string** | A description of the IP zone. | [optional] 
**IpAddresses** | **[]string** | The IP addresses or CIDR ranges in the IP zone. | 
**CreatedByUserId** | **string** | The ID of the user that created the IP zone. | 
**CreatedAt** | **time.Time** | The date the IP zone was created. | 
**UpdatedAt** | **time.Time** | The date the IP zone was last updated. | 

## Methods

### NewIpZone

`func NewIpZone(ipZoneId string, name string, ipAddresses []string, createdByUserId string, createdAt time.Time, updatedAt time.Time, ) *IpZone`

NewIpZone instantiates a new IpZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpZoneWithDefaults

`func NewIpZoneWithDefaults() *IpZone`

NewIpZoneWithDefaults instantiates a new IpZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpZoneId

`func (o *IpZone) GetIpZoneId() string`

GetIpZoneId returns the IpZoneId field if non-nil, zero value otherwise.

### GetIpZoneIdOk

`func (o *IpZone) GetIpZoneIdOk() (*string, bool)`

GetIpZoneIdOk returns a tuple with the IpZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpZoneId

`func (o *IpZone) SetIpZoneId(v string)`

SetIpZoneId sets IpZoneId field to given value.


### GetName

`func (o *IpZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpZone) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IpZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddresses

`func (o *IpZone) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *IpZone) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *IpZone) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.


### GetCreatedByUserId

`func (o *IpZone) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *IpZone) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *IpZone) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.


### GetCreatedAt

`func (o *IpZone) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpZone) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpZone) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IpZone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpZone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpZone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


