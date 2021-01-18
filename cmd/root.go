package cmd

import (
	"fmt"
	"os"

	"github.com/Stavrospanakakis/lfl/internal/services"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lfl",
	Short: "lfl opens the link of the lecture of your choice in zoom",
	Long: `lfl (Links for Lectures) opens the link of the lecture of your choice
in zoom. You add the subjects' info in your config file and you are done.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := services.MakeService()
		configPath, err := s.GetConfigPath()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if s.ConfigurationFileExists(configPath) {
			s.Run(configPath)
		} else {
			err := s.GenerateConfig(configPath)
			if err != nil {
				fmt.Println(err)
			}
			os.Exit(1)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
