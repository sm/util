package json

import (
	"bufio"
	"fmt"
	"encoding/json"
	"errors"
	"reflect"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var value string
var jsonObject map[string]interface{}
var lg,dbg *log.Logger

func init() {
	lg = log.New(os.Stdout, "DEBUG: ", log.Lshortfile)
}

func unit(data interface{}) (value string, err error){
	v := data
	switch v.(type) {
	case string:
		value = v.(string)
	case int,int32,int64:
		value = strconv.Itoa(v.(int))
	case float32, float64:
		value = strconv.FormatFloat(v.(float64),'f',-1,64)
		//var b []byte
		//k := v.(float64)
		//s := strconv.AppendFloat(b, k, 'g', 5, 64)
		//value = string(s)
	default:
		err = errors.New(fmt.Sprintf("ERROR: unit(): Unknown type of '%v': %v",v,reflect.TypeOf(v)))
	}
	return 
}

func list(data []interface{},path []string,segment int) (value string, err error){
	values := []string{}
	for _, v := range data {
		switch v.(type) {
		case string,int,int32,int64,float32,float64:
			value, err = unit(v.(interface{}))
			values = append(values,value)
		case []interface{}:
			if segment < len(path) {
				value, err = list(v.([]interface{}),path,segment+1)
				values = append(values,value)
			} else {
				err = errors.New(fmt.Sprintf("ERROR: list(): more path segment (%i), yet next is a {{List}} '%v' not a {{Value}}...\n",segment, v))
			}
		case map[string]interface{}:
			if segment < len(path) {
				value, err = dict(v.(map[string]interface{}),path,segment+1)
				values = append(values,value)
			} else {
				err = errors.New(fmt.Sprintf("ERROR: list(): more path segment (%i), yet next is a {{Dict}} '%v' not a {{Value}}...\n",segment, v))
			}
		default:
			err = errors.New(fmt.Sprintf("ERROR: list(): Unknown '%v' of type '%v'\n",v,reflect.TypeOf(v)))
		}
	}
	value = strings.Join(values," ")
	return
}

func dict(data map[string]interface{}, path []string, segment int) (value string, err error) {
	key := path[segment]
	v := data[key]

	if v == nil {
		err = errors.New(fmt.Sprintf("ERROR: No key '%s' exists for segment '%i' in path '%v'\n",key,segment,path))
		return
	}

	switch v.(type) {
	case []interface{}:
		// if segment < len(path) {
		value, err = list(v.([]interface{}),path,segment+1)
		// }
	case map[string]interface{}:
		if segment < len(path) {
			value, err = dict(v.(map[string]interface{}),path,segment+1)
		} else {
			err = errors.New(fmt.Sprintf("ERROR: dict(): end of path, yet next is a {{Dict}} '%v' not a {{Value}}...\n",segment, v))
		}
	case string,int,int32,int64,float32,float64,interface{}:
		value, err = unit(v.(interface{}))
	default:
		err = errors.New(fmt.Sprintf("ERROR: dict(): Unknown '%v' of type '%v'\n", v,reflect.TypeOf(v)))
	}
	//}
	return
}

func Run(uri string, path string) (value string, err error) {
	if uri == "{uri}" {
		lg.Fatal("ERROR: A json file location must be specified for reading using --uri={{uri}}")
	}
	file, err := os.Open(uri)
	if err != nil {
		lg.Fatalf("ERROR: The json file at '%s' cannot be opened", uri)
	}

	bfile := bufio.NewReader(file)
	buf, err := ioutil.ReadAll(bfile)
	if file == nil && err != nil {
		lg.Fatalf("ERROR: A json does not exist at '%s'", uri)
	}

	jsonError := json.Unmarshal(buf, &jsonObject)

	if jsonError != nil {
		lg.Fatalf("ERROR: Unable to parse json file located at '%s'", uri)
	}

	pathArray := strings.Split(path, "/")
	if len(pathArray) == 0 {
		err = errors.New("-path length must be non-zero (a/b/c...)")
		return
	}

	value, err = dict(jsonObject, pathArray, 0)

	return
}
