# DeleteGroupUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingUserId** | Pointer to **string** | Record this Opal user as the actor of the membership removal in the audit event instead of the caller. Only for service users an admin has allowed to act for users, in organizations with the feature enabled. The named user must be a person (not a service user) in the caller&#39;s organization and active in the identity provider. Authorization is unchanged: the caller still needs permission to edit assignments on the group. The named user&#39;s own permissions on the group are not evaluated. Unknown body fields are rejected, so a misspelled field fails instead of removing access as the caller. | [optional] 

## Methods

### NewDeleteGroupUserRequest

`func NewDeleteGroupUserRequest() *DeleteGroupUserRequest`

NewDeleteGroupUserRequest instantiates a new DeleteGroupUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteGroupUserRequestWithDefaults

`func NewDeleteGroupUserRequestWithDefaults() *DeleteGroupUserRequest`

NewDeleteGroupUserRequestWithDefaults instantiates a new DeleteGroupUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActingUserId

`func (o *DeleteGroupUserRequest) GetActingUserId() string`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *DeleteGroupUserRequest) GetActingUserIdOk() (*string, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *DeleteGroupUserRequest) SetActingUserId(v string)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *DeleteGroupUserRequest) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


