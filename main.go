package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(rootCmd,
		cli.Tree(helpCmd),
		cli.Tree(versionCmd),
		cli.Tree(runCmd),
		cli.Tree(buildCmd),
		cli.Tree(transpileCmd),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
