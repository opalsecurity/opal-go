# GroupRemoteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryGroup** | Pointer to [**GroupRemoteInfoActiveDirectoryGroup**](GroupRemoteInfoActiveDirectoryGroup.md) |  | [optional] 
**GithubTeam** | Pointer to [**GroupRemoteInfoGithubTeam**](GroupRemoteInfoGithubTeam.md) |  | [optional] 
**GitlabGroup** | Pointer to [**GroupRemoteInfoGitlabGroup**](GroupRemoteInfoGitlabGroup.md) |  | [optional] 
**GoogleGroup** | Pointer to [**GroupRemoteInfoGoogleGroup**](GroupRemoteInfoGoogleGroup.md) |  | [optional] 
**LdapGroup** | Pointer to [**GroupRemoteInfoLdapGroup**](GroupRemoteInfoLdapGroup.md) |  | [optional] 
**OktaGroup** | Pointer to [**GroupRemoteInfoOktaGroup**](GroupRemoteInfoOktaGroup.md) |  | [optional] 
**DuoGroup** | Pointer to [**GroupRemoteInfoDuoGroup**](GroupRemoteInfoDuoGroup.md) |  | [optional] 
**AzureAdSecurityGroup** | Pointer to [**GroupRemoteInfoAzureAdSecurityGroup**](GroupRemoteInfoAzureAdSecurityGroup.md) |  | [optional] 
**AzureAdMicrosoft365Group** | Pointer to [**GroupRemoteInfoAzureAdMicrosoft365Group**](GroupRemoteInfoAzureAdMicrosoft365Group.md) |  | [optional] 
**SnowflakeRole** | Pointer to [**GroupRemoteInfoSnowflakeRole**](GroupRemoteInfoSnowflakeRole.md) |  | [optional] 
**OktaGroupRule** | Pointer to [**GroupRemoteInfoOktaGroupRule**](GroupRemoteInfoOktaGroupRule.md) |  | [optional] 
**WorkdayUserSecurityGroup** | Pointer to [**GroupRemoteInfoWorkdayUserSecurityGroup**](GroupRemoteInfoWorkdayUserSecurityGroup.md) |  | [optional] 

## Methods

### NewGroupRemoteInfo

`func NewGroupRemoteInfo() *GroupRemoteInfo`

NewGroupRemoteInfo instantiates a new GroupRemoteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRemoteInfoWithDefaults

`func NewGroupRemoteInfoWithDefaults() *GroupRemoteInfo`

NewGroupRemoteInfoWithDefaults instantiates a new GroupRemoteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryGroup

`func (o *GroupRemoteInfo) GetActiveDirectoryGroup() GroupRemoteInfoActiveDirectoryGroup`

GetActiveDirectoryGroup returns the ActiveDirectoryGroup field if non-nil, zero value otherwise.

### GetActiveDirectoryGroupOk

`func (o *GroupRemoteInfo) GetActiveDirectoryGroupOk() (*GroupRemoteInfoActiveDirectoryGroup, bool)`

GetActiveDirectoryGroupOk returns a tuple with the ActiveDirectoryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryGroup

`func (o *GroupRemoteInfo) SetActiveDirectoryGroup(v GroupRemoteInfoActiveDirectoryGroup)`

SetActiveDirectoryGroup sets ActiveDirectoryGroup field to given value.

### HasActiveDirectoryGroup

`func (o *GroupRemoteInfo) HasActiveDirectoryGroup() bool`

HasActiveDirectoryGroup returns a boolean if a field has been set.

### GetGithubTeam

`func (o *GroupRemoteInfo) GetGithubTeam() GroupRemoteInfoGithubTeam`

GetGithubTeam returns the GithubTeam field if non-nil, zero value otherwise.

### GetGithubTeamOk

`func (o *GroupRemoteInfo) GetGithubTeamOk() (*GroupRemoteInfoGithubTeam, bool)`

GetGithubTeamOk returns a tuple with the GithubTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubTeam

`func (o *GroupRemoteInfo) SetGithubTeam(v GroupRemoteInfoGithubTeam)`

SetGithubTeam sets GithubTeam field to given value.

### HasGithubTeam

`func (o *GroupRemoteInfo) HasGithubTeam() bool`

HasGithubTeam returns a boolean if a field has been set.

### GetGitlabGroup

`func (o *GroupRemoteInfo) GetGitlabGroup() GroupRemoteInfoGitlabGroup`

GetGitlabGroup returns the GitlabGroup field if non-nil, zero value otherwise.

### GetGitlabGroupOk

`func (o *GroupRemoteInfo) GetGitlabGroupOk() (*GroupRemoteInfoGitlabGroup, bool)`

GetGitlabGroupOk returns a tuple with the GitlabGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabGroup

`func (o *GroupRemoteInfo) SetGitlabGroup(v GroupRemoteInfoGitlabGroup)`

SetGitlabGroup sets GitlabGroup field to given value.

### HasGitlabGroup

`func (o *GroupRemoteInfo) HasGitlabGroup() bool`

HasGitlabGroup returns a boolean if a field has been set.

### GetGoogleGroup

`func (o *GroupRemoteInfo) GetGoogleGroup() GroupRemoteInfoGoogleGroup`

GetGoogleGroup returns the GoogleGroup field if non-nil, zero value otherwise.

### GetGoogleGroupOk

`func (o *GroupRemoteInfo) GetGoogleGroupOk() (*GroupRemoteInfoGoogleGroup, bool)`

GetGoogleGroupOk returns a tuple with the GoogleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleGroup

`func (o *GroupRemoteInfo) SetGoogleGroup(v GroupRemoteInfoGoogleGroup)`

SetGoogleGroup sets GoogleGroup field to given value.

### HasGoogleGroup

`func (o *GroupRemoteInfo) HasGoogleGroup() bool`

HasGoogleGroup returns a boolean if a field has been set.

### GetLdapGroup

`func (o *GroupRemoteInfo) GetLdapGroup() GroupRemoteInfoLdapGroup`

GetLdapGroup returns the LdapGroup field if non-nil, zero value otherwise.

### GetLdapGroupOk

`func (o *GroupRemoteInfo) GetLdapGroupOk() (*GroupRemoteInfoLdapGroup, bool)`

GetLdapGroupOk returns a tuple with the LdapGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroup

`func (o *GroupRemoteInfo) SetLdapGroup(v GroupRemoteInfoLdapGroup)`

SetLdapGroup sets LdapGroup field to given value.

### HasLdapGroup

`func (o *GroupRemoteInfo) HasLdapGroup() bool`

HasLdapGroup returns a boolean if a field has been set.

### GetOktaGroup

`func (o *GroupRemoteInfo) GetOktaGroup() GroupRemoteInfoOktaGroup`

GetOktaGroup returns the OktaGroup field if non-nil, zero value otherwise.

### GetOktaGroupOk

`func (o *GroupRemoteInfo) GetOktaGroupOk() (*GroupRemoteInfoOktaGroup, bool)`

GetOktaGroupOk returns a tuple with the OktaGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaGroup

`func (o *GroupRemoteInfo) SetOktaGroup(v GroupRemoteInfoOktaGroup)`

SetOktaGroup sets OktaGroup field to given value.

### HasOktaGroup

`func (o *GroupRemoteInfo) HasOktaGroup() bool`

HasOktaGroup returns a boolean if a field has been set.

### GetDuoGroup

`func (o *GroupRemoteInfo) GetDuoGroup() GroupRemoteInfoDuoGroup`

GetDuoGroup returns the DuoGroup field if non-nil, zero value otherwise.

### GetDuoGroupOk

`func (o *GroupRemoteInfo) GetDuoGroupOk() (*GroupRemoteInfoDuoGroup, bool)`

GetDuoGroupOk returns a tuple with the DuoGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuoGroup

`func (o *GroupRemoteInfo) SetDuoGroup(v GroupRemoteInfoDuoGroup)`

SetDuoGroup sets DuoGroup field to given value.

### HasDuoGroup

`func (o *GroupRemoteInfo) HasDuoGroup() bool`

HasDuoGroup returns a boolean if a field has been set.

### GetAzureAdSecurityGroup

`func (o *GroupRemoteInfo) GetAzureAdSecurityGroup() GroupRemoteInfoAzureAdSecurityGroup`

GetAzureAdSecurityGroup returns the AzureAdSecurityGroup field if non-nil, zero value otherwise.

### GetAzureAdSecurityGroupOk

`func (o *GroupRemoteInfo) GetAzureAdSecurityGroupOk() (*GroupRemoteInfoAzureAdSecurityGroup, bool)`

GetAzureAdSecurityGroupOk returns a tuple with the AzureAdSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdSecurityGroup

`func (o *GroupRemoteInfo) SetAzureAdSecurityGroup(v GroupRemoteInfoAzureAdSecurityGroup)`

SetAzureAdSecurityGroup sets AzureAdSecurityGroup field to given value.

### HasAzureAdSecurityGroup

`func (o *GroupRemoteInfo) HasAzureAdSecurityGroup() bool`

HasAzureAdSecurityGroup returns a boolean if a field has been set.

### GetAzureAdMicrosoft365Group

`func (o *GroupRemoteInfo) GetAzureAdMicrosoft365Group() GroupRemoteInfoAzureAdMicrosoft365Group`

GetAzureAdMicrosoft365Group returns the AzureAdMicrosoft365Group field if non-nil, zero value otherwise.

### GetAzureAdMicrosoft365GroupOk

`func (o *GroupRemoteInfo) GetAzureAdMicrosoft365GroupOk() (*GroupRemoteInfoAzureAdMicrosoft365Group, bool)`

GetAzureAdMicrosoft365GroupOk returns a tuple with the AzureAdMicrosoft365Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdMicrosoft365Group

`func (o *GroupRemoteInfo) SetAzureAdMicrosoft365Group(v GroupRemoteInfoAzureAdMicrosoft365Group)`

SetAzureAdMicrosoft365Group sets AzureAdMicrosoft365Group field to given value.

### HasAzureAdMicrosoft365Group

`func (o *GroupRemoteInfo) HasAzureAdMicrosoft365Group() bool`

HasAzureAdMicrosoft365Group returns a boolean if a field has been set.

### GetSnowflakeRole

`func (o *GroupRemoteInfo) GetSnowflakeRole() GroupRemoteInfoSnowflakeRole`

GetSnowflakeRole returns the SnowflakeRole field if non-nil, zero value otherwise.

### GetSnowflakeRoleOk

`func (o *GroupRemoteInfo) GetSnowflakeRoleOk() (*GroupRemoteInfoSnowflakeRole, bool)`

GetSnowflakeRoleOk returns a tuple with the SnowflakeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeRole

`func (o *GroupRemoteInfo) SetSnowflakeRole(v GroupRemoteInfoSnowflakeRole)`

SetSnowflakeRole sets SnowflakeRole field to given value.

### HasSnowflakeRole

`func (o *GroupRemoteInfo) HasSnowflakeRole() bool`

HasSnowflakeRole returns a boolean if a field has been set.

### GetOktaGroupRule

`func (o *GroupRemoteInfo) GetOktaGroupRule() GroupRemoteInfoOktaGroupRule`

GetOktaGroupRule returns the OktaGroupRule field if non-nil, zero value otherwise.

### GetOktaGroupRuleOk

`func (o *GroupRemoteInfo) GetOktaGroupRuleOk() (*GroupRemoteInfoOktaGroupRule, bool)`

GetOktaGroupRuleOk returns a tuple with the OktaGroupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaGroupRule

`func (o *GroupRemoteInfo) SetOktaGroupRule(v GroupRemoteInfoOktaGroupRule)`

SetOktaGroupRule sets OktaGroupRule field to given value.

### HasOktaGroupRule

`func (o *GroupRemoteInfo) HasOktaGroupRule() bool`

HasOktaGroupRule returns a boolean if a field has been set.

### GetWorkdayUserSecurityGroup

`func (o *GroupRemoteInfo) GetWorkdayUserSecurityGroup() GroupRemoteInfoWorkdayUserSecurityGroup`

GetWorkdayUserSecurityGroup returns the WorkdayUserSecurityGroup field if non-nil, zero value otherwise.

### GetWorkdayUserSecurityGroupOk

`func (o *GroupRemoteInfo) GetWorkdayUserSecurityGroupOk() (*GroupRemoteInfoWorkdayUserSecurityGroup, bool)`

GetWorkdayUserSecurityGroupOk returns a tuple with the WorkdayUserSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkdayUserSecurityGroup

`func (o *GroupRemoteInfo) SetWorkdayUserSecurityGroup(v GroupRemoteInfoWorkdayUserSecurityGroup)`

SetWorkdayUserSecurityGroup sets WorkdayUserSecurityGroup field to given value.

### HasWorkdayUserSecurityGroup

`func (o *GroupRemoteInfo) HasWorkdayUserSecurityGroup() bool`

HasWorkdayUserSecurityGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


