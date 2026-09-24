# ViewerCampaignItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignItemReviewId** | **string** | The ID of the campaign item review. | 
**Decision** | Pointer to [**ViewerCampaignItemReviewDecisionEnum**](ViewerCampaignItemReviewDecisionEnum.md) | The reviewer&#39;s decision. Null when undecided. | [optional] 
**Note** | Pointer to **string** | Optional note from the reviewer. | [optional] 
**DecidedAt** | Pointer to **time.Time** | When the decision was finalized. | [optional] 
**PendingDecision** | Pointer to [**ViewerCampaignItemReviewDecisionEnum**](ViewerCampaignItemReviewDecisionEnum.md) | Pending decision staged by the viewer, if any. | [optional] 
**PendingNote** | Pointer to **string** | Pending note attached to the reviewer&#39;s pending decision. | [optional] 
**UpdatedAccessLevelRemoteId** | Pointer to **string** | Target access level remote ID when decision is CHANGE_ROLE. | [optional] 
**PendingUpdatedAccessLevelRemoteId** | Pointer to **string** | Pending target access level remote ID when decision is CHANGE_ROLE. | [optional] 
**ReassignedToReviewerIds** | Pointer to **[]string** | User IDs this review was reassigned to. | [optional] 
**PrincipalId** | **string** | Principal ID from the role assignment. | 
**PrincipalType** | [**EntityTypeEnum**](EntityTypeEnum.md) |  | 
**EntityId** | **string** | Entity ID from the role assignment. | 
**EntityType** | [**EntityTypeEnum**](EntityTypeEnum.md) |  | 
**AccessLevelName** | Pointer to **string** | Denormalized access level name. | [optional] 
**AccessLevelRemoteId** | Pointer to **string** | Access level remote ID from the role assignment. | [optional] 
**GrantedAt** | **time.Time** | When the access was originally granted. | 
**ExpiresAt** | Pointer to **time.Time** | When the access expires, if set. | [optional] 
**RoleAssignmentId** | **string** | ID of the underlying role assignment. | 

## Methods

### NewViewerCampaignItem

`func NewViewerCampaignItem(campaignItemReviewId string, principalId string, principalType EntityTypeEnum, entityId string, entityType EntityTypeEnum, grantedAt time.Time, roleAssignmentId string, ) *ViewerCampaignItem`

NewViewerCampaignItem instantiates a new ViewerCampaignItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewerCampaignItemWithDefaults

`func NewViewerCampaignItemWithDefaults() *ViewerCampaignItem`

NewViewerCampaignItemWithDefaults instantiates a new ViewerCampaignItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignItemReviewId

`func (o *ViewerCampaignItem) GetCampaignItemReviewId() string`

GetCampaignItemReviewId returns the CampaignItemReviewId field if non-nil, zero value otherwise.

### GetCampaignItemReviewIdOk

`func (o *ViewerCampaignItem) GetCampaignItemReviewIdOk() (*string, bool)`

GetCampaignItemReviewIdOk returns a tuple with the CampaignItemReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignItemReviewId

`func (o *ViewerCampaignItem) SetCampaignItemReviewId(v string)`

SetCampaignItemReviewId sets CampaignItemReviewId field to given value.


### GetDecision

`func (o *ViewerCampaignItem) GetDecision() ViewerCampaignItemReviewDecisionEnum`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ViewerCampaignItem) GetDecisionOk() (*ViewerCampaignItemReviewDecisionEnum, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ViewerCampaignItem) SetDecision(v ViewerCampaignItemReviewDecisionEnum)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *ViewerCampaignItem) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetNote

`func (o *ViewerCampaignItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ViewerCampaignItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ViewerCampaignItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ViewerCampaignItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetDecidedAt

`func (o *ViewerCampaignItem) GetDecidedAt() time.Time`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *ViewerCampaignItem) GetDecidedAtOk() (*time.Time, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *ViewerCampaignItem) SetDecidedAt(v time.Time)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *ViewerCampaignItem) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetPendingDecision

`func (o *ViewerCampaignItem) GetPendingDecision() ViewerCampaignItemReviewDecisionEnum`

GetPendingDecision returns the PendingDecision field if non-nil, zero value otherwise.

### GetPendingDecisionOk

`func (o *ViewerCampaignItem) GetPendingDecisionOk() (*ViewerCampaignItemReviewDecisionEnum, bool)`

GetPendingDecisionOk returns a tuple with the PendingDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDecision

`func (o *ViewerCampaignItem) SetPendingDecision(v ViewerCampaignItemReviewDecisionEnum)`

SetPendingDecision sets PendingDecision field to given value.

### HasPendingDecision

`func (o *ViewerCampaignItem) HasPendingDecision() bool`

HasPendingDecision returns a boolean if a field has been set.

### GetPendingNote

`func (o *ViewerCampaignItem) GetPendingNote() string`

GetPendingNote returns the PendingNote field if non-nil, zero value otherwise.

### GetPendingNoteOk

`func (o *ViewerCampaignItem) GetPendingNoteOk() (*string, bool)`

GetPendingNoteOk returns a tuple with the PendingNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingNote

`func (o *ViewerCampaignItem) SetPendingNote(v string)`

SetPendingNote sets PendingNote field to given value.

### HasPendingNote

`func (o *ViewerCampaignItem) HasPendingNote() bool`

HasPendingNote returns a boolean if a field has been set.

### GetUpdatedAccessLevelRemoteId

`func (o *ViewerCampaignItem) GetUpdatedAccessLevelRemoteId() string`

GetUpdatedAccessLevelRemoteId returns the UpdatedAccessLevelRemoteId field if non-nil, zero value otherwise.

### GetUpdatedAccessLevelRemoteIdOk

`func (o *ViewerCampaignItem) GetUpdatedAccessLevelRemoteIdOk() (*string, bool)`

GetUpdatedAccessLevelRemoteIdOk returns a tuple with the UpdatedAccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAccessLevelRemoteId

`func (o *ViewerCampaignItem) SetUpdatedAccessLevelRemoteId(v string)`

SetUpdatedAccessLevelRemoteId sets UpdatedAccessLevelRemoteId field to given value.

### HasUpdatedAccessLevelRemoteId

`func (o *ViewerCampaignItem) HasUpdatedAccessLevelRemoteId() bool`

HasUpdatedAccessLevelRemoteId returns a boolean if a field has been set.

### GetPendingUpdatedAccessLevelRemoteId

`func (o *ViewerCampaignItem) GetPendingUpdatedAccessLevelRemoteId() string`

GetPendingUpdatedAccessLevelRemoteId returns the PendingUpdatedAccessLevelRemoteId field if non-nil, zero value otherwise.

### GetPendingUpdatedAccessLevelRemoteIdOk

`func (o *ViewerCampaignItem) GetPendingUpdatedAccessLevelRemoteIdOk() (*string, bool)`

GetPendingUpdatedAccessLevelRemoteIdOk returns a tuple with the PendingUpdatedAccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdatedAccessLevelRemoteId

`func (o *ViewerCampaignItem) SetPendingUpdatedAccessLevelRemoteId(v string)`

SetPendingUpdatedAccessLevelRemoteId sets PendingUpdatedAccessLevelRemoteId field to given value.

### HasPendingUpdatedAccessLevelRemoteId

`func (o *ViewerCampaignItem) HasPendingUpdatedAccessLevelRemoteId() bool`

HasPendingUpdatedAccessLevelRemoteId returns a boolean if a field has been set.

### GetReassignedToReviewerIds

`func (o *ViewerCampaignItem) GetReassignedToReviewerIds() []string`

GetReassignedToReviewerIds returns the ReassignedToReviewerIds field if non-nil, zero value otherwise.

### GetReassignedToReviewerIdsOk

`func (o *ViewerCampaignItem) GetReassignedToReviewerIdsOk() (*[]string, bool)`

GetReassignedToReviewerIdsOk returns a tuple with the ReassignedToReviewerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignedToReviewerIds

`func (o *ViewerCampaignItem) SetReassignedToReviewerIds(v []string)`

SetReassignedToReviewerIds sets ReassignedToReviewerIds field to given value.

### HasReassignedToReviewerIds

`func (o *ViewerCampaignItem) HasReassignedToReviewerIds() bool`

HasReassignedToReviewerIds returns a boolean if a field has been set.

### GetPrincipalId

`func (o *ViewerCampaignItem) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *ViewerCampaignItem) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *ViewerCampaignItem) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPrincipalType

`func (o *ViewerCampaignItem) GetPrincipalType() EntityTypeEnum`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *ViewerCampaignItem) GetPrincipalTypeOk() (*EntityTypeEnum, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *ViewerCampaignItem) SetPrincipalType(v EntityTypeEnum)`

SetPrincipalType sets PrincipalType field to given value.


### GetEntityId

`func (o *ViewerCampaignItem) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ViewerCampaignItem) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ViewerCampaignItem) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *ViewerCampaignItem) GetEntityType() EntityTypeEnum`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ViewerCampaignItem) GetEntityTypeOk() (*EntityTypeEnum, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ViewerCampaignItem) SetEntityType(v EntityTypeEnum)`

SetEntityType sets EntityType field to given value.


### GetAccessLevelName

`func (o *ViewerCampaignItem) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *ViewerCampaignItem) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *ViewerCampaignItem) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *ViewerCampaignItem) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.

### GetAccessLevelRemoteId

`func (o *ViewerCampaignItem) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *ViewerCampaignItem) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *ViewerCampaignItem) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *ViewerCampaignItem) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetGrantedAt

`func (o *ViewerCampaignItem) GetGrantedAt() time.Time`

GetGrantedAt returns the GrantedAt field if non-nil, zero value otherwise.

### GetGrantedAtOk

`func (o *ViewerCampaignItem) GetGrantedAtOk() (*time.Time, bool)`

GetGrantedAtOk returns a tuple with the GrantedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedAt

`func (o *ViewerCampaignItem) SetGrantedAt(v time.Time)`

SetGrantedAt sets GrantedAt field to given value.


### GetExpiresAt

`func (o *ViewerCampaignItem) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ViewerCampaignItem) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ViewerCampaignItem) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ViewerCampaignItem) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRoleAssignmentId

`func (o *ViewerCampaignItem) GetRoleAssignmentId() string`

GetRoleAssignmentId returns the RoleAssignmentId field if non-nil, zero value otherwise.

### GetRoleAssignmentIdOk

`func (o *ViewerCampaignItem) GetRoleAssignmentIdOk() (*string, bool)`

GetRoleAssignmentIdOk returns a tuple with the RoleAssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignmentId

`func (o *ViewerCampaignItem) SetRoleAssignmentId(v string)`

SetRoleAssignmentId sets RoleAssignmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


