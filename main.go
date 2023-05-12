package main

import (
	"fmt"
	"os"

	"github.com/ginanjarrahmansyah/tix-cli/cmd/command1"
	"github.com/ginanjarrahmansyah/tix-cli/cmd/command2"
	"github.com/ginanjarrahmansyah/tix-cli/cmd/command3"
	"github.com/ginanjarrahmansyah/tix-cli/cmd/gcpls"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tix-cli",
	Short: "A collection of TIX CLI tools",
}

func init() {
	rootCmd.AddCommand(gcpls.NewCmdGCPLS())
	rootCmd.AddCommand(command1.NewCmdCommand1())
	rootCmd.AddCommand(command2.NewCmdCommand2())
	rootCmd.AddCommand(command3.NewCmdCommand3())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
