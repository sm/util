package json

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"

	"strings"
)

var f map[string]interface{}
var value string

func findValue(key []string, data map[string]interface{}, index int) string {
	switch v := data[key[index]].(type) {
	case string:
		return v
	case float32, float64:
		var b []byte
		k := v.(float64)
		s := strconv.AppendFloat(b, k, 'g', 5, 64)
		return string(s)
	case map[string]interface{}:
		return findValue(key, v, index+1)
	default:
		log.Fatal("ERROR: Error with value %s, %s", v, reflect.TypeOf(v))
	}
	return ""
}

func findMapValue(data map[string]interface{}) string {
	var value string
	for k, v := range data {
		value += k + " " + v.(string) + " "
	}
	return value
}

func Run(uri *string, path *string) string {
	if *uri == "{uri}" {
		log.Fatal("ERROR: A json file location must be specified for reading using --uri={{uri}}")
	} else {
		file, err := os.Open(*uri)
		if err != nil {
			log.Fatalf("ERROR: The json file at '%s' cannot be opened", *uri)
		}

		bfile := bufio.NewReader(file)
		buf, err := ioutil.ReadAll(bfile)
		if file == nil && err != nil {
			log.Fatalf("ERROR: A json does not exist at '%s'", *uri)
		}

		jsonError := json.Unmarshal(buf, &f)

		if jsonError != nil {
			log.Fatalf("ERROR: Unable to parse json file located at '%s'", *uri)
		}

		// Parse path
		pathArray := strings.Split(*path, "/")
		for i := range pathArray {
			switch v := f[pathArray[i]].(type) {
			case string:
				if len(pathArray) == i+1 {
					value = v
				} else {
					log.Fatalf("ERROR: Unable to traverse the full path before encoutering a value at '%s'.", v)
				}
			case map[string]interface{}:
				if len(pathArray) == i+1 {
					value = findMapValue(v)
				} else {
					value = findValue(pathArray, v, i+1)
				}
				break
			case []interface{}:
				// @TODO make implementation better
				for index, data := range v {
					switch MapData := data.(type) {
					case map[string]interface{}:
						value += findMapValue(MapData)
					default:
						// @TODO: A more informative error message here.
						log.Fatal("ERROR: JSON Data is not correct for index ", index)
					}
				}
			default:
			}
			if len(value) > 0 {
				break
			}
		}
	}
	return value
}
