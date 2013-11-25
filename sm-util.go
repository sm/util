package main

import (
	"./json"
	"./mustache"
	"flag"
	"fmt"
	"log"
	"os"
)

var data = flag.String("data", "key1=value1:~key2=value2", "key=value data pair, overrides json file data")
var jsonString = flag.String("json", "{string}", "json data string")
var output = flag.String("output", "{file}", "path to output file")
var path = flag.String("path", "{path}", "path of variable")
var template = flag.String("template", "{file}", "path to template file")
var uri = flag.String("uri", "{uri}", "json uri or path to file")

func main() {
	flag.Parse()

	command := os.Args[0]

	if command == "mustache" {
		if *template == "{file}" {
			log.Fatal("ERROR: A template file location must be specified with --template={{path to template file}}")
		}
		mustache.Run(*template, *data, *output)
	} else if command == "json" {
		value := json.Run(uri, path)
		fmt.Println(value)
	} else {
		log.Fatal("ERROR: Unknown Command as first argument, expected {mustache|json}")
	}
}
