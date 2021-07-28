# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | The ID of the connection. | 
**UserId** | **string** | The ID of the user. | 
**ResourceId** | **string** | The ID of the resource. | 
**AccessLevel** | [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | 
**ExpirationDate** | **time.Time** | The day and time the user&#39;s access will expire. | 
**IsGlobal** | **bool** | Whether the session is of type global impersonation. | 

## Methods

### NewSession

`func NewSession(connectionId string, userId string, resourceId string, accessLevel ResourceAccessLevel, expirationDate time.Time, isGlobal bool, ) *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *Session) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Session) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Session) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetUserId

`func (o *Session) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Session) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Session) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetResourceId

`func (o *Session) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Session) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Session) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetAccessLevel

`func (o *Session) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *Session) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *Session) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.


### GetExpirationDate

`func (o *Session) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Session) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Session) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### GetIsGlobal

`func (o *Session) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *Session) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *Session) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


