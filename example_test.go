package cli_test

import "github.com/gowww/cli"

var (
	address    string
	production bool
)

func run() {
	// Run app.
}

func watch() {
	// Detect changes and rerun app.
}

func Example() {
	cli.Description = "My app command line interface."

	cli.String(&address, "address", ":8080", "The address to listen and serve on.")

	cli.Command("run", run, "Run app.").
		String(&address, "address", ":8080", "The address to listen and serve on.").
		Bool(&production, "docker", false, `Run the server in production environment.`)

	cli.Command("watch", watch, "Detect changes and rerun app.")

	cli.Parse()
}
