package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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
		rt, err := runtime.Connect()
		if err != nil {
			log.Fatalln(err)
		}

		if err := listContainers(rt, os.Stdout); err != nil {
			log.Fatalln(err)
		}
	},
}

func listContainers(engine runtime.ContainerEngine, out io.Writer) error {
	tentContainers, err := engine.ListTentContainers()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	w.Init(out, 5, 5, 5, ' ', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t", "CONTAINER", "IMAGE", "PORTS")

	for _, tentContainer := range tentContainers {
		var portParts []string
		for _, p := range tentContainer.Ports {
			portParts = append(portParts, strconv.Itoa(int(p.HostPort))+"->"+strconv.Itoa(int(p.ContainerPort))+"/"+p.Protocol)
		}
		ports := strings.Join(portParts, ", ")

		fmt.Fprintf(w, "\n %s\t%s\t%s\t", tentContainer.Name, tentContainer.Image, ports)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
