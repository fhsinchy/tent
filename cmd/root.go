package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tent",
	Short: "Podman (https://podman.io/) based development-only dependency manager for Linux.",
	Long: `Setting up different development dependencies such as MySQL, MongoDB or Redis has always been a pain for developers.
Tent is a tool designed for making that process easier. It allows you to set-up several development dependencies by executing simple commands.
i.e. the following command will start a functional MySQL server on your local system.

$ tent start mysql --default

Tent leverages the power of containerization for achieving it's goals. All the available services (as they are called inside tent) are just pre-configured containers.
These containers are created from OCI compliant images and use Podman as their container runtime.

Tent is heavily inspired from tighten/takeout (https://github.com/tighten/takeout) and is an experimental project. Hence care should be taken if you're using it in a critical environment.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tent.yaml)")
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

		// Search config in home directory with name ".tent" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tent")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
