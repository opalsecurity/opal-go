# EntityDescriptionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**String** | **string** | The substring to match against the entity description. | 
**Not** | Pointer to **bool** | Invert the match — return resources/groups whose description does NOT contain the substring. | [optional] 

## Methods

### NewEntityDescriptionFilter

`func NewEntityDescriptionFilter(string_ string, ) *EntityDescriptionFilter`

NewEntityDescriptionFilter instantiates a new EntityDescriptionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityDescriptionFilterWithDefaults

`func NewEntityDescriptionFilterWithDefaults() *EntityDescriptionFilter`

NewEntityDescriptionFilterWithDefaults instantiates a new EntityDescriptionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetString

`func (o *EntityDescriptionFilter) GetString() string`

GetString returns the String field if non-nil, zero value otherwise.

### GetStringOk

`func (o *EntityDescriptionFilter) GetStringOk() (*string, bool)`

GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetString

`func (o *EntityDescriptionFilter) SetString(v string)`

SetString sets String field to given value.


### GetNot

`func (o *EntityDescriptionFilter) GetNot() bool`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *EntityDescriptionFilter) GetNotOk() (*bool, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *EntityDescriptionFilter) SetNot(v bool)`

SetNot sets Not field to given value.

### HasNot

`func (o *EntityDescriptionFilter) HasNot() bool`

HasNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


