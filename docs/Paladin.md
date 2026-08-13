# Paladin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaladinId** | **string** | The ID of the Paladin. Use this value as a reviewer in a request configuration&#39;s service_user_ids. | 
**Name** | **string** | The name of the Paladin. | 

## Methods

### NewPaladin

`func NewPaladin(paladinId string, name string, ) *Paladin`

NewPaladin instantiates a new Paladin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaladinWithDefaults

`func NewPaladinWithDefaults() *Paladin`

NewPaladinWithDefaults instantiates a new Paladin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaladinId

`func (o *Paladin) GetPaladinId() string`

GetPaladinId returns the PaladinId field if non-nil, zero value otherwise.

### GetPaladinIdOk

`func (o *Paladin) GetPaladinIdOk() (*string, bool)`

GetPaladinIdOk returns a tuple with the PaladinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaladinId

`func (o *Paladin) SetPaladinId(v string)`

SetPaladinId sets PaladinId field to given value.


### GetName

`func (o *Paladin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Paladin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Paladin) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


