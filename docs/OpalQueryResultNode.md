# OpalQueryResultNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The entity&#39;s unique identifier. | 
**Name** | **string** | The display name of the entity. | 
**EntityType** | **string** | The top-level entity type. | 
**EntityItemType** | [**EntityItemTypeEnum**](EntityItemTypeEnum.md) |  | 

## Methods

### NewOpalQueryResultNode

`func NewOpalQueryResultNode(id string, name string, entityType string, entityItemType EntityItemTypeEnum, ) *OpalQueryResultNode`

NewOpalQueryResultNode instantiates a new OpalQueryResultNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpalQueryResultNodeWithDefaults

`func NewOpalQueryResultNodeWithDefaults() *OpalQueryResultNode`

NewOpalQueryResultNodeWithDefaults instantiates a new OpalQueryResultNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpalQueryResultNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpalQueryResultNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpalQueryResultNode) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OpalQueryResultNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpalQueryResultNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpalQueryResultNode) SetName(v string)`

SetName sets Name field to given value.


### GetEntityType

`func (o *OpalQueryResultNode) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *OpalQueryResultNode) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *OpalQueryResultNode) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetEntityItemType

`func (o *OpalQueryResultNode) GetEntityItemType() EntityItemTypeEnum`

GetEntityItemType returns the EntityItemType field if non-nil, zero value otherwise.

### GetEntityItemTypeOk

`func (o *OpalQueryResultNode) GetEntityItemTypeOk() (*EntityItemTypeEnum, bool)`

GetEntityItemTypeOk returns a tuple with the EntityItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityItemType

`func (o *OpalQueryResultNode) SetEntityItemType(v EntityItemTypeEnum)`

SetEntityItemType sets EntityItemType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


