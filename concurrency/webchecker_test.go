package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"waat://furhurterwe.geds",
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	want := map[string]bool{
		"https://google.com":      true,
		"https://facebook.com":    true,
		"waat://furhurterwe.geds": false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v wanted %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)
	for i := range 100 {
		urls[i] = "a url"
	}
	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
