# SyncTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the sync task. | 
**CompletedAt** | **time.Time** | The time when the sync task was completed. | 

## Methods

### NewSyncTask

`func NewSyncTask(id string, completedAt time.Time, ) *SyncTask`

NewSyncTask instantiates a new SyncTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncTaskWithDefaults

`func NewSyncTaskWithDefaults() *SyncTask`

NewSyncTaskWithDefaults instantiates a new SyncTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyncTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncTask) SetId(v string)`

SetId sets Id field to given value.


### GetCompletedAt

`func (o *SyncTask) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *SyncTask) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *SyncTask) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


