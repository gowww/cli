package cli_test

import "github.com/gowww/cli"

var (
	flagForMain    string // Flag "-m"
	flagForCommand string // Flag "-c"
	flagForAll     string // Flag "-a"
)

func Example() {
	cli.Description = "Command line interface example."

	cli.String(&flagForMain, "m", "", "Example flag for main function.")

	cli.Command("command", command, "Example command.").
		String(&flagForCommand, "c", "", `Example flag for this command only.`)

	cli.String(&flagForAll, "a", "", "Example flag for main function and all commands defined previously.")

	cli.Parse()
}

func command() {
	// Do the command job.
}
