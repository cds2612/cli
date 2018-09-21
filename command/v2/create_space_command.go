package v2

import (
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/translatableerror"
	"code.cloudfoundry.org/cli/command/v2/shared"
)

//go:generate counterfeiter . CreateSpaceActor

type CreateSpaceActor interface {
	CreateSpace(spaceName, orgName, quotaName string) (v2action.Space, v2action.Warnings, error)
	//GrantOrgManagerByUsername(guid string, username string) (v2action.Warnings, error)
}

type CreateSpaceCommand struct {
	RequiredArgs    flag.Space  `positional-args:"yes"`
	Organization    string      `short:"o" description:"Organization"`
	Quota           string      `short:"q" description:"Quota to assign to the newly created space"`
	usage           interface{} `usage:"CF_NAME create-space SPACE [-o ORG] [-q SPACE_QUOTA]"`
	relatedCommands interface{} `related_commands:"set-space-isolation-segment, space-quotas, spaces, target"`

	UI          command.UI
	Config      command.Config
	Actor       CreateSpaceActor
	SharedActor command.SharedActor
}

func (cmd *CreateSpaceCommand) Setup(config command.Config, ui command.UI) error {
	ccClient, uaaClient, err := shared.NewClients(config, ui, true)
	if err != nil {
		return err
	}

	cmd.Actor = v2action.NewActor(ccClient, uaaClient, config)
	cmd.SharedActor = sharedaction.NewActor(config)
	cmd.Config = config
	cmd.UI = ui
	return nil
}

func (cmd CreateSpaceCommand) Execute(args []string) error {
	if !cmd.Config.Experimental() {
		return translatableerror.UnrefactoredCommandError{}
	}

	spaceName := cmd.RequiredArgs.Space
	userName, err := cmd.SharedActor.RequireCurrentUser()
	if err != nil {
		return err
	}

	var orgName string
	if cmd.Organization == "" {
		orgName, err = cmd.SharedActor.RequireTargetedOrg()
		if err != nil {
			return err
		}
	} else {
		orgName = cmd.Organization
	}

	cmd.UI.DisplayTextWithFlavor("Creating space {{.Space}} in org {{.Org}} as {{.User}}...", map[string]interface{}{
		"Space": spaceName,
		"Org":   orgName,
		"User":  userName,
	})

	_, warnings, err := cmd.Actor.CreateSpace(spaceName, orgName, cmd.Quota)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	cmd.UI.DisplayOK()
	cmd.UI.DisplayNewline()

	return nil
}
