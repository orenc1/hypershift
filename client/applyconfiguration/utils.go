/*


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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/openshift/hypershift/api/certificates/v1alpha1"
	hypershiftv1alpha1 "github.com/openshift/hypershift/api/hypershift/v1alpha1"
	v1beta1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
	schedulingv1alpha1 "github.com/openshift/hypershift/api/scheduling/v1alpha1"
	certificatesv1alpha1 "github.com/openshift/hypershift/client/applyconfiguration/certificates/v1alpha1"
	applyconfigurationhypershiftv1alpha1 "github.com/openshift/hypershift/client/applyconfiguration/hypershift/v1alpha1"
	hypershiftv1beta1 "github.com/openshift/hypershift/client/applyconfiguration/hypershift/v1beta1"
	applyconfigurationschedulingv1alpha1 "github.com/openshift/hypershift/client/applyconfiguration/scheduling/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=certificates.hypershift.openshift.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("CertificateRevocationRequest"):
		return &certificatesv1alpha1.CertificateRevocationRequestApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CertificateRevocationRequestSpec"):
		return &certificatesv1alpha1.CertificateRevocationRequestSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CertificateRevocationRequestStatus"):
		return &certificatesv1alpha1.CertificateRevocationRequestStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CertificateSigningRequestApproval"):
		return &certificatesv1alpha1.CertificateSigningRequestApprovalApplyConfiguration{}

		// Group=hypershift.openshift.io, Version=v1alpha1
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AESCBCSpec"):
		return &applyconfigurationhypershiftv1alpha1.AESCBCSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AgentNodePoolPlatform"):
		return &applyconfigurationhypershiftv1alpha1.AgentNodePoolPlatformApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AgentPlatformSpec"):
		return &applyconfigurationhypershiftv1alpha1.AgentPlatformSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("APIEndpoint"):
		return &applyconfigurationhypershiftv1alpha1.APIEndpointApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("APIServerNetworking"):
		return &applyconfigurationhypershiftv1alpha1.APIServerNetworkingApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSCloudProviderConfig"):
		return &applyconfigurationhypershiftv1alpha1.AWSCloudProviderConfigApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSKMSAuthSpec"):
		return &applyconfigurationhypershiftv1alpha1.AWSKMSAuthSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSKMSKeyEntry"):
		return &applyconfigurationhypershiftv1alpha1.AWSKMSKeyEntryApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSKMSSpec"):
		return &applyconfigurationhypershiftv1alpha1.AWSKMSSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSNodePoolPlatform"):
		return &applyconfigurationhypershiftv1alpha1.AWSNodePoolPlatformApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSPlatformSpec"):
		return &applyconfigurationhypershiftv1alpha1.AWSPlatformSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSPlatformStatus"):
		return &applyconfigurationhypershiftv1alpha1.AWSPlatformStatusApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSResourceReference"):
		return &applyconfigurationhypershiftv1alpha1.AWSResourceReferenceApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSResourceTag"):
		return &applyconfigurationhypershiftv1alpha1.AWSResourceTagApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSRoleCredentials"):
		return &applyconfigurationhypershiftv1alpha1.AWSRoleCredentialsApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSRolesRef"):
		return &applyconfigurationhypershiftv1alpha1.AWSRolesRefApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AWSServiceEndpoint"):
		return &applyconfigurationhypershiftv1alpha1.AWSServiceEndpointApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AzureKMSSpec"):
		return &applyconfigurationhypershiftv1alpha1.AzureKMSSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AzureNodePoolPlatform"):
		return &applyconfigurationhypershiftv1alpha1.AzureNodePoolPlatformApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("AzurePlatformSpec"):
		return &applyconfigurationhypershiftv1alpha1.AzurePlatformSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ClusterAutoscaling"):
		return &applyconfigurationhypershiftv1alpha1.ClusterAutoscalingApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ClusterConfiguration"):
		return &applyconfigurationhypershiftv1alpha1.ClusterConfigurationApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ClusterNetworkEntry"):
		return &applyconfigurationhypershiftv1alpha1.ClusterNetworkEntryApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ClusterNetworking"):
		return &applyconfigurationhypershiftv1alpha1.ClusterNetworkingApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ClusterVersionStatus"):
		return &applyconfigurationhypershiftv1alpha1.ClusterVersionStatusApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("DNSSpec"):
		return &applyconfigurationhypershiftv1alpha1.DNSSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("EtcdSpec"):
		return &applyconfigurationhypershiftv1alpha1.EtcdSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("EtcdTLSConfig"):
		return &applyconfigurationhypershiftv1alpha1.EtcdTLSConfigApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("Filter"):
		return &applyconfigurationhypershiftv1alpha1.FilterApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("HostedCluster"):
		return &applyconfigurationhypershiftv1alpha1.HostedClusterApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("HostedClusterSpec"):
		return &applyconfigurationhypershiftv1alpha1.HostedClusterSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("HostedClusterStatus"):
		return &applyconfigurationhypershiftv1alpha1.HostedClusterStatusApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("IBMCloudKMSAuthSpec"):
		return &applyconfigurationhypershiftv1alpha1.IBMCloudKMSAuthSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("IBMCloudKMSKeyEntry"):
		return &applyconfigurationhypershiftv1alpha1.IBMCloudKMSKeyEntryApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("IBMCloudKMSSpec"):
		return &applyconfigurationhypershiftv1alpha1.IBMCloudKMSSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("IBMCloudKMSUnmanagedAuthSpec"):
		return &applyconfigurationhypershiftv1alpha1.IBMCloudKMSUnmanagedAuthSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("IBMCloudPlatformSpec"):
		return &applyconfigurationhypershiftv1alpha1.IBMCloudPlatformSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ImageContentSource"):
		return &applyconfigurationhypershiftv1alpha1.ImageContentSourceApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("InPlaceUpgrade"):
		return &applyconfigurationhypershiftv1alpha1.InPlaceUpgradeApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KMSSpec"):
		return &applyconfigurationhypershiftv1alpha1.KMSSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubeconfigSecretRef"):
		return &applyconfigurationhypershiftv1alpha1.KubeconfigSecretRefApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtCachingStrategy"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtCachingStrategyApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtCompute"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtComputeApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtDiskImage"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtDiskImageApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtManualStorageDriverConfig"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtManualStorageDriverConfigApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtNetwork"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtNetworkApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtNodePoolPlatform"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtNodePoolPlatformApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubeVirtNodePoolStatus"):
		return &applyconfigurationhypershiftv1alpha1.KubeVirtNodePoolStatusApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtPersistentVolume"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtPersistentVolumeApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtPlatformCredentials"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtPlatformCredentialsApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtPlatformSpec"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtPlatformSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtRootVolume"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtRootVolumeApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtStorageClassMapping"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtStorageClassMappingApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtStorageDriverSpec"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtStorageDriverSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("KubevirtVolume"):
		return &applyconfigurationhypershiftv1alpha1.KubevirtVolumeApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("LoadBalancerPublishingStrategy"):
		return &applyconfigurationhypershiftv1alpha1.LoadBalancerPublishingStrategyApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("MachineNetworkEntry"):
		return &applyconfigurationhypershiftv1alpha1.MachineNetworkEntryApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ManagedEtcdSpec"):
		return &applyconfigurationhypershiftv1alpha1.ManagedEtcdSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ManagedEtcdStorageSpec"):
		return &applyconfigurationhypershiftv1alpha1.ManagedEtcdStorageSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("NodePool"):
		return &applyconfigurationhypershiftv1alpha1.NodePoolApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("NodePoolAutoScaling"):
		return &applyconfigurationhypershiftv1alpha1.NodePoolAutoScalingApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("NodePoolCondition"):
		return &applyconfigurationhypershiftv1alpha1.NodePoolConditionApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("NodePoolManagement"):
		return &applyconfigurationhypershiftv1alpha1.NodePoolManagementApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("NodePoolPlatform"):
		return &applyconfigurationhypershiftv1alpha1.NodePoolPlatformApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("NodePoolPlatformStatus"):
		return &applyconfigurationhypershiftv1alpha1.NodePoolPlatformStatusApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("NodePoolSpec"):
		return &applyconfigurationhypershiftv1alpha1.NodePoolSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("NodePoolStatus"):
		return &applyconfigurationhypershiftv1alpha1.NodePoolStatusApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("NodePortPublishingStrategy"):
		return &applyconfigurationhypershiftv1alpha1.NodePortPublishingStrategyApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("PersistentVolumeEtcdStorageSpec"):
		return &applyconfigurationhypershiftv1alpha1.PersistentVolumeEtcdStorageSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("PlatformSpec"):
		return &applyconfigurationhypershiftv1alpha1.PlatformSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("PlatformStatus"):
		return &applyconfigurationhypershiftv1alpha1.PlatformStatusApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("PowerVSNodePoolPlatform"):
		return &applyconfigurationhypershiftv1alpha1.PowerVSNodePoolPlatformApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("PowerVSPlatformSpec"):
		return &applyconfigurationhypershiftv1alpha1.PowerVSPlatformSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("PowerVSResourceReference"):
		return &applyconfigurationhypershiftv1alpha1.PowerVSResourceReferenceApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("PowerVSVPC"):
		return &applyconfigurationhypershiftv1alpha1.PowerVSVPCApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("Release"):
		return &applyconfigurationhypershiftv1alpha1.ReleaseApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ReplaceUpgrade"):
		return &applyconfigurationhypershiftv1alpha1.ReplaceUpgradeApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("RollingUpdate"):
		return &applyconfigurationhypershiftv1alpha1.RollingUpdateApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("RoutePublishingStrategy"):
		return &applyconfigurationhypershiftv1alpha1.RoutePublishingStrategyApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("SecretEncryptionSpec"):
		return &applyconfigurationhypershiftv1alpha1.SecretEncryptionSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ServiceNetworkEntry"):
		return &applyconfigurationhypershiftv1alpha1.ServiceNetworkEntryApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ServicePublishingStrategy"):
		return &applyconfigurationhypershiftv1alpha1.ServicePublishingStrategyApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("ServicePublishingStrategyMapping"):
		return &applyconfigurationhypershiftv1alpha1.ServicePublishingStrategyMappingApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("Taint"):
		return &applyconfigurationhypershiftv1alpha1.TaintApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("UnmanagedEtcdSpec"):
		return &applyconfigurationhypershiftv1alpha1.UnmanagedEtcdSpecApplyConfiguration{}
	case hypershiftv1alpha1.SchemeGroupVersion.WithKind("Volume"):
		return &applyconfigurationhypershiftv1alpha1.VolumeApplyConfiguration{}

		// Group=hypershift.openshift.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("AESCBCSpec"):
		return &hypershiftv1beta1.AESCBCSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AgentNodePoolPlatform"):
		return &hypershiftv1beta1.AgentNodePoolPlatformApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AgentPlatformSpec"):
		return &hypershiftv1beta1.AgentPlatformSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("APIEndpoint"):
		return &hypershiftv1beta1.APIEndpointApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("APIServerNetworking"):
		return &hypershiftv1beta1.APIServerNetworkingApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSCloudProviderConfig"):
		return &hypershiftv1beta1.AWSCloudProviderConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSKMSAuthSpec"):
		return &hypershiftv1beta1.AWSKMSAuthSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSKMSKeyEntry"):
		return &hypershiftv1beta1.AWSKMSKeyEntryApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSKMSSpec"):
		return &hypershiftv1beta1.AWSKMSSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSNodePoolPlatform"):
		return &hypershiftv1beta1.AWSNodePoolPlatformApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSPlatformSpec"):
		return &hypershiftv1beta1.AWSPlatformSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSPlatformStatus"):
		return &hypershiftv1beta1.AWSPlatformStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSResourceReference"):
		return &hypershiftv1beta1.AWSResourceReferenceApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSResourceTag"):
		return &hypershiftv1beta1.AWSResourceTagApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSRolesRef"):
		return &hypershiftv1beta1.AWSRolesRefApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AWSServiceEndpoint"):
		return &hypershiftv1beta1.AWSServiceEndpointApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AzureKMSSpec"):
		return &hypershiftv1beta1.AzureKMSSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AzureNodePoolPlatform"):
		return &hypershiftv1beta1.AzureNodePoolPlatformApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AzurePlatformSpec"):
		return &hypershiftv1beta1.AzurePlatformSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CertificateSigningRequestApproval"):
		return &hypershiftv1beta1.CertificateSigningRequestApprovalApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterAutoscaling"):
		return &hypershiftv1beta1.ClusterAutoscalingApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterConfiguration"):
		return &hypershiftv1beta1.ClusterConfigurationApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterNetworkEntry"):
		return &hypershiftv1beta1.ClusterNetworkEntryApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterNetworking"):
		return &hypershiftv1beta1.ClusterNetworkingApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterVersionStatus"):
		return &hypershiftv1beta1.ClusterVersionStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("DNSSpec"):
		return &hypershiftv1beta1.DNSSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("EtcdSpec"):
		return &hypershiftv1beta1.EtcdSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("EtcdTLSConfig"):
		return &hypershiftv1beta1.EtcdTLSConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Filter"):
		return &hypershiftv1beta1.FilterApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("HostedCluster"):
		return &hypershiftv1beta1.HostedClusterApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("HostedClusterSpec"):
		return &hypershiftv1beta1.HostedClusterSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("HostedClusterStatus"):
		return &hypershiftv1beta1.HostedClusterStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("HostedControlPlane"):
		return &hypershiftv1beta1.HostedControlPlaneApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("HostedControlPlaneSpec"):
		return &hypershiftv1beta1.HostedControlPlaneSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("HostedControlPlaneStatus"):
		return &hypershiftv1beta1.HostedControlPlaneStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("IBMCloudKMSAuthSpec"):
		return &hypershiftv1beta1.IBMCloudKMSAuthSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("IBMCloudKMSKeyEntry"):
		return &hypershiftv1beta1.IBMCloudKMSKeyEntryApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("IBMCloudKMSSpec"):
		return &hypershiftv1beta1.IBMCloudKMSSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("IBMCloudKMSUnmanagedAuthSpec"):
		return &hypershiftv1beta1.IBMCloudKMSUnmanagedAuthSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("IBMCloudPlatformSpec"):
		return &hypershiftv1beta1.IBMCloudPlatformSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ImageContentSource"):
		return &hypershiftv1beta1.ImageContentSourceApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("InPlaceUpgrade"):
		return &hypershiftv1beta1.InPlaceUpgradeApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KMSSpec"):
		return &hypershiftv1beta1.KMSSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubeconfigSecretRef"):
		return &hypershiftv1beta1.KubeconfigSecretRefApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtCachingStrategy"):
		return &hypershiftv1beta1.KubevirtCachingStrategyApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtCompute"):
		return &hypershiftv1beta1.KubevirtComputeApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtDiskImage"):
		return &hypershiftv1beta1.KubevirtDiskImageApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtManualStorageDriverConfig"):
		return &hypershiftv1beta1.KubevirtManualStorageDriverConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtNetwork"):
		return &hypershiftv1beta1.KubevirtNetworkApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtNodePoolPlatform"):
		return &hypershiftv1beta1.KubevirtNodePoolPlatformApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubeVirtNodePoolStatus"):
		return &hypershiftv1beta1.KubeVirtNodePoolStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtPersistentVolume"):
		return &hypershiftv1beta1.KubevirtPersistentVolumeApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtPlatformCredentials"):
		return &hypershiftv1beta1.KubevirtPlatformCredentialsApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtPlatformSpec"):
		return &hypershiftv1beta1.KubevirtPlatformSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtRootVolume"):
		return &hypershiftv1beta1.KubevirtRootVolumeApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtStorageClassMapping"):
		return &hypershiftv1beta1.KubevirtStorageClassMappingApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtStorageDriverSpec"):
		return &hypershiftv1beta1.KubevirtStorageDriverSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubevirtVolume"):
		return &hypershiftv1beta1.KubevirtVolumeApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LoadBalancerPublishingStrategy"):
		return &hypershiftv1beta1.LoadBalancerPublishingStrategyApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MachineNetworkEntry"):
		return &hypershiftv1beta1.MachineNetworkEntryApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ManagedEtcdSpec"):
		return &hypershiftv1beta1.ManagedEtcdSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ManagedEtcdStorageSpec"):
		return &hypershiftv1beta1.ManagedEtcdStorageSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodePool"):
		return &hypershiftv1beta1.NodePoolApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodePoolAutoScaling"):
		return &hypershiftv1beta1.NodePoolAutoScalingApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodePoolCondition"):
		return &hypershiftv1beta1.NodePoolConditionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodePoolManagement"):
		return &hypershiftv1beta1.NodePoolManagementApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodePoolPlatform"):
		return &hypershiftv1beta1.NodePoolPlatformApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodePoolPlatformStatus"):
		return &hypershiftv1beta1.NodePoolPlatformStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodePoolSpec"):
		return &hypershiftv1beta1.NodePoolSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodePoolStatus"):
		return &hypershiftv1beta1.NodePoolStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodePortPublishingStrategy"):
		return &hypershiftv1beta1.NodePortPublishingStrategyApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PersistentVolumeEtcdStorageSpec"):
		return &hypershiftv1beta1.PersistentVolumeEtcdStorageSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PlatformSpec"):
		return &hypershiftv1beta1.PlatformSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PlatformStatus"):
		return &hypershiftv1beta1.PlatformStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PowerVSNodePoolPlatform"):
		return &hypershiftv1beta1.PowerVSNodePoolPlatformApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PowerVSPlatformSpec"):
		return &hypershiftv1beta1.PowerVSPlatformSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PowerVSResourceReference"):
		return &hypershiftv1beta1.PowerVSResourceReferenceApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PowerVSVPC"):
		return &hypershiftv1beta1.PowerVSVPCApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Release"):
		return &hypershiftv1beta1.ReleaseApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ReplaceUpgrade"):
		return &hypershiftv1beta1.ReplaceUpgradeApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("RollingUpdate"):
		return &hypershiftv1beta1.RollingUpdateApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("RoutePublishingStrategy"):
		return &hypershiftv1beta1.RoutePublishingStrategyApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("SecretEncryptionSpec"):
		return &hypershiftv1beta1.SecretEncryptionSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ServiceNetworkEntry"):
		return &hypershiftv1beta1.ServiceNetworkEntryApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ServicePublishingStrategy"):
		return &hypershiftv1beta1.ServicePublishingStrategyApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ServicePublishingStrategyMapping"):
		return &hypershiftv1beta1.ServicePublishingStrategyMappingApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Taint"):
		return &hypershiftv1beta1.TaintApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("UnmanagedEtcdSpec"):
		return &hypershiftv1beta1.UnmanagedEtcdSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Volume"):
		return &hypershiftv1beta1.VolumeApplyConfiguration{}

		// Group=scheduling.hypershift.openshift.io, Version=v1alpha1
	case schedulingv1alpha1.SchemeGroupVersion.WithKind("ClusterSizingConfiguration"):
		return &applyconfigurationschedulingv1alpha1.ClusterSizingConfigurationApplyConfiguration{}
	case schedulingv1alpha1.SchemeGroupVersion.WithKind("ClusterSizingConfigurationSpec"):
		return &applyconfigurationschedulingv1alpha1.ClusterSizingConfigurationSpecApplyConfiguration{}
	case schedulingv1alpha1.SchemeGroupVersion.WithKind("ClusterSizingConfigurationStatus"):
		return &applyconfigurationschedulingv1alpha1.ClusterSizingConfigurationStatusApplyConfiguration{}
	case schedulingv1alpha1.SchemeGroupVersion.WithKind("ConcurrencyConfiguration"):
		return &applyconfigurationschedulingv1alpha1.ConcurrencyConfigurationApplyConfiguration{}
	case schedulingv1alpha1.SchemeGroupVersion.WithKind("Effects"):
		return &applyconfigurationschedulingv1alpha1.EffectsApplyConfiguration{}
	case schedulingv1alpha1.SchemeGroupVersion.WithKind("NodeCountCriteria"):
		return &applyconfigurationschedulingv1alpha1.NodeCountCriteriaApplyConfiguration{}
	case schedulingv1alpha1.SchemeGroupVersion.WithKind("SizeConfiguration"):
		return &applyconfigurationschedulingv1alpha1.SizeConfigurationApplyConfiguration{}
	case schedulingv1alpha1.SchemeGroupVersion.WithKind("TransitionDelayConfiguration"):
		return &applyconfigurationschedulingv1alpha1.TransitionDelayConfigurationApplyConfiguration{}

	}
	return nil
}
