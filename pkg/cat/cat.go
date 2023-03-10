package cat

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// type
func Cat(cmd *cobra.Command, args []string) {
	n, _ := cmd.Flags().GetBool("name")
	if len(args) == 0 {
		fmt.Println("no file provided")
		os.Exit(1)
	} else {
		f := &file{}
		if n {
			f.n = 1
		} else {
			f.n = 0
		}
		for i := 0; i < len(args); i++ {
			f.name = args[i]
			f.readFile()
			// f.n = last_line
		}
	}
}

type file struct {
	name string
	n    int
}

func (f *file) readFile() {
	content, err := os.Open(f.name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		if f.n == 0 {
			fmt.Println(scanner.Text())
		} else {
			fmt.Println(f.n, " ", scanner.Text())
			f.n++

		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
