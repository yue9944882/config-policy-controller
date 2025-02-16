// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package e2e

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"open-cluster-management.io/config-policy-controller/test/utils"
)

const (
	case4ConfigPolicyName       string = "openshift-upgrade-channel-e2e"
	case4PolicyYaml             string = "../resources/case4_clusterversion/case4_clusterversion_create.yaml"
	case4ConfigPolicyNameInform string = "openshift-upgrade-channel-inform"
	case4PolicyYamlInform       string = "../resources/case4_clusterversion/case4_clusterversion_inform.yaml"
	case4ConfigPolicyNamePatch  string = "openshift-upgrade-channel-patch"
	case4PolicyYamlPatch        string = "../resources/case4_clusterversion/case4_clusterversion_patch.yaml"
)

var _ = Describe("Test cluster version obj template handling", func() {
	Describe("enforce patch on unnamespaced resource clusterversion "+testNamespace, func() {
		It("should be created properly on the managed cluster", func() {
			By("Creating " + case4ConfigPolicyName + " on managed")
			utils.Kubectl("apply", "-f", case4PolicyYaml, "-n", testNamespace)
			plc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
				case4ConfigPolicyName, testNamespace, true, defaultTimeoutSeconds)
			Expect(plc).NotTo(BeNil())
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case4ConfigPolicyName, testNamespace, true, defaultTimeoutSeconds)

				return utils.GetComplianceState(managedPlc)
			}, defaultTimeoutSeconds, 1).Should(Equal("Compliant"))
			utils.Kubectl("delete", "configurationpolicy", case4ConfigPolicyName, "-n", testNamespace)
		})
		It("should be patched properly on the managed cluster", func() {
			By("Creating " + case4ConfigPolicyNamePatch + " on managed")
			utils.Kubectl("apply", "-f", case4PolicyYamlPatch, "-n", testNamespace)
			plc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
				case4ConfigPolicyNamePatch, testNamespace, true, defaultTimeoutSeconds)
			Expect(plc).NotTo(BeNil())
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case4ConfigPolicyNamePatch, testNamespace, true, defaultTimeoutSeconds)

				return utils.GetComplianceState(managedPlc)
			}, defaultTimeoutSeconds, 1).Should(Equal("Compliant"))
		})
		It("should be generate status properly for cluster-level resources", func() {
			By("Creating " + case4ConfigPolicyNameInform + " on managed")
			utils.Kubectl("apply", "-f", case4PolicyYamlInform, "-n", testNamespace)
			plc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
				case4ConfigPolicyNameInform, testNamespace, true, defaultTimeoutSeconds)
			Expect(plc).NotTo(BeNil())
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case4ConfigPolicyNameInform, testNamespace, true, defaultTimeoutSeconds)

				return utils.GetStatusMessage(managedPlc)
			}, 120, 1).Should(Equal(
				"clusterversions [version] found as specified, therefore this Object template is compliant"))
		})
	})
})
