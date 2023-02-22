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

// checks if the ResourceRemoteInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRemoteInfo{}

// ResourceRemoteInfo Remote info that is required for the creation of remote resources. Providing this will allow you to omit deprecated remote_id and metadata fields.
type ResourceRemoteInfo struct {
	AwsIamRole *ResourceRemoteInfoAwsIamRole `json:"aws_iam_role,omitempty"`
	AwsEc2Instance *ResourceRemoteInfoAwsEc2Instance `json:"aws_ec2_instance,omitempty"`
	AwsRdsInstance *ResourceRemoteInfoAwsRdsInstance `json:"aws_rds_instance,omitempty"`
	AwsEksCluster *ResourceRemoteInfoAwsEksCluster `json:"aws_eks_cluster,omitempty"`
	GithubRepo *ResourceRemoteInfoGithubRepo `json:"github_repo,omitempty"`
	GitlabProject *ResourceRemoteInfoGitlabProject `json:"gitlab_project,omitempty"`
	OktaApp *ResourceRemoteInfoOktaApp `json:"okta_app,omitempty"`
	OktaStandardRole *ResourceRemoteInfoOktaStandardRole `json:"okta_standard_role,omitempty"`
	OktaCustomRole *ResourceRemoteInfoOktaCustomRole `json:"okta_custom_role,omitempty"`
	TeleportRole *ResourceRemoteInfoTeleportRole `json:"teleport_role,omitempty"`
}

// NewResourceRemoteInfo instantiates a new ResourceRemoteInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRemoteInfo() *ResourceRemoteInfo {
	this := ResourceRemoteInfo{}
	return &this
}

// NewResourceRemoteInfoWithDefaults instantiates a new ResourceRemoteInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRemoteInfoWithDefaults() *ResourceRemoteInfo {
	this := ResourceRemoteInfo{}
	return &this
}

// GetAwsIamRole returns the AwsIamRole field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetAwsIamRole() ResourceRemoteInfoAwsIamRole {
	if o == nil || IsNil(o.AwsIamRole) {
		var ret ResourceRemoteInfoAwsIamRole
		return ret
	}
	return *o.AwsIamRole
}

// GetAwsIamRoleOk returns a tuple with the AwsIamRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetAwsIamRoleOk() (*ResourceRemoteInfoAwsIamRole, bool) {
	if o == nil || IsNil(o.AwsIamRole) {
		return nil, false
	}
	return o.AwsIamRole, true
}

// HasAwsIamRole returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasAwsIamRole() bool {
	if o != nil && !IsNil(o.AwsIamRole) {
		return true
	}

	return false
}

// SetAwsIamRole gets a reference to the given ResourceRemoteInfoAwsIamRole and assigns it to the AwsIamRole field.
func (o *ResourceRemoteInfo) SetAwsIamRole(v ResourceRemoteInfoAwsIamRole) {
	o.AwsIamRole = &v
}

// GetAwsEc2Instance returns the AwsEc2Instance field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetAwsEc2Instance() ResourceRemoteInfoAwsEc2Instance {
	if o == nil || IsNil(o.AwsEc2Instance) {
		var ret ResourceRemoteInfoAwsEc2Instance
		return ret
	}
	return *o.AwsEc2Instance
}

// GetAwsEc2InstanceOk returns a tuple with the AwsEc2Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetAwsEc2InstanceOk() (*ResourceRemoteInfoAwsEc2Instance, bool) {
	if o == nil || IsNil(o.AwsEc2Instance) {
		return nil, false
	}
	return o.AwsEc2Instance, true
}

// HasAwsEc2Instance returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasAwsEc2Instance() bool {
	if o != nil && !IsNil(o.AwsEc2Instance) {
		return true
	}

	return false
}

// SetAwsEc2Instance gets a reference to the given ResourceRemoteInfoAwsEc2Instance and assigns it to the AwsEc2Instance field.
func (o *ResourceRemoteInfo) SetAwsEc2Instance(v ResourceRemoteInfoAwsEc2Instance) {
	o.AwsEc2Instance = &v
}

// GetAwsRdsInstance returns the AwsRdsInstance field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetAwsRdsInstance() ResourceRemoteInfoAwsRdsInstance {
	if o == nil || IsNil(o.AwsRdsInstance) {
		var ret ResourceRemoteInfoAwsRdsInstance
		return ret
	}
	return *o.AwsRdsInstance
}

// GetAwsRdsInstanceOk returns a tuple with the AwsRdsInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetAwsRdsInstanceOk() (*ResourceRemoteInfoAwsRdsInstance, bool) {
	if o == nil || IsNil(o.AwsRdsInstance) {
		return nil, false
	}
	return o.AwsRdsInstance, true
}

// HasAwsRdsInstance returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasAwsRdsInstance() bool {
	if o != nil && !IsNil(o.AwsRdsInstance) {
		return true
	}

	return false
}

// SetAwsRdsInstance gets a reference to the given ResourceRemoteInfoAwsRdsInstance and assigns it to the AwsRdsInstance field.
func (o *ResourceRemoteInfo) SetAwsRdsInstance(v ResourceRemoteInfoAwsRdsInstance) {
	o.AwsRdsInstance = &v
}

// GetAwsEksCluster returns the AwsEksCluster field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetAwsEksCluster() ResourceRemoteInfoAwsEksCluster {
	if o == nil || IsNil(o.AwsEksCluster) {
		var ret ResourceRemoteInfoAwsEksCluster
		return ret
	}
	return *o.AwsEksCluster
}

// GetAwsEksClusterOk returns a tuple with the AwsEksCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetAwsEksClusterOk() (*ResourceRemoteInfoAwsEksCluster, bool) {
	if o == nil || IsNil(o.AwsEksCluster) {
		return nil, false
	}
	return o.AwsEksCluster, true
}

// HasAwsEksCluster returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasAwsEksCluster() bool {
	if o != nil && !IsNil(o.AwsEksCluster) {
		return true
	}

	return false
}

// SetAwsEksCluster gets a reference to the given ResourceRemoteInfoAwsEksCluster and assigns it to the AwsEksCluster field.
func (o *ResourceRemoteInfo) SetAwsEksCluster(v ResourceRemoteInfoAwsEksCluster) {
	o.AwsEksCluster = &v
}

// GetGithubRepo returns the GithubRepo field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetGithubRepo() ResourceRemoteInfoGithubRepo {
	if o == nil || IsNil(o.GithubRepo) {
		var ret ResourceRemoteInfoGithubRepo
		return ret
	}
	return *o.GithubRepo
}

// GetGithubRepoOk returns a tuple with the GithubRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetGithubRepoOk() (*ResourceRemoteInfoGithubRepo, bool) {
	if o == nil || IsNil(o.GithubRepo) {
		return nil, false
	}
	return o.GithubRepo, true
}

// HasGithubRepo returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasGithubRepo() bool {
	if o != nil && !IsNil(o.GithubRepo) {
		return true
	}

	return false
}

// SetGithubRepo gets a reference to the given ResourceRemoteInfoGithubRepo and assigns it to the GithubRepo field.
func (o *ResourceRemoteInfo) SetGithubRepo(v ResourceRemoteInfoGithubRepo) {
	o.GithubRepo = &v
}

// GetGitlabProject returns the GitlabProject field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetGitlabProject() ResourceRemoteInfoGitlabProject {
	if o == nil || IsNil(o.GitlabProject) {
		var ret ResourceRemoteInfoGitlabProject
		return ret
	}
	return *o.GitlabProject
}

// GetGitlabProjectOk returns a tuple with the GitlabProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetGitlabProjectOk() (*ResourceRemoteInfoGitlabProject, bool) {
	if o == nil || IsNil(o.GitlabProject) {
		return nil, false
	}
	return o.GitlabProject, true
}

// HasGitlabProject returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasGitlabProject() bool {
	if o != nil && !IsNil(o.GitlabProject) {
		return true
	}

	return false
}

// SetGitlabProject gets a reference to the given ResourceRemoteInfoGitlabProject and assigns it to the GitlabProject field.
func (o *ResourceRemoteInfo) SetGitlabProject(v ResourceRemoteInfoGitlabProject) {
	o.GitlabProject = &v
}

// GetOktaApp returns the OktaApp field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetOktaApp() ResourceRemoteInfoOktaApp {
	if o == nil || IsNil(o.OktaApp) {
		var ret ResourceRemoteInfoOktaApp
		return ret
	}
	return *o.OktaApp
}

// GetOktaAppOk returns a tuple with the OktaApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetOktaAppOk() (*ResourceRemoteInfoOktaApp, bool) {
	if o == nil || IsNil(o.OktaApp) {
		return nil, false
	}
	return o.OktaApp, true
}

// HasOktaApp returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasOktaApp() bool {
	if o != nil && !IsNil(o.OktaApp) {
		return true
	}

	return false
}

// SetOktaApp gets a reference to the given ResourceRemoteInfoOktaApp and assigns it to the OktaApp field.
func (o *ResourceRemoteInfo) SetOktaApp(v ResourceRemoteInfoOktaApp) {
	o.OktaApp = &v
}

// GetOktaStandardRole returns the OktaStandardRole field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetOktaStandardRole() ResourceRemoteInfoOktaStandardRole {
	if o == nil || IsNil(o.OktaStandardRole) {
		var ret ResourceRemoteInfoOktaStandardRole
		return ret
	}
	return *o.OktaStandardRole
}

// GetOktaStandardRoleOk returns a tuple with the OktaStandardRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetOktaStandardRoleOk() (*ResourceRemoteInfoOktaStandardRole, bool) {
	if o == nil || IsNil(o.OktaStandardRole) {
		return nil, false
	}
	return o.OktaStandardRole, true
}

// HasOktaStandardRole returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasOktaStandardRole() bool {
	if o != nil && !IsNil(o.OktaStandardRole) {
		return true
	}

	return false
}

// SetOktaStandardRole gets a reference to the given ResourceRemoteInfoOktaStandardRole and assigns it to the OktaStandardRole field.
func (o *ResourceRemoteInfo) SetOktaStandardRole(v ResourceRemoteInfoOktaStandardRole) {
	o.OktaStandardRole = &v
}

// GetOktaCustomRole returns the OktaCustomRole field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetOktaCustomRole() ResourceRemoteInfoOktaCustomRole {
	if o == nil || IsNil(o.OktaCustomRole) {
		var ret ResourceRemoteInfoOktaCustomRole
		return ret
	}
	return *o.OktaCustomRole
}

// GetOktaCustomRoleOk returns a tuple with the OktaCustomRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetOktaCustomRoleOk() (*ResourceRemoteInfoOktaCustomRole, bool) {
	if o == nil || IsNil(o.OktaCustomRole) {
		return nil, false
	}
	return o.OktaCustomRole, true
}

// HasOktaCustomRole returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasOktaCustomRole() bool {
	if o != nil && !IsNil(o.OktaCustomRole) {
		return true
	}

	return false
}

// SetOktaCustomRole gets a reference to the given ResourceRemoteInfoOktaCustomRole and assigns it to the OktaCustomRole field.
func (o *ResourceRemoteInfo) SetOktaCustomRole(v ResourceRemoteInfoOktaCustomRole) {
	o.OktaCustomRole = &v
}

// GetTeleportRole returns the TeleportRole field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetTeleportRole() ResourceRemoteInfoTeleportRole {
	if o == nil || IsNil(o.TeleportRole) {
		var ret ResourceRemoteInfoTeleportRole
		return ret
	}
	return *o.TeleportRole
}

// GetTeleportRoleOk returns a tuple with the TeleportRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetTeleportRoleOk() (*ResourceRemoteInfoTeleportRole, bool) {
	if o == nil || IsNil(o.TeleportRole) {
		return nil, false
	}
	return o.TeleportRole, true
}

// HasTeleportRole returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasTeleportRole() bool {
	if o != nil && !IsNil(o.TeleportRole) {
		return true
	}

	return false
}

// SetTeleportRole gets a reference to the given ResourceRemoteInfoTeleportRole and assigns it to the TeleportRole field.
func (o *ResourceRemoteInfo) SetTeleportRole(v ResourceRemoteInfoTeleportRole) {
	o.TeleportRole = &v
}

func (o ResourceRemoteInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRemoteInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsIamRole) {
		toSerialize["aws_iam_role"] = o.AwsIamRole
	}
	if !IsNil(o.AwsEc2Instance) {
		toSerialize["aws_ec2_instance"] = o.AwsEc2Instance
	}
	if !IsNil(o.AwsRdsInstance) {
		toSerialize["aws_rds_instance"] = o.AwsRdsInstance
	}
	if !IsNil(o.AwsEksCluster) {
		toSerialize["aws_eks_cluster"] = o.AwsEksCluster
	}
	if !IsNil(o.GithubRepo) {
		toSerialize["github_repo"] = o.GithubRepo
	}
	if !IsNil(o.GitlabProject) {
		toSerialize["gitlab_project"] = o.GitlabProject
	}
	if !IsNil(o.OktaApp) {
		toSerialize["okta_app"] = o.OktaApp
	}
	if !IsNil(o.OktaStandardRole) {
		toSerialize["okta_standard_role"] = o.OktaStandardRole
	}
	if !IsNil(o.OktaCustomRole) {
		toSerialize["okta_custom_role"] = o.OktaCustomRole
	}
	if !IsNil(o.TeleportRole) {
		toSerialize["teleport_role"] = o.TeleportRole
	}
	return toSerialize, nil
}

type NullableResourceRemoteInfo struct {
	value *ResourceRemoteInfo
	isSet bool
}

func (v NullableResourceRemoteInfo) Get() *ResourceRemoteInfo {
	return v.value
}

func (v *NullableResourceRemoteInfo) Set(val *ResourceRemoteInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRemoteInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRemoteInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRemoteInfo(val *ResourceRemoteInfo) *NullableResourceRemoteInfo {
	return &NullableResourceRemoteInfo{value: val, isSet: true}
}

func (v NullableResourceRemoteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRemoteInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


