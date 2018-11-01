package cmd

import (
	"fmt"
	"strconv"

	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
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
