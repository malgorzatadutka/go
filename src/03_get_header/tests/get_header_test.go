package tests

import (
	f "03_get_header/functions"
	"testing"
)

func TestGetHeader(t *testing.T) {

	correctHeader, _ := f.GetHeader("https://onet.pl", "Access-Control-Max-Age")
	recievedHeader := "60"

	if correctHeader != recievedHeader {
		t.Errorf("Correct header: %q, Recieved header: %q", correctHeader, recievedHeader)
	}
}
