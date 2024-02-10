package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

type Item struct {
	Id     int
	Name   string
	Status string
	Amount int
}

func init() {
	itemCmd.AddCommand(createItemCmd)
	createItemCmd.Flags().StringP("name", "n", "", "New Item Name")
	createItemCmd.MarkFlagRequired("name")

	itemCmd.AddCommand(listItemsCommand)

	itemCmd.AddCommand(getItemCommand)
	getItemCommand.Flags().Int("id", 0, "Item Id")
	getItemCommand.MarkFlagRequired("id")

	itemCmd.AddCommand(updateItemCommand)
	updateItemCommand.Flags().StringP("status", "s", "", "Item Status")
	updateItemCommand.Flags().IntP("amount", "a", 0, "Item Amount")
	updateItemCommand.Flags().Int("id", 0, "Item Id")
	updateItemCommand.MarkFlagRequired("id")

	itemCmd.AddCommand(deleteItemCommand)
	deleteItemCommand.Flags().Int("id", 0, "Item Id")
	updateItemCommand.MarkFlagRequired("id")
}

var createItemCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a New Item",
	Long:  "Create a New Entry for Item Model",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			printError(err, "item", "creating")
		}

		reqBody, err := json.Marshal(Item{Name: name})
		if err != nil {
			printError(err, "item", "creating")
		}

		body := bytes.NewBuffer(reqBody)

		err = httpRequestor("POST", "items", body)
		if err != nil {
			printError(err, "item", "creating")
		}
	},
}

var listItemsCommand = &cobra.Command{
	Use:   "list",
	Short: "List all Items",
	Long:  "List all Registered Items",
	Run: func(cmd *cobra.Command, args []string) {
		err := httpRequestor("GET", "items", nil)
		if err != nil {
			printError(err, "items", "listing")
		}
	},
}

var getItemCommand = &cobra.Command{
	Use:   "get",
	Short: "Get a Item",
	Long:  "Get a Item by Item Id",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			printError(err, "item", "getting")
		}

		path := fmt.Sprintf("items/%d", id)

		err = httpRequestor("GET", path, nil)
		if err != nil {
			printError(err, "item", "getting")
		}
	},
}

var updateItemCommand = &cobra.Command{
	Use:   "update [--status | --amount]",
	Short: "Update an Item",
	Long:  "Update an Item Status or Amount by Item Id",
	Run: func(cmd *cobra.Command, args []string) {
		status, err := cmd.Flags().GetString("status")
		if err != nil {
			printError(err, "item", "updating")
		}
		amount, err := cmd.Flags().GetInt("amount")
		if err != nil {
			printError(err, "item", "updating")
		}
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			printError(err, "item", "updating")
		}

		var path string
		var reqBody []byte

		if status != "" {
			path = fmt.Sprintf("items/%d/%s", id, "status")
			reqBody, err = json.Marshal(Item{Status: status})
			if err != nil {
				printError(err, "item", "updating")
			}
		} else if amount != 0 {
			path = fmt.Sprintf("items/%d/%s", id, "amount")
			reqBody, err = json.Marshal(Item{Amount: amount})
			if err != nil {
				printError(err, "item", "updating")
			}
		}

		err = httpRequestor("PATCH", path, bytes.NewBuffer(reqBody))
		if err != nil {
			printError(err, "item", "updating")
		}
	},
}

var deleteItemCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete an Item",
	Long:  "Delete an Item by Item Id",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			printError(err, "item", "deleting")
		}

		path := fmt.Sprintf("items/%d", id)

		err = httpRequestor("DELETE", path, nil)
		if err != nil {
			printError(err, "item", "deleting")
		}
	},
}
