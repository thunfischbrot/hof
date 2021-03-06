package cmddatamodel

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"

	"github.com/hofstadter-io/hof/lib/datamodel"
)

var visualizeLong = `visualize a data model`

func VisualizeRun(args []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = datamodel.RunVisualizeFromArgs(args)

	return err
}

var VisualizeCmd = &cobra.Command{

	Use: "visualize",

	Aliases: []string{
		"v",
		"viz",
		"show",
	},

	Short: "visualize a data model",

	Long: visualizeLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendCommandPath(cmd.CommandPath())

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = VisualizeRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {

	help := VisualizeCmd.HelpFunc()
	usage := VisualizeCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
		return usage(cmd)
	}
	VisualizeCmd.SetHelpFunc(thelp)
	VisualizeCmd.SetUsageFunc(tusage)

}