package fakerp

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/openshift/openshift-azure/test/clients/azure"
	"github.com/openshift/openshift-azure/test/sanity"
)

var _ = Describe("Change OpenShift Component Log Level E2E tests [ChangeLogLevel][LongRunning]", func() {
	var (
		ctx = context.Background()
	)
	It("should be possible for an SRE to update the OpenShift component log level of a cluster", func() {
		By("Reading the internal config before the log level update")
		before, err := azure.RPClient.OpenShiftManagedClustersAdmin.Get(ctx, os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"))
		Expect(err).NotTo(HaveOccurred())
		Expect(before).NotTo(BeNil())

		By("Executing a cluster update with updated log levels.")
		before.Config.ComponentLogLevel.APIServer = to.IntPtr(*before.Config.ComponentLogLevel.APIServer - 2)
		before.Config.ComponentLogLevel.ControllerManager = to.IntPtr(*before.Config.ComponentLogLevel.ControllerManager - 2)
		before.Config.ComponentLogLevel.Node = to.IntPtr(*before.Config.ComponentLogLevel.Node - 2)
		update, err := azure.RPClient.OpenShiftManagedClustersAdmin.CreateOrUpdate(ctx, os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"), before)
		Expect(err).NotTo(HaveOccurred())
		Expect(update).NotTo(BeNil())

		By("Reading the internal config after the log level update")
		after, err := azure.RPClient.OpenShiftManagedClustersAdmin.Get(ctx, os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"))
		Expect(err).NotTo(HaveOccurred())
		Expect(after).NotTo(BeNil())

		By("Verifying that the cluster log level has been updated")
		Expect(after.Config.ComponentLogLevel).To(Equal(before.Config.ComponentLogLevel))

		By("Validating the cluster")
		errs := sanity.Checker.ValidateCluster(context.Background())
		Expect(errs).To(BeEmpty())
	})
})
