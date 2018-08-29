# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) cli [![GoDoc](https://godoc.org/github.com/gowww/cli?status.svg)](https://godoc.org/github.com/gowww/cli) [![Build](https://travis-ci.org/gowww/cli.svg?branch=master)](https://travis-ci.org/gowww/cli) [![Coverage](https://coveralls.io/repos/github/gowww/cli/badge.svg?branch=master)](https://coveralls.io/github/gowww/cli?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/cli)](https://goreportcard.com/report/github.com/gowww/cli) ![Status Testing](https://img.shields.io/badge/status-testing-orange.svg)

Package [cli](https://godoc.org/github.com/gowww/cli) wraps the standard [flag](https://golang.org/pkg/flag/) package for a cleaner command line interface.

## Installing

1. Get package:

	```Shell
	go get -u github.com/gowww/cli
	```

2. Import it in your code:

	```Go
	import "github.com/gowww/cli"
	```

## Usage

Henceforth, by "command" we mean "subcommand" (like the `build` part in `go build`)…

The order in which you define commands and flags is important!  
When you define a main flag, it will be added to the top-level flag set but also to all commands already defined.

Obviously, each command can also define its own flags.

For the sake of clarity for the developer and ease of use for the final user, the usage pattern is simple and always the same : `program [command] [flags]`. No flags before command, and no commands of commands.

### Example

```Go
package main

import "github.com/gowww/cli"

var (
	flagForMain    string // Flag "-m"
	flagForCommand string // Flag "-c"
	flagForAll     string // Flag "-a"
)

func main() {
	cli.SetUsageText("Command line interface example.")

	cli.String(&flagForMain, "m", "", "Example flag for main function.")

	cli.Command("command", command, "Example command.").
		String(&flagForCommand, "c", "", `Example flag for this command only.`)

	cli.String(&flagForAll, "a", "", "Example flag for main function and all commands defined previously.")

	cli.Parse()
}

func command() {
	// Do the command job.
}
```

#### Usage output

##### For `example -help`

```
Command line interface example.

Usage:

	example [command] [flags]

Commands:

	command  Example command.

Flags:

	-a  Example flag for main function and all commands defined previously.
	-m  Example flag for main function.
```

##### For `example command -help`

```
Example command.

Usage:

	example command [flags]

Flags:

	-a  Example flag for main function and all commands defined previously.
	-c  Example flag for this command only.
```
