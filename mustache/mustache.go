package mustache

import (
	"github.com/sm/go-mustache"
	"log"
	"os"
	"strings"
)

var mapData map[string]string

func Run(template string, data string, output string) {

	mapData = make(map[string]string)

	//Error checking on arguements passed in
	file, err := os.Open(template)
	if file == nil && err != nil {
		log.Fatalf("ERROR: Unable to open template file '%s'", template)
	}

	//Split up the data set passed in
	sets := strings.Split(data, ":~")

	for i := range sets {
		splitData := strings.Split(sets[i], "=")
		if len(splitData) > 1 {
			mapData[splitData[0]] = splitData[1]
		}
	}

	fileData := mustache.RenderFile(template, mapData)

	if output != "{file}" {
		outputFile, err := os.Create(output)
		if err != nil {
			log.Fatalf("ERROR: Unable to create output file %s", output)
		}
		outputFile.WriteString(fileData)
	}
}
