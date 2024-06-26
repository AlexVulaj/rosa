package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/openshift/rosa/tests/ci/labels"
	"github.com/openshift/rosa/tests/utils/exec/rosacli"
	ph "github.com/openshift/rosa/tests/utils/profilehandler"
)

var _ = Describe("ROSA CLI Test", func() {
	It("DestroyClusterByProfile",
		labels.Critical,
		labels.Destroy,
		func() {
			client := rosacli.NewClient()
			profile := ph.LoadProfileYamlFileByENV()
			var errs = ph.DestroyResourceByProfile(profile, client)
			Expect(len(errs)).To(Equal(0))
		})
})
