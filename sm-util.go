package main

import (
	"flag"
	"fmt"
	"errors"
	"log"
	"os"
	"github.com/sm/util/json"
	"github.com/sm/util/mustache"
)

var (
	command      string
	mustacheData string
	outputFile   string
	jsonPath     string
	templateFile string
	jsonUri      string
	jsonString   string
)

func init() {
	flag.StringVar(&command, "c", "help", "Command to run, one of {mustache,json}")
	flag.StringVar(&mustacheData, "data", "{data}", "string of key=value pairs separated by a :~")
	flag.StringVar(&outputFile, "output", "{file}", "path to output file")
	flag.StringVar(&jsonPath, "path", "{path}", "path of variable")
	flag.StringVar(&templateFile, "template", "{file}", "path to template file")
	flag.StringVar(&jsonUri, "uri", "{uri}", "json uri or path to file")
	flag.StringVar(&jsonString, "json", "{}", "json string")
	flag.Parse()
}

func errN(msg string) (err error){
	log.Fatal(msg)
	err = errors.New(msg)
	return
}

func main() {
	var err error
	value := ""

	if command == "mustache" {
		if templateFile == "{file}" {
			err = errN("ERROR: A template file location must be specified with -template {{path to template file}}")
		}
		if mustacheData == "{data}" {
			err = errN("ERROR: -data must be given as a string of key=value pairs separated by a :~")
		}
		if outputFile == "{file}" {
			err = errN("ERROR: -output filename with path must be given.")
		}
		value, err = mustache.Run(templateFile, mustacheData, outputFile)
	} else if command == "json" {
		value, err = json.Run(jsonUri, jsonPath)
	} else {
		err = errN("ERROR: Unknown Command or not specified: -c {mustache|json}")
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(value)
	os.Exit(0)
}
