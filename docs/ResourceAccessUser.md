# ResourceAccessUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource. | 
**UserId** | **string** | The ID of the user. | 
**AccessLevel** | [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | 
**FullName** | **string** | The user&#39;s full name. | 
**Email** | **string** | The user&#39;s email. | 
**ExpirationDate** | **NullableTime** | The day and time the user&#39;s access will expire. | 
**HasDirectAccess** | **bool** | The user has direct access to this resources (vs. indirectly, like through a group). | 
**NumAccessPaths** | **int32** | The number of ways in which the user has access through this resource (directly and indirectly). | 

## Methods

### NewResourceAccessUser

`func NewResourceAccessUser(resourceId string, userId string, accessLevel ResourceAccessLevel, fullName string, email string, expirationDate NullableTime, hasDirectAccess bool, numAccessPaths int32, ) *ResourceAccessUser`

NewResourceAccessUser instantiates a new ResourceAccessUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAccessUserWithDefaults

`func NewResourceAccessUserWithDefaults() *ResourceAccessUser`

NewResourceAccessUserWithDefaults instantiates a new ResourceAccessUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *ResourceAccessUser) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceAccessUser) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceAccessUser) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetUserId

`func (o *ResourceAccessUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResourceAccessUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResourceAccessUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetAccessLevel

`func (o *ResourceAccessUser) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ResourceAccessUser) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ResourceAccessUser) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.


### GetFullName

`func (o *ResourceAccessUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ResourceAccessUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ResourceAccessUser) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEmail

`func (o *ResourceAccessUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResourceAccessUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResourceAccessUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpirationDate

`func (o *ResourceAccessUser) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ResourceAccessUser) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ResourceAccessUser) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *ResourceAccessUser) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *ResourceAccessUser) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetHasDirectAccess

`func (o *ResourceAccessUser) GetHasDirectAccess() bool`

GetHasDirectAccess returns the HasDirectAccess field if non-nil, zero value otherwise.

### GetHasDirectAccessOk

`func (o *ResourceAccessUser) GetHasDirectAccessOk() (*bool, bool)`

GetHasDirectAccessOk returns a tuple with the HasDirectAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDirectAccess

`func (o *ResourceAccessUser) SetHasDirectAccess(v bool)`

SetHasDirectAccess sets HasDirectAccess field to given value.


### GetNumAccessPaths

`func (o *ResourceAccessUser) GetNumAccessPaths() int32`

GetNumAccessPaths returns the NumAccessPaths field if non-nil, zero value otherwise.

### GetNumAccessPathsOk

`func (o *ResourceAccessUser) GetNumAccessPathsOk() (*int32, bool)`

GetNumAccessPathsOk returns a tuple with the NumAccessPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAccessPaths

`func (o *ResourceAccessUser) SetNumAccessPaths(v int32)`

SetNumAccessPaths sets NumAccessPaths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


