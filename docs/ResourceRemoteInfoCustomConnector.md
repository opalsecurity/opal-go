# ResourceRemoteInfoCustomConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteResourceId** | **string** | The id of the resource in the end system | 
**CanHaveUsageEvents** | **bool** | A bool representing whether or not the resource can have usage data. | 

## Methods

### NewResourceRemoteInfoCustomConnector

`func NewResourceRemoteInfoCustomConnector(remoteResourceId string, canHaveUsageEvents bool, ) *ResourceRemoteInfoCustomConnector`

NewResourceRemoteInfoCustomConnector instantiates a new ResourceRemoteInfoCustomConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoCustomConnectorWithDefaults

`func NewResourceRemoteInfoCustomConnectorWithDefaults() *ResourceRemoteInfoCustomConnector`

NewResourceRemoteInfoCustomConnectorWithDefaults instantiates a new ResourceRemoteInfoCustomConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteResourceId

`func (o *ResourceRemoteInfoCustomConnector) GetRemoteResourceId() string`

GetRemoteResourceId returns the RemoteResourceId field if non-nil, zero value otherwise.

### GetRemoteResourceIdOk

`func (o *ResourceRemoteInfoCustomConnector) GetRemoteResourceIdOk() (*string, bool)`

GetRemoteResourceIdOk returns a tuple with the RemoteResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteResourceId

`func (o *ResourceRemoteInfoCustomConnector) SetRemoteResourceId(v string)`

SetRemoteResourceId sets RemoteResourceId field to given value.


### GetCanHaveUsageEvents

`func (o *ResourceRemoteInfoCustomConnector) GetCanHaveUsageEvents() bool`

GetCanHaveUsageEvents returns the CanHaveUsageEvents field if non-nil, zero value otherwise.

### GetCanHaveUsageEventsOk

`func (o *ResourceRemoteInfoCustomConnector) GetCanHaveUsageEventsOk() (*bool, bool)`

GetCanHaveUsageEventsOk returns a tuple with the CanHaveUsageEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanHaveUsageEvents

`func (o *ResourceRemoteInfoCustomConnector) SetCanHaveUsageEvents(v bool)`

SetCanHaveUsageEvents sets CanHaveUsageEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


