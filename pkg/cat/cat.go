package cat

import (
	"fmt"

	"github.com/spf13/cobra"
)
// type 
func Cat (cmd *cobra.Command, args []string) {
	for i, arg := range args {
		fmt.Printf("%d, %s\n", i, arg)
	}
}