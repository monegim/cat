package cat

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// type
func Cat(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("no file provided")
		os.Exit(1)
	}
	if len(args) == 1 {
		content, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer content.Close()
		scanner := bufio.NewScanner(content)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
