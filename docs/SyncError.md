# SyncError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstSeen** | **time.Time** | The time when this error was first seen. | 
**LastSeen** | **time.Time** | The time when this error was most recently seen. | 
**ErrorMessage** | **string** | The error message associated with the sync error. | 
**AppId** | Pointer to **string** | The ID of the app that the error occured for. | [optional] 

## Methods

### NewSyncError

`func NewSyncError(firstSeen time.Time, lastSeen time.Time, errorMessage string, ) *SyncError`

NewSyncError instantiates a new SyncError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncErrorWithDefaults

`func NewSyncErrorWithDefaults() *SyncError`

NewSyncErrorWithDefaults instantiates a new SyncError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstSeen

`func (o *SyncError) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *SyncError) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *SyncError) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.


### GetLastSeen

`func (o *SyncError) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *SyncError) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *SyncError) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.


### GetErrorMessage

`func (o *SyncError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SyncError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SyncError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetAppId

`func (o *SyncError) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SyncError) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SyncError) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *SyncError) HasAppId() bool`

HasAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


