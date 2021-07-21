# ResourceAccessStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource. | 
**UserId** | **string** | The ID of the user. | 
**AccessLevel** | Pointer to [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | [optional] 
**Status** | [**ResourceAccessStatusEnum**](ResourceAccessStatusEnum.md) |  | 
**ExpirationDate** | **NullableTime** | The day and time the user&#39;s access will expire. | 

## Methods

### NewResourceAccessStatus

`func NewResourceAccessStatus(resourceId string, userId string, status ResourceAccessStatusEnum, expirationDate NullableTime, ) *ResourceAccessStatus`

NewResourceAccessStatus instantiates a new ResourceAccessStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAccessStatusWithDefaults

`func NewResourceAccessStatusWithDefaults() *ResourceAccessStatus`

NewResourceAccessStatusWithDefaults instantiates a new ResourceAccessStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *ResourceAccessStatus) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceAccessStatus) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceAccessStatus) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetUserId

`func (o *ResourceAccessStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResourceAccessStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResourceAccessStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetAccessLevel

`func (o *ResourceAccessStatus) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ResourceAccessStatus) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ResourceAccessStatus) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *ResourceAccessStatus) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceAccessStatus) GetStatus() ResourceAccessStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceAccessStatus) GetStatusOk() (*ResourceAccessStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceAccessStatus) SetStatus(v ResourceAccessStatusEnum)`

SetStatus sets Status field to given value.


### GetExpirationDate

`func (o *ResourceAccessStatus) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ResourceAccessStatus) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ResourceAccessStatus) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *ResourceAccessStatus) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *ResourceAccessStatus) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


