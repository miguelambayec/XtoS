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

var xmlStrings []string
var outputStrings []string

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

	initXmlStrings()

	testTable := []testmodel.TestTable{}

	for i := 0; i < 4; i++ {

		notag := false

		if i%2 != 0 {
			notag = true
		}

		testTable = append(testTable, testmodel.TestTable{
			XmlString:      xmlStrings[i],
			Commands:       initCommands(notag),
			ExpectedOutput: outputStrings[i],
		})
	}

	return testTable
}

func initXmlStrings() {

	xmlStrings = []string{
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
		  <soapenv:Header>
		    <ns1:RequestHeader soapenv:actor="http://schemas.xmlsoap.org/soap/actor/next" soapenv:mustUnderstand="0" xmlns:ns1="https://www.google.com/apis/ads/publisher/v201808">
		      <ns1:networkCode>123456</ns1:networkCode>
		      <ns1:applicationName>DfpApi-Java-2.1.0-dfp_test</ns1:applicationName>
		    </ns1:RequestHeader>
		  </soapenv:Header>
		  <soapenv:Body>
		    <getAdUnitsByStatement xmlns="https://www.google.com/apis/ads/publisher/v201808">
		      <filterStatement>
		        <query>WHERE parentId IS NULL LIMIT 500</query>
		      </filterStatement>
		    </getAdUnitsByStatement>
		  </soapenv:Body>
		</soapenv:Envelope>`,

		`<catalog>
		   <product description="Cardigan Sweater" product_image="cardigan.jpg">
		      <catalog_item gender="Men's">
		         <item_number>QWZ5671</item_number>
		         <price>39.95</price>
		         <size description="Medium">
		            <color_swatch image="red_cardigan.jpg">Red</color_swatch>
		            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
		         </size>
		         <size description="Large">
		            <color_swatch image="red_cardigan.jpg">Red</color_swatch>
		            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
		         </size>
		      </catalog_item>
		      <catalog_item gender="Women's">
		         <item_number>RRX9856</item_number>
		         <price>42.50</price>
		         <size description="Small">
		            <color_swatch image="red_cardigan.jpg">Red</color_swatch>
		            <color_swatch image="navy_cardigan.jpg">Navy</color_swatch>
		            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
		         </size>
		         <size description="Medium">
		            <color_swatch image="red_cardigan.jpg">Red</color_swatch>
		            <color_swatch image="navy_cardigan.jpg">Navy</color_swatch>
		            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
		            <color_swatch image="black_cardigan.jpg">Black</color_swatch>
		         </size>
		         <size description="Large">
		            <color_swatch image="navy_cardigan.jpg">Navy</color_swatch>
		            <color_swatch image="black_cardigan.jpg">Black</color_swatch>
		         </size>
		         <size description="Extra Large">
		            <color_swatch image="burgundy_cardigan.jpg">Burgundy</color_swatch>
		            <color_swatch image="black_cardigan.jpg">Black</color_swatch>
		         </size>
		      </catalog_item>
		   </product>
		</catalog>`,

		`<Root xmlns="http://www.adventure-works.com">
		  <Customers>
		    <Customer CustomerID="GREAL">
		      <CompanyName>Great Lakes Food Market</CompanyName>
		      <ContactName>Howard Snyder</ContactName>
		      <ContactTitle>Marketing Manager</ContactTitle>
		      <Phone>(503) 555-7555</Phone>
		      <FullAddress>
		        <Address>2732 Baker Blvd.</Address>
		        <City>Eugene</City>
		        <Region>OR</Region>
		        <PostalCode>97403</PostalCode>
		        <Country>USA</Country>
		      </FullAddress>
		    </Customer>
		    <Customer CustomerID="HUNGC">
		      <CompanyName>Hungry Coyote Import Store</CompanyName>
		      <ContactName>Yoshi Latimer</ContactName>
		      <ContactTitle>Sales Representative</ContactTitle>
		      <Phone>(503) 555-6874</Phone>
		      <Fax>(503) 555-2376</Fax>
		      <FullAddress>
		        <Address>City Center Plaza 516 Main St.</Address>
		        <City>Elgin</City>
		        <Region>OR</Region>
		        <PostalCode>97827</PostalCode>
		        <Country>USA</Country>
		      </FullAddress>
		    </Customer>
		  </Customers>
		  <Orders>
		    <Order>
		      <CustomerID>LETSS</CustomerID>
		      <EmployeeID>6</EmployeeID>
		      <OrderDate>1997-11-10T00:00:00</OrderDate>
		      <RequiredDate>1997-12-08T00:00:00</RequiredDate>
		      <ShipInfo ShippedDate="1997-11-21T00:00:00">
		        <ShipVia>2</ShipVia>
		        <Freight>45.97</Freight>
		        <ShipName>Let's Stop N Shop</ShipName>
		        <ShipAddress>87 Polk St. Suite 5</ShipAddress>
		        <ShipCity>San Francisco</ShipCity>
		        <ShipRegion>CA</ShipRegion>
		        <ShipPostalCode>94117</ShipPostalCode>
		        <ShipCountry>USA</ShipCountry>
		      </ShipInfo>
		    </Order>
		    <Order>
		      <CustomerID>LETSS</CustomerID>
		      <EmployeeID>4</EmployeeID>
		      <OrderDate>1998-02-12T00:00:00</OrderDate>
		      <RequiredDate>1998-03-12T00:00:00</RequiredDate>
		      <ShipInfo ShippedDate="1998-02-13T00:00:00">
		        <ShipVia>2</ShipVia>
		        <Freight>90.97</Freight>
		        <ShipName>Let's Stop N Shop</ShipName>
		        <ShipAddress>87 Polk St. Suite 5</ShipAddress>
		        <ShipCity>San Francisco</ShipCity>
		        <ShipRegion>CA</ShipRegion>
		        <ShipPostalCode>94117</ShipPostalCode>
		        <ShipCountry>USA</ShipCountry>
		      </ShipInfo>
		    </Order>
		  </Orders>
		</Root>`,

		`<bookstore specialty="novel">
		  <book style="autobiography">
		    <author>
		      <first-name>Joe</first-name>
		      <last-name>Bob</last-name>
		      <award>Trenton Literary Review Honorable Mention</award>
		    </author>
		    <price>12</price>
		  </book>
		  <book style="textbook">
		    <author>
		      <first-name>Mary</first-name>
		      <last-name>Bob</last-name>
		      <publication>Selected Short Stories of        <first-name>Mary</first-name>
		      <last-name>Bob</last-name>
		    </publication>
		  </author>
		  <editor>
		    <first-name>Britney</first-name>
		    <last-name>Bob</last-name>
		  </editor>
		  <price>55</price>
		</book>
		<magazine style="glossy" frequency="monthly">
		  <price>2.50</price>
		  <subscription price="24" per="year"/>
		</magazine>
		<book style="novel" id="myfave">
		  <author>
		    <first-name>Toni</first-name>
		    <last-name>Bob</last-name>
		    <degree from="Trenton U">B.A.</degree>
		    <degree from="Harvard">Ph.D.</degree>
		    <award>Pulitzer</award>
		    <publication>Still in Trenton</publication>
		    <publication>Trenton Forever</publication>
		  </author>
		  <price intl="Canada" exchange="0.7">6.50</price>
		  <excerpt>
		    <p>It was a dark and stormy night.</p>
		    <p>But then all nights in Trenton seem dark and      stormy to someone who has gone through what      <emph>I</emph> have.</p>
		    <definition-list>
		      <term>Trenton</term>
		      <definition>misery</definition>
		    </definition-list>
		  </excerpt>
		</book>
		<my:book xmlns:my="uri:mynamespace" style="leather" price="29.50">
		  <my:title>Who's Who in Trenton</my:title>
		  <my:author>Robert Bob</my:author>
		</my:book>
		</bookstore>`,
	}

	outputStrings = []string{
		`package model

		import "encoding/xml"

		type Envelope struct {
			XMLName xml.Name ` + "`" + `xml:"soapenv:Envelope"` + "`" + `
			Body    Body     ` + "`" + `xml:"soapenv:Body"` + "`" + `
			Header  Header   ` + "`" + `xml:"soapenv:Header"` + "`" + `
			Soapenv string   ` + "`" + `xml:"xmlns:soapenv,attr"` + "`" + `
			Xsd     string   ` + "`" + `xml:"xmlns:xsd,attr"` + "`" + `
			Xsi     string   ` + "`" + `xml:"xmlns:xsi,attr"` + "`" + `
		}

		type Body struct {
			XMLName               xml.Name              ` + "`" + `xml:"soapenv:Body"` + "`" + `
			GetAdUnitsByStatement GetAdUnitsByStatement ` + "`" + `xml:"getAdUnitsByStatement"` + "`" + `
		}

		type GetAdUnitsByStatement struct {
			XMLName         xml.Name        ` + "`" + `xml:"getAdUnitsByStatement"` + "`" + `
			FilterStatement FilterStatement ` + "`" + `xml:"filterStatement"` + "`" + `
			Xmlns           string          ` + "`" + `xml:"xmlns,attr"` + "`" + `
		}

		type FilterStatement struct {
			XMLName xml.Name ` + "`" + `xml:"filterStatement"` + "`" + `
			Query   string   ` + "`" + `xml:"query"` + "`" + `
		}

		type Header struct {
			XMLName       xml.Name      ` + "`" + `xml:"soapenv:Header"` + "`" + `
			RequestHeader RequestHeader ` + "`" + `xml:"ns1:RequestHeader"` + "`" + `
		}

		type RequestHeader struct {
			XMLName         xml.Name ` + "`" + `xml:"ns1:RequestHeader"` + "`" + `
			Actor           string   ` + "`" + `xml:"soapenv:actor,attr"` + "`" + `
			ApplicationName string   ` + "`" + `xml:"ns1:applicationName"` + "`" + `
			MustUnderstand  int64    ` + "`" + `xml:"soapenv:mustUnderstand,attr"` + "`" + `
			NetworkCode     int64    ` + "`" + `xml:"ns1:networkCode"` + "`" + `
			Ns1             string   ` + "`" + `xml:"xmlns:ns1,attr"` + "`" + `
		}`,

		`package model

		type Catalog struct {
			Product Product
		}

		type Product struct {
			CatalogItem  []CatalogItem
			Description  string ` + "`" + `xml:",attr"` + "`" + `
			ProductImage string ` + "`" + `xml:",attr"` + "`" + `
		}

		type CatalogItem struct {
			Gender     string ` + "`" + `xml:",attr"` + "`" + `
			ItemNumber string
			Price      float64
			Size       []Size
		}

		type Size struct {
			ColorSwatch []ColorSwatch
			Description string ` + "`" + `xml:",attr"` + "`" + `
		}

		type ColorSwatch struct {
			Image string ` + "`" + `xml:",attr"` + "`" + `
		}`,

		`package model

		import "encoding/xml"

		type Root struct {
			XMLName   xml.Name  ` + "`" + `xml:"Root"` + "`" + `
			Customers Customers ` + "`" + `xml:"Customers"` + "`" + `
			Orders    Orders    ` + "`" + `xml:"Orders"` + "`" + `
			Xmlns     string    ` + "`" + `xml:"xmlns,attr"` + "`" + `
		}

		type Orders struct {
			XMLName xml.Name ` + "`" + `xml:"Orders"` + "`" + `
			Order   []Order  ` + "`" + `xml:"Order"` + "`" + `
		}

		type Order struct {
			XMLName      xml.Name ` + "`" + `xml:"Order"` + "`" + `
			CustomerID   string   ` + "`" + `xml:"CustomerID"` + "`" + `
			EmployeeID   int64    ` + "`" + `xml:"EmployeeID"` + "`" + `
			OrderDate    string   ` + "`" + `xml:"OrderDate"` + "`" + `
			ShipInfo     ShipInfo ` + "`" + `xml:"ShipInfo"` + "`" + `
			RequiredDate string   ` + "`" + `xml:"RequiredDate"` + "`" + `
		}

		type ShipInfo struct {
			XMLName        xml.Name ` + "`" + `xml:"ShipInfo"` + "`" + `
			Freight        float64  ` + "`" + `xml:"Freight"` + "`" + `
			ShipAddress    string   ` + "`" + `xml:"ShipAddress"` + "`" + `
			ShipCity       string   ` + "`" + `xml:"ShipCity"` + "`" + `
			ShipName       string   ` + "`" + `xml:"ShipName"` + "`" + `
			ShipPostalCode int64    ` + "`" + `xml:"ShipPostalCode"` + "`" + `
			ShipCountry    string   ` + "`" + `xml:"ShipCountry"` + "`" + `
			ShipVia        int64    ` + "`" + `xml:"ShipVia"` + "`" + `
			ShippedDate    string   ` + "`" + `xml:"ShippedDate,attr"` + "`" + `
			ShipRegion     string   ` + "`" + `xml:"ShipRegion"` + "`" + `
		}

		type Customers struct {
			XMLName  xml.Name   ` + "`" + `xml:"Customers"` + "`" + `
			Customer []Customer ` + "`" + `xml:"Customer"` + "`" + `
		}

		type Customer struct {
			XMLName      xml.Name    ` + "`" + `xml:"Customer"` + "`" + `
			CompanyName  string      ` + "`" + `xml:"CompanyName"` + "`" + `
			ContactName  string      ` + "`" + `xml:"ContactName"` + "`" + `
			ContactTitle string      ` + "`" + `xml:"ContactTitle"` + "`" + `
			CustomerID   string      ` + "`" + `xml:"CustomerID,attr"` + "`" + `
			Phone        string      ` + "`" + `xml:"Phone"` + "`" + `
			FullAddress  FullAddress ` + "`" + `xml:"FullAddress"` + "`" + `
		}

		type FullAddress struct {
			XMLName    xml.Name ` + "`" + `xml:"FullAddress"` + "`" + `
			Address    string   ` + "`" + `xml:"Address"` + "`" + `
			City       string   ` + "`" + `xml:"City"` + "`" + `
			Country    string   ` + "`" + `xml:"Country"` + "`" + `
			PostalCode int64    ` + "`" + `xml:"PostalCode"` + "`" + `
			Region     string   ` + "`" + `xml:"Region"` + "`" + `
		}`,

		`package model

		type Bookstore struct {
			Book      []Book
			Magazine  Magazine
			Specialty string ` + "`" + `xml:",attr"` + "`" + `
		}

		type Excerpt struct {
			DefinitionList DefinitionList
			P              []string
		}

		type DefinitionList struct {
			Definition string
			Term       string
		}

		type P struct {
			Emph string
		}

		type Price struct {
			Exchange float64 ` + "`" + `xml:",attr"` + "`" + `
			Intl     string  ` + "`" + `xml:",attr"` + "`" + `
		}

		type Degree struct {
			From string ` + "`" + `xml:",attr"` + "`" + `
		}

		type Magazine struct {
			Frequency    string ` + "`" + `xml:",attr"` + "`" + `
			Price        float64
			Style        string ` + "`" + `xml:",attr"` + "`" + `
			Subscription Subscription
		}

		type Subscription struct {
			Per   string ` + "`" + `xml:",attr"` + "`" + `
			Price int64  ` + "`" + `xml:",attr"` + "`" + `
		}

		type Editor struct {
			FirstName string
			LastName  string
		}

		type Book struct {
			Author Author
			Price  int64
			Style  string ` + "`" + `xml:",attr"` + "`" + `
		}

		type Author struct {
			Award     string
			FirstName string
			LastName  string
		}`,
	}
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
