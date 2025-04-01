# BundleResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** | The ID of the bundle containing the resource. | [optional] [readonly] 
**ResourceId** | Pointer to **string** | The ID of the resource within a bundle. | [optional] [readonly] 
**AccessLevelName** | Pointer to **string** | The access level of the resource within a bundle. | [optional] 
**AccessLevelRemoteId** | Pointer to **string** | The remote ID of the access level of the resource within a bundle. | [optional] 

## Methods

### NewBundleResource

`func NewBundleResource() *BundleResource`

NewBundleResource instantiates a new BundleResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleResourceWithDefaults

`func NewBundleResourceWithDefaults() *BundleResource`

NewBundleResourceWithDefaults instantiates a new BundleResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *BundleResource) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *BundleResource) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *BundleResource) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *BundleResource) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetResourceId

`func (o *BundleResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BundleResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BundleResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BundleResource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetAccessLevelName

`func (o *BundleResource) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *BundleResource) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *BundleResource) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *BundleResource) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.

### GetAccessLevelRemoteId

`func (o *BundleResource) GetAccessLevelRemoteId() string`

GetAccessLevelRemoteId returns the AccessLevelRemoteId field if non-nil, zero value otherwise.

### GetAccessLevelRemoteIdOk

`func (o *BundleResource) GetAccessLevelRemoteIdOk() (*string, bool)`

GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelRemoteId

`func (o *BundleResource) SetAccessLevelRemoteId(v string)`

SetAccessLevelRemoteId sets AccessLevelRemoteId field to given value.

### HasAccessLevelRemoteId

`func (o *BundleResource) HasAccessLevelRemoteId() bool`

HasAccessLevelRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


