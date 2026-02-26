package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
	rootCmd.Version = version
}
