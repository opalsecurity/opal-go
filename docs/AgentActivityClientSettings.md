# AgentActivityClientSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchWaitSeconds** | **float64** | The minimum gap between two sends from one machine. Send at once when the last send was at least this long ago; otherwise wait until this long after the last send, and send what arrived meanwhile as one request. Between 0 and 300. | 
**MaxEventsPerRequest** | **int32** | The most events one request to &#x60;/agent-activity&#x60; may carry. | 
**MaxBytesPerRequest** | **int64** | The largest body, in bytes, &#x60;/agent-activity&#x60; accepts for a request with more than one event. A request with one event may be up to 1 MiB. | 
**MaxRequestsPerSecond** | **float64** | The rate of batched requests one client should stay under. Can be below 1. | 
**BurstRequests** | **int32** | How many batched requests a client may send at once before the rate applies. At least 1. | 
**CacheSeconds** | **int32** | How long these settings may be kept before asking again. | 

## Methods

### NewAgentActivityClientSettings

`func NewAgentActivityClientSettings(batchWaitSeconds float64, maxEventsPerRequest int32, maxBytesPerRequest int64, maxRequestsPerSecond float64, burstRequests int32, cacheSeconds int32, ) *AgentActivityClientSettings`

NewAgentActivityClientSettings instantiates a new AgentActivityClientSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityClientSettingsWithDefaults

`func NewAgentActivityClientSettingsWithDefaults() *AgentActivityClientSettings`

NewAgentActivityClientSettingsWithDefaults instantiates a new AgentActivityClientSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchWaitSeconds

`func (o *AgentActivityClientSettings) GetBatchWaitSeconds() float64`

GetBatchWaitSeconds returns the BatchWaitSeconds field if non-nil, zero value otherwise.

### GetBatchWaitSecondsOk

`func (o *AgentActivityClientSettings) GetBatchWaitSecondsOk() (*float64, bool)`

GetBatchWaitSecondsOk returns a tuple with the BatchWaitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchWaitSeconds

`func (o *AgentActivityClientSettings) SetBatchWaitSeconds(v float64)`

SetBatchWaitSeconds sets BatchWaitSeconds field to given value.


### GetMaxEventsPerRequest

`func (o *AgentActivityClientSettings) GetMaxEventsPerRequest() int32`

GetMaxEventsPerRequest returns the MaxEventsPerRequest field if non-nil, zero value otherwise.

### GetMaxEventsPerRequestOk

`func (o *AgentActivityClientSettings) GetMaxEventsPerRequestOk() (*int32, bool)`

GetMaxEventsPerRequestOk returns a tuple with the MaxEventsPerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventsPerRequest

`func (o *AgentActivityClientSettings) SetMaxEventsPerRequest(v int32)`

SetMaxEventsPerRequest sets MaxEventsPerRequest field to given value.


### GetMaxBytesPerRequest

`func (o *AgentActivityClientSettings) GetMaxBytesPerRequest() int64`

GetMaxBytesPerRequest returns the MaxBytesPerRequest field if non-nil, zero value otherwise.

### GetMaxBytesPerRequestOk

`func (o *AgentActivityClientSettings) GetMaxBytesPerRequestOk() (*int64, bool)`

GetMaxBytesPerRequestOk returns a tuple with the MaxBytesPerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytesPerRequest

`func (o *AgentActivityClientSettings) SetMaxBytesPerRequest(v int64)`

SetMaxBytesPerRequest sets MaxBytesPerRequest field to given value.


### GetMaxRequestsPerSecond

`func (o *AgentActivityClientSettings) GetMaxRequestsPerSecond() float64`

GetMaxRequestsPerSecond returns the MaxRequestsPerSecond field if non-nil, zero value otherwise.

### GetMaxRequestsPerSecondOk

`func (o *AgentActivityClientSettings) GetMaxRequestsPerSecondOk() (*float64, bool)`

GetMaxRequestsPerSecondOk returns a tuple with the MaxRequestsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestsPerSecond

`func (o *AgentActivityClientSettings) SetMaxRequestsPerSecond(v float64)`

SetMaxRequestsPerSecond sets MaxRequestsPerSecond field to given value.


### GetBurstRequests

`func (o *AgentActivityClientSettings) GetBurstRequests() int32`

GetBurstRequests returns the BurstRequests field if non-nil, zero value otherwise.

### GetBurstRequestsOk

`func (o *AgentActivityClientSettings) GetBurstRequestsOk() (*int32, bool)`

GetBurstRequestsOk returns a tuple with the BurstRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstRequests

`func (o *AgentActivityClientSettings) SetBurstRequests(v int32)`

SetBurstRequests sets BurstRequests field to given value.


### GetCacheSeconds

`func (o *AgentActivityClientSettings) GetCacheSeconds() int32`

GetCacheSeconds returns the CacheSeconds field if non-nil, zero value otherwise.

### GetCacheSecondsOk

`func (o *AgentActivityClientSettings) GetCacheSecondsOk() (*int32, bool)`

GetCacheSecondsOk returns a tuple with the CacheSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSeconds

`func (o *AgentActivityClientSettings) SetCacheSeconds(v int32)`

SetCacheSeconds sets CacheSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


