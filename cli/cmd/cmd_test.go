package cmd

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	fmt.Println("Register")
	registerCmd.Flags().Set("user", "root")
	registerCmd.Flags().Set("password", "pass")
	registerCmd.Flags().Set("mail", "mail")
	registerCmd.Flags().Set("phone", "phone")
	registerCmd.Run(registerCmd, nil)
}

func TestLogin(t *testing.T) {
	fmt.Println("Login")
	loginCmd.Flags().Set("name", "root")
	loginCmd.Flags().Set("password", "pass")
	loginCmd.Run(loginCmd, nil)
}
