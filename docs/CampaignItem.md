# CampaignItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignItemId** | **string** | The ID of the campaign item. | 
**CampaignId** | **string** | The ID of the parent campaign. | 
**TargetType** | [**CampaignItemTargetTypeEnum**](CampaignItemTargetTypeEnum.md) |  | 
**RoleAssignmentId** | **string** | The role assignment ID this item reviews (&#x60;target_id&#x60; when &#x60;target_type&#x60; is ROLE_ASSIGNMENT).  | 
**Status** | [**CampaignItemStatusEnum**](CampaignItemStatusEnum.md) |  | 
**CreatedAt** | **time.Time** | When the item was created (snapshotted into the campaign). | 
**IsTargetDeleted** | **bool** | True when the underlying role assignment is soft-deleted. | 
**PrincipalId** | Pointer to **string** | Principal ID from the role assignment. Null if the RA is missing. | [optional] 
**PrincipalType** | Pointer to [**EntityTypeEnum**](EntityTypeEnum.md) | Principal entity type. Null if the RA is missing. | [optional] 
**EntityId** | Pointer to **string** | Entitlement entity ID. Null if the RA is missing. | [optional] 
**EntityType** | Pointer to [**EntityTypeEnum**](EntityTypeEnum.md) | Entitlement entity type. Null if the RA is missing. | [optional] 
**AccessLevel** | Pointer to [**ResourceAccessLevel**](ResourceAccessLevel.md) | Access level on the role assignment, if any. | [optional] 
**AccessLevelName** | Pointer to **string** | Denormalized access level name stored on the campaign item. | [optional] 
**Reviews** | [**[]CampaignItemReview**](CampaignItemReview.md) | Reviewer assignments and decisions for this item. | 
**AdminOverride** | Pointer to [**CampaignItemAdminOverride**](CampaignItemAdminOverride.md) | Admin override, if any. | [optional] 

## Methods

### NewCampaignItem

`func NewCampaignItem(campaignItemId string, campaignId string, targetType CampaignItemTargetTypeEnum, roleAssignmentId string, status CampaignItemStatusEnum, createdAt time.Time, isTargetDeleted bool, reviews []CampaignItemReview, ) *CampaignItem`

NewCampaignItem instantiates a new CampaignItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignItemWithDefaults

`func NewCampaignItemWithDefaults() *CampaignItem`

NewCampaignItemWithDefaults instantiates a new CampaignItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignItemId

`func (o *CampaignItem) GetCampaignItemId() string`

GetCampaignItemId returns the CampaignItemId field if non-nil, zero value otherwise.

### GetCampaignItemIdOk

`func (o *CampaignItem) GetCampaignItemIdOk() (*string, bool)`

GetCampaignItemIdOk returns a tuple with the CampaignItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignItemId

`func (o *CampaignItem) SetCampaignItemId(v string)`

SetCampaignItemId sets CampaignItemId field to given value.


### GetCampaignId

`func (o *CampaignItem) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignItem) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignItem) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetTargetType

`func (o *CampaignItem) GetTargetType() CampaignItemTargetTypeEnum`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CampaignItem) GetTargetTypeOk() (*CampaignItemTargetTypeEnum, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CampaignItem) SetTargetType(v CampaignItemTargetTypeEnum)`

SetTargetType sets TargetType field to given value.


### GetRoleAssignmentId

`func (o *CampaignItem) GetRoleAssignmentId() string`

GetRoleAssignmentId returns the RoleAssignmentId field if non-nil, zero value otherwise.

### GetRoleAssignmentIdOk

`func (o *CampaignItem) GetRoleAssignmentIdOk() (*string, bool)`

GetRoleAssignmentIdOk returns a tuple with the RoleAssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignmentId

`func (o *CampaignItem) SetRoleAssignmentId(v string)`

SetRoleAssignmentId sets RoleAssignmentId field to given value.


### GetStatus

`func (o *CampaignItem) GetStatus() CampaignItemStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignItem) GetStatusOk() (*CampaignItemStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignItem) SetStatus(v CampaignItemStatusEnum)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *CampaignItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsTargetDeleted

`func (o *CampaignItem) GetIsTargetDeleted() bool`

GetIsTargetDeleted returns the IsTargetDeleted field if non-nil, zero value otherwise.

### GetIsTargetDeletedOk

`func (o *CampaignItem) GetIsTargetDeletedOk() (*bool, bool)`

GetIsTargetDeletedOk returns a tuple with the IsTargetDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTargetDeleted

`func (o *CampaignItem) SetIsTargetDeleted(v bool)`

SetIsTargetDeleted sets IsTargetDeleted field to given value.


### GetPrincipalId

`func (o *CampaignItem) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *CampaignItem) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *CampaignItem) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *CampaignItem) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### GetPrincipalType

`func (o *CampaignItem) GetPrincipalType() EntityTypeEnum`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *CampaignItem) GetPrincipalTypeOk() (*EntityTypeEnum, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *CampaignItem) SetPrincipalType(v EntityTypeEnum)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *CampaignItem) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.

### GetEntityId

`func (o *CampaignItem) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *CampaignItem) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *CampaignItem) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *CampaignItem) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *CampaignItem) GetEntityType() EntityTypeEnum`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CampaignItem) GetEntityTypeOk() (*EntityTypeEnum, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CampaignItem) SetEntityType(v EntityTypeEnum)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CampaignItem) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetAccessLevel

`func (o *CampaignItem) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *CampaignItem) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *CampaignItem) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *CampaignItem) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetAccessLevelName

`func (o *CampaignItem) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *CampaignItem) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *CampaignItem) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *CampaignItem) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.

### GetReviews

`func (o *CampaignItem) GetReviews() []CampaignItemReview`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *CampaignItem) GetReviewsOk() (*[]CampaignItemReview, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *CampaignItem) SetReviews(v []CampaignItemReview)`

SetReviews sets Reviews field to given value.


### GetAdminOverride

`func (o *CampaignItem) GetAdminOverride() CampaignItemAdminOverride`

GetAdminOverride returns the AdminOverride field if non-nil, zero value otherwise.

### GetAdminOverrideOk

`func (o *CampaignItem) GetAdminOverrideOk() (*CampaignItemAdminOverride, bool)`

GetAdminOverrideOk returns a tuple with the AdminOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOverride

`func (o *CampaignItem) SetAdminOverride(v CampaignItemAdminOverride)`

SetAdminOverride sets AdminOverride field to given value.

### HasAdminOverride

`func (o *CampaignItem) HasAdminOverride() bool`

HasAdminOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


