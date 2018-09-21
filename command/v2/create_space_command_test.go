package v2_test

import (
	"errors"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/translatableerror"
	. "code.cloudfoundry.org/cli/command/v2"
	"code.cloudfoundry.org/cli/command/v2/v2fakes"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("CreateSpaceCommand", func() {
	var (
		fakeConfig      *commandfakes.FakeConfig
		fakeActor       *v2fakes.FakeCreateSpaceActor
		fakeSharedActor *commandfakes.FakeSharedActor
		testUI          *ui.UI
		spaceName       string
		cmd             CreateSpaceCommand

		executeErr error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeActor = new(v2fakes.FakeCreateSpaceActor)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		spaceName = "some-space"

		cmd = CreateSpaceCommand{
			UI:           testUI,
			Config:       fakeConfig,
			Actor:        fakeActor,
			SharedActor:  fakeSharedActor,
			RequiredArgs: flag.Space{Space: spaceName},
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	//TODO: Delete me when this command is no longer experimental
	When("The experimental flag is not set", func() {
		It("is experimental", func() {
			Expect(executeErr).To(MatchError(translatableerror.UnrefactoredCommandError{}))
		})
	})

	When("the experimental flag is set", func() {
		BeforeEach(func() {
			fakeConfig.ExperimentalReturns(true)
		})

		It("checks for user being logged in", func() {
			Expect(fakeSharedActor.RequireCurrentUserCallCount()).To(Equal(1))
		})

		When("user is not logged in", func() {
			expectedErr := errors.New("not logged in and/or can't verify login because of error")

			BeforeEach(func() {
				fakeSharedActor.RequireCurrentUserReturns("", expectedErr)
			})

			It("returns the error", func() {
				Expect(executeErr).To(MatchError(expectedErr))
			})
		})

		When("user is logged in", func() {
			var username string

			BeforeEach(func() {
				username = "some-guy"
				fakeSharedActor.RequireCurrentUserReturns(username, nil)
			})

			When("user specifies an org using the -o flag", func() {
				var specifiedOrgName string

				BeforeEach(func() {
					specifiedOrgName = "specified-org"
					cmd.Organization = specifiedOrgName
				})

				It("does not require a targeted org", func() {
					Expect(fakeSharedActor.RequireTargetedOrgCallCount()).To(Equal(0))
				})

				It("creates a space in the specified org with no errors", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(testUI.Out).To(Say(`Creating space %s in org %s as %s\.\.\.`, spaceName, specifiedOrgName, username))

					Expect(testUI.Out).To(Say("OK\n\n"))

					Expect(fakeActor.CreateSpaceCallCount()).To(Equal(1))
					inputSpace, inputOrg, _ := fakeActor.CreateSpaceArgsForCall(0)
					Expect(inputSpace).To(Equal(spaceName))
					Expect(inputOrg).To(Equal(specifiedOrgName))
				})

				When("an org is targeted", func() {
					BeforeEach(func() {
						fakeSharedActor.RequireTargetedOrgReturns("do-not-use-this-org", nil)
					})

					It("uses the specified org, not the targeted org", func() {
						Expect(testUI.Out).To(Say(`Creating space %s in org %s as %s\.\.\.`, spaceName, specifiedOrgName, username))
						_, inputOrg, _ := fakeActor.CreateSpaceArgsForCall(0)
						Expect(inputOrg).To(Equal(specifiedOrgName))
					})
				})
			})

			When("no org is specified using the -o flag", func() {
				It("requires a targeted org", func() {
					Expect(fakeSharedActor.RequireTargetedOrgCallCount()).To(Equal(1))
				})

				When("no org is targeted", func() {
					BeforeEach(func() {
						fakeSharedActor.RequireTargetedOrgReturns("", errors.New("check target error"))
					})

					It("returns an error", func() {
						Expect(executeErr).To(MatchError("check target error"))
					})
				})

				When("an org is targeted", func() {
					var targetedOrgName string

					BeforeEach(func() {
						fakeSharedActor.RequireTargetedOrgReturns(targetedOrgName, nil)
					})

					It("attempts to create a space with the specified name in the targeted org", func() {
						Expect(fakeActor.CreateSpaceCallCount()).To(Equal(1))
						inputSpace, inputOrg, _ := fakeActor.CreateSpaceArgsForCall(0)
						Expect(inputSpace).To(Equal(spaceName))
						Expect(inputOrg).To(Equal(targetedOrgName))
					})

					When("creating the space succeeds", func() {
						BeforeEach(func() {
							fakeActor.CreateSpaceReturns(
								v2action.Space{},
								v2action.Warnings{"warn-1", "warn-2"},
								nil,
							)
						})

						It("creates the org and displays warnings", func() {
							Expect(executeErr).ToNot(HaveOccurred())
							Expect(testUI.Out).To(Say(`Creating space %s in org %s as %s\.\.\.`, spaceName, targetedOrgName, username))

							Expect(testUI.Err).To(Say("warn-1\nwarn-2\n"))
							Expect(testUI.Out).To(Say("OK\n\n"))
						})

					})

					When("creating the space fails", func() {
						BeforeEach(func() {
							fakeActor.CreateSpaceReturns(
								v2action.Space{},
								v2action.Warnings{"some warning"},
								errors.New("some error"),
							)
						})

						It("should print the warnings and return the error", func() {
							Expect(executeErr).To(MatchError("some error"))
							Expect(testUI.Err).To(Say("some warning\n"))
						})
					})

					When("quota is not specified", func() {
						BeforeEach(func() {
							cmd.Quota = ""
						})
						It("attempts to create the space with no quota", func() {
							_, _, inputQuota := fakeActor.CreateSpaceArgsForCall(0)
							Expect(inputQuota).To(BeEmpty())
						})

					})

					When("quota is specified", func() {
						BeforeEach(func() {
							cmd.Quota = "some-quota"
						})

						It("attempts to create the space with no quota", func() {
							_, _, inputQuota := fakeActor.CreateSpaceArgsForCall(0)
							Expect(inputQuota).To(Equal("some-quota"))
						})
					})
				})
			})
		})
	})
})
