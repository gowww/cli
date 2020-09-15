package cli

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	usageText   string
	usageSuffix string
	commands    []*CommandUnit
	args        []string
	subArgs     []string
)

// SetUsageText sets the CLI description for the usage help.
func SetUsageText(s string) {
	usageText = s
}

// SetUsageSuffix sets the suffix added to the the usage line.
// It can be used to documentate command arguments.
func SetUsageSuffix(s string) {
	usageSuffix = s
}

// Parse parses the command arguments.
func Parse() {
	args = os.Args[1:]

	// Parse subarguments (after "--") for a subprocess.
	for i := 0; i < len(args); i++ {
		if args[i] == "--" {
			subArgs = args[i+1:]
			args = args[:i]
		}
	}

	// Check if first argument is a command and parse its flags.
	if len(args) >= 1 {
		for _, c := range commands {
			if c.flagSet.Name() != args[0] {
				continue
			}
			c.flagSet.Parse(args[1:])
			args = c.flagSet.Args()
			c.f()
			os.Exit(0)
		}
	}

	// Otherwise, parse main flags.
	flag.Usage = Usage
	flag.CommandLine.Parse(args)
	args = flag.Args()
}

// Parsed reports whether the command-line flags have been parsed.
func Parsed() bool {
	return flag.Parsed()
}

// Arg returns the i'th non-flag CLI argument.
func Arg(i int) string {
	if i < 0 || i >= len(args) {
		return ""
	}
	return args[i]
}

// Args returns the non-flag CLI arguments.
func Args() []string {
	return args
}

// SubArgs returns the arguments after "--".
func SubArgs() []string {
	return subArgs
}

// CleanLines removes n lines from terminal.
func CleanLines(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("\033[1A\033[0K")
	}
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func Bool(p *bool, name string, value bool, usage string) {
	flag.BoolVar(p, name, value, usage)
	for _, c := range commands {
		c.Bool(p, name, value, usage)
	}
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func Duration(p *time.Duration, name string, value time.Duration, usage string) {
	flag.DurationVar(p, name, value, usage)
	for _, c := range commands {
		c.Duration(p, name, value, usage)
	}
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func Float64(p *float64, name string, value float64, usage string) {
	flag.Float64Var(p, name, value, usage)
	for _, c := range commands {
		c.Float64(p, name, value, usage)
	}
}

// Int defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func Int(p *int, name string, value int, usage string) {
	flag.IntVar(p, name, value, usage)
	for _, c := range commands {
		c.Int(p, name, value, usage)
	}
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func Int64(p *int64, name string, value int64, usage string) {
	flag.Int64Var(p, name, value, usage)
	for _, c := range commands {
		c.Int64(p, name, value, usage)
	}
}

// String defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func String(p *string, name string, value string, usage string) {
	flag.StringVar(p, name, value, usage)
	for _, c := range commands {
		c.String(p, name, value, usage)
	}
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func Uint(p *uint, name string, value uint, usage string) {
	flag.UintVar(p, name, value, usage)
	for _, c := range commands {
		c.Uint(p, name, value, usage)
	}
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func Uint64(p *uint64, name string, value uint64, usage string) {
	flag.Uint64Var(p, name, value, usage)
	for _, c := range commands {
		c.Uint64(p, name, value, usage)
	}
}
