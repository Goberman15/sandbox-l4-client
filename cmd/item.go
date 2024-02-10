package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(itemCmd)
}

var itemCmd = &cobra.Command{
	Use: "item",
	Short: "Item Crud",
	Long: "Crud Operation for Item Model",
}