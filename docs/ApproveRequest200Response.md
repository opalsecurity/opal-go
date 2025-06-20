# ApproveRequest200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**Request**](Request.md) |  | [optional] 
**TaskId** | Pointer to **string** | ID of the task created for propagating access | [optional] 

## Methods

### NewApproveRequest200Response

`func NewApproveRequest200Response() *ApproveRequest200Response`

NewApproveRequest200Response instantiates a new ApproveRequest200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveRequest200ResponseWithDefaults

`func NewApproveRequest200ResponseWithDefaults() *ApproveRequest200Response`

NewApproveRequest200ResponseWithDefaults instantiates a new ApproveRequest200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *ApproveRequest200Response) GetRequest() Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ApproveRequest200Response) GetRequestOk() (*Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ApproveRequest200Response) SetRequest(v Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ApproveRequest200Response) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetTaskId

`func (o *ApproveRequest200Response) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ApproveRequest200Response) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ApproveRequest200Response) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ApproveRequest200Response) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


