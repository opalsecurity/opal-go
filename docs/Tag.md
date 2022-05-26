# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | **string** | The ID of the tag. | 
**CreatedAt** | Pointer to **time.Time** | The date the tag was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date the tag was last updated. | [optional] 
**UserCreatorId** | Pointer to **string** | The ID of the user that created the tag. | [optional] 
**OwnerTeamId** | Pointer to **string** | The ID of the team that owns the tag. | [optional] 
**Key** | Pointer to **string** | The key of the tag. | [optional] 
**Value** | Pointer to **string** | The value of the tag. | [optional] 

## Methods

### NewTag

`func NewTag(tagId string, ) *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *Tag) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *Tag) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *Tag) SetTagId(v string)`

SetTagId sets TagId field to given value.


### GetCreatedAt

`func (o *Tag) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Tag) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Tag) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Tag) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Tag) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Tag) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Tag) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Tag) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserCreatorId

`func (o *Tag) GetUserCreatorId() string`

GetUserCreatorId returns the UserCreatorId field if non-nil, zero value otherwise.

### GetUserCreatorIdOk

`func (o *Tag) GetUserCreatorIdOk() (*string, bool)`

GetUserCreatorIdOk returns a tuple with the UserCreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCreatorId

`func (o *Tag) SetUserCreatorId(v string)`

SetUserCreatorId sets UserCreatorId field to given value.

### HasUserCreatorId

`func (o *Tag) HasUserCreatorId() bool`

HasUserCreatorId returns a boolean if a field has been set.

### GetOwnerTeamId

`func (o *Tag) GetOwnerTeamId() string`

GetOwnerTeamId returns the OwnerTeamId field if non-nil, zero value otherwise.

### GetOwnerTeamIdOk

`func (o *Tag) GetOwnerTeamIdOk() (*string, bool)`

GetOwnerTeamIdOk returns a tuple with the OwnerTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerTeamId

`func (o *Tag) SetOwnerTeamId(v string)`

SetOwnerTeamId sets OwnerTeamId field to given value.

### HasOwnerTeamId

`func (o *Tag) HasOwnerTeamId() bool`

HasOwnerTeamId returns a boolean if a field has been set.

### GetKey

`func (o *Tag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Tag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Tag) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Tag) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *Tag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Tag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Tag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Tag) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


