# CreatePaladinContextSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceKind** | [**PaladinContextSourceKind**](PaladinContextSourceKind.md) |  | 
**ThirdPartyProvider** | [**PaladinContextSourceProvider**](PaladinContextSourceProvider.md) |  | 
**RemoteId** | **string** | The provider&#39;s identifier for the source. The Slack channel ID for a channel, or the Notion/Confluence page ID for a document. | 
**Name** | Pointer to **string** | An optional human-readable name for the source. | [optional] 
**Url** | Pointer to **string** | An optional link to the source. | [optional] 

## Methods

### NewCreatePaladinContextSourceInfo

`func NewCreatePaladinContextSourceInfo(sourceKind PaladinContextSourceKind, thirdPartyProvider PaladinContextSourceProvider, remoteId string, ) *CreatePaladinContextSourceInfo`

NewCreatePaladinContextSourceInfo instantiates a new CreatePaladinContextSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaladinContextSourceInfoWithDefaults

`func NewCreatePaladinContextSourceInfoWithDefaults() *CreatePaladinContextSourceInfo`

NewCreatePaladinContextSourceInfoWithDefaults instantiates a new CreatePaladinContextSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceKind

`func (o *CreatePaladinContextSourceInfo) GetSourceKind() PaladinContextSourceKind`

GetSourceKind returns the SourceKind field if non-nil, zero value otherwise.

### GetSourceKindOk

`func (o *CreatePaladinContextSourceInfo) GetSourceKindOk() (*PaladinContextSourceKind, bool)`

GetSourceKindOk returns a tuple with the SourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKind

`func (o *CreatePaladinContextSourceInfo) SetSourceKind(v PaladinContextSourceKind)`

SetSourceKind sets SourceKind field to given value.


### GetThirdPartyProvider

`func (o *CreatePaladinContextSourceInfo) GetThirdPartyProvider() PaladinContextSourceProvider`

GetThirdPartyProvider returns the ThirdPartyProvider field if non-nil, zero value otherwise.

### GetThirdPartyProviderOk

`func (o *CreatePaladinContextSourceInfo) GetThirdPartyProviderOk() (*PaladinContextSourceProvider, bool)`

GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyProvider

`func (o *CreatePaladinContextSourceInfo) SetThirdPartyProvider(v PaladinContextSourceProvider)`

SetThirdPartyProvider sets ThirdPartyProvider field to given value.


### GetRemoteId

`func (o *CreatePaladinContextSourceInfo) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreatePaladinContextSourceInfo) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreatePaladinContextSourceInfo) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetName

`func (o *CreatePaladinContextSourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePaladinContextSourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePaladinContextSourceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePaladinContextSourceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *CreatePaladinContextSourceInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreatePaladinContextSourceInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreatePaladinContextSourceInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreatePaladinContextSourceInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


