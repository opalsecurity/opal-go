# EventStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventStreamId** | **string** | The ID of the event stream. | 
**Connection** | [**EventStreamConnection**](EventStreamConnection.md) |  | 

## Methods

### NewEventStream

`func NewEventStream(eventStreamId string, connection EventStreamConnection, ) *EventStream`

NewEventStream instantiates a new EventStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStreamWithDefaults

`func NewEventStreamWithDefaults() *EventStream`

NewEventStreamWithDefaults instantiates a new EventStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventStreamId

`func (o *EventStream) GetEventStreamId() string`

GetEventStreamId returns the EventStreamId field if non-nil, zero value otherwise.

### GetEventStreamIdOk

`func (o *EventStream) GetEventStreamIdOk() (*string, bool)`

GetEventStreamIdOk returns a tuple with the EventStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStreamId

`func (o *EventStream) SetEventStreamId(v string)`

SetEventStreamId sets EventStreamId field to given value.


### GetConnection

`func (o *EventStream) GetConnection() EventStreamConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *EventStream) GetConnectionOk() (*EventStreamConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *EventStream) SetConnection(v EventStreamConnection)`

SetConnection sets Connection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


