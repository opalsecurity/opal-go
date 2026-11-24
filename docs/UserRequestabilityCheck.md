# UserRequestabilityCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **string** | The ID of the resource that was checked. Set for resource checks. | [optional] 
**GroupId** | Pointer to **string** | The ID of the group that was checked. Set for group checks. | [optional] 
**AppId** | Pointer to **string** | The ID of the app that was checked. Set for app checks. | [optional] 
**UserId** | **string** | The ID of the user that was checked. | 
**Requestable** | **bool** |  | 
**Reason** | [**RequestabilityReasonEnum**](RequestabilityReasonEnum.md) |  | 
**AccessLevels** | Pointer to [**[]RequestableAccessLevel**](RequestableAccessLevel.md) | Per-access-level requestability. Omitted for apps; empty remote ID for assets without levels. | [optional] 

## Methods

### NewUserRequestabilityCheck

`func NewUserRequestabilityCheck(userId string, requestable bool, reason RequestabilityReasonEnum, ) *UserRequestabilityCheck`

NewUserRequestabilityCheck instantiates a new UserRequestabilityCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRequestabilityCheckWithDefaults

`func NewUserRequestabilityCheckWithDefaults() *UserRequestabilityCheck`

NewUserRequestabilityCheckWithDefaults instantiates a new UserRequestabilityCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *UserRequestabilityCheck) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *UserRequestabilityCheck) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *UserRequestabilityCheck) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *UserRequestabilityCheck) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetGroupId

`func (o *UserRequestabilityCheck) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UserRequestabilityCheck) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UserRequestabilityCheck) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UserRequestabilityCheck) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAppId

`func (o *UserRequestabilityCheck) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UserRequestabilityCheck) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UserRequestabilityCheck) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UserRequestabilityCheck) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetUserId

`func (o *UserRequestabilityCheck) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserRequestabilityCheck) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserRequestabilityCheck) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetRequestable

`func (o *UserRequestabilityCheck) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *UserRequestabilityCheck) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *UserRequestabilityCheck) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.


### GetReason

`func (o *UserRequestabilityCheck) GetReason() RequestabilityReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UserRequestabilityCheck) GetReasonOk() (*RequestabilityReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UserRequestabilityCheck) SetReason(v RequestabilityReasonEnum)`

SetReason sets Reason field to given value.


### GetAccessLevels

`func (o *UserRequestabilityCheck) GetAccessLevels() []RequestableAccessLevel`

GetAccessLevels returns the AccessLevels field if non-nil, zero value otherwise.

### GetAccessLevelsOk

`func (o *UserRequestabilityCheck) GetAccessLevelsOk() (*[]RequestableAccessLevel, bool)`

GetAccessLevelsOk returns a tuple with the AccessLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevels

`func (o *UserRequestabilityCheck) SetAccessLevels(v []RequestableAccessLevel)`

SetAccessLevels sets AccessLevels field to given value.

### HasAccessLevels

`func (o *UserRequestabilityCheck) HasAccessLevels() bool`

HasAccessLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


