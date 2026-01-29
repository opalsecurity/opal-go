# ResourceRemoteInfoAwsRdsCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | The clusterId of the RDS cluster. | 
**Region** | **string** | The region of the RDS cluster. | 
**ResourceId** | **string** | The resourceId of the RDS cluster. | 
**AccountId** | **string** | The id of the AWS account. Required for AWS Organizations. | 
**DatabaseName** | **string** | The name of the database in the RDS cluster. This can be the value of the tag &#x60;opal:database-name&#x60; or the database name. | 
**Engine** | [**RDSEngineEnum**](RDSEngineEnum.md) |  | 

## Methods

### NewResourceRemoteInfoAwsRdsCluster

`func NewResourceRemoteInfoAwsRdsCluster(clusterId string, region string, resourceId string, accountId string, databaseName string, engine RDSEngineEnum, ) *ResourceRemoteInfoAwsRdsCluster`

NewResourceRemoteInfoAwsRdsCluster instantiates a new ResourceRemoteInfoAwsRdsCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoAwsRdsClusterWithDefaults

`func NewResourceRemoteInfoAwsRdsClusterWithDefaults() *ResourceRemoteInfoAwsRdsCluster`

NewResourceRemoteInfoAwsRdsClusterWithDefaults instantiates a new ResourceRemoteInfoAwsRdsCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ResourceRemoteInfoAwsRdsCluster) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ResourceRemoteInfoAwsRdsCluster) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ResourceRemoteInfoAwsRdsCluster) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetRegion

`func (o *ResourceRemoteInfoAwsRdsCluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ResourceRemoteInfoAwsRdsCluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ResourceRemoteInfoAwsRdsCluster) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetResourceId

`func (o *ResourceRemoteInfoAwsRdsCluster) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceRemoteInfoAwsRdsCluster) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceRemoteInfoAwsRdsCluster) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetAccountId

`func (o *ResourceRemoteInfoAwsRdsCluster) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResourceRemoteInfoAwsRdsCluster) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResourceRemoteInfoAwsRdsCluster) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetDatabaseName

`func (o *ResourceRemoteInfoAwsRdsCluster) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ResourceRemoteInfoAwsRdsCluster) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ResourceRemoteInfoAwsRdsCluster) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetEngine

`func (o *ResourceRemoteInfoAwsRdsCluster) GetEngine() RDSEngineEnum`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ResourceRemoteInfoAwsRdsCluster) GetEngineOk() (*RDSEngineEnum, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ResourceRemoteInfoAwsRdsCluster) SetEngine(v RDSEngineEnum)`

SetEngine sets Engine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


