# ResourceRemoteInfoAwsEksCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** | The ARN of the EKS cluster. | 
**AccountId** | Pointer to **string** | The id of the AWS account. Required for AWS Organizations. | [optional] 

## Methods

### NewResourceRemoteInfoAwsEksCluster

`func NewResourceRemoteInfoAwsEksCluster(arn string, ) *ResourceRemoteInfoAwsEksCluster`

NewResourceRemoteInfoAwsEksCluster instantiates a new ResourceRemoteInfoAwsEksCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoAwsEksClusterWithDefaults

`func NewResourceRemoteInfoAwsEksClusterWithDefaults() *ResourceRemoteInfoAwsEksCluster`

NewResourceRemoteInfoAwsEksClusterWithDefaults instantiates a new ResourceRemoteInfoAwsEksCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ResourceRemoteInfoAwsEksCluster) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ResourceRemoteInfoAwsEksCluster) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ResourceRemoteInfoAwsEksCluster) SetArn(v string)`

SetArn sets Arn field to given value.


### GetAccountId

`func (o *ResourceRemoteInfoAwsEksCluster) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResourceRemoteInfoAwsEksCluster) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResourceRemoteInfoAwsEksCluster) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ResourceRemoteInfoAwsEksCluster) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


