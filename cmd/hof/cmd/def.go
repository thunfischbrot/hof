package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"

	"github.com/hofstadter-io/hof/lib/cuetils"
)

var defLong = `print consolidated definitions`

func DefRun(args []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = cuetils.RunDefFromArgs(args)

	return err
}

var DefCmd = &cobra.Command{

	Use: "def",

	Short: "print consolidated definitions",

	Long: defLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "", 0)

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = DefRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {

	help := DefCmd.HelpFunc()
	usage := DefCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/help", "", 0)
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/usage", "", 0)
		return usage(cmd)
	}
	DefCmd.SetHelpFunc(thelp)
	DefCmd.SetUsageFunc(tusage)

}