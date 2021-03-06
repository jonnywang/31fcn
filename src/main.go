package main

import (
	"flag"
	"fmt"
	"os"
	"spider"
)

var optionConfigFile = flag.String("config", "./config.xml", "configure xml file")
var version = flag.Bool("version", false, "print current version")

func usage() {
	fmt.Printf("Usage: %s options\nOptions:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 {
		usage()
	}

	if *version {
		fmt.Printf("%s\n", spider.VERSION)
		os.Exit(0)
	}

	_, err := spider.ParseXmlConfig(*optionConfigFile)
	if err != nil {
		spider.Logger.Print(err)
		os.Exit(1)
	}

	spider.Run()
}
