# ResourceUserAccessStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource. | 
**UserId** | **string** | The ID of the user. | 
**AccessLevel** | Pointer to [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | [optional] 
**Status** | [**ResourceUserAccessStatusEnum**](ResourceUserAccessStatusEnum.md) |  | 
**ExpirationDate** | Pointer to **time.Time** | The day and time the user&#39;s access will expire. | [optional] 

## Methods

### NewResourceUserAccessStatus

`func NewResourceUserAccessStatus(resourceId string, userId string, status ResourceUserAccessStatusEnum, ) *ResourceUserAccessStatus`

NewResourceUserAccessStatus instantiates a new ResourceUserAccessStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUserAccessStatusWithDefaults

`func NewResourceUserAccessStatusWithDefaults() *ResourceUserAccessStatus`

NewResourceUserAccessStatusWithDefaults instantiates a new ResourceUserAccessStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *ResourceUserAccessStatus) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceUserAccessStatus) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceUserAccessStatus) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetUserId

`func (o *ResourceUserAccessStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResourceUserAccessStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResourceUserAccessStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetAccessLevel

`func (o *ResourceUserAccessStatus) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ResourceUserAccessStatus) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ResourceUserAccessStatus) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *ResourceUserAccessStatus) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceUserAccessStatus) GetStatus() ResourceUserAccessStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceUserAccessStatus) GetStatusOk() (*ResourceUserAccessStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceUserAccessStatus) SetStatus(v ResourceUserAccessStatusEnum)`

SetStatus sets Status field to given value.


### GetExpirationDate

`func (o *ResourceUserAccessStatus) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ResourceUserAccessStatus) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ResourceUserAccessStatus) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ResourceUserAccessStatus) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


