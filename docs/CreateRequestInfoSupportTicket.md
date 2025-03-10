# CreateRequestInfoSupportTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TicketingProvider** | [**TicketingProviderEnum**](TicketingProviderEnum.md) |  | 
**RemoteId** | **string** |  | 
**Identifier** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewCreateRequestInfoSupportTicket

`func NewCreateRequestInfoSupportTicket(ticketingProvider TicketingProviderEnum, remoteId string, identifier string, url string, ) *CreateRequestInfoSupportTicket`

NewCreateRequestInfoSupportTicket instantiates a new CreateRequestInfoSupportTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestInfoSupportTicketWithDefaults

`func NewCreateRequestInfoSupportTicketWithDefaults() *CreateRequestInfoSupportTicket`

NewCreateRequestInfoSupportTicketWithDefaults instantiates a new CreateRequestInfoSupportTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicketingProvider

`func (o *CreateRequestInfoSupportTicket) GetTicketingProvider() TicketingProviderEnum`

GetTicketingProvider returns the TicketingProvider field if non-nil, zero value otherwise.

### GetTicketingProviderOk

`func (o *CreateRequestInfoSupportTicket) GetTicketingProviderOk() (*TicketingProviderEnum, bool)`

GetTicketingProviderOk returns a tuple with the TicketingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketingProvider

`func (o *CreateRequestInfoSupportTicket) SetTicketingProvider(v TicketingProviderEnum)`

SetTicketingProvider sets TicketingProvider field to given value.


### GetRemoteId

`func (o *CreateRequestInfoSupportTicket) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreateRequestInfoSupportTicket) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreateRequestInfoSupportTicket) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetIdentifier

`func (o *CreateRequestInfoSupportTicket) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateRequestInfoSupportTicket) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateRequestInfoSupportTicket) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetUrl

`func (o *CreateRequestInfoSupportTicket) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateRequestInfoSupportTicket) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateRequestInfoSupportTicket) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


