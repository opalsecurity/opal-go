# CreateRequestCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingUserId** | Pointer to **string** | Comment as this Opal user instead of the caller. Only for service users an admin has allowed to act for users, in organizations with the feature enabled. The user must be active in the identity provider and be the request&#39;s requester, target or initiated_by user, an admin, or a reviewer who can act on or has acted on the request. | [optional] 
**Comment** | **string** | comment | 

## Methods

### NewCreateRequestCommentRequest

`func NewCreateRequestCommentRequest(comment string, ) *CreateRequestCommentRequest`

NewCreateRequestCommentRequest instantiates a new CreateRequestCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestCommentRequestWithDefaults

`func NewCreateRequestCommentRequestWithDefaults() *CreateRequestCommentRequest`

NewCreateRequestCommentRequestWithDefaults instantiates a new CreateRequestCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActingUserId

`func (o *CreateRequestCommentRequest) GetActingUserId() string`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *CreateRequestCommentRequest) GetActingUserIdOk() (*string, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *CreateRequestCommentRequest) SetActingUserId(v string)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *CreateRequestCommentRequest) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.

### GetComment

`func (o *CreateRequestCommentRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateRequestCommentRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateRequestCommentRequest) SetComment(v string)`

SetComment sets Comment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


