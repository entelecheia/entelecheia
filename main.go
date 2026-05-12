package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/entelecheia/entelecheia/internal/profile"
)

var Version = "dev"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr, profile.Open))
}

func run(args []string, stdout io.Writer, stderr io.Writer, open func(string) error) int {
	color := profile.SupportsColor(os.Stdout)

	if len(args) == 0 {
		profile.PrintIntro(stdout, color)
		return 0
	}

	switch args[0] {
	case "intro":
		if len(args) > 1 {
			return usageError(stderr, "intro does not accept arguments")
		}
		profile.PrintIntro(stdout, color)
		return 0
	case "open":
		if len(args) > 2 {
			return usageError(stderr, "open accepts at most one site alias")
		}
		alias := "home"
		if len(args) == 2 {
			alias = args[1]
		}
		site, ok := profile.FindSite(alias)
		if !ok {
			fmt.Fprintf(stderr, "unknown site alias: %s\n", alias)
			fmt.Fprintln(stderr, "Run 'entelecheia links' to see available aliases.")
			return 1
		}
		if err := open(site.URL); err != nil {
			fmt.Fprintf(stderr, "failed to open %s: %v\n", site.URL, err)
			return 1
		}
		fmt.Fprintf(stdout, "Opening %s: %s\n", site.Label, site.URL)
		return 0
	case "links":
		if len(args) > 1 {
			return usageError(stderr, "links does not accept arguments")
		}
		profile.PrintLinks(stdout)
		return 0
	case "version":
		if len(args) > 1 {
			return usageError(stderr, "version does not accept arguments")
		}
		fmt.Fprintln(stdout, Version)
		return 0
	case "-v", "--version":
		if len(args) > 1 {
			return usageError(stderr, "version flag does not accept arguments")
		}
		fmt.Fprintln(stdout, Version)
		return 0
	case "-h", "--help", "help":
		printHelp(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown command: %s\n", args[0])
		printHelp(stderr)
		return 1
	}
}

func usageError(stderr io.Writer, msg string) int {
	fmt.Fprintf(stderr, "%s\n\n", msg)
	printHelp(stderr)
	return 1
}

func printHelp(w io.Writer) {
	fmt.Fprint(w, strings.TrimSpace(`
entelecheia - Young Joon Lee's GitHub profile CLI

Usage:
  entelecheia              Show intro
  entelecheia intro        Show intro
  entelecheia open [alias] Open a site in the default browser
  entelecheia links        List site aliases and URLs
  entelecheia version      Show version
  entelecheia --version    Show version
  entelecheia --help       Show help

Examples:
  entelecheia open
  entelecheia open halla
  entelecheia links
`)+"\n")
}
