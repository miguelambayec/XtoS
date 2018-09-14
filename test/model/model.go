package testmodel

import (
	"xtos/src/model"
)

type TestTable struct {
	XmlString      string
	Commands       []model.Command
	ExpectedOutput string
}
