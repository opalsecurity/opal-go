# AgentActivityIngestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Received** | **int32** | Events parsed out of the request body. | 
**Stored** | **int32** | Rows actually written. Lower than &#x60;received&#x60; when a retried delivery was deduplicated on &#x60;tool_use_id&#x60;. | 
**Attributed** | **int32** | Events that resolved to an agent in this organization. | 
**Unattributed** | **int32** | Events that resolved to no agent, and are therefore visible only in the unattributed activity list. A persistently high value means the reporting clients are not sending an identity Opal can resolve. | 

## Methods

### NewAgentActivityIngestResult

`func NewAgentActivityIngestResult(received int32, stored int32, attributed int32, unattributed int32, ) *AgentActivityIngestResult`

NewAgentActivityIngestResult instantiates a new AgentActivityIngestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityIngestResultWithDefaults

`func NewAgentActivityIngestResultWithDefaults() *AgentActivityIngestResult`

NewAgentActivityIngestResultWithDefaults instantiates a new AgentActivityIngestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceived

`func (o *AgentActivityIngestResult) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *AgentActivityIngestResult) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *AgentActivityIngestResult) SetReceived(v int32)`

SetReceived sets Received field to given value.


### GetStored

`func (o *AgentActivityIngestResult) GetStored() int32`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *AgentActivityIngestResult) GetStoredOk() (*int32, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *AgentActivityIngestResult) SetStored(v int32)`

SetStored sets Stored field to given value.


### GetAttributed

`func (o *AgentActivityIngestResult) GetAttributed() int32`

GetAttributed returns the Attributed field if non-nil, zero value otherwise.

### GetAttributedOk

`func (o *AgentActivityIngestResult) GetAttributedOk() (*int32, bool)`

GetAttributedOk returns a tuple with the Attributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributed

`func (o *AgentActivityIngestResult) SetAttributed(v int32)`

SetAttributed sets Attributed field to given value.


### GetUnattributed

`func (o *AgentActivityIngestResult) GetUnattributed() int32`

GetUnattributed returns the Unattributed field if non-nil, zero value otherwise.

### GetUnattributedOk

`func (o *AgentActivityIngestResult) GetUnattributedOk() (*int32, bool)`

GetUnattributedOk returns a tuple with the Unattributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnattributed

`func (o *AgentActivityIngestResult) SetUnattributed(v int32)`

SetUnattributed sets Unattributed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


