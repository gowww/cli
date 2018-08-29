package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Usage prints the CLI usage.
func Usage() {
	printUsage(nil)
}

func printUsage(command *CommandUnit) {
	// Print description
	var description string
	if command == nil {
		if Description == "" {
			description = strings.Title(filepath.Base(os.Args[0]))
		} else {
			description = Description
		}
	} else {
		if command.description == "" {
			description = strings.Title(command.flagSet.Name())
		} else {
			description = command.description
		}
	}

	// Get flags
	var flags []*Flag
	flagVisitFunc := func(f *flag.Flag) {
		flags = append(flags, &Flag{*f})
	}
	if command == nil {
		flag.VisitAll(flagVisitFunc)
	} else {
		command.flagSet.VisitAll(flagVisitFunc)
	}

	// Print usage
	fmt.Printf("\n%s\n\nUsage:\n\n\t%s", description, os.Args[0])
	if command == nil {
		if len(commands) > 0 {
			fmt.Print(" [command]")
		}
		if len(flags) > 0 {
			fmt.Print(" [flags]")
		}
	} else {
		fmt.Printf(" %s", command.flagSet.Name())
		if len(flags) > 0 {
			fmt.Print(" [flags]")
		}
	}
	fmt.Print("\n\n")

	// Print commands
	if command == nil && len(commands) > 0 {
		fmt.Print("Commands:\n\n")
		sort.Slice(commands, func(i, j int) bool { return commands[i].flagSet.Name() < commands[j].flagSet.Name() })
		l := maxCommandLen(commands)
		for _, cmd := range commands {
			fmt.Printf("\t%s%s  %s\n", cmd.flagSet.Name(), strings.Repeat(" ", l-len(cmd.flagSet.Name())), cmd.description)
		}
		fmt.Print("\n")
	}

	// Print flags
	if len(flags) > 0 {
		fmt.Print("Flags:\n\n")
		l := maxFlagLen(flags)
		for _, f := range flags {
			fmt.Printf("\t-%s%s  %s\n", f.NameWithDefValue(), strings.Repeat(" ", l-len(f.NameWithDefValue())), f.Usage)
		}
		fmt.Print("\n")
	}
}

func maxCommandLen(cmds []*CommandUnit) (l int) {
	for _, c := range cmds {
		if len(c.flagSet.Name()) > l {
			l = len(c.flagSet.Name())
		}
	}
	return
}

func maxFlagLen(flags []*Flag) (l int) {
	for _, f := range flags {
		name := f.NameWithDefValue()
		if len(name) > l {
			l = len(name)
		}
	}
	return
}
