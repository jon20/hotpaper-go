package hotpepper

import (
	"os"
	"testing"
)

func TestGourmetSearch(t *testing.T) {
	api := New(os.Getenv("HOTPAPER_TOKEN"))
	api.GourmetSearch(WithID(2), WithAddress("sample"), WithLat(0.1))
	t.Fatal("aa")
}
