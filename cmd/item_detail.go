package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(itemDetailCmd)
}

var itemDetailCmd = &cobra.Command{
	Use: "detail",
	Short: "Item Detail Crud",
	Long: "Crud Operation for Item Detail Model",
}