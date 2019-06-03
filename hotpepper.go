package hotpepper

const (
	APIURL = ""
)

type Client struct {
	token string
	host  string
}

func New(token string) *Client {
	s := &Client{
		token: token,
	}
	return s
}
