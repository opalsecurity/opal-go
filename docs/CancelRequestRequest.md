# CancelRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingUserId** | Pointer to **string** | Cancel as this Opal user instead of the caller. Only for service users an admin has allowed to act for users, in organizations with the feature enabled. The user must be active in the identity provider and be an Opal admin or the request&#39;s requester, target or initiated_by user. Unknown body fields are ignored, so a misspelled field cancels as the caller. | [optional] 

## Methods

### NewCancelRequestRequest

`func NewCancelRequestRequest() *CancelRequestRequest`

NewCancelRequestRequest instantiates a new CancelRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelRequestRequestWithDefaults

`func NewCancelRequestRequestWithDefaults() *CancelRequestRequest`

NewCancelRequestRequestWithDefaults instantiates a new CancelRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActingUserId

`func (o *CancelRequestRequest) GetActingUserId() string`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *CancelRequestRequest) GetActingUserIdOk() (*string, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *CancelRequestRequest) SetActingUserId(v string)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *CancelRequestRequest) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


