package cat

import (
	"fmt"
	"os"

	"github.com/monegim/cat/pkg/cat"
	"github.com/spf13/cobra"
)
var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "cat",
	Short: "cat - cat a file",
	Long:  "It acts like linux cat",
	Args: cobra.ArbitraryArgs,
	Version: version,
	Run: cat.Cat,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
