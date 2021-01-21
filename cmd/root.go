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
	Short: "lfl opens the lectures virtual meeting of your choice with just one command.",
	Long: `lfl (Links for Lectures) is a CLI which opens the lectures virtual meeting of your choice
from terminal. It currently supports only Zoom & Webex links and it is tested only in Linux environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := services.MakeService()
		err := s.SetConfigPath()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		wantToAddLectures, _ := cmd.Flags().GetBool("add")
		wantToRemoveLectures, _ := cmd.Flags().GetBool("remove")

		if wantToAddLectures {
			s.AddNewLectures()
		} else if wantToRemoveLectures {
			s.RemoveLectures()
		} else {
			if s.ConfigurationFileExists() || s.Lectures != nil {
				err := s.Run()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else {
				err := s.GenerateConfig()
				if err != nil {
					fmt.Println(err)
				}
				os.Exit(1)
			}
		}

	},
}

func init() {
	rootCmd.PersistentFlags().Bool("add", false, "add new lectures")
	rootCmd.PersistentFlags().Bool("remove", false, "remove lecture")
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
