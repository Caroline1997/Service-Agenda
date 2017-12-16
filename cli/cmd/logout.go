package cmd

import (
	"Service-Agenda/cli/service"

	"github.com/spf13/cobra"
	"os"
	"log"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout current user",
	Long:  `the usage of the command is to log out current user`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")
		service.Logout(name, password)
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)
	logoutCmd.Flags().StringP("name", "n", "", "logout username")
	logoutCmd.Flags().StringP("password", "p", "", "password")
}
