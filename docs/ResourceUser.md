# ResourceUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource. | 
**UserId** | **string** | The ID of the user. | 
**AccessLevel** | [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | 
**FullName** | **string** | The user&#39;s full name. | 
**Email** | **string** | The user&#39;s email. | 
**ExpirationDate** | Pointer to **time.Time** | The day and time the user&#39;s access will expire. | [optional] 

## Methods

### NewResourceUser

`func NewResourceUser(resourceId string, userId string, accessLevel ResourceAccessLevel, fullName string, email string, ) *ResourceUser`

NewResourceUser instantiates a new ResourceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUserWithDefaults

`func NewResourceUserWithDefaults() *ResourceUser`

NewResourceUserWithDefaults instantiates a new ResourceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *ResourceUser) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceUser) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceUser) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetUserId

`func (o *ResourceUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResourceUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResourceUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetAccessLevel

`func (o *ResourceUser) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ResourceUser) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ResourceUser) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.


### GetFullName

`func (o *ResourceUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ResourceUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ResourceUser) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEmail

`func (o *ResourceUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResourceUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResourceUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpirationDate

`func (o *ResourceUser) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ResourceUser) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ResourceUser) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ResourceUser) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


