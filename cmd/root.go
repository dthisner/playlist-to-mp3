/*
Copyright © 2021 Dennis Thisner <dthisner@protonmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	constans "github.com/dthisner/m3u-to-drive/constants"
)

var cfgFile, destination, m3uLocation, origin string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "m3u-to-drive",
	Short: "Help transfering your music to your device",
	Long: `M3U to Drive is here to help you transfer your 
	music to your device by reading an M3U file that you
	have exported from examle iTunes (Music). `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.m3u-to-drive.yaml)")
	rootCmd.PersistentFlags().StringVar(&destination, "dest", "", "Where should the files be copied to")
	rootCmd.PersistentFlags().StringVar(&m3uLocation, "m3u", "", "The location for the m3u file")
	rootCmd.PersistentFlags().StringVar(&origin, "origin", "", "verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("json")
		viper.SetConfigName(".m3u-to-drive")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(constans.MissingConfig)
		} else {
			log.Fatal(constans.ReadingConfig, err)
		}
	}

	viper.BindPFlag(constans.Dest, rootCmd.Flags().Lookup("dest"))
	viper.BindPFlag(constans.Origin, rootCmd.Flags().Lookup("origin"))
	viper.BindPFlag(constans.M3uLocation, rootCmd.Flags().Lookup("m3u"))

	checkEnvVariables()
}

func checkEnvVariables() {
	if !viper.IsSet(constans.Dest) {
		log.Fatal(constans.MissingDest)
	}
	if !viper.IsSet(constans.Origin) {
		log.Fatal(constans.MissingOrigin)
	}
	if !viper.IsSet(constans.M3uLocation) {
		log.Fatal(constans.MissingM3uLoc)
	}
}
