package test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"xtos/src/model"
	"xtos/test/model"

	xtos "xtos/src"
)

func TestExecuteXtos(t *testing.T) {

	testTable := initTestTable()

	for _, test := range testTable {

		xtosOutput := xtos.ExecuteXtos(test.XmlString, test.Commands)
		xtos.Reset()

		if removeTabsAndNewLines(xtosOutput) != removeTabsAndNewLines(test.ExpectedOutput) {
			t.Errorf("Output was incorrect!\n\nGot:\n\n%s\n\nWant:\n\n%s", xtosOutput, test.ExpectedOutput)
			break
		}
	}

}

func initTestTable() []testmodel.TestTable {

	testTable := []testmodel.TestTable{

		testmodel.TestTable{
			XmlString:      getDataFromFile("xml/sample1.xml"),
			Commands:       initCommands(false),
			ExpectedOutput: getDataFromFile("struct/sample1.go"),
		},
		testmodel.TestTable{
			XmlString:      getDataFromFile("xml/sample2.xml"),
			Commands:       initCommands(true),
			ExpectedOutput: getDataFromFile("struct/sample2.go"),
		},
		testmodel.TestTable{
			XmlString:      getDataFromFile("xml/sample3.xml"),
			Commands:       initCommands(false),
			ExpectedOutput: getDataFromFile("struct/sample3.go"),
		},
		testmodel.TestTable{
			XmlString:      getDataFromFile("xml/sample4.xml"),
			Commands:       initCommands(true),
			ExpectedOutput: getDataFromFile("struct/sample4.go"),
		},
	}

	return testTable
}

func removeTabsAndNewLines(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, " ", "", -1)

	return str
}

func getDataFromFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func initCommands(notag bool) (Commands []model.Command) {

	Commands = append(Commands, model.Command{
		Name:   "-nt",
		Status: notag,
	})

	return Commands
}
