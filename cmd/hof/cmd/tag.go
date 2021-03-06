package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"

	"github.com/hofstadter-io/hof/lib/workspace"
)

var tagLong = `create, list, delete or verify a tag object signed with GPG`

func TagRun(args []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = workspace.RunTagFromArgs(args)

	return err
}

var TagCmd = &cobra.Command{

	Use: "tag",

	Short: "create, list, delete or verify a tag object signed with GPG",

	Long: tagLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendCommandPath(cmd.CommandPath())

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = TagRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {

	help := TagCmd.HelpFunc()
	usage := TagCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
		return usage(cmd)
	}
	TagCmd.SetHelpFunc(thelp)
	TagCmd.SetUsageFunc(tusage)

}