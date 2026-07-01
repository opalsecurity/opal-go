# ResourceRemoteInfoAwsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The id of the AWS account. | 
**OrganizationalUnitId** | Pointer to **string** | The id of the AWS organizational unit. Required only if customer has OUs enabled. | [optional] 

## Methods

### NewResourceRemoteInfoAwsAccount

`func NewResourceRemoteInfoAwsAccount(accountId string, ) *ResourceRemoteInfoAwsAccount`

NewResourceRemoteInfoAwsAccount instantiates a new ResourceRemoteInfoAwsAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRemoteInfoAwsAccountWithDefaults

`func NewResourceRemoteInfoAwsAccountWithDefaults() *ResourceRemoteInfoAwsAccount`

NewResourceRemoteInfoAwsAccountWithDefaults instantiates a new ResourceRemoteInfoAwsAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ResourceRemoteInfoAwsAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResourceRemoteInfoAwsAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResourceRemoteInfoAwsAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetOrganizationalUnitId

`func (o *ResourceRemoteInfoAwsAccount) GetOrganizationalUnitId() string`

GetOrganizationalUnitId returns the OrganizationalUnitId field if non-nil, zero value otherwise.

### GetOrganizationalUnitIdOk

`func (o *ResourceRemoteInfoAwsAccount) GetOrganizationalUnitIdOk() (*string, bool)`

GetOrganizationalUnitIdOk returns a tuple with the OrganizationalUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnitId

`func (o *ResourceRemoteInfoAwsAccount) SetOrganizationalUnitId(v string)`

SetOrganizationalUnitId sets OrganizationalUnitId field to given value.

### HasOrganizationalUnitId

`func (o *ResourceRemoteInfoAwsAccount) HasOrganizationalUnitId() bool`

HasOrganizationalUnitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


