package hotpepper

// Shop Name search api query struct
type ShopSearch struct {
	Keyword string `xml:"keyword"`
	Tel     string `xml:"tel"`
	Start   int64  `xml:"start"`
	Count   int64  `xml: "count"`
	Format  string `xml: "format"`
}

type ShopName struct {
	Results           string `xml:"results"`
	Api_version       string `xml:"api_version"`
	Results_available string `xml:"results_available"`
	Results_returned  string `xml:"results_returned"`
	Results_start     string `xml:"results_start"`
	Shops             []Shop `xml:"Shop"`
}

type ShopSearchShop struct {
	Id        string `xml:"id"`
	Name      string `xml:"name"`
	Name_kana string `xml:"name_kana"`
	Address   string `xml:"address"`
	Desc      int64  `xml:"desc"`
}
