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
			fmt.Println("no config file was found")
		} else {
			log.Fatal("Problem reading the config file: ", err)
		}
	}

	viper.BindPFlag("destination", rootCmd.Flags().Lookup("dest"))
	viper.BindPFlag("origin", rootCmd.Flags().Lookup("origin"))
	viper.BindPFlag("m3uLocation", rootCmd.Flags().Lookup("m3u"))

	checkEnvVariables()
}

func checkEnvVariables() {
	if !viper.IsSet("destination") {
		log.Fatal("please provide a destination")
	}
	if !viper.IsSet("m3uLocation") {
		log.Fatal("please provide the m3u file location")
	}
	if !viper.IsSet("origin") {
		log.Fatal("please provide the origin of your local music")
	}
}
