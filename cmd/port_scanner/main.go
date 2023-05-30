package main

import (
	"fmt"

	"github.com/spf13/cobra"
	portscanner "github.com/yhorbachov/black-hat-go/pkg/port_scanner"
)

var from int
var to int

var rootCmd = &cobra.Command{
	Use:   "portscan example.com",
	Short: "portscan is a very basic port scanner",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]

		openports := portscanner.Scan(host, from, to)
		for _, port := range openports {
			fmt.Printf("%d open\n", port)
		}
	},
}

func init() {
	rootCmd.Flags().IntVarP(&from, "from", "f", 1, "beggining of the scan range")
	rootCmd.Flags().IntVarP(&to, "to", "t", 1024, "end of the scan range")
}

func main() {
	rootCmd.Execute()
}
