/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"github.com/spf13/cobra"

	// common imports for subcommands
	cmdgenerate "guestbook-app/cmd/gbkctl/commands/generate"
	cmdinit "guestbook-app/cmd/gbkctl/commands/init"
	cmdversion "guestbook-app/cmd/gbkctl/commands/version"

	// specific imports for workloads
	generateapps "guestbook-app/cmd/gbkctl/commands/generate/apps"
	initapps "guestbook-app/cmd/gbkctl/commands/init/apps"
	versionapps "guestbook-app/cmd/gbkctl/commands/version/apps"
	//+kubebuilder:scaffold:operator-builder:subcommands:imports
)

// GbkctlCommand represents the base command when called without any subcommands.
type GbkctlCommand struct {
	*cobra.Command
}

// NewGbkctlCommand returns an instance of the GbkctlCommand.
func NewGbkctlCommand() *GbkctlCommand {
	c := &GbkctlCommand{
		Command: &cobra.Command{
			Use:   "gbkctl",
			Short: "Deploys a sample kubernetes guestbook application",
			Long:  "Deploys a sample kubernetes guestbook application",
		},
	}

	c.addSubCommands()

	return c
}

// Run represents the main entry point into the command
// This is called by main.main() to execute the root command.
func (c *GbkctlCommand) Run() {
	cobra.CheckErr(c.Execute())
}

func (c *GbkctlCommand) newInitSubCommand() {
	parentCommand := cmdinit.GetParent(c.Command)
	_ = parentCommand

	// add the init subcommands
	initapps.NewGuestbookSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:init
}

func (c *GbkctlCommand) newGenerateSubCommand() {
	parentCommand := cmdgenerate.GetParent(c.Command)
	_ = parentCommand

	// add the generate subcommands
	generateapps.NewGuestbookSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:generate
}

func (c *GbkctlCommand) newVersionSubCommand() {
	parentCommand := cmdversion.GetParent(c.Command)
	_ = parentCommand

	// add the version subcommands
	versionapps.NewGuestbookSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:version
}

// addSubCommands adds any additional subCommands to the root command.
func (c *GbkctlCommand) addSubCommands() {
	c.newInitSubCommand()
	c.newGenerateSubCommand()
	c.newVersionSubCommand()
}
