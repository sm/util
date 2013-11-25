package main

import (
	"flag"
	"fmt"
	"github.com/sm/util/json"
	"github.com/sm/util/mustache"
	"log"
	//"os"
)

var (
	command      string
	mustacheData string
	outputFile   string
	jsonPath     string
	templateFile string
	jsonUri      string
)

func main() {
	flag.StringVar(&command, "c", "help", "Command to run, one of {mustache,json}")
	flag.StringVar(&mustacheData, "data", "{data}", "string of key=value pairs separated by a :~")
	flag.StringVar(&outputFile, "output", "{file}", "path to output file")
	flag.StringVar(&jsonPath, "path", "{path}", "path of variable")
	flag.StringVar(&templateFile, "template", "{file}", "path to template file")
	flag.StringVar(&jsonUri, "uri", "{uri}", "json uri or path to file")
	flag.Parse()

	fmt.Println(command)
	fmt.Println(mustacheData)
	fmt.Println(templateFile)

	if command == "mustache" {
		fmt.Println(templateFile)
		if templateFile == "{file}" {
			log.Fatal("ERROR: A template file location must be specified with -template {{path to template file}}")
		}
		if mustacheData == "{data}" {
			log.Fatal("ERROR: -data must be given as a string of key=value pairs separated by a :~")
		}
		if outputFile == "{file}" {
			log.Fatal("ERROR: -output filename with path must be given.")
		}
		mustache.Run(templateFile, mustacheData, outputFile)
	} else if command == "json" {
		value := json.Run(jsonUri, jsonPath)
		fmt.Println(value)
	} else {
		log.Fatal("ERROR: Unknown Command or not specified: -c {mustache|json}")
	}
}
