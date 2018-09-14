package model

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Body    Body     `xml:"soapenv:Body"`
	Header  Header   `xml:"soapenv:Header"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
}

type Body struct {
	XMLName               xml.Name              `xml:"soapenv:Body"`
	GetAdUnitsByStatement GetAdUnitsByStatement `xml:"getAdUnitsByStatement"`
}

type GetAdUnitsByStatement struct {
	XMLName         xml.Name        `xml:"getAdUnitsByStatement"`
	FilterStatement FilterStatement `xml:"filterStatement"`
	Xmlns           string          `xml:"xmlns,attr"`
}

type FilterStatement struct {
	XMLName xml.Name `xml:"filterStatement"`
	Query   string   `xml:"query"`
}

type Header struct {
	XMLName       xml.Name      `xml:"soapenv:Header"`
	RequestHeader RequestHeader `xml:"ns1:RequestHeader"`
}

type RequestHeader struct {
	XMLName         xml.Name `xml:"ns1:RequestHeader"`
	Actor           string   `xml:"soapenv:actor,attr"`
	ApplicationName string   `xml:"ns1:applicationName"`
	MustUnderstand  int64    `xml:"soapenv:mustUnderstand,attr"`
	NetworkCode     int64    `xml:"ns1:networkCode"`
	Ns1             string   `xml:"xmlns:ns1,attr"`
}
