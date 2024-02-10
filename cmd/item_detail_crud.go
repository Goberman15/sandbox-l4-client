package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

type ItemDetail struct {
	Id     int
	ItemId int
	Name   string
}

func init() {
	itemDetailCmd.AddCommand(createItemDetailCmd)
	createItemDetailCmd.Flags().StringP("name", "n", "", "Item Detail Name")
	createItemDetailCmd.Flags().IntP("item-id", "i", 0, "Item Id")
	createItemDetailCmd.MarkFlagRequired("name")
	createItemDetailCmd.MarkFlagRequired("item-id")

	itemDetailCmd.AddCommand(listItemDetailsCommand)
	listItemDetailsCommand.Flags().IntP("item-id", "i", 0, "Item Id")
	listItemDetailsCommand.MarkFlagRequired("item-id")

	itemDetailCmd.AddCommand(updateItemDetailCmd)
	updateItemDetailCmd.Flags().Int("id", 0, "Item Detail Id")
	updateItemDetailCmd.Flags().StringP("name", "n", "", "Item Detail Name")
	updateItemDetailCmd.MarkFlagRequired("id")
	updateItemDetailCmd.MarkFlagRequired("name")

	itemDetailCmd.AddCommand(deleteItemDetailCmd)
	deleteItemDetailCmd.Flags().Int("id", 0, "Item Detail Id")
	deleteItemDetailCmd.MarkFlagRequired("id")
}

var createItemDetailCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Item Detail",
	Long:  "Create a New Entry for Item Detail Model",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			printError(err, "item detail", "creating")
		}
		itemId, err := cmd.Flags().GetInt("item-id")
		if err != nil {
			printError(err, "item detail", "creating")
		}

		reqBody, err := json.Marshal(ItemDetail{Name: name, ItemId: itemId})
		if err != nil {
			printError(err, "item detail", "creating")
		}

		body := bytes.NewBuffer(reqBody)

		err = httpRequestor("POST", "item-details", body)
		if err != nil {
			printError(err, "item detail", "creating")
		}
	},
}

var listItemDetailsCommand = &cobra.Command{
	Use:   "list",
	Short: "List Item Details",
	Long:  "List All Item Details by Item Id",
	Run: func(cmd *cobra.Command, args []string) {
		itemId, err := cmd.Flags().GetInt("item-id")
		if err != nil {
			printError(err, "item details", "listing")
		}

		path := fmt.Sprintf("item-details/%d", itemId)

		err = httpRequestor("GET", path, nil)
		if err != nil {
			printError(err, "item details", "listing")
		}
	},
}

var updateItemDetailCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Item Detail",
	Long:  "Update Item Detail by Id",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			printError(err, "item detail", "updating")
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			printError(err, "item detail", "updating")
		}

		reqBody, err := json.Marshal(ItemDetail{Name: name})
		if err != nil {
			printError(err, "item detail", "updating")
		}

		path := fmt.Sprintf("item-details/%d", id)

		body := bytes.NewBuffer(reqBody)

		err = httpRequestor("PUT", path, body)
		if err != nil {
			printError(err, "item detail", "updating")
		}
	},
}

var deleteItemDetailCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Item Detail",
	Long:  "Delete Item Detail by Id",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			printError(err, "item detail", "deleting")
		}

		path := fmt.Sprintf("item-details/%d", id)

		err = httpRequestor("DELETE", path, nil)
		if err != nil {
			printError(err, "item detail", "deleting")
		}
	},
}
