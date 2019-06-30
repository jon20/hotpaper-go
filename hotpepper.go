package hotpepper

const (
	APIURL = "http://webservice.recruit.co.jp/hotpepper/gourmet/v1/"
)

type Client struct {
	token   string
	baseURL string
}

func New(token string) *Client {
	s := &Client{
		token:   token,
		baseURL: APIURL,
	}
	return s
}
