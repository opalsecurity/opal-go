/*
Opal API

Your Home For Developer Resources.

API version: 1.0
Contact: hello@opal.dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opal

import (
	"encoding/json"
)

// CreateResourceInfo # CreateResourceInfo Object ### Description The `CreateResourceInfo` object is used to store creation info for a resource.  ### Usage Example Use in the `POST Resources` endpoint.
type CreateResourceInfo struct {
	// The name of the remote resource.
	Name string `json:"name"`
	// A description of the remote resource.
	Description *string `json:"description,omitempty"`
	ResourceType ResourceTypeEnum `json:"resource_type"`
	// The ID of the app for the resource.
	AppId string `json:"app_id"`
	RemoteInfo *ResourceRemoteInfo `json:"remote_info,omitempty"`
	// Deprecated - use remote_info instead. The ID of the resource on the remote system. Include only for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details on how to specify this field.
	// Deprecated
	RemoteResourceId *string `json:"remote_resource_id,omitempty"`
	// Deprecated - use remote_info instead.  JSON metadata about the remote resource. Include only for items linked to remote systems. See [this guide](https://docs.opal.dev/reference/end-system-objects) for details on how to specify this field. The required format is dependent on resource_type and should have the following schema: <style type=\"text/css\"> code {max-height:300px !important} </style> ```json {   \"$schema\": \"http://json-schema.org/draft-04/schema#\",   \"title\": \"Resource Metadata\",   \"properties\": {     \"aws_ec2_instance\": {       \"properties\": {         \"instance_id\": {           \"type\": \"string\"         },         \"region\": {           \"type\": \"string\"         }       },       \"required\": [\"instance_id\", \"region\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"AWS EC2 Instance\"     },     \"aws_eks_cluster\": {       \"properties\": {         \"cluster_name\": {           \"type\": \"string\"         },         \"cluster_region\": {           \"type\": \"string\"         },         \"cluster_arn\": {           \"type\": \"string\"         }       },       \"required\": [\"cluster_name\", \"cluster_region\", \"cluster_arn\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"AWS EKS Cluster\"     },     \"aws_rds_instance\": {       \"properties\": {         \"instance_id\": {           \"type\": \"string\"         },         \"engine\": {           \"type\": \"string\"         },         \"region\": {           \"type\": \"string\"         },         \"resource_id\": {           \"type\": \"string\"         },         \"database_name\": {           \"type\": \"string\"         }       },       \"required\": [         \"instance_id\",         \"engine\",         \"region\",         \"resource_id\",         \"database_name\"       ],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"AWS RDS Instance\"     },     \"aws_role\": {       \"properties\": {         \"arn\": {           \"type\": \"string\"         },         \"name\": {           \"type\": \"string\"         }       },       \"required\": [\"arn\", \"name\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"AWS Role\"     },     \"gcp_bucket\": {       \"properties\": {         \"bucket_id\": {           \"type\": \"string\"         }       },       \"required\": [\"bucket_id\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"GCP Bucket\"     },     \"gcp_compute_instance\": {       \"properties\": {         \"instance_id\": {           \"type\": \"string\"         },         \"project_id\": {           \"type\": \"string\"         },         \"zone\": {           \"type\": \"string\"         }       },       \"required\": [\"instance_id\", \"project_id\", \"zone\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"GCP Compute Instance\"     },     \"gcp_folder\": {       \"properties\": {         \"folder_id\": {           \"type\": \"string\"         }       },       \"required\": [\"folder_id\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"GCP Folder\"     },     \"gcp_gke_cluster\": {       \"properties\": {         \"cluster_name\": {           \"type\": \"string\"         }       },       \"required\": [\"cluster_name\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"GCP GKE Cluster\"     },     \"gcp_project\": {       \"properties\": {         \"project_id\": {           \"type\": \"string\"         }       },       \"required\": [\"project_id\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"GCP Project\"     },     \"gcp_sql_instance\": {       \"properties\": {         \"instance_id\": {           \"type\": \"string\"         },         \"project_id\": {           \"type\": \"string\"         }       },       \"required\": [\"instance_id\", \"project_id\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"GCP SQL Instance\"     },     \"git_hub_repo\": {       \"properties\": {         \"org_name\": {           \"type\": \"string\"         },         \"repo_name\": {           \"type\": \"string\"         }       },       \"required\": [\"org_name\", \"repo_name\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"GitHub Repo\"     },     \"okta_directory_app\": {       \"properties\": {         \"app_id\": {           \"type\": \"string\"         },         \"logo_url\": {           \"type\": \"string\"         }       },       \"required\": [\"app_id\", \"logo_url\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"Okta Directory App\"     },     \"okta_directory_role\": {       \"properties\": {         \"role_type\": {           \"type\": \"string\"         },         \"role_id\": {           \"type\": \"string\"         }       },       \"required\": [\"role_type\", \"role_id\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"Okta Directory Role\"     },     \"salesforce_profile\": {       \"properties\": {         \"user_license\": {           \"type\": \"string\"         }       },       \"required\": [\"user_license\"],       \"additionalProperties\": false,       \"type\": \"object\",       \"title\": \"Salesforce Profile\"     }   },   \"additionalProperties\": false,   \"minProperties\": 1,   \"maxProperties\": 1,   \"type\": \"object\" } ```
	// Deprecated
	Metadata *string `json:"metadata,omitempty"`
}

// NewCreateResourceInfo instantiates a new CreateResourceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateResourceInfo(name string, resourceType ResourceTypeEnum, appId string) *CreateResourceInfo {
	this := CreateResourceInfo{}
	this.Name = name
	this.ResourceType = resourceType
	this.AppId = appId
	return &this
}

// NewCreateResourceInfoWithDefaults instantiates a new CreateResourceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateResourceInfoWithDefaults() *CreateResourceInfo {
	this := CreateResourceInfo{}
	return &this
}

// GetName returns the Name field value
func (o *CreateResourceInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateResourceInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateResourceInfo) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateResourceInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateResourceInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateResourceInfo) SetDescription(v string) {
	o.Description = &v
}

// GetResourceType returns the ResourceType field value
func (o *CreateResourceInfo) GetResourceType() ResourceTypeEnum {
	if o == nil {
		var ret ResourceTypeEnum
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CreateResourceInfo) GetResourceTypeOk() (*ResourceTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *CreateResourceInfo) SetResourceType(v ResourceTypeEnum) {
	o.ResourceType = v
}

// GetAppId returns the AppId field value
func (o *CreateResourceInfo) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *CreateResourceInfo) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *CreateResourceInfo) SetAppId(v string) {
	o.AppId = v
}

// GetRemoteInfo returns the RemoteInfo field value if set, zero value otherwise.
func (o *CreateResourceInfo) GetRemoteInfo() ResourceRemoteInfo {
	if o == nil || o.RemoteInfo == nil {
		var ret ResourceRemoteInfo
		return ret
	}
	return *o.RemoteInfo
}

// GetRemoteInfoOk returns a tuple with the RemoteInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceInfo) GetRemoteInfoOk() (*ResourceRemoteInfo, bool) {
	if o == nil || o.RemoteInfo == nil {
		return nil, false
	}
	return o.RemoteInfo, true
}

// HasRemoteInfo returns a boolean if a field has been set.
func (o *CreateResourceInfo) HasRemoteInfo() bool {
	if o != nil && o.RemoteInfo != nil {
		return true
	}

	return false
}

// SetRemoteInfo gets a reference to the given ResourceRemoteInfo and assigns it to the RemoteInfo field.
func (o *CreateResourceInfo) SetRemoteInfo(v ResourceRemoteInfo) {
	o.RemoteInfo = &v
}

// GetRemoteResourceId returns the RemoteResourceId field value if set, zero value otherwise.
// Deprecated
func (o *CreateResourceInfo) GetRemoteResourceId() string {
	if o == nil || o.RemoteResourceId == nil {
		var ret string
		return ret
	}
	return *o.RemoteResourceId
}

// GetRemoteResourceIdOk returns a tuple with the RemoteResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateResourceInfo) GetRemoteResourceIdOk() (*string, bool) {
	if o == nil || o.RemoteResourceId == nil {
		return nil, false
	}
	return o.RemoteResourceId, true
}

// HasRemoteResourceId returns a boolean if a field has been set.
func (o *CreateResourceInfo) HasRemoteResourceId() bool {
	if o != nil && o.RemoteResourceId != nil {
		return true
	}

	return false
}

// SetRemoteResourceId gets a reference to the given string and assigns it to the RemoteResourceId field.
// Deprecated
func (o *CreateResourceInfo) SetRemoteResourceId(v string) {
	o.RemoteResourceId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
// Deprecated
func (o *CreateResourceInfo) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateResourceInfo) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateResourceInfo) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
// Deprecated
func (o *CreateResourceInfo) SetMetadata(v string) {
	o.Metadata = &v
}

func (o CreateResourceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["resource_type"] = o.ResourceType
	}
	if true {
		toSerialize["app_id"] = o.AppId
	}
	if o.RemoteInfo != nil {
		toSerialize["remote_info"] = o.RemoteInfo
	}
	if o.RemoteResourceId != nil {
		toSerialize["remote_resource_id"] = o.RemoteResourceId
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableCreateResourceInfo struct {
	value *CreateResourceInfo
	isSet bool
}

func (v NullableCreateResourceInfo) Get() *CreateResourceInfo {
	return v.value
}

func (v *NullableCreateResourceInfo) Set(val *CreateResourceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateResourceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateResourceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateResourceInfo(val *CreateResourceInfo) *NullableCreateResourceInfo {
	return &NullableCreateResourceInfo{value: val, isSet: true}
}

func (v NullableCreateResourceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateResourceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


