package testmodel

import (
	"xtos/src/model"
)

// for testing
type TestTable struct {
	XmlString      string
	Commands       []model.Command
	ExpectedOutput string
}
