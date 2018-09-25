package v2action_test

import (
	"errors"

	"code.cloudfoundry.org/cli/actor/actionerror"
	. "code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v2action/v2actionfakes"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Space", func() {
	Describe("SpaceNotFoundError#Error", func() {
		When("the name is specified", func() {
			It("returns an error message with the name of the missing space", func() {
				err := actionerror.SpaceNotFoundError{
					Name: "some-space",
				}
				Expect(err.Error()).To(Equal("Space 'some-space' not found."))
			})
		})

		When("the name is not specified, but the GUID is specified", func() {
			It("returns an error message with the GUID of the missing space", func() {
				err := actionerror.SpaceNotFoundError{
					GUID: "some-space-guid",
				}
				Expect(err.Error()).To(Equal("Space with GUID 'some-space-guid' not found."))
			})
		})

		When("neither the name nor the GUID is specified", func() {
			It("returns a generic error message for the missing space", func() {
				err := actionerror.SpaceNotFoundError{}
				Expect(err.Error()).To(Equal("Space '' not found."))
			})
		})
	})

	Describe("Actions", func() {
		var (
			actor                     *Actor
			fakeCloudControllerClient *v2actionfakes.FakeCloudControllerClient
		)

		BeforeEach(func() {
			fakeCloudControllerClient = new(v2actionfakes.FakeCloudControllerClient)
			actor = NewActor(fakeCloudControllerClient, nil, nil)
		})

		Describe("CreateSpace", func() {
			When("the org is not found", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetOrganizationsReturns(
						[]ccv2.Organization{},
						ccv2.Warnings{
							"warning-1",
							"warning-2",
						},
						actionerror.OrganizationNotFoundError{Name: "not-an-org-name"},
					)
				})

				It("returns an OrganizationNotFoundError", func() {
					_, warnings, err := actor.CreateSpace("some-space", "not-an-org-name", "quota-name")
					Expect(err).To(MatchError(actionerror.OrganizationNotFoundError{Name: "not-an-org-name"}))
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
				})
			})

			When("the org is found", func() {
				var expectedSpace Space
				BeforeEach(func() {
					expectedSpace = Space{
						GUID: "some-space-guid",
						Name: "some-space",
					}

					fakeCloudControllerClient.CreateSpaceReturns(
						ccv2.Space{
							GUID: "some-space-guid",
							Name: "some-space",
						},
						ccv2.Warnings{"create-space-warning-1", "create-space-warning-2"},
						nil)

					fakeCloudControllerClient.GetOrganizationsReturns([]ccv2.Organization{
						ccv2.Organization{GUID: "some-org-guid", Name: "some-org-name"}},
						ccv2.Warnings{},
						nil,
					)
				})

				It("should return the space and all its warnings", func() {
					space, warnings, err := actor.CreateSpace("some-space", "some-org-name", "")
					spaceName, orgGuid := fakeCloudControllerClient.CreateSpaceArgsForCall(0)
					Expect(spaceName).To(Equal("some-space"))
					Expect(orgGuid).To(Equal("some-org-guid"))

					Expect(space).To(Equal(expectedSpace))
					Expect(warnings).To(ConsistOf("create-space-warning-1", "create-space-warning-2"))
					Expect(err).ToNot(HaveOccurred())
				})

				When("a space quota is provided", func() {
					When("the quota is not found", func() {
						BeforeEach(func() {
							fakeCloudControllerClient.GetSpaceQuotasReturns(
								[]ccv2.SpaceQuota{},
								ccv2.Warnings{
									"get-quota-warning-1",
									"get-quota-warning-2",
								},
								nil,
							)
						})

						It("returns a SpaceQuotaNotFoundError", func() {
							_, warnings, err := actor.CreateSpace("some-space", "some-org-name", "not-a-space-quota")
							Expect(err).To(MatchError(actionerror.SpaceQuotaNotFoundError{Name: "not-a-space-quota"}))
							Expect(warnings).To(ContainElement("get-quota-warning-1"))
							Expect(warnings).To(ContainElement("get-quota-warning-2"))
						})
					})

					When("the quota is found", func() {
						var spaceQuotaName = "space-quota-name"
						var spaceQuotaGuid = "space-quota-guid"
						BeforeEach(func() {
							fakeCloudControllerClient.GetSpaceQuotasReturns(
								[]ccv2.SpaceQuota{
									ccv2.SpaceQuota{Name: spaceQuotaName, GUID: spaceQuotaGuid},
								},
								ccv2.Warnings{
									"get-quota-warning-1",
									"get-quota-warning-2",
								},
								nil,
							)
						})

						It("sets the quota on the space", func() {
							_, warnings, err := actor.CreateSpace("some-space", "some-org", spaceQuotaName)
							Expect(warnings).To(ContainElement("get-quota-warning-1"))
							Expect(warnings).To(ContainElement("get-quota-warning-2"))
							Expect(err).ToNot(HaveOccurred())
							Expect(fakeCloudControllerClient.SetSpaceQuotaCallCount()).To(Equal(1))
						})
					})
				})

				When("a space quota is not provided", func() {
					It("does not call the cloud controller method to get space quotas", func() {
						_, _, _ = actor.CreateSpace("some-space", "some-org-name", "")
						Expect(fakeCloudControllerClient.GetSpaceQuotasCallCount()).To(Equal(0))
					})
				})
			})

		})

		Describe("DeleteSpaceByNameAndOrganizationName", func() {
			var (
				warnings Warnings
				err      error
			)

			JustBeforeEach(func() {
				warnings, err = actor.DeleteSpaceByNameAndOrganizationName("some-space", "some-org")
			})

			When("the org is not found", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetOrganizationsReturns(
						[]ccv2.Organization{},
						ccv2.Warnings{
							"warning-1",
							"warning-2",
						},
						nil,
					)
				})

				It("returns an OrganizationNotFoundError", func() {
					Expect(err).To(MatchError(actionerror.OrganizationNotFoundError{Name: "some-org"}))
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
				})
			})

			When("the org is found", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetOrganizationsReturns(
						[]ccv2.Organization{{Name: "some-org", GUID: "some-org-guid"}},
						ccv2.Warnings{"warning-1", "warning-2"},
						nil,
					)
				})

				When("the space is not found", func() {
					BeforeEach(func() {
						fakeCloudControllerClient.GetSpacesReturns(
							[]ccv2.Space{},
							ccv2.Warnings{"warning-3", "warning-4"},
							nil,
						)
					})

					It("returns an SpaceNotFoundError", func() {
						Expect(err).To(MatchError(actionerror.SpaceNotFoundError{Name: "some-space"}))
						Expect(warnings).To(ConsistOf("warning-1", "warning-2", "warning-3", "warning-4"))
					})
				})

				When("the space is found", func() {
					BeforeEach(func() {
						fakeCloudControllerClient.GetSpacesReturns(
							[]ccv2.Space{{GUID: "some-space-guid"}},
							ccv2.Warnings{"warning-3", "warning-4"},
							nil,
						)
					})

					When("the delete returns an error", func() {
						var expectedErr error

						BeforeEach(func() {
							expectedErr = errors.New("some delete space error")
							fakeCloudControllerClient.DeleteSpaceJobReturns(
								ccv2.Job{},
								ccv2.Warnings{"warning-5", "warning-6"},
								expectedErr,
							)
						})

						It("returns the error", func() {
							Expect(err).To(Equal(expectedErr))
							Expect(warnings).To(ConsistOf("warning-1", "warning-2", "warning-3", "warning-4", "warning-5", "warning-6"))
						})
					})

					When("the delete returns a job", func() {
						BeforeEach(func() {
							fakeCloudControllerClient.DeleteSpaceJobReturns(
								ccv2.Job{GUID: "some-job-guid"},
								ccv2.Warnings{"warning-5", "warning-6"},
								nil,
							)
						})

						When("polling errors", func() {
							var expectedErr error

							BeforeEach(func() {
								expectedErr = errors.New("Never expected, by anyone")
								fakeCloudControllerClient.PollJobReturns(ccv2.Warnings{"warning-7", "warning-8"}, expectedErr)
							})

							It("returns the error", func() {
								Expect(err).To(Equal(expectedErr))
								Expect(warnings).To(ConsistOf("warning-1", "warning-2", "warning-3", "warning-4", "warning-5", "warning-6", "warning-7", "warning-8"))
							})
						})

						When("the job is successful", func() {
							BeforeEach(func() {
								fakeCloudControllerClient.PollJobReturns(ccv2.Warnings{"warning-7", "warning-8"}, nil)
							})

							It("returns warnings and no error", func() {
								Expect(err).ToNot(HaveOccurred())
								Expect(warnings).To(ConsistOf("warning-1", "warning-2", "warning-3", "warning-4", "warning-5", "warning-6", "warning-7", "warning-8"))

								Expect(fakeCloudControllerClient.GetOrganizationsCallCount()).To(Equal(1))
								Expect(fakeCloudControllerClient.GetOrganizationsArgsForCall(0)).To(Equal([]ccv2.Filter{{
									Type:     constant.NameFilter,
									Operator: constant.EqualOperator,
									Values:   []string{"some-org"},
								}}))

								Expect(fakeCloudControllerClient.GetSpacesCallCount()).To(Equal(1))
								Expect(fakeCloudControllerClient.GetSpacesArgsForCall(0)).To(Equal([]ccv2.Filter{{
									Type:     constant.NameFilter,
									Operator: constant.EqualOperator,
									Values:   []string{"some-space"},
								},
									{
										Type:     constant.OrganizationGUIDFilter,
										Operator: constant.EqualOperator,
										Values:   []string{"some-org-guid"},
									},
								}))

								Expect(fakeCloudControllerClient.DeleteSpaceJobCallCount()).To(Equal(1))
								Expect(fakeCloudControllerClient.DeleteSpaceJobArgsForCall(0)).To(Equal("some-space-guid"))

								Expect(fakeCloudControllerClient.PollJobCallCount()).To(Equal(1))
								Expect(fakeCloudControllerClient.PollJobArgsForCall(0)).To(Equal(ccv2.Job{GUID: "some-job-guid"}))
							})
						})
					})
				})
			})
		})

		Describe("GetOrganizationSpaces", func() {
			When("there are spaces in the org", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetSpacesReturns(
						[]ccv2.Space{
							{
								GUID:                     "space-1-guid",
								Name:                     "space-1",
								AllowSSH:                 true,
								SpaceQuotaDefinitionGUID: "some-space-quota-guid",
							},
							{
								GUID:     "space-2-guid",
								Name:     "space-2",
								AllowSSH: false,
							},
						},
						ccv2.Warnings{"warning-1", "warning-2"},
						nil)
				})

				It("returns all spaces and all warnings", func() {
					spaces, warnings, err := actor.GetOrganizationSpaces("some-org-guid")

					Expect(err).ToNot(HaveOccurred())
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
					Expect(spaces).To(Equal(
						[]Space{
							{
								GUID:                     "space-1-guid",
								Name:                     "space-1",
								AllowSSH:                 true,
								SpaceQuotaDefinitionGUID: "some-space-quota-guid",
							},
							{
								GUID:     "space-2-guid",
								Name:     "space-2",
								AllowSSH: false,
							},
						}))

					Expect(fakeCloudControllerClient.GetSpacesCallCount()).To(Equal(1))
					Expect(fakeCloudControllerClient.GetSpacesArgsForCall(0)).To(Equal(
						[]ccv2.Filter{
							{
								Type:     constant.OrganizationGUIDFilter,
								Operator: constant.EqualOperator,
								Values:   []string{"some-org-guid"},
							},
						}))
				})
			})

			When("an error is encountered", func() {
				var returnedErr error

				BeforeEach(func() {
					returnedErr = errors.New("cc-get-spaces-error")
					fakeCloudControllerClient.GetSpacesReturns(
						[]ccv2.Space{},
						ccv2.Warnings{"warning-1", "warning-2"},
						returnedErr,
					)
				})

				It("returns the error and all warnings", func() {
					_, warnings, err := actor.GetOrganizationSpaces("some-org-guid")

					Expect(err).To(MatchError(returnedErr))
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
				})
			})
		})

		Describe("GetSpaceByOrganizationAndName", func() {
			When("the space exists", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetSpacesReturns(
						[]ccv2.Space{
							{
								GUID:                     "some-space-guid",
								Name:                     "some-space",
								AllowSSH:                 true,
								SpaceQuotaDefinitionGUID: "some-space-quota-guid",
							},
						},
						ccv2.Warnings{"warning-1", "warning-2"},
						nil)
				})

				It("returns the space and all warnings", func() {
					space, warnings, err := actor.GetSpaceByOrganizationAndName("some-org-guid", "some-space")

					Expect(err).ToNot(HaveOccurred())
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
					Expect(space).To(Equal(Space{
						GUID:                     "some-space-guid",
						Name:                     "some-space",
						AllowSSH:                 true,
						SpaceQuotaDefinitionGUID: "some-space-quota-guid",
					}))

					Expect(fakeCloudControllerClient.GetSpacesCallCount()).To(Equal(1))
					Expect(fakeCloudControllerClient.GetSpacesArgsForCall(0)).To(ConsistOf(
						[]ccv2.Filter{
							{
								Type:     constant.OrganizationGUIDFilter,
								Operator: constant.EqualOperator,
								Values:   []string{"some-org-guid"},
							},
							{
								Type:     constant.NameFilter,
								Operator: constant.EqualOperator,
								Values:   []string{"some-space"},
							},
						}))
				})
			})

			When("an error is encountered", func() {
				var returnedErr error

				BeforeEach(func() {
					returnedErr = errors.New("cc-get-spaces-error")
					fakeCloudControllerClient.GetSpacesReturns(
						[]ccv2.Space{},
						ccv2.Warnings{"warning-1", "warning-2"},
						returnedErr,
					)
				})

				It("return the error and all warnings", func() {
					_, warnings, err := actor.GetSpaceByOrganizationAndName("some-org-guid", "some-space")

					Expect(err).To(MatchError(returnedErr))
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
				})
			})

			When("the space does not exist", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetSpacesReturns(
						[]ccv2.Space{},
						nil,
						nil,
					)
				})

				It("returns SpaceNotFoundError", func() {
					_, _, err := actor.GetSpaceByOrganizationAndName("some-org-guid", "some-space")

					Expect(err).To(MatchError(actionerror.SpaceNotFoundError{
						Name: "some-space",
					}))
				})
			})

			When("multiple spaces exists", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetSpacesReturns(
						[]ccv2.Space{
							{
								GUID:     "some-space-guid",
								Name:     "some-space",
								AllowSSH: true,
							},
							{
								GUID:     "another-space-guid",
								Name:     "another-space",
								AllowSSH: true,
							},
						},
						nil,
						nil,
					)
				})

				It("returns MultipleSpacesFoundError", func() {
					_, _, err := actor.GetSpaceByOrganizationAndName("some-org-guid", "some-space")

					Expect(err).To(MatchError(actionerror.MultipleSpacesFoundError{
						OrgGUID: "some-org-guid",
						Name:    "some-space",
					}))
				})
			})
		})
	})
})
