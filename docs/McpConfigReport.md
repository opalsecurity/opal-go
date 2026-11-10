# McpConfigReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | [**[]McpConfigServer**](McpConfigServer.md) |  | 
**Reporter** | Pointer to [**McpConfigReporter**](McpConfigReporter.md) |  | [optional] 

## Methods

### NewMcpConfigReport

`func NewMcpConfigReport(servers []McpConfigServer, ) *McpConfigReport`

NewMcpConfigReport instantiates a new McpConfigReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConfigReportWithDefaults

`func NewMcpConfigReportWithDefaults() *McpConfigReport`

NewMcpConfigReportWithDefaults instantiates a new McpConfigReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *McpConfigReport) GetServers() []McpConfigServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *McpConfigReport) GetServersOk() (*[]McpConfigServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *McpConfigReport) SetServers(v []McpConfigServer)`

SetServers sets Servers field to given value.


### GetReporter

`func (o *McpConfigReport) GetReporter() McpConfigReporter`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *McpConfigReport) GetReporterOk() (*McpConfigReporter, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *McpConfigReport) SetReporter(v McpConfigReporter)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *McpConfigReport) HasReporter() bool`

HasReporter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


