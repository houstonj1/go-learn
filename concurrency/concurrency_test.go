package concurrency

import (
	"reflect"
	"testing"
)

/* Existing code */
// mockWebsitesChecker function
func mockWebsitesChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

// TestCheckWebsites function
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsitesChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

/* End existing code */
