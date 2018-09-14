package xtos

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"xtos/src/model"
)

var StructHolder []model.Struct
var output string

func ExecuteXtos(xmlString string, Commands []model.Command) string {

	notag := Commands[0].Status

	// RECURSION OF CREATION OF STRUCT
	createStruct(xmlString, "MAIN", []model.Attribute{}, notag)

	// DISPLAYS THE STRUCT ON THE
	buildOutput(notag)

	return output
}

func Reset() {
	output = ""
	StructHolder = []model.Struct{}
}

func DisplayOutput(Commands []model.Command) {

	writetofile := Commands[1].Status

	if writetofile {
		writeToFile("xml_to_struct_output.go")
		fmt.Println("Successfully generated a file containing the struct for the xml")
	} else {
		fmt.Println("\n" + output + "\n")
	}
}

// INITIALIZE Commands
// ALL ARE DEACTIVATED AT FIRST
func InitializeCommands() (Commands []model.Command) {
	Commands = append(Commands, model.Command{
		Name:   "-nt",
		Status: false,
	})

	Commands = append(Commands, model.Command{
		Name:   "-wo",
		Status: false,
	})

	return Commands
}

// 	ESPECIALLY FOR SOAP XML
func removePrefix(name string) string {

	if strings.Contains(name, ":") {
		return strings.Split(name, ":")[1]
	}

	return name
}

func writeToFile(path string) {

	// write the whole body at once
	err := ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}

func GetXmlStringAndCommands(Commands []model.Command) (string, []model.Command) {
	b, err := ioutil.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Print(err)
	}

	// Detect commands if called
	for key, Command := range Commands {

		for i := 1; i <= len(os.Args)-1; i++ {

			if os.Args[i] == Command.Name {

				Commands[key].Status = true
			}
		}
	}

	return string(b), Commands
}

func createStruct(str string, structName string, structAttrs []model.Attribute, notag bool) {

	// CREATION OF STRUCT TO BE ADDED IN THE HOLDER

	Struct := model.Struct{}
	Struct.StructName = formatName(structName)

	if !notag {
		structAttrs = append(structAttrs, model.Attribute{
			Name:     "XMLName",
			DataType: "xml.Name",
			Tag:      "`" + `xml:"` + structName + `"` + "`",
		})
	}
	Struct.Attributes.Attribute = structAttrs

	// IF STRINGS STILL HAS STRUCT VALUE INSIDE
	for hasStructValue(str) {

		var re = regexp.MustCompile(`<([^<]*)>`)

		// GET PARENT TAG
		rawParent := re.FindStringSubmatch(str)[1]

		// GET ATTR OF RAW PARENT
		separatedParent, attrs := getAttrs(rawParent, notag)

		// GET PARENT ATTRIBUTES
		re = regexp.MustCompile("<" + rawParent + `>([\s\S]*?)</` + separatedParent + ">")

		content := re.FindAllStringSubmatch(str, -1)

		// IF SELF CLOSING TAG
		if content == nil {

			re = regexp.MustCompile("<" + rawParent + ">")
			content = re.FindAllStringSubmatch(str, -1)

			// OPTIONAL
			if content == nil {
				log.Println("This is nil <" + rawParent + ">")
			}

			innerStringValue := content[0][0]

			// SET ATTRIBUTE
			tmpAttribute := model.Attribute{}
			tmpAttribute.Name = formatName(separatedParent)

			if !notag {
				tmpAttribute.Tag = "`" + `xml:"` + separatedParent + `"` + "`"
			}

			if len(attrs) != 0 {
				tmpAttribute.DataType = tmpAttribute.Name
			} else {
				tmpAttribute.DataType = "string"
			}

			// ADD THE ATTRIBUTE TO THE STRUCT
			Struct.Attributes.Attribute = append(Struct.Attributes.Attribute, tmpAttribute)

			// DELETE STRUCT ATTRIBUTE FROM THE STRING
			str = strings.Replace(str, innerStringValue, "", -1)

			if len(attrs) != 0 {
				/* STRUCT CREATION FOR SELF CLOSING ATTRIBUTE */

				// CREATE STRUCT FOR SELF CLOSING
				selfClosingStruct := model.Struct{}
				selfClosingStruct.StructName = formatName(separatedParent)

				if !notag {
					attrs = append(attrs, model.Attribute{
						Name:     "XMLName",
						DataType: "xml.Name",
						Tag:      "`" + `xml:"` + separatedParent + `"` + "`",
					})
				}

				selfClosingStruct.Attributes.Attribute = attrs

				// ADD STRUCT TO HOLDER
				if !structExists(selfClosingStruct) {
					StructHolder = append(StructHolder, selfClosingStruct)
				}

				/* END OF CREATION */
			}

		} else {

			innerStringValue := content[0][1]

			// SET ATTRIBUTE
			tmpAttribute := model.Attribute{}
			tmpAttribute.Name = formatName(separatedParent)

			if !notag {
				tmpAttribute.Tag = "`" + `xml:"` + separatedParent + `"` + "`"
			}

			if hasStructValue(innerStringValue) || hasAttrs(rawParent) {
				tmpAttribute.DataType = tmpAttribute.Name
				createStruct(innerStringValue, separatedParent, attrs, notag)
			} else {
				tmpAttribute.DataType = detectDataType(string(innerStringValue))
			}

			// ADD THE ATTRIBUTE TO THE STRUCT
			Struct.Attributes.Attribute = append(Struct.Attributes.Attribute, tmpAttribute)

			// DELETE STRUCT ATTRIBUTE FROM THE STRING
			str = strings.Replace(str, content[0][0], "", -1)

		}

	}

	// ADD STRUCT TO HOLDER
	if !structExists(Struct) {
		StructHolder = append(StructHolder, Struct)
	}

}

// CHECKS IF THE RAW PARENT HAS ATTRS
func hasAttrs(rawParent string) bool {
	return strings.Contains(rawParent, "=")
}

// REMOVES SPECIAL CHARACTERS FROM A STRING
// THIS IS USED FOR STRUCT NAMES AND ATTRIBUTE NAMES
func formatName(name string) string {
	name = removePrefix(name)
	name = strings.Title(name)
	re := regexp.MustCompile("[$&+,:;=?@#|'<>.^*()%!-/_]")
	specialCharacters := re.FindAllStringSubmatch(name, -1)

	for _, character := range specialCharacters {

		var tmpName string
		separatedNames := strings.Split(name, character[0])
		for _, name := range separatedNames {
			tmpName += strings.Title(name)
		}
		name = tmpName
		// name = strings.Replace(name, character[0], "", -1)
	}
	return name
}

// USED FOR DISPLAYING ALL THE STRUCTS IN THE TERMINAL
func buildOutput(notag bool) {

	output += "package model\n\n"

	if !notag {
		output += `import "encoding/xml"` + "\n\n"
	}

	for i := len(StructHolder) - 2; i >= 0; i-- {

		output += "type " + StructHolder[i].StructName + " struct {\n"

		// DETECTS ARRAY ATTRIBUTE
		// IF ARRAY, ATTRIBUTE WILL BE DISPLAYED AS []<Attribute Name>
		detectArrayAttribute(StructHolder[i].Attributes.Attribute, notag)

		output += "}\n\n"

	}
}

// DETECTS ARRAY ATTRIBUTE
// IF ARRAY, ATTRIBUTE WILL BE DISPLAYED AS []<Attribute Name>
func detectArrayAttribute(Attributes []model.Attribute, notag bool) {

	// GET THE COUNT OF ATTRIBUTE
	AttributeCount := []model.AttributeCount{}

	for _, Attribute := range Attributes {

		count := 0
		for _, attr := range Attributes {

			if Attribute.Name == attr.Name {
				count++
			}
		}

		tmp := model.AttributeCount{
			Name:  Attribute.Name,
			Count: count,
		}

		if !elementInArray(tmp, AttributeCount) {
			AttributeCount = append(AttributeCount, tmp)
		}
	}

	// SORT ATTRIBUTES
	for key, data := range AttributeCount {
		for cpy_key, cpy_data := range AttributeCount {

			if cpy_data.Name > data.Name {

				tmp := AttributeCount[cpy_key]
				AttributeCount[cpy_key] = AttributeCount[key]
				AttributeCount[key] = tmp
			}

		}
	}

	// DISPLAY XMLNAME AS THE FIRST ATTRIBUTE
	for _, Attribute := range Attributes {

		if "XMLName" == Attribute.Name {

			output += attributeIfNoTag(notag, Attribute, "single")
		}
	}

	// DISPLAY THE REST OF ATTRIBUTES
	for _, attrData := range AttributeCount {

		for _, Attribute := range Attributes {

			if attrData.Name == Attribute.Name && attrData.Name != "XMLName" {

				if attrData.Count > 1 {

					output += attributeIfNoTag(notag, Attribute, "array")
					break
				} else {

					output += attributeIfNoTag(notag, Attribute, "single")
				}
			}
		}
	}

}

func attributeIfNoTag(notag bool, Attribute model.Attribute, attrType string) string {

	// if notag {
	//
	// 	if attrType == "array" {
	// 		return "\t" + Attribute.Name + " []" + Attribute.DataType + "\n"
	// 	} else {
	// 		return "\t" + Attribute.Name + " " + Attribute.DataType + "\n"
	// 	}
	// } else {
	//
	// 	if attrType == "array" {
	// 		return "\t" + Attribute.Name + " []" + Attribute.DataType + " `" + Attribute.Tag + "`\n"
	// 	} else {
	// 		return "\t" + Attribute.Name + " " + Attribute.DataType + " `" + Attribute.Tag + "`\n"
	// 	}
	// }

	if attrType == "array" {
		return "\t" + Attribute.Name + " []" + Attribute.DataType + " " + Attribute.Tag + "\n"
	} else {
		return "\t" + Attribute.Name + " " + Attribute.DataType + " " + Attribute.Tag + "\n"
	}
}

// CHECKS IF AN ELEMENT EXISTS IN THE LIST
func elementInArray(element model.AttributeCount, list []model.AttributeCount) bool {

	for _, data := range list {
		if element.Name == data.Name {
			return true
		}
	}

	return false
}

// GETS THE ATTRS OF THE RAW PARENT
func getAttrs(rawParent string, notag bool) (string, []model.Attribute) {

	/* STRING ADJUSMENTS */
	// THIS IS DUE TO XML EXTRA SPACES, USAGE OF SINGLE QUOTE INSTEAD OF DOUBLE QUOTES IN ATTRIBUTE VALUES

	// REPLACE ' TO "
	rawParent = strings.Replace(rawParent, "'", `"`, -1)

	// remove tabs and new lines
	rawParent = strings.Replace(rawParent, "\t", " ", -1)
	rawParent = strings.Replace(rawParent, "\n", " ", -1)

	for spacesMorethanOne(rawParent) {
		rawParent = strings.Replace(rawParent, "  ", " ", -1)
	}
	rawParent = strings.Replace(rawParent, `" /`, `"/`, -1)

	// REMOVE " " AT THE END
	if string(rawParent[len(rawParent)-1]) == " " {
		rawParent = rawParent[:len(rawParent)-1]
	}

	/* END OF STRING ADJUSTMENTS */

	// REPLACE SPACES INTO DIFFERENT INSIDE THE VALUES
	// THIS IS DONE BECAUSE I WILL USE SPLIT FUNCTION TO SEPARATE ALL THE ATTRS INTO ARRAY
	re := regexp.MustCompile(`"([\s\S]*?)"`)
	values := re.FindAllStringSubmatch(rawParent, -1)

	// COPY THE VALUES
	var valuesCpy = [][]string{}
	for _, val := range values {
		valuesCpy = append(valuesCpy, []string{
			val[0],
			val[1],
		})
	}

	// REPLACE SPACES TO DIFFERENT CHARACTER
	for key, val := range valuesCpy {
		val[1] = strings.Replace(val[1], " ", "~", -1)
		rawParent = strings.Replace(rawParent, values[key][1], val[1], -1)
	}

	// SEPARATE XMLNAME TO ATTRS
	separatedAttr := strings.Split(rawParent, " ")

	// GET ATTRS
	attributes := []model.Attribute{}
	for i := 1; i < len(separatedAttr); i++ {

		// SET ATTRIBUTE
		tmpAttribute := model.Attribute{}

		// SET ATTR NAME
		attr := strings.Split(separatedAttr[i], "=")
		tmpAttribute.Name = formatName(attr[0])

		if tmpAttribute.Name != "" {

			// GET ATTR VALUE
			rawAttrValue := attr[1]
			re = regexp.MustCompile(`"([\s\S]*?)"`)
			attrValue := re.FindStringSubmatch(rawAttrValue)[1]
			tmpAttribute.DataType = detectDataType(attrValue)

			// SET ATTR TA
			if !notag {
				tmpAttribute.Tag = "`" + `xml:"` + attr[0] + `,attr"` + "`"
			} else {
				tmpAttribute.Tag = "`" + `xml:",attr"` + "`"
			}

			// ADD THE ATTRIBUTE TO THE STRUCT
			attributes = append(attributes, tmpAttribute)
		}

	}

	return separatedAttr[0], attributes
}

// CHECKS IF STRINGS HAS MORE THAN ONE SPACES
// THIS FUNCTION WILL BE USED TO REDUCE ALL THE SPACES MORE THAN ONE
func spacesMorethanOne(str string) bool {
	consecutive := 0
	for i := 0; i < len(str); i++ {

		if string(str[i]) == " " {
			consecutive++

			if consecutive > 1 {
				return true
			}
		} else {
			consecutive = 0
		}
	}
	return false
}

// CHECKS THE STRUCT EXISTS IN THE HOLDER
// USED FOR PREVENTION OF STRUCT REDUNDANCY
func structExists(paramStruct model.Struct) bool {

	for _, Struct := range StructHolder {

		if paramStruct.StructName == Struct.StructName {
			return true
		}
	}

	return false
}

// DETECTS THE DATA TYPE OF A VALUE
func detectDataType(variable string) string {

	// FLOAT
	f, err := strconv.ParseFloat(variable, 64)
	if err == nil {

		if strings.Contains(variable, ".") {
			return reflect.TypeOf(f).String()
		}
	}

	// INT
	i, err := strconv.ParseInt(variable, 10, 64)
	if err == nil {
		return reflect.TypeOf(i).String()
	}

	// BOOL
	b, err := strconv.ParseBool(variable)
	if err == nil {
		return reflect.TypeOf(b).String()
	}

	return "string"

}

// CHECKS IF THE STRING HAS A STRUCT VALUE
func hasStructValue(str string) bool {

	return strings.Contains(str, "<")
}
