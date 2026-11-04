# ResourceRemoteInfoDatabricksCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetastoreId** | **string** | The ID of the Unity Catalog metastore the catalog belongs to. | 
**FullName** | **string** | The dot-qualified full name of the catalog (e.g. \&quot;catalog\&quot;). | 

## Methods

### NewResourceRemoteInfoDatabricksCatalog

`func NewResourceRemoteInfoDatabricksCatalog(metastoreId string, fullName string, ) *ResourceRemoteInfoDatabricksCatalog`

NewResourceRemoteInfoDatabricksCatalog instantiates a new ResourceRemoteInfoDatabricksCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoDatabricksCatalogWithDefaults

`func NewResourceRemoteInfoDatabricksCatalogWithDefaults() *ResourceRemoteInfoDatabricksCatalog`

NewResourceRemoteInfoDatabricksCatalogWithDefaults instantiates a new ResourceRemoteInfoDatabricksCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetastoreId

`func (o *ResourceRemoteInfoDatabricksCatalog) GetMetastoreId() string`

GetMetastoreId returns the MetastoreId field if non-nil, zero value otherwise.

### GetMetastoreIdOk

`func (o *ResourceRemoteInfoDatabricksCatalog) GetMetastoreIdOk() (*string, bool)`

GetMetastoreIdOk returns a tuple with the MetastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetastoreId

`func (o *ResourceRemoteInfoDatabricksCatalog) SetMetastoreId(v string)`

SetMetastoreId sets MetastoreId field to given value.


### GetFullName

`func (o *ResourceRemoteInfoDatabricksCatalog) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ResourceRemoteInfoDatabricksCatalog) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ResourceRemoteInfoDatabricksCatalog) SetFullName(v string)`

SetFullName sets FullName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


