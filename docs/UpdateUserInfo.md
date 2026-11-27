# UpdateUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagerId** | Pointer to **string** | The user&#39;s manager ID. Set to null to remove the manager. | [optional] 
**Position** | Pointer to **string** | The user&#39;s position. | [optional] 

## Methods

### NewUpdateUserInfo

`func NewUpdateUserInfo() *UpdateUserInfo`

NewUpdateUserInfo instantiates a new UpdateUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserInfoWithDefaults

`func NewUpdateUserInfoWithDefaults() *UpdateUserInfo`

NewUpdateUserInfoWithDefaults instantiates a new UpdateUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagerId

`func (o *UpdateUserInfo) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *UpdateUserInfo) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *UpdateUserInfo) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *UpdateUserInfo) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### GetPosition

`func (o *UpdateUserInfo) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UpdateUserInfo) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UpdateUserInfo) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UpdateUserInfo) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


