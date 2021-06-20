# PermissionAccessStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionId** | **string** | The ID of the permission. | 
**UserId** | **string** | The ID of the user. | 
**Status** | [**PermissionAccessStatusEnum**](PermissionAccessStatusEnum.md) |  | 
**ExpirationDate** | **NullableTime** | The day and time the user&#39;s access will expire. | 

## Methods

### NewPermissionAccessStatus

`func NewPermissionAccessStatus(permissionId string, userId string, status PermissionAccessStatusEnum, expirationDate NullableTime, ) *PermissionAccessStatus`

NewPermissionAccessStatus instantiates a new PermissionAccessStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionAccessStatusWithDefaults

`func NewPermissionAccessStatusWithDefaults() *PermissionAccessStatus`

NewPermissionAccessStatusWithDefaults instantiates a new PermissionAccessStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionId

`func (o *PermissionAccessStatus) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *PermissionAccessStatus) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *PermissionAccessStatus) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.


### GetUserId

`func (o *PermissionAccessStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PermissionAccessStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PermissionAccessStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetStatus

`func (o *PermissionAccessStatus) GetStatus() PermissionAccessStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PermissionAccessStatus) GetStatusOk() (*PermissionAccessStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PermissionAccessStatus) SetStatus(v PermissionAccessStatusEnum)`

SetStatus sets Status field to given value.


### GetExpirationDate

`func (o *PermissionAccessStatus) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PermissionAccessStatus) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PermissionAccessStatus) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *PermissionAccessStatus) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PermissionAccessStatus) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


