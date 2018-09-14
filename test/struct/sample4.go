package model

type Bookstore struct {
	Book      []Book
	Magazine  Magazine
	Specialty string `xml:",attr"`
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
	Exchange float64 `xml:",attr"`
	Intl     string  `xml:",attr"`
}

type Degree struct {
	From string `xml:",attr"`
}

type Magazine struct {
	Frequency    string `xml:",attr"`
	Price        float64
	Style        string `xml:",attr"`
	Subscription Subscription
}

type Subscription struct {
	Per   string `xml:",attr"`
	Price int64  `xml:",attr"`
}

type Editor struct {
	FirstName string
	LastName  string
}

type Book struct {
	Author Author
	Price  int64
	Style  string `xml:",attr"`
}

type Author struct {
	Award     string
	FirstName string
	LastName  string
}
