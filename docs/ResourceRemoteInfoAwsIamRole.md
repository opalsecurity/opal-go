# ResourceRemoteInfoAwsIamRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** | The ARN of the IAM role. | 
**AccountId** | Pointer to **string** | The id of the AWS account. Required for AWS Organizations. | [optional] 

## Methods

### NewResourceRemoteInfoAwsIamRole

`func NewResourceRemoteInfoAwsIamRole(arn string, ) *ResourceRemoteInfoAwsIamRole`

NewResourceRemoteInfoAwsIamRole instantiates a new ResourceRemoteInfoAwsIamRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoAwsIamRoleWithDefaults

`func NewResourceRemoteInfoAwsIamRoleWithDefaults() *ResourceRemoteInfoAwsIamRole`

NewResourceRemoteInfoAwsIamRoleWithDefaults instantiates a new ResourceRemoteInfoAwsIamRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *ResourceRemoteInfoAwsIamRole) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *ResourceRemoteInfoAwsIamRole) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *ResourceRemoteInfoAwsIamRole) SetArn(v string)`

SetArn sets Arn field to given value.


### GetAccountId

`func (o *ResourceRemoteInfoAwsIamRole) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResourceRemoteInfoAwsIamRole) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResourceRemoteInfoAwsIamRole) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ResourceRemoteInfoAwsIamRole) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


