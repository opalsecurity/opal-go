# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The ID of the event. | 
**ActorUserId** | **string** | The ID of the actor user. | 
**ActorName** | **interface{}** |  | 
**ActorEmail** | Pointer to **string** | The email of the actor user. | [optional] 
**EventType** | **string** | The event type. | 
**CreatedAt** | **time.Time** | The day and time the event was created. | 
**ActorIpAddress** | Pointer to **string** | The IP address of the event actor. | [optional] 
**ApiTokenName** | Pointer to **string** | The name of the API token used to create the event. | [optional] 
**ApiTokenPreview** | Pointer to **string** | The preview of the API token used to create the event. | [optional] 
**SubEvents** | Pointer to [**[]SubEvent**](SubEvent.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent(eventId string, actorUserId string, actorName interface{}, eventType string, createdAt time.Time, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *Event) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Event) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Event) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetActorUserId

`func (o *Event) GetActorUserId() string`

GetActorUserId returns the ActorUserId field if non-nil, zero value otherwise.

### GetActorUserIdOk

`func (o *Event) GetActorUserIdOk() (*string, bool)`

GetActorUserIdOk returns a tuple with the ActorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUserId

`func (o *Event) SetActorUserId(v string)`

SetActorUserId sets ActorUserId field to given value.


### GetActorName

`func (o *Event) GetActorName() interface{}`

GetActorName returns the ActorName field if non-nil, zero value otherwise.

### GetActorNameOk

`func (o *Event) GetActorNameOk() (*interface{}, bool)`

GetActorNameOk returns a tuple with the ActorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorName

`func (o *Event) SetActorName(v interface{})`

SetActorName sets ActorName field to given value.


### SetActorNameNil

`func (o *Event) SetActorNameNil(b bool)`

 SetActorNameNil sets the value for ActorName to be an explicit nil

### UnsetActorName
`func (o *Event) UnsetActorName()`

UnsetActorName ensures that no value is present for ActorName, not even an explicit nil
### GetActorEmail

`func (o *Event) GetActorEmail() string`

GetActorEmail returns the ActorEmail field if non-nil, zero value otherwise.

### GetActorEmailOk

`func (o *Event) GetActorEmailOk() (*string, bool)`

GetActorEmailOk returns a tuple with the ActorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorEmail

`func (o *Event) SetActorEmail(v string)`

SetActorEmail sets ActorEmail field to given value.

### HasActorEmail

`func (o *Event) HasActorEmail() bool`

HasActorEmail returns a boolean if a field has been set.

### GetEventType

`func (o *Event) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Event) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Event) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetCreatedAt

`func (o *Event) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Event) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Event) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActorIpAddress

`func (o *Event) GetActorIpAddress() string`

GetActorIpAddress returns the ActorIpAddress field if non-nil, zero value otherwise.

### GetActorIpAddressOk

`func (o *Event) GetActorIpAddressOk() (*string, bool)`

GetActorIpAddressOk returns a tuple with the ActorIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorIpAddress

`func (o *Event) SetActorIpAddress(v string)`

SetActorIpAddress sets ActorIpAddress field to given value.

### HasActorIpAddress

`func (o *Event) HasActorIpAddress() bool`

HasActorIpAddress returns a boolean if a field has been set.

### GetApiTokenName

`func (o *Event) GetApiTokenName() string`

GetApiTokenName returns the ApiTokenName field if non-nil, zero value otherwise.

### GetApiTokenNameOk

`func (o *Event) GetApiTokenNameOk() (*string, bool)`

GetApiTokenNameOk returns a tuple with the ApiTokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenName

`func (o *Event) SetApiTokenName(v string)`

SetApiTokenName sets ApiTokenName field to given value.

### HasApiTokenName

`func (o *Event) HasApiTokenName() bool`

HasApiTokenName returns a boolean if a field has been set.

### GetApiTokenPreview

`func (o *Event) GetApiTokenPreview() string`

GetApiTokenPreview returns the ApiTokenPreview field if non-nil, zero value otherwise.

### GetApiTokenPreviewOk

`func (o *Event) GetApiTokenPreviewOk() (*string, bool)`

GetApiTokenPreviewOk returns a tuple with the ApiTokenPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenPreview

`func (o *Event) SetApiTokenPreview(v string)`

SetApiTokenPreview sets ApiTokenPreview field to given value.

### HasApiTokenPreview

`func (o *Event) HasApiTokenPreview() bool`

HasApiTokenPreview returns a boolean if a field has been set.

### GetSubEvents

`func (o *Event) GetSubEvents() []SubEvent`

GetSubEvents returns the SubEvents field if non-nil, zero value otherwise.

### GetSubEventsOk

`func (o *Event) GetSubEventsOk() (*[]SubEvent, bool)`

GetSubEventsOk returns a tuple with the SubEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubEvents

`func (o *Event) SetSubEvents(v []SubEvent)`

SetSubEvents sets SubEvents field to given value.

### HasSubEvents

`func (o *Event) HasSubEvents() bool`

HasSubEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


