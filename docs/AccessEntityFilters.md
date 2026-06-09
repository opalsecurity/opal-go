# AccessEntityFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityTypes** | Pointer to **[]string** | Filter by entity type. Only RESOURCE, GROUP, and USER are queryable via OpalQuery. | [optional] 
**EntityItemTypes** | Pointer to [**[]EntityItemTypeEnum**](EntityItemTypeEnum.md) | Filter by entity item types. | [optional] 
**EntityName** | Pointer to [**EntityNameFilter**](EntityNameFilter.md) |  | [optional] 
**EntityTag** | Pointer to [**EntityTagFilter**](EntityTagFilter.md) |  | [optional] 
**EntityIDs** | Pointer to **[]string** | Filter by specific entity UUIDs. | [optional] 
**ImportedFromApp** | Pointer to **[]string** | Filter by app IDs from which returned nodes will be imported from. | [optional] 
**RoleRemoteIds** | Pointer to **[]string** | Filter by role remote IDs. Can only be applied within a hasAccessTo clause. | [optional] 
**RoleNames** | Pointer to **[]string** | Filter by role display names (e.g. \&quot;Admin\&quot;, \&quot;Read\&quot;). Can only be applied within a hasAccessTo clause. | [optional] 
**AllOf** | Pointer to [**[]AccessEntityFilters**](AccessEntityFilters.md) | A list of nested filters that must all match (logical AND). Each  item has the same shape as this object — scalar fields like  &#x60;entityTypes&#x60; or &#x60;entityTag&#x60;, and can further nest &#x60;allOf&#x60;,  &#x60;anyOf&#x60;, or &#x60;not&#x60;.  | [optional] 
**AnyOf** | Pointer to [**[]AccessEntityFilters**](AccessEntityFilters.md) | A list of nested filters where at least one must match (logical  OR). Each item has the same shape as this object.  | [optional] 
**Not** | Pointer to **map[string]interface{}** | Excludes entities matching the embedded filter (logical NOT). Pass a filter object with the same shape as this one — typically a single scalar field, like &#x60;{not: {entityTypes: [\&quot;RESOURCE\&quot;]}}&#x60; to exclude resources.  | [optional] 

## Methods

### NewAccessEntityFilters

`func NewAccessEntityFilters() *AccessEntityFilters`

NewAccessEntityFilters instantiates a new AccessEntityFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessEntityFiltersWithDefaults

`func NewAccessEntityFiltersWithDefaults() *AccessEntityFilters`

NewAccessEntityFiltersWithDefaults instantiates a new AccessEntityFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityTypes

`func (o *AccessEntityFilters) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *AccessEntityFilters) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *AccessEntityFilters) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *AccessEntityFilters) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### GetEntityItemTypes

`func (o *AccessEntityFilters) GetEntityItemTypes() []EntityItemTypeEnum`

GetEntityItemTypes returns the EntityItemTypes field if non-nil, zero value otherwise.

### GetEntityItemTypesOk

`func (o *AccessEntityFilters) GetEntityItemTypesOk() (*[]EntityItemTypeEnum, bool)`

GetEntityItemTypesOk returns a tuple with the EntityItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityItemTypes

`func (o *AccessEntityFilters) SetEntityItemTypes(v []EntityItemTypeEnum)`

SetEntityItemTypes sets EntityItemTypes field to given value.

### HasEntityItemTypes

`func (o *AccessEntityFilters) HasEntityItemTypes() bool`

HasEntityItemTypes returns a boolean if a field has been set.

### GetEntityName

`func (o *AccessEntityFilters) GetEntityName() EntityNameFilter`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *AccessEntityFilters) GetEntityNameOk() (*EntityNameFilter, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *AccessEntityFilters) SetEntityName(v EntityNameFilter)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *AccessEntityFilters) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetEntityTag

`func (o *AccessEntityFilters) GetEntityTag() EntityTagFilter`

GetEntityTag returns the EntityTag field if non-nil, zero value otherwise.

### GetEntityTagOk

`func (o *AccessEntityFilters) GetEntityTagOk() (*EntityTagFilter, bool)`

GetEntityTagOk returns a tuple with the EntityTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTag

`func (o *AccessEntityFilters) SetEntityTag(v EntityTagFilter)`

SetEntityTag sets EntityTag field to given value.

### HasEntityTag

`func (o *AccessEntityFilters) HasEntityTag() bool`

HasEntityTag returns a boolean if a field has been set.

### GetEntityIDs

`func (o *AccessEntityFilters) GetEntityIDs() []string`

GetEntityIDs returns the EntityIDs field if non-nil, zero value otherwise.

### GetEntityIDsOk

`func (o *AccessEntityFilters) GetEntityIDsOk() (*[]string, bool)`

GetEntityIDsOk returns a tuple with the EntityIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIDs

`func (o *AccessEntityFilters) SetEntityIDs(v []string)`

SetEntityIDs sets EntityIDs field to given value.

### HasEntityIDs

`func (o *AccessEntityFilters) HasEntityIDs() bool`

HasEntityIDs returns a boolean if a field has been set.

### GetImportedFromApp

`func (o *AccessEntityFilters) GetImportedFromApp() []string`

GetImportedFromApp returns the ImportedFromApp field if non-nil, zero value otherwise.

### GetImportedFromAppOk

`func (o *AccessEntityFilters) GetImportedFromAppOk() (*[]string, bool)`

GetImportedFromAppOk returns a tuple with the ImportedFromApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFromApp

`func (o *AccessEntityFilters) SetImportedFromApp(v []string)`

SetImportedFromApp sets ImportedFromApp field to given value.

### HasImportedFromApp

`func (o *AccessEntityFilters) HasImportedFromApp() bool`

HasImportedFromApp returns a boolean if a field has been set.

### GetRoleRemoteIds

`func (o *AccessEntityFilters) GetRoleRemoteIds() []string`

GetRoleRemoteIds returns the RoleRemoteIds field if non-nil, zero value otherwise.

### GetRoleRemoteIdsOk

`func (o *AccessEntityFilters) GetRoleRemoteIdsOk() (*[]string, bool)`

GetRoleRemoteIdsOk returns a tuple with the RoleRemoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleRemoteIds

`func (o *AccessEntityFilters) SetRoleRemoteIds(v []string)`

SetRoleRemoteIds sets RoleRemoteIds field to given value.

### HasRoleRemoteIds

`func (o *AccessEntityFilters) HasRoleRemoteIds() bool`

HasRoleRemoteIds returns a boolean if a field has been set.

### GetRoleNames

`func (o *AccessEntityFilters) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *AccessEntityFilters) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *AccessEntityFilters) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *AccessEntityFilters) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.

### GetAllOf

`func (o *AccessEntityFilters) GetAllOf() []AccessEntityFilters`

GetAllOf returns the AllOf field if non-nil, zero value otherwise.

### GetAllOfOk

`func (o *AccessEntityFilters) GetAllOfOk() (*[]AccessEntityFilters, bool)`

GetAllOfOk returns a tuple with the AllOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOf

`func (o *AccessEntityFilters) SetAllOf(v []AccessEntityFilters)`

SetAllOf sets AllOf field to given value.

### HasAllOf

`func (o *AccessEntityFilters) HasAllOf() bool`

HasAllOf returns a boolean if a field has been set.

### GetAnyOf

`func (o *AccessEntityFilters) GetAnyOf() []AccessEntityFilters`

GetAnyOf returns the AnyOf field if non-nil, zero value otherwise.

### GetAnyOfOk

`func (o *AccessEntityFilters) GetAnyOfOk() (*[]AccessEntityFilters, bool)`

GetAnyOfOk returns a tuple with the AnyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyOf

`func (o *AccessEntityFilters) SetAnyOf(v []AccessEntityFilters)`

SetAnyOf sets AnyOf field to given value.

### HasAnyOf

`func (o *AccessEntityFilters) HasAnyOf() bool`

HasAnyOf returns a boolean if a field has been set.

### GetNot

`func (o *AccessEntityFilters) GetNot() map[string]interface{}`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *AccessEntityFilters) GetNotOk() (*map[string]interface{}, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *AccessEntityFilters) SetNot(v map[string]interface{})`

SetNot sets Not field to given value.

### HasNot

`func (o *AccessEntityFilters) HasNot() bool`

HasNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


