package hotpepper

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
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
	ID                int      `param:"id"`
	Name              string   `param:"name"`
	NameKana          string   `param:"name_kana"`
	NameAny           string   `param:"name_any"`
	Tel               string   `param:"tel"`
	Address           string   `param:"address"`
	Special           string   `param:"special"`
	SpecialOR         string   `param:"special_or"`
	SpecialCategory   string   `param:"special_category`
	SpecialCategoryOR string   `param:"special_category_or`
	LargeServiceArea  string   `param:"large_service_area"`
	ServiceArea       string   `param:"service_area"`
	LargeArea         string   `param:"large_area"`
	MiddleArea        string   `param:"middle_area"`
	SmallArea         string   `param:"small_area"`
	Keyword           string   `param:"keyword"`
	Lat               float64  `param:"lat"`
	Lng               float64  `param:"lng"`
	Range             int      `param:"range"`
	Datum             string   `param:"datum"`
	KtaiCoupon        int      `param:"ktai_coupon"`
	Genre             int      `param:"genre"`
	Budget            string   `param:"budget"`
	PartyCapacity     int      `param:"party_capacity"`
	Wifi              int      `param:"wifi"`
	Wedding           int      `param:"wedding"`
	Course            int      `param:"course"`
	Free_drink        int      `param:"free_drink"`
	Free_food         int      `param:"free_food"`
	Private_room      int      `param:"private_room"`
	Horigotatsu       int      `param:"horigotatsu"`
	Tatami            int      `param:"tatami"`
	Cocktail          int      `param:"cocktail"`
	Shochu            int      `param:"shochu"`
	Sake              int      `param:"sake"`
	Wine              int      `param:"wine"`
	Card              int      `param:"card"`
	Non_smoking       int      `param:"non_smoking"`
	Charter           int      `param:"charter"`
	Ktai              int      `param:"ktai"`
	Parking           int      `param:"parking"`
	Barrier_free      int      `param:"barrier_free"`
	Sommelier         int      `param:"sommelier"`
	Night_view        int      `param:"night_view"`
	Open_air          int      `param:"open_air"`
	Show              int      `param:"show"`
	Equipment         int      `param:"equipment"`
	Karaoke           int      `param:"karaoke"`
	Band              int      `param:"band"`
	Tv                int      `param:"tv"`
	Lunch             int      `param:"lunch"`
	Midnight          int      `param:"midnight"`
	Midnight_meal     int      `param:"midnight_meal"`
	English           int      `param:"english"`
	Pet               int      `param:"pet"`
	Child             int      `param:"child"`
	Credit_card       []string `param:"credit_card"`
	Type              string   `param:"type"`
	Order             int      `param:"order"`
	Start             int      `param:"start"`
	Count             int      `param:"count"`
	Format            string   `param:"format"`
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
func (cli Client) newGourmetRequest(gourmet *Gourmet) (*Hotpaper, error) {
	req, err := http.NewRequest("GET", cli.baseURL, nil)
	if err != nil {
		err = errors.New("aa")
		return nil, err
	}
	q := req.URL.Query()
	//v := reflect.Indirect(reflect.ValueOf(gourmet))
	//t := v.Type()
	//t := reflect.ValueOf(gourmet).Elem()
	//typeOfT := t.Type()

	q.Add("key", cli.token)
	rt, rv := reflect.TypeOf(*gourmet), reflect.ValueOf(*gourmet)
	for i := 0; i < rt.NumField(); i++ {
		value := rv.Field(i).Interface()
		f := rt.Field(i)
		switch v := value.(type) {
		case int:
			if v == 0 {
				continue
			}
			param := strconv.Itoa(value.(int))
			q.Add(f.Tag.Get("param"), param)
		case string:
			if v == "" || value == "0" {
				continue
			}
			param := rv.Field(i).String()
			q.Add(f.Tag.Get("param"), param)
		case float64:
			if v == 0.0 {
				continue
			}
			param := value.(float64)
			convstr := fmt.Sprintf("%f", param)
			q.Add(f.Tag.Get("param"), convstr)
		case []string:
			if len(v) == 0 {
				continue
			}
		}
	}

	req.URL.RawQuery = q.Encode()
	response, err := http.Get(req.URL.String())
	if err != nil {
		err = errors.New(ErrorConnectionFailed)
		return nil, err
	}
	defer response.Body.Close()
	byteArray, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := new(Hotpaper)
	err = xml.Unmarshal(byteArray, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GourmetSearch will returns specified list
func (cli Client) GourmetSearch(opts ...Option) (*Hotpaper, error) {
	c := new(Gourmet)
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	response, err := cli.newGourmetRequest(c)
	if err != nil {
		return nil, err
	}
	return response, nil

}
