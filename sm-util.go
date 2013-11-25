package main

import (
	"flag"
	"fmt"
	"github.com/sm/util/json"
	"github.com/sm/util/mustache"
	"log"
	"os"
)

var (
	data     string
	output   string
	path     string
	template string
	uri      string
)

func Init() {
	//data = flag.String("data", "key1=value1:~key2=value2", "key=value data pair, overrides json file data")
	flag.StringVar(&data, "data", "key1=value1:~key2=value2", "key=value data pair, overrides json file data")

	//output = flag.String("output", "{file}", "path to output file")
	flag.StringVar(&output, "output", "{file}", "path to output file")
	//path = flag.String("path", "{path}", "path of variable")
	flag.StringVar(&path, "path", "{path}", "path of variable")
	//template = flag.String("template", "{file}", "path to template file")
	flag.StringVar(&template, "template", "{file}", "path to template file")
	//uri = flag.String("uri", "{uri}", "json uri or path to file")
	flag.StringVar(&uri, "uri", "{uri}", "json uri or path to file")

	flag.Parse()
}

func main() {
	command := os.Args[1]

	if command == "mustache" {
		fmt.Println(template)
		if template == "{file}" {
			log.Fatal("ERROR: A template file location must be specified with -template {{path to template file}}")
		}
		mustache.Run(template, data, output)
	} else if command == "json" {
		value := json.Run(uri, path)
		fmt.Println(value)
	} else {
		log.Fatal("ERROR: Unknown Command as first argument, expected {mustache|json}")
	}
}
