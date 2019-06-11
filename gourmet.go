package hotpepper

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Hotpaper struct {
	Api_version       string `xml:"api_version"`
	Results_available string `xml:"results_available"`
	Results_returned  string `xml:"results_returned"`
	Results_start     string `xml:"results_start"`
	Shops             []Shop `xml:"shop"`
}
type Shop struct {
	Id           string `xml:"id"`
	Name         string `xml:"name"`
	Logo_image   string `xml: "logo_image"`
	Name_kana    string `xml: "name_kana"`
	Address      string `xml: "address"`
	Station_name string `xml: "station_name"`
	Ktai_coupon  string `xml: "ktai_conpon"`
}

func (cli *Client) GourmetSearch() error {
	req, err := http.NewRequest("GET", cli.baseURL, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	fmt.Println(cli.token)
	q.Add("key", cli.token)
	q.Add("large_area", "Z011")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	a, _ := http.Get(req.URL.String())
	defer a.Body.Close()
	fmt.Println(a.Body)
	byteArray, _ := ioutil.ReadAll(a.Body)
	//fmt.Println(string(byteArray))
	data := new(Hotpaper)
	err = xml.Unmarshal(byteArray, data)
	fmt.Println(data.Shops)
	return nil
}
