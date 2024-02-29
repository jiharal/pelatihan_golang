package cmd

import (
	"pelatihan_golang/services/auth"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:  "auth",
	Long: "Service for authentications",
	Run: func(cmd *cobra.Command, args []string) {
		auth.StartServerOauth2()
	},
}
