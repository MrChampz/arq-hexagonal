package cmd

import (
	"fmt"

	"github.com/MrChampz/arq-hexagonal/adapters/web/server"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewWebServer(productService)
		fmt.Println("WebServer has been started")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
