package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "practice",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
`,
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute is
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
