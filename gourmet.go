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
	Id                 string             `xml:"id"`
	Name               string             `xml:"name"`
	Logo_image         string             `xml:"logo_image"`
	Name_kana          string             `xml:"name_kana"`
	Address            string             `xml:"address"`
	Station_name       string             `xml:"station_name"`
	Ktai_coupon        int                `xml:"ktai_conpon"`
	Large_service_area Large_service_area `xml:"large_service_area"`
	Service_area       Service_area       `xml:"service_area"`
	Middle_area        Middle_area        `xml:"middle_area"`
	Lat                float64            `xml:"lat"`
	Lng                float64            `xml:"lng"`
	Genre              Genre              `xml:"genre"`
	Sub_Genre          Sub_Genre          `xml:"sub_genre"`
	Budget             Budget             `xml:"budget"`
	Budget_memo        string             `xml:"budget_memo"`
	Catch              string             `xml:"catch"`
	Capacity           string             `xml:"capacity"`
	Access             string             `xml:"access"`
	Mobile_access      string             `xml:"mobile_access"`
	Urls               Urls               `xml:"urls"`
	Photo              Photo              `xml:"photo"`
	Open               string             `xml:"open"`
	Close              string             `xml:"close"`
	Party_capacity     string             `xml:"party_capacity"`
	Wifi               string             `xml:"wifi"`
	Wedding            string             `xml:"wedding"`
	Course             string             `xml:"course"`
	Free_drink         string             `xml:"free_drink"`
	Free_food          string             `xml:"free_food"`
	Private_room       string             `xml:"private_room"`
	Horigotatsu        string             `xml:"horigotatsu"`
	Tatami             string             `xml:"tatami"`
	Card               string             `xml:"card"`
	Non_smoking        string             `xml:"non_smoking"`
	Charter            string             `xml:"charter"`
	Ktai               string             `xml:"ktai"`
	Parking            string             `xml:"parking"`
	Barrier_free       string             `xml:"barrier_free"`
	Other_memo         string             `xml:"other_memo"`
	Sommelier          string             `xml:"sommelier"`
	Open_air           string             `xml:"open_air"`
	Show               string             `xml:"show"`
	Equipment          string             `xml:"equipment"`
	Karaoke            string             `xml:"karaoke"`
	Band               string             `xml:"band"`
	Tv                 string             `xml:"tv"`
	English            string             `xml:"english`
	Pet                string             `xml:"pet"`
	Child              string             `xml:"child"`
	Lunch              string             `xml:"lunch"`
	Midnight           string             `xml:"midnight"`
	Shop_detail_memo   string             `xml:"shop_detail_memo"`
	Coupon_urls        Coupon_urls        `xml:"coupon_urls"`
}
type Large_service_area struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}
type Service_area struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}

type Middle_area struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}
type Genre struct {
	Code  string `xml:"code"`
	Name  string `xml:"name"`
	Catch string `xml:"catch"`
}
type Sub_Genre struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}

type Budget struct {
	Code    string `xml:"code"`
	Name    string `xml:"name"`
	Average string `xml:"average"`
}

type Urls struct {
	Pc string `xml:"pc"`
}

type Photo struct {
	Pc     Pc     `xml:"pc"`
	Mobile Mobile `xml:"mobile"`
}
type Pc struct {
	L string `xml:"l"`
	M string `xml:"m"`
	S string `xml:"s"`
}
type Mobile struct {
	L string `xml:"l"`
	S string `xml:"s"`
}
type Coupon_urls struct {
	Pc string `xml:"pc"`
	Sp string `xml:"sp"`
}

type Option func(*Gourmet) error

type Gourmet struct {
	ID                int
	Name              string
	NameKana          string
	NameAny           string
	Tel               string
	Address           string
	Special           string
	SpecialOR         string
	SpecialCategory   string
	SpecialCategoryOR string
	LargeServiceArea  string
	ServiceArea       string
	LargeArea         string
	MiddleArea        string
	SmallArea         string
	Keyword           string
	Lat               float64
	Lng               float64
	Range             int
	Datum             string
	KtaiCoupon        int
	Genre             int
	Budget            string
	PartyCapacity     int
	Wifi              int
	Wedding           int
	Course            int
	Free_drink        int
	Free_food         int
	Private_room      int
	Horigotatsu       int
	Tatami            int
	Cocktail          int
	Shochu            int
	Sake              int
	Wine              int
	Card              int
	Non_smoking       int
	Charter           int
	Ktai              int
	Parking           int
	Barrier_free      int
	Sommelier         int
	Night_view        int
	Open_air          int
	Show              int
	Equipment         int
	Karaoke           int
	Band              int
	Tv                int
	Lunch             int
	Midnight          int
	Midnight_meal     int
	English           int
	Pet               int
	Child             int
	Credit_card       []string
	Type              string
	Order             int
	Start             int
	Count             int
	Format            string
}

func WithID(id int) Option {
	return func(g *Gourmet) error {
		g.ID = id
		return nil
	}
}

func WithName(name string) Option {
	return func(g *Gourmet) error {
		g.Name = name
		return nil
	}
}

func WithNamekana(namekana string) Option {
	return func(g *Gourmet) error {
		g.NameKana = namekana
		return nil
	}
}

func WithNameAny(nameany string) Option {
	return func(g *Gourmet) error {
		g.NameAny = nameany
		return nil
	}
}
func WithTel(tel string) Option {
	return func(g *Gourmet) error {
		g.Tel = tel
		return nil
	}
}
func WithAddress(address string) Option {
	return func(g *Gourmet) error {
		g.Address = address
		return nil
	}
}
func WithSpacial(special string) Option {
	return func(g *Gourmet) error {
		g.Special = special
		return nil
	}
}
func WithSpacialOR(specialOR string) Option {
	return func(g *Gourmet) error {
		g.SpecialOR = specialOR
		return nil
	}
}
func WithSpacialCategory(specialCategory string) Option {
	return func(g *Gourmet) error {
		g.SpecialCategory = specialCategory
		return nil
	}
}
func WithSpacialCategoryOR(specialCategoryOR string) Option {
	return func(g *Gourmet) error {
		g.SpecialCategoryOR = specialCategoryOR
		return nil
	}
}
func WithLargeServiceArea(largeServiceArea string) Option {
	return func(g *Gourmet) error {
		g.LargeServiceArea = largeServiceArea
		return nil
	}
}
func WithServiceArea(serviceArea string) Option {
	return func(g *Gourmet) error {
		g.ServiceArea = serviceArea
		return nil
	}
}
func WithLargeArea(largeArea string) Option {
	return func(g *Gourmet) error {
		g.LargeArea = largeArea
		return nil
	}
}
func WithMiddleArea(middleArea string) Option {
	return func(g *Gourmet) error {
		g.MiddleArea = middleArea
		return nil
	}
}
func WithSmallArea(smallArea string) Option {
	return func(g *Gourmet) error {
		g.SmallArea = smallArea
		return nil
	}
}
func WithKeyword(keyword string) Option {
	return func(g *Gourmet) error {
		g.Keyword = keyword
		return nil
	}
}

func WithLat(lat float64) Option {
	return func(g *Gourmet) error {
		g.Lat = lat
		return nil
	}
}
func WithLng(lng float64) Option {
	return func(g *Gourmet) error {
		g.Lng = lng
		return nil
	}
}
func WithRange(withRange int) Option {
	return func(g *Gourmet) error {
		g.Range = withRange
		return nil
	}
}
func WithDatum(datum string) Option {
	return func(g *Gourmet) error {
		g.Datum = datum
		return nil
	}
}
func WithKtaiCoupon(ktaiCoupon int) Option {
	return func(g *Gourmet) error {
		g.KtaiCoupon = ktaiCoupon
		return nil
	}
}

func WithGenre(genre int) Option {
	return func(g *Gourmet) error {
		g.Genre = genre
		return nil
	}
}
func WithBudget(budget string) Option {
	return func(g *Gourmet) error {
		g.Budget = budget
		return nil
	}
}
func WithPartyCapacity(partyCapacity int) Option {
	return func(g *Gourmet) error {
		g.PartyCapacity = partyCapacity
		return nil
	}
}

func WithWifi(wifi int) Option {
	return func(g *Gourmet) error {
		g.Wifi = wifi
		return nil
	}
}
func WithWedding(wedding int) Option {
	return func(g *Gourmet) error {
		g.Wedding = wedding
		return nil
	}
}
func WithCourse(course int) Option {
	return func(g *Gourmet) error {
		g.Course = course
		return nil
	}
}
func WithFreeDrink(freeDrink int) Option {
	return func(g *Gourmet) error {
		g.Free_drink = freeDrink
		return nil
	}
}
func WithPrivateRoom(privateRoom int) Option {
	return func(g *Gourmet) error {
		g.Private_room = privateRoom
		return nil
	}
}
func WithHirigotatsu(horigotatsu int) Option {
	return func(g *Gourmet) error {
		g.Horigotatsu = horigotatsu
		return nil
	}
}
func WithTatami(tatami int) Option {
	return func(g *Gourmet) error {
		g.Tatami = tatami
		return nil
	}
}
func WithCocktail(coktail int) Option {
	return func(g *Gourmet) error {
		g.Cocktail = coktail
		return nil
	}
}
func WithShochu(shochu int) Option {
	return func(g *Gourmet) error {
		g.Shochu = shochu
		return nil
	}
}
func WithSake(sake int) Option {
	return func(g *Gourmet) error {
		g.Sake = sake
		return nil
	}
}
func WithWine(wine int) Option {
	return func(g *Gourmet) error {
		g.Wine = wine
		return nil
	}
}
func WithCard(card int) Option {
	return func(g *Gourmet) error {
		g.Card = card
		return nil
	}
}

func WithNonSmoking(nonSmoking int) Option {
	return func(g *Gourmet) error {
		g.Non_smoking = nonSmoking
		return nil
	}
}

func WithCharter(charter int) Option {
	return func(g *Gourmet) error {
		g.Charter = charter
		return nil
	}
}
func WithKtai(ktai int) Option {
	return func(g *Gourmet) error {
		g.Ktai = ktai
		return nil
	}
}
func WithParking(parking int) Option {
	return func(g *Gourmet) error {
		g.Parking = parking
		return nil
	}
}
func WithBarrierFree(barrierfree int) Option {
	return func(g *Gourmet) error {
		g.Barrier_free = barrierfree
		return nil
	}
}

func WithSommelier(sommelier int) Option {
	return func(g *Gourmet) error {
		g.Sommelier = sommelier
		return nil
	}
}
func WithNightView(nightView int) Option {
	return func(g *Gourmet) error {
		g.Night_view = nightView
		return nil
	}
}
func WithOpenAir(openAir int) Option {
	return func(g *Gourmet) error {
		g.Open_air = openAir
		return nil
	}
}
func WithShow(show int) Option {
	return func(g *Gourmet) error {
		g.Show = show
		return nil
	}
}
func WithEquipment(equipment int) Option {
	return func(g *Gourmet) error {
		g.Equipment = equipment
		return nil
	}
}
func WithKaraoke(karaoke int) Option {
	return func(g *Gourmet) error {
		g.Karaoke = karaoke
		return nil
	}
}

func WithBand(band int) Option {
	return func(g *Gourmet) error {
		g.Band = band
		return nil
	}
}
func WithTV(tv int) Option {
	return func(g *Gourmet) error {
		g.Tv = tv
		return nil
	}
}
func WithLunch(lunch int) Option {
	return func(g *Gourmet) error {
		g.Lunch = lunch
		return nil
	}
}
func WithMidnight(midnight int) Option {
	return func(g *Gourmet) error {
		g.Midnight = midnight
		return nil
	}
}
func WithEnglish(english int) Option {
	return func(g *Gourmet) error {
		g.English = english
		return nil
	}
}
func WithPet(pet int) Option {
	return func(g *Gourmet) error {
		g.Pet = pet
		return nil
	}
}
func WithChild(child int) Option {
	return func(g *Gourmet) error {
		g.Child = child
		return nil
	}
}
func WithCreditCard(creditCard []string) Option {
	return func(g *Gourmet) error {
		g.Credit_card = creditCard
		return nil

	}
}
func WithType(withtype string) Option {
	return func(g *Gourmet) error {
		g.Type = withtype
		return nil
	}
}
func WithOrder(order int) Option {
	return func(g *Gourmet) error {
		g.Order = order
		return nil
	}
}
func WithStart(start int) Option {
	return func(g *Gourmet) error {
		g.Start = start
		return nil
	}
}
func WithCount(count int) Option {
	return func(g *Gourmet) error {
		g.Count = count
		return nil
	}
}
func WithFormat(format string) Option {
	return func(g *Gourmet) error {
		g.Format = format
		return nil
	}
}

func (cli Client) GourmetSearch(opts ...Option) error {
	c := new(Gourmet)
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	fmt.Println(c.ID)
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
	for _, item := range data.Shops {
		fmt.Println(item.Coupon_urls)
	}
	return nil

}
