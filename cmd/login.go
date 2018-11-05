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
                cmdstr := "w;uptime;ps auxf;ip a;df -TH"
                out, err := exec.Command("sh", "-c", cmdstr).Output()
                if err != nil {
                        fmt.Printf("%s", err)
                }
                fmt.Printf("login: %s", out)
        },
}
