package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	// version and githash are used to inject build-time values i.e.
	// go build -ldflags "-X main.date=`date -I` -X main.githash=`git rev-parse --short HEAD`"
	version = "0.0.1"
	date    = "use -ldflags \"-X main.date=`date -I`\""
	githash = "use -X main.githash=`git rev-parse --short HEAD`\""
)

func main() {
	confFile := flag.String("config", "", "(absolute) path and name of the configuration file")
	versionFlag := flag.Bool("version", false, "version of the running code")
	writeConfig := flag.Bool("write_config", false, "write sample config.json file to local directory")
	flag.Parse()
	if *versionFlag {
		fmt.Printf("%s\tversion=%s\tdate=%s\tgithash=%s\n", os.Args[0], version, date, githash)
		os.Exit(0)
	}
	if *writeConfig {
		writeConfigFile()
		os.Exit(0)
	}
	cf := strings.TrimSpace(*confFile)
	if cf == "" {
		fmt.Println("No config file specified. Use -help to see the options")
		os.Exit(0)
	}

	app := initConfig(cf)

	app.runListenAndServeDaemon()

	// Notnagel
	app.logFile.Close()
}
