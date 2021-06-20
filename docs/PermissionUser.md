# PermissionUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionId** | **string** | The ID of the permission. | 
**UserId** | **string** | The ID of the user. | 
**FullName** | **string** | The user&#39;s full name. | 
**Email** | **string** | The user&#39;s email. | 
**ExpirationDate** | **NullableTime** | The day and time the user&#39;s access will expire. | 

## Methods

### NewPermissionUser

`func NewPermissionUser(permissionId string, userId string, fullName string, email string, expirationDate NullableTime, ) *PermissionUser`

NewPermissionUser instantiates a new PermissionUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionUserWithDefaults

`func NewPermissionUserWithDefaults() *PermissionUser`

NewPermissionUserWithDefaults instantiates a new PermissionUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionId

`func (o *PermissionUser) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *PermissionUser) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *PermissionUser) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.


### GetUserId

`func (o *PermissionUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PermissionUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PermissionUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFullName

`func (o *PermissionUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PermissionUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PermissionUser) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEmail

`func (o *PermissionUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PermissionUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PermissionUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpirationDate

`func (o *PermissionUser) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PermissionUser) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PermissionUser) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *PermissionUser) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PermissionUser) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


