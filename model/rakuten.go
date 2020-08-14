package model

// TypeCategory holds the primary and secondary categories
type TypeCategory struct {
	Primary   string `xml:"primary"`
	Secondary string `xml:"secondary"`
}

// TypePrice holds the price tag and its attribute
type TypePrice struct {
	Price    float64 `xml:",chardata"`
	Currency string  `xml:"currency,attr"`
}

// TypeDescription holds the description tags
type TypeDescription struct {
	Short string `xml:"short"`
	Long  string `xml:"long"`
}

// CourseItem is the structure to hold each of the elements returned
type CourseItem struct {
	Mid          int             `xml:"mid"`
	MerchantName string          `xml:"merchantname"`
	Linkid       int             `xml:"linkid"`
	Createdon    string          `xml:"createdon"`
	Sku          string          `xml:"sku"`
	ProductName  string          `xml:"productname"`
	Categories   TypeCategory    `xml:"category"`
	Price        TypePrice       `xml:"price"`
	SalePrice    TypePrice       `xml:"saleprice"`
	Upcode       string          `xml:"upcode"`
	Description  TypeDescription `xml:"description"`
	Keywords     string          `xml:"keywords"`
	Imageurl     string          `xml:"imageurl"`
	Linkurl      string          `xml:"linkurl"`
}

// ErrorCall is a struct to hold api errors
type ErrorCall struct {
	ErrorID   int    `xml:"ErrorID"`
	ErrorText string `xml:"ErrorText"`
}

// ProductResult is the structure to hold the response
type ProductResult struct {
	TotalMatches int
	TotalPages   int
	PageNumber   int
	Items        []CourseItem `xml:"item"`
	Errors       ErrorCall    `xml:"Errors,omitempty"`
}
