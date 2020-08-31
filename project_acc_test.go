package aiven

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"strconv"
)

var _ = Describe("Projects", func() {
	var (
		projectName string
		project     *Project
		err         error
	)

	BeforeEach(func() {
		projectName = "test-acc-pr" + strconv.Itoa(rand.Int())
		project, err = client.Projects.Create(CreateProjectRequest{
			Project:         projectName,
			BillingCurrency: "EUR",
			TechnicalEmails: ContactEmailFromStringSlice([]string{"test@example.com"}),
			BillingEmails:   ContactEmailFromStringSlice([]string{"test123@example.com"}),
		})
	})

	Context("Create new project", func() {
		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should populate fields properly", func() {
			Expect(project).NotTo(BeNil())

			if project != nil {
				Expect(project.Name).NotTo(BeEmpty())
				Expect(project.AccountId).To(BeEmpty())
				Expect(project.BillingCurrency).NotTo(BeEmpty())
				Expect(project.GetBillingEmailsAsStringSlice()).NotTo(BeEmpty())
				Expect(project.GetTechnicalEmailsAsStringSlice()).NotTo(BeEmpty())
			}
		})
	})

	AfterEach(func() {
		err = client.Projects.Delete(projectName)
		if err != nil {
			Fail("cannot delete project : " + err.Error())
		}
	})
})
