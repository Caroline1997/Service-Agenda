package cmd

import (
	"fmt"
	"Service-Agenda/cli/service"
	"github.com/spf13/cobra"
	"os"
	"log"
)

var deleteCmd = &cobra.Command{
	Use:   "delete user",
	Short: "for user to delete the account",
	Long: `the usage is to delete his own acccount`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")
		service.Delete_user(name, password)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("name","n","","user_name")
	deleteCmd.Flags().StringP("password", "p", "", "password")
}
