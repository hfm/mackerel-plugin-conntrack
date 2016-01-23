package main

import (
	"flag"
	"fmt"
	"os"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

// graphdef is Graph definition for mackerelplugin.
var graphdef = map[string](mp.Graphs){
	"conntrack.count": mp.Graphs{
		Label: "Conntrack Count",
		Unit:  "integer",
		Metrics: [](mp.Metrics){
			mp.Metrics{Name: "*", Label: "%1", Diff: false, Stacked: true, Type: "uint64"},
		},
	},
}

// ConntrackPlugin mackerel plugin for *_conntrack.
type ConntrackPlugin struct {}

// GraphDefinition interface for mackerelplugin.
func (c ConntrackPlugin) GraphDefinition() map[string](mp.Graphs) {
	return graphdef
}

// FetchMetrics interface for mackerelplugin.
func (c ConntrackPlugin) FetchMetrics() (map[string]interface{}, error) {
	conntrack_count, err := CurrentValue(ConntrackCountPaths)
	if err != nil {
		return nil, err
	}

	conntrack_max, err := CurrentValue(ConntrackMaxPaths)
	if err != nil {
		return nil, err
	}

	stat := make(map[string]interface{})
	stat["conntrack.count.used"] = conntrack_count
	stat["conntrack.count.free"] = (conntrack_max - conntrack_count)

	return stat, nil
}

// Parse flags and Run helper (MackerelPlugin) with the given arguments.
func main() {
	// Flags
	var (
		tempfile string
		version  bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.BoolVar(&version, "version", false, "Print version information and quit.")
	flags.StringVar(&tempfile, "tempfile", "/tmp/mackerel-plugin-conntrack", "Temp file name")

	// Parse commandline flag
	if err := flags.Parse(os.Args[1:]); err != nil {
		os.Exit(ExitCodeParseFlagError)
	}

	// Show version
	if version {
		fmt.Fprintf(os.Stderr, "%s version %s\n", Name, Version)
		os.Exit(ExitCodeOK)
	}

	// Create MackerelPlugin for Conntrack
	var c ConntrackPlugin
	helper := mp.NewMackerelPlugin(c)
	helper.Tempfile = tempfile

	helper.Run()
}
