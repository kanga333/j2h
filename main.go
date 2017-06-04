package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const name string = "j2h"

var version string

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s is a tool to convert json to hive ddl\n\n", name)
		fmt.Fprintf(os.Stderr, "Usage: %s <option>\n", name)
		flag.PrintDefaults()
	}
}

func main() {
	var (
		showVersion = flag.Bool("version", false, "Print version information.")
		path        = flag.String("path", "", "Path of json file.")
	)

	flag.Parse()

	if *showVersion {
		fmt.Fprintln(os.Stdout, "Version: ", version)
		os.Exit(0)
	}

	if _, err := os.Stat(*path); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "json file is not exsit in %v\n", *path)
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read json file: %v\n", err)
		os.Exit(1)
	}

	ddl, err := ConvertJSONTOHQL(string(b))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to convert json to hql: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, ddl)
}
