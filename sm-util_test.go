package main

import (
	"./json"
	"./mustache"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"testing"
)

func TestMustache(t *testing.T) {
	Convey("Given a mustache template file containing '{{key1}}\n{{key2}}'", t, func() {
		wd, _ := os.Getwd()
		outputFile, _ := ioutil.TempFile(wd, "test-output")
		outputFileName := outputFile.Name()
		templateFile, _ := ioutil.TempFile(wd, "test-template")
		templateFileName := templateFile.Name()

		Convey(`The output file should contain 'value1\nvalue2'. `, func() {
			ioutil.WriteFile(templateFileName, []byte("{{key1}}\n{{key2}}"), 0644)

			mustache.Run(templateFileName, "key1=value1:~key2=value2", outputFileName)

			content, _ := ioutil.ReadFile(outputFileName)

			So(string(content), ShouldEqual, "value1\nvalue2")

			os.Remove(outputFileName)
			os.Remove(templateFileName)
		})
	})
}

func TestJSON(t *testing.T) {
	Convey(`Given a json file containing {"a": {"b": 2, "c": "d"}}`, t, func() {
		jsonData := `{"a": {"b": 2, "c": "d"}}`
		wd, _ := os.Getwd()

		Convey("path a/b should be '2'", func() {
			jsonFile, _ := ioutil.TempFile(wd, "test-json")
			jsonFileName := jsonFile.Name()
			ioutil.WriteFile(jsonFileName, []byte(jsonData), 0644)
			defer os.Remove(jsonFileName)
			jsonPath := "a/b"
			value := json.Run(jsonFileName, jsonPath)
			So(value, ShouldEqual, "2")
		})

		Convey("path a/c should be 'd'", func() {
			jsonFile, _ := ioutil.TempFile(wd, "test-json")
			jsonFileName := jsonFile.Name()
			ioutil.WriteFile(jsonFileName, []byte(jsonData), 0644)
			defer os.Remove(jsonFileName)
			jsonPath := "a/c"
			value := json.Run(jsonFileName, jsonPath)
			So(value, ShouldEqual, "d")
		})

	})
}
