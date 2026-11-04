# PaladinContextSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the context source. | 
**PaladinId** | **string** | The ID of the Paladin this source belongs to. | 
**SourceKind** | [**PaladinContextSourceKind**](PaladinContextSourceKind.md) |  | 
**ThirdPartyProvider** | [**PaladinContextSourceProvider**](PaladinContextSourceProvider.md) |  | 
**RemoteId** | **string** | The provider&#39;s identifier for the source. The Slack channel ID for a channel, or the Notion/Confluence page ID for a document. | 
**Name** | **string** | A human-readable name for the source. | 
**Url** | **string** | A link to the source. | 

## Methods

### NewPaladinContextSource

`func NewPaladinContextSource(id string, paladinId string, sourceKind PaladinContextSourceKind, thirdPartyProvider PaladinContextSourceProvider, remoteId string, name string, url string, ) *PaladinContextSource`

NewPaladinContextSource instantiates a new PaladinContextSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaladinContextSourceWithDefaults

`func NewPaladinContextSourceWithDefaults() *PaladinContextSource`

NewPaladinContextSourceWithDefaults instantiates a new PaladinContextSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaladinContextSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaladinContextSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaladinContextSource) SetId(v string)`

SetId sets Id field to given value.


### GetPaladinId

`func (o *PaladinContextSource) GetPaladinId() string`

GetPaladinId returns the PaladinId field if non-nil, zero value otherwise.

### GetPaladinIdOk

`func (o *PaladinContextSource) GetPaladinIdOk() (*string, bool)`

GetPaladinIdOk returns a tuple with the PaladinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaladinId

`func (o *PaladinContextSource) SetPaladinId(v string)`

SetPaladinId sets PaladinId field to given value.


### GetSourceKind

`func (o *PaladinContextSource) GetSourceKind() PaladinContextSourceKind`

GetSourceKind returns the SourceKind field if non-nil, zero value otherwise.

### GetSourceKindOk

`func (o *PaladinContextSource) GetSourceKindOk() (*PaladinContextSourceKind, bool)`

GetSourceKindOk returns a tuple with the SourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKind

`func (o *PaladinContextSource) SetSourceKind(v PaladinContextSourceKind)`

SetSourceKind sets SourceKind field to given value.


### GetThirdPartyProvider

`func (o *PaladinContextSource) GetThirdPartyProvider() PaladinContextSourceProvider`

GetThirdPartyProvider returns the ThirdPartyProvider field if non-nil, zero value otherwise.

### GetThirdPartyProviderOk

`func (o *PaladinContextSource) GetThirdPartyProviderOk() (*PaladinContextSourceProvider, bool)`

GetThirdPartyProviderOk returns a tuple with the ThirdPartyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyProvider

`func (o *PaladinContextSource) SetThirdPartyProvider(v PaladinContextSourceProvider)`

SetThirdPartyProvider sets ThirdPartyProvider field to given value.


### GetRemoteId

`func (o *PaladinContextSource) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PaladinContextSource) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PaladinContextSource) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetName

`func (o *PaladinContextSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaladinContextSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaladinContextSource) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *PaladinContextSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaladinContextSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaladinContextSource) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


