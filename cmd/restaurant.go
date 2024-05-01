/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// restaurantCmd represents the restaurant command
var restaurantCmd = &cobra.Command{
	Use:   "restaurant",
	Short: "List and modify restaurant data",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("fancy cli started")
	},
}

func init() {
	rootCmd.AddCommand(restaurantCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restaurantCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restaurantCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
