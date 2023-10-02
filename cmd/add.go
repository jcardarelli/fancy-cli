package cmd

import (
	"log"

	db "github.com/jcardarelli/fancy-cli/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [restaurant-name]",
	Short: "Add a new restaurant to the database",
	Long:  `Add new restaurant so that it can be queried with the command 'fancy restaurants get [restaurant_name]'.`,
	Run: func(cmd *cobra.Command, args []string) {
		// insert new restaurant
		restaurant_name := args[0]
		address := "San Francisco, CA"
		michelin_stars := 2
		_, err := db.InsertRestaurant(restaurant_name, address, michelin_stars)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	restaurantCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
