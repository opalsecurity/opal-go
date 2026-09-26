# OpalAccessPathDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **[]string** | Entity IDs along the path from principal to entitlement. | 
**Expiration** | Pointer to **time.Time** | This path&#39;s expiration, if any. | [optional] 

## Methods

### NewOpalAccessPathDetail

`func NewOpalAccessPathDetail(path []string, ) *OpalAccessPathDetail`

NewOpalAccessPathDetail instantiates a new OpalAccessPathDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalAccessPathDetailWithDefaults

`func NewOpalAccessPathDetailWithDefaults() *OpalAccessPathDetail`

NewOpalAccessPathDetailWithDefaults instantiates a new OpalAccessPathDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *OpalAccessPathDetail) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OpalAccessPathDetail) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OpalAccessPathDetail) SetPath(v []string)`

SetPath sets Path field to given value.


### GetExpiration

`func (o *OpalAccessPathDetail) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *OpalAccessPathDetail) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *OpalAccessPathDetail) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *OpalAccessPathDetail) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


