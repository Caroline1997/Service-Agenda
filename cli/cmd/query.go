package cmd

import (
	"fmt"
	"github.com/Caroline1997/Service-Agenda/cli/service"
	"github.com/spf13/cobra"
	"log"
	"os"
	//"io"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "for user to query",
	Long: `the usage is to query`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")
		service.Query_user(name,password)
	},
}

func init() {
	RootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringP("name","n","","user_name")
	queryCmd.Flags().StringP("password", "p", "", "password")
}
