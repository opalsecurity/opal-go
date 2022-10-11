# CreateResourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the remote resource. | 
**Description** | Pointer to **string** | A description of the remote resource. | [optional] 
**ResourceType** | [**ResourceTypeEnum**](ResourceTypeEnum.md) |  | 
**AppId** | **string** | The ID of the app for the resource. | 
**RemoteInfo** | Pointer to [**ResourceRemoteInfo**](ResourceRemoteInfo.md) |  | [optional] 
**RemoteResourceId** | Pointer to **string** | Deprecated - use remote_info instead. The ID of the resource on the remote system. Include only for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details on how to specify this field. | [optional] 
**Metadata** | Pointer to **string** | Deprecated - use remote_info instead.  JSON metadata about the remote resource. Include only for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details on how to specify this field. The required format is dependent on resource_type and should have the following schema: &lt;style type&#x3D;\&quot;text/css\&quot;&gt; code {max-height:300px !important} &lt;/style&gt; &#x60;&#x60;&#x60;json {   \&quot;$schema\&quot;: \&quot;http://json-schema.org/draft-04/schema#\&quot;,   \&quot;title\&quot;: \&quot;Resource Metadata\&quot;,   \&quot;properties\&quot;: {     \&quot;aws_ec2_instance\&quot;: {       \&quot;properties\&quot;: {         \&quot;instance_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;region\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;instance_id\&quot;, \&quot;region\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;AWS EC2 Instance\&quot;     },     \&quot;aws_eks_cluster\&quot;: {       \&quot;properties\&quot;: {         \&quot;cluster_name\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;cluster_region\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;cluster_arn\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;cluster_name\&quot;, \&quot;cluster_region\&quot;, \&quot;cluster_arn\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;AWS EKS Cluster\&quot;     },     \&quot;aws_rds_instance\&quot;: {       \&quot;properties\&quot;: {         \&quot;instance_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;engine\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;region\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;resource_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;database_name\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [         \&quot;instance_id\&quot;,         \&quot;engine\&quot;,         \&quot;region\&quot;,         \&quot;resource_id\&quot;,         \&quot;database_name\&quot;       ],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;AWS RDS Instance\&quot;     },     \&quot;aws_role\&quot;: {       \&quot;properties\&quot;: {         \&quot;arn\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;name\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;arn\&quot;, \&quot;name\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;AWS Role\&quot;     },     \&quot;gcp_bucket\&quot;: {       \&quot;properties\&quot;: {         \&quot;bucket_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;bucket_id\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;GCP Bucket\&quot;     },     \&quot;gcp_compute_instance\&quot;: {       \&quot;properties\&quot;: {         \&quot;instance_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;project_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;zone\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;instance_id\&quot;, \&quot;project_id\&quot;, \&quot;zone\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;GCP Compute Instance\&quot;     },     \&quot;gcp_folder\&quot;: {       \&quot;properties\&quot;: {         \&quot;folder_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;folder_id\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;GCP Folder\&quot;     },     \&quot;gcp_gke_cluster\&quot;: {       \&quot;properties\&quot;: {         \&quot;cluster_name\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;cluster_name\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;GCP GKE Cluster\&quot;     },     \&quot;gcp_project\&quot;: {       \&quot;properties\&quot;: {         \&quot;project_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;project_id\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;GCP Project\&quot;     },     \&quot;gcp_sql_instance\&quot;: {       \&quot;properties\&quot;: {         \&quot;instance_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;project_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;instance_id\&quot;, \&quot;project_id\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;GCP SQL Instance\&quot;     },     \&quot;git_hub_repo\&quot;: {       \&quot;properties\&quot;: {         \&quot;org_name\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;repo_name\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;org_name\&quot;, \&quot;repo_name\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;GitHub Repo\&quot;     },     \&quot;okta_directory_app\&quot;: {       \&quot;properties\&quot;: {         \&quot;app_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;logo_url\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;app_id\&quot;, \&quot;logo_url\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;Okta Directory App\&quot;     },     \&quot;okta_directory_role\&quot;: {       \&quot;properties\&quot;: {         \&quot;role_type\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         },         \&quot;role_id\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;role_type\&quot;, \&quot;role_id\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;Okta Directory Role\&quot;     },     \&quot;salesforce_profile\&quot;: {       \&quot;properties\&quot;: {         \&quot;user_license\&quot;: {           \&quot;type\&quot;: \&quot;string\&quot;         }       },       \&quot;required\&quot;: [\&quot;user_license\&quot;],       \&quot;additionalProperties\&quot;: false,       \&quot;type\&quot;: \&quot;object\&quot;,       \&quot;title\&quot;: \&quot;Salesforce Profile\&quot;     }   },   \&quot;additionalProperties\&quot;: false,   \&quot;minProperties\&quot;: 1,   \&quot;maxProperties\&quot;: 1,   \&quot;type\&quot;: \&quot;object\&quot; } &#x60;&#x60;&#x60; | [optional] 

## Methods

### NewCreateResourceInfo

`func NewCreateResourceInfo(name string, resourceType ResourceTypeEnum, appId string, ) *CreateResourceInfo`

NewCreateResourceInfo instantiates a new CreateResourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceInfoWithDefaults

`func NewCreateResourceInfoWithDefaults() *CreateResourceInfo`

NewCreateResourceInfoWithDefaults instantiates a new CreateResourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateResourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResourceInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateResourceInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateResourceInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateResourceInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateResourceInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceType

`func (o *CreateResourceInfo) GetResourceType() ResourceTypeEnum`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateResourceInfo) GetResourceTypeOk() (*ResourceTypeEnum, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateResourceInfo) SetResourceType(v ResourceTypeEnum)`

SetResourceType sets ResourceType field to given value.


### GetAppId

`func (o *CreateResourceInfo) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateResourceInfo) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateResourceInfo) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetRemoteInfo

`func (o *CreateResourceInfo) GetRemoteInfo() ResourceRemoteInfo`

GetRemoteInfo returns the RemoteInfo field if non-nil, zero value otherwise.

### GetRemoteInfoOk

`func (o *CreateResourceInfo) GetRemoteInfoOk() (*ResourceRemoteInfo, bool)`

GetRemoteInfoOk returns a tuple with the RemoteInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteInfo

`func (o *CreateResourceInfo) SetRemoteInfo(v ResourceRemoteInfo)`

SetRemoteInfo sets RemoteInfo field to given value.

### HasRemoteInfo

`func (o *CreateResourceInfo) HasRemoteInfo() bool`

HasRemoteInfo returns a boolean if a field has been set.

### GetRemoteResourceId

`func (o *CreateResourceInfo) GetRemoteResourceId() string`

GetRemoteResourceId returns the RemoteResourceId field if non-nil, zero value otherwise.

### GetRemoteResourceIdOk

`func (o *CreateResourceInfo) GetRemoteResourceIdOk() (*string, bool)`

GetRemoteResourceIdOk returns a tuple with the RemoteResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteResourceId

`func (o *CreateResourceInfo) SetRemoteResourceId(v string)`

SetRemoteResourceId sets RemoteResourceId field to given value.

### HasRemoteResourceId

`func (o *CreateResourceInfo) HasRemoteResourceId() bool`

HasRemoteResourceId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateResourceInfo) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateResourceInfo) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateResourceInfo) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateResourceInfo) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


