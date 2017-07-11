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
	printUsage(Description, "", commands, flagsUsage)
}

func printUsage(description, command string, commands []*CommandUnit, flags map[string]string) {
	if description == "" {
		if command == "" {
			description = strings.Title(filepath.Base(os.Args[0]))
		} else {
			description = strings.Title(command)
		}
	}
	fmt.Printf("\n%s\n\nUsage:\n\n\t%s", description, os.Args[0])
	if command != "" {
		fmt.Printf(" %s", command)
	}
	if len(commands) > 0 {
		fmt.Printf(" [command]")
	}
	if len(flags) > 0 {
		fmt.Printf(" [flags]")
	}
	fmt.Print("\n\n")
	if len(commands) > 0 {
		fmt.Print("Commands:\n\n")
		sort.Slice(commands, func(i, j int) bool { return commands[i].name < commands[j].name })
		l := maxCommandLen(commands)
		for _, cmd := range commands {
			fmt.Printf("\t%s%s  %s\n", cmd.name, strings.Repeat(" ", l-len(cmd.name)), cmd.description)
		}
		fmt.Print("\n")
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
