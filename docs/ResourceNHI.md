# ResourceNHI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The ID of the resource. | 
**NonHumanIdentityId** | **string** | The resource ID of the non-human identity. | 
**AccessLevel** | Pointer to [**ResourceAccessLevel**](ResourceAccessLevel.md) |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** | The day and time the non-human identity&#39;s access will expire. | [optional] 

## Methods

### NewResourceNHI

`func NewResourceNHI(resourceId string, nonHumanIdentityId string, ) *ResourceNHI`

NewResourceNHI instantiates a new ResourceNHI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceNHIWithDefaults

`func NewResourceNHIWithDefaults() *ResourceNHI`

NewResourceNHIWithDefaults instantiates a new ResourceNHI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *ResourceNHI) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceNHI) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceNHI) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetNonHumanIdentityId

`func (o *ResourceNHI) GetNonHumanIdentityId() string`

GetNonHumanIdentityId returns the NonHumanIdentityId field if non-nil, zero value otherwise.

### GetNonHumanIdentityIdOk

`func (o *ResourceNHI) GetNonHumanIdentityIdOk() (*string, bool)`

GetNonHumanIdentityIdOk returns a tuple with the NonHumanIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonHumanIdentityId

`func (o *ResourceNHI) SetNonHumanIdentityId(v string)`

SetNonHumanIdentityId sets NonHumanIdentityId field to given value.


### GetAccessLevel

`func (o *ResourceNHI) GetAccessLevel() ResourceAccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ResourceNHI) GetAccessLevelOk() (*ResourceAccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ResourceNHI) SetAccessLevel(v ResourceAccessLevel)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *ResourceNHI) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ResourceNHI) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ResourceNHI) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ResourceNHI) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ResourceNHI) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


