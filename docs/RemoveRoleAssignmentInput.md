# RemoveRoleAssignmentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | The ID of the entity (resource or group) to revoke access from. | 
**EntityType** | [**EntityTypeEnum**](EntityTypeEnum.md) |  | 
**PrincipalId** | **string** | The ID of the principal (user, group, or resource) losing access. | 
**PrincipalType** | [**EntityTypeEnum**](EntityTypeEnum.md) |  | 
**AccessLevelRemoteId** | Pointer to **string** | Remote ID of the specific access-level grant to revoke. The lookup is exact-match on this value against stored grants.  Required when the entity has access levels defined (the server returns 422 if omitted or sent as &#x60;null&#x60;/empty for such an entity). Must be omitted or &#x60;null&#x60; for entities with no access-level catalog (group memberships and resources that have no defined access levels); supplying a value in that case also returns 422. These rules mirror those of &#x60;POST /v1/role-assignments&#x60;.  A specified value that does not match any stored grant for this &#x60;(entity_id, principal_id)&#x60; pair is a successful idempotent no-op: HTTP 200 is returned and the entry is absent from &#x60;entries&#x60;.  | [optional] 

## Methods

### NewRemoveRoleAssignmentInput

`func NewRemoveRoleAssignmentInput(entityId string, entityType EntityTypeEnum, principalId string, principalType EntityTypeEnum, ) *RemoveRoleAssignmentInput`

NewRemoveRoleAssignmentInput instantiates a new RemoveRoleAssignmentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveRoleAssignmentInputWithDefaults

`func NewRemoveRoleAssignmentInputWithDefaults() *RemoveRoleAssignmentInput`

NewRemoveRoleAssignmentInputWithDefaults instantiates a new RemoveRoleAssignmentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *RemoveRoleAssignmentInput) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *RemoveRoleAssignmentInput) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *RemoveRoleAssignmentInput) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *RemoveRoleAssignmentInput) GetEntityType() EntityTypeEnum`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *RemoveRoleAssignmentInput) GetEntityTypeOk() (*EntityTypeEnum, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *RemoveRoleAssignmentInput) SetEntityType(v EntityTypeEnum)`

SetEntityType sets EntityType field to given value.


### GetPrincipalId

`func (o *RemoveRoleAssignmentInput) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *RemoveRoleAssignmentInput) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *RemoveRoleAssignmentInput) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPrincipalType

`func (o *RemoveRoleAssignmentInput) GetPrincipalType() EntityTypeEnum`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *RemoveRoleAssignmentInput) GetPrincipalTypeOk() (*EntityTypeEnum, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *RemoveRoleAssignmentInput) SetPrincipalType(v EntityTypeEnum)`

SetPrincipalType sets PrincipalType field to given value.


### GetAccessLevelRemoteId

`func (o *RemoveRoleAssignmentInput) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *RemoveRoleAssignmentInput) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *RemoveRoleAssignmentInput) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *RemoveRoleAssignmentInput) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


