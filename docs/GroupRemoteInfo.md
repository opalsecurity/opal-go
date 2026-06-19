# GroupRemoteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryGroup** | Pointer to [**GroupRemoteInfoActiveDirectoryGroup**](GroupRemoteInfoActiveDirectoryGroup.md) |  | [optional] 
**TailscaleGroup** | Pointer to [**GroupRemoteInfoTailscaleGroup**](GroupRemoteInfoTailscaleGroup.md) |  | [optional] 
**TwingateGroup** | Pointer to [**GroupRemoteInfoTwingateGroup**](GroupRemoteInfoTwingateGroup.md) |  | [optional] 
**TwingateGroupSynced** | Pointer to [**GroupRemoteInfoTwingateGroupSynced**](GroupRemoteInfoTwingateGroupSynced.md) |  | [optional] 
**AwsSsoGroup** | Pointer to [**GroupRemoteInfoAwsSsoGroup**](GroupRemoteInfoAwsSsoGroup.md) |  | [optional] 
**DatabricksAccountGroup** | Pointer to [**GroupRemoteInfoDatabricksAccountGroup**](GroupRemoteInfoDatabricksAccountGroup.md) |  | [optional] 
**ConnectorGroup** | Pointer to [**GroupRemoteInfoConnectorGroup**](GroupRemoteInfoConnectorGroup.md) |  | [optional] 
**GithubTeam** | Pointer to [**GroupRemoteInfoGithubTeam**](GroupRemoteInfoGithubTeam.md) |  | [optional] 
**GithubEnterpriseTeam** | Pointer to [**GroupRemoteInfoGithubEnterpriseTeam**](GroupRemoteInfoGithubEnterpriseTeam.md) |  | [optional] 
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
**PagerdutyOnCallSchedule** | Pointer to [**GroupRemoteInfoPagerdutyOnCallSchedule**](GroupRemoteInfoPagerdutyOnCallSchedule.md) |  | [optional] 
**IncidentioOnCallSchedule** | Pointer to [**GroupRemoteInfoIncidentioOnCallSchedule**](GroupRemoteInfoIncidentioOnCallSchedule.md) |  | [optional] 
**RootlyOnCallSchedule** | Pointer to [**GroupRemoteInfoRootlyOnCallSchedule**](GroupRemoteInfoRootlyOnCallSchedule.md) |  | [optional] 
**DevinGroup** | Pointer to [**GroupRemoteInfoDevinGroup**](GroupRemoteInfoDevinGroup.md) |  | [optional] 
**ClickhouseRole** | Pointer to [**GroupRemoteInfoClickhouseRole**](GroupRemoteInfoClickhouseRole.md) |  | [optional] 
**GrafanaTeam** | Pointer to [**GroupRemoteInfoGrafanaTeam**](GroupRemoteInfoGrafanaTeam.md) |  | [optional] 
**ZendeskGroup** | Pointer to [**GroupRemoteInfoZendeskGroup**](GroupRemoteInfoZendeskGroup.md) |  | [optional] 
**SlackUserGroup** | Pointer to [**GroupRemoteInfoSlackUserGroup**](GroupRemoteInfoSlackUserGroup.md) |  | [optional] 
**ZendeskOrganization** | Pointer to [**GroupRemoteInfoZendeskOrganization**](GroupRemoteInfoZendeskOrganization.md) |  | [optional] 

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

### GetTailscaleGroup

`func (o *GroupRemoteInfo) GetTailscaleGroup() GroupRemoteInfoTailscaleGroup`

GetTailscaleGroup returns the TailscaleGroup field if non-nil, zero value otherwise.

### GetTailscaleGroupOk

`func (o *GroupRemoteInfo) GetTailscaleGroupOk() (*GroupRemoteInfoTailscaleGroup, bool)`

GetTailscaleGroupOk returns a tuple with the TailscaleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTailscaleGroup

`func (o *GroupRemoteInfo) SetTailscaleGroup(v GroupRemoteInfoTailscaleGroup)`

SetTailscaleGroup sets TailscaleGroup field to given value.

### HasTailscaleGroup

`func (o *GroupRemoteInfo) HasTailscaleGroup() bool`

HasTailscaleGroup returns a boolean if a field has been set.

### GetTwingateGroup

`func (o *GroupRemoteInfo) GetTwingateGroup() GroupRemoteInfoTwingateGroup`

GetTwingateGroup returns the TwingateGroup field if non-nil, zero value otherwise.

### GetTwingateGroupOk

`func (o *GroupRemoteInfo) GetTwingateGroupOk() (*GroupRemoteInfoTwingateGroup, bool)`

GetTwingateGroupOk returns a tuple with the TwingateGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwingateGroup

`func (o *GroupRemoteInfo) SetTwingateGroup(v GroupRemoteInfoTwingateGroup)`

SetTwingateGroup sets TwingateGroup field to given value.

### HasTwingateGroup

`func (o *GroupRemoteInfo) HasTwingateGroup() bool`

HasTwingateGroup returns a boolean if a field has been set.

### GetTwingateGroupSynced

`func (o *GroupRemoteInfo) GetTwingateGroupSynced() GroupRemoteInfoTwingateGroupSynced`

GetTwingateGroupSynced returns the TwingateGroupSynced field if non-nil, zero value otherwise.

### GetTwingateGroupSyncedOk

`func (o *GroupRemoteInfo) GetTwingateGroupSyncedOk() (*GroupRemoteInfoTwingateGroupSynced, bool)`

GetTwingateGroupSyncedOk returns a tuple with the TwingateGroupSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwingateGroupSynced

`func (o *GroupRemoteInfo) SetTwingateGroupSynced(v GroupRemoteInfoTwingateGroupSynced)`

SetTwingateGroupSynced sets TwingateGroupSynced field to given value.

### HasTwingateGroupSynced

`func (o *GroupRemoteInfo) HasTwingateGroupSynced() bool`

HasTwingateGroupSynced returns a boolean if a field has been set.

### GetAwsSsoGroup

`func (o *GroupRemoteInfo) GetAwsSsoGroup() GroupRemoteInfoAwsSsoGroup`

GetAwsSsoGroup returns the AwsSsoGroup field if non-nil, zero value otherwise.

### GetAwsSsoGroupOk

`func (o *GroupRemoteInfo) GetAwsSsoGroupOk() (*GroupRemoteInfoAwsSsoGroup, bool)`

GetAwsSsoGroupOk returns a tuple with the AwsSsoGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSsoGroup

`func (o *GroupRemoteInfo) SetAwsSsoGroup(v GroupRemoteInfoAwsSsoGroup)`

SetAwsSsoGroup sets AwsSsoGroup field to given value.

### HasAwsSsoGroup

`func (o *GroupRemoteInfo) HasAwsSsoGroup() bool`

HasAwsSsoGroup returns a boolean if a field has been set.

### GetDatabricksAccountGroup

`func (o *GroupRemoteInfo) GetDatabricksAccountGroup() GroupRemoteInfoDatabricksAccountGroup`

GetDatabricksAccountGroup returns the DatabricksAccountGroup field if non-nil, zero value otherwise.

### GetDatabricksAccountGroupOk

`func (o *GroupRemoteInfo) GetDatabricksAccountGroupOk() (*GroupRemoteInfoDatabricksAccountGroup, bool)`

GetDatabricksAccountGroupOk returns a tuple with the DatabricksAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabricksAccountGroup

`func (o *GroupRemoteInfo) SetDatabricksAccountGroup(v GroupRemoteInfoDatabricksAccountGroup)`

SetDatabricksAccountGroup sets DatabricksAccountGroup field to given value.

### HasDatabricksAccountGroup

`func (o *GroupRemoteInfo) HasDatabricksAccountGroup() bool`

HasDatabricksAccountGroup returns a boolean if a field has been set.

### GetConnectorGroup

`func (o *GroupRemoteInfo) GetConnectorGroup() GroupRemoteInfoConnectorGroup`

GetConnectorGroup returns the ConnectorGroup field if non-nil, zero value otherwise.

### GetConnectorGroupOk

`func (o *GroupRemoteInfo) GetConnectorGroupOk() (*GroupRemoteInfoConnectorGroup, bool)`

GetConnectorGroupOk returns a tuple with the ConnectorGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorGroup

`func (o *GroupRemoteInfo) SetConnectorGroup(v GroupRemoteInfoConnectorGroup)`

SetConnectorGroup sets ConnectorGroup field to given value.

### HasConnectorGroup

`func (o *GroupRemoteInfo) HasConnectorGroup() bool`

HasConnectorGroup returns a boolean if a field has been set.

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

### GetGithubEnterpriseTeam

`func (o *GroupRemoteInfo) GetGithubEnterpriseTeam() GroupRemoteInfoGithubEnterpriseTeam`

GetGithubEnterpriseTeam returns the GithubEnterpriseTeam field if non-nil, zero value otherwise.

### GetGithubEnterpriseTeamOk

`func (o *GroupRemoteInfo) GetGithubEnterpriseTeamOk() (*GroupRemoteInfoGithubEnterpriseTeam, bool)`

GetGithubEnterpriseTeamOk returns a tuple with the GithubEnterpriseTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubEnterpriseTeam

`func (o *GroupRemoteInfo) SetGithubEnterpriseTeam(v GroupRemoteInfoGithubEnterpriseTeam)`

SetGithubEnterpriseTeam sets GithubEnterpriseTeam field to given value.

### HasGithubEnterpriseTeam

`func (o *GroupRemoteInfo) HasGithubEnterpriseTeam() bool`

HasGithubEnterpriseTeam returns a boolean if a field has been set.

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

### GetPagerdutyOnCallSchedule

`func (o *GroupRemoteInfo) GetPagerdutyOnCallSchedule() GroupRemoteInfoPagerdutyOnCallSchedule`

GetPagerdutyOnCallSchedule returns the PagerdutyOnCallSchedule field if non-nil, zero value otherwise.

### GetPagerdutyOnCallScheduleOk

`func (o *GroupRemoteInfo) GetPagerdutyOnCallScheduleOk() (*GroupRemoteInfoPagerdutyOnCallSchedule, bool)`

GetPagerdutyOnCallScheduleOk returns a tuple with the PagerdutyOnCallSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerdutyOnCallSchedule

`func (o *GroupRemoteInfo) SetPagerdutyOnCallSchedule(v GroupRemoteInfoPagerdutyOnCallSchedule)`

SetPagerdutyOnCallSchedule sets PagerdutyOnCallSchedule field to given value.

### HasPagerdutyOnCallSchedule

`func (o *GroupRemoteInfo) HasPagerdutyOnCallSchedule() bool`

HasPagerdutyOnCallSchedule returns a boolean if a field has been set.

### GetIncidentioOnCallSchedule

`func (o *GroupRemoteInfo) GetIncidentioOnCallSchedule() GroupRemoteInfoIncidentioOnCallSchedule`

GetIncidentioOnCallSchedule returns the IncidentioOnCallSchedule field if non-nil, zero value otherwise.

### GetIncidentioOnCallScheduleOk

`func (o *GroupRemoteInfo) GetIncidentioOnCallScheduleOk() (*GroupRemoteInfoIncidentioOnCallSchedule, bool)`

GetIncidentioOnCallScheduleOk returns a tuple with the IncidentioOnCallSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentioOnCallSchedule

`func (o *GroupRemoteInfo) SetIncidentioOnCallSchedule(v GroupRemoteInfoIncidentioOnCallSchedule)`

SetIncidentioOnCallSchedule sets IncidentioOnCallSchedule field to given value.

### HasIncidentioOnCallSchedule

`func (o *GroupRemoteInfo) HasIncidentioOnCallSchedule() bool`

HasIncidentioOnCallSchedule returns a boolean if a field has been set.

### GetRootlyOnCallSchedule

`func (o *GroupRemoteInfo) GetRootlyOnCallSchedule() GroupRemoteInfoRootlyOnCallSchedule`

GetRootlyOnCallSchedule returns the RootlyOnCallSchedule field if non-nil, zero value otherwise.

### GetRootlyOnCallScheduleOk

`func (o *GroupRemoteInfo) GetRootlyOnCallScheduleOk() (*GroupRemoteInfoRootlyOnCallSchedule, bool)`

GetRootlyOnCallScheduleOk returns a tuple with the RootlyOnCallSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootlyOnCallSchedule

`func (o *GroupRemoteInfo) SetRootlyOnCallSchedule(v GroupRemoteInfoRootlyOnCallSchedule)`

SetRootlyOnCallSchedule sets RootlyOnCallSchedule field to given value.

### HasRootlyOnCallSchedule

`func (o *GroupRemoteInfo) HasRootlyOnCallSchedule() bool`

HasRootlyOnCallSchedule returns a boolean if a field has been set.

### GetDevinGroup

`func (o *GroupRemoteInfo) GetDevinGroup() GroupRemoteInfoDevinGroup`

GetDevinGroup returns the DevinGroup field if non-nil, zero value otherwise.

### GetDevinGroupOk

`func (o *GroupRemoteInfo) GetDevinGroupOk() (*GroupRemoteInfoDevinGroup, bool)`

GetDevinGroupOk returns a tuple with the DevinGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevinGroup

`func (o *GroupRemoteInfo) SetDevinGroup(v GroupRemoteInfoDevinGroup)`

SetDevinGroup sets DevinGroup field to given value.

### HasDevinGroup

`func (o *GroupRemoteInfo) HasDevinGroup() bool`

HasDevinGroup returns a boolean if a field has been set.

### GetClickhouseRole

`func (o *GroupRemoteInfo) GetClickhouseRole() GroupRemoteInfoClickhouseRole`

GetClickhouseRole returns the ClickhouseRole field if non-nil, zero value otherwise.

### GetClickhouseRoleOk

`func (o *GroupRemoteInfo) GetClickhouseRoleOk() (*GroupRemoteInfoClickhouseRole, bool)`

GetClickhouseRoleOk returns a tuple with the ClickhouseRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouseRole

`func (o *GroupRemoteInfo) SetClickhouseRole(v GroupRemoteInfoClickhouseRole)`

SetClickhouseRole sets ClickhouseRole field to given value.

### HasClickhouseRole

`func (o *GroupRemoteInfo) HasClickhouseRole() bool`

HasClickhouseRole returns a boolean if a field has been set.

### GetGrafanaTeam

`func (o *GroupRemoteInfo) GetGrafanaTeam() GroupRemoteInfoGrafanaTeam`

GetGrafanaTeam returns the GrafanaTeam field if non-nil, zero value otherwise.

### GetGrafanaTeamOk

`func (o *GroupRemoteInfo) GetGrafanaTeamOk() (*GroupRemoteInfoGrafanaTeam, bool)`

GetGrafanaTeamOk returns a tuple with the GrafanaTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaTeam

`func (o *GroupRemoteInfo) SetGrafanaTeam(v GroupRemoteInfoGrafanaTeam)`

SetGrafanaTeam sets GrafanaTeam field to given value.

### HasGrafanaTeam

`func (o *GroupRemoteInfo) HasGrafanaTeam() bool`

HasGrafanaTeam returns a boolean if a field has been set.

### GetZendeskGroup

`func (o *GroupRemoteInfo) GetZendeskGroup() GroupRemoteInfoZendeskGroup`

GetZendeskGroup returns the ZendeskGroup field if non-nil, zero value otherwise.

### GetZendeskGroupOk

`func (o *GroupRemoteInfo) GetZendeskGroupOk() (*GroupRemoteInfoZendeskGroup, bool)`

GetZendeskGroupOk returns a tuple with the ZendeskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskGroup

`func (o *GroupRemoteInfo) SetZendeskGroup(v GroupRemoteInfoZendeskGroup)`

SetZendeskGroup sets ZendeskGroup field to given value.

### HasZendeskGroup

`func (o *GroupRemoteInfo) HasZendeskGroup() bool`

HasZendeskGroup returns a boolean if a field has been set.

### GetSlackUserGroup

`func (o *GroupRemoteInfo) GetSlackUserGroup() GroupRemoteInfoSlackUserGroup`

GetSlackUserGroup returns the SlackUserGroup field if non-nil, zero value otherwise.

### GetSlackUserGroupOk

`func (o *GroupRemoteInfo) GetSlackUserGroupOk() (*GroupRemoteInfoSlackUserGroup, bool)`

GetSlackUserGroupOk returns a tuple with the SlackUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUserGroup

`func (o *GroupRemoteInfo) SetSlackUserGroup(v GroupRemoteInfoSlackUserGroup)`

SetSlackUserGroup sets SlackUserGroup field to given value.

### HasSlackUserGroup

`func (o *GroupRemoteInfo) HasSlackUserGroup() bool`

HasSlackUserGroup returns a boolean if a field has been set.

### GetZendeskOrganization

`func (o *GroupRemoteInfo) GetZendeskOrganization() GroupRemoteInfoZendeskOrganization`

GetZendeskOrganization returns the ZendeskOrganization field if non-nil, zero value otherwise.

### GetZendeskOrganizationOk

`func (o *GroupRemoteInfo) GetZendeskOrganizationOk() (*GroupRemoteInfoZendeskOrganization, bool)`

GetZendeskOrganizationOk returns a tuple with the ZendeskOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskOrganization

`func (o *GroupRemoteInfo) SetZendeskOrganization(v GroupRemoteInfoZendeskOrganization)`

SetZendeskOrganization sets ZendeskOrganization field to given value.

### HasZendeskOrganization

`func (o *GroupRemoteInfo) HasZendeskOrganization() bool`

HasZendeskOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


