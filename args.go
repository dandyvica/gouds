package main

import (
	"fmt"
	"flag"
)

// this will hold all cli options/flags
type CliOptions struct {
	domain string
	output string
	stdout bool
}

func GetArgs() CliOptions {
	var options CliOptions

	flag.StringVar(&options.domain, "domain", "uds.sock", "path to the domain socket")
	flag.StringVar(&options.output, "output", "uds.txt", "name of the output file. If not specified, use stdin")
	flag.Parse()

	if options.output == "" {
		options.stdout = true
	}

	fmt.Printf("%#v\n", options)
	
	return options


}