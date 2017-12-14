package cmd

import (
	"Service-Agenda/cli/service"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new account",
	Long:  `usage of using this command is to register a new account`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		mail, _ := cmd.Flags().GetString("mail")
		phone, _ := cmd.Flags().GetString("phone")
		service.Register(username, password, mail, phone)
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "", "Username")
	registerCmd.Flags().StringP("password", "p", "", "Password")
	registerCmd.Flags().StringP("mail", "m", "", "email_address")
	registerCmd.Flags().StringP("phone", "n", "", "phone_number")
}
