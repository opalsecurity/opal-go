# DeleteResourceUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingUserId** | Pointer to **string** | Record this Opal user as the actor of the access removal in the audit event instead of the caller. Only for service users an admin has allowed to act for users, in organizations with the feature enabled. The named user must be a person (not a service user) in the caller&#39;s organization and active in the identity provider. Authorization is unchanged: the caller still needs permission to edit assignments on the resource. The named user&#39;s own permissions on the resource are not evaluated. Unknown body fields are rejected, so a misspelled field fails instead of removing access as the caller. | [optional] 

## Methods

### NewDeleteResourceUserRequest

`func NewDeleteResourceUserRequest() *DeleteResourceUserRequest`

NewDeleteResourceUserRequest instantiates a new DeleteResourceUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteResourceUserRequestWithDefaults

`func NewDeleteResourceUserRequestWithDefaults() *DeleteResourceUserRequest`

NewDeleteResourceUserRequestWithDefaults instantiates a new DeleteResourceUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActingUserId

`func (o *DeleteResourceUserRequest) GetActingUserId() string`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *DeleteResourceUserRequest) GetActingUserIdOk() (*string, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *DeleteResourceUserRequest) SetActingUserId(v string)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *DeleteResourceUserRequest) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


