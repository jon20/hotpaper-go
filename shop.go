package hotpepper

// Shop Name search api query struct
type ShopSearch struct {
	Keyword string `xml:"keyword"`
	Tel     string `xml:"tel"`
	Start   int64  `xml:"start"`
	Count   int64  `xml: "count"`
	Format  string `xml: "format"`
}
