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
	} else{
	for _, path := range(args) {
		readFile(path)
	}
}
}

func readFile(name string)  {
	content, err := os.Open(name)
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
