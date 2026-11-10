# ResourceRemoteInfoDatabricksSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetastoreId** | **string** | The ID of the Unity Catalog metastore the schema belongs to. | 
**FullName** | **string** | The dot-qualified full name of the schema (e.g. \&quot;catalog.schema\&quot;). | 

## Methods

### NewResourceRemoteInfoDatabricksSchema

`func NewResourceRemoteInfoDatabricksSchema(metastoreId string, fullName string, ) *ResourceRemoteInfoDatabricksSchema`

NewResourceRemoteInfoDatabricksSchema instantiates a new ResourceRemoteInfoDatabricksSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoDatabricksSchemaWithDefaults

`func NewResourceRemoteInfoDatabricksSchemaWithDefaults() *ResourceRemoteInfoDatabricksSchema`

NewResourceRemoteInfoDatabricksSchemaWithDefaults instantiates a new ResourceRemoteInfoDatabricksSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetastoreId

`func (o *ResourceRemoteInfoDatabricksSchema) GetMetastoreId() string`

GetMetastoreId returns the MetastoreId field if non-nil, zero value otherwise.

### GetMetastoreIdOk

`func (o *ResourceRemoteInfoDatabricksSchema) GetMetastoreIdOk() (*string, bool)`

GetMetastoreIdOk returns a tuple with the MetastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetastoreId

`func (o *ResourceRemoteInfoDatabricksSchema) SetMetastoreId(v string)`

SetMetastoreId sets MetastoreId field to given value.


### GetFullName

`func (o *ResourceRemoteInfoDatabricksSchema) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ResourceRemoteInfoDatabricksSchema) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ResourceRemoteInfoDatabricksSchema) SetFullName(v string)`

SetFullName sets FullName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


