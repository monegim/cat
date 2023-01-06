package cat

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "cat",
	Short: "cat - cat a file",
	Long: "It acts like linux cat",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}
