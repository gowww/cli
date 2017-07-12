package cli

import (
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
	var description string
	if command == nil {
		if Description == "" {
			description = strings.Title(filepath.Base(os.Args[0]))
		} else {
			description = Description
		}
	} else {
		if command.description == "" {
			description = strings.Title(command.name)
		} else {
			description = command.description
		}
	}

	fmt.Printf("\n%s\n\nUsage:\n\n\t%s", description, os.Args[0])
	if len(mainFlagsUsage) > 0 {
		fmt.Print(" [flags]")
	}
	if command == nil {
		if len(mainCommands) > 0 {
			fmt.Printf(" [command]")
		}
	} else {
		fmt.Printf(" %s", command.name)
		if len(command.flagsUsage) > 0 {
			fmt.Print(" [flags]")
		}
	}
	fmt.Print("\n\n")

	if command == nil && len(mainCommands) > 0 {
		fmt.Print("Commands:\n\n")
		sort.Slice(mainCommands, func(i, j int) bool { return mainCommands[i].name < mainCommands[j].name })
		l := maxCommandLen(mainCommands)
		for _, cmd := range mainCommands {
			fmt.Printf("\t%s%s  %s\n", cmd.name, strings.Repeat(" ", l-len(cmd.name)), cmd.description)
		}
		fmt.Print("\n")
	}

	var flags map[string]string
	if command == nil {
		flags = mainFlagsUsage
	} else {
		flags = command.flagsUsage
	}
	if len(flags) > 0 {
		fmt.Print("Flags:\n\n")
		ff := sortedFlags(flags)
		l := maxFlagLen(flags)
		for _, f := range ff {
			fmt.Printf("\t-%s%s  %s\n", f, strings.Repeat(" ", l-len(f)), flags[f])
		}
		fmt.Print("\n")
	}
}

func maxCommandLen(cmds []*CommandUnit) (l int) {
	for _, c := range cmds {
		if len(c.name) > l {
			l = len(c.name)
		}
	}
	return
}

func maxFlagLen(flags map[string]string) (l int) {
	for f := range flags {
		if len(f) > l {
			l = len(f)
		}
	}
	return
}

func sortedFlags(ss map[string]string) []string {
	keys := make([]string, 0, len(ss))
	for k := range ss {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
