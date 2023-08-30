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

// ResourceRemoteInfo Information that defines the remote resource. This replaces the deprecated remote_id and metadata fields.
type ResourceRemoteInfo struct {
	AwsAccount *ResourceRemoteInfoAwsAccount `json:"aws_account,omitempty"`
	AwsPermissionSet *ResourceRemoteInfoAwsPermissionSet `json:"aws_permission_set,omitempty"`
	AwsIamRole *ResourceRemoteInfoAwsIamRole `json:"aws_iam_role,omitempty"`
	AwsEc2Instance *ResourceRemoteInfoAwsEc2Instance `json:"aws_ec2_instance,omitempty"`
	AwsRdsInstance *ResourceRemoteInfoAwsRdsInstance `json:"aws_rds_instance,omitempty"`
	AwsEksCluster *ResourceRemoteInfoAwsEksCluster `json:"aws_eks_cluster,omitempty"`
	GcpBucket *ResourceRemoteInfoGcpBucket `json:"gcp_bucket,omitempty"`
	GcpComputeInstance *ResourceRemoteInfoGcpComputeInstance `json:"gcp_compute_instance,omitempty"`
	GcpFolder *ResourceRemoteInfoGcpFolder `json:"gcp_folder,omitempty"`
	GcpGkeCluster *ResourceRemoteInfoGcpGkeCluster `json:"gcp_gke_cluster,omitempty"`
	GcpProject *ResourceRemoteInfoGcpProject `json:"gcp_project,omitempty"`
	GcpSqlInstance *ResourceRemoteInfoGcpSqlInstance `json:"gcp_sql_instance,omitempty"`
	GithubRepo *ResourceRemoteInfoGithubRepo `json:"github_repo,omitempty"`
	GitlabProject *ResourceRemoteInfoGitlabProject `json:"gitlab_project,omitempty"`
	OktaApp *ResourceRemoteInfoOktaApp `json:"okta_app,omitempty"`
	OktaStandardRole *ResourceRemoteInfoOktaStandardRole `json:"okta_standard_role,omitempty"`
	OktaCustomRole *ResourceRemoteInfoOktaCustomRole `json:"okta_custom_role,omitempty"`
	PagerdutyRole *ResourceRemoteInfoPagerdutyRole `json:"pagerduty_role,omitempty"`
	SalesforcePermissionSet *ResourceRemoteInfoSalesforcePermissionSet `json:"salesforce_permission_set,omitempty"`
	SalesforceProfile *ResourceRemoteInfoSalesforceProfile `json:"salesforce_profile,omitempty"`
	SalesforceRole *ResourceRemoteInfoSalesforceRole `json:"salesforce_role,omitempty"`
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

// GetAwsAccount returns the AwsAccount field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetAwsAccount() ResourceRemoteInfoAwsAccount {
	if o == nil || IsNil(o.AwsAccount) {
		var ret ResourceRemoteInfoAwsAccount
		return ret
	}
	return *o.AwsAccount
}

// GetAwsAccountOk returns a tuple with the AwsAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetAwsAccountOk() (*ResourceRemoteInfoAwsAccount, bool) {
	if o == nil || IsNil(o.AwsAccount) {
		return nil, false
	}
	return o.AwsAccount, true
}

// HasAwsAccount returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasAwsAccount() bool {
	if o != nil && !IsNil(o.AwsAccount) {
		return true
	}

	return false
}

// SetAwsAccount gets a reference to the given ResourceRemoteInfoAwsAccount and assigns it to the AwsAccount field.
func (o *ResourceRemoteInfo) SetAwsAccount(v ResourceRemoteInfoAwsAccount) {
	o.AwsAccount = &v
}

// GetAwsPermissionSet returns the AwsPermissionSet field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetAwsPermissionSet() ResourceRemoteInfoAwsPermissionSet {
	if o == nil || IsNil(o.AwsPermissionSet) {
		var ret ResourceRemoteInfoAwsPermissionSet
		return ret
	}
	return *o.AwsPermissionSet
}

// GetAwsPermissionSetOk returns a tuple with the AwsPermissionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetAwsPermissionSetOk() (*ResourceRemoteInfoAwsPermissionSet, bool) {
	if o == nil || IsNil(o.AwsPermissionSet) {
		return nil, false
	}
	return o.AwsPermissionSet, true
}

// HasAwsPermissionSet returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasAwsPermissionSet() bool {
	if o != nil && !IsNil(o.AwsPermissionSet) {
		return true
	}

	return false
}

// SetAwsPermissionSet gets a reference to the given ResourceRemoteInfoAwsPermissionSet and assigns it to the AwsPermissionSet field.
func (o *ResourceRemoteInfo) SetAwsPermissionSet(v ResourceRemoteInfoAwsPermissionSet) {
	o.AwsPermissionSet = &v
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

// GetGcpBucket returns the GcpBucket field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetGcpBucket() ResourceRemoteInfoGcpBucket {
	if o == nil || IsNil(o.GcpBucket) {
		var ret ResourceRemoteInfoGcpBucket
		return ret
	}
	return *o.GcpBucket
}

// GetGcpBucketOk returns a tuple with the GcpBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetGcpBucketOk() (*ResourceRemoteInfoGcpBucket, bool) {
	if o == nil || IsNil(o.GcpBucket) {
		return nil, false
	}
	return o.GcpBucket, true
}

// HasGcpBucket returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasGcpBucket() bool {
	if o != nil && !IsNil(o.GcpBucket) {
		return true
	}

	return false
}

// SetGcpBucket gets a reference to the given ResourceRemoteInfoGcpBucket and assigns it to the GcpBucket field.
func (o *ResourceRemoteInfo) SetGcpBucket(v ResourceRemoteInfoGcpBucket) {
	o.GcpBucket = &v
}

// GetGcpComputeInstance returns the GcpComputeInstance field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetGcpComputeInstance() ResourceRemoteInfoGcpComputeInstance {
	if o == nil || IsNil(o.GcpComputeInstance) {
		var ret ResourceRemoteInfoGcpComputeInstance
		return ret
	}
	return *o.GcpComputeInstance
}

// GetGcpComputeInstanceOk returns a tuple with the GcpComputeInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetGcpComputeInstanceOk() (*ResourceRemoteInfoGcpComputeInstance, bool) {
	if o == nil || IsNil(o.GcpComputeInstance) {
		return nil, false
	}
	return o.GcpComputeInstance, true
}

// HasGcpComputeInstance returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasGcpComputeInstance() bool {
	if o != nil && !IsNil(o.GcpComputeInstance) {
		return true
	}

	return false
}

// SetGcpComputeInstance gets a reference to the given ResourceRemoteInfoGcpComputeInstance and assigns it to the GcpComputeInstance field.
func (o *ResourceRemoteInfo) SetGcpComputeInstance(v ResourceRemoteInfoGcpComputeInstance) {
	o.GcpComputeInstance = &v
}

// GetGcpFolder returns the GcpFolder field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetGcpFolder() ResourceRemoteInfoGcpFolder {
	if o == nil || IsNil(o.GcpFolder) {
		var ret ResourceRemoteInfoGcpFolder
		return ret
	}
	return *o.GcpFolder
}

// GetGcpFolderOk returns a tuple with the GcpFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetGcpFolderOk() (*ResourceRemoteInfoGcpFolder, bool) {
	if o == nil || IsNil(o.GcpFolder) {
		return nil, false
	}
	return o.GcpFolder, true
}

// HasGcpFolder returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasGcpFolder() bool {
	if o != nil && !IsNil(o.GcpFolder) {
		return true
	}

	return false
}

// SetGcpFolder gets a reference to the given ResourceRemoteInfoGcpFolder and assigns it to the GcpFolder field.
func (o *ResourceRemoteInfo) SetGcpFolder(v ResourceRemoteInfoGcpFolder) {
	o.GcpFolder = &v
}

// GetGcpGkeCluster returns the GcpGkeCluster field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetGcpGkeCluster() ResourceRemoteInfoGcpGkeCluster {
	if o == nil || IsNil(o.GcpGkeCluster) {
		var ret ResourceRemoteInfoGcpGkeCluster
		return ret
	}
	return *o.GcpGkeCluster
}

// GetGcpGkeClusterOk returns a tuple with the GcpGkeCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetGcpGkeClusterOk() (*ResourceRemoteInfoGcpGkeCluster, bool) {
	if o == nil || IsNil(o.GcpGkeCluster) {
		return nil, false
	}
	return o.GcpGkeCluster, true
}

// HasGcpGkeCluster returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasGcpGkeCluster() bool {
	if o != nil && !IsNil(o.GcpGkeCluster) {
		return true
	}

	return false
}

// SetGcpGkeCluster gets a reference to the given ResourceRemoteInfoGcpGkeCluster and assigns it to the GcpGkeCluster field.
func (o *ResourceRemoteInfo) SetGcpGkeCluster(v ResourceRemoteInfoGcpGkeCluster) {
	o.GcpGkeCluster = &v
}

// GetGcpProject returns the GcpProject field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetGcpProject() ResourceRemoteInfoGcpProject {
	if o == nil || IsNil(o.GcpProject) {
		var ret ResourceRemoteInfoGcpProject
		return ret
	}
	return *o.GcpProject
}

// GetGcpProjectOk returns a tuple with the GcpProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetGcpProjectOk() (*ResourceRemoteInfoGcpProject, bool) {
	if o == nil || IsNil(o.GcpProject) {
		return nil, false
	}
	return o.GcpProject, true
}

// HasGcpProject returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasGcpProject() bool {
	if o != nil && !IsNil(o.GcpProject) {
		return true
	}

	return false
}

// SetGcpProject gets a reference to the given ResourceRemoteInfoGcpProject and assigns it to the GcpProject field.
func (o *ResourceRemoteInfo) SetGcpProject(v ResourceRemoteInfoGcpProject) {
	o.GcpProject = &v
}

// GetGcpSqlInstance returns the GcpSqlInstance field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetGcpSqlInstance() ResourceRemoteInfoGcpSqlInstance {
	if o == nil || IsNil(o.GcpSqlInstance) {
		var ret ResourceRemoteInfoGcpSqlInstance
		return ret
	}
	return *o.GcpSqlInstance
}

// GetGcpSqlInstanceOk returns a tuple with the GcpSqlInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetGcpSqlInstanceOk() (*ResourceRemoteInfoGcpSqlInstance, bool) {
	if o == nil || IsNil(o.GcpSqlInstance) {
		return nil, false
	}
	return o.GcpSqlInstance, true
}

// HasGcpSqlInstance returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasGcpSqlInstance() bool {
	if o != nil && !IsNil(o.GcpSqlInstance) {
		return true
	}

	return false
}

// SetGcpSqlInstance gets a reference to the given ResourceRemoteInfoGcpSqlInstance and assigns it to the GcpSqlInstance field.
func (o *ResourceRemoteInfo) SetGcpSqlInstance(v ResourceRemoteInfoGcpSqlInstance) {
	o.GcpSqlInstance = &v
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

// GetPagerdutyRole returns the PagerdutyRole field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetPagerdutyRole() ResourceRemoteInfoPagerdutyRole {
	if o == nil || IsNil(o.PagerdutyRole) {
		var ret ResourceRemoteInfoPagerdutyRole
		return ret
	}
	return *o.PagerdutyRole
}

// GetPagerdutyRoleOk returns a tuple with the PagerdutyRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetPagerdutyRoleOk() (*ResourceRemoteInfoPagerdutyRole, bool) {
	if o == nil || IsNil(o.PagerdutyRole) {
		return nil, false
	}
	return o.PagerdutyRole, true
}

// HasPagerdutyRole returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasPagerdutyRole() bool {
	if o != nil && !IsNil(o.PagerdutyRole) {
		return true
	}

	return false
}

// SetPagerdutyRole gets a reference to the given ResourceRemoteInfoPagerdutyRole and assigns it to the PagerdutyRole field.
func (o *ResourceRemoteInfo) SetPagerdutyRole(v ResourceRemoteInfoPagerdutyRole) {
	o.PagerdutyRole = &v
}

// GetSalesforcePermissionSet returns the SalesforcePermissionSet field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetSalesforcePermissionSet() ResourceRemoteInfoSalesforcePermissionSet {
	if o == nil || IsNil(o.SalesforcePermissionSet) {
		var ret ResourceRemoteInfoSalesforcePermissionSet
		return ret
	}
	return *o.SalesforcePermissionSet
}

// GetSalesforcePermissionSetOk returns a tuple with the SalesforcePermissionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetSalesforcePermissionSetOk() (*ResourceRemoteInfoSalesforcePermissionSet, bool) {
	if o == nil || IsNil(o.SalesforcePermissionSet) {
		return nil, false
	}
	return o.SalesforcePermissionSet, true
}

// HasSalesforcePermissionSet returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasSalesforcePermissionSet() bool {
	if o != nil && !IsNil(o.SalesforcePermissionSet) {
		return true
	}

	return false
}

// SetSalesforcePermissionSet gets a reference to the given ResourceRemoteInfoSalesforcePermissionSet and assigns it to the SalesforcePermissionSet field.
func (o *ResourceRemoteInfo) SetSalesforcePermissionSet(v ResourceRemoteInfoSalesforcePermissionSet) {
	o.SalesforcePermissionSet = &v
}

// GetSalesforceProfile returns the SalesforceProfile field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetSalesforceProfile() ResourceRemoteInfoSalesforceProfile {
	if o == nil || IsNil(o.SalesforceProfile) {
		var ret ResourceRemoteInfoSalesforceProfile
		return ret
	}
	return *o.SalesforceProfile
}

// GetSalesforceProfileOk returns a tuple with the SalesforceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetSalesforceProfileOk() (*ResourceRemoteInfoSalesforceProfile, bool) {
	if o == nil || IsNil(o.SalesforceProfile) {
		return nil, false
	}
	return o.SalesforceProfile, true
}

// HasSalesforceProfile returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasSalesforceProfile() bool {
	if o != nil && !IsNil(o.SalesforceProfile) {
		return true
	}

	return false
}

// SetSalesforceProfile gets a reference to the given ResourceRemoteInfoSalesforceProfile and assigns it to the SalesforceProfile field.
func (o *ResourceRemoteInfo) SetSalesforceProfile(v ResourceRemoteInfoSalesforceProfile) {
	o.SalesforceProfile = &v
}

// GetSalesforceRole returns the SalesforceRole field value if set, zero value otherwise.
func (o *ResourceRemoteInfo) GetSalesforceRole() ResourceRemoteInfoSalesforceRole {
	if o == nil || IsNil(o.SalesforceRole) {
		var ret ResourceRemoteInfoSalesforceRole
		return ret
	}
	return *o.SalesforceRole
}

// GetSalesforceRoleOk returns a tuple with the SalesforceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceRemoteInfo) GetSalesforceRoleOk() (*ResourceRemoteInfoSalesforceRole, bool) {
	if o == nil || IsNil(o.SalesforceRole) {
		return nil, false
	}
	return o.SalesforceRole, true
}

// HasSalesforceRole returns a boolean if a field has been set.
func (o *ResourceRemoteInfo) HasSalesforceRole() bool {
	if o != nil && !IsNil(o.SalesforceRole) {
		return true
	}

	return false
}

// SetSalesforceRole gets a reference to the given ResourceRemoteInfoSalesforceRole and assigns it to the SalesforceRole field.
func (o *ResourceRemoteInfo) SetSalesforceRole(v ResourceRemoteInfoSalesforceRole) {
	o.SalesforceRole = &v
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
	if !IsNil(o.AwsAccount) {
		toSerialize["aws_account"] = o.AwsAccount
	}
	if !IsNil(o.AwsPermissionSet) {
		toSerialize["aws_permission_set"] = o.AwsPermissionSet
	}
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
	if !IsNil(o.GcpBucket) {
		toSerialize["gcp_bucket"] = o.GcpBucket
	}
	if !IsNil(o.GcpComputeInstance) {
		toSerialize["gcp_compute_instance"] = o.GcpComputeInstance
	}
	if !IsNil(o.GcpFolder) {
		toSerialize["gcp_folder"] = o.GcpFolder
	}
	if !IsNil(o.GcpGkeCluster) {
		toSerialize["gcp_gke_cluster"] = o.GcpGkeCluster
	}
	if !IsNil(o.GcpProject) {
		toSerialize["gcp_project"] = o.GcpProject
	}
	if !IsNil(o.GcpSqlInstance) {
		toSerialize["gcp_sql_instance"] = o.GcpSqlInstance
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
	if !IsNil(o.PagerdutyRole) {
		toSerialize["pagerduty_role"] = o.PagerdutyRole
	}
	if !IsNil(o.SalesforcePermissionSet) {
		toSerialize["salesforce_permission_set"] = o.SalesforcePermissionSet
	}
	if !IsNil(o.SalesforceProfile) {
		toSerialize["salesforce_profile"] = o.SalesforceProfile
	}
	if !IsNil(o.SalesforceRole) {
		toSerialize["salesforce_role"] = o.SalesforceRole
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


