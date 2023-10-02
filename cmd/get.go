package cmd

import (
	"fmt"

	db "github.com/jcardarelli/fancy-cli/database"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [restaurantName]",
	Short: "Get restaurant info from the database",
	Long:  `Get restaurant info.`,
	Run: func(cmd *cobra.Command, args []string) {
		var restaurantName string

		if len(args) > 0 {
			restaurantName = args[0]
		}

		allFlag, _ := cmd.Flags().GetBool("all")
		if !allFlag {
			fmt.Println(db.GetRestaurant(restaurantName))
		} else {
			db.GetAllRestaurants()
		}
	},
}

func init() {
	restaurantCmd.AddCommand(getCmd)

	// Cobra supports Persistent Flags which will work for this command
	getCmd.PersistentFlags().Bool("all", false, "Get all restaurants")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
