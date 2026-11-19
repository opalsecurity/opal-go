# CreateTokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenLabel** | **string** | A human-readable label for the token. | 
**AccessLevel** | [**ApiAccessLevelEnum**](ApiAccessLevelEnum.md) |  | 
**UserId** | Pointer to **string** | The ID of the user the token should authenticate as. Omit or set to the caller&#39;s ID to create a personal token. Set to a service user&#39;s ID to mint a token for that service user.  | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Optional expiration time for the token. | [optional] 

## Methods

### NewCreateTokenInfo

`func NewCreateTokenInfo(tokenLabel string, accessLevel ApiAccessLevelEnum, ) *CreateTokenInfo`

NewCreateTokenInfo instantiates a new CreateTokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenInfoWithDefaults

`func NewCreateTokenInfoWithDefaults() *CreateTokenInfo`

NewCreateTokenInfoWithDefaults instantiates a new CreateTokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenLabel

`func (o *CreateTokenInfo) GetTokenLabel() string`

GetTokenLabel returns the TokenLabel field if non-nil, zero value otherwise.

### GetTokenLabelOk

`func (o *CreateTokenInfo) GetTokenLabelOk() (*string, bool)`

GetTokenLabelOk returns a tuple with the TokenLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLabel

`func (o *CreateTokenInfo) SetTokenLabel(v string)`

SetTokenLabel sets TokenLabel field to given value.


### GetAccessLevel

`func (o *CreateTokenInfo) GetAccessLevel() ApiAccessLevelEnum`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *CreateTokenInfo) GetAccessLevelOk() (*ApiAccessLevelEnum, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *CreateTokenInfo) SetAccessLevel(v ApiAccessLevelEnum)`

SetAccessLevel sets AccessLevel field to given value.


### GetUserId

`func (o *CreateTokenInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateTokenInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateTokenInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateTokenInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateTokenInfo) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateTokenInfo) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateTokenInfo) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateTokenInfo) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


