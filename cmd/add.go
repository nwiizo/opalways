package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Calculator of addition.",
	Long:  "Calculator to perform the addition.",
	Run: func(cmd *cobra.Command, args []string) {
		var n1 int
		var n2 int
		n1, _ = strconv.Atoi(args[0])
		n2, _ = strconv.Atoi(args[1])
		fmt.Println(n1 + n2)
	},
}
