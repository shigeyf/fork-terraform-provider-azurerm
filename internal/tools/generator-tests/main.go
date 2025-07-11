// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"log"
	"os"

	"github.com/hashicorp/terraform-provider-azurerm/internal/tools/generator-tests/generators"
	"github.com/mitchellh/cli"
)

func main() {
	os.Exit(realMain(os.Args[1:]))
}

func realMain(args []string) int {
	var ui cli.Ui = &cli.ColoredUi{
		ErrorColor: cli.UiColorRed,
		WarnColor:  cli.UiColorYellow,
		InfoColor:  cli.UiColorNone,

		Ui: &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
	}

	commands := map[string]cli.CommandFactory{
		"resourceidentity": func() (cli.Command, error) {
			return &generators.ResourceIdentityCommand{
				Ui: ui,
			}, nil
		},
	}

	gen := cli.CLI{
		Args:     args,
		Commands: commands,
		Name:     "testgenerator",
		Version:  "0.1",
	}

	exitStatus, err := gen.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}
