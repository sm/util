package json

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"testing"
)

var jsonFileName,jsonData string

func init() {
	os.Mkdir("test",0755)
	jsonFileName = "./test/json"
	jsonData = `{"a": {"b": 2, "c": "d", "e": 1.1},"aa": [1,2,3,4]}`
	ioutil.WriteFile(jsonFileName, []byte(jsonData), 0644)
}

func TestJSONIntegration(t *testing.T) {
	Convey("Given a json file", t, func() {
		Convey("path a/b should be '2'", func() {
			value,_ := Run(jsonFileName, "a/b")
			So(value, ShouldEqual, "2")
		})
		Convey("path a/c should be 'd'", func() {
			value,_ := Run(jsonFileName, "a/c")
			So(value, ShouldEqual, "d")
		})
		Convey("path a/e should return 1.1",func() {
			value,_ := Run(jsonFileName, "a/e")
			So(value, ShouldEqual, "1.1")
		})
		Convey("path aa should return '1 2 3 4'",func() {
			value,_ := Run(jsonFileName, "aa")
			So(value, ShouldEqual, "1 2 3 4")
		})
	})
}
