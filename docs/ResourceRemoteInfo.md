# ResourceRemoteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsIamRole** | Pointer to [**ResourceRemoteInfoAwsIamRole**](ResourceRemoteInfoAwsIamRole.md) |  | [optional] 
**AwsEc2Instance** | Pointer to [**ResourceRemoteInfoAwsEc2Instance**](ResourceRemoteInfoAwsEc2Instance.md) |  | [optional] 
**AwsRdsInstance** | Pointer to [**ResourceRemoteInfoAwsRdsInstance**](ResourceRemoteInfoAwsRdsInstance.md) |  | [optional] 
**AwsEksCluster** | Pointer to [**ResourceRemoteInfoAwsEksCluster**](ResourceRemoteInfoAwsEksCluster.md) |  | [optional] 
**GithubRepo** | Pointer to [**ResourceRemoteInfoGithubRepo**](ResourceRemoteInfoGithubRepo.md) |  | [optional] 
**OktaApp** | Pointer to [**ResourceRemoteInfoOktaApp**](ResourceRemoteInfoOktaApp.md) |  | [optional] 
**OktaStandardRole** | Pointer to [**ResourceRemoteInfoOktaStandardRole**](ResourceRemoteInfoOktaStandardRole.md) |  | [optional] 
**OktaCustomRole** | Pointer to [**ResourceRemoteInfoOktaCustomRole**](ResourceRemoteInfoOktaCustomRole.md) |  | [optional] 

## Methods

### NewResourceRemoteInfo

`func NewResourceRemoteInfo() *ResourceRemoteInfo`

NewResourceRemoteInfo instantiates a new ResourceRemoteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoWithDefaults

`func NewResourceRemoteInfoWithDefaults() *ResourceRemoteInfo`

NewResourceRemoteInfoWithDefaults instantiates a new ResourceRemoteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsIamRole

`func (o *ResourceRemoteInfo) GetAwsIamRole() ResourceRemoteInfoAwsIamRole`

GetAwsIamRole returns the AwsIamRole field if non-nil, zero value otherwise.

### GetAwsIamRoleOk

`func (o *ResourceRemoteInfo) GetAwsIamRoleOk() (*ResourceRemoteInfoAwsIamRole, bool)`

GetAwsIamRoleOk returns a tuple with the AwsIamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIamRole

`func (o *ResourceRemoteInfo) SetAwsIamRole(v ResourceRemoteInfoAwsIamRole)`

SetAwsIamRole sets AwsIamRole field to given value.

### HasAwsIamRole

`func (o *ResourceRemoteInfo) HasAwsIamRole() bool`

HasAwsIamRole returns a boolean if a field has been set.

### GetAwsEc2Instance

`func (o *ResourceRemoteInfo) GetAwsEc2Instance() ResourceRemoteInfoAwsEc2Instance`

GetAwsEc2Instance returns the AwsEc2Instance field if non-nil, zero value otherwise.

### GetAwsEc2InstanceOk

`func (o *ResourceRemoteInfo) GetAwsEc2InstanceOk() (*ResourceRemoteInfoAwsEc2Instance, bool)`

GetAwsEc2InstanceOk returns a tuple with the AwsEc2Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEc2Instance

`func (o *ResourceRemoteInfo) SetAwsEc2Instance(v ResourceRemoteInfoAwsEc2Instance)`

SetAwsEc2Instance sets AwsEc2Instance field to given value.

### HasAwsEc2Instance

`func (o *ResourceRemoteInfo) HasAwsEc2Instance() bool`

HasAwsEc2Instance returns a boolean if a field has been set.

### GetAwsRdsInstance

`func (o *ResourceRemoteInfo) GetAwsRdsInstance() ResourceRemoteInfoAwsRdsInstance`

GetAwsRdsInstance returns the AwsRdsInstance field if non-nil, zero value otherwise.

### GetAwsRdsInstanceOk

`func (o *ResourceRemoteInfo) GetAwsRdsInstanceOk() (*ResourceRemoteInfoAwsRdsInstance, bool)`

GetAwsRdsInstanceOk returns a tuple with the AwsRdsInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRdsInstance

`func (o *ResourceRemoteInfo) SetAwsRdsInstance(v ResourceRemoteInfoAwsRdsInstance)`

SetAwsRdsInstance sets AwsRdsInstance field to given value.

### HasAwsRdsInstance

`func (o *ResourceRemoteInfo) HasAwsRdsInstance() bool`

HasAwsRdsInstance returns a boolean if a field has been set.

### GetAwsEksCluster

`func (o *ResourceRemoteInfo) GetAwsEksCluster() ResourceRemoteInfoAwsEksCluster`

GetAwsEksCluster returns the AwsEksCluster field if non-nil, zero value otherwise.

### GetAwsEksClusterOk

`func (o *ResourceRemoteInfo) GetAwsEksClusterOk() (*ResourceRemoteInfoAwsEksCluster, bool)`

GetAwsEksClusterOk returns a tuple with the AwsEksCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEksCluster

`func (o *ResourceRemoteInfo) SetAwsEksCluster(v ResourceRemoteInfoAwsEksCluster)`

SetAwsEksCluster sets AwsEksCluster field to given value.

### HasAwsEksCluster

`func (o *ResourceRemoteInfo) HasAwsEksCluster() bool`

HasAwsEksCluster returns a boolean if a field has been set.

### GetGithubRepo

`func (o *ResourceRemoteInfo) GetGithubRepo() ResourceRemoteInfoGithubRepo`

GetGithubRepo returns the GithubRepo field if non-nil, zero value otherwise.

### GetGithubRepoOk

`func (o *ResourceRemoteInfo) GetGithubRepoOk() (*ResourceRemoteInfoGithubRepo, bool)`

GetGithubRepoOk returns a tuple with the GithubRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubRepo

`func (o *ResourceRemoteInfo) SetGithubRepo(v ResourceRemoteInfoGithubRepo)`

SetGithubRepo sets GithubRepo field to given value.

### HasGithubRepo

`func (o *ResourceRemoteInfo) HasGithubRepo() bool`

HasGithubRepo returns a boolean if a field has been set.

### GetOktaApp

`func (o *ResourceRemoteInfo) GetOktaApp() ResourceRemoteInfoOktaApp`

GetOktaApp returns the OktaApp field if non-nil, zero value otherwise.

### GetOktaAppOk

`func (o *ResourceRemoteInfo) GetOktaAppOk() (*ResourceRemoteInfoOktaApp, bool)`

GetOktaAppOk returns a tuple with the OktaApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaApp

`func (o *ResourceRemoteInfo) SetOktaApp(v ResourceRemoteInfoOktaApp)`

SetOktaApp sets OktaApp field to given value.

### HasOktaApp

`func (o *ResourceRemoteInfo) HasOktaApp() bool`

HasOktaApp returns a boolean if a field has been set.

### GetOktaStandardRole

`func (o *ResourceRemoteInfo) GetOktaStandardRole() ResourceRemoteInfoOktaStandardRole`

GetOktaStandardRole returns the OktaStandardRole field if non-nil, zero value otherwise.

### GetOktaStandardRoleOk

`func (o *ResourceRemoteInfo) GetOktaStandardRoleOk() (*ResourceRemoteInfoOktaStandardRole, bool)`

GetOktaStandardRoleOk returns a tuple with the OktaStandardRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaStandardRole

`func (o *ResourceRemoteInfo) SetOktaStandardRole(v ResourceRemoteInfoOktaStandardRole)`

SetOktaStandardRole sets OktaStandardRole field to given value.

### HasOktaStandardRole

`func (o *ResourceRemoteInfo) HasOktaStandardRole() bool`

HasOktaStandardRole returns a boolean if a field has been set.

### GetOktaCustomRole

`func (o *ResourceRemoteInfo) GetOktaCustomRole() ResourceRemoteInfoOktaCustomRole`

GetOktaCustomRole returns the OktaCustomRole field if non-nil, zero value otherwise.

### GetOktaCustomRoleOk

`func (o *ResourceRemoteInfo) GetOktaCustomRoleOk() (*ResourceRemoteInfoOktaCustomRole, bool)`

GetOktaCustomRoleOk returns a tuple with the OktaCustomRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaCustomRole

`func (o *ResourceRemoteInfo) SetOktaCustomRole(v ResourceRemoteInfoOktaCustomRole)`

SetOktaCustomRole sets OktaCustomRole field to given value.

### HasOktaCustomRole

`func (o *ResourceRemoteInfo) HasOktaCustomRole() bool`

HasOktaCustomRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


