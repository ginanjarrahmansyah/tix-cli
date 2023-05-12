package command1

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdCommand1() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command1",
		Short: "Command1",
		Run:   runCommand1,
	}

	return cmd
}

func runCommand1(cmd *cobra.Command, args []string) {
	fmt.Println("Executing command1")
	// add command implementation here
}
