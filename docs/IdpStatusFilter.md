# IdpStatusFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**[]UserHrIdpStatusEnum**](UserHrIdpStatusEnum.md) | Match users whose HR/IDP status is one of these values. | [optional] 
**Not** | Pointer to **bool** | Invert the match within the user domain (e.g. \&quot;IDP status is NOT active\&quot;). | [optional] 

## Methods

### NewIdpStatusFilter

`func NewIdpStatusFilter() *IdpStatusFilter`

NewIdpStatusFilter instantiates a new IdpStatusFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpStatusFilterWithDefaults

`func NewIdpStatusFilterWithDefaults() *IdpStatusFilter`

NewIdpStatusFilterWithDefaults instantiates a new IdpStatusFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *IdpStatusFilter) GetStatuses() []UserHrIdpStatusEnum`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *IdpStatusFilter) GetStatusesOk() (*[]UserHrIdpStatusEnum, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *IdpStatusFilter) SetStatuses(v []UserHrIdpStatusEnum)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *IdpStatusFilter) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetNot

`func (o *IdpStatusFilter) GetNot() bool`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *IdpStatusFilter) GetNotOk() (*bool, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *IdpStatusFilter) SetNot(v bool)`

SetNot sets Not field to given value.

### HasNot

`func (o *IdpStatusFilter) HasNot() bool`

HasNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


