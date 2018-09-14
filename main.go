package main

import "xtos/src"

func main() {
	Commands := xtos.InitializeCommands()

	// GET XML STRING FROM OTHER FILE
	xmlString, Commands := xtos.GetXmlStringAndCommands(Commands)

	xtos.ExecuteXtos(xmlString, Commands)

	xtos.DisplayOutput(Commands)
}
