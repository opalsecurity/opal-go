# CreateRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | [**[]CreateRequestInfoResourcesInner**](CreateRequestInfoResourcesInner.md) |  | 
**Groups** | [**[]CreateRequestInfoGroupsInner**](CreateRequestInfoGroupsInner.md) |  | 
**TargetUserId** | Pointer to **string** | The ID of the user to be granted access. Should not be specified if target_group_id is specified. | [optional] 
**TargetGroupId** | Pointer to **string** | The ID of the group the request is for.  Should not be specified if target_user_id is specified. | [optional] 
**Reason** | **string** |  | 
**SupportTicket** | Pointer to [**CreateRequestInfoSupportTicket**](CreateRequestInfoSupportTicket.md) |  | [optional] 
**DurationMinutes** | **int32** | The duration of the request in minutes. -1 represents an indefinite duration | 
**CustomMetadata** | Pointer to [**[]CreateRequestInfoCustomMetadataInner**](CreateRequestInfoCustomMetadataInner.md) |  | [optional] 

## Methods

### NewCreateRequestInfo

`func NewCreateRequestInfo(resources []CreateRequestInfoResourcesInner, groups []CreateRequestInfoGroupsInner, reason string, durationMinutes int32, ) *CreateRequestInfo`

NewCreateRequestInfo instantiates a new CreateRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestInfoWithDefaults

`func NewCreateRequestInfoWithDefaults() *CreateRequestInfo`

NewCreateRequestInfoWithDefaults instantiates a new CreateRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *CreateRequestInfo) GetResources() []CreateRequestInfoResourcesInner`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CreateRequestInfo) GetResourcesOk() (*[]CreateRequestInfoResourcesInner, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CreateRequestInfo) SetResources(v []CreateRequestInfoResourcesInner)`

SetResources sets Resources field to given value.


### GetGroups

`func (o *CreateRequestInfo) GetGroups() []CreateRequestInfoGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateRequestInfo) GetGroupsOk() (*[]CreateRequestInfoGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateRequestInfo) SetGroups(v []CreateRequestInfoGroupsInner)`

SetGroups sets Groups field to given value.


### GetTargetUserId

`func (o *CreateRequestInfo) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *CreateRequestInfo) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *CreateRequestInfo) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.

### HasTargetUserId

`func (o *CreateRequestInfo) HasTargetUserId() bool`

HasTargetUserId returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *CreateRequestInfo) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *CreateRequestInfo) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *CreateRequestInfo) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *CreateRequestInfo) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetReason

`func (o *CreateRequestInfo) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateRequestInfo) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateRequestInfo) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSupportTicket

`func (o *CreateRequestInfo) GetSupportTicket() CreateRequestInfoSupportTicket`

GetSupportTicket returns the SupportTicket field if non-nil, zero value otherwise.

### GetSupportTicketOk

`func (o *CreateRequestInfo) GetSupportTicketOk() (*CreateRequestInfoSupportTicket, bool)`

GetSupportTicketOk returns a tuple with the SupportTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicket

`func (o *CreateRequestInfo) SetSupportTicket(v CreateRequestInfoSupportTicket)`

SetSupportTicket sets SupportTicket field to given value.

### HasSupportTicket

`func (o *CreateRequestInfo) HasSupportTicket() bool`

HasSupportTicket returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *CreateRequestInfo) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *CreateRequestInfo) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *CreateRequestInfo) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.


### GetCustomMetadata

`func (o *CreateRequestInfo) GetCustomMetadata() []CreateRequestInfoCustomMetadataInner`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *CreateRequestInfo) GetCustomMetadataOk() (*[]CreateRequestInfoCustomMetadataInner, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *CreateRequestInfo) SetCustomMetadata(v []CreateRequestInfoCustomMetadataInner)`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *CreateRequestInfo) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


