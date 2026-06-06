# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The ID of the API token. | 
**CreatedAt** | **time.Time** | The date and time the token was created. | 
**TokenLabel** | **string** | A human-readable label for the token. | 
**CreatorUserId** | **string** | The ID of the user who created the token. | 
**UserId** | **string** | The ID of the user the token authenticates as. | 
**LastUsedAt** | Pointer to **time.Time** | The date and time the token was last used. | [optional] 
**AccessLevel** | [**ApiAccessLevelEnum**](ApiAccessLevelEnum.md) |  | 
**ExpiresAt** | Pointer to **time.Time** | The date and time the token expires. | [optional] 

## Methods

### NewToken

`func NewToken(tokenId string, createdAt time.Time, tokenLabel string, creatorUserId string, userId string, accessLevel ApiAccessLevelEnum, ) *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *Token) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Token) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Token) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetCreatedAt

`func (o *Token) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Token) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Token) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTokenLabel

`func (o *Token) GetTokenLabel() string`

GetTokenLabel returns the TokenLabel field if non-nil, zero value otherwise.

### GetTokenLabelOk

`func (o *Token) GetTokenLabelOk() (*string, bool)`

GetTokenLabelOk returns a tuple with the TokenLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLabel

`func (o *Token) SetTokenLabel(v string)`

SetTokenLabel sets TokenLabel field to given value.


### GetCreatorUserId

`func (o *Token) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *Token) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *Token) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.


### GetUserId

`func (o *Token) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Token) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Token) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetLastUsedAt

`func (o *Token) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *Token) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *Token) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *Token) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetAccessLevel

`func (o *Token) GetAccessLevel() ApiAccessLevelEnum`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *Token) GetAccessLevelOk() (*ApiAccessLevelEnum, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *Token) SetAccessLevel(v ApiAccessLevelEnum)`

SetAccessLevel sets AccessLevel field to given value.


### GetExpiresAt

`func (o *Token) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Token) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Token) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Token) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


