package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/cmd/label"

	"github.com/hofstadter-io/hof/cmd/hof/ga"
)

var labelLong = `manage labels for resources and more`

var LabelCmd = &cobra.Command{

	Use: "label",

	Aliases: []string{
		"l",
		"labels",
		"attrs",
	},

	Short: "manage labels for resources and more",

	Long: labelLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendCommandPath(cmd.CommandPath())

	},
}

func init() {

	help := LabelCmd.HelpFunc()
	usage := LabelCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
		return usage(cmd)
	}
	LabelCmd.SetHelpFunc(thelp)
	LabelCmd.SetUsageFunc(tusage)

	LabelCmd.AddCommand(cmdlabel.InfoCmd)
	LabelCmd.AddCommand(cmdlabel.CreateCmd)
	LabelCmd.AddCommand(cmdlabel.GetCmd)
	LabelCmd.AddCommand(cmdlabel.SetCmd)
	LabelCmd.AddCommand(cmdlabel.EditCmd)
	LabelCmd.AddCommand(cmdlabel.DeleteCmd)
	LabelCmd.AddCommand(cmdlabel.ApplyCmd)
	LabelCmd.AddCommand(cmdlabel.RemoveCmd)

}
