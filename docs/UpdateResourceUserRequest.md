# UpdateResourceUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMinutes** | **int32** | The updated duration for which the resource can be accessed (in minutes). Use 0 for indefinite. | 
**AccessLevelRemoteId** | Pointer to **string** | The updated remote ID of the access level granted to this user. | [optional] 
**ActingUserId** | Pointer to **string** | Record this Opal user as the actor of the access change in the audit event instead of the caller. Only for service users an admin has allowed to act for users, in organizations with the feature enabled. The named user must be a person (not a service user) in the caller&#39;s organization and active in the identity provider. Authorization is unchanged: the caller still needs permission to edit assignments on the resource. The named user&#39;s own permissions on the resource are not evaluated. | [optional] 

## Methods

### NewUpdateResourceUserRequest

`func NewUpdateResourceUserRequest(durationMinutes int32, ) *UpdateResourceUserRequest`

NewUpdateResourceUserRequest instantiates a new UpdateResourceUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceUserRequestWithDefaults

`func NewUpdateResourceUserRequestWithDefaults() *UpdateResourceUserRequest`

NewUpdateResourceUserRequestWithDefaults instantiates a new UpdateResourceUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMinutes

`func (o *UpdateResourceUserRequest) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *UpdateResourceUserRequest) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *UpdateResourceUserRequest) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.


### GetAccessLevelRemoteId

`func (o *UpdateResourceUserRequest) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *UpdateResourceUserRequest) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *UpdateResourceUserRequest) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *UpdateResourceUserRequest) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetActingUserId

`func (o *UpdateResourceUserRequest) GetActingUserId() string`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *UpdateResourceUserRequest) GetActingUserIdOk() (*string, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *UpdateResourceUserRequest) SetActingUserId(v string)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *UpdateResourceUserRequest) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


