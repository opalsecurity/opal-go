# UpdateGroupUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMinutes** | **int32** | The updated duration for which the group can be accessed (in minutes). Use 0 for indefinite, or a negative value to revoke access. | 
**AccessLevelRemoteId** | Pointer to **string** | The updated remote ID of the access level granted to this user. | [optional] 
**ActingUserId** | Pointer to **string** | Record this Opal user as the actor of the membership change in the audit event instead of the caller. Only for service users an admin has allowed to act for users, in organizations with the feature enabled. The named user must be a person (not a service user) in the caller&#39;s organization and active in the identity provider. Authorization is unchanged: the caller still needs permission to edit assignments on the group. The named user&#39;s own permissions on the group are not evaluated. | [optional] 

## Methods

### NewUpdateGroupUserRequest

`func NewUpdateGroupUserRequest(durationMinutes int32, ) *UpdateGroupUserRequest`

NewUpdateGroupUserRequest instantiates a new UpdateGroupUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupUserRequestWithDefaults

`func NewUpdateGroupUserRequestWithDefaults() *UpdateGroupUserRequest`

NewUpdateGroupUserRequestWithDefaults instantiates a new UpdateGroupUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMinutes

`func (o *UpdateGroupUserRequest) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *UpdateGroupUserRequest) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *UpdateGroupUserRequest) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.


### GetAccessLevelRemoteId

`func (o *UpdateGroupUserRequest) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *UpdateGroupUserRequest) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *UpdateGroupUserRequest) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *UpdateGroupUserRequest) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetActingUserId

`func (o *UpdateGroupUserRequest) GetActingUserId() string`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *UpdateGroupUserRequest) GetActingUserIdOk() (*string, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *UpdateGroupUserRequest) SetActingUserId(v string)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *UpdateGroupUserRequest) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


