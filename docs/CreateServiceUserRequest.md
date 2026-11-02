# CreateServiceUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name of the service user. | 
**OwnerId** | **string** | The ID of the owner (team of record) that owns the service user. Must reference an existing owner in the caller&#39;s organization. | 

## Methods

### NewCreateServiceUserRequest

`func NewCreateServiceUserRequest(name string, ownerId string, ) *CreateServiceUserRequest`

NewCreateServiceUserRequest instantiates a new CreateServiceUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceUserRequestWithDefaults

`func NewCreateServiceUserRequestWithDefaults() *CreateServiceUserRequest`

NewCreateServiceUserRequestWithDefaults instantiates a new CreateServiceUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateServiceUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *CreateServiceUserRequest) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreateServiceUserRequest) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreateServiceUserRequest) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


