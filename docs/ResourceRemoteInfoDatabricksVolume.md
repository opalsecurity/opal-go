# ResourceRemoteInfoDatabricksVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetastoreId** | **string** | The ID of the Unity Catalog metastore the volume belongs to. | 
**FullName** | **string** | The dot-qualified full name of the volume (e.g. \&quot;catalog.schema.volume\&quot;). | 

## Methods

### NewResourceRemoteInfoDatabricksVolume

`func NewResourceRemoteInfoDatabricksVolume(metastoreId string, fullName string, ) *ResourceRemoteInfoDatabricksVolume`

NewResourceRemoteInfoDatabricksVolume instantiates a new ResourceRemoteInfoDatabricksVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoDatabricksVolumeWithDefaults

`func NewResourceRemoteInfoDatabricksVolumeWithDefaults() *ResourceRemoteInfoDatabricksVolume`

NewResourceRemoteInfoDatabricksVolumeWithDefaults instantiates a new ResourceRemoteInfoDatabricksVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetastoreId

`func (o *ResourceRemoteInfoDatabricksVolume) GetMetastoreId() string`

GetMetastoreId returns the MetastoreId field if non-nil, zero value otherwise.

### GetMetastoreIdOk

`func (o *ResourceRemoteInfoDatabricksVolume) GetMetastoreIdOk() (*string, bool)`

GetMetastoreIdOk returns a tuple with the MetastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetastoreId

`func (o *ResourceRemoteInfoDatabricksVolume) SetMetastoreId(v string)`

SetMetastoreId sets MetastoreId field to given value.


### GetFullName

`func (o *ResourceRemoteInfoDatabricksVolume) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ResourceRemoteInfoDatabricksVolume) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ResourceRemoteInfoDatabricksVolume) SetFullName(v string)`

SetFullName sets FullName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


