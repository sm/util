package mustache

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"testing"
)

var mustacheFile, mustacheData, outputFile string

func init() {
	os.Mkdir("test",0755)
	mustacheFile = "./test/mustache"
	mustacheData = "{{key1}}\n{{key2}}"
	outputFile = "./test/mustache-output"
	ioutil.WriteFile(mustacheFile, []byte(mustacheData), 0644)
}

func TestMustache(t *testing.T) {

	Convey("Given a mustache template file", t, func() {
		Convey("The output file should contain 'value1\nvalue2'. ", func() {
			Run(mustacheFile, "key1=value1:~key2=value2", outputFile)
			content, _ := ioutil.ReadFile(outputFile)
			So(string(content), ShouldEqual, "value1\nvalue2")
		})
	})
}

