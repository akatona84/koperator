package e2e

import (
	"fmt"

	"github.com/gruntwork-io/terratest/modules/k8s"
)

type dependencyCRDsType struct {
	zookeeper   []string
	prometheus  []string
	certManager []string
}

func (c *dependencyCRDsType) Zookeeper() []string {
	return c.zookeeper
}
func (c *dependencyCRDsType) Prometheus() []string {
	return c.prometheus
}
func (c *dependencyCRDsType) CertManager() []string {
	return c.certManager
}

func (c *dependencyCRDsType) Initialize(kubectlOptions k8s.KubectlOptions) error {
	var err error
	c.certManager, err = listK8sResourceKinds(kubectlOptions, apiGroupKoperatorDependencies()["cert-manager"])
	if len(c.certManager) == 0 {
		if err != nil {
			return fmt.Errorf("initialize Cert-manager CRDs error: %w", err)
		}
		return fmt.Errorf("Cert-manager CRDs %w", ErrorNotFound)
	}
	c.prometheus, err = listK8sResourceKinds(kubectlOptions, apiGroupKoperatorDependencies()["prometheus"])
	if len(c.prometheus) == 0 {
		if err != nil {
			return fmt.Errorf("initialize Prometheus CRDs error: %w", err)
		}
		return fmt.Errorf("Prometheus CRDs %w", ErrorNotFound)
	}
	c.zookeeper, err = listK8sResourceKinds(kubectlOptions, apiGroupKoperatorDependencies()["zookeeper"])
	if len(c.zookeeper) == 0 {
		if err != nil {
			return fmt.Errorf("initialize Zookeeper CRDs error: %w", err)
		}
		return fmt.Errorf("Zookeeper CRDs %w", ErrorNotFound)
	}
	return nil
}