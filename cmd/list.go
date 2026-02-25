package cmd

import (
	"fmt"
	"os"
	"strconv"

	"text/tabwriter"

	"github.com/fhsinchy/tent/runtime"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all running services",
	Long:  `The list command lists all the running services in a nicely formatted table.`,
	Run: func(cmd *cobra.Command, args []string) {
		rt := runtime.Connect()

		tentContainers := rt.ListTentContainers()

		w := new(tabwriter.Writer)

		w.Init(os.Stdout, 5, 5, 5, ' ', 0)

		defer w.Flush()

		fmt.Fprintf(w, "\n %s\t%s\t%s\t", "CONTAINER", "IMAGE", "PORTS")

		for _, tentContainer := range tentContainers {
			ports := strconv.Itoa(int(tentContainer.Ports[0].HostPort)) + "->" + strconv.Itoa(int(tentContainer.Ports[0].ContainerPort)) + "/" + tentContainer.Ports[0].Protocol

			fmt.Fprintf(w, "\n %s\t%s\t%s\t", tentContainer.Name, tentContainer.Image, ports)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
