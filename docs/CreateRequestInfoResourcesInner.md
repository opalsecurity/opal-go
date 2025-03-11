# CreateRequestInfoResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the resource requested. Should not be specified if group_id is specified. | [optional] 
**AccessLevelRemoteId** | Pointer to **string** | The ID of the access level requested on the remote system. | [optional] 
**AccessLevelName** | Pointer to **string** | The ID of the access level requested on the remote system. | [optional] 

## Methods

### NewCreateRequestInfoResourcesInner

`func NewCreateRequestInfoResourcesInner() *CreateRequestInfoResourcesInner`

NewCreateRequestInfoResourcesInner instantiates a new CreateRequestInfoResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestInfoResourcesInnerWithDefaults

`func NewCreateRequestInfoResourcesInnerWithDefaults() *CreateRequestInfoResourcesInner`

NewCreateRequestInfoResourcesInnerWithDefaults instantiates a new CreateRequestInfoResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRequestInfoResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRequestInfoResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRequestInfoResourcesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateRequestInfoResourcesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessLevelRemoteId

`func (o *CreateRequestInfoResourcesInner) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *CreateRequestInfoResourcesInner) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *CreateRequestInfoResourcesInner) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *CreateRequestInfoResourcesInner) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetAccessLevelName

`func (o *CreateRequestInfoResourcesInner) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *CreateRequestInfoResourcesInner) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *CreateRequestInfoResourcesInner) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *CreateRequestInfoResourcesInner) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


