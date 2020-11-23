package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/Stavrospanakakis/lfl/services"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lfl",
	Short: "lfl opens the link of the lecture of your choice in zoom",
	Long: `lfl (Links for Lectures) opens the link of the lecture of your choice
in zoom. You add the subjects' info in your config file and you are done.
Current configuration file only supports the lectures of the 3rd Year of Department
of Informatics at Ionian University.`,
	Run: func(cmd *cobra.Command, args []string) {
		lectureID, _ := cmd.Flags().GetString("lecture")
		err := services.GetLinkByID(lectureID)
		if err != nil {
			fmt.Println(err)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// It also shows a list of the lectures
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lFlag, _ := rootCmd.Flags().GetString("lecture")
	if lFlag == "" {
		err := services.ShowLecturesList()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	var lecture string

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lfl.json)")
	rootCmd.Flags().StringVarP(&lecture, "lecture", "l", "", "The id of the lecture of your choice (Type lfl -i to see the available ids)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".lfl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".lfl.json")
	}

	viper.AutomaticEnv() // read in environment variables that match
}
