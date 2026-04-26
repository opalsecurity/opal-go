# ResourceRemoteInfoClickhouseTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | **string** | The name of the ClickHouse database containing the table. | 
**TableName** | **string** | The name of the ClickHouse table. | 

## Methods

### NewResourceRemoteInfoClickhouseTable

`func NewResourceRemoteInfoClickhouseTable(databaseName string, tableName string, ) *ResourceRemoteInfoClickhouseTable`

NewResourceRemoteInfoClickhouseTable instantiates a new ResourceRemoteInfoClickhouseTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoClickhouseTableWithDefaults

`func NewResourceRemoteInfoClickhouseTableWithDefaults() *ResourceRemoteInfoClickhouseTable`

NewResourceRemoteInfoClickhouseTableWithDefaults instantiates a new ResourceRemoteInfoClickhouseTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *ResourceRemoteInfoClickhouseTable) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ResourceRemoteInfoClickhouseTable) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ResourceRemoteInfoClickhouseTable) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetTableName

`func (o *ResourceRemoteInfoClickhouseTable) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *ResourceRemoteInfoClickhouseTable) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *ResourceRemoteInfoClickhouseTable) SetTableName(v string)`

SetTableName sets TableName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


