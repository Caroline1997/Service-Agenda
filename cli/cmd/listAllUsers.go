package cmd

import (
	"fmt"

	"github.com/Caroline1997/Service-Agenda/cli/service"
	"github.com/spf13/cobra"
	//"log"
	//"os"
)

// listAllUsersCmd represents the listAllUsers command
var listAllUsersCmd = &cobra.Command{
	Use:   "listAllUsers",
	Short: "To list all users",
	Long:  `To list all users`,
	Run: func(cmd *cobra.Command, args []string) {
		username, err := service.List_all_users()
		if err != nil {
			fmt.Println(err)
		}
		if len(username) == 0 {
			fmt.Println("There is no user!")
			return
		}
		fmt.Println("All users in this Agenda")
		for _, user := range username {
			fmt.Println("Username:", user.Username)
		}
	},
}

func init() {
	RootCmd.AddCommand(listAllUsersCmd)
}
