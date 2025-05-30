package cluster

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test validation functions", func() {

	validTestArn := "arn:aws:iam::123456789012:role/test"
	invalidTestArn := "123arn:aws?"

	Context("Validation of HCP shared VPC", func() {
		When("isHostedCp", func() {
			It("OK: Passes validation (all flags filled out and correct format", func() {
				err := validateHcpSharedVpcArgs(validTestArn, validTestArn, "123", "123", false)
				Expect(err).ToNot(HaveOccurred())
			})
			It("KO: Invalid Route53 ARN", func() {
				err := validateHcpSharedVpcArgs(invalidTestArn, validTestArn, "123", "123", false)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(isNotValidArnErrorMsg, route53RoleArnFlag)))
			})
			It("KO: invalid VPC Endpoint ARN", func() {
				err := validateHcpSharedVpcArgs(validTestArn, invalidTestArn, "123", "123", false)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(isNotValidArnErrorMsg, vpcEndpointRoleArnFlag)))
			})
			It("KO: Route53 ARN flag not populated", func() {
				err := validateHcpSharedVpcArgs("", validTestArn, "123", "123", false)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(hcpSharedVpcFlagNotFilledErrorMsg,
					route53RoleArnFlag)))
			})
			It("KO: VPC Endpoint ARN flag not populated", func() {
				err := validateHcpSharedVpcArgs(validTestArn, "", "123", "123", false)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(hcpSharedVpcFlagNotFilledErrorMsg,
					vpcEndpointRoleArnFlag)))
			})
			It("KO: Ingress Private Hosted Zone ID flag not populated", func() {
				err := validateHcpSharedVpcArgs(validTestArn, validTestArn, "", "123", false)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(hcpSharedVpcFlagNotFilledErrorMsg,
					ingressPrivateHostedZoneIdFlag)))
			})
			It("KO: HCP Internal Communication Hosted Zone ID flag not populated", func() {
				err := validateHcpSharedVpcArgs(validTestArn, validTestArn, "123", "", false)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(hcpSharedVpcFlagNotFilledErrorMsg,
					hcpInternalCommunicationHostedZoneIdFlag)))
			})
			It("KO: HCP Shared VPC cluster fails to validate on govcloud/fedramp", func() {
				err := validateHcpSharedVpcArgs(validTestArn, validTestArn, "123", "123", true)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(isNotGovcloudFeature,
					fmt.Sprintf(pleaseRemoveFlags, ""),
				),
				))
				// Must be `ContainSubstring` because for some reason `strings.Join()` puts strings in a random order
				Expect(err.Error()).To(ContainSubstring("'" + route53RoleArnFlag + "'"))
				Expect(err.Error()).To(ContainSubstring("'" + vpcEndpointRoleArnFlag + "'"))
				Expect(err.Error()).To(ContainSubstring("'" + ingressPrivateHostedZoneIdFlag + "'"))
				Expect(err.Error()).To(ContainSubstring("'" + hcpInternalCommunicationHostedZoneIdFlag + "'"))
			})
		})
	})
})
