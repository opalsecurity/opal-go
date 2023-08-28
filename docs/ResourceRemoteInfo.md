# ResourceRemoteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccount** | Pointer to [**ResourceRemoteInfoAwsAccount**](ResourceRemoteInfoAwsAccount.md) |  | [optional] 
**AwsPermissionSet** | Pointer to [**ResourceRemoteInfoAwsPermissionSet**](ResourceRemoteInfoAwsPermissionSet.md) |  | [optional] 
**AwsIamRole** | Pointer to [**ResourceRemoteInfoAwsIamRole**](ResourceRemoteInfoAwsIamRole.md) |  | [optional] 
**AwsEc2Instance** | Pointer to [**ResourceRemoteInfoAwsEc2Instance**](ResourceRemoteInfoAwsEc2Instance.md) |  | [optional] 
**AwsRdsInstance** | Pointer to [**ResourceRemoteInfoAwsRdsInstance**](ResourceRemoteInfoAwsRdsInstance.md) |  | [optional] 
**AwsEksCluster** | Pointer to [**ResourceRemoteInfoAwsEksCluster**](ResourceRemoteInfoAwsEksCluster.md) |  | [optional] 
**GcpBucket** | Pointer to [**ResourceRemoteInfoGcpBucket**](ResourceRemoteInfoGcpBucket.md) |  | [optional] 
**GcpComputeInstance** | Pointer to [**ResourceRemoteInfoGcpComputeInstance**](ResourceRemoteInfoGcpComputeInstance.md) |  | [optional] 
**GcpFolder** | Pointer to [**ResourceRemoteInfoGcpFolder**](ResourceRemoteInfoGcpFolder.md) |  | [optional] 
**GcpGkeCluster** | Pointer to [**ResourceRemoteInfoGcpGkeCluster**](ResourceRemoteInfoGcpGkeCluster.md) |  | [optional] 
**GcpProject** | Pointer to [**ResourceRemoteInfoGcpProject**](ResourceRemoteInfoGcpProject.md) |  | [optional] 
**GcpSqlInstance** | Pointer to [**ResourceRemoteInfoGcpSqlInstance**](ResourceRemoteInfoGcpSqlInstance.md) |  | [optional] 
**GithubRepo** | Pointer to [**ResourceRemoteInfoGithubRepo**](ResourceRemoteInfoGithubRepo.md) |  | [optional] 
**GitlabProject** | Pointer to [**ResourceRemoteInfoGitlabProject**](ResourceRemoteInfoGitlabProject.md) |  | [optional] 
**OktaApp** | Pointer to [**ResourceRemoteInfoOktaApp**](ResourceRemoteInfoOktaApp.md) |  | [optional] 
**OktaStandardRole** | Pointer to [**ResourceRemoteInfoOktaStandardRole**](ResourceRemoteInfoOktaStandardRole.md) |  | [optional] 
**OktaCustomRole** | Pointer to [**ResourceRemoteInfoOktaCustomRole**](ResourceRemoteInfoOktaCustomRole.md) |  | [optional] 
**PagerdutyRole** | Pointer to [**ResourceRemoteInfoPagerdutyRole**](ResourceRemoteInfoPagerdutyRole.md) |  | [optional] 
**SalesforcePermissionSet** | Pointer to [**ResourceRemoteInfoSalesforcePermissionSet**](ResourceRemoteInfoSalesforcePermissionSet.md) |  | [optional] 
**SalesforceProfile** | Pointer to [**ResourceRemoteInfoSalesforceProfile**](ResourceRemoteInfoSalesforceProfile.md) |  | [optional] 
**SalesforceRole** | Pointer to [**ResourceRemoteInfoSalesforceRole**](ResourceRemoteInfoSalesforceRole.md) |  | [optional] 
**TeleportRole** | Pointer to [**ResourceRemoteInfoTeleportRole**](ResourceRemoteInfoTeleportRole.md) |  | [optional] 

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

### GetAwsAccount

`func (o *ResourceRemoteInfo) GetAwsAccount() ResourceRemoteInfoAwsAccount`

GetAwsAccount returns the AwsAccount field if non-nil, zero value otherwise.

### GetAwsAccountOk

`func (o *ResourceRemoteInfo) GetAwsAccountOk() (*ResourceRemoteInfoAwsAccount, bool)`

GetAwsAccountOk returns a tuple with the AwsAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccount

`func (o *ResourceRemoteInfo) SetAwsAccount(v ResourceRemoteInfoAwsAccount)`

SetAwsAccount sets AwsAccount field to given value.

### HasAwsAccount

`func (o *ResourceRemoteInfo) HasAwsAccount() bool`

HasAwsAccount returns a boolean if a field has been set.

### GetAwsPermissionSet

`func (o *ResourceRemoteInfo) GetAwsPermissionSet() ResourceRemoteInfoAwsPermissionSet`

GetAwsPermissionSet returns the AwsPermissionSet field if non-nil, zero value otherwise.

### GetAwsPermissionSetOk

`func (o *ResourceRemoteInfo) GetAwsPermissionSetOk() (*ResourceRemoteInfoAwsPermissionSet, bool)`

GetAwsPermissionSetOk returns a tuple with the AwsPermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsPermissionSet

`func (o *ResourceRemoteInfo) SetAwsPermissionSet(v ResourceRemoteInfoAwsPermissionSet)`

SetAwsPermissionSet sets AwsPermissionSet field to given value.

### HasAwsPermissionSet

`func (o *ResourceRemoteInfo) HasAwsPermissionSet() bool`

HasAwsPermissionSet returns a boolean if a field has been set.

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

### GetGcpBucket

`func (o *ResourceRemoteInfo) GetGcpBucket() ResourceRemoteInfoGcpBucket`

GetGcpBucket returns the GcpBucket field if non-nil, zero value otherwise.

### GetGcpBucketOk

`func (o *ResourceRemoteInfo) GetGcpBucketOk() (*ResourceRemoteInfoGcpBucket, bool)`

GetGcpBucketOk returns a tuple with the GcpBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpBucket

`func (o *ResourceRemoteInfo) SetGcpBucket(v ResourceRemoteInfoGcpBucket)`

SetGcpBucket sets GcpBucket field to given value.

### HasGcpBucket

`func (o *ResourceRemoteInfo) HasGcpBucket() bool`

HasGcpBucket returns a boolean if a field has been set.

### GetGcpComputeInstance

`func (o *ResourceRemoteInfo) GetGcpComputeInstance() ResourceRemoteInfoGcpComputeInstance`

GetGcpComputeInstance returns the GcpComputeInstance field if non-nil, zero value otherwise.

### GetGcpComputeInstanceOk

`func (o *ResourceRemoteInfo) GetGcpComputeInstanceOk() (*ResourceRemoteInfoGcpComputeInstance, bool)`

GetGcpComputeInstanceOk returns a tuple with the GcpComputeInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpComputeInstance

`func (o *ResourceRemoteInfo) SetGcpComputeInstance(v ResourceRemoteInfoGcpComputeInstance)`

SetGcpComputeInstance sets GcpComputeInstance field to given value.

### HasGcpComputeInstance

`func (o *ResourceRemoteInfo) HasGcpComputeInstance() bool`

HasGcpComputeInstance returns a boolean if a field has been set.

### GetGcpFolder

`func (o *ResourceRemoteInfo) GetGcpFolder() ResourceRemoteInfoGcpFolder`

GetGcpFolder returns the GcpFolder field if non-nil, zero value otherwise.

### GetGcpFolderOk

`func (o *ResourceRemoteInfo) GetGcpFolderOk() (*ResourceRemoteInfoGcpFolder, bool)`

GetGcpFolderOk returns a tuple with the GcpFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpFolder

`func (o *ResourceRemoteInfo) SetGcpFolder(v ResourceRemoteInfoGcpFolder)`

SetGcpFolder sets GcpFolder field to given value.

### HasGcpFolder

`func (o *ResourceRemoteInfo) HasGcpFolder() bool`

HasGcpFolder returns a boolean if a field has been set.

### GetGcpGkeCluster

`func (o *ResourceRemoteInfo) GetGcpGkeCluster() ResourceRemoteInfoGcpGkeCluster`

GetGcpGkeCluster returns the GcpGkeCluster field if non-nil, zero value otherwise.

### GetGcpGkeClusterOk

`func (o *ResourceRemoteInfo) GetGcpGkeClusterOk() (*ResourceRemoteInfoGcpGkeCluster, bool)`

GetGcpGkeClusterOk returns a tuple with the GcpGkeCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpGkeCluster

`func (o *ResourceRemoteInfo) SetGcpGkeCluster(v ResourceRemoteInfoGcpGkeCluster)`

SetGcpGkeCluster sets GcpGkeCluster field to given value.

### HasGcpGkeCluster

`func (o *ResourceRemoteInfo) HasGcpGkeCluster() bool`

HasGcpGkeCluster returns a boolean if a field has been set.

### GetGcpProject

`func (o *ResourceRemoteInfo) GetGcpProject() ResourceRemoteInfoGcpProject`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *ResourceRemoteInfo) GetGcpProjectOk() (*ResourceRemoteInfoGcpProject, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *ResourceRemoteInfo) SetGcpProject(v ResourceRemoteInfoGcpProject)`

SetGcpProject sets GcpProject field to given value.

### HasGcpProject

`func (o *ResourceRemoteInfo) HasGcpProject() bool`

HasGcpProject returns a boolean if a field has been set.

### GetGcpSqlInstance

`func (o *ResourceRemoteInfo) GetGcpSqlInstance() ResourceRemoteInfoGcpSqlInstance`

GetGcpSqlInstance returns the GcpSqlInstance field if non-nil, zero value otherwise.

### GetGcpSqlInstanceOk

`func (o *ResourceRemoteInfo) GetGcpSqlInstanceOk() (*ResourceRemoteInfoGcpSqlInstance, bool)`

GetGcpSqlInstanceOk returns a tuple with the GcpSqlInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSqlInstance

`func (o *ResourceRemoteInfo) SetGcpSqlInstance(v ResourceRemoteInfoGcpSqlInstance)`

SetGcpSqlInstance sets GcpSqlInstance field to given value.

### HasGcpSqlInstance

`func (o *ResourceRemoteInfo) HasGcpSqlInstance() bool`

HasGcpSqlInstance returns a boolean if a field has been set.

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

### GetGitlabProject

`func (o *ResourceRemoteInfo) GetGitlabProject() ResourceRemoteInfoGitlabProject`

GetGitlabProject returns the GitlabProject field if non-nil, zero value otherwise.

### GetGitlabProjectOk

`func (o *ResourceRemoteInfo) GetGitlabProjectOk() (*ResourceRemoteInfoGitlabProject, bool)`

GetGitlabProjectOk returns a tuple with the GitlabProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabProject

`func (o *ResourceRemoteInfo) SetGitlabProject(v ResourceRemoteInfoGitlabProject)`

SetGitlabProject sets GitlabProject field to given value.

### HasGitlabProject

`func (o *ResourceRemoteInfo) HasGitlabProject() bool`

HasGitlabProject returns a boolean if a field has been set.

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

### GetPagerdutyRole

`func (o *ResourceRemoteInfo) GetPagerdutyRole() ResourceRemoteInfoPagerdutyRole`

GetPagerdutyRole returns the PagerdutyRole field if non-nil, zero value otherwise.

### GetPagerdutyRoleOk

`func (o *ResourceRemoteInfo) GetPagerdutyRoleOk() (*ResourceRemoteInfoPagerdutyRole, bool)`

GetPagerdutyRoleOk returns a tuple with the PagerdutyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerdutyRole

`func (o *ResourceRemoteInfo) SetPagerdutyRole(v ResourceRemoteInfoPagerdutyRole)`

SetPagerdutyRole sets PagerdutyRole field to given value.

### HasPagerdutyRole

`func (o *ResourceRemoteInfo) HasPagerdutyRole() bool`

HasPagerdutyRole returns a boolean if a field has been set.

### GetSalesforcePermissionSet

`func (o *ResourceRemoteInfo) GetSalesforcePermissionSet() ResourceRemoteInfoSalesforcePermissionSet`

GetSalesforcePermissionSet returns the SalesforcePermissionSet field if non-nil, zero value otherwise.

### GetSalesforcePermissionSetOk

`func (o *ResourceRemoteInfo) GetSalesforcePermissionSetOk() (*ResourceRemoteInfoSalesforcePermissionSet, bool)`

GetSalesforcePermissionSetOk returns a tuple with the SalesforcePermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesforcePermissionSet

`func (o *ResourceRemoteInfo) SetSalesforcePermissionSet(v ResourceRemoteInfoSalesforcePermissionSet)`

SetSalesforcePermissionSet sets SalesforcePermissionSet field to given value.

### HasSalesforcePermissionSet

`func (o *ResourceRemoteInfo) HasSalesforcePermissionSet() bool`

HasSalesforcePermissionSet returns a boolean if a field has been set.

### GetSalesforceProfile

`func (o *ResourceRemoteInfo) GetSalesforceProfile() ResourceRemoteInfoSalesforceProfile`

GetSalesforceProfile returns the SalesforceProfile field if non-nil, zero value otherwise.

### GetSalesforceProfileOk

`func (o *ResourceRemoteInfo) GetSalesforceProfileOk() (*ResourceRemoteInfoSalesforceProfile, bool)`

GetSalesforceProfileOk returns a tuple with the SalesforceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesforceProfile

`func (o *ResourceRemoteInfo) SetSalesforceProfile(v ResourceRemoteInfoSalesforceProfile)`

SetSalesforceProfile sets SalesforceProfile field to given value.

### HasSalesforceProfile

`func (o *ResourceRemoteInfo) HasSalesforceProfile() bool`

HasSalesforceProfile returns a boolean if a field has been set.

### GetSalesforceRole

`func (o *ResourceRemoteInfo) GetSalesforceRole() ResourceRemoteInfoSalesforceRole`

GetSalesforceRole returns the SalesforceRole field if non-nil, zero value otherwise.

### GetSalesforceRoleOk

`func (o *ResourceRemoteInfo) GetSalesforceRoleOk() (*ResourceRemoteInfoSalesforceRole, bool)`

GetSalesforceRoleOk returns a tuple with the SalesforceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesforceRole

`func (o *ResourceRemoteInfo) SetSalesforceRole(v ResourceRemoteInfoSalesforceRole)`

SetSalesforceRole sets SalesforceRole field to given value.

### HasSalesforceRole

`func (o *ResourceRemoteInfo) HasSalesforceRole() bool`

HasSalesforceRole returns a boolean if a field has been set.

### GetTeleportRole

`func (o *ResourceRemoteInfo) GetTeleportRole() ResourceRemoteInfoTeleportRole`

GetTeleportRole returns the TeleportRole field if non-nil, zero value otherwise.

### GetTeleportRoleOk

`func (o *ResourceRemoteInfo) GetTeleportRoleOk() (*ResourceRemoteInfoTeleportRole, bool)`

GetTeleportRoleOk returns a tuple with the TeleportRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeleportRole

`func (o *ResourceRemoteInfo) SetTeleportRole(v ResourceRemoteInfoTeleportRole)`

SetTeleportRole sets TeleportRole field to given value.

### HasTeleportRole

`func (o *ResourceRemoteInfo) HasTeleportRole() bool`

HasTeleportRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


