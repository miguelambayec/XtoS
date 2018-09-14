package model

import "encoding/xml"

type Root struct {
	XMLName   xml.Name  `xml:"Root"`
	Customers Customers `xml:"Customers"`
	Orders    Orders    `xml:"Orders"`
	Xmlns     string    `xml:"xmlns,attr"`
}

type Orders struct {
	XMLName xml.Name `xml:"Orders"`
	Order   []Order  `xml:"Order"`
}

type Order struct {
	XMLName      xml.Name `xml:"Order"`
	CustomerID   string   `xml:"CustomerID"`
	EmployeeID   int64    `xml:"EmployeeID"`
	OrderDate    string   `xml:"OrderDate"`
	ShipInfo     ShipInfo `xml:"ShipInfo"`
	RequiredDate string   `xml:"RequiredDate"`
}

type ShipInfo struct {
	XMLName        xml.Name `xml:"ShipInfo"`
	Freight        float64  `xml:"Freight"`
	ShipAddress    string   `xml:"ShipAddress"`
	ShipCity       string   `xml:"ShipCity"`
	ShipName       string   `xml:"ShipName"`
	ShipPostalCode int64    `xml:"ShipPostalCode"`
	ShipCountry    string   `xml:"ShipCountry"`
	ShipVia        int64    `xml:"ShipVia"`
	ShippedDate    string   `xml:"ShippedDate,attr"`
	ShipRegion     string   `xml:"ShipRegion"`
}

type Customers struct {
	XMLName  xml.Name   `xml:"Customers"`
	Customer []Customer `xml:"Customer"`
}

type Customer struct {
	XMLName      xml.Name    `xml:"Customer"`
	CompanyName  string      `xml:"CompanyName"`
	ContactName  string      `xml:"ContactName"`
	ContactTitle string      `xml:"ContactTitle"`
	CustomerID   string      `xml:"CustomerID,attr"`
	Phone        string      `xml:"Phone"`
	FullAddress  FullAddress `xml:"FullAddress"`
}

type FullAddress struct {
	XMLName    xml.Name `xml:"FullAddress"`
	Address    string   `xml:"Address"`
	City       string   `xml:"City"`
	Country    string   `xml:"Country"`
	PostalCode int64    `xml:"PostalCode"`
	Region     string   `xml:"Region"`
}
