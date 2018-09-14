# xtos
Xml to Struct generator for Go


## Install
`go get github.com/miguelambayec/xtos/`


## Usage


**Syntax**

`xtos <settings> <filename>`


**Settings**
```
-nt   no tag
-wo   generate xml to file
```

**Example**


Prints xml output to the terminal with no tags


sample.xml

```
<Person status="broken">
  <name>Arnold</name>
  <age>50</age>
</Person>
```

Execute XtoS to terminal

`$ xtos sample.xml`


Output to terminal

```
type Person struct {
  XMLName xml.Name `xml:"Person"`
  Status string `xml:"status,attr"`
  Name string `xml:"name"`
  Age int `xml:"age"`
}
```

