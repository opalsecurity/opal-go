# EntityNameFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StringMatchType** | [**StringMatchType**](StringMatchType.md) |  | 
**String** | **string** | The string value to match against the entity name. | 

## Methods

### NewEntityNameFilter

`func NewEntityNameFilter(stringMatchType StringMatchType, string_ string, ) *EntityNameFilter`

NewEntityNameFilter instantiates a new EntityNameFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityNameFilterWithDefaults

`func NewEntityNameFilterWithDefaults() *EntityNameFilter`

NewEntityNameFilterWithDefaults instantiates a new EntityNameFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStringMatchType

`func (o *EntityNameFilter) GetStringMatchType() StringMatchType`

GetStringMatchType returns the StringMatchType field if non-nil, zero value otherwise.

### GetStringMatchTypeOk

`func (o *EntityNameFilter) GetStringMatchTypeOk() (*StringMatchType, bool)`

GetStringMatchTypeOk returns a tuple with the StringMatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringMatchType

`func (o *EntityNameFilter) SetStringMatchType(v StringMatchType)`

SetStringMatchType sets StringMatchType field to given value.


### GetString

`func (o *EntityNameFilter) GetString() string`

GetString returns the String field if non-nil, zero value otherwise.

### GetStringOk

`func (o *EntityNameFilter) GetStringOk() (*string, bool)`

GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetString

`func (o *EntityNameFilter) SetString(v string)`

SetString sets String field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


