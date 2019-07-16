package hotpepper

import (
	"os"
	"strconv"
	"testing"
)

func TestGourmetSearchError(t *testing.T) {
	api := New(os.Getenv("HOTPAPER_TOKEN"))
	_, err := api.GourmetSearch(WithAddress("東京"))
	if err != nil {
		t.Fatalf("Expected Error, got %v", err)
	}
}
func TestGourmetSearch(t *testing.T) {
	api := New(os.Getenv("HOTPAPER_TOKEN"))
	hotpaper, err := api.GourmetSearch(WithAddress("東京"))
	if err != nil {
		t.Fatal(err)
	}
	totalHit, _ := strconv.Atoi(hotpaper.Results_available)
	if totalHit <= 0 {
		t.Fatalf("Expected hit more than 100")
	}
}
