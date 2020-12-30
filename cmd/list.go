package cmd

import (
	"fmt"
	"os"
	"strconv"

	"text/tabwriter"

	"github.com/fhsinchy/tent/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all running services",
	Long:  `The list command lists all the running services in a nicely formatted table.`,
	Run: func(cmd *cobra.Command, args []string) {
		connText := utils.GetContext()

		tentContainers := utils.ListTentContainers(connText)

		w := new(tabwriter.Writer)

		w.Init(os.Stdout, 5, 5, 5, ' ', 0)

		defer w.Flush()

		fmt.Fprintf(w, "\n %s\t%s\t%s\t", "CONTAINER", "IMAGE", "PORTS")

		for _, tentContainer := range tentContainers {
			ports := strconv.Itoa(int(tentContainer.Ports[0].HostPort)) + "->" + strconv.Itoa(int(tentContainer.Ports[0].ContainerPort)) + "/" + tentContainer.Ports[0].Protocol

			fmt.Fprintf(w, "\n %s\t%s\t%s\t", tentContainer.Names[0], tentContainer.Image, ports)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
