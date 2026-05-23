# ResourceRemoteInfoGithubOrgRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | **string** | The id of the role. | 
**OrgName** | Pointer to **string** | GitHub org role&#39;s org name, required only for Enterprise. | [optional] 

## Methods

### NewResourceRemoteInfoGithubOrgRole

`func NewResourceRemoteInfoGithubOrgRole(roleId string, ) *ResourceRemoteInfoGithubOrgRole`

NewResourceRemoteInfoGithubOrgRole instantiates a new ResourceRemoteInfoGithubOrgRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoGithubOrgRoleWithDefaults

`func NewResourceRemoteInfoGithubOrgRoleWithDefaults() *ResourceRemoteInfoGithubOrgRole`

NewResourceRemoteInfoGithubOrgRoleWithDefaults instantiates a new ResourceRemoteInfoGithubOrgRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *ResourceRemoteInfoGithubOrgRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ResourceRemoteInfoGithubOrgRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ResourceRemoteInfoGithubOrgRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetOrgName

`func (o *ResourceRemoteInfoGithubOrgRole) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ResourceRemoteInfoGithubOrgRole) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ResourceRemoteInfoGithubOrgRole) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ResourceRemoteInfoGithubOrgRole) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


