// Copyright © 2020 Banzai Cloud
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

package tests

import (
	"context"
	"fmt"
	"strconv"

	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/banzaicloud/kafka-operator/api/v1beta1"
	"github.com/banzaicloud/kafka-operator/pkg/resources/kafkamonitoring"
	"github.com/banzaicloud/kafka-operator/pkg/util"
)

func expectKafka(kafkaCluster *v1beta1.KafkaCluster) {
	expectKafkaAllBrokerService(kafkaCluster)
	expectKafkaPDB(kafkaCluster)
	expectKafkaPVC(kafkaCluster)

	for _, broker := range kafkaCluster.Spec.Brokers {
		expectKafkaBrokerConfigmap(kafkaCluster, broker)
		expectKafkaBrokerService(kafkaCluster, broker)
		expectKafkaBrokerPod(kafkaCluster, broker)
	}

	// TODO test reconcile PKI?
	// TODO test reconcileKafkaPodDelete
}

func expectKafkaAllBrokerService(kafkaCluster *v1beta1.KafkaCluster) {
	service := &corev1.Service{}
	Eventually(func() error {
		return k8sClient.Get(context.Background(), types.NamespacedName{
			Namespace: kafkaCluster.Namespace,
			Name:      fmt.Sprintf("%s-all-broker", kafkaCluster.Name),
		}, service)
	}).Should(Succeed())

	Expect(service.Labels).To(HaveKeyWithValue("app", "kafka"))
	Expect(service.Labels).To(HaveKeyWithValue("kafka_cr", kafkaCluster.Name))

	Expect(service.Spec.Type).To(Equal(corev1.ServiceTypeClusterIP))
	Expect(service.Spec.SessionAffinity).To(Equal(corev1.ServiceAffinityNone))
	Expect(service.Spec.Selector).To(HaveKeyWithValue("app", "kafka"))
	Expect(service.Spec.Selector).To(HaveKeyWithValue("kafka_cr", kafkaCluster.Name))
	Expect(service.Spec.Ports).To(ConsistOf(
		corev1.ServicePort{
			Name:       "tcp-internal",
			Protocol:   "TCP",
			Port:       29092,
			TargetPort: intstr.FromInt(29092),
		},
		corev1.ServicePort{
			Name:       "tcp-controller",
			Protocol:   "TCP",
			Port:       29093,
			TargetPort: intstr.FromInt(29093),
		},
		corev1.ServicePort{
			Name:       "tcp-test",
			Protocol:   "TCP",
			Port:       9733,
			TargetPort: intstr.FromInt(9733),
		}))
}

func expectKafkaPDB(kafkaCluster *v1beta1.KafkaCluster) {
	// get current CR
	err := k8sClient.Get(context.TODO(), types.NamespacedName{Name: kafkaCluster.Name, Namespace: kafkaCluster.Namespace}, kafkaCluster)

	// set PDB and reset status
	kafkaCluster.Spec.DisruptionBudget = v1beta1.DisruptionBudget{
		Create: true,
		Budget: "20%",
	}
	kafkaCluster.Status = v1beta1.KafkaClusterStatus{}

	// update CR
	err = k8sClient.Update(context.TODO(), kafkaCluster)
	Expect(err).NotTo(HaveOccurred())

	// wait until reconcile finishes
	waitForClusterRunningState(kafkaCluster, kafkaCluster.Namespace)

	// get created PDB
	pdb := policyv1beta1.PodDisruptionBudget{}
	Eventually(func() error {
		return k8sClient.Get(context.Background(), types.NamespacedName{
			Namespace: kafkaCluster.Namespace,
			Name:      fmt.Sprintf("%s-pdb", kafkaCluster.Name),
		}, &pdb)
	}).Should(Succeed())

	// make assertions
	Expect(pdb.Labels).To(HaveKeyWithValue("app", "kafka"))
	Expect(pdb.Labels).To(HaveKeyWithValue("kafka_cr", kafkaCluster.Name))
	Expect(pdb.Spec.MinAvailable).To(Equal(util.IntstrPointer(1)))
	Expect(pdb.Spec.Selector).NotTo(BeNil())
	Expect(pdb.Spec.Selector.MatchLabels).To(HaveKeyWithValue("app", "kafka"))
	Expect(pdb.Spec.Selector.MatchLabels).To(HaveKeyWithValue("kafka_cr", kafkaCluster.Name))
}

func expectKafkaPVC(kafkaCluster *v1beta1.KafkaCluster) {
	// get PVCs
	pvcs := corev1.PersistentVolumeClaimList{}
	Eventually(func() error {
		return k8sClient.List(context.Background(), &pvcs,
			client.ListOption(client.InNamespace(kafkaCluster.Namespace)),
			client.ListOption(client.MatchingLabels(map[string]string{"app": "kafka", "kafka_cr": kafkaCluster.Name})))
	}).Should(Succeed())

	Expect(pvcs.Items).To(HaveLen(1))
	pvc := pvcs.Items[0]
	Expect(pvc.GenerateName).To(Equal(fmt.Sprintf("%s-0-storage-0-", kafkaCluster.Name)))
	Expect(pvc.Labels).To(HaveKeyWithValue("app", "kafka"))
	Expect(pvc.Labels).To(HaveKeyWithValue("brokerId", "0"))
	Expect(pvc.Labels).To(HaveKeyWithValue("kafka_cr", kafkaCluster.Name))
	Expect(pvc.Annotations).To(HaveKeyWithValue("mountPath", "/kafka-logs"))
	Expect(pvc.Spec.AccessModes).To(ConsistOf(corev1.ReadWriteOnce))
	Expect(pvc.Spec.Resources).To(Equal(corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			"storage": resource.MustParse("10Gi"),
		},
	}))
}

func expectKafkaBrokerConfigmap(kafkaCluster *v1beta1.KafkaCluster, broker v1beta1.Broker) {
	configMap := corev1.ConfigMap{}
	Eventually(func() error {
		return k8sClient.Get(context.Background(), types.NamespacedName{
			Namespace: kafkaCluster.Namespace,
			Name:      fmt.Sprintf("%s-config-%d", kafkaCluster.Name, broker.Id),
		}, &configMap)
	}).Should(Succeed())

	Expect(configMap.Labels).To(HaveKeyWithValue("app", "kafka"))
	Expect(configMap.Labels).To(HaveKeyWithValue("kafka_cr", kafkaCluster.Name))
	Expect(configMap.Labels).To(HaveKeyWithValue("brokerId", strconv.Itoa(int(broker.Id))))

	Expect(configMap.Data).To(HaveKeyWithValue("broker-config", fmt.Sprintf(`advertised.listeners=TEST://test-host:11202,INTERNAL://kafkacluster-1-0.kafka-1.svc.cluster.local:29092,CONTROLLER://kafkacluster-1-0.kafka-1.svc.cluster.local:29093
broker.id=%d
control.plane.listener.name=CONTROLLER
cruise.control.metrics.reporter.bootstrap.servers=INTERNAL://kafkacluster-1-0.kafka-1.svc.cluster.local:29092
cruise.control.metrics.reporter.kubernetes.mode=true
inter.broker.listener.name=INTERNAL
listener.security.protocol.map=INTERNAL:PLAINTEXT,CONTROLLER:PLAINTEXT,TEST:
listeners=INTERNAL://:29092,CONTROLLER://:29093,TEST://:9733
log.dirs=/kafka-logs/kafka
metric.reporters=com.linkedin.kafka.cruisecontrol.metricsreporter.CruiseControlMetricsReporter
zookeeper.connect=/`, broker.Id)))

	// assert log4j?
}

func expectKafkaBrokerService(kafkaCluster *v1beta1.KafkaCluster, broker v1beta1.Broker) {
	service := corev1.Service{}
	Eventually(func() error {
		return k8sClient.Get(context.Background(), types.NamespacedName{
			Namespace: kafkaCluster.Namespace,
			Name:      fmt.Sprintf("%s-%d", kafkaCluster.Name, broker.Id),
		}, &service)
	}).Should(Succeed())

	Expect(service.Labels).To(HaveKeyWithValue("app", "kafka"))
	Expect(service.Labels).To(HaveKeyWithValue("kafka_cr", kafkaCluster.Name))
	Expect(service.Labels).To(HaveKeyWithValue("brokerId", strconv.Itoa(int(broker.Id))))

	Expect(service.Spec.Ports).To(ConsistOf(
		corev1.ServicePort{
			Name:       "tcp-internal",
			Protocol:   "TCP",
			Port:       29092,
			TargetPort: intstr.FromInt(29092),
		},
		corev1.ServicePort{
			Name:       "tcp-controller",
			Protocol:   "TCP",
			Port:       29093,
			TargetPort: intstr.FromInt(29093),
		},
		corev1.ServicePort{
			Name:       "tcp-test",
			Protocol:   "TCP",
			Port:       9733,
			TargetPort: intstr.FromInt(9733),
		},
		corev1.ServicePort{
			Name:       "metrics",
			Protocol:   "TCP",
			Port:       9020,
			TargetPort: intstr.FromInt(9020),
		}))

	Expect(service.Spec.Selector).To(HaveKeyWithValue("app", "kafka"))
	Expect(service.Spec.Selector).To(HaveKeyWithValue("kafka_cr", kafkaCluster.Name))
	Expect(service.Spec.Selector).To(HaveKeyWithValue("brokerId", strconv.Itoa(int(broker.Id))))
	Expect(service.Spec.Type).To(Equal(corev1.ServiceTypeClusterIP))
}

func expectKafkaBrokerPod(kafkaCluster *v1beta1.KafkaCluster, broker v1beta1.Broker) {
	podList := corev1.PodList{}
	Eventually(func() ([]corev1.Pod, error) {
		err := k8sClient.List(context.Background(), &podList,
			client.ListOption(client.InNamespace(kafkaCluster.Namespace)),
			client.ListOption(client.MatchingLabels(map[string]string{"app": "kafka", "kafka_cr": kafkaCluster.Name})))
		return podList.Items, err
	}).Should(HaveLen(1))

	pod := podList.Items[0]

	Expect(pod.GenerateName).To(Equal(fmt.Sprintf("%s-%d-", kafkaCluster.Name, broker.Id)))
	Expect(pod.Labels).To(HaveKeyWithValue("brokerId", strconv.Itoa(int(broker.Id))))
	getContainerName := func(c corev1.Container) string { return c.Name }
	Expect(pod.Spec.InitContainers).To(ConsistOf(
		WithTransform(getContainerName, Equal("cruise-control-reporter")),
		WithTransform(getContainerName, Equal("jmx-exporter"))))

	Expect(pod.Spec.Affinity).NotTo(BeNil())
	Expect(pod.Spec.Affinity.PodAntiAffinity).NotTo(BeNil())

	Expect(pod.Spec.Containers).To(HaveLen(1))
	container := pod.Spec.Containers[0]
	Expect(container.Name).To(Equal("kafka"))
	Expect(container.Image).To(Equal("ghcr.io/banzaicloud/kafka:2.13-2.6.0-bzc.1"))
	Expect(container.Lifecycle).NotTo(BeNil())
	Expect(container.Lifecycle.PreStop).NotTo(BeNil())
	getEnvName := func(c corev1.EnvVar) string { return c.Name }
	Expect(container.Env).To(ConsistOf(
		// the exact value is not interesting
		WithTransform(getEnvName, Equal("CLASSPATH")),
		WithTransform(getEnvName, Equal("KAFKA_OPTS")),
		WithTransform(getEnvName, Equal("KAFKA_JVM_PERFORMANCE_OPTS")),

		// the exact value should be checked
		corev1.EnvVar{
			Name: "ENVOY_SIDECAR_STATUS",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					APIVersion: "v1",
					FieldPath:  `metadata.annotations['sidecar.istio.io/status']`,
				},
			},
		},
		corev1.EnvVar{
			Name:  "KAFKA_HEAP_OPTS",
			Value: "-Xmx2G -Xms2G",
		},
	))

	getVolumeName := func(vol corev1.Volume) string { return vol.Name }
	Expect(pod.Spec.Volumes).To(ConsistOf(
		corev1.Volume{
			Name: "exitfile",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		},
		corev1.Volume{
			Name: "broker-config",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{Name: fmt.Sprintf("%s-config-%d", kafkaCluster.Name, broker.Id)},
					DefaultMode:          util.Int32Pointer(0644),
				},
			},
		},
		corev1.Volume{
			Name: "extensions",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		},
		corev1.Volume{
			Name: fmt.Sprintf(kafkamonitoring.BrokerJmxTemplate, kafkaCluster.Name),
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{Name: fmt.Sprintf(kafkamonitoring.BrokerJmxTemplate, kafkaCluster.Name)},
					DefaultMode:          util.Int32Pointer(0644),
				},
			},
		},
		corev1.Volume{
			Name: "jmx-jar-data",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		},

		// the name of the PVC is dynamically created - no exact match
		WithTransform(getVolumeName, Equal(fmt.Sprintf("kafka-data-%d", broker.Id))),
	))

	Expect(pod.Spec.RestartPolicy).To(Equal(corev1.RestartPolicyNever))
	Expect(pod.Spec.TerminationGracePeriodSeconds).To(Equal(util.Int64Pointer(120)))

	// expect some other fields
}