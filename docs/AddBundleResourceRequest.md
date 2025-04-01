# AddBundleResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource to add. | 
**AccessLevelRemoteId** | Pointer to **string** | The remote ID of the access level to grant to this user. Required if the resource being added requires an access level. If omitted, the default access level remote ID value (empty string) is used. | [optional] 
**AccessLevelName** | Pointer to **string** | The name of the access level to grant to this user. If omitted, the default access level name value (empty string) is used. | [optional] 

## Methods

### NewAddBundleResourceRequest

`func NewAddBundleResourceRequest(resourceId string, ) *AddBundleResourceRequest`

NewAddBundleResourceRequest instantiates a new AddBundleResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBundleResourceRequestWithDefaults

`func NewAddBundleResourceRequestWithDefaults() *AddBundleResourceRequest`

NewAddBundleResourceRequestWithDefaults instantiates a new AddBundleResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *AddBundleResourceRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AddBundleResourceRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AddBundleResourceRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetAccessLevelRemoteId

`func (o *AddBundleResourceRequest) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *AddBundleResourceRequest) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *AddBundleResourceRequest) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *AddBundleResourceRequest) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.

### GetAccessLevelName

`func (o *AddBundleResourceRequest) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *AddBundleResourceRequest) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *AddBundleResourceRequest) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *AddBundleResourceRequest) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


