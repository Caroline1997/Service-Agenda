package cmd

import (
	"fmt"
	"github.com/Caroline1997/Service-Agenda/cli/service"
	"github.com/spf13/cobra"
	"os"
	"log"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "for user to log in",
	Long: `the usage is to log in`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")
		err := service.Login(name, password)
		if err != nil {
			  fmt.Println(err)
		}
		fmt.Println("Login successfully!")
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("name","n","","user_name")
	loginCmd.Flags().StringP("password", "p", "", "password")
}
