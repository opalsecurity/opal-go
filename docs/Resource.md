# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource. | 
**AppId** | Pointer to **string** | The ID of the app. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**Description** | Pointer to **string** | A description of the resource. | [optional] 
**AdminOwnerId** | Pointer to **string** | The ID of the owner of the resource. | [optional] 
**RemoteResourceId** | Pointer to **string** | The ID of the resource on the remote system. | [optional] 
**RemoteResourceName** | Pointer to **string** | The name of the resource on the remote system. | [optional] 
**ResourceType** | Pointer to [**ResourceTypeEnum**](ResourceTypeEnum.md) |  | [optional] 
**MaxDuration** | Pointer to **int32** | The maximum duration for which the resource can be requested (in minutes). | [optional] 
**RecommendedDuration** | Pointer to **int32** | The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration. | [optional] 
**RequireManagerApproval** | Pointer to **bool** | A bool representing whether or not access requests to the resource require manager approval. | [optional] 
**RequireSupportTicket** | Pointer to **bool** | A bool representing whether or not access requests to the resource require an access ticket. | [optional] 
**FolderId** | Pointer to **string** | The ID of the folder that the resource is located in. | [optional] 
**RequireMfaToApprove** | Pointer to **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this resource. | [optional] 
**RequireMfaToConnect** | Pointer to **bool** | A bool representing whether or not to require MFA to connect to this resource. | [optional] 
**AutoApproval** | Pointer to **bool** | A bool representing whether or not to automatically approve requests to this resource. | [optional] 
**RequestTemplateId** | Pointer to **string** | The ID of the associated request template. | [optional] 
**Metadata** | Pointer to **string** | JSON metadata about the remote resource. Only set for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details. | [optional] 

## Methods

### NewResource

`func NewResource(resourceId string, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *Resource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Resource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Resource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetAppId

`func (o *Resource) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Resource) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Resource) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Resource) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Resource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Resource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Resource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Resource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Resource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *Resource) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *Resource) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *Resource) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.

### HasAdminOwnerId

`func (o *Resource) HasAdminOwnerId() bool`

HasAdminOwnerId returns a boolean if a field has been set.

### GetRemoteResourceId

`func (o *Resource) GetRemoteResourceId() string`

GetRemoteResourceId returns the RemoteResourceId field if non-nil, zero value otherwise.

### GetRemoteResourceIdOk

`func (o *Resource) GetRemoteResourceIdOk() (*string, bool)`

GetRemoteResourceIdOk returns a tuple with the RemoteResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteResourceId

`func (o *Resource) SetRemoteResourceId(v string)`

SetRemoteResourceId sets RemoteResourceId field to given value.

### HasRemoteResourceId

`func (o *Resource) HasRemoteResourceId() bool`

HasRemoteResourceId returns a boolean if a field has been set.

### GetRemoteResourceName

`func (o *Resource) GetRemoteResourceName() string`

GetRemoteResourceName returns the RemoteResourceName field if non-nil, zero value otherwise.

### GetRemoteResourceNameOk

`func (o *Resource) GetRemoteResourceNameOk() (*string, bool)`

GetRemoteResourceNameOk returns a tuple with the RemoteResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteResourceName

`func (o *Resource) SetRemoteResourceName(v string)`

SetRemoteResourceName sets RemoteResourceName field to given value.

### HasRemoteResourceName

`func (o *Resource) HasRemoteResourceName() bool`

HasRemoteResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *Resource) GetResourceType() ResourceTypeEnum`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Resource) GetResourceTypeOk() (*ResourceTypeEnum, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Resource) SetResourceType(v ResourceTypeEnum)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Resource) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetMaxDuration

`func (o *Resource) GetMaxDuration() int32`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *Resource) GetMaxDurationOk() (*int32, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *Resource) SetMaxDuration(v int32)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *Resource) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.

### GetRecommendedDuration

`func (o *Resource) GetRecommendedDuration() int32`

GetRecommendedDuration returns the RecommendedDuration field if non-nil, zero value otherwise.

### GetRecommendedDurationOk

`func (o *Resource) GetRecommendedDurationOk() (*int32, bool)`

GetRecommendedDurationOk returns a tuple with the RecommendedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedDuration

`func (o *Resource) SetRecommendedDuration(v int32)`

SetRecommendedDuration sets RecommendedDuration field to given value.

### HasRecommendedDuration

`func (o *Resource) HasRecommendedDuration() bool`

HasRecommendedDuration returns a boolean if a field has been set.

### GetRequireManagerApproval

`func (o *Resource) GetRequireManagerApproval() bool`

GetRequireManagerApproval returns the RequireManagerApproval field if non-nil, zero value otherwise.

### GetRequireManagerApprovalOk

`func (o *Resource) GetRequireManagerApprovalOk() (*bool, bool)`

GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireManagerApproval

`func (o *Resource) SetRequireManagerApproval(v bool)`

SetRequireManagerApproval sets RequireManagerApproval field to given value.

### HasRequireManagerApproval

`func (o *Resource) HasRequireManagerApproval() bool`

HasRequireManagerApproval returns a boolean if a field has been set.

### GetRequireSupportTicket

`func (o *Resource) GetRequireSupportTicket() bool`

GetRequireSupportTicket returns the RequireSupportTicket field if non-nil, zero value otherwise.

### GetRequireSupportTicketOk

`func (o *Resource) GetRequireSupportTicketOk() (*bool, bool)`

GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSupportTicket

`func (o *Resource) SetRequireSupportTicket(v bool)`

SetRequireSupportTicket sets RequireSupportTicket field to given value.

### HasRequireSupportTicket

`func (o *Resource) HasRequireSupportTicket() bool`

HasRequireSupportTicket returns a boolean if a field has been set.

### GetFolderId

`func (o *Resource) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Resource) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Resource) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Resource) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetRequireMfaToApprove

`func (o *Resource) GetRequireMfaToApprove() bool`

GetRequireMfaToApprove returns the RequireMfaToApprove field if non-nil, zero value otherwise.

### GetRequireMfaToApproveOk

`func (o *Resource) GetRequireMfaToApproveOk() (*bool, bool)`

GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToApprove

`func (o *Resource) SetRequireMfaToApprove(v bool)`

SetRequireMfaToApprove sets RequireMfaToApprove field to given value.

### HasRequireMfaToApprove

`func (o *Resource) HasRequireMfaToApprove() bool`

HasRequireMfaToApprove returns a boolean if a field has been set.

### GetRequireMfaToConnect

`func (o *Resource) GetRequireMfaToConnect() bool`

GetRequireMfaToConnect returns the RequireMfaToConnect field if non-nil, zero value otherwise.

### GetRequireMfaToConnectOk

`func (o *Resource) GetRequireMfaToConnectOk() (*bool, bool)`

GetRequireMfaToConnectOk returns a tuple with the RequireMfaToConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToConnect

`func (o *Resource) SetRequireMfaToConnect(v bool)`

SetRequireMfaToConnect sets RequireMfaToConnect field to given value.

### HasRequireMfaToConnect

`func (o *Resource) HasRequireMfaToConnect() bool`

HasRequireMfaToConnect returns a boolean if a field has been set.

### GetAutoApproval

`func (o *Resource) GetAutoApproval() bool`

GetAutoApproval returns the AutoApproval field if non-nil, zero value otherwise.

### GetAutoApprovalOk

`func (o *Resource) GetAutoApprovalOk() (*bool, bool)`

GetAutoApprovalOk returns a tuple with the AutoApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproval

`func (o *Resource) SetAutoApproval(v bool)`

SetAutoApproval sets AutoApproval field to given value.

### HasAutoApproval

`func (o *Resource) HasAutoApproval() bool`

HasAutoApproval returns a boolean if a field has been set.

### GetRequestTemplateId

`func (o *Resource) GetRequestTemplateId() string`

GetRequestTemplateId returns the RequestTemplateId field if non-nil, zero value otherwise.

### GetRequestTemplateIdOk

`func (o *Resource) GetRequestTemplateIdOk() (*string, bool)`

GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplateId

`func (o *Resource) SetRequestTemplateId(v string)`

SetRequestTemplateId sets RequestTemplateId field to given value.

### HasRequestTemplateId

`func (o *Resource) HasRequestTemplateId() bool`

HasRequestTemplateId returns a boolean if a field has been set.

### GetMetadata

`func (o *Resource) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Resource) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Resource) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Resource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


