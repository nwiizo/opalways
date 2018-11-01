package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "github.com/spf13/viper"
	"os/exec"
)

func init() {
	RootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Calculator of addition.",
	Long:  "Calculator to perform the addition.",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("ls", "-la").Output()
	},
}
