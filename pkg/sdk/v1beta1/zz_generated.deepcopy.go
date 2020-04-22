// +build !ignore_autogenerated

// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/banzaicloud/istio-client-go/pkg/networking/v1alpha3"
	metav1 "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertManagerConfig) DeepCopyInto(out *AlertManagerConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertManagerConfig.
func (in *AlertManagerConfig) DeepCopy() *AlertManagerConfig {
	if in == nil {
		return nil
	}
	out := new(AlertManagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Broker) DeepCopyInto(out *Broker) {
	*out = *in
	if in.BrokerConfig != nil {
		in, out := &in.BrokerConfig, &out.BrokerConfig
		*out = new(BrokerConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Broker.
func (in *Broker) DeepCopy() *Broker {
	if in == nil {
		return nil
	}
	out := new(Broker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerConfig) DeepCopyInto(out *BrokerConfig) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageConfigs != nil {
		in, out := &in.StorageConfigs, &out.StorageConfigs
		*out = make([]StorageConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BrokerAnnotations != nil {
		in, out := &in.BrokerAnnotations, &out.BrokerAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerConfig.
func (in *BrokerConfig) DeepCopy() *BrokerConfig {
	if in == nil {
		return nil
	}
	out := new(BrokerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerState) DeepCopyInto(out *BrokerState) {
	*out = *in
	in.GracefulActionState.DeepCopyInto(&out.GracefulActionState)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerState.
func (in *BrokerState) DeepCopy() *BrokerState {
	if in == nil {
		return nil
	}
	out := new(BrokerState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonListenerSpec) DeepCopyInto(out *CommonListenerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonListenerSpec.
func (in *CommonListenerSpec) DeepCopy() *CommonListenerSpec {
	if in == nil {
		return nil
	}
	out := new(CommonListenerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CruiseControlConfig) DeepCopyInto(out *CruiseControlConfig) {
	*out = *in
	out.CruiseControlTaskSpec = in.CruiseControlTaskSpec
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopicConfig != nil {
		in, out := &in.TopicConfig, &out.TopicConfig
		*out = new(TopicConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CruiseControlConfig.
func (in *CruiseControlConfig) DeepCopy() *CruiseControlConfig {
	if in == nil {
		return nil
	}
	out := new(CruiseControlConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CruiseControlTaskSpec) DeepCopyInto(out *CruiseControlTaskSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CruiseControlTaskSpec.
func (in *CruiseControlTaskSpec) DeepCopy() *CruiseControlTaskSpec {
	if in == nil {
		return nil
	}
	out := new(CruiseControlTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfig) DeepCopyInto(out *EnvoyConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfig.
func (in *EnvoyConfig) DeepCopy() *EnvoyConfig {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalListenerConfig) DeepCopyInto(out *ExternalListenerConfig) {
	*out = *in
	out.CommonListenerSpec = in.CommonListenerSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalListenerConfig.
func (in *ExternalListenerConfig) DeepCopy() *ExternalListenerConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalListenerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GracefulActionState) DeepCopyInto(out *GracefulActionState) {
	*out = *in
	if in.VolumeStates != nil {
		in, out := &in.VolumeStates, &out.VolumeStates
		*out = make(map[string]VolumeState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GracefulActionState.
func (in *GracefulActionState) DeepCopy() *GracefulActionState {
	if in == nil {
		return nil
	}
	out := new(GracefulActionState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalListenerConfig) DeepCopyInto(out *InternalListenerConfig) {
	*out = *in
	out.CommonListenerSpec = in.CommonListenerSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalListenerConfig.
func (in *InternalListenerConfig) DeepCopy() *InternalListenerConfig {
	if in == nil {
		return nil
	}
	out := new(InternalListenerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioIngressConfig) DeepCopyInto(out *IstioIngressConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TLSOptions != nil {
		in, out := &in.TLSOptions, &out.TLSOptions
		*out = new(v1alpha3.TLSOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.VirtualServiceAnnotations != nil {
		in, out := &in.VirtualServiceAnnotations, &out.VirtualServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioIngressConfig.
func (in *IstioIngressConfig) DeepCopy() *IstioIngressConfig {
	if in == nil {
		return nil
	}
	out := new(IstioIngressConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaCluster) DeepCopyInto(out *KafkaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaCluster.
func (in *KafkaCluster) DeepCopy() *KafkaCluster {
	if in == nil {
		return nil
	}
	out := new(KafkaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterList) DeepCopyInto(out *KafkaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterList.
func (in *KafkaClusterList) DeepCopy() *KafkaClusterList {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterSpec) DeepCopyInto(out *KafkaClusterSpec) {
	*out = *in
	in.ListenersConfig.DeepCopyInto(&out.ListenersConfig)
	if in.ZKAddresses != nil {
		in, out := &in.ZKAddresses, &out.ZKAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RackAwareness != nil {
		in, out := &in.RackAwareness, &out.RackAwareness
		*out = new(RackAwareness)
		(*in).DeepCopyInto(*out)
	}
	if in.BrokerConfigGroups != nil {
		in, out := &in.BrokerConfigGroups, &out.BrokerConfigGroups
		*out = make(map[string]BrokerConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Brokers != nil {
		in, out := &in.Brokers, &out.Brokers
		*out = make([]Broker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.RollingUpgradeConfig = in.RollingUpgradeConfig
	in.CruiseControlConfig.DeepCopyInto(&out.CruiseControlConfig)
	in.EnvoyConfig.DeepCopyInto(&out.EnvoyConfig)
	out.MonitoringConfig = in.MonitoringConfig
	out.VaultConfig = in.VaultConfig
	if in.AlertManagerConfig != nil {
		in, out := &in.AlertManagerConfig, &out.AlertManagerConfig
		*out = new(AlertManagerConfig)
		**out = **in
	}
	in.IstioIngressConfig.DeepCopyInto(&out.IstioIngressConfig)
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterSpec.
func (in *KafkaClusterSpec) DeepCopy() *KafkaClusterSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaClusterStatus) DeepCopyInto(out *KafkaClusterStatus) {
	*out = *in
	if in.BrokersState != nil {
		in, out := &in.BrokersState, &out.BrokersState
		*out = make(map[string]BrokerState, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.RollingUpgrade = in.RollingUpgrade
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaClusterStatus.
func (in *KafkaClusterStatus) DeepCopy() *KafkaClusterStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenersConfig) DeepCopyInto(out *ListenersConfig) {
	*out = *in
	if in.ExternalListeners != nil {
		in, out := &in.ExternalListeners, &out.ExternalListeners
		*out = make([]ExternalListenerConfig, len(*in))
		copy(*out, *in)
	}
	if in.InternalListeners != nil {
		in, out := &in.InternalListeners, &out.InternalListeners
		*out = make([]InternalListenerConfig, len(*in))
		copy(*out, *in)
	}
	if in.SSLSecrets != nil {
		in, out := &in.SSLSecrets, &out.SSLSecrets
		*out = new(SSLSecrets)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenersConfig.
func (in *ListenersConfig) DeepCopy() *ListenersConfig {
	if in == nil {
		return nil
	}
	out := new(ListenersConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfig) DeepCopyInto(out *MonitoringConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfig.
func (in *MonitoringConfig) DeepCopy() *MonitoringConfig {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackAwareness) DeepCopyInto(out *RackAwareness) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackAwareness.
func (in *RackAwareness) DeepCopy() *RackAwareness {
	if in == nil {
		return nil
	}
	out := new(RackAwareness)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpgradeConfig) DeepCopyInto(out *RollingUpgradeConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpgradeConfig.
func (in *RollingUpgradeConfig) DeepCopy() *RollingUpgradeConfig {
	if in == nil {
		return nil
	}
	out := new(RollingUpgradeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpgradeStatus) DeepCopyInto(out *RollingUpgradeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpgradeStatus.
func (in *RollingUpgradeStatus) DeepCopy() *RollingUpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(RollingUpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSLSecrets) DeepCopyInto(out *SSLSecrets) {
	*out = *in
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(metav1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSLSecrets.
func (in *SSLSecrets) DeepCopy() *SSLSecrets {
	if in == nil {
		return nil
	}
	out := new(SSLSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	if in.PvcSpec != nil {
		in, out := &in.PvcSpec, &out.PvcSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicConfig) DeepCopyInto(out *TopicConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicConfig.
func (in *TopicConfig) DeepCopy() *TopicConfig {
	if in == nil {
		return nil
	}
	out := new(TopicConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultConfig) DeepCopyInto(out *VaultConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultConfig.
func (in *VaultConfig) DeepCopy() *VaultConfig {
	if in == nil {
		return nil
	}
	out := new(VaultConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeState) DeepCopyInto(out *VolumeState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeState.
func (in *VolumeState) DeepCopy() *VolumeState {
	if in == nil {
		return nil
	}
	out := new(VolumeState)
	in.DeepCopyInto(out)
	return out
}