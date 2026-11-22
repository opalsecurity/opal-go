# McpConfigReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Received** | **int32** | Servers in the report. | 
**Recorded** | **int32** | Servers recorded in the registry. Lower than &#x60;received&#x60; when an entry had neither a name Opal could use nor a URL or command hash. | 

## Methods

### NewMcpConfigReportResult

`func NewMcpConfigReportResult(received int32, recorded int32, ) *McpConfigReportResult`

NewMcpConfigReportResult instantiates a new McpConfigReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConfigReportResultWithDefaults

`func NewMcpConfigReportResultWithDefaults() *McpConfigReportResult`

NewMcpConfigReportResultWithDefaults instantiates a new McpConfigReportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceived

`func (o *McpConfigReportResult) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *McpConfigReportResult) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *McpConfigReportResult) SetReceived(v int32)`

SetReceived sets Received field to given value.


### GetRecorded

`func (o *McpConfigReportResult) GetRecorded() int32`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *McpConfigReportResult) GetRecordedOk() (*int32, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *McpConfigReportResult) SetRecorded(v int32)`

SetRecorded sets Recorded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


