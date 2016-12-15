package main

import (
	"os"
	"flag"
	"fmt"
	"logmarked/logmarker"
)

const VERSION = "0.1.0"

var logging_config = LoggingConfig{Format:DEFAULT_FORMAT, Level:"DEBUG"}

func Usage() {
	guide := `
Usage:
  logmarked [OPTIONS,...]

Options:
`
	os.Stdout.Write([]byte(guide))
	flag.PrintDefaults()
}


func main() {
	option := logmarker.Configuration{}
	log_file := ""
	log_level := "WARN"
	// Global Arguments ================================================================================================
	//Log_File string
	flag.StringVar(&log_file, "l", "", "Logging output file. Default empty means 'stdout'.")
	//Log_Level string
	flag.StringVar(&log_level, "i", "WARN", "Logging level. ")
	// Functional Arguments ============================================================================================
	var config string
	flag.StringVar(&config, "c", "", "Configuration file instead of command line parameters. Default empty means using parameters.")
	var check bool
	flag.BoolVar(&check, "z", false, "Check options.")
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "Display version info.")

	flag.Parse()

	if showVersion {
		os.Stderr.Write([]byte(fmt.Sprintf("logmarked v%v\n", VERSION)))
		os.Exit(0)
	}
	os.Stderr.Write([]byte(fmt.Sprintf("logmarked v%v - A log processor.\n", VERSION)))
	os.Stderr.Write([]byte("Copyright (C) 2015 Mingcai SHEN <archsh@gmail.com>. Licensed for use under the GNU GPL version 3.\n"))
	if config == "" {
		os.Stderr.Write([]byte("\n\n!!! Configuration file is required!\n"))
		Usage()
		os.Exit(1)
	}
	if e:= logmarker.LoadConfiguration(config, &option); e != nil {
		os.Stderr.Write([]byte(fmt.Sprintf("\n\n!!! Load Configuration file '%s' failed: %s \n", config, e)))
		os.Exit(1)
	}

	if check {
		logmarker.CheckConfiguration(&option, os.Stderr)
		os.Exit(0)
	}

	logging_config.Filename = log_file
	logging_config.Level = log_level
	if log_file != "" {
		InitializeLogging(&logging_config, false, logging_config.Level)
	}else{
		InitializeLogging(&logging_config, true, logging_config.Level)
	}
	defer DestroyLogging()

	//os.Stderr.Write([]byte(fmt.Sprintf("Config: \n %+v \n", option)))

	if e := logmarker.Run(&option); nil != e {
		os.Stderr.Write([]byte(fmt.Sprintf("\n\n!!! Run LogMarkerStart failed: %s \n", e)))
		os.Exit(1)
	}else{
		os.Stderr.Write([]byte("\n\n!!! Mission Acomplished!\n"))
	}
}
