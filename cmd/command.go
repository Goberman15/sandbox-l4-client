package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const baseUrl = "http://localhost:8080/api"


var rootCmd = &cobra.Command{
	Use: "crud",
	Short: "Simple CRUD for Item and Item Details",
	Long: "Simple CRUD for Item and Item Details as Part of Sandbox L4 Task",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
