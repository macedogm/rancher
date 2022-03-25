/*
Copyright 2022 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=management.cattle.io
package v3

import (
	management "github.com/rancher/rancher/pkg/apis/management.cattle.io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	ActiveDirectoryProviderResourceName                 = "activedirectoryproviders"
	AuthConfigResourceName                              = "authconfigs"
	AuthProviderResourceName                            = "authproviders"
	AuthTokenResourceName                               = "authtokens"
	AzureADProviderResourceName                         = "azureadproviders"
	CatalogResourceName                                 = "catalogs"
	CatalogTemplateResourceName                         = "catalogtemplates"
	CatalogTemplateVersionResourceName                  = "catalogtemplateversions"
	CisBenchmarkVersionResourceName                     = "cisbenchmarkversions"
	CisConfigResourceName                               = "cisconfigs"
	CloudCredentialResourceName                         = "cloudcredentials"
	ClusterResourceName                                 = "clusters"
	ClusterAlertResourceName                            = "clusteralerts"
	ClusterAlertGroupResourceName                       = "clusteralertgroups"
	ClusterAlertRuleResourceName                        = "clusteralertrules"
	ClusterCatalogResourceName                          = "clustercatalogs"
	ClusterLoggingResourceName                          = "clusterloggings"
	ClusterMonitorGraphResourceName                     = "clustermonitorgraphs"
	ClusterRegistrationTokenResourceName                = "clusterregistrationtokens"
	ClusterRoleTemplateBindingResourceName              = "clusterroletemplatebindings"
	ClusterScanResourceName                             = "clusterscans"
	ClusterTemplateResourceName                         = "clustertemplates"
	ClusterTemplateRevisionResourceName                 = "clustertemplaterevisions"
	ComposeConfigResourceName                           = "composeconfigs"
	DynamicSchemaResourceName                           = "dynamicschemas"
	EtcdBackupResourceName                              = "etcdbackups"
	FeatureResourceName                                 = "features"
	FleetWorkspaceResourceName                          = "fleetworkspaces"
	FreeIpaProviderResourceName                         = "freeipaproviders"
	GithubProviderResourceName                          = "githubproviders"
	GlobalDnsResourceName                               = "globaldnses"
	GlobalDnsProviderResourceName                       = "globaldnsproviders"
	GlobalRoleResourceName                              = "globalroles"
	GlobalRoleBindingResourceName                       = "globalrolebindings"
	GoogleOAuthProviderResourceName                     = "googleoauthproviders"
	GroupResourceName                                   = "groups"
	GroupMemberResourceName                             = "groupmembers"
	KontainerDriverResourceName                         = "kontainerdrivers"
	LocalProviderResourceName                           = "localproviders"
	MonitorMetricResourceName                           = "monitormetrics"
	MultiClusterAppResourceName                         = "multiclusterapps"
	MultiClusterAppRevisionResourceName                 = "multiclusterapprevisions"
	NodeResourceName                                    = "nodes"
	NodeDriverResourceName                              = "nodedrivers"
	NodePoolResourceName                                = "nodepools"
	NodeTemplateResourceName                            = "nodetemplates"
	NotifierResourceName                                = "notifiers"
	OpenLdapProviderResourceName                        = "openldapproviders"
	PodSecurityPolicyTemplateResourceName               = "podsecuritypolicytemplates"
	PodSecurityPolicyTemplateProjectBindingResourceName = "podsecuritypolicytemplateprojectbindings"
	PreferenceResourceName                              = "preferences"
	PrincipalResourceName                               = "principals"
	ProjectResourceName                                 = "projects"
	ProjectAlertResourceName                            = "projectalerts"
	ProjectAlertGroupResourceName                       = "projectalertgroups"
	ProjectAlertRuleResourceName                        = "projectalertrules"
	ProjectCatalogResourceName                          = "projectcatalogs"
	ProjectLoggingResourceName                          = "projectloggings"
	ProjectMonitorGraphResourceName                     = "projectmonitorgraphs"
	ProjectNetworkPolicyResourceName                    = "projectnetworkpolicies"
	ProjectRoleTemplateBindingResourceName              = "projectroletemplatebindings"
	RkeAddonResourceName                                = "rkeaddons"
	RkeK8sServiceOptionResourceName                     = "rkek8sserviceoptions"
	RkeK8sSystemImageResourceName                       = "rkek8ssystemimages"
	RoleTemplateResourceName                            = "roletemplates"
	SamlProviderResourceName                            = "samlproviders"
	SamlTokenResourceName                               = "samltokens"
	SettingResourceName                                 = "settings"
	TemplateResourceName                                = "templates"
	TemplateContentResourceName                         = "templatecontents"
	TemplateVersionResourceName                         = "templateversions"
	TokenResourceName                                   = "tokens"
	UserResourceName                                    = "users"
	UserAttributeResourceName                           = "userattributes"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: management.GroupName, Version: "v3"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&ActiveDirectoryProvider{},
		&ActiveDirectoryProviderList{},
		&AuthConfig{},
		&AuthConfigList{},
		&AuthProvider{},
		&AuthProviderList{},
		&AuthToken{},
		&AuthTokenList{},
		&AzureADProvider{},
		&AzureADProviderList{},
		&Catalog{},
		&CatalogList{},
		&CatalogTemplate{},
		&CatalogTemplateList{},
		&CatalogTemplateVersion{},
		&CatalogTemplateVersionList{},
		&CisBenchmarkVersion{},
		&CisBenchmarkVersionList{},
		&CisConfig{},
		&CisConfigList{},
		&CloudCredential{},
		&CloudCredentialList{},
		&Cluster{},
		&ClusterList{},
		&ClusterAlert{},
		&ClusterAlertList{},
		&ClusterAlertGroup{},
		&ClusterAlertGroupList{},
		&ClusterAlertRule{},
		&ClusterAlertRuleList{},
		&ClusterCatalog{},
		&ClusterCatalogList{},
		&ClusterLogging{},
		&ClusterLoggingList{},
		&ClusterMonitorGraph{},
		&ClusterMonitorGraphList{},
		&ClusterRegistrationToken{},
		&ClusterRegistrationTokenList{},
		&ClusterRoleTemplateBinding{},
		&ClusterRoleTemplateBindingList{},
		&ClusterScan{},
		&ClusterScanList{},
		&ClusterTemplate{},
		&ClusterTemplateList{},
		&ClusterTemplateRevision{},
		&ClusterTemplateRevisionList{},
		&ComposeConfig{},
		&ComposeConfigList{},
		&DynamicSchema{},
		&DynamicSchemaList{},
		&EtcdBackup{},
		&EtcdBackupList{},
		&Feature{},
		&FeatureList{},
		&FleetWorkspace{},
		&FleetWorkspaceList{},
		&FreeIpaProvider{},
		&FreeIpaProviderList{},
		&GithubProvider{},
		&GithubProviderList{},
		&GlobalDns{},
		&GlobalDnsList{},
		&GlobalDnsProvider{},
		&GlobalDnsProviderList{},
		&GlobalRole{},
		&GlobalRoleList{},
		&GlobalRoleBinding{},
		&GlobalRoleBindingList{},
		&GoogleOAuthProvider{},
		&GoogleOAuthProviderList{},
		&Group{},
		&GroupList{},
		&GroupMember{},
		&GroupMemberList{},
		&KontainerDriver{},
		&KontainerDriverList{},
		&LocalProvider{},
		&LocalProviderList{},
		&MonitorMetric{},
		&MonitorMetricList{},
		&MultiClusterApp{},
		&MultiClusterAppList{},
		&MultiClusterAppRevision{},
		&MultiClusterAppRevisionList{},
		&Node{},
		&NodeList{},
		&NodeDriver{},
		&NodeDriverList{},
		&NodePool{},
		&NodePoolList{},
		&NodeTemplate{},
		&NodeTemplateList{},
		&Notifier{},
		&NotifierList{},
		&OpenLdapProvider{},
		&OpenLdapProviderList{},
		&PodSecurityPolicyTemplate{},
		&PodSecurityPolicyTemplateList{},
		&PodSecurityPolicyTemplateProjectBinding{},
		&PodSecurityPolicyTemplateProjectBindingList{},
		&Preference{},
		&PreferenceList{},
		&Principal{},
		&PrincipalList{},
		&Project{},
		&ProjectList{},
		&ProjectAlert{},
		&ProjectAlertList{},
		&ProjectAlertGroup{},
		&ProjectAlertGroupList{},
		&ProjectAlertRule{},
		&ProjectAlertRuleList{},
		&ProjectCatalog{},
		&ProjectCatalogList{},
		&ProjectLogging{},
		&ProjectLoggingList{},
		&ProjectMonitorGraph{},
		&ProjectMonitorGraphList{},
		&ProjectNetworkPolicy{},
		&ProjectNetworkPolicyList{},
		&ProjectRoleTemplateBinding{},
		&ProjectRoleTemplateBindingList{},
		&RkeAddon{},
		&RkeAddonList{},
		&RkeK8sServiceOption{},
		&RkeK8sServiceOptionList{},
		&RkeK8sSystemImage{},
		&RkeK8sSystemImageList{},
		&RoleTemplate{},
		&RoleTemplateList{},
		&SamlProvider{},
		&SamlProviderList{},
		&SamlToken{},
		&SamlTokenList{},
		&Setting{},
		&SettingList{},
		&Template{},
		&TemplateList{},
		&TemplateContent{},
		&TemplateContentList{},
		&TemplateVersion{},
		&TemplateVersionList{},
		&Token{},
		&TokenList{},
		&User{},
		&UserList{},
		&UserAttribute{},
		&UserAttributeList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}