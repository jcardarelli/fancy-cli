package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	// Use the name of the binary generated from `go build -o name` as the
	// root cmd that cobra shows in the help menu
	Use:   filepath.Base(os.Args[0]),
	Short: "Restaurant database manager",
	Long:  `Restaurant database`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Print startup banner with go-figure
	fmt.Println()
	banner := figure.NewFigure("fancy", "ticks", true)
	banner.Print()
	fmt.Println()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fancy-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
