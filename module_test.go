// Package module unit test
// @author Elhussein Ali <hussein.ali@sadeem.io>
package main

import (
	"fmt"
	"testing"

	"github.com/Ahmah2009/myhttp/lib"
)

func TestModule(t *testing.T) {
	t.Log("testing myhttp module started")

	// tes url with https scheme
	url, err := lib.ParseURL("https://google.com")
	assertNotError(t, err, "")
	assertEqual(t, url, "https://google.com", "invalid url parsing")

	// test add scheme to url
	url, err = lib.ParseURL("http://google.com")
	assertNotError(t, err, "")
	assertEqual(t, url, "http://google.com", "invalid url parsing")

}

func assertError(t *testing.T, a interface{}, message string) {
	switch a.(type) {
	case error:
		return
	case nil:
		t.Fatal("error is expected here,", message) // here v has type int
	}
}

func assertNotError(t *testing.T, a interface{}, message string) {
	switch a.(type) {
	case error:
		t.Fatal("un-expected error,", message) // here v has type int
	case nil:
		return
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
