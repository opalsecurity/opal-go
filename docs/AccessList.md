# AccessList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Access**](Access.md) |  | [optional] 

## Methods

### NewAccessList

`func NewAccessList() *AccessList`

NewAccessList instantiates a new AccessList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessListWithDefaults

`func NewAccessListWithDefaults() *AccessList`

NewAccessListWithDefaults instantiates a new AccessList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *AccessList) GetResults() []Access`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AccessList) GetResultsOk() (*[]Access, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AccessList) SetResults(v []Access)`

SetResults sets Results field to given value.

### HasResults

`func (o *AccessList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


