package model

type Catalog struct {
	Product Product
}

type Product struct {
	CatalogItem  []CatalogItem
	Description  string `xml:",attr"`
	ProductImage string `xml:",attr"`
}

type CatalogItem struct {
	Gender     string `xml:",attr"`
	ItemNumber string
	Price      float64
	Size       []Size
}

type Size struct {
	ColorSwatch []ColorSwatch
	Description string `xml:",attr"`
}

type ColorSwatch struct {
	Image string `xml:",attr"`
}
