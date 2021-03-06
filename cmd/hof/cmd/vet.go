package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"

	"github.com/hofstadter-io/hof/lib/cuetils"
)

var vetLong = `validate data`

func VetRun(args []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = cuetils.RunVetFromArgs(args)

	return err
}

var VetCmd = &cobra.Command{

	Use: "vet",

	Short: "validate data",

	Long: vetLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendCommandPath(cmd.CommandPath())

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = VetRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {

	help := VetCmd.HelpFunc()
	usage := VetCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
		return usage(cmd)
	}
	VetCmd.SetHelpFunc(thelp)
	VetCmd.SetUsageFunc(tusage)

}