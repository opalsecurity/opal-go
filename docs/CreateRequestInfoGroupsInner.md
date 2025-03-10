# CreateRequestInfoGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the group requested. Should not be specified if resource_id is specified. | 
**AccessLevelRemoteId** | Pointer to **string** | The ID of the access level requested on the remote system. | [optional] 
**AccessLevelName** | Pointer to **string** | The ID of the access level requested on the remote system. | [optional] 

## Methods

### NewCreateRequestInfoGroupsInner

`func NewCreateRequestInfoGroupsInner(id string, ) *CreateRequestInfoGroupsInner`

NewCreateRequestInfoGroupsInner instantiates a new CreateRequestInfoGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestInfoGroupsInnerWithDefaults

`func NewCreateRequestInfoGroupsInnerWithDefaults() *CreateRequestInfoGroupsInner`

NewCreateRequestInfoGroupsInnerWithDefaults instantiates a new CreateRequestInfoGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRequestInfoGroupsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRequestInfoGroupsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRequestInfoGroupsInner) SetId(v string)`

SetId sets Id field to given value.


### GetAccessLevelRemoteId

`func (o *CreateRequestInfoGroupsInner) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *CreateRequestInfoGroupsInner) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *CreateRequestInfoGroupsInner) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *CreateRequestInfoGroupsInner) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetAccessLevelName

`func (o *CreateRequestInfoGroupsInner) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *CreateRequestInfoGroupsInner) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *CreateRequestInfoGroupsInner) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *CreateRequestInfoGroupsInner) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


