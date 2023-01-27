# GroupUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group. | 
**UserId** | **string** | The ID of the user. | 
**AccessLevel** | Pointer to [**GroupAccessLevel**](GroupAccessLevel.md) |  | [optional] 
**FullName** | **string** | The user&#39;s full name. | 
**Email** | **string** | The user&#39;s email. | 
**ExpirationDate** | **NullableTime** | The day and time the user&#39;s access will expire. | 

## Methods

### NewGroupUser

`func NewGroupUser(groupId string, userId string, fullName string, email string, expirationDate NullableTime, ) *GroupUser`

NewGroupUser instantiates a new GroupUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserWithDefaults

`func NewGroupUserWithDefaults() *GroupUser`

NewGroupUserWithDefaults instantiates a new GroupUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupUser) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupUser) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupUser) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetUserId

`func (o *GroupUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetAccessLevel

`func (o *GroupUser) GetAccessLevel() GroupAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *GroupUser) GetAccessLevelOk() (*GroupAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *GroupUser) SetAccessLevel(v GroupAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *GroupUser) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetFullName

`func (o *GroupUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GroupUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GroupUser) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEmail

`func (o *GroupUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GroupUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GroupUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpirationDate

`func (o *GroupUser) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *GroupUser) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *GroupUser) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *GroupUser) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *GroupUser) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


