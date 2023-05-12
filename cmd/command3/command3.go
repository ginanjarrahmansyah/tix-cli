package command3

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdCommand3() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command3",
		Short: "Command3",
		Run:   runCommand3,
	}

	return cmd
}

func runCommand3(cmd *cobra.Command, args []string) {
	fmt.Println("Executing command3")
	// add command implementation here
}
