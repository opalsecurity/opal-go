# ResourceRemoteInfoAwsEc2Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | The instanceId of the EC2 instance. | 
**Region** | **string** | The region of the EC2 instance. | 
**AccountId** | Pointer to **string** | The id of the AWS account. Required for AWS Organizations. | [optional] 

## Methods

### NewResourceRemoteInfoAwsEc2Instance

`func NewResourceRemoteInfoAwsEc2Instance(instanceId string, region string, ) *ResourceRemoteInfoAwsEc2Instance`

NewResourceRemoteInfoAwsEc2Instance instantiates a new ResourceRemoteInfoAwsEc2Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoAwsEc2InstanceWithDefaults

`func NewResourceRemoteInfoAwsEc2InstanceWithDefaults() *ResourceRemoteInfoAwsEc2Instance`

NewResourceRemoteInfoAwsEc2InstanceWithDefaults instantiates a new ResourceRemoteInfoAwsEc2Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *ResourceRemoteInfoAwsEc2Instance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ResourceRemoteInfoAwsEc2Instance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ResourceRemoteInfoAwsEc2Instance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetRegion

`func (o *ResourceRemoteInfoAwsEc2Instance) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ResourceRemoteInfoAwsEc2Instance) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ResourceRemoteInfoAwsEc2Instance) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccountId

`func (o *ResourceRemoteInfoAwsEc2Instance) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResourceRemoteInfoAwsEc2Instance) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResourceRemoteInfoAwsEc2Instance) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ResourceRemoteInfoAwsEc2Instance) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


