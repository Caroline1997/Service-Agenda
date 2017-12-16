package cmd

import (
	"fmt"
	"github.com/Caroline1997/Service-Agenda/cli/service"
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
		err := service.Delete_user(name, password)
		if err != nil {
			  fmt.Println(err)
		}
		fmt.Println("Delete user successfully!")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("name","n","","user_name")
	deleteCmd.Flags().StringP("password", "p", "", "password")
}
