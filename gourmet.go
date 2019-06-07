package hotpepper

import (
	"fmt"

	"github.com/dghubble/sling"
)

func (cli *Client) GourmetSearch() {
	sling.New().Get(cli.host)
	fmt.Println("a")
}
