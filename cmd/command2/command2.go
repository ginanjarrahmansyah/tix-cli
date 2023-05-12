package command2

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdCommand2() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command2",
		Short: "Command2",
		Run:   runCommand2,
	}

	return cmd
}

func runCommand2(cmd *cobra.Command, args []string) {
	fmt.Println("Executing command2")
	// add command implementation here
}
